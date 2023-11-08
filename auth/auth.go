package auth

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"regexp"
	"rest-go-demo/database"
	"rest-go-demo/entity"
	"rest-go-demo/wc_analytics"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spruceid/siwe-go"

	//"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"github.com/haltingstate/secp256k1-go"

	delegatecash "rest-go-demo/contracts" // for demo

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	_ "rest-go-demo/docs"

	"github.com/0xsequence/go-sequence/api"

	"blockwatch.cc/tzgo/tezos"
)

var (
	ErrUserNotExists  = errors.New("Authuser does not exist")
	ErrUserExists     = errors.New("Authuser already exists")
	ErrInvalidAddress = errors.New("invalid address")
	ErrInvalidNonce   = errors.New("invalid nonce")
	ErrInvalidSIWE    = errors.New("invalid format of SIWE message")
	ErrInvalidDomain  = errors.New("invalid domain in SIWE message")
	ErrMissingSig     = errors.New("signature is missing")
	ErrAuthError      = errors.New("authentication error")
)

type JwtHmacProvider struct {
	hmacSecret []byte
	issuer     string
	duration   time.Duration
}

func NewJwtHmacProvider(hmacSecret string, issuer string, duration time.Duration) *JwtHmacProvider {
	ans := JwtHmacProvider{
		hmacSecret: []byte(hmacSecret),
		issuer:     issuer,
		duration:   duration,
	}
	return &ans
}

func (j *JwtHmacProvider) CreateStandard(subject string) (string, error) {
	now := time.Now()
	claims := jwt.RegisteredClaims{
		Issuer:    j.issuer,
		Subject:   subject,
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(j.duration)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.hmacSecret)
}

func (j *JwtHmacProvider) Verify(tokenString string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.hmacSecret, nil
	})
	if err != nil {
		return nil, ErrAuthError
	}
	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrAuthError
}

type Authuser struct {
	Address string
	Nonce   string
}

func CreateIfNotExists(u Authuser) error {
	var checkUser Authuser
	dbQuery := database.Connector.Where("address = ?", u.Address).Find(&checkUser)

	if dbQuery.RowsAffected > 0 {
		return ErrUserExists
	}

	//create the item in the database
	dbQuery = database.Connector.Create(&u)
	if dbQuery.RowsAffected == 0 {
		fmt.Println("error creating user / nonce for user: ", u.Address)
	}
	return nil
}

func Get(address string) (Authuser, error) {
	var checkUser Authuser
	dbQuery := database.Connector.Where("address = ?", address).Find(&checkUser)

	if dbQuery.RowsAffected == 0 {
		return checkUser, ErrUserNotExists
	}

	return checkUser, nil
}

func Update(user Authuser) error {

	database.Connector.Model(&Authuser{}).
		Where("address = ?", user.Address).
		Update("nonce", user.Nonce)

	return nil
}

// ============================================================================

var (
	//TODO - might need to pass in chain type for
	//hexRegexEVM   *regexp.Regexp = regexp.MustCompile(`^0x[a-fA-F0-9]{40}$`)
	nonceRegex *regexp.Regexp = regexp.MustCompile(`^[0-9]+$`)
)

type RegisterPayload struct {
	Address string `json:"address"`
}

func (p RegisterPayload) Validate() error {
	// if !hexRegexEVM.MatchString(p.Address) {
	// 	return ErrInvalidAddress
	// }
	return nil
}

//Legacy - not needed anymore
func RegisterHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestBody, _ := ioutil.ReadAll(r.Body)
		var p RegisterPayload
		if err := json.Unmarshal(requestBody, &p); err != nil { // Parse []byte to the go struct pointer
			fmt.Println("Can not unmarshal JSON in RegisterHandler")
		}
		if err := p.Validate(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		nonce, err := GetNonce()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		u := Authuser{
			Address: strings.ToLower(p.Address), // let's only store lower case
			Nonce:   nonce,
		}
		if err := CreateIfNotExists(u); err != nil {
			switch errors.Is(err, ErrUserExists) {
			case true:
				w.WriteHeader(http.StatusConflict)
			default:
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

// UserNonceHandler godoc
// @Summary If the current wallet doesn't have a valid local JWT, need to request a new nonce to sign
// @Description As part of the login process, we need a user to sign a nonce genrated from the API, to prove the user in fact
// @Description the owner of the wallet they are siging in from.  JWT currently set to 24 hour validity (could change this upon request)
// @Tags Auth
// @Accept  json
// @Produce json
// @Param address path string true "wallet address to get nonce to sign"
// @Success 200 {} Authuser
// @Router /users/{address}/nonce [get]
func UserNonceHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		address := vars["address"]

		// if !hexRegexEVM.MatchString(address) {
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	return
		// }

		if !strings.HasPrefix(address, "tz") && !strings.HasPrefix(address, "SP") { //Tezos accounts are case senstive
			address = strings.ToLower(address)
		}

		//combining /register and /users (no need to call both and check each time)
		nonce, err := GetNonce()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		user := Authuser{
			Address: address, // let's only store lower case
			Nonce:   nonce,
		}
		CreateIfNotExists(user)
		//end of copied /register functionality

		Authuser, err := Get(address)
		if err != nil {
			switch errors.Is(err, ErrUserNotExists) {
			case true:
				w.WriteHeader(http.StatusNotFound)
			default:
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
		resp := struct {
			Nonce string
		}{
			Nonce: Authuser.Nonce,
		}
		renderJson(r, w, http.StatusOK, resp)
	}
}

type SigninPayload struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Nonce   string `json:"nonce"`
	Sig     string `json:"sig"`
	Msg     string `json:"msg"`
}

func (s SigninPayload) Validate() error {
	var message *siwe.Message
	var err error

	message, err = siwe.ParseMessage(s.Msg)
	if err != nil {
		return ErrInvalidSIWE
	}

	domainMatch := false
	domainList := strings.Split(os.Getenv("ALLOWED_DOMAINS"), ",")
	for _, domain := range domainList {
		if strings.HasSuffix(message.GetDomain(), domain) {
			domainMatch = true
			fmt.Println("Signin From Domain: ", domain, message.GetAddress())
			break
		}
	}
	if !domainMatch {
		fmt.Println("Unauthorized Domain: ", message.GetDomain())
		return ErrInvalidDomain
	}

	if !nonceRegex.MatchString(s.Nonce) {
		fmt.Println("Invalid Nonce: ", s.Nonce)
		return ErrInvalidNonce
	}
	if len(s.Sig) == 0 {
		fmt.Println("Missing Sig")
		return ErrMissingSig
	}
	return nil
}

// SigninHandler godoc
// @Summary Sign In with signed nonce value, currently JWT token returned should be valid for 24 hours
// @Description Every call the to API after this signin should present the JWT Bearer token for authenticated access.
// @Description Upon request we can change the timeout to greater than 24 hours, or setup an addtional dedicated API for
// @Description an agreed upon development and maintenance cost
// @Tags Auth
// @Accept  json
// @Produce json
// @Param message body SigninPayload true "json input containing signed message and append nonce for easy processing"
// @Success 200 {integer} int
// @Router /signin [post]
func SigninHandler(jwtProvider *JwtHmacProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p SigninPayload
		requestBody, _ := ioutil.ReadAll(r.Body)
		if err := json.Unmarshal(requestBody, &p); err != nil { // Parse []byte to the go struct pointer
			fmt.Println("Can not unmarshal JSON in SigninHandler")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := p.Validate(); err != nil {
			fmt.Println("Some fields were invalid")
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		//Tezos addresses are case sensitive (here address is the full public key)
		address := p.Address
		if !strings.HasPrefix(p.Address, "edpk") {
			address = strings.ToLower(p.Address)
		}

		Authuser, err := Authenticate(p.Name, address, p.Nonce, p.Msg, p.Sig)
		switch err {
		case nil:
		case ErrAuthError:
			fmt.Println("Auth Error: ", address, err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if strings.HasPrefix(Authuser.Address, "0x") {
			//Not sure yet if this is the best way, but lets try it
			delegates := GetDelegationsByDelegate(Authuser.Address)
			if delegates != nil {
				//fmt.Println("Wallet Delegates in Auth: ", delegates)

				for _, delegateWallet := range delegates {
					if delegateWallet.Type == 1 {
						fmt.Println("Wallet Full Delegate Authorized: ", delegateWallet.Vault.Hex())
						Authuser.Address = delegateWallet.Vault.Hex()
						break
					}
				}
			}
		}

		wc_analytics.SendCustomEvent(Authuser.Address, "CONNECT_WALLET_SIGNIN")

		signedToken, err := jwtProvider.CreateStandard(Authuser.Address)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:  "Authorization",
			Value: signedToken,
			// Path:     "/",
			// SameSite: 4,
			// Secure:   true,
			// MaxAge:   86400,
			// true means no scripts, http requests only. This has
			// nothing to do with https vs http
			HttpOnly: true,
		})
		resp := struct {
			AccessToken string `json:"access"`
		}{
			AccessToken: signedToken,
		}
		renderJson(r, w, http.StatusOK, resp)
		// renderJsonWithCookie(r, w, http.StatusOK, http.Cookie{
		// 	Name:  "jwt",
		// 	Value: signedToken,
		// 	// true means no scripts, http requests only. This has
		// 	// nothing to do with https vs http
		// 	HttpOnly: true,
		// }, resp)
	}
}

func WelcomeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		Authuser := GetUserFromReqContext(r)
		//fmt.Println("getting Authuser: ", Authuser)
		resp := struct {
			Msg string `json:"msg"`
		}{
			Msg: "Congrats " + Authuser.Address + " you made it",
		}
		var addrnameDB entity.Addrnameitem
		var dbQuery = database.Connector.Where("address = ?", Authuser.Address).Find(&addrnameDB)
		if dbQuery.RowsAffected > 0 {
			resp.Msg = "Welcome:" + addrnameDB.Name + ":Addr:" + Authuser.Address
			wc_analytics.SendCustomEvent(Authuser.Address, "CONNECT_WALLET_VALID_JWT")
		}

		renderJson(r, w, http.StatusOK, resp)
	}
}

// ============================================================================

func GetUserFromReqContext(r *http.Request) Authuser {
	ctx := r.Context()
	key := ctx.Value("Authuser").(Authuser)
	return key
}

var apiTrackerCnt = make(map[string]int32)

func AuthMiddleware(jwtProvider *JwtHmacProvider) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			headerValue := r.Header.Get("Authorization")
			if len(headerValue) > 0 {
				const prefix = "Bearer "
				if len(headerValue) < len(prefix) {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
				//fmt.Println("Found JWT in Authorization header")
				headerValue = headerValue[len(prefix):]
			} else {
				tokenCookie, err := r.Cookie("Authorization")
				if err != nil {
					//log.Fatalf("Error occured while reading cookie")
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
				fmt.Println("Found JWT in Cookie")
				headerValue = tokenCookie.Value
			}
			// fmt.Println("Authorization: ", headerValue)
			// fmt.Println("headerValue: ", r.Header)

			tokenString := headerValue //headerValue[len(prefix):]
			if len(tokenString) < 32 {
				//fmt.Println("auth.go - Invalid Bearer Token - too short")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			if strings.Contains(os.Getenv("ADMIN_API_KEY_LIST"), tokenString) {
				var authAdmin Authuser
				authAdmin.Address = tokenString[0:16]
				authAdmin.Nonce = "none"

				//count POST requests per user
				if r.Method == "POST" {
					SHORT_API_ID := authAdmin.Address //lets only grab the first 16 so any log doesn't contain full API keys
					val, ok := apiTrackerCnt[SHORT_API_ID]
					// If the key exists
					if ok {
						apiTrackerCnt[SHORT_API_ID] = val + 1
					} else {
						apiTrackerCnt[SHORT_API_ID] = 1
					}

					//fmt.Println("POST Count Update for Addr: ", SHORT_API_ID, apiTrackerCnt[SHORT_API_ID])
					wc_analytics.SendCustomEvent(authAdmin.Address, "POST_COUNT")

				}

				ctx := context.WithValue(r.Context(), "Authuser", authAdmin)
				next.ServeHTTP(w, r.WithContext(ctx))

				wc_analytics.SendCustomEvent(authAdmin.Address, "ADMIN_API_AUTH")
			}
			claims, err := jwtProvider.Verify(tokenString)
			if err != nil {
				//w.WriteHeader(http.StatusUnauthorized)
				return
			}

			Authuser, err := Get(claims.Subject)
			if err != nil {
				if errors.Is(err, ErrUserNotExists) {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			//count POST requests per user
			if r.Method == "POST" {
				apiTrackerCnt[Authuser.Address] = apiTrackerCnt[Authuser.Address] + 1
				//fmt.Println("POST Count Update for Addr: ", Authuser.Address, apiTrackerCnt[Authuser.Address])
				wc_analytics.SendCustomEvent(Authuser.Address, "POST_COUNT")
			}

			ctx := context.WithValue(r.Context(), "Authuser", Authuser)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetCountsAPI() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		apiKey := r.Header.Get("Authorization")
		if len(apiKey) > 0 {
			const prefix = "Bearer "
			if len(apiKey) < len(prefix) {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			apiKey = apiKey[len(prefix):]
		}

		if strings.Contains(os.Getenv("WALLETCHAT_KEY_LIST"), apiKey) {
			renderJson(r, w, http.StatusOK, apiTrackerCnt)
		}
	}
}

func ValidateMessageSignatureSequenceWallet(chainID string, walletAddress string, signature string, message string) bool {
	seqAPI := api.NewAPIClient("https://api.sequence.app", http.DefaultClient)

	isValid, err := seqAPI.IsValidMessageSignature(context.Background(), chainID, walletAddress, message, signature)
	if !isValid {
		fmt.Println("Invalid Sequence Wallet Signature!", chainID, walletAddress, message, signature)
	}
	if err != nil {
		fmt.Println("Failed to Verify Sequence Wallet Signature?", err)
		isValid = false
	}
	//fmt.Println("isValid?", isValid)
	return isValid
}

func ValidateMessageSignatureTezosWallet(key, sig, msg string) string {
	pk, err := tezos.ParseKey(key)
	fmt.Println("Tezos Key Input", key, pk.Address())
	if err != nil {
		fmt.Println("Tezos Validate Error: ", err)
		return "fail"
	}
	fmt.Println("Tezos Key OK", pk)
	s, err := tezos.ParseSignature(sig)
	if err != nil {
		fmt.Println("Tezos Validate Error: ", err)
		return "fail"
	}
	fmt.Println("Tezos Sig Correct Format")
	m, err := hex.DecodeString(msg) //input as ASCII HEX from Beacon Payload data
	if err != nil {
		fmt.Println("Tezos Validate Error: ", err)
		return "fail"
	}
	digest := tezos.Digest([]byte(m))
	if err := pk.Verify(digest[:], s); err == nil {
		fmt.Println("Tezos Signature OK")
	} else {
		fmt.Println("Tezos Validate Error: ", err)
		return "fail"
	}
	return pk.Address().String()
}

func ValidateMessageSignatureNearWallet(key, sig, msg string) bool {
	msgBytes := []byte(msg) //message is just a string here

	// Decode the public key from the signature
	var pubKey ed25519.PublicKey
	keyBytes, err := hex.DecodeString(key) //
	if err != nil {
		fmt.Println("NEAR Validate Error 1: ", err, key)
	}
	copy(pubKey, keyBytes)

	sigBytes, err := hex.DecodeString(sig) //
	if err != nil {
		fmt.Println("NEAR Validate Error 2: ", err, sig)
	}

	// Verify the signature
	valid := ed25519.Verify(keyBytes, msgBytes, sigBytes)
	return valid
}

func Authenticate(walletName string, address string, nonce string, message string, sigHex string) (Authuser, error) {
	//removed print 9/29/2023 - SSO page might be authenticated a lot over and over though, TODO
	//fmt.Println("Authenticate: walletname: " + walletName + " \r\n address" + address + "\r\n msg: " + message + " sig: " + sigHex)

	pubKey := " "
	if strings.HasPrefix(address, "edpk") {
		pubKey = address
		pk, err := tezos.ParseKey(address)
		fmt.Println("Tezos Key", pk.Address().String())
		if err != nil {
			fmt.Println("Tezos Validate Error: ", err)
		}
		address = pk.Address().String()
	}
	//Massage data for NEAR wallets.  They have a common name and pub key natively
	if strings.HasSuffix(walletName, ".near") || strings.HasSuffix(walletName, ".testnet") ||
		(len(walletName) == 64 && !strings.HasPrefix(walletName, "0x")) ||
		strings.HasPrefix(walletName, "SP") {
		pubKey = address
		address = walletName
	}

	Authuser, err := Get(address)
	if err != nil {
		fmt.Println("Get Address Validate Error: ", address, err)
		return Authuser, err
	}
	//For SSO either need to require users use this nonce + check for its presence in the message
	// if Authuser.Nonce != nonce {
	// 	fmt.Println("Invalid Nonce Error: ", Authuser.Nonce, nonce)
	// 	return Authuser, ErrAuthError
	// }

	recoveredAddr := " "
	//fmt.Println("Signature Length: ", len(sigHex))
	if len(sigHex) > 400 { //594 without the 0x to be exact, but we can clean this up TODO: should be something more specific
		fmt.Println("Using Smart Contract Wallet Signature")
		chain := "mainnet"
		if walletName != "" {
			chain = walletName
		}
		isValidSequenceWalletSig := ValidateMessageSignatureSequenceWallet(chain, address, sigHex, message)

		if isValidSequenceWalletSig {
			recoveredAddr = address
			fmt.Println("Smart Contract Wallet Signature Valid!")
		}
	} else if strings.HasSuffix(walletName, ".near") || strings.HasSuffix(walletName, ".testnet") ||
		(len(walletName) == 64 && !strings.HasPrefix(walletName, "0x")) { //NEAR wallet
		fmt.Println("Using NEAR Wallet Signature")
		isValidNearWalletSig := ValidateMessageSignatureNearWallet(pubKey, sigHex, message)
		if isValidNearWalletSig {
			recoveredAddr = walletName
			fmt.Println("NEAR Wallet Signature Valid!", walletName)
		}
	} else if strings.HasPrefix(pubKey, "edpk") { //Tezos Wallet
		fmt.Println("Using Tezos Wallet Signature")
		returnsKeyForSuccess := ValidateMessageSignatureTezosWallet(pubKey, sigHex, message)
		if returnsKeyForSuccess == "fail" {
			fmt.Println("failed to recover Tezos Signature")
			return Authuser, err
		}
		recoveredAddr = returnsKeyForSuccess
		fmt.Println("Tezos Wallet Signature Valid!", returnsKeyForSuccess)
	} else if strings.HasPrefix(walletName, "SP") || strings.HasSuffix(walletName, ".btc") { //Stacks/BTC Wallet
		fmt.Println("Using Stacks/BTC Wallet Signature")
		// Decode hex-encoded serialized signature
		sigBytes, err := hex.DecodeString(sigHex)
		if err != nil {
			fmt.Println("sig err", err)
			return Authuser, err
		}
		keyBytes, err := hex.DecodeString(pubKey)
		if err != nil {
			fmt.Println("key err", err)
			return Authuser, err
		}
		hashMsgBytes, err := hex.DecodeString(message) //for STX this actually is the msg hash as input
		if err != nil {
			fmt.Println("hash err", err)
			return Authuser, err
		}

		//isValid := secp256k1.VerifySignature(keyBytes, msgBytes, sigBytes)
		isValid := secp256k1.VerifySignature(hashMsgBytes, sigBytes, keyBytes)
		if isValid == 0 {
			fmt.Println("sig error code: ", secp256k1.SignatureErrorString(hashMsgBytes, sigBytes, keyBytes), keyBytes) //hex.EncodeToString(secp256k1.RecoverPubkey(msgBytes, sigBytes)))
			fmt.Println("sig verify failed!")
			error := errors.New("stacks/btc sig verify failed")
			return Authuser, error
		}
		recoveredAddr = walletName
		fmt.Println("Stacks/BTC Wallet Signature Valid!", recoveredAddr)
	} else {
		sig := hexutil.MustDecode(sigHex)
		// https://github.com/ethereum/go-ethereum/blob/master/internal/ethapi/api.go#L516
		// check here why I am subtracting 27 from the last byte
		sig[crypto.RecoveryIDOffset] -= 27
		msg := accounts.TextHash([]byte(message))
		recovered, err := crypto.SigToPub(msg, sig)
		//fmt.Println("EVM signature ", sig)
		if err != nil {
			fmt.Println("EVM signature initial error: ", err)
			err = nil //reset error
			//this is a workaround for Ledger+Metamask - which has a known implementation difference to Ledger Live alone.
			sig[crypto.RecoveryIDOffset] += 27
			recovered, err = crypto.SigToPub(msg, sig)
			fmt.Println("EVM signature (mod)", sig)
			if err != nil {
				fmt.Println("failed to recover EVM signature ", err)
				return Authuser, err
			}
		}

		recoveredAddr = strings.ToLower(crypto.PubkeyToAddress(*recovered).Hex())
	}

	if !strings.EqualFold(Authuser.Address, recoveredAddr) {
		fmt.Println("Authuser (JWT Token) vs. Recovered Addr ", Authuser.Address, recoveredAddr)
		return Authuser, ErrAuthError
	}

	// update the nonce here so that the signature cannot be resused
	nonce, err = GetNonce()
	if err != nil {
		return Authuser, err
	}
	Authuser.Nonce = nonce
	Update(Authuser)

	return Authuser, nil
}

var (
	max  *big.Int
	once sync.Once
)

func GetNonce() (string, error) {
	once.Do(func() {
		max = new(big.Int)
		max.Exp(big.NewInt(2), big.NewInt(130), nil).Sub(max, big.NewInt(1))
	})
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}
	return n.Text(10), nil
}

func renderJson(r *http.Request, w http.ResponseWriter, statusCode int, res interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8 ")
	var body []byte
	if res != nil {
		var err error
		body, err = json.Marshal(res)
		if err != nil { // TODO handle me better
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
	w.WriteHeader(statusCode)
	if len(body) > 0 {
		w.Write(body)
	}
}

//call DelegateCash function
func GetDelegationsByDelegate(addressDelegateWallet string) []delegatecash.IDelegationRegistryDelegationInfo {
	// Connect to an ethereum node
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/" + os.Getenv("INFURA_V3"))
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// Create an instance of the contract
	contractAddress := common.HexToAddress("0x00000000000076A84feF008CDAbe6409d2FE638B")
	instance, err := delegatecash.NewDelegatecash(contractAddress, client)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	delegateAddress := common.HexToAddress(addressDelegateWallet)
	// Call the contract method
	var result []delegatecash.IDelegationRegistryDelegationInfo
	result, err = instance.GetDelegationsByDelegate(nil, delegateAddress)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	//[{x x x x}, {x x x x}]
	//{1 <cold_addr> <delegate_addr> 0 0} //delegate full wallet
	//{3 <cold_addr> <delegate_addr> <nft_addr> <nft_id>} //Delegate for single NFT
	//fmt.Println(result)
	return result
}

// func renderJsonWithCookie(r *http.Request, w http.ResponseWriter, statusCode int, cookie http.Cookie, res interface{}) {
// 	w.Header().Set("Content-Type", "application/json; charset=utf-8 ")
// 	var body []byte
// 	if res != nil {
// 		var err error
// 		body, err = json.Marshal(res)
// 		if err != nil { // TODO handle me better
// 			w.WriteHeader(http.StatusInternalServerError)
// 		}
// 	}
// 	w.WriteHeader(statusCode)
// 	if len(body) > 0 {
// 		w.Write(body)
// 	}
// 		// Finally, we set the client cookie for "token" as the JWT we just generated
// 	// we also set an expiry time which is the same as the token itself
// 	http.SetCookie(w, &cookie)
// }
