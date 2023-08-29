package referrals

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rest-go-demo/auth"
	"rest-go-demo/database"
	_ "rest-go-demo/docs"
	"rest-go-demo/entity"
)

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
func CreateReferralCodes(w http.ResponseWriter, r *http.Request) {
	Authuser := auth.GetUserFromReqContext(r)
	walletaddr := Authuser.Address

	fmt.Printf("Create refferal code for wallet: %#v\n", walletaddr)

	//get all items that relate to passed in owner/address
	var code []entity.Referralcode
	database.Connector.Where("walletaddr = ?", walletaddr).Find(&code)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	json.NewEncoder(w).Encode(code)
}
