package referrals

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"math/rand"
	"net/http"
	"os"
	"rest-go-demo/auth"
	"rest-go-demo/database"
	_ "rest-go-demo/docs"
	"rest-go-demo/entity"
	"rest-go-demo/wc_analytics"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

var currentLeaderboard []ChatStatistics
var currentOuraLeaderboard []OuraChatStatistics

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func InitRandom() {
	rand.Seed(time.Now().UnixNano())
}
func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GetLeaderboardDataGlobal() []ChatStatistics {
	return currentLeaderboard
}

func GetOuraLeaderboardDataGlobal() []OuraChatStatistics {
	return currentOuraLeaderboard
}

// GetInboxByOwner godoc
// @Summary     Get Inbox Summary With Last Message
// @Description Get Each 1-on-1 Conversation, NFT and Community Chat For Display in Inbox
// @Tags        Inbox
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       address path    string true "Wallet Address"
// @Success     200     {array} entity.Chatiteminbox
// @Router      /v1/get_referral_code [get]
func GetReferralCode(w http.ResponseWriter, r *http.Request) {
	Authuser := auth.GetUserFromReqContext(r)
	walletaddr := Authuser.Address

	fmt.Printf("GetReferralCode for wallet: %#v\n", walletaddr)

	//get all items that relate to passed in owner/address
	var code []entity.Referralcode
	database.Connector.Where("walletaddr = ?", walletaddr).Find(&code)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	json.NewEncoder(w).Encode(code)
}

func GetReferralCodeAddr(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	walletaddr := vars["address"]

	fmt.Printf("GetReferralCode for wallet: %#v\n", walletaddr)

	//get all items that relate to passed in owner/address
	var code []entity.Referralcode
	database.Connector.Where("walletaddr = ?", walletaddr).Find(&code)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	json.NewEncoder(w).Encode(code)
}

// just to test with mainly
func CreateReferralCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	walletaddr := vars["address"]
	apiKey := r.Header.Get("Authorization")
	if len(apiKey) > 0 {
		const prefix = "Bearer "
		if len(apiKey) < len(prefix) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		apiKey = apiKey[len(prefix):]
		if strings.Contains(os.Getenv("ADMIN_API_KEY_LIST"), apiKey) {
			fmt.Printf("Create referral code for wallet via ADMIN: %#v\n", walletaddr)

			//get all items that relate to passed in owner/address
			var code entity.Referralcode
			code.Code = "wc-" + randSeq(10)
			code.Walletaddr = walletaddr
			code.Date = time.Now()
			database.Connector.Create(&code)

			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("X-Content-Type-Options", "nosniff")
			json.NewEncoder(w).Encode(code)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}

// not called from API - called upon new user signup
func CreateReferralCodeInternal(walletaddr string) {
	fmt.Printf("Create 3 New Referral Codes for Wallet: %#v\n", walletaddr)

	var code entity.Referralcode
	code.Code = "va-" + randSeq(10)
	code.Walletaddr = walletaddr
	code.Date = time.Now()
	database.Connector.Create(&code)

	var code1 entity.Referralcode
	code1.Code = "va-" + randSeq(10)
	code1.Walletaddr = walletaddr
	code1.Date = time.Now()
	database.Connector.Create(&code1)

	var code2 entity.Referralcode
	code2.Code = "va-" + randSeq(10)
	code2.Walletaddr = walletaddr
	code2.Date = time.Now()
	database.Connector.Create(&code2)
}

func CreateDailyReferralCodes() {
	fmt.Println("Creating daily referral codes!")

	//only create new daily codes for those users who have no unused codes remaining
	var result []entity.Referralcode
	database.Connector.Raw("CALL InsertVanaReferralCodes()").Scan(&result)

	//gorm results were not showing correct number of rows returned, so I had to manually do this in the SP (UGLY AF)
	fmt.Println("Number of New Daily Referral Codes Created: ", len(result))
}

func RedeemReferralCode(w http.ResponseWriter, r *http.Request) {
	Authuser := auth.GetUserFromReqContext(r)
	walletaddr := Authuser.Address
	vars := mux.Vars(r)
	referral_code := vars["code"]

	//get all items that relate to passed in referral code
	var code []entity.Referralcode
	dbQuery := database.Connector.Where("code = ?", referral_code).Where("redeemed = ?", 0).Find(&code)

	//allow users to sign in without a code, just don't get referral points
	if referral_code == "wc-test" {
		var uservalid entity.Referraluser
		uservalid.Referralcode = "wc-test"
		uservalid.Walletaddr = walletaddr
		database.Connector.Create(&uservalid)

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		json.NewEncoder(w).Encode(code)
		return
	}

	//don't let people redeem their own codes
	if dbQuery.RowsAffected > 0 && code[0].Walletaddr != walletaddr {

		var result = database.Connector.Model(&entity.Referralcode{}).
			Where("code = ?", referral_code).
			Update("redeemed", true)

		//set user as validated in the referral code table (used separate table in the case we drop this in future)
		var uservalid entity.Referraluser
		uservalid.Referralcode = referral_code
		uservalid.Walletaddr = walletaddr
		database.Connector.Create(&uservalid)

		code[0].Redeemed = true //for a proper return value - not sure if we will actually use it

		if result.RowsAffected > 0 {
			fmt.Printf("Redeemed referral code for wallet: %#v\n", code[0].Walletaddr)
			if code[0].Walletaddr != "0xtest" {
				wc_analytics.SendCustomEvent(Authuser.Address, "REFERRAL_CODE_REDEEMED")
			}
		} else {
			fmt.Printf("Redeemed referral failed!!!!")
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		json.NewEncoder(w).Encode(code)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

type ChatStatistics struct {
	Walletaddr    string
	Username      string
	Pfpdata       string
	MessagesTx    int
	MessagesRx    int
	GroupMessages int
	UniqueConvos  int
	Installedsnap string
	RedeemedCount int
	Points        int
}

type OuraChatStatistics struct {
	Wallet     string
	Numuploads int
	Tokens     string
	Name       string
}

type ChatStatisticsReturn struct {
	Walletaddr string
	Username   string
	Pfpdata    string
	// MessagesTx    int
	// MessagesRx    int
	// UniqueConvos  int
	// Installedsnap string
	// RedeemedCount int
	Points int
}

func GetLeaderboardDataCronJob() {
	var results []ChatStatistics
	dbQuery := database.Connector.Raw("CALL get_leaderboard_data()").Scan(&results)
	//fmt.Println("get leaderboard: ", dbQuery.Error, results)

	if dbQuery.Error != nil {
		return
	}

	currentLeaderboard = results
}

func GetLeaderboardData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	json.NewEncoder(w).Encode(currentLeaderboard)
}

// code for eth_call
// ethCallRequest is the JSON-RPC request payload.
type ethCallRequest struct {
	JSONRPC string        `json:"jsonrpc"`
	ID      int           `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

// ethCallResponse is the JSON-RPC response payload.
type ethCallResponse struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  string `json:"result"`
	Error   *struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

// balanceOf performs an eth_call against the given rpcURL and tokenAddr.
// walletAddr must be a 0x-prefixed hex string. It returns the raw balance (big.Int).
func balanceOf(rpcURL, tokenAddr, walletAddr string) (*big.Int, error) {
	// 1) Method ID for balanceOf(address) is 0x70a08231
	methodID, _ := hex.DecodeString("70a08231")

	// 2) Remove “0x” prefix from wallet address and decode
	if len(walletAddr) < 2 || walletAddr[:2] != "0x" {
		return nil, fmt.Errorf("invalid wallet address: %s", walletAddr)
	}
	addrBytes, err := hex.DecodeString(walletAddr[2:])
	if err != nil {
		return nil, fmt.Errorf("invalid wallet address: %w", err)
	}

	// 3) Build the data: methodID (4 bytes) + 32-byte padded address
	data := append(methodID, make([]byte, 32-len(addrBytes))...)
	data = append(data, addrBytes...)

	// 4) Construct the JSON-RPC request
	reqBody := ethCallRequest{
		JSONRPC: "2.0",
		ID:      1,
		Method:  "eth_call",
		Params: []interface{}{
			map[string]interface{}{
				"to":   tokenAddr,
				"data": "0x" + hex.EncodeToString(data),
			},
			"latest",
		},
	}
	payload, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	// 5) Send HTTP POST
	resp, err := http.Post(rpcURL, "application/json", bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 6) Decode the response
	var respBody ethCallResponse
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return nil, err
	}
	if respBody.Error != nil {
		return nil, fmt.Errorf("rpc error %d: %s", respBody.Error.Code, respBody.Error.Message)
	}

	// 7) Parse hex result into *big.Int
	//    Trim "0x" prefix and decode
	resultHex := respBody.Result
	if len(resultHex) >= 2 && resultHex[:2] == "0x" {
		resultHex = resultHex[2:]
	}
	bal := new(big.Int)
	bal.SetString(resultHex, 16)
	return bal, nil
}

func getIntraTokenBalance(walletAddress string) string {
	rpcURL := os.Getenv("VANA_RPC_URL") // replace with actual RPC URL
	tokenAddr := os.Getenv("VANA_DLP_TOKEN")
	walletAddr := walletAddress

	bal, err := balanceOf(rpcURL, tokenAddr, walletAddr)
	if err != nil {
		fmt.Println("balanceOf failed: %v", err)
		return ""
	}

	// Print raw balance
	//fmt.Printf("Raw balance: %s\n", bal.String())

	// If the token has 18 decimals, convert to human-readable
	human := new(big.Rat).SetFrac(bal, big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil))
	fmt.Println("INTRA Token Balance: ", walletAddr, human.FloatString(18))
	return human.FloatString(18)
}

func GetOuraLeaderboardDataCronJob() {
	var results []OuraChatStatistics
	dbQuery := database.Connector.Raw("CALL get_intra_leaderboard_info()").Scan(&results)
	fmt.Println("get intra leaderboard: ", dbQuery.Error, results)

	if dbQuery.Error != nil {
		return
	}

	for i, result := range results {
		results[i].Tokens = getIntraTokenBalance(result.Wallet)
	}

	currentOuraLeaderboard = results
}

func GetOuraLeaderboardData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	json.NewEncoder(w).Encode(currentOuraLeaderboard)
}

func GetOuraLeaderboardDataSingle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["address"]
	var results OuraChatStatistics
	for i := 0; i < len(currentOuraLeaderboard); i++ {
		if strings.EqualFold(currentOuraLeaderboard[i].Wallet, address) {
			//fmt.Println("get leaderboard single - found address: ", currentLeaderboard[i])
			results = currentOuraLeaderboard[i]
			break
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	json.NewEncoder(w).Encode(results)
}

func GetLeaderboardDataSingle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["address"]
	var results ChatStatistics
	for i := 0; i < len(currentLeaderboard); i++ {
		if strings.EqualFold(currentLeaderboard[i].Walletaddr, address) {
			//fmt.Println("get leaderboard single - found address: ", currentLeaderboard[i])
			results = currentLeaderboard[i]
			break
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	json.NewEncoder(w).Encode(results)
}

// this is used upon login to check if a user has entered a valid code or not in the past
// used similar to getting user name so we don't prompt them if its already set.
func GetHasEnteredValidCode(w http.ResponseWriter, r *http.Request) {
	Authuser := auth.GetUserFromReqContext(r)
	walletaddr := Authuser.Address

	//get all items that relate to passed in referral code
	var uservalid []entity.Referraluser
	database.Connector.Where("walletaddr = ?", walletaddr).Find(&uservalid)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(uservalid)
}
