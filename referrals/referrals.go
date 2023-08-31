package referrals

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"rest-go-demo/auth"
	"rest-go-demo/database"
	_ "rest-go-demo/docs"
	"rest-go-demo/entity"
	"time"

	"github.com/gorilla/mux"
)

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

// GetInboxByOwner godoc
// @Summary Get Inbox Summary With Last Message
// @Description Get Each 1-on-1 Conversation, NFT and Community Chat For Display in Inbox
// @Tags Inbox
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param address path string true "Wallet Address"
// @Success 200 {array} entity.Chatiteminbox
// @Router /v1/get_referral_code [get]
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

//just to test with postman for now - either we will do this for all addresses periodcially or need to take wallet address as input
//possibly need to use an admin API key for authentication here as well not the user JWT
func CreateReferralCode(w http.ResponseWriter, r *http.Request) {
	Authuser := auth.GetUserFromReqContext(r)
	walletaddr := Authuser.Address

	fmt.Printf("Create referral code for wallet: %#v\n", walletaddr)

	//get all items that relate to passed in owner/address
	var code entity.Referralcode
	code.Code = "wc-" + randSeq(10)
	code.Walletaddr = walletaddr
	code.Date = time.Now()
	database.Connector.Create(&code)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	json.NewEncoder(w).Encode(code)
}

// func CreateReferralCodesForAll() {

// 	//get all items that relate to passed in owner/address
// 	var code entity.Referralcode
// 	code.Code = "wc-" + randSeq(10)
// 	code.Walletaddr = walletaddr
// 	code.Date = time.Now()
// 	database.Connector.Create(&code)

// }

func RedeemReferralCode(w http.ResponseWriter, r *http.Request) {
	Authuser := auth.GetUserFromReqContext(r)
	walletaddr := Authuser.Address
	vars := mux.Vars(r)
	referral_code := vars["code"]

	//get all items that relate to passed in referral code
	var code []entity.Referralcode
	dbQuery := database.Connector.Where("code = ?", referral_code).Find(&code)

	//don't let people redeem their own codes
	if dbQuery.RowsAffected > 0 && code[0].Walletaddr != walletaddr {

		var result = database.Connector.Model(&entity.Referralcode{}).
			Where("code = ?", referral_code).
			Update("redeemed", true)

		if result.RowsAffected > 0 {
			fmt.Printf("Redeemed referral code for wallet: %#v\n", code[0].Walletaddr)
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
	Walletaddr   string
	MessagesTx   int
	MessagesRx   int
	UniqueConvos int
}

func GetLeaderboardData(w http.ResponseWriter, r *http.Request) {
	var results []ChatStatistics
	dbQuery := database.Connector.Raw("CALL get_leaderboard_data()").Scan(&results)
	//fmt.Println("get leaderboard: ", dbQuery.Error, results)

	if dbQuery.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	json.NewEncoder(w).Encode(results)
}
