package controllers

import (
	"archive/zip"
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"rest-go-demo/auth"
	"rest-go-demo/database"
	"rest-go-demo/email"
	"rest-go-demo/entity"
	"rest-go-demo/referrals"
	"rest-go-demo/vanaencrypt"
	"rest-go-demo/vanatransact"
	"rest-go-demo/wc_analytics"
	"sync"

	"strconv"
	"strings"
	"time"

	_ "rest-go-demo/docs"

	goaway "github.com/TwiN/go-away"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	ens "github.com/wealdtech/go-ens/v3"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gorilla/mux"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var telegramUpdateOffset = 0
var throttleInboxCounterPerUser = make(map[string]int64) //only access via mutex function!

// Retrieve the environment variable value with an array-like data as a comma-separated string
var tgSupportWalletsCsvString = ""
var tgSupportWalletArray []string

var tgSupporChatIdsCsvString = ""
var tgSupportChatIdsArray []string

var tgSupportAdminsCsvString = ""
var tgSupportAdminsArray []string

var tgSupportBlockedUsers []string

var giveawayMessageCounter = 0

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func InitGlobals() {
	tgSupportWalletsCsvString = strings.ToLower(os.Getenv("TG_SUPPORT_WALLETS")) //TODO - non EVM addres may require this to be case sensitive
	tgSupportWalletArray = strings.Split(tgSupportWalletsCsvString, ",")

	tgSupporChatIdsCsvString = os.Getenv("TG_SUPPORT_CHAT_IDS")
	tgSupportChatIdsArray = strings.Split(tgSupporChatIdsCsvString, ",")

	tgSupportAdminsCsvString = os.Getenv("TG_SUPPORT_ADMIN_IDS")
	tgSupportAdminsArray = strings.Split(tgSupportAdminsCsvString, ",")
}

// This function is used for MM Snaps specifically
// @Router /v1/get_latest_unread/{address} [get]
func GetLastMsgToOwner(w http.ResponseWriter, r *http.Request) {
	Authuser := auth.GetUserFromReqContext(r)
	key := Authuser.Address

	var chat []entity.Chatitem
	dbResult := database.Connector.Where("toaddr = ?", key).Where("msgread != ?", true).Find(&chat)

	var chatReturn entity.Chatitem
	if dbResult.RowsAffected > 0 {
		chatReturn = chat[0]
	}

	//record that the user has used snaps
	var settings []entity.Settings
	database.Connector.Where("walletaddr = ?", key).Where("installedsnap != ?", "true").Find(&settings)
	if len(settings) > 0 {
		fmt.Println("New Metamask Snaps Install!! ", Authuser.Address)
		database.Connector.Model(&entity.Settings{}).Where("walletaddr = ?", key).Where("installedsnap != ?", "true").Update("installedsnap", "true")
		wc_analytics.SendCustomEvent(Authuser.Address, "METAMASK_SNAPS_INSTALL")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	json.NewEncoder(w).Encode(chatReturn)
}

var mu sync.Mutex

func incrementAndCheck(key string) bool {
	// Step 1: Acquire the lock at the beginning
	mu.Lock()
	// Step 2: Defer the unlock until the end, ensuring it happens even if there's an early return
	defer mu.Unlock()

	// Critical section: Increment the counter for the specified key
	throttleInboxCounterPerUser[key]++

	// Check if the incremented value is divisible by 25
	return throttleInboxCounterPerUser[key]%25 == 0
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
// @Router      /v1/get_inbox/{address} [get]
func GetInboxByOwner(w http.ResponseWriter, r *http.Request) {

	//GetInboxByID returns the latest message for each unique conversation
	//vars := mux.Vars(r)
	Authuser := auth.GetUserFromReqContext(r)
	key := Authuser.Address //vars["address"] //owner of the inbox

	//fmt.Printf("GetInboxByOwner: %#v\n", key)

	//get all items that relate to passed in owner/address
	// var chat []entity.Chatitem
	// database.Connector.Where("fromaddr = ?", key).Or("toaddr = ?", key).Find(&chat)

	//get unique conversation addresses
	var uniqueChatMembers []entity.Chatiteminboxconvos
	//right now this is limiting the inbox to 100 messages (TODO: allow for scroll and update)
	database.Connector.Raw("CALL GetUniqueAddressesWithInput(?)", key).Scan(&uniqueChatMembers)

	// for _, chatitem := range chat {
	// 	//fmt.Printf("search for unique addrs")
	// 	if chatitem.Fromaddr != key {
	// 		if !stringInSlice(chatitem.Fromaddr, uniqueChatMembers) {
	// 			uniqueChatMembers = append(uniqueChatMembers, chatitem.Fromaddr)
	// 		}
	// 	}
	// 	if chatitem.Toaddr != key {
	// 		if !stringInSlice(chatitem.Toaddr, uniqueChatMembers) {
	// 			uniqueChatMembers = append(uniqueChatMembers, chatitem.Toaddr)
	// 		}
	// 	}
	// }

	//fmt.Printf("find first message now")
	//for each unique chat member that is not the owner addr, get the latest message
	var userInbox []entity.Chatiteminbox
	for _, chatmember := range uniqueChatMembers {
		//fmt.Println("Unique Chat Addrs Result: ", chatmember.Address)
		// //add Unread msg count to both first/second items since we don't know which one is newer yet
		var chatCount []entity.Chatitem
		database.Connector.Where("fromaddr = ?", chatmember.Address).Where("toaddr = ?", key).Where("msgread != ?", true).Find(&chatCount)

		// //get name for return val
		var addrname entity.Addrnameitem
		database.Connector.Where("address = ?", chatmember.Address).Find(&addrname)

		//database view - local code replaced 7/14
		var vchatitem entity.V_chatitem
		var dbQuery = database.Connector.Where("fromaddr = ? AND toaddr = ?", key, chatmember.Address).Find(&vchatitem)
		//var dbQuery = database.Connector.Raw("select * from v_chatitems WHERE fromaddr in('0xcafebabe', '0xdeadbeef');").Scan(&testView)

		var itemToInsert entity.Chatiteminbox
		if dbQuery.RowsAffected > 0 {
			itemToInsert.Id = vchatitem.Id
			itemToInsert.Fromaddr = vchatitem.Fromaddr
			itemToInsert.Toaddr = vchatitem.Toaddr
			itemToInsert.Timestamp = vchatitem.Timestamp
			itemToInsert.Timestamp_dtm = vchatitem.Timestamp_dtm
			itemToInsert.Msgread = vchatitem.Msgread
			itemToInsert.Message = vchatitem.Message
			itemToInsert.Unreadcnt = len(chatCount)
			itemToInsert.Contexttype = entity.DM
			itemToInsert.Type = entity.Message
			itemToInsert.Sendername = addrname.Name
			itemToInsert.Encryptsymkey = vchatitem.Encryptsymkey
			itemToInsert.Litaccesscond = vchatitem.Litaccesscond
			//fmt.Printf("encrypted symmetric LIT key: %#v %#v %#v\n", vchatitem.Encryptsymkey, vchatitem.Toaddr, vchatitem.Fromaddr)

			var imgname entity.Imageitem
			var result = database.Connector.Where("addr = ?", chatmember.Address).Find(&imgname)
			if result.RowsAffected > 0 {
				itemToInsert.LogoData = imgname.Base64data
			}

			found := false
			for i := 0; i < len(userInbox); i++ {
				if itemToInsert.Timestamp_dtm.After(userInbox[i].Timestamp_dtm) {
					userInbox = append(userInbox[:i+1], userInbox[i:]...)
					userInbox[i] = itemToInsert
					found = true
					break
				}
			}
			if !found {
				userInbox = append(userInbox, itemToInsert)
			}
			//end timesort the append
		}
	}

	//now get bookmarked/joined groups as well but fit it into the inbox return val type
	var bookmarks []entity.Bookmarkitem
	database.Connector.Where("walletaddr = ?", key).Find(&bookmarks)

	//TODO: need to throttle these 2 calls to auto-join?
	//should auto-join them to the community chat
	if strings.HasPrefix(key, "0x") || strings.HasSuffix(key, ".eth") {
		//golang maps are not concurrently accessible, must use locks if read/writing in threads
		updateUserInfo := incrementAndCheck(key)
		if updateUserInfo {
			AutoJoinCommunitiesByChainWithDelegates(key, "ethereum") //Moralis uses "eth" instead of "ethereum" but we change this at lower level
			//AutoJoinCommunitiesByChainWithDelegates(key, "polygon")
			AutoJoinPoapChats(key)
		}
	}

	//now add last message from group chat this bookmark is for
	var gchat []entity.Groupchatitem //even though I use this in a Last() function I need to store as an array, or subsequenct DB queries fail!
	for idx := 0; idx < len(bookmarks); idx++ {
		//fmt.Printf("bookmarks: %#v\n", bookmarks[i])
		//fmt.Printf("\nnftaddr: %#v\n", bookmarks[idx].Nftaddr)
		dbQuery := database.Connector.Where("nftaddr = ?", bookmarks[idx].Nftaddr).Last(&gchat)
		//fmt.Printf("dbQuery: %#v\n", dbQuery.Error)

		var returnItem entity.Chatiteminbox
		if dbQuery.RowsAffected == 0 {
			//if this chat is new/empty just return the basic info
			returnItem.Nftaddr = bookmarks[idx].Nftaddr
			returnItem.Contexttype = entity.Community
			if strings.HasPrefix(returnItem.Nftaddr, "0x") {
				returnItem.Contexttype = entity.Nft
				returnItem.Chain = bookmarks[idx].Chain
			}
			if strings.HasPrefix(returnItem.Nftaddr, "poap_") {
				returnItem.Contexttype = entity.Nft
				returnItem.Chain = bookmarks[idx].Chain
			}
			userInbox = append(userInbox, returnItem)
			continue
		}
		//fmt.Printf("bookmarkchat: %#v\n", gchat)

		var groupchat = gchat[0]

		//get num unread messages
		//var chatCnt []entity.Groupchatitem
		var chatCount int
		var chatReadTime entity.Groupchatreadtime
		dbQuery = database.Connector.Where("fromaddr = ?", key).Where("nftaddr = ?", groupchat.Nftaddr).Find(&chatReadTime)
		//if no respsonse to this query, its the first time a user is reading the chat history, send it all
		if dbQuery.RowsAffected == 0 {
			//fmt.Printf("sending all values! \n")
			database.Connector.Model(&entity.Groupchatitem{}).Where("nftaddr = ?", groupchat.Nftaddr).Count(&chatCount)
			//database.Connector.Where("nftaddr = ?", groupchat.Nftaddr).Find(&chatCnt)
		} else {
			database.Connector.Model(&entity.Groupchatitem{}).Where("timestamp_dtm > ?", chatReadTime.Readtimestamp_dtm).Where("nftaddr = ?", groupchat.Nftaddr).Count(&chatCount)
			//database.Connector.Where("timestamp_dtm > ?", chatReadTime.Readtimestamp_dtm).Where("nftaddr = ?", groupchat.Nftaddr).Find(&chatCnt)
			//fmt.Printf("sending time based count \n")
		}
		//end get num unread messages

		returnItem.Id = groupchat.Id
		returnItem.Message = groupchat.Message
		returnItem.Timestamp = groupchat.Timestamp
		returnItem.Timestamp_dtm = groupchat.Timestamp_dtm
		returnItem.Nftaddr = groupchat.Nftaddr
		returnItem.Fromaddr = groupchat.Fromaddr
		returnItem.Unreadcnt = chatCount
		returnItem.Type = groupchat.Type
		returnItem.Chain = bookmarks[idx].Chain
		//retrofit old messages prior to setting Type
		if returnItem.Type != entity.Message && returnItem.Type != entity.Welcome {
			returnItem.Type = entity.Message
		}
		returnItem.Contexttype = entity.Community

		//get common name from nftaddress
		var addrname entity.Addrnameitem
		var result = database.Connector.Where("address = ?", groupchat.Nftaddr).Find(&addrname)
		if result.RowsAffected > 0 {
			returnItem.Name = addrname.Name
		}
		//not sure if long term we will store by name (WalletChat HQ) or nftaddr (walletchat)
		var imgname entity.Imageitem
		result = database.Connector.Where("addr = ?", groupchat.Nftaddr).Find(&imgname)
		if result.RowsAffected > 0 {
			returnItem.LogoData = imgname.Base64data
		}

		//until we fix up old tables, we can hack this to double check
		if strings.HasPrefix(returnItem.Nftaddr, "0x") {
			returnItem.Contexttype = entity.Nft
		}
		if strings.HasPrefix(returnItem.Nftaddr, "poap_") {
			returnItem.Contexttype = entity.Nft
		}

		returnItem.Sendername = ""
		if returnItem.Message == "" {
			var unsetTime time.Time
			var noInt int
			returnItem.Unreadcnt = noInt
			returnItem.Timestamp = unsetTime.String()
		}

		//timesort the append
		found := false
		for i := 0; i < len(userInbox); i++ {
			if returnItem.Timestamp_dtm.After(userInbox[i].Timestamp_dtm) {
				userInbox = append(userInbox[:i+1], userInbox[i:]...)
				userInbox[i] = returnItem
				found = true
				break
			}
		}
		if !found {
			userInbox = append(userInbox, returnItem)
		}
		//userInbox = append(userInbox, returnItem)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	json.NewEncoder(w).Encode(userInbox)
}

// GetUnreadMsgCntTotal godoc
// @Summary     Get all unread messages TO a specific user, used for total count notification at top notification bar
// @Description Get Each 1-on-1 Conversation, NFT and Community Chat For Display in Inbox
// @Tags        Inbox
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       address path      string true "Wallet Address"
// @Success     200     {integer} int
// @Router      /v1/get_unread_cnt/{address} [get]
func GetUnreadMsgCntTotal(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	Authuser := auth.GetUserFromReqContext(r)
	key := Authuser.Address //vars["address"]

	var chat []entity.Chatitem
	database.Connector.Where("toaddr = ?", key).Where("msgread != ?", true).Find(&chat)

	//get group chat unread items as well

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(len(chat))
}

// GetUnreadMsgCntTotalExternal godoc
// @Summary     Unread DM count - unauthenticated for user with Android App
// @Description Unread DM count
// @Tags        Inbox
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       address path      string true "Wallet Address"
// @Success     200     {integer} int
// @Router      /get_unread_cnt/{address} [get]
func GetUnreadMsgCntTotalExternal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["address"]

	var chat []entity.Chatitem
	database.Connector.Where("toaddr = ?", key).Where("msgread != ?", true).Find(&chat)

	//get group chat unread items as well

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(len(chat))
}

// GetUnreadMsgCntTotalByType godoc
// @Summary     Get all unread messages TO a specific user, used for total count notification at top notification bar
// @Description Get Each 1-on-1 Conversation, NFT and Community Chat For Display in Inbox
// @Tags        Unused/Legacy
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       address path      string true "Wallet Address"
// @Param       type    path      string true "Message Type - nft|community|dm|all"
// @Success     200     {integer} int
// @Router      /v1/get_unread_cnt_by_type/{address}/{type} [get]
func GetUnreadMsgCntTotalByType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Authuser := auth.GetUserFromReqContext(r)
	key := Authuser.Address //key := vars["address"]
	msgtype := vars["type"] //nft/community/DM/ALL

	msgCntTotal := 0

	var bookmarks []entity.Bookmarkitem
	database.Connector.Where("walletaddr = ?", key).Find(&bookmarks)

	//now add last message from group chat this bookmark is for
	var gchat []entity.Groupchatitem //even though I use this in a Last() function I need to store as an array, or subsequenct DB queries fail!
	if msgtype == entity.Nft || msgtype == entity.Community || msgtype == entity.All {
		for idx := 0; idx < len(bookmarks); idx++ {
			dbQuery := database.Connector.Where("nftaddr = ?", bookmarks[idx].Nftaddr).Last(&gchat)
			if dbQuery.RowsAffected == 0 {
				continue
			}
			var groupchat = gchat[0]

			//get num unread messages
			var chatCnt []entity.Groupchatitem
			var chatReadTime entity.Groupchatreadtime
			dbQuery = database.Connector.Where("fromaddr = ?", key).Where("nftaddr = ?", groupchat.Nftaddr).Find(&chatReadTime)
			//if no respsonse to this query, its the first time a user is reading the chat history
			if dbQuery.RowsAffected == 0 {
				database.Connector.Where("nftaddr = ?", groupchat.Nftaddr).Find(&chatCnt)
			} else {
				database.Connector.Where("timestamp_dtm > ?", chatReadTime.Readtimestamp_dtm).Where("nftaddr = ?", groupchat.Nftaddr).Find(&chatCnt)
			}
			//end get num unread messages

			if strings.HasPrefix(groupchat.Nftaddr, "0x") {
				if msgtype == entity.Nft || msgtype == entity.All {
					msgCntTotal += len(chatCnt)
				}
			} else if msgtype == entity.Community || msgtype == entity.All {
				msgCntTotal += len(chatCnt)
			}
		}
	}
}

// func PutUnreadcnt(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	walletaddr := vars["address"]

// 	requestBody, _ := ioutil.ReadAll(r.Body)
// 	var config entity.Unreadcountitem
// 	json.Unmarshal(requestBody, &config)

// 	var findConfig entity.Unreadcountitem
// 	var dbQuery = database.Connector.Where("walletaddr = ?", walletaddr).Find(&findConfig)

// 	if dbQuery.RowsAffected == 0 {
// 		config.Walletaddr = walletaddr
// 		database.Connector.Create(&config)
// 	} else {
// 		database.Connector.Model(&entity.Unreadcountitem{}).Where("walletaddr = ?", walletaddr).Update("dm", config.Dm)
// 		database.Connector.Model(&entity.Unreadcountitem{}).Where("walletaddr = ?", walletaddr).Update("nft", config.Nft)
// 		database.Connector.Model(&entity.Unreadcountitem{}).Where("walletaddr = ?", walletaddr).Update("community", config.Community)
// 	}

// 	w.Header().Set("Content-Type", "application/json")
//  w.Header().Set("X-Content-Type-Options", "nosniff")
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(true)
// }

func LocalGetUnread(address string) entity.Unreadcountitem {
	var config entity.Unreadcountitem
	var bookmarks []entity.Bookmarkitem
	database.Connector.Where("walletaddr = ?", address).Find(&bookmarks)

	//now add last message from group chat this bookmark is for
	var gchat []entity.Groupchatitem //even though I use this in a Last() function I need to store as an array, or subsequenct DB queries fail!
	for idx := 0; idx < len(bookmarks); idx++ {
		dbQuery := database.Connector.Where("nftaddr = ?", bookmarks[idx].Nftaddr).Last(&gchat)
		if dbQuery.RowsAffected == 0 {
			continue
		}
		var groupchat = gchat[0]

		//get num unread messages
		var chatCnt []entity.Groupchatitem
		var chatReadTime entity.Groupchatreadtime
		dbQuery = database.Connector.Where("fromaddr = ?", address).Where("nftaddr = ?", groupchat.Nftaddr).Find(&chatReadTime)
		//if no respsonse to this query, its the first time a user is reading the chat history
		if dbQuery.RowsAffected == 0 {
			database.Connector.Where("nftaddr = ?", groupchat.Nftaddr).Find(&chatCnt)
		} else {
			database.Connector.Where("timestamp_dtm > ?", chatReadTime.Readtimestamp_dtm).Where("nftaddr = ?", groupchat.Nftaddr).Find(&chatCnt)
		}
		//end get num unread messages

		if strings.HasPrefix(groupchat.Nftaddr, "0x") {
			config.Nft += len(chatCnt)
		} else {
			config.Community += len(chatCnt)
		}
	}

	var chat []entity.Chatitem
	database.Connector.Where("toaddr = ?", address).Where("msgread != ?", true).Find(&chat)
	config.Dm = len(chat)

	return config
}

// GetUnreadcnt godoc
// @Summary     Get all unread messages TO a specific user, used for total count notification at top notification bar
// @Description Get Unread count just given an address
// @Tags        Inbox
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       address path      string true "Wallet Address"
// @Success     200     {integer} int
// @Router      /v1/unreadcount/{address} [get]
func GetUnreadcnt(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	Authuser := auth.GetUserFromReqContext(r)
	key := Authuser.Address
	//key := vars["address"]

	//get configured items from DB
	config := LocalGetUnread(key)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(config)
}

// GetUnreadMsgCntNft godoc
// @Summary     Get all unread messages for a specific NFT context
// @Description Get Unread count for specifc NFT context given a wallet address and specific NFT
// @Tags        NFT
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       address path      string true "Wallet Address"
// @Param       nftaddr path      string true "NFT Contract Address"
// @Param       nftid   path      string true "NFT ID"
// @Success     200     {integer} int
// @Router      /v1/get_unread_cnt/{address}/{nftaddr}/{nftid} [get]
func GetUnreadMsgCntNft(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Authuser := auth.GetUserFromReqContext(r)
	key := Authuser.Address
	//key := vars["address"]
	addr := vars["nftaddr"]
	id := vars["nftid"]

	var chat []entity.Chatitem
	database.Connector.Where("toaddr = ?", key).Where("nftaddr = ?", addr).Where("nftid = ?", id).Where("msgread = ?", false).Find(&chat)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(len(chat))
}

// GetUnreadMsgCntNft godoc
// @Summary     Get all unread messages for all NFT related chats for given user
// @Description Get Unread count for all NFT contexts given a wallet address
// @Tags        Unused/Legacy
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       address path      string true "Wallet Address"
// @Success     200     {integer} int
// @Router      /v1/get_unread_cnt_nft/{address} [get]
func GetUnreadMsgCntNftAllByAddr(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// key := vars["address"]
	Authuser := auth.GetUserFromReqContext(r)
	key := Authuser.Address

	var chat []entity.Chatitem
	database.Connector.Where("toaddr = ?", key).Where("nftid != ?", 0).Find(&chat)

	//first we need to find unique senders (has to be a better way to use SQL db for this)
	var senderlist []string
	for i := 0; i < len(chat); i++ {
		if !stringInSlice(chat[i].Fromaddr, senderlist) {
			fmt.Printf("Found Unique Sender: %#v\n", chat[i].Fromaddr)
			senderlist = append(senderlist, chat[i].Fromaddr)
		}
	}

	//now for each sender we need get unique nft contract addresses
	var nftretval []entity.Nftsidebar

	for i := 0; i < len(senderlist); i++ {
		var senderAddr = senderlist[i]
		var chatUniqueNft []entity.Chatitem
		database.Connector.Where("toaddr = ?", key).Where("nftid != ?", 0).Where("fromaddr = ?", senderAddr).Find(&chatUniqueNft)

		var uniquecontracts []string
		for j := 0; j < len(chatUniqueNft); j++ {
			if !stringInSlice(chatUniqueNft[i].Nftaddr, uniquecontracts) {
				fmt.Printf("Found Unique NFT Contract: %#v\n", chatUniqueNft[i].Nftaddr)
				//for the given senderAddr this is unique list of contract addresses
				uniquecontracts = append(uniquecontracts, chatUniqueNft[i].Nftaddr)
			}
		}

		//now for each unqiue sender, and unique nft contract address, get unique NFT ids
		for k := 0; k < len(uniquecontracts); k++ {
			var uniqueNftAddr = uniquecontracts[k]
			var chatUniqueNftIds []entity.Chatitem
			database.Connector.Where("toaddr = ?", key).Where("nftid != ?", 0).Where("fromaddr = ?", senderAddr).Where("nftaddr = ?", uniqueNftAddr).Find(&chatUniqueNftIds)

			var uniquenftids []string
			for l := 0; l < len(chatUniqueNftIds); l++ {
				var nftid = chatUniqueNftIds[l].Nftid
				var chatNftId []entity.Chatitem
				fmt.Printf("Unique NFT ID : %#v\n", nftid)

				database.Connector.Where("toaddr = ?", key).
					Where("nftid = ?", nftid).Where("fromaddr = ?", senderAddr).
					Where("nftaddr = ?", uniqueNftAddr).
					Where("msgread = ?", false).Find(&chatNftId)

				if !stringInSlice(nftid, uniquenftids) {
					uniquenftids = append(uniquenftids, nftid)

					var sbitem entity.Nftsidebar
					sbitem.Fromaddr = senderAddr
					sbitem.Nftaddr = uniqueNftAddr
					sbitem.Nftid = nftid
					sbitem.Unread = len(chatNftId)

					nftretval = append(nftretval, sbitem)
				}
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(nftretval)
}

// GetUnreadMsgCnt godoc
// @Summary     Get all unread messages between two addresses
// @Description Get Unread count for DMs
// @Tags        Unused/Legacy
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       toaddr path      string true "TO: Wallet Address"
// @Param       from   path      string true "FROM: Wallet Address"
// @Success     200    {integer} int
// @Router      /v1/get_unread_cnt/{fromaddr}/{toaddr} [get]
func GetUnreadMsgCnt(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//to := vars["toaddr"]
	Authuser := auth.GetUserFromReqContext(r)
	to := Authuser.Address
	owner := vars["fromaddr"]

	var chat []entity.Chatitem
	database.Connector.Where("toaddr = ?", to).Where("fromaddr = ?", owner).Where("msgread != ?", true).Find(&chat)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(len(chat))
}

// GetChatFromAddress godoc
// @Summary     Get Chat Item For Given Wallet Address
// @Description Get all Chat Items for DMs for a given wallet address
// @Tags        Unused/Legacy
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       toaddr path    string true "Wallet Address"
// @Success     200    {array} entity.Chatitem
// @Router      /v1/getall_chatitems/{address} [get]
func GetChatFromAddress(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// key := vars["address"]
	Authuser := auth.GetUserFromReqContext(r)
	key := Authuser.Address

	var chat []entity.Chatitem
	database.Connector.Where("fromaddr = ?", key).Or("toaddr = ?", key).Find(&chat)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	json.NewEncoder(w).Encode(chat)
}

// GetNftChatFromAddress godoc
// @Summary     Get NFT Related Chat Items For Given Wallet Address
// @Description Get ALL NFT context items for a given wallet address
// @Tags        Unused/Legacy
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       toaddr path    string true "Wallet Address"
// @Success     200    {array} entity.Chatitem
// @Router      /v1/getnft_chatitems/{address} [get]
func GetNftChatFromAddress(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// key := vars["address"]
	Authuser := auth.GetUserFromReqContext(r)
	key := Authuser.Address

	var chat []entity.Chatitem
	database.Connector.Where("fromaddr = ?", key).Where("nftid != ?", 0).Or("toaddr = ?", key).Where("nftid != ?", 0).Find(&chat)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	json.NewEncoder(w).Encode(chat)
}

// Could combine this with GetAll of we change FE to send in 0 or something for ALL
// @Router /v1/get_n_chatitems/{fromaddr}/{toaddr}/{count} [get]
func GetNChatFromAddressToAddr(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//from := vars["fromaddr"]
	to := vars["toaddr"]
	Authuser := auth.GetUserFromReqContext(r)
	from := Authuser.Address
	count, _ := strconv.Atoi(vars["count"])

	var chat []entity.Chatitem
	database.Connector.Where("fromaddr = ?", from).Where("toaddr = ?", to).Find(&chat)

	var chat2 []entity.Chatitem
	database.Connector.Where("fromaddr = ?", to).Where("toaddr = ?", from).Find(&chat2)

	for _, chatmember := range chat2 {
		currTime := chatmember.Timestamp_dtm
		found := false
		//both lists are already sorted, so we can use the assumption here
		for i := 0; i < len(chat); i++ {
			ret_time := chat[i].Timestamp_dtm
			if currTime.Before(ret_time) {
				chat = append(chat[:i+1], chat[i:]...)
				chat[i] = chatmember
				found = true
				break
			}
		}
		if !found {
			chat = append(chat, chatmember)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	if len(chat) < count {
		json.NewEncoder(w).Encode(chat)
	} else {
		json.NewEncoder(w).Encode(chat[(len(chat) - count):])

		//since snaps only shows 6 lines (TODO fix via UI or some other method)
		//we have to mark the earlier messages as read - rare but possible situation

		//first get number of unread messages (could be 21 total convo messages, 9 unread, 6 grabbed for snap)
		var chatUnreadLength []entity.Chatitem
		database.Connector.Where("toaddr = ?", from).Where("msgread != ?", true).Find(&chatUnreadLength)
		numUnread := len(chatUnreadLength)
		msgRead := true
		//if numUnread > count {
		//simplify for now - just mark all unread as read from this specific address when using snaps
		//the above if statement works but worried there might be a race condition we miss, so this is safer for now
		//and just marks the messages read a few seconds before they would be marked read anyway in the snap
		for i := 0; i < (numUnread); i++ {
			//fmt.Println("forcing read item update at position: ", i, chatUnreadLength[i].Message)
			database.Connector.Model(&entity.Chatitem{}).
				Where("fromaddr = ?", chatUnreadLength[i].Fromaddr).
				Where("toaddr = ?", chatUnreadLength[i].Toaddr).
				Where("timestamp = ?", chatUnreadLength[i].Timestamp).
				Update("msgread", msgRead)
		}
		//}
	}
}

// GetChatFromAddressToAddr godoc
// @Summary     Get Chat Data Between Two Addresses
// @Description Get chat data between the given two addresses, TO and FROM and interchangable here
// @Tags        DMs
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       toaddr path    string true "TO: Wallet Address"
// @Param       from   path    string true "FROM: Wallet Address"
// @Success     200    {array} entity.Chatitem
// @Router      /v1/getall_chatitems/{fromaddr}/{toaddr} [get]
func GetAllChatFromAddressToAddr(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//from := vars["fromaddr"]
	to := vars["toaddr"]
	Authuser := auth.GetUserFromReqContext(r)
	from := Authuser.Address

	var chat []entity.Chatitem
	database.Connector.Where("fromaddr = ?", from).Where("toaddr = ?", to).Find(&chat)

	var chat2 []entity.Chatitem
	database.Connector.Where("fromaddr = ?", to).Where("toaddr = ?", from).Find(&chat2)

	for _, chatmember := range chat2 {
		currTime := chatmember.Timestamp_dtm
		found := false
		//both lists are already sorted, so we can use the assumption here
		for i := 0; i < len(chat); i++ {
			ret_time := chat[i].Timestamp_dtm
			if currTime.Before(ret_time) {
				chat = append(chat[:i+1], chat[i:]...)
				chat[i] = chatmember
				found = true
				break
			}
		}
		if !found {
			chat = append(chat, chatmember)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	json.NewEncoder(w).Encode(chat)
}

// GetReadChatFromAddressToAddr godoc
// @Summary     Get Recently Read Messages
// @Description Get newly read messages to update READ status for lazy loading
// @Tags        DMs
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       toaddr path    string true "TO: Wallet Address"
// @Param       from   path    string true "FROM: Wallet Address"
// @Success     200    {array} entity.Chatitem
// @Router      /v1/getread_chatitems/{fromaddr}/{toaddr} [get]
func GetReadChatFromAddressToAddr(w http.ResponseWriter, r *http.Request) {
	//we only want to return values here once (don't repeatedly report newly read messages)
	//need to keep track of current prev messages and detect changes

	vars := mux.Vars(r)
	//from := vars["fromaddr"]
	to := vars["toaddr"]

	Authuser := auth.GetUserFromReqContext(r)
	from := Authuser.Address

	var readIDs []int
	database.Connector.Model(&entity.Chatitem{}).Where("fromaddr = ?", from).Where("toaddr = ?", to).Where("msgread = ?", true).Pluck("id", &readIDs)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	json.NewEncoder(w).Encode(readIDs)
}

// GetChatFromAddressToAddr godoc
// @Summary     Get Chat Data Between Two Addresses
// @Description Get chat data between the given two addresses, TO and FROM and interchangable here
// @Tags        DMs
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       toaddr path    string true "TO: Wallet Address"
// @Param       from   path    string true "FROM: Wallet Address"
// @Param       time   path    string true "Load only messages after this time"
// @Success     200    {array} entity.Chatitem
// @Router      /v1/getall_chatitems/{fromaddr}/{toaddr}/${time} [get]
func GetNewChatFromAddressToAddr(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//from := vars["fromaddr"]
	to := vars["toaddr"]
	timeStamp := vars["time"]
	Authuser := auth.GetUserFromReqContext(r)
	from := Authuser.Address

	decodedStr, err := url.QueryUnescape(timeStamp)
	if err != nil {
		fmt.Printf("Error decoding the string %v\r\n", err)
	}
	//fmt.Printf("Input Timestamp: %v", decodedStr)

	layout := "2006-01-02T15:04:05.000Z"
	formattedTime, err := time.Parse(layout, decodedStr)
	if err != nil {
		log.Println(err)
	}
	//fmt.Printf("Formatted Timestamp: %v", formattedTime.Add(time.Second))

	//add a second to timestamp sent in, because conversion rounds up sometimes
	formattedTime = formattedTime.Add(time.Second)
	var chat []entity.Chatitem
	database.Connector.
		Where("fromaddr = ?", from).
		Where("toaddr = ?", to).
		Where("timestamp_dtm > ?", formattedTime).
		Find(&chat)

	var chat2 []entity.Chatitem
	database.Connector.
		Where("fromaddr = ?", to).
		Where("toaddr = ?", from).
		Where("timestamp_dtm > ?", formattedTime).
		Find(&chat2)

	for _, chatmember := range chat2 {
		currTime := chatmember.Timestamp_dtm
		found := false
		//both lists are already sorted, so we can use the assumption here
		for i := 0; i < len(chat); i++ {
			ret_time := chat[i].Timestamp_dtm
			if currTime.Before(ret_time) {
				chat = append(chat[:i+1], chat[i:]...)
				chat[i] = chatmember
				found = true
				break
			}
		}
		if !found {
			chat = append(chat, chatmember)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	json.NewEncoder(w).Encode(chat)
}

// GetChatNftContext godoc
// @Summary     Get NFT Related Chat Items For Given NFT Contract and ID
// @Description Get ALL NFT context items for a given wallet address
// @Tags        Unused/Legacy
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       nftaddr path    string true "NFT Contract Address"
// @Param       nftid   path    string true "NFT ID"
// @Success     200     {array} entity.Chatitem
// @Router      /v1/getnft_chatitems/{nftaddr}/{nftid} [get]
func GetChatNftContext(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nftaddr := vars["nftaddr"]
	nftid := vars["nftid"]

	var chat []entity.Chatitem
	database.Connector.Where("nftaddr = ?", nftaddr).Where("nftid = ?", nftid).Find(&chat)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	json.NewEncoder(w).Encode(chat)
}

// GetChatNftContext godoc
// @Summary     Get NFT Related Chat Items For Given NFT Contract and ID, between two wallet addresses (TO and FROM are interchangable)
// @Description Get ALL NFT context items for a specifc NFT context convo between two wallets
// @Tags        NFT
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       nftaddr path    string true "NFT Contract Address"
// @Param       nftid   path    string true "NFT ID"
// @Param       toaddr  path    string true "TO: Wallet Address"
// @Param       from    path    string true "FROM: Wallet Address"
// @Success     200     {array} entity.Chatitem
// @Router      /v1/getnft_chatitems/{fromaddr}/{toaddr}/{nftaddr}/{nftid} [get]
func GetChatNftAllItemsFromAddrAndNFT(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	from := vars["fromaddr"]

	Authuser := auth.GetUserFromReqContext(r)
	to := Authuser.Address
	//to := vars["toaddr"]
	addr := vars["nftaddr"]
	id := vars["nftid"]

	var chat []entity.Chatitem
	database.Connector.Where("fromaddr = ?", from).Where("toaddr = ?", to).Where("nftaddr = ?", addr).Where("nftid = ?", id).Find(&chat)
	//fmt.Printf("Chat Items: %#v\n", chat)

	var chat2 []entity.Chatitem
	database.Connector.Where("fromaddr = ?", to).Where("toaddr = ?", from).Where("nftaddr = ?", addr).Where("nftid = ?", id).Find(&chat2)
	//fmt.Printf("Chat2 Items: %#v\n", chat2)

	//TODO: should be a way to called a stored proc for this to sort in MySQL using timestamp
	for _, chatmember := range chat2 {
		currTime := chatmember.Timestamp_dtm
		found := false
		//both lists are already sorted, so we can use the assumption here
		for i := 0; i < len(chat); i++ {
			ret_time := chat[i].Timestamp_dtm

			if currTime.Before(ret_time) {
				chat = append(chat[:i+1], chat[i:]...)
				chat[i] = chatmember
				found = true
				break
			}
		}
		if !found {
			chat = append(chat, chatmember)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	json.NewEncoder(w).Encode(chat)
}

// GetChatNftAllItemsFromAddr godoc
// @Summary     Get NFT Related Chat Items For Given NFT Contract and ID, relating to one wallet
// @Description Get all specified NFT contract and ID items for a given wallet address
// @Tags        Unused/Legacy
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       address path    string true "Wallet Address"
// @Param       nftaddr path    string true "NFT Contract Address"
// @Param       nftid   path    string true "NFT ID"
// @Success     200     {array} entity.Chatitem
// @Router      /v1/getnft_chatitems/{address}/{nftaddr}/{nftid} [get]
func GetChatNftAllItemsFromAddr(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//walletaddr := vars["address"]
	Authuser := auth.GetUserFromReqContext(r)
	walletaddr := Authuser.Address
	addr := vars["nftaddr"]
	id := vars["nftid"]

	var chat []entity.Chatitem
	database.Connector.Where("fromaddr = ?", walletaddr).Where("nftaddr = ?", addr).Where("nftid = ?", id).Find(&chat)

	var chat2 []entity.Chatitem
	database.Connector.Where("toaddr = ?", walletaddr).Where("nftaddr = ?", addr).Where("nftid = ?", id).Find(&chat2)

	//TODO, should do this and similar sorts in a stored proc probably which sort (call 2 queries above with and ORDER)
	for _, chatmember := range chat2 {
		currTime := chatmember.Timestamp_dtm
		found := false
		//both lists are already sorted, so we can use the assumption here
		for i := 0; i < len(chat); i++ {
			ret_time := chat[i].Timestamp_dtm
			if currTime.Before(ret_time) {
				chat = append(chat[:i+1], chat[i:]...)
				chat[i] = chatmember
				found = true
				break
			}
		}
		if !found {
			chat = append(chat, chatmember)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	json.NewEncoder(w).Encode(chat)
}

func findStrIndexInArray(s string, arr []string) int {
	for i, item := range arr {
		if item == s {
			return i
		}
	}
	return -1 // Return -1 if the string is not found in the slice.
}

// CreateChatitem godoc
// @Summary     Create/Insert DM Chat Message (1-to-1 messaging)
// @Description For DMs, Chatitem data struct is used to store each message and associated info.
// @Description REQUIRED: fromaddr, toaddr, message (see data struct section at bottom of page for more detailed info on each paramter)
// @Description Other fields are generally filled in by the backed REST API and used as return parameters
// @Description ID is auto generated and should never be used as input.
// @Tags        DMs
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       message body    entity.Chatitem true "Direct Message Chat Data"
// @Success     200     {array} entity.Chatitem
// @Router      /v1/create_chatitem [post]
func CreateChatitem(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var chat entity.Chatitem
	json.Unmarshal(requestBody, &chat)

	//increment our  giveawaymessageCounter
	giveawayMessageCounter = giveawayMessageCounter + 1

	//added this because from API doc it was throwing error w/o this
	//TODO: we should sort out if we really need this as an input or output only
	chat.Timestamp = time.Now().Format("2006-01-02T15:04:05.000Z")
	//I think can remove this too since Oliver added a DB trigger
	chat.Timestamp_dtm = time.Now()

	//ensure user in body is same as user in JWT
	Authuser := auth.GetUserFromReqContext(r)
	walletaddr := Authuser.Address

	if strings.EqualFold(walletaddr, chat.Fromaddr) {
		//ensure its not from a blocked user
		var blockedUser entity.Blockeduser
		blockedQuery := database.Connector.Where("blockedaddress = ?", chat.Fromaddr).Where("owneraddress = ?", chat.Toaddr).Find(&blockedUser)
		if blockedQuery.RowsAffected > 0 {
			fmt.Println("block user trying to send messages: ", chat.Fromaddr, chat.Toaddr)
			w.WriteHeader(http.StatusForbidden)
			return
		}

		dbQuery := database.Connector.Create(&chat)
		if dbQuery.RowsAffected == 0 {
			fmt.Println(dbQuery.Error)
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("X-Content-Type-Options", "nosniff")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(dbQuery.Error)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("X-Content-Type-Options", "nosniff")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(chat)

			//fmt.Println("message check for support", tgSupportWalletsCsvString, tgSupporChatIdsCsvString, chat.Toaddr)
			if strings.Contains(tgSupportWalletsCsvString, chat.Toaddr) {
				//first find the index - we could have this be the initial check, but might be slow if the list gets long
				index := findStrIndexInArray(chat.Toaddr, tgSupportWalletArray)
				chat_id := tgSupportChatIdsArray[index]

				blockedUserIndex := findStrIndexInArray(strings.ToLower(chat.Fromaddr), tgSupportBlockedUsers)

				if blockedUserIndex < 0 {
					//fmt.Println("sending to TG group message")
					var fromAddrSettings entity.Settings
					database.Connector.Where("walletaddr = ?", chat.Fromaddr).Find(&fromAddrSettings)
					SendTelegramMessage("_WalletChat Ticket #("+strconv.Itoa(fromAddrSettings.ID)+")_\r\n"+chat.Message, chat_id)
				}
			}
			//manage support messages
			// if strings.EqualFold(os.Getenv("SUPPORT_WALLET"), chat.Toaddr) {
			// 	url := os.Getenv("SUPPORT_WEBOOK_URL")
			// 	messageToWebhook := "From: " + chat.Fromaddr + " Messge: " + chat.Message
			// 	method := "POST"

			// 	jsonBody := `{"content":"` + messageToWebhook + `"}`
			// 	payload := strings.NewReader(jsonBody)

			// 	client := &http.Client{}
			// 	req, err := http.NewRequest(method, url, payload)

			// 	if err != nil {
			// 		fmt.Println(err)
			// 		return
			// 	}
			// 	req.Header.Add("Content-Type", "application/json")

			// 	res, err := client.Do(req)
			// 	if err != nil {
			// 		fmt.Println(err)
			// 		return
			// 	}
			// 	defer res.Body.Close()

			// 	body, err := ioutil.ReadAll(res.Body)
			// 	if err != nil {
			// 		fmt.Println(err)
			// 		return
			// 	}
			// 	fmt.Println(string(body))
			// }

			//send DM that the user should claim the prize on Twitter
			// if giveawayMessageCounter%50 == 0 {
			// 	var chat entity.Chatitem
			// 	chat.Timestamp = time.Now().Format("2006-01-02T15:04:05.000Z")
			// 	chat.Timestamp_dtm = time.Now()
			// 	chat.Fromaddr = strings.ToLower(os.Getenv("SUPPORT_WALLET"))
			// 	chat.Toaddr = strings.ToLower(walletaddr)
			// 	chat.Nftid = "0"
			// 	//auto-send a message to the user to check out the leaderboard
			// 	chat.Message = "Tweet @wallet_chat with #chat2earn and #chat2win and be creative!\nThe tweet of the day with most impressions wins 5 USDC!"
			// 	database.Connector.Create(&chat)
			// 	fmt.Println("$$$ Prize Print for - JWT Address: ", Authuser.Address)

			// 	//send an email that the user should claim the prize on Twitter - TBD need new email template
			// 	var settings entity.Settings
			// 	var dbResult = database.Connector.Where("walletaddr = ?", walletaddr).Find(&settings)
			// 	if dbResult.RowsAffected > 0 {
			// 		fmt.Println("$$$ Prize Email for: ", walletaddr, settings.Email)
			// 		if strings.Contains(settings.Email, "@") {
			// 			var fromAddrname entity.Addrnameitem
			// 			database.Connector.Where("address = ?", walletaddr).Find(&fromAddrname)

			// 			from := mail.NewEmail("WalletChat Prize Notifications", "contact@walletchat.fun")
			// 			subject := "Prize Notfication For WalletChat!"
			// 			to := mail.NewEmail(fromAddrname.Name, settings.Email)
			// 			plainTextContent := "You have a message from" + fromAddrname.Name + " : \r\n" + chat.Message + "\r\n Please login via the app at https://app.walletchat.fun to read!"
			// 			htmlContent := email.NotificationEmailTwitter(walletaddr, fromAddrname.Name, settings.Email)
			// 			message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
			// 			client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
			// 			response, err := client.Send(message)
			// 			if err != nil {
			// 				log.Println(err)
			// 			} else {
			// 				_ = response
			// 			}
			// 		}
			// 	}
			// }

			//also notify the TO user of a new message (need to throttle this somehow)
			var settings entity.Settings
			var dbResult = database.Connector.Where("walletaddr = ?", chat.Toaddr).Find(&settings)
			wc_analytics.SendCustomEventWithSignupSite(Authuser.Address, "SEND_MESSAGE", settings.Signupsite)
			if dbResult.RowsAffected > 0 && strings.EqualFold("true", settings.Verified) {
				if strings.Contains(settings.Email, "@") && strings.EqualFold(settings.Notifydm, "true") {
					var fromAddrname entity.Addrnameitem
					database.Connector.Where("address = ?", chat.Fromaddr).Find(&fromAddrname)
					var toAddrname entity.Addrnameitem
					database.Connector.Where("address = ?", chat.Toaddr).Find(&toAddrname)

					from := mail.NewEmail("WalletChat Notifications", "contact@walletchat.fun")
					subject := "Message Waiting In WalletChat"
					to := mail.NewEmail(toAddrname.Name, settings.Email)
					plainTextContent := "You have a message from" + fromAddrname.Name + " : \r\n" + chat.Message + "\r\n Please login via the app at https://app.walletchat.fun to read!"
					htmlContent := email.NotificationEmailDM(toAddrname.Address, fromAddrname.Address, toAddrname.Name, fromAddrname.Name, settings.Email, chat.Message)
					message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
					client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
					response, err := client.Send(message)
					if err != nil {
						log.Println(err)
					} else {
						_ = response
					}
				}
			}
			if dbResult.RowsAffected > 0 && settings.Telegramid != "" {
				if strings.EqualFold(settings.Notifydm, "true") {
					var fromAddrname entity.Addrnameitem
					database.Connector.Where("address = ?", chat.Fromaddr).Find(&fromAddrname)

					var message string = "You have a message waiting in WalletChat from: " + fromAddrname.Name + "(" + fromAddrname.Address + ")"
					SendTelegramMessage(message, settings.Telegramid)
				}
			}
		}
	} else {
		fmt.Println("create_chatitem - JWT Address: ", Authuser.Address)
		fmt.Println("create_chatitem - POST Address: ", chat.Fromaddr)
		w.WriteHeader(http.StatusForbidden)
	}
}

func SendNotificationEmails() {
	fmt.Println("** Sending Daily Notifications **")
	var settings []entity.Settings
	database.Connector.Find(&settings)
	for i := 0; i < len(settings); i++ {
		config := LocalGetUnread(settings[i].Walletaddr)
		if config.Dm > 0 {
			var addrnameDB entity.Addrnameitem
			var dbQuery = database.Connector.Where("address = ?", settings[i].Walletaddr).Find(&addrnameDB)

			if dbQuery.RowsAffected > 0 && strings.EqualFold(settings[i].Notify24, "true") && strings.EqualFold("true", settings[i].Verified) {
				from := mail.NewEmail("WalletChat Notifications", "contact@walletchat.fun")
				subject := "Message Waiting In WalletChat"
				to := mail.NewEmail(addrnameDB.Name, settings[i].Email)
				plainTextContent := "You have " + strconv.Itoa(config.Dm) + " unread DM(s), " + strconv.Itoa(config.Nft) + " unread NFT group chat messages, and " + strconv.Itoa(config.Community) + " unread custom community chat messages waiting in WalletChat. Please login via the app at https://app.walletchat.fun to read!"
				htmlContent := email.NotificationEmail24(addrnameDB.Address, addrnameDB.Name, strconv.Itoa(config.Dm), strconv.Itoa(config.Nft), strconv.Itoa(config.Community), settings[i].Email)
				message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
				client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
				response, err := client.Send(message)
				if err != nil {
					log.Println(err)
				} else {
					_ = response
				}
			}
		}
	}
}

func getTelegrameUrl() string {
	return fmt.Sprintf("https://api.telegram.org/bot%s", os.Getenv("TELEGRAM_BOT_TOKEN"))
}
func SendTelegramMessage(text string, chatId string) (bool, error) {
	// Global variables
	var err error
	var response *http.Response

	// Send the message
	url := fmt.Sprintf("%s/sendMessage", getTelegrameUrl())
	body, _ := json.Marshal(map[string]string{
		"chat_id":    chatId,
		"text":       text,
		"parse_mode": "Markdown",
	})
	response, err = http.Post(
		url,
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return false, err
	}

	// Close the request at the end
	defer response.Body.Close()

	// Return
	return true, nil
}

func isFieldSet(i interface{}) bool {
	return !reflect.ValueOf(i).IsNil()
}

// func extractAddress(input string) string {
// 	// Define a regular expression to match the text after "WalletChat User: "
// 	re := regexp.MustCompile(`WalletChat User: (.*)`)

// 	// Find the first match in the input string
// 	match := re.FindStringSubmatch(input)

// 	if len(match) >= 2 {
// 		// Extract the text after the delimiter
// 		return strings.TrimSpace(match[1])
// 	}

//		return ""
//	}
func extractNumber(input string) string {
	re := regexp.MustCompile(`\((\d+)\)`)
	match := re.FindStringSubmatch(input)
	if len(match) != 2 {
		return "0"
	}
	return match[1]
}

// TODO: should be done by webhook eventually so we don't have to loop, and can do additional verifications
func UpdateTelegramNotifications() {
	//poll for new users setting up telegram notifications (can be a webhook someday for better performance)
	var err error
	var response *http.Response

	// Send the message
	url := fmt.Sprintf("%s/getUpdates?offset=%d", getTelegrameUrl(), telegramUpdateOffset)
	response, err = http.Get(
		url,
	)
	if err != nil {
		fmt.Println("update telegram error", err)
		return
	}

	// Close the request at the end
	defer response.Body.Close()

	// Body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("update telegram error", err)
	}

	var updatedNotifsData TelegramUpdateNotifsData
	json.Unmarshal(body, &updatedNotifsData)

	for i := 0; i < len(updatedNotifsData.Result); i++ {
		//fmt.Println("full message: ", updatedNotifsData.Result[i])
		if isFieldSet(updatedNotifsData.Result[i].Message.ReplyToMessage) {
			//if its a reply message, we need to send the user the reply (but only permissioned admins can do this)
			//TODO: need a different permission list per TG group
			//fmt.Println("response permissions (fromID) (adminsArray)", strconv.FormatInt(updatedNotifsData.Result[i].Message.From.ID, 10), tgSupportAdminsArray)
			if findStrIndexInArray(strconv.FormatInt(updatedNotifsData.Result[i].Message.From.ID, 10), tgSupportAdminsArray) > -1 {
				//TODO: make this an actual ticketing system instead, not just one id
				var origSenderSettings entity.Settings
				settingsIdMsgSender := extractNumber(updatedNotifsData.Result[i].Message.ReplyToMessage.Text)
				database.Connector.Where("id = ?", settingsIdMsgSender).Find(&origSenderSettings)
				origMsgSender := origSenderSettings.Walletaddr
				//fmt.Println("GD Admin Replied To Message from / with:", origMsgSender, updatedNotifsData.Result[i].Message)

				//see if its a block user request
				if strings.EqualFold(updatedNotifsData.Result[i].Message.Text, "BLOCK_USER") {
					fmt.Println("TG Admin Blocked User: ", origMsgSender)
					tgSupportBlockedUsers = append(tgSupportBlockedUsers, strings.ToLower(origMsgSender))
					telegramUpdateOffset = updatedNotifsData.Result[i].UpdateID + 1
					return
				}

				//find corresponding support wallet for given chat_id
				indexOfChatId := findStrIndexInArray(strconv.Itoa(updatedNotifsData.Result[i].Message.ReplyToMessage.Chat.ID), tgSupportChatIdsArray)

				if indexOfChatId > -1 {
					var chat entity.Chatitem
					chat.Timestamp = time.Now().Format("2006-01-02T15:04:05.000Z")
					chat.Timestamp_dtm = time.Now()
					chat.Fromaddr = tgSupportWalletArray[indexOfChatId]
					chat.Toaddr = origMsgSender
					chat.Message = updatedNotifsData.Result[i].Message.Text
					chat.Nftid = "0"
					//fmt.Println("creating TG response chat item", chat)
					database.Connector.Create(&chat)
				}
			}

		} else {
			//fmt.Println("Results For Telegram Check", updatedNotifsData.Result[i])
			verifCode := updatedNotifsData.Result[i].Message.Text
			//fmt.Println("Results For Telegram Verification Code: ", verifCode)
			var settings []entity.Settings
			dbResult := database.Connector.Where("telegramcode = ?", verifCode).Find(&settings)
			if dbResult.RowsAffected > 0 && verifCode != "" {
				chatId := updatedNotifsData.Result[i].Message.Chat.ID

				fmt.Println("Updating Telegram Chat ID for WalletAddr/chatID: ", settings[0].Walletaddr, strconv.FormatInt(chatId, 10))
				database.Connector.Model(&entity.Settings{}).Where("walletaddr = ?", settings[0].Walletaddr).Update("telegramid", strconv.FormatInt(chatId, 10))
				database.Connector.Model(&entity.Settings{}).Where("walletaddr = ?", settings[0].Walletaddr).Update("telegramcode", "")

				var message string = "You have successfully setup notifications in WalletChat for: " + settings[0].Walletaddr
				SendTelegramMessage(message, strconv.FormatInt(chatId, 10))
			}
		}
		// Update the offset to the next update ID
		//fmt.Println("updated offset TG: ", telegramUpdateOffset)
		telegramUpdateOffset = updatedNotifsData.Result[i].UpdateID + 1
	}
}

// CreateGroupChatitem godoc
// @Summary     Create/Insert chat message for NFT Group Messaging
// @Description Currently used for NFT Gated Chats
// @Tags        GroupChat
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       message body    entity.Groupchatitem true "NFT Group Message Chat Data"
// @Success     200     {array} entity.Groupchatitem
// @Router      /v1/create_groupchatitem [post]
func CreateGroupChatitem(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var chat entity.Groupchatitem
	json.Unmarshal(requestBody, &chat)

	//these will get overwritten as needed when returning data
	chat.Contexttype = entity.Nft
	chat.Type = entity.Message

	//probably can removed now with DB trigger
	chat.Timestamp = time.Now().Format("2006-01-02T15:04:05.000Z")
	chat.Timestamp_dtm = time.Now()

	Authuser := auth.GetUserFromReqContext(r)

	//ensure user holds the NFT first
	isHolder := false
	if strings.HasPrefix(chat.Nftaddr, "0x") {
		//TODO: we should send in chain along with message
		isHolder = IsOwnerOfNFT(chat.Nftaddr, chat.Fromaddr, "ethereum")
		if !isHolder {
			isHolder = IsOwnerOfNFT(chat.Nftaddr, chat.Fromaddr, "polygon")
		}
	} else if !isHolder && (strings.HasSuffix(chat.Fromaddr, ".near") || strings.HasSuffix(chat.Fromaddr, ".testnet")) ||
		(len(chat.Fromaddr) == 64 && !strings.HasPrefix(chat.Fromaddr, "0x")) { //NEAR check
		isHolder = IsOwnerOfNFT(chat.Nftaddr, chat.Fromaddr, "near")
	} else if !isHolder && strings.HasPrefix(chat.Fromaddr, "tz") { //Tezos check
		isHolder = IsOwnerOfNFT(chat.Nftaddr, chat.Fromaddr, "tezos")
	} else if !isHolder && strings.HasPrefix(chat.Nftaddr, "poap_") {
		split := strings.Split(chat.Nftaddr, "_")
		isHolder = IsOwnerOfPOAP(split[1], chat.Fromaddr)
	}

	if strings.EqualFold(chat.Fromaddr, Authuser.Address) && isHolder {
		//public chats are not encrpyted and we implement a basic censor
		cleanMessage := goaway.Censor(chat.Message)
		chat.Message = cleanMessage

		database.Connector.Create(&chat)

		wc_analytics.SendCustomEvent(Authuser.Address, "SEND_MESSAGE_NFTGROUP")

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(chat)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}

// ChangeCommunityConditions godoc
// @Summary     Change community access conditions
// @Description Change community access conditions
// @Tags        GroupChat
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       message body    entity.Communityaccesscondition true "Create/Edit Community/Group Access Conditions"
// @Success     200     {array} entity.Communityaccesscondition
// @Router      /v1/community/conditions [post]
func ChangeCommunityConditions(w http.ResponseWriter, r *http.Request) {
	//todo - should maybe just accept JSON and process later.
	requestBody, _ := ioutil.ReadAll(r.Body)
	var accessCondition entity.Communityaccesscondition
	json.Unmarshal(requestBody, &accessCondition)
	fmt.Println("ChangeCommunityConditions", accessCondition)

	Authuser := auth.GetUserFromReqContext(r)

	//ensure the caller is an admin for the group
	var adminForCommunity []entity.Communityadmin
	dbQuery := database.Connector.Where("adminaddr = ?", Authuser.Address).Find(&adminForCommunity)
	isAdmin := false
	for i := 0; i < int(dbQuery.RowsAffected); i++ {
		if adminForCommunity[i].Slug == accessCondition.Slug {
			isAdmin = true
			break
		}
	}
	if !isAdmin {
		fmt.Println("ChangeCommunityConditions not an admin", accessCondition.Slug)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	var accessConditionToUpdate entity.Communityaccesscondition
	dbQuery = database.Connector.Where("slug = ?", accessCondition.Slug).Find(&accessConditionToUpdate)
	resultCnt := 0
	if dbQuery.RowsAffected == 0 {
		dbQuery = database.Connector.Create(&accessCondition)
		resultCnt += int(dbQuery.RowsAffected)
	} else {
		dbQuery = database.Connector.Model(&entity.Communityaccesscondition{}).
			Where("slug = ?", accessCondition.Slug).
			Update("nftaddr", accessCondition.Nftaddr)
		//fmt.Println("ChangeCommunityConditions 1", accessCondition, dbQuery.RowsAffected)
		resultCnt += int(dbQuery.RowsAffected)

		dbQuery = database.Connector.Model(&entity.Communityaccesscondition{}).
			Where("slug = ?", accessCondition.Slug).
			Update("count", accessCondition.Count)
		//fmt.Println("ChangeCommunityConditions 2", accessCondition, dbQuery.RowsAffected)
		resultCnt += int(dbQuery.RowsAffected)
	}

	if resultCnt > 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
	} else {
		fmt.Println("failed ChangeCommunityConditions", dbQuery.RowsAffected)
		w.WriteHeader(http.StatusForbidden)
	}
}

// CreateCommunity godoc
// @Summary     CreateCommunity creates new custom community chat
// @Description Community Chat Creation
// @Tags        GroupChat
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       message body    entity.Createcommunityitem true "Community/Group Creation"
// @Success     200     {array} entity.Createcommunityitem
// @Router      /v1/create_community [post]
func CreateCommunity(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var communityInfo entity.Createcommunityitem
	if err := json.Unmarshal(requestBody, &communityInfo); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON in CreateCommunityChat", requestBody)
	}

	Authuser := auth.GetUserFromReqContext(r)

	//don't want groups named "new" as that is the path for creating new groups
	if strings.EqualFold(communityInfo.Name, "new") {
		fmt.Println("blacklisted community name", communityInfo.Name)
		w.WriteHeader(http.StatusForbidden)
		return
	}
	fmt.Println("input data create community: ", communityInfo)

	var addrname entity.Addrnameitem
	addrname.Name = communityInfo.Name //communityInfo.Title
	//auto-generate the slug (unique URL safe group name)
	slug := url.QueryEscape(addrname.Name)
	fmt.Println("Create Community Slug: ", slug)
	addrname.Address = slug //Slug

	var mappings []entity.Addrnameitem
	dbQuery := database.Connector.Where("address = ?", addrname.Address).Find(&mappings)
	//currently, community chat is in the addrname mapping table in the DB
	for i := 0; i < 100; i++ {
		if dbQuery.RowsAffected == 0 {
			database.Connector.Create(&addrname)
			break
		}
		addrname.Address = addrname.Address + "_" + strconv.Itoa(i)
		dbQuery = database.Connector.Where("address = ?", addrname.Address).Find(&mappings)
	}

	if communityInfo.Image != "" {
		var imageaddr entity.Imageitem
		imageaddr.Addr = slug
		imageaddr.Base64data = communityInfo.Image
		database.Connector.Create(&imageaddr)
	}

	for i := 0; i < len(communityInfo.Social); i++ {
		var social entity.Communitysocial
		social.Community = addrname.Address //slug
		social.Type = communityInfo.Social[i].Type
		social.Name = communityInfo.Social[i].Name
		database.Connector.Create(&social)
	}

	//currently the community creator is admin (lots of room for progress)
	var groupadmin entity.Communityadmin
	groupadmin.Adminaddr = Authuser.Address
	groupadmin.Slug = addrname.Address //slug
	groupadmin.Accesslevel = "admin"   //eventually an enum - admin/moderator/?
	dbQuery = database.Connector.Create(&groupadmin)
	//TODO: do we need to limit the people that can create community chat groups?
	//Authuser := auth.GetUserFromReqContext(r)
	//if strings.EqualFold(chat.Fromaddr, Authuser.Address) {
	if dbQuery.RowsAffected != 0 {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(slug)
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}

// UpdateCommunity godoc
// @Summary     UpdateCommunity updates  custom community chat
// @Description Community Chat Update - input slug, and any updates to Name, Socials
// @Tags        GroupChat
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       message body    entity.Createcommunityitem true "Community/Group Update"
// @Success     200     {array} entity.Createcommunityitem
// @Router      /v1/update_community [post]
func UpdateCommunity(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var communityInfo entity.Createcommunityitem
	if err := json.Unmarshal(requestBody, &communityInfo); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON in CreateCommunityChat", requestBody)
	}

	Authuser := auth.GetUserFromReqContext(r)

	//don't want groups named "new" as that is the path for creating new groups
	if strings.EqualFold(communityInfo.Slug, "new") {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	fmt.Println("input data update community: ", communityInfo)

	var groupadmin entity.Communityadmin
	database.Connector.Where("slug = ?", communityInfo.Slug).Find(&groupadmin)

	if groupadmin.Accesslevel == "admin" && strings.EqualFold(groupadmin.Adminaddr, Authuser.Address) {
		var mappings []entity.Addrnameitem
		database.Connector.Where("address = ?", communityInfo.Slug).Find(&mappings)

		//for update we modify the common name if its different
		database.Connector.Model(&entity.Addrnameitem{}).
			Where("address = ?", communityInfo.Slug).
			Update("name", communityInfo.Name)

		//delete all communitysocials and just add back in what is passsed in, this allows for deletion
		var socialsToDelete []entity.Communitysocial
		database.Connector.Where("community = ?", communityInfo.Slug).Find(&socialsToDelete)
		for i := 0; i < len(socialsToDelete); i++ {
			database.Connector.Delete(&socialsToDelete[i])
		}

		for i := 0; i < len(communityInfo.Social); i++ {
			var social entity.Communitysocial
			social.Community = communityInfo.Slug
			social.Type = communityInfo.Social[i].Type
			social.Name = communityInfo.Social[i].Name
			database.Connector.Create(&social)
		}

		if communityInfo.Image != "" {
			var imageaddr entity.Imageitem
			imageaddr.Addr = communityInfo.Slug
			imageaddr.Base64data = communityInfo.Image

			//just delete what is there, if anything is there. (saves checking what exists to do create or update)
			database.Connector.Where("addr = ?", communityInfo.Slug).Delete(&imageaddr)

			database.Connector.Create(&imageaddr)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}

// CreateCommunityChatitem godoc
// @Summary     CreateCommunityChatitem creates GroupChatitem just with community tag (likely could be consolidated)
// @Description Community Chat Data
// @Tags        GroupChat
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       message body    entity.Groupchatitem true "Community Message Chat Data"
// @Success     200     {array} entity.Groupchatitem
// @Router      /v1/community [post]
func CreateCommunityChatItem(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var chat entity.Groupchatitem
	if err := json.Unmarshal(requestBody, &chat); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON in CreateCommunityChat", requestBody)
	}

	//set type (could hack this in GET side but this is probably cleaner?)
	if chat.Type != entity.Welcome {
		chat.Type = entity.Message
	}

	chat.Timestamp = time.Now().Format("2006-01-02T15:04:05.000Z")
	chat.Timestamp_dtm = time.Now()

	Authuser := auth.GetUserFromReqContext(r)
	if strings.EqualFold(chat.Fromaddr, Authuser.Address) {
		cleanMessage := goaway.Censor(chat.Message)
		chat.Message = cleanMessage

		database.Connector.Create(&chat)

		wc_analytics.SendCustomEvent(Authuser.Address, "SEND_MESSAGE_COMMUNITY")

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(chat)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}

// CreateBookmarkItem godoc
// @Summary     Join an NFT or Community group chat
// @Description Bookmarks keep an NFT/Community group chat in the sidebar
// @Tags        GroupChat
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       message body    entity.Bookmarkitem true "Add Bookmark from Community Group Chat"
// @Success     200     {array} entity.Bookmarkitem
// @Router      /v1/create_bookmark [post]
func CreateBookmarkItem(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var bookmark entity.Bookmarkitem
	json.Unmarshal(requestBody, &bookmark)

	Authuser := auth.GetUserFromReqContext(r)

	if strings.EqualFold(bookmark.Walletaddr, Authuser.Address) {
		//fmt.Printf("Bookmark Item: %#v\n", chat)
		bookmark.Chain = "none"

		if strings.HasPrefix(bookmark.Nftaddr, "poap_") {
			bookmark.Chain = "xdai"
		} else {
			var result = IsOnChain(bookmark.Nftaddr, "ethereum")
			if result {
				bookmark.Chain = "ethereum"
			} else {
				var result = IsOnChain(bookmark.Nftaddr, "polygon")
				if result {
					bookmark.Chain = "polygon"
				}
			}
		}

		database.Connector.Create(&bookmark)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(bookmark)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}

// DeleteBookmarkItem godoc
// @Summary     Leave an NFT or Community group chat
// @Description Bookmarks keep an NFT/Community group chat in the sidebar
// @Tags        GroupChat
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       message body    entity.Bookmarkitem true "Remove Bookmark from Community Group Chat"
// @Success     200     {array} entity.Bookmarkitem
// @Router      /v1/delete_bookmark [post]
func DeleteBookmarkItem(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var bookmark entity.Bookmarkitem
	json.Unmarshal(requestBody, &bookmark)

	Authuser := auth.GetUserFromReqContext(r)

	if strings.EqualFold(bookmark.Walletaddr, Authuser.Address) {
		var success = database.Connector.Where("nftaddr = ?", bookmark.Nftaddr).Where("walletaddr = ?", bookmark.Walletaddr).Delete(bookmark)

		var returnval bool
		if success.RowsAffected > 0 {
			returnval = true
		}

		//set the fact the user has manually unjoined this NFT
		var tempUserUnjoined entity.Userunjoined
		var checkUser = database.Connector.Where("nftaddr = ?", bookmark.Nftaddr).Where("walletaddr = ?", bookmark.Walletaddr).Find(&tempUserUnjoined)

		if checkUser.RowsAffected > 0 {
			database.Connector.Model(&entity.Userunjoined{}).
				Where("walletaddr = ?", bookmark.Walletaddr).
				Where("nftaddr = ?", bookmark.Nftaddr).
				Update("unjoined", true)
		} else {
			tempUserUnjoined.Nftaddr = bookmark.Nftaddr
			tempUserUnjoined.Walletaddr = bookmark.Walletaddr
			tempUserUnjoined.Unjoined = true
			database.Connector.Create(&tempUserUnjoined)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(returnval)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}

// IsBookmarkItem godoc
// @Summary     Check if a wallet address has bookmarked/joined given NFT contract
// @Description This used for UI purposes, checking if a user/wallet has bookmarked a community.
// @Tags        GroupChat
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       walletaddr path   string true "Wallet Address"
// @Param       nftaddr    path   string true "NFT Contract Address"
// @Success     200        {bool} bool
// @Router      /v1/get_bookmarks/{walletaddr}/{nftaddr} [get]
func IsBookmarkItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//walletaddr := vars["walletaddr"]
	nftaddr := vars["nftaddr"]

	Authuser := auth.GetUserFromReqContext(r)
	walletaddr := Authuser.Address

	var bookmarks []entity.Bookmarkitem
	database.Connector.Where("walletaddr = ?", walletaddr).Where("nftaddr = ?", nftaddr).Find(&bookmarks)

	var returnval bool
	if len(bookmarks) > 0 {
		returnval = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(returnval)
}

// GetBookmarkItems godoc
// @Summary     Check if a wallet address has bookmarked/joined given NFT contract
// @Description This used for UI purposes, checking if a user/wallet has bookmarked a community.
// @Tags        GroupChat
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       address path    string true "Wallet Address"
// @Success     200     {array} entity.Bookmarkitem
// @Router      /v1/get_bookmarks/{address}/ [get]
func GetBookmarkItems(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// key := vars["address"]
	Authuser := auth.GetUserFromReqContext(r)
	key := Authuser.Address

	var bookmarks []entity.Bookmarkitem
	database.Connector.Where("walletaddr = ?", key).Find(&bookmarks)

	//now add last message from group chat this bookmark is for
	var returnItems []entity.BookmarkReturnItem
	var chat entity.Groupchatitem
	for i := 0; i < len(bookmarks); i++ {
		chat.Message = ""
		//chat.Timestamp
		database.Connector.Where("nftaddr = ?", bookmarks[i].Nftaddr).Find(&chat)

		//get num unread messages
		var chatCnt []entity.Groupchatitem
		var chatReadTime entity.Groupchatreadtime
		var dbQuery = database.Connector.Where("fromaddr = ?", key).Find(&chatReadTime)
		//if no respsonse to this query, its the first time a user is reading the chat history, send it all
		if dbQuery.RowsAffected == 0 {
			database.Connector.Where("nftaddr = ?", chat.Nftaddr).Find(&chatCnt)
		} else {
			database.Connector.Where("timestamp_dtm > ?", chatReadTime.Readtimestamp_dtm).Where("nftaddr = ?", chat.Nftaddr).Find(&chatCnt)
		}
		//end get num unread messages

		var returnItem entity.BookmarkReturnItem
		returnItem.Id = chat.Id
		returnItem.Lastmsg = chat.Message
		returnItem.Lasttimestamp = chat.Timestamp
		returnItem.Lasttimestamp_dtm = chat.Timestamp_dtm
		returnItem.Nftaddr = bookmarks[i].Nftaddr
		returnItem.Walletaddr = bookmarks[i].Walletaddr
		returnItem.Unreadcnt = len(chatCnt)
		if returnItem.Lastmsg == "" {
			var unsetTimeDtm time.Time
			var unsetTime string
			var noInt int
			returnItem.Unreadcnt = noInt
			returnItem.Lasttimestamp = unsetTime
			returnItem.Lasttimestamp_dtm = unsetTimeDtm
		}
		returnItems = append(returnItems, returnItem)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	json.NewEncoder(w).Encode(returnItems)
}

// CreateImageItem godoc
// @Summary     Store Image in SQL DB for later user
// @Description Currently used for the WC HQ Logo, stores the base64 raw data of the profile image for a community
// @Tags        Common
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       message body    entity.Imageitem true "Profile Thumbnail Pic"
// @Success     200     {array} entity.Imageitem
// @Router      /v1/image [post]
func CreateImageItem(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var imageaddr entity.Imageitem
	json.Unmarshal(requestBody, &imageaddr)

	Authuser := auth.GetUserFromReqContext(r)

	if strings.EqualFold(Authuser.Address, imageaddr.Addr) {
		database.Connector.Create(&imageaddr)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(imageaddr)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}

// CreateRawImageItem godoc
// @Summary     Store Image in Bucket Storage
// @Description private image storage for photo uploads in DMS
// @Description imageid should follow format: <fromAddr>_<toAddr>_<random 10 digit number>
// @Description the random number is passed in instead of returned, to make it easier for the FE to save this in message data
// @Tags        DMs
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       message body entity.ImageitemPhoto true "Raw Images"
// @Success     200
// @Router      /v1/imageraw [post]
func CreateRawImageItem(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var imageaddr entity.ImageitemPhoto
	err := json.Unmarshal(requestBody, &imageaddr)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	Authuser := auth.GetUserFromReqContext(r)

	//only allow users to create images for themselves in DMs
	if strings.Contains(strings.ToLower(imageaddr.Imageid), Authuser.Address) {
		// Step 2: Define the parameters for the session you want to create.
		key := "DO00CLQBPDAEHFUTYMGR"        // Access key pair. You can create access key pairs using the control panel or API.
		secret := os.Getenv("SPACES_SECRET") // Secret access key defined through an environment variable.

		s3Config := &aws.Config{
			Credentials:      credentials.NewStaticCredentials(key, secret, ""), // Specifies your credentials.
			Endpoint:         aws.String("https://sgp1.digitaloceanspaces.com"), // Find your endpoint in the control panel, under Settings. Prepend "https://".
			S3ForcePathStyle: aws.Bool(false),                                   // // Configures to use subdomain/virtual calling format. Depending on your version, alternatively use o.UsePathStyle = false
			Region:           aws.String("us-east-1"),                           // Must be "us-east-1" when creating new Spaces. Otherwise, use the region in your endpoint, such as "nyc3".
		}

		// Step 3: The new session validates your request and directs it to your Space's specified endpoint using the AWS SDK.
		newSession, errs := session.NewSession(s3Config)
		if errs != nil {
			fmt.Println(errs.Error())
		}
		s3Client := s3.New(newSession)

		// Step 4: Define the parameters of the object you want to upload.
		object := s3.PutObjectInput{
			Bucket: aws.String("walletchat-pfp-storage"),    // The path to the directory you want to upload the object to, starting with your Space name.
			Key:    aws.String(imageaddr.Imageid),           // Object key, referenced whenever you want to access this file later.
			Body:   strings.NewReader(imageaddr.Base64data), // The object's contents.
			ACL:    aws.String("private"),                   // Defines Access-control List (ACL) permissions, such as private or public.
			Metadata: map[string]*string{ // Required. Defines metadata tags.
				"x-amz-meta-my-key": aws.String("your-value"),
			},
		}

		// Step 5: Run the PutObject function with your parameters, catching for errors.
		_, err := s3Client.PutObject(&object)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusCreated)
		}
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}

// CreatePublicImageItem godoc
// @Summary     Store Image in Bucket Storage
// @Description public image storage for photo uploads, still use GUID or random string to make it link-only private
// @Description imageid should follow format: <random 10 digit number>
// @Description the random number is passed in instead of returned, to make it easier for the FE to save this in message data
// @Tags        DMs
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       message body entity.ImageitemPhoto true "Raw Images"
// @Success     200
// @Router      /v1/imageraw [post]
func CreatePublicImageItem(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var imageaddr entity.ImageitemPhoto
	err_ := json.Unmarshal(requestBody, &imageaddr)
	if err_ != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	// Step 2: Define the parameters for the session you want to create.
	key := "DO00CLQBPDAEHFUTYMGR"        // Access key pair. You can create access key pairs using the control panel or API.
	secret := os.Getenv("SPACES_SECRET") // Secret access key defined through an environment variable.

	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(key, secret, ""), // Specifies your credentials.
		Endpoint:         aws.String("https://sgp1.digitaloceanspaces.com"), // Find your endpoint in the control panel, under Settings. Prepend "https://".
		S3ForcePathStyle: aws.Bool(false),                                   // // Configures to use subdomain/virtual calling format. Depending on your version, alternatively use o.UsePathStyle = false
		Region:           aws.String("us-east-1"),                           // Must be "us-east-1" when creating new Spaces. Otherwise, use the region in your endpoint, such as "nyc3".
	}

	// Step 3: The new session validates your request and directs it to your Space's specified endpoint using the AWS SDK.
	newSession, errs := session.NewSession(s3Config)
	if errs != nil {
		fmt.Println(errs.Error())
	}
	s3Client := s3.New(newSession)

	//public image we want stored in raw so it can be rendered properly by any client
	// Extract the base64 data without the header
	dataParts := strings.Split(imageaddr.Base64data, ";base64,")
	if len(dataParts) != 2 {
		fmt.Println("Invalid base64 data format")
		return
	}

	// Decode base64 data
	decoded, errr := base64.StdEncoding.DecodeString(dataParts[1])
	if errr != nil {
		fmt.Println("Error decoding base64:", errr)
		return
	}

	// Step 4: Define the parameters of the object you want to upload.
	object := s3.PutObjectInput{
		Bucket: aws.String("walletchat-pfp-storage"), // The path to the directory you want to upload the object to, starting with your Space name.
		Key:    aws.String(imageaddr.Imageid),        // Object key, referenced whenever you want to access this file later.
		Body:   bytes.NewReader(decoded),             // The object's contents.
		ACL:    aws.String("public"),                 // Defines Access-control List (ACL) permissions, such as private or public.
		Metadata: map[string]*string{ // Required. Defines metadata tags.
			"x-amz-meta-my-key": aws.String("your-value"),
		},
	}

	// Step 5: Run the PutObject function with your parameters, catching for errors.
	_, err := s3Client.PutObject(&object)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusConflict)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func SaveFileToSpaces(fileData []byte, fileName string) (string, error) {
	key := "DO00CLQBPDAEHFUTYMGR"        // Access key pair.
	secret := os.Getenv("SPACES_SECRET") // Secret access key defined through an environment variable.

	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(key, secret, ""),
		Endpoint:         aws.String("https://sgp1.digitaloceanspaces.com"),
		S3ForcePathStyle: aws.Bool(false),
		Region:           aws.String("us-east-1"), // Must be "us-east-1" when creating new Spaces.
	}

	// Create a new session
	newSession, err := session.NewSession(s3Config)
	if err != nil {
		return "", fmt.Errorf("failed to create session: %v", err)
	}
	s3Client := s3.New(newSession)

	// Define the parameters of the object you want to upload.
	object := s3.PutObjectInput{
		Bucket: aws.String("walletchat-pfp-storage"), // Change to your bucket name
		Key:    aws.String(fileName),                 // Object key
		Body:   bytes.NewReader(fileData),            // The object's contents
		ACL:    aws.String("public-read"),            // Set to public-read for public access
		Metadata: map[string]*string{ // Optional metadata
			"x-amz-meta-my-key": aws.String("your-value"),
		},
	}

	// Run the PutObject function with your parameters
	_, err = s3Client.PutObject(&object)
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %v", err)
	}

	// Construct the public URL
	publicURL := fmt.Sprintf("https://%s.%s/%s", "walletchat-pfp-storage", "sgp1.digitaloceanspaces.com", fileName)
	return publicURL, nil
}

// GetRawImageItem godoc
// @Summary     Store Image in Bucket Storage
// @Description private image storage for photo uploads in DMs
// @Description Images can only be fetched by address to (recipient), or address from (sender)
// @Tags        DMs
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       imageid path    string true "unique image id to fetch"
// @Success     200     {array} entity.Imageitem
// @Router      /v1/imageraw/{imageid} [get]
func GetRawImageItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	imageid := vars["imageid"]

	//check addresses out of it to ensure user can get this file
	Authuser := auth.GetUserFromReqContext(r)
	//fmt.Println("auth check: ", imageid, Authuser.Address)
	if strings.Contains(strings.ToLower(imageid), Authuser.Address) {
		// Step 2: Define the parameters for the session you want to create.
		key := "DO00CLQBPDAEHFUTYMGR"        // Access key pair. You can create access key pairs using the control panel or API.
		secret := os.Getenv("SPACES_SECRET") // Secret access key defined through an environment variable.

		s3Config := &aws.Config{
			Credentials:      credentials.NewStaticCredentials(key, secret, ""), // Specifies your credentials.
			Endpoint:         aws.String("https://sgp1.digitaloceanspaces.com"), // Find your endpoint in the control panel, under Settings. Prepend "https://".
			S3ForcePathStyle: aws.Bool(false),                                   // // Configures to use subdomain/virtual calling format. Depending on your version, alternatively use o.UsePathStyle = false
			Region:           aws.String("us-east-1"),                           // Must be "us-east-1" when creating new Spaces. Otherwise, use the region in your endpoint, such as "nyc3".
		}

		// Step 3: The new session validates your request and directs it to your Space's specified endpoint using the AWS SDK.
		newSession, errs := session.NewSession(s3Config)
		if errs != nil {
			fmt.Println(errs.Error())
		}
		s3Client := s3.New(newSession)

		// Step 4: Define the parameters of the object you want to upload.
		object := s3.GetObjectInput{
			Bucket: aws.String("walletchat-pfp-storage"), // The path to the directory you want to upload the object to, starting with your Space name.
			Key:    aws.String(imageid),                  // Object key, referenced whenever you want to access this file later.
		}

		// Step 5: Run the PutObject function with your parameters, catching for errors.
		imageObject, err := s3Client.GetObject(&object)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		// Read the contents of the object's body
		content, err := ioutil.ReadAll(imageObject.Body)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(string(content))
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}

// UpdateImageItem godoc
// @Summary     Store Image in DB for later user (update existing photo)
// @Description Currently used for the WC HQ Logo, stores the base64 raw data of the profile image for a community
// @Tags        Common
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       message body    entity.Imageitem true "Profile Thumbnail Pic"
// @Success     200     {array} entity.Bookmarkitem
// @Router      /v1/image [put]
func UpdateImageItem(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var imageaddr entity.Imageitem
	json.Unmarshal(requestBody, &imageaddr)

	var returnval bool
	Authuser := auth.GetUserFromReqContext(r)
	if strings.EqualFold(Authuser.Address, imageaddr.Addr) {
		var result = database.Connector.Model(&entity.Imageitem{}).
			Where("addr = ?", imageaddr.Addr).
			Update("base64data", imageaddr.Base64data)

		if result.RowsAffected > 0 {
			returnval = true
		} else {
			database.Connector.Create(&imageaddr)
		}
	} else {
		w.WriteHeader(http.StatusForbidden)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(returnval)
}

// GetImageItem godoc
// @Summary     Get Thumbnail Image Data
// @Description Retreive image data for use with user/community/nft group dislayed icon
// @Tags        Common
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       name path    string true "Wallet/NFT Address Mapped to User/Community"
// @Success     200  {array} entity.Imageitem
// @Router      /v1/image/{addr} [get]
func GetImageItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	addr := vars["addr"]

	var imgaddr []entity.Imageitem

	database.Connector.Where("addr = ?", addr).Find(&imgaddr)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(imgaddr)
}

func TrackEventGA4(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	event := vars["event"]
	email := vars["email"]
	addr := vars["addr"]

	fmt.Println("TrackEventGA4: ", event, addr, email)

	wc_analytics.SendCustomEventWithEmail(addr, event, email)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusCreated)
	//json.NewEncoder(w).Encode(imgaddr)
}

// CreateAddrNameItem godoc
// @Summary     give a common name to a user address, or NFT collection
// @Description Give a common name (Kevin.eth, BillyTheKid, etc) to an Address (*** Will return 403 if address doesn't own ENS name! ***)
// @Description Accepts ADMIN_API_KEY for integrated sign-in and creating new WC users (see top of page for examples)
// @Tags        Common
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       message body      entity.Addrnamesignupitem true "Address and Name to map together"
// @Success     200     {integer} int
// @Router      /v1/name [post]
func CreateAddrNameItem(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var addrnameSignup entity.Addrnamesignupitem
	err := json.Unmarshal(requestBody, &addrnameSignup)
	if err != nil {
		fmt.Println("unmarshal error: ", err)
	}
	var isAdmin bool = false

	var addrname entity.Addrnameitem
	addrname.Address = addrnameSignup.Address
	addrname.Name = addrnameSignup.Name
	apiKey := r.Header.Get("Authorization")
	if len(apiKey) > 0 {
		const prefix = "Bearer "
		if len(apiKey) < len(prefix) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		apiKey = apiKey[len(prefix):]
		if strings.Contains(os.Getenv("ADMIN_API_KEY_LIST"), apiKey) {
			isAdmin = true
			fmt.Println("Found API key in Authorization header")
			wc_analytics.SendCustomEvent(apiKey[:16], "ADMIN_UPDATE_NAME")
		}
	}

	//ensure if user is trying to use .eth that they own it
	if strings.HasSuffix(addrname.Name, ".eth") {
		client, err := ethclient.Dial("https://mainnet.infura.io/v3/" + os.Getenv("INFURA_V3"))
		if err != nil {
			fmt.Println(err)
		}

		// Resolve a name to an address.
		address, err := ens.Resolve(client, addrname.Name)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Address of %s is %s\n", addrname.Name, address.Hex())

		if !strings.EqualFold(address.Hex(), addrname.Address) {
			fmt.Printf("Addresses do not match! %s is %s\n", addrname.Address, address.Hex())
			w.WriteHeader(http.StatusForbidden)
			return
		}
	}
	//end ensuring .eth name is owned by sender

	Authuser := auth.GetUserFromReqContext(r)
	if strings.EqualFold(Authuser.Address, addrname.Address) || isAdmin {
		//create or update in one function is easier
		var addrnameDB entity.Addrnameitem
		var dbQuery = database.Connector.Where("address = ?", addrname.Address).Find(&addrnameDB)

		var affectedRows = 0
		if dbQuery.RowsAffected == 0 {
			var result = database.Connector.Create(&addrname)
			affectedRows = int(result.RowsAffected)
			fmt.Printf("creating addr->name item: %s <-> %s\n", addrname.Address, addrname.Name)
			wc_analytics.SendCustomEvent(addrname.Address, "NEW_SIGNUP")

			//create a settings entry as well to save signupsite, could be combined upon redesign
			fmt.Printf("Signup Site: %s \n", addrnameSignup.Signupsite)
			var settings entity.Settings
			var dbResults = database.Connector.Where("walletaddr = ?", addrname.Address).Find(&settings)
			if dbResults.RowsAffected == 0 {
				if isAdmin {
					var uservalid entity.Referraluser
					uservalid.Referralcode = "wc-admin"
					uservalid.Walletaddr = addrname.Address
					database.Connector.Create(&uservalid)

					if addrnameSignup.Email != "" {
						fmt.Println("Update from Admin: ", apiKey[:16])
						settings.Email = addrnameSignup.Email
						settings.Verified = "true"
						settings.Notify24 = "false"
						settings.Notifydm = "true"
					}
				}
				settings.Domain = addrnameSignup.Domain
				settings.Signupsite = addrnameSignup.Signupsite
				settings.Walletaddr = addrname.Address
				database.Connector.Create(&settings)

				//give new users 3 new referral codes
				referrals.CreateReferralCodeInternal(settings.Walletaddr)

				var chat entity.Chatitem
				chat.Timestamp = time.Now().Format("2006-01-02T15:04:05.000Z")
				chat.Timestamp_dtm = time.Now()
				chat.Fromaddr = strings.ToLower(os.Getenv("SUPPORT_WALLET"))
				chat.Toaddr = strings.ToLower(addrname.Address)
				chat.Nftid = "0"
				//auto-send a message to the user to check out the leaderboard
				if strings.Contains(addrnameSignup.Domain, "app.walletchat.fun") {
					chat.Message = "Welcome to WalletChat!  Head over to the leadboard via the trophy icon and grab your referral codes to invite your frens to start earning today! \nMake sure to follow @wallet_chat on Twitter to increase your chances of winning!"
				} else if strings.Contains(addrnameSignup.Domain, "good") {
					chat.Fromaddr = strings.ToLower(os.Getenv("GD_SUPPORT_WALLET"))
					chat.Message = "Welcome to WalletChat! Here you can talk to GoodDollar support directly!"
				} else {
					chat.Message = "Welcome to WalletChat! Feel free to reach out to us with any questions!"
				}
				database.Connector.Create(&chat)
			}
		} else {
			var result = database.Connector.Model(&entity.Addrnameitem{}).
				Where("address = ?", addrname.Address).
				Update("name", addrname.Name)
			affectedRows = int(result.RowsAffected)
			fmt.Printf("updating addr->name item: %s <-> %s\n", addrname.Address, addrname.Name)

			if isAdmin && addrnameSignup.Email != "" {
				fmt.Println("Update from Admin: ", apiKey[:16])
				fmt.Printf("updating addr->email item: %s <-> %s\n", addrname.Address, addrnameSignup.Email)
				database.Connector.Model(&entity.Settings{}).Where("walletaddr = ?", addrname.Address).Update("email", addrnameSignup.Email)
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(affectedRows)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}

// GetAddrNameItem godoc
// @Summary     get the common name which has been mapped to an address
// @Description get the given a common name (Kevin.eth, BillyTheKid, etc) what has already been mapped to an Address
// @Tags        Common
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       address path    string true "Get Name for given address"
// @Success     200     {array} entity.Addrnameitem
// @Router      /v1/name/{addr} [get]
func GetAddrNameItem(w http.ResponseWriter, r *http.Request) {
	//Authuser := auth.GetUserFromReqContext(r)
	//address := Authuser.Address
	vars := mux.Vars(r)
	address := vars["address"]

	var addrname []entity.Addrnameitem

	database.Connector.Where("address = ?", address).Find(&addrname)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(addrname)
}
func OuraCreateAddrNameItem(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var addrNameItem entity.Addrnameitem
	err := json.Unmarshal(requestBody, &addrNameItem)
	if err != nil {
		fmt.Println("unmarshal error: ", err)
	}

	var addrnameDB entity.Addrnameitem
	var dbQuery = database.Connector.Where("address = ?", addrNameItem.Address).Find(&addrnameDB)
	if dbQuery.RowsAffected == 0 {
		database.Connector.Create(&addrNameItem)
	} else {
		database.Connector.Model(&entity.Addrnameitem{}).
			Where("address = ?", addrNameItem.Address).
			Update("name", addrNameItem.Name)

		fmt.Printf("updating addr->name item: %s <-> %s\n", addrNameItem.Address, addrNameItem.Name)
	}

	//this isn't great long-term, but for smaller table it allows new usernames to show up in table instead of waiting for once daily update
	referrals.GetOuraLeaderboardDataCronJob()

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
}

func OuraTestFile(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	// Print the raw JSON or POST body
	fmt.Printf("Received POST body: %s\n", string(requestBody))

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
}

// func UpdateAddrNameItem(w http.ResponseWriter, r *http.Request) {
// 	requestBody, _ := ioutil.ReadAll(r.Body)
// 	var addrname entity.Addrnameitem
// 	json.Unmarshal(requestBody, &addrname)

// 	Authuser := auth.GetUserFromReqContext(r)
// 	if strings.EqualFold(Authuser.Address, addrname.Address) {
// 		//ensure if user is trying to use .eth that they own it
// 		if strings.HasSuffix(addrname.Name, ".eth") {
// 			client, err := ethclient.Dial("https://mainnet.infura.io/v3/" + os.Getenv("INFURA_V3"))
// 			if err != nil {
// 				fmt.Println(err)
// 			}

// 			// Resolve a name to an address.
// 			address, err := ens.Resolve(client, addrname.Name)
// 			if err != nil {
// 				fmt.Println(err)
// 			}
// 			fmt.Printf("Address of %s is %s\n", addrname.Name, address.Hex())

// 			if strings.ToLower(address.Hex()) != strings.ToLower(addrname.Address) {
// 				w.WriteHeader(http.StatusForbidden)
// 				return
// 			}
// 		}
// 		//end ensuring .eth name is owned by sender

// 		var result = database.Connector.Model(&entity.Addrnameitem{}).
// 			Where("address = ?", addrname.Address).
// 			Update("name", addrname.Name)

// 		var returnval bool
// 		if result.RowsAffected > 0 {
// 			returnval = true
// 		}

// 		w.Header().Set("Content-Type", "application/json")
//    w.Header().Set("X-Content-Type-Options", "nosniff")
// 		w.WriteHeader(http.StatusCreated)
// 		json.NewEncoder(w).Encode(returnval)
// 	} else {
// 		w.WriteHeader(http.StatusForbidden)
// 	}
// }

// GetGroupChatItemsByAddr godoc
// @Summary     Get group chat items, given a wallt FROM address and NFT Contract Address
// @Description Get all group chat items for a given wallet (useraddress) for a given NFT Contract Address (TODO: fix up var names)
// @Tags        NFT
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       address     path    string true "NFT Address"
// @Param       useraddress path    string true "FROM: wallet address"
// @Success     200         {array} entity.Groupchatitem
// @Router      /v1/get_groupchatitems/{address}/{useraddress} [get]
func GetGroupChatItemsByAddr(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nftaddr := vars["address"]
	//fromaddr := vars["useraddress"]
	Authuser := auth.GetUserFromReqContext(r)
	fromaddr := Authuser.Address

	var chat []entity.Groupchatitem

	//TODO this will use up API calls fast if we are polling all the time
	//ensure user holds the NFT first
	isHolder := false
	if strings.HasPrefix(nftaddr, "0x") {
		//TODO: we should send in chain along with message
		isHolder = IsOwnerOfNFT(nftaddr, fromaddr, "ethereum")
		if !isHolder {
			isHolder = IsOwnerOfNFT(nftaddr, fromaddr, "polygon")
		}
	} else if !isHolder && (strings.HasSuffix(fromaddr, ".near") || strings.HasSuffix(fromaddr, ".testnet")) ||
		(len(fromaddr) == 64 && !strings.HasPrefix(fromaddr, "0x")) { //NEAR check
		isHolder = IsOwnerOfNFT(nftaddr, fromaddr, "near")
	} else if !isHolder && strings.HasPrefix(fromaddr, "tz") { //Tezos check
		isHolder = IsOwnerOfNFT(nftaddr, fromaddr, "tezos")
	} else if !isHolder && strings.HasPrefix(nftaddr, "poap_") {
		split := strings.Split(nftaddr, "_")
		isHolder = IsOwnerOfPOAP(split[1], fromaddr)
	}

	//if user is not a holder, can't get the messages
	if isHolder {
		var chatReadTime entity.Groupchatreadtime
		var dbQuery = database.Connector.Where("fromaddr = ?", fromaddr).Where("nftaddr = ?", nftaddr).Find(&chatReadTime)

		//fmt.Printf("Group Chat Get By Addr Result: %#v\n", chatReadTime)

		//if no respsonse to this query, its the first time a user is reading the chat history, send it all
		if dbQuery.RowsAffected == 0 {
			//database.Connector.Where("nftaddr = ?", nftaddr).Find(&chat)  //mana requests all data for now

			//add the first read element to the group timestamp table cross reference
			chatReadTime.Fromaddr = fromaddr
			chatReadTime.Nftaddr = nftaddr
			chatReadTime.Readtimestamp_dtm = time.Now()

			database.Connector.Create(&chatReadTime)
		} else {
			//database.Connector.Where("timestamp > ?", chatReadTime.Readtimestamp_dtm).Where("nftaddr = ?", nftaddr).Find(&chat) //mana requests all data for now
			//set timestamp when this was last grabbed
			currtime := time.Now()
			database.Connector.Model(&entity.Groupchatreadtime{}).Where("fromaddr = ?", fromaddr).Where("nftaddr = ?", nftaddr).Update("readtimestamp_dtm", currtime)
		}
		//this line goes away if we selectively load data in the future
		database.Connector.Where("nftaddr = ?", nftaddr).Find(&chat) //manapixels requests all data for now

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		json.NewEncoder(w).Encode(chat)
	} else {
		fmt.Printf("Not holder....: ")
		w.WriteHeader(http.StatusForbidden)
	}
}

// GetGroupChatItemsByAddrLen godoc
// @Summary     Get Unread Groupchat Items (TODO: cleanup naming convention here)
// @Description For group chat unread counts, currently the database stores a timestamp for each time a user enters a group chat.
// @Description We though in the design it would be impractical to keep a read/unread count copy per user per message, but if this
// @Description method doesn't proof to be fine grained enough, we could add a boolean relational table of read messgages per user.
// @Tags        Common
// @Accept      json
// @Produce     plain
// @Param       name path      string true "Common Name Mapped to User/Community"
// @Success     200  {integer} int
// @Router      /v1/get_groupchatitems_unreadcnt/{address}/{useraddress} [get]
func GetGroupChatItemsByAddrLen(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nftaddr := vars["address"]
	//fromaddr := vars["useraddress"]
	Authuser := auth.GetUserFromReqContext(r)
	fromaddr := Authuser.Address

	var chat []entity.Groupchatitem

	var chatReadTime entity.Groupchatreadtime
	var dbQuery = database.Connector.Where("fromaddr = ?", fromaddr).Where("nftaddr = ?", nftaddr).Find(&chatReadTime)

	//fmt.Printf("Group Chat Get By Addr Result: %#v\n", chatReadTime.Readtimestamp_dtm)

	//if no respsonse to this query, its the first time a user is reading the chat history, send it all
	if dbQuery.RowsAffected == 0 {
		database.Connector.Where("nftaddr = ?", nftaddr).Find(&chat)
	} else {
		database.Connector.Where("timestamp_dtm > ?", chatReadTime.Readtimestamp_dtm).Where("nftaddr = ?", nftaddr).Find(&chat)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	json.NewEncoder(w).Encode(len(chat))
}

// CreateAddrNameItem godoc
// @Summary     Update Message Read Status of a given DM chat message
// @Description Currently this only update the message read/unread status.  It could update the entire JSON struct
// @Description upon request, however we only needed this functionality currently and it saved re-encryption of the data.
// @Description TODO: TO/FROM address in the URL is not needed/not used anymore.
// @Tags        DMs
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       message body    entity.Chatitem true "chat item JSON struct to update msg read status"
// @Success     200     {array} entity.Chatitem
// @Router      /v1/update_chatitem/{fromaddr}/{toaddr} [put]
func UpdateChatitemByOwner(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var chat entity.Chatitem

	json.Unmarshal(requestBody, &chat)
	Authuser := auth.GetUserFromReqContext(r)
	if strings.EqualFold(Authuser.Address, chat.Toaddr) {
		//for now only support updating the read status
		//we would need to re-encrypt the data on message update (not hard just need to add it)
		// database.Connector.Model(&entity.Chatitem{}).
		// 	Where("fromaddr = ?", chat.Fromaddr).
		// 	Where("toaddr = ?", chat.Toaddr).
		// 	Where("timestamp = ?", chat.Timestamp).
		// 	Update("message", chat.Message)
		database.Connector.Model(&entity.Chatitem{}).
			Where("fromaddr = ?", chat.Fromaddr).
			Where("toaddr = ?", chat.Toaddr).
			Where("timestamp = ?", chat.Timestamp).
			Update("msgread", chat.Msgread)

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(chat)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}

// DeleteAllChatitemsToAddressByOwner godoc
// @Summary     Delete All Chat Items (DMs) between sender (from JWT) given addresses
// @Description Currently deletes all chat items between these two addresses
// @Tags        Unused/Legacy
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       address   path string true "Delete convo with Wallet Address"
// @Success     204
// @Router      /v1/deleteall_chatitems/{address} [delete]
func DeleteAllChatitemsToAddressByOwner(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	to := vars["address"]

	var chat entity.Chatitem
	Authuser := auth.GetUserFromReqContext(r)
	owner := Authuser.Address

	rowsAff := 0
	dbQuery := database.Connector.Where("toaddr = ?", to).Where("fromaddr = ?", owner).Delete(&chat)
	rowsAff += int(dbQuery.RowsAffected)
	dbQuery = database.Connector.Where("fromaddr = ?", to).Where("toaddr = ?", owner).Delete(&chat)
	rowsAff += int(dbQuery.RowsAffected)

	if rowsAff > 0 {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

// DeleteAllChatitemsToAddressByOwner godoc
// @Summary     Delete Single Chat Item (DM)
// @Description Can only delete messages sent, cannot delete incoming messages
// @Tags        Unused/Legacy
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       id path string true "message ID"
// @Success     204
// @Router      /v1/delete_chatitem/{id} [delete]
func DeleteChatitem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var chat entity.Chatitem
	Authuser := auth.GetUserFromReqContext(r)
	owner := Authuser.Address

	dbQuery := database.Connector.Where("id = ?", id).Where("fromaddr = ?", owner).Delete(&chat)

	if dbQuery.RowsAffected > 0 {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

// BlockUser godoc
// @Summary     Block User From DMing (DM)
// @Description Unblocking takes manual request from end user
// @Tags        Security
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       address path string true "Address to Block"
// @Success     204
// @Router      /v1/block_user/{address} [get]
func BlockUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	addressToBlock := vars["address"]

	var blockedUser entity.Blockeduser
	Authuser := auth.GetUserFromReqContext(r)
	owner := Authuser.Address

	blockedUser.Owneraddress = owner
	blockedUser.Blockedaddress = addressToBlock

	dbQuery := database.Connector.Create(&blockedUser)

	if dbQuery.RowsAffected > 0 {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

// IsModerator godoc
// @Summary     Used to check for Verified Tick
// @Description Check to see if a wallet is a verified moderator for a given domain (company)
// @Tags        Security
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       address path string true "Address to check if is moderator"
// @Param       company path string true "domain/company to check"
// @Success     204
// @Router      /v1/is_moderator/{company}/{address} [get]
func IsModerator(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	company := vars["company"]
	address := vars["address"]

	var moderator entity.Moderator
	dbQuery := database.Connector.Where("address = ?", address).Where("company = ?", company).Find(&moderator)

	if dbQuery.RowsAffected > 0 {
		json.NewEncoder(w).Encode(true)
		//w.WriteHeader(http.StatusOK)
	} else {
		json.NewEncoder(w).Encode(false)
	}
}

// func CreateSettings(w http.ResponseWriter, r *http.Request) {
// 	requestBody, _ := ioutil.ReadAll(r.Body)
// 	var settings entity.Settings
// 	json.Unmarshal(requestBody, &settings)

// 	Authuser := auth.GetUserFromReqContext(r)
// 	if strings.EqualFold(Authuser.Address, settings.Walletaddr) {
// 		database.Connector.Create(&settings)
// 		w.Header().Set("Content-Type", "application/json")
//    w.Header().Set("X-Content-Type-Options", "nosniff")
// 		w.WriteHeader(http.StatusCreated)
// 		json.NewEncoder(w).Encode(settings)
// 	} else {
// 		w.WriteHeader(http.StatusForbidden)
// 	}
// }

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

// UpdateSettings godoc
// @Summary     Settings hold a user address and the email address for notifications if they opt-in
// @Description Update settings, email address, daily notifications and per DM notifications
// @Tags        Common
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       message body    entity.Settings true "update struct"
// @Success     200     {array} entity.Settings
// @Router      /v1/update_settings [POST]
func UpdateSettings(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var settingsRX entity.Settings
	json.Unmarshal(requestBody, &settingsRX)

	//strip HTTPS prefix and trailing /
	if strings.HasPrefix(settingsRX.Signupsite, "https://") {
		split := strings.Split(settingsRX.Signupsite, "https://")
		settingsRX.Signupsite = split[1]
		settingsRX.Signupsite = strings.TrimSuffix(settingsRX.Signupsite, "/")
	}
	fmt.Println("RX Settings", settingsRX)
	addr := strings.ToLower(settingsRX.Walletaddr)
	Authuser := auth.GetUserFromReqContext(r)
	if strings.EqualFold(Authuser.Address, addr) {
		var settings entity.Settings
		var dbResults = database.Connector.Where("walletaddr = ?", addr).Find(&settings)

		if dbResults.RowsAffected == 0 {
			dbResults = database.Connector.Create(&settingsRX)
			log.Println("Created New Settings")

			//create Telegram Link/Login Code
			if settingsRX.Telegramhandle != "" {
				var telegramVerificationCode = randSeq(20)
				dbResults = database.Connector.Model(&entity.Settings{}).Where("walletaddr = ?", addr).Update("telegramcode", telegramVerificationCode)
				if dbResults.RowsAffected == 0 {
					log.Println("Did not update verification code item for: ", addr)
				}
			}

			//send verification email
			if strings.Contains(settingsRX.Email, "@") {
				var toAddrname entity.Addrnameitem
				dbResults = database.Connector.Where("address = ?", settingsRX.Walletaddr).Find(&toAddrname)
				if dbResults.RowsAffected == 0 {
					log.Println("Did not find addrname item for: ", settingsRX.Walletaddr)
				}
				var verificationCode = randSeq(10)
				dbResults = database.Connector.Model(&entity.Settings{}).Where("walletaddr = ?", addr).Update("verified", verificationCode)
				if dbResults.RowsAffected == 0 {
					log.Println("Did not update verification code item for: ", addr)
				}
				from := mail.NewEmail("WalletChat Notifications", "contact@walletchat.fun")
				subject := "Please Verify Email for " + settingsRX.Signupsite
				to := mail.NewEmail(toAddrname.Name, settingsRX.Email)
				if settingsRX.Signupsite == "" {
					settingsRX.Signupsite = settingsRX.Domain //from the main webapp, domain and signup site is the same
				}
				plainTextContent := "Please verify your email entered at " + settingsRX.Signupsite + " by clicking here: " + settingsRX.Domain + "/verify-email?email=" + settings.Email + "&code=" + verificationCode
				htmlContent := email.NotificationEmailVerify(toAddrname.Address, toAddrname.Name, "Email Verification", settingsRX.Email, verificationCode, settingsRX.Signupsite, settingsRX.Domain)
				message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
				client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
				response, err := client.Send(message)
				if err != nil {
					log.Println(err)
				} else {
					_ = response
				}
			}
			wc_analytics.SendCustomEvent(settingsRX.Walletaddr, "UPDATE_SETTINGS")
		} else {
			if settingsRX.Telegramhandle != "" {
				//create Telegram Link/Login Code
				var telegramVerificationCode = randSeq(20)
				dbResults = database.Connector.Model(&entity.Settings{}).Where("walletaddr = ?", addr).Update("telegramcode", telegramVerificationCode)
				if dbResults.RowsAffected == 0 {
					fmt.Println("Did not update verification code item for: ", addr)
				}

				log.Println("Updating Telegram Handle ", settingsRX.Telegramhandle)

				//technically we don't need the handle since the chatId is really what is used, but this helps with login flow and double checks
				dbResults = database.Connector.Model(&entity.Settings{}).Where("walletaddr = ?", addr).Update("telegramhandle", settingsRX.Telegramhandle)
			}
			if settingsRX.Email != "" {
				fmt.Println("Updating Email", settingsRX.Email)
				dbResults = database.Connector.Model(&entity.Settings{}).Where("walletaddr = ?", addr).Update("email", settingsRX.Email)
				//send verification email
				if strings.Contains(settingsRX.Email, "@") {
					var toAddrname entity.Addrnameitem
					database.Connector.Where("address = ?", settingsRX.Walletaddr).Find(&toAddrname)

					var verificationCode = randSeq(10)
					database.Connector.Model(&entity.Settings{}).Where("walletaddr = ?", addr).Update("verified", verificationCode)

					from := mail.NewEmail("WalletChat Notifications", "contact@walletchat.fun")
					if settingsRX.Signupsite == "" {
						settingsRX.Signupsite = settingsRX.Domain //from the main webapp, domain and signup site is the same
					}
					if settingsRX.Signupsite != "" {
						settings.Signupsite = settingsRX.Signupsite //use the received one over past saved signup site.
					}
					subject := "Please Verify Email for " + settings.Signupsite
					to := mail.NewEmail(toAddrname.Name, settingsRX.Email)
					plainTextContent := "Please verify your email entered at " + settings.Signupsite + " by clicking here: " + settings.Domain + "/verify-email?email=" + settings.Email + "&code=" + verificationCode
					htmlContent := email.NotificationEmailVerify(toAddrname.Address, toAddrname.Name, "Email Verification", settingsRX.Email, verificationCode, settingsRX.Signupsite, settingsRX.Domain)
					message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
					client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
					response, err := client.Send(message)
					if err != nil {
						log.Println(err)
					} else {
						_ = response
					}
				}
			}
			if settingsRX.Verified != "" {
				log.Println("Updating Verifed Email Status", settingsRX.Verified)
				dbResults = database.Connector.Model(&entity.Settings{}).Where("walletaddr = ?", addr).Update("verified", settingsRX.Verified)
			}
			if settingsRX.Notifydm != "" {
				log.Println("Updating Daily Notifications", settingsRX.Notifydm)
				dbResults = database.Connector.Model(&entity.Settings{}).Where("walletaddr = ?", addr).Update("notifydm", settingsRX.Notifydm)
			}
			if settingsRX.Notify24 != "" {
				log.Println("Updating Daily Notifications", settingsRX.Notify24)
				dbResults = database.Connector.Model(&entity.Settings{}).Where("walletaddr = ?", addr).Update("notify24", settingsRX.Notify24)
			}
			if settingsRX.Signupsite != "" {
				log.Println("Updating Signup Site", settingsRX.Signupsite)
				//strip HTTPS prefix and trailing /
				if strings.HasPrefix(settingsRX.Signupsite, "https://") {
					split := strings.Split(settingsRX.Signupsite, "https://")
					settingsRX.Signupsite = split[1]
					settingsRX.Signupsite = strings.TrimSuffix(settingsRX.Signupsite, "/")
				}
				dbResults = database.Connector.Model(&entity.Settings{}).Where("walletaddr = ?", addr).Update("signupsite", settingsRX.Signupsite)
			}
			if settingsRX.Domain != "" {
				log.Println("Updating Domain", settingsRX.Domain)
				dbResults = database.Connector.Model(&entity.Settings{}).Where("walletaddr = ?", addr).Update("domain", settingsRX.Domain)
			}
			if settingsRX.Twitteruser != "" {
				//we want to ensure consistency with the @ before the username.
				//we store without the leading @ because twitter doesn't store it this way either.
				//We can verify on client side but should double check here
				settingsRX.Twitteruser = strings.TrimPrefix(settingsRX.Twitteruser, "@")

				log.Println("Updating Twitter User", settingsRX.Twitteruser)
				database.Connector.Model(&entity.Settings{}).Where("walletaddr = ?", addr).Update("twitteruser", settingsRX.Twitteruser)
				//when a user updates the username - must re-verify
				dbResults = database.Connector.Model(&entity.Settings{}).Where("walletaddr = ?", addr).Update("twitterverified", "false")
			}
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		json.NewEncoder(w).Encode(dbResults.RowsAffected)
	} else {
		fmt.Println("UpdateSettings - JWT Address: ", Authuser.Address)
		fmt.Println("UpdateSettings - POST Address: ", addr)
		w.WriteHeader(http.StatusForbidden)
	}
}

// UpdateSettings godoc
// @Summary     Link a user can click in email to verify email address, will have unique code
// @Description Users will get an email when signing-up to verify email, to ensure we do not send spam
// @Tags        Common
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       message body    entity.Settings true "update struct"
// @Success     200     {array} entity.Settings
// @Router      /v1/verify_email/{email}/{verification_code} [GET]
func VerifyEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]
	code := vars["code"]
	log.Println("Verify Email: ", email)
	var settingsRX []entity.Settings
	var dbResults = database.Connector.Where("email = ?", email).Find(&settingsRX)
	if dbResults.RowsAffected == 0 {
		w.WriteHeader(http.StatusForbidden)
	} else {
		dbResults.RowsAffected = 0
		for i := 0; i < len(settingsRX); i++ {
			if settingsRX[i].Verified == code {
				log.Println("Updating Verifed Email Status", settingsRX[i].Verified)
				dbResults = database.Connector.Model(&entity.Settings{}).Where("walletaddr = ?", settingsRX[i].Walletaddr).Update("verified", "true")
				break
			}
		}
		if dbResults.RowsAffected == 0 {
			w.WriteHeader(http.StatusForbidden)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("X-Content-Type-Options", "nosniff")
			json.NewEncoder(w).Encode(dbResults.RowsAffected)
		}
	}
}

// DeleteSettings godoc
// @Summary     Delete Settings Info
// @Description TODO: not yet used
// @Tags        Unused/Legacy
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       address path string true "Wallet Address"
// @Success     204
// @Router      /v1/delete_settings/{address} [delete]
func DeleteSettings(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//key := vars["address"]
	Authuser := auth.GetUserFromReqContext(r)
	key := Authuser.Address

	var settings entity.Settings

	database.Connector.Where("walletaddr = ?", key).Delete(&settings)
	w.WriteHeader(http.StatusNoContent)
}

// GetSettings godoc
// @Summary     Get Settings Info
// @Description TODO: not yet used
// @Tags        Unused/Legacy
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       address path    string true "Wallet Address"
// @Success     200     {array} entity.Settings
// @Router      /v1/get_settings/{address} [get]
func GetSettings(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// key := vars["address"]
	Authuser := auth.GetUserFromReqContext(r)
	key := Authuser.Address

	var settings []entity.Settings
	database.Connector.Where("walletaddr = ?", key).Find(&settings)

	//if there is a verification code, make sure to clear it out
	//or this would be a vulnerability that people could verify other email addresses
	if len(settings) > 0 && len(settings[0].Verified) > 9 {
		settings[0].Verified = "false"
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	json.NewEncoder(w).Encode(settings)
}

func Erc20TokenOverlap(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tokenAddress := vars["contract_address"]

	url := "https://api.dune.com/api/v1/query/3615247/execute"

	// Create the POST request payload
	payload := map[string]interface{}{
		"query_parameters": map[string]string{
			"token_address": tokenAddress,
		},
		"performance": "medium",
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "Error preparing request", http.StatusInternalServerError)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		http.Error(w, "Error creating request", http.StatusInternalServerError)
		return
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-Dune-API-Key", os.Getenv("DUNE_API_KEY"))
	req.Header.Add("Content-Type", "application/json")

	// Send the POST request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "Error sending request", http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	// Read and parse the response
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		http.Error(w, "Error reading response", http.StatusInternalServerError)
		return
	}

	var response struct {
		ExecutionID string `json:"execution_id"`
		State       string `json:"state"`
	}
	if err := json.Unmarshal(body, &response); err != nil {
		http.Error(w, "Error parsing response", http.StatusInternalServerError)
		return
	}

	// Check the status of the execution periodically
	for {
		statusURL := "https://api.dune.com/api/v1/execution/" + response.ExecutionID + "/results"
		statusReq, err := http.NewRequest("GET", statusURL, nil)
		if err != nil {
			http.Error(w, "Error creating status request", http.StatusInternalServerError)
			return
		}
		statusReq.Header.Add("X-Dune-API-Key", os.Getenv("DUNE_API_KEY"))

		statusRes, err := http.DefaultClient.Do(statusReq)
		if err != nil {
			http.Error(w, "Error sending status request", http.StatusInternalServerError)
			return
		}

		statusBody, err := ioutil.ReadAll(statusRes.Body)
		statusRes.Body.Close()
		if err != nil {
			http.Error(w, "Error reading status response", http.StatusInternalServerError)
			return
		}

		var statusResponse struct {
			IsExecutionFinished bool `json:"is_execution_finished"`
		}
		json.Unmarshal(statusBody, &statusResponse)

		fmt.Println("requestID current status: ", response.ExecutionID, statusResponse)

		if statusResponse.IsExecutionFinished {
			// Once finished, respond with the final results
			json.NewEncoder(w).Encode(string(statusBody))
			break
		}
		time.Sleep(30 * time.Second)
	}
}

func SolTokenOverlap(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tokenAddress := vars["contract_address"]

	url := "https://api.dune.com/api/v1/query/3623869/execute"

	// Create the POST request payload
	payload := map[string]interface{}{
		"query_parameters": map[string]string{
			"token_address": tokenAddress,
		},
		"performance": "medium",
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "Error preparing request", http.StatusInternalServerError)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		http.Error(w, "Error creating request", http.StatusInternalServerError)
		return
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-Dune-API-Key", os.Getenv("DUNE_API_KEY"))
	req.Header.Add("Content-Type", "application/json")

	// Send the POST request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "Error sending request", http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	// Read and parse the response
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		http.Error(w, "Error reading response", http.StatusInternalServerError)
		return
	}

	var response struct {
		ExecutionID string `json:"execution_id"`
		State       string `json:"state"`
	}
	if err := json.Unmarshal(body, &response); err != nil {
		http.Error(w, "Error parsing response", http.StatusInternalServerError)
		return
	}

	// Check the status of the execution periodically
	for {
		statusURL := "https://api.dune.com/api/v1/execution/" + response.ExecutionID + "/results"
		statusReq, err := http.NewRequest("GET", statusURL, nil)
		if err != nil {
			http.Error(w, "Error creating status request", http.StatusInternalServerError)
			return
		}
		statusReq.Header.Add("X-Dune-API-Key", os.Getenv("DUNE_API_KEY"))

		statusRes, err := http.DefaultClient.Do(statusReq)
		if err != nil {
			http.Error(w, "Error sending status request", http.StatusInternalServerError)
			return
		}

		statusBody, err := ioutil.ReadAll(statusRes.Body)
		statusRes.Body.Close()
		if err != nil {
			http.Error(w, "Error reading status response", http.StatusInternalServerError)
			return
		}

		var statusResponse struct {
			IsExecutionFinished bool `json:"is_execution_finished"`
		}
		json.Unmarshal(statusBody, &statusResponse)

		fmt.Println("requestID current status: ", response.ExecutionID, statusResponse)

		if statusResponse.IsExecutionFinished {
			// Once finished, respond with the final results
			json.NewEncoder(w).Encode(string(statusBody))
			break
		}
		time.Sleep(30 * time.Second)
	}
}

// GetSettings godoc
// @Summary     Generic Resolve Name Service
// @Description Resolve .ETH, .BNB, .ARB names
// @Tags
// @Accept   json
// @Produce     json
// @Security    BearerAuth
// @Param    address path    string true "ENS/BNB/ARB/TEZ/NEAR/BTC Name"
// @Success  200     {array} entity.Settings
// @Router   /v1/resolve_name/{name} [get]
func ResolveName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nameToResolve := vars["name"]
	// Authuser := auth.GetUserFromReqContext(r)
	// key := Authuser.Address
	fmt.Println("ENS: ", nameToResolve)

	if strings.HasSuffix(nameToResolve, ".eth") {
		url := "https://deep-index.moralis.io/api/v2/resolve/ens/" + nameToResolve

		req, _ := http.NewRequest("GET", url, nil)

		req.Header.Add("Accept", "application/json")
		req.Header.Add("X-API-Key", os.Getenv("MORALIS_NFT_API_KEY"))

		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		var parsed any
		json.Unmarshal(body, &parsed)

		json.NewEncoder(w).Encode(parsed)
	} else if strings.HasSuffix(nameToResolve, ".btc") {
		url := "https://stacks-node-api.mainnet.stacks.co/v1/names/" + nameToResolve

		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Add("Accept", "application/json")
		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		var parsed any
		json.Unmarshal(body, &parsed)

		json.NewEncoder(w).Encode(parsed)
	} else if strings.HasSuffix(nameToResolve, ".bnb") || strings.HasSuffix(nameToResolve, ".arb") {
		//https://docs.space.id/developer-guide/web3-name-sdk/sid-api
		chain := "bnb"
		if strings.HasSuffix(nameToResolve, ".arb") {
			chain = "arb1"
		}
		url := "https://api.prd.space.id/v1/getAddress?tld=" + chain + "&domain=" + nameToResolve

		// Create a new request using http
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Add("Content-Type", "application/json")

		// Send req using http Client
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println("Error on response.\n[ERROR] -", err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error while reading the response bytes:", err)
		}

		var parsed any
		json.Unmarshal(body, &parsed)

		json.NewEncoder(w).Encode(parsed)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
}

// CreateComments godoc
// @Summary     Comments are used within an NFT community chat
// @Description Comments are meant to be public, someday having an up/downvote method for auto-moderation
// @Tags        NFT
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       message body    entity.Comments true "create struct"
// @Success     200     {array} entity.Comments
// @Router      /v1/create_comments [post]
func CreateComments(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var comment entity.Comments
	json.Unmarshal(requestBody, &comment)

	Authuser := auth.GetUserFromReqContext(r)
	if strings.EqualFold(Authuser.Address, comment.Fromaddr) {
		database.Connector.Create(&comment)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(comment)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}

func DebugPrint(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)

	// Define a variable of type json.RawMessage
	var raw json.RawMessage

	// Unmarshal the JSON string into the json.RawMessage variable
	if err := json.Unmarshal([]byte(requestBody), &raw); err != nil {
		//[object Object]
		objectObject := []byte{91, 111, 98, 106, 101, 99, 116, 32, 79, 98, 106, 101, 99, 116, 93}
		// Compare the two byte arrays
		if bytes.Equal(objectObject, requestBody) {
			//fmt.Println("Couldn't unmarshal :", requestBody, err)
			return
		} else {
			fmt.Println("Couldn't unmarshal :", requestBody, err)
			return
		}
	}

	// Print the raw JSON as a string
	fmt.Println("DEBUG Raw:", string(raw))

	w.WriteHeader(http.StatusOK)
}

// func UpdateComments(w http.ResponseWriter, r *http.Request) {
// 	requestBody, _ := ioutil.ReadAll(r.Body)
// 	var comment entity.Comment

// 	json.Unmarshal(requestBody, &comment)
// 	database.Connector.Model(&entity.Settings{}).Where("walletaddr = ?", settings.Walletaddr).Update("publickey", settings.Publickey)

// 	w.Header().Set("Content-Type", "application/json")
//    w.Header().Set("X-Content-Type-Options", "nosniff")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(comment)
// }

// DeleteComments godoc
// @Summary     Delete Public Comments for given FROM wallet address, NFT Contract and ID
// @Description NFTs have a public comment section
// @Tags        NFT
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       address path string true "FROM Wallet Address"
// @Param       nftaddr path string true "NFT Contract Address"
// @Param       nftid   path string true "NFT ID"
// @Success     204
// @Router      /v1/delete_comments/{fromaddr}/{nftaddr}/{nftid} [delete]
func DeleteComments(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//fromaddr := vars["address"]
	nftaddr := vars["nftaddr"]
	nftid := vars["nftid"]
	Authuser := auth.GetUserFromReqContext(r)
	fromaddr := Authuser.Address

	var comment entity.Comments

	database.Connector.Where("fromaddr = ?", fromaddr).Where("nftaddr = ?", nftaddr).Where("nftid = ?", nftid).Delete(&comment)
	w.WriteHeader(http.StatusNoContent)
}

// GetComments godoc
// @Summary     Get Public Comments for given NFT Contract and ID
// @Description NFTs have a public comment section
// @Tags        NFT
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       address path    string true "Wallet Address"
// @Success     200     {array} entity.Comments
// @Router      /v1/get_comments/{nftaddr}/{nftid} [get]
func GetComments(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["nftid"]
	addr := vars["nftaddr"]

	var comment []entity.Comments
	database.Connector.Where("nftaddr = ?", addr).Where("nftid = ?", id).Find(&comment)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	json.NewEncoder(w).Encode(comment)
}

// GetCommentsCount godoc
// @Summary     Get Public Comments Count for given NFT Contract and ID
// @Description NFTs have a public comment section
// @Tags        NFT
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       nftaddr path      string true "NFT Contract Address"
// @Param       nftid   path      string true "NFT ID"
// @Success     200     {integer} int
// @Router      /v1/get_comments_cnt/{nftaddr}/{nftid} [get]
func GetCommentsCount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["nftid"]
	addr := vars["nftaddr"]

	var comment []entity.Comments
	database.Connector.Where("nftaddr = ?", addr).Where("nftid = ?", id).Find(&comment)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	json.NewEncoder(w).Encode(len(comment))
}

// func GetAllComments(w http.ResponseWriter, r *http.Request) {
// 	var comment []entity.Comments
// 	database.Connector.Find(&comment)

// 	//make sure to get the name if it wasn't there (not there by default now)
// 	var addrname entity.Addrnameitem
// 	for i := 0; i < len(comment); i++ {
// 		var result = database.Connector.Where("address = ?", comment[i].Fromaddr).Find(&addrname)
// 		if result.RowsAffected > 0 {
// 			comment[i].Name = addrname.Name
// 		}
// 	}
// 	//end of adding names for fromaddr

// 	w.Header().Set("Content-Type", "application/json")
//    w.Header().Set("X-Content-Type-Options", "nosniff")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(comment)
// }

func GetOpenseaCollectionSlug(contractAddr string) string {
	//url := "https://api.opensea.io/api/v1/asset_contract/" + contractAddr
	url := "https://api.opensea.io/api/v2/chain/ethereum/contract/" + contractAddr

	// Create a new request using http
	req, _ := http.NewRequest("GET", url, nil)
	osKey := os.Getenv("OPENSEA_API_KEY")
	req.Header.Add("X-API-KEY", osKey)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("OSea API CALL - Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	// Create a variable of type map[string]interface{} to hold the JSON data
	var data map[string]interface{}
	body, _ := ioutil.ReadAll(resp.Body)
	// Unmarshal the JSON response into the data map
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return ""
	}

	// Access the "slug" field using the map
	if collection, ok := data["collection"].(string); ok {
		//fmt.Println("Slug:", collection)
		return collection
	}

	return ""
}

func GetOpenseaCollectionStats(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	contractAddr := vars["contract"]
	slug := GetOpenseaCollectionSlug(contractAddr)
	//fmt.Println("slug: ", slug)
	//https://api.opensea.io/api/v2/collections/{collection_slug}/stats
	url := "https://api.opensea.io/api/v2/collections/" + slug + "/stats"

	// Create a new request using http
	req, _ := http.NewRequest("GET", url, nil)
	osKey := os.Getenv("OPENSEA_API_KEY")
	req.Header.Add("X-API-KEY", osKey)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("OSea API CALL - Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var parsed any
	json.Unmarshal(body, &parsed)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	json.NewEncoder(w).Encode(parsed)
}

// mainly for internal use - API integrations should get own API key
func GetOpenseaAssetContract(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	contractAddr := vars["contract"]
	slug := GetOpenseaCollectionSlug(contractAddr)
	url := "https://api.opensea.io/api/v2/collections/" + slug

	// Create a new request using http
	req, _ := http.NewRequest("GET", url, nil)
	osKey := os.Getenv("OPENSEA_API_KEY")
	req.Header.Add("X-API-KEY", osKey)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("OSea API CALL - Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var parsed any
	json.Unmarshal(body, &parsed)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	json.NewEncoder(w).Encode(parsed)
}

// mainly for internal use - API integrations should get own API key
func GetOpenseaAsset(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nftaddr := vars["nftaddr"]
	nftid := vars["nftid"]
	//address := vars["address"]
	//https://api.opensea.io/api/v1/asset/${msg.nftAddr}/${msg.nftId}?account_address=${account}
	//url := "https://api.opensea.io/api/v1/asset/" + nftaddr + "/" + nftid + "/?account_address=" + address
	//https://api.opensea.io/api/v2/chain/{chain}/contract/{address}/nfts/{identifier}
	url := "https://api.opensea.io/api/v2/chain/ethereum/contract/" + nftaddr + "/nfts/" + nftid

	// Create a new request using http
	req, _ := http.NewRequest("GET", url, nil)
	osKey := os.Getenv("OPENSEA_API_KEY")
	req.Header.Add("X-API-KEY", osKey)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("OSea API CALL asset - Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var parsed any
	json.Unmarshal(body, &parsed)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	json.NewEncoder(w).Encode(parsed)
}

// mainly for internal use - API integrations should get own API key
func GetOpenseaAssetOwner(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	owner := vars["address"]
	//https://api.opensea.io/api/v1/assets?owner=${account}
	//https://api.opensea.io/api/v2/chain/{chain}/account/{address}/nfts
	url := "https://api.opensea.io/api/v2/chain/ethereum/account/" + owner + "/nfts"

	// Create a new request using http
	req, _ := http.NewRequest("GET", url, nil)
	osKey := os.Getenv("OPENSEA_API_KEY")
	req.Header.Add("X-API-KEY", osKey)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("OSea API CALL asset - Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var parsed any
	json.Unmarshal(body, &parsed)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	json.NewEncoder(w).Encode(parsed)
}

// mainly for internal use - API integrations should get own API key
func GetOpenseaAssetOwnerENS(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	owner := vars["address"]
	//https://api.opensea.io/api/v1/assets?owner=${account}&collection=ens
	//url := "https://api.opensea.io/api/v1/assets?owner=" + owner + "&collection=ens"
	//https://api.opensea.io/api/v2/chain/ethereum/account//nfts?collection=ens
	url := "https://api.opensea.io/api/v2/chain/ethereum/account/" + owner + "/nfts?collection=ens"

	// Create a new request using http
	req, _ := http.NewRequest("GET", url, nil)
	osKey := os.Getenv("OPENSEA_API_KEY")
	req.Header.Add("X-API-KEY", osKey)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("OSea API CALL asset - Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var parsed any
	json.Unmarshal(body, &parsed)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	json.NewEncoder(w).Encode(parsed)
}

func GetTwitter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	contract := vars["contract"]

	slug := GetOpenseaCollectionSlug(contract)
	handle := GetTwitterHandle(slug)
	twitterID := GetTwitterID(handle)
	tweets := GetTweetsFromAPI(twitterID)
	formatted := FormatTwitterData(tweets)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	json.NewEncoder(w).Encode(formatted)
}

func GetTwitterCount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	contract := vars["contract"]

	slug := GetOpenseaCollectionSlug(contract)
	handle := GetTwitterHandle(slug)
	twitterID := GetTwitterID(handle)
	tweets := GetTweetsFromAPI(twitterID)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	json.NewEncoder(w).Encode(len(tweets.Data))
}

func GetTwitterHandle(slug string) string {
	//https://api.opensea.io/api/v2/collections/{collection_slug}
	url := "https://api.opensea.io/api/v2/collections/" + slug

	// Create a new request using http
	req, _ := http.NewRequest("GET", url, nil)
	osKey := os.Getenv("OPENSEA_API_KEY")
	req.Header.Add("X-API-KEY", osKey)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	var result OpenseaData
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON - GetTwitterHandle")
	}

	collection := result.Collection.TwitterUsername

	fmt.Printf("get twitter username: %#v\n", collection)

	return collection
}

func GetTwitterID(twitterHandle string) string {
	url := "https://api.twitter.com/2/users/by/username/" + twitterHandle

	// Create a Bearer string by appending string access token
	bearer := "Bearer " + os.Getenv("TWITTER_BEARER")

	// Create a new request using http
	req, _ := http.NewRequest("GET", url, nil)

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	var result TwitterIdResp
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON - GetTwitterID", body)
	}

	twitterID := result.Data.Id

	//fmt.Printf("get twitter ID: %#v\n", twitterID)

	return twitterID
}

func GetTweetsFromAPI(twitterID string) TwitterTweetsData {
	//url := "https://api.twitter.com/2/users/" + twitterID + "/tweets"
	url := "https://api.twitter.com/2/users/" + twitterID + "/tweets?media.fields=height,width,url,preview_image_url,type&tweet.fields=attachments,created_at&user.fields=profile_image_url,username&expansions=author_id,attachments.media_keys&exclude=retweets"

	// Create a Bearer string by appending string access token
	bearer := "Bearer " + os.Getenv("TWITTER_BEARER")

	// Create a new request using http
	req, _ := http.NewRequest("GET", url, nil)

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error fetching twitter: ", err)
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error fetching twitter bytes: ", err)
		log.Println("Error while reading the response bytes:", err)
	}

	//fmt.Println("body twitter: ", body)

	var result TwitterTweetsData
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		//fmt.Println("Can not unmarshal JSON - GetTweetsFromAPI: ", twitterID)
		//TODO: data struct has likely changed
		err = nil
	}
	//fmt.Println("length twitter: ", len(result.Data))

	return result
}

func FormatTwitterData(data TwitterTweetsData) []TweetType {
	var tweets []TweetType
	if len(data.Data) > 0 {
		var user User
		if len(data.Includes.Users) > 0 {
			user = data.Includes.Users[0]
		}

		//for i, item := range data.data {
		//first copy just data.data stuff
		for i := 0; i < len(data.Data); i++ {
			// Text        string `json:"text"`
			// ID          string `json:"id"`
			// Attachments struct {
			// 	MediaKeys []string `json:"media_keys"`
			// } `json:"attachments"`
			// AuthorID  string    `json:"author_id"`
			// CreatedAt time.Time `json:"created_at"`
			var initData TweetType
			initData.Text = data.Data[i].Text
			initData.ID = data.Data[i].ID
			initData.Attachments = data.Data[i].Attachments
			initData.AuthorID = data.Data[i].AuthorID
			initData.CreatedAt = data.Data[i].CreatedAt
			tweets = append(tweets, initData)
		}

		for i := 0; i < len(data.Data); i++ {
			tweets[i].User = user

			if len(data.Data[i].Attachments.MediaKeys) > 0 {
				var localAttachment Attachments
				for j := 0; j < len(tweets[i].Attachments.MediaKeys); j++ {
					var mediaKey = tweets[i].Attachments.MediaKeys[j]
					if len(data.Includes.Media) > 0 {
						//var matched = data.includes.media.find((item => item.media_key === mediaKey))
						for _, v := range data.Includes.Media {
							if v.MediaKey == mediaKey {
								if v.URL != "" {
									localAttachment.MediaKeys = append(localAttachment.MediaKeys, v.URL)
								}
							}
						}
					}
				}
				if len(localAttachment.MediaKeys) > 0 {
					tweets[i].Media = localAttachment
				}
			}
		}
	}
	return tweets
}

// GetCommunityChat godoc
// @Summary     Get Community Chat Landing Page Info
// @Description TODO: need a creation API for communities, which includes specificied welcome message text, Twitter handle, page title
// @Tags        GroupChat
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       address path    string true "Wallet Address"
// @Param       address path    string true "Wallet Address"
// @Success     200       {array} LandingPageItems
// @Router      /v1/community/{community}/{address} [get]
func GetCommunityChat(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	community := vars["community"]
	key := vars["address"]
	var landingData LandingPageItems

	//for now, the walletchat living room is all users by default
	var memberCount []entity.Bookmarkitem
	database.Connector.Where("nftaddr = ?", community).Find(&memberCount)
	landingData.MemberCount = len(memberCount)

	//need to re-architect this - will be very slow
	// for i := 0; i < landingData.MemberCount; i++ {
	// 	var localMember CommunityMember

	// 	localMember.Address = memberCount[i].Walletaddr
	// 	var localAddrName entity.Addrnameitem
	// 	database.Connector.Where("address = ?", localMember.Address).Find(&localAddrName)
	// 	localMember.Name = localAddrName.Name

	// 	var localImage entity.Imageitem
	// 	database.Connector.Where("addr = ?", localMember.Address).Find(&localImage)
	// 	localMember.Image = localImage.Base64data

	// 	localMember.Admin = false
	// 	var isAdmin entity.Communityadmin
	// 	database.Connector.Where("slug = ?", community).Find(&isAdmin)
	// 	if strings.EqualFold(localMember.Address, isAdmin.Adminaddr) {
	// 		localMember.Admin = true
	// 	}

	// 	landingData.Members = append(landingData.Members, localMember)
	// }

	//name (this might be better moved to a different table someday)
	var addrname entity.Addrnameitem
	database.Connector.Where("address = ?", community).Find(&addrname)
	landingData.Name = addrname.Name

	//logo base64 data (url requires other changes)
	var imgaddr entity.Imageitem
	database.Connector.Where("addr = ?", community).Find(&imgaddr)
	landingData.Logo = imgaddr.Base64data

	//WalletChat is verified of course
	landingData.Verified = true

	//auto-join new users to WalletChat community (they can leave later) - might need to break this out
	var bookmarks []entity.Bookmarkitem
	var dbQuery = database.Connector.Where("nftaddr = ?", community).Where("walletaddr = ?", key).Find(&bookmarks)
	if dbQuery.RowsAffected == 0 {
		var bookmark entity.Bookmarkitem
		bookmark.Nftaddr = community
		bookmark.Walletaddr = key
		bookmark.Chain = "none"

		database.Connector.Create(&bookmark)

		//by default everyone is joined to Walletchat
		landingData.Joined = true
		//create the welcome message, save it
		var newgroupchatuser entity.Groupchatitem
		newgroupchatuser.Type = entity.Welcome
		newgroupchatuser.Contexttype = entity.Community
		newgroupchatuser.Fromaddr = key
		newgroupchatuser.Nftaddr = community
		newgroupchatuser.Message = "Welcome " + key + " to " + landingData.Name + "!"
		newgroupchatuser.Timestamp_dtm = time.Now()
		newgroupchatuser.Timestamp = time.Now().Format("2006-01-02T15:04:05.000Z")

		//add it to the database
		database.Connector.Create(&newgroupchatuser)
	} else {
		//We don't have a way for users to get back to WC HQ if they leave (shouldn't need to use above block to re-welcome them)
		landingData.Joined = true
	}

	//check messages read for this user address because this GetCommunityChat is being called
	//separately each time (I thought it would be filled from bookmarks)
	var groupchat []entity.Groupchatitem
	database.Connector.Where("nftaddr = ?", community).Where("fromaddr = ?", key).Find(&groupchat)
	//redoing some things already done in getGroupChatItemsByAddr
	var chatReadTime entity.Groupchatreadtime
	dbQuery = database.Connector.Where("fromaddr = ?", key).Where("nftaddr = ?", community).Find(&chatReadTime)
	if dbQuery.RowsAffected == 0 {
		//add the first read element to the group timestamp table cross reference
		chatReadTime.Fromaddr = key
		chatReadTime.Nftaddr = community
		chatReadTime.Readtimestamp_dtm = time.Now()

		database.Connector.Create(&chatReadTime)
	} else {
		//database.Connector.Where("timestamp > ?", chatReadTime.Readtimestamp_dtm).Where("nftaddr = ?", nftaddr).Find(&chat) //mana requests all data for now
		//set timestamp when this was last grabbed
		currtime := time.Now()
		database.Connector.Model(&entity.Groupchatreadtime{}).Where("fromaddr = ?", key).Where("nftaddr = ?", community).Update("readtimestamp_dtm", currtime)
	}

	var hasMessaged bool
	if len(groupchat) > 0 {
		hasMessaged = true
	} else {
		hasMessaged = false
	}
	landingData.Messaged = hasMessaged

	//grab all the data for walletchat group
	database.Connector.Where("nftaddr = ?", community).Order("id desc").Limit(100).Find(&groupchat)
	landingData.Messages = groupchat

	//get social media info
	var socialMediaMatches []entity.Communitysocial
	database.Connector.Where("community = ?", community).Find(&socialMediaMatches)
	for i := 0; i < len(socialMediaMatches); i++ {
		if socialMediaMatches[i].Type == "twitter" {
			//fmt.Println("adding Twitter social: ", socialmedia.Name)
			//get twitter data
			twitterID := GetTwitterID(socialMediaMatches[i].Name)
			tweets := GetTweetsFromAPI(twitterID)
			formatted := FormatTwitterData(tweets)
			landingData.Tweets = formatted
		}
		//social data
		var insertSocial SocialMsg
		insertSocial.Type = socialMediaMatches[i].Type
		insertSocial.Username = socialMediaMatches[i].Name
		landingData.Social = append(landingData.Social, insertSocial)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	json.NewEncoder(w).Encode(landingData)
}

// GetCommunityChatAfterTime godoc
// @Summary     Get Community Chat Items When Scrolling
// @Description Get Community Chat Items When Scrolling
// @Tags        GroupChat
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       community path    string true "community slug"
// @Param       time      path    string true "Timestamp of last message in current community chat"
// @Param       count     path    string true "Number of Messages To Get (1-1000)"
// @Success     200     {array} LandingPageItems
// @Router      /v1/community/{community}/{time}/{count} [get]
func GetCommunityChatAfterTime(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	community := vars["community"]
	time := vars["time"]

	count, _ := strconv.Atoi(vars["count"])
	if count > 1000 || count < 1 {
		count = 100
	}

	var messages []entity.Groupchatitem
	database.Connector.Where("nftaddr = ?", community).Order("id desc").Where("timestamp_dtm < ?", time).Limit(count).Find(&messages)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	json.NewEncoder(w).Encode(messages)
}

// GetCommunityChatPage godoc
// @Summary     Get Community Chat Items When Scrolling by page
// @Description Get Community Chat Items When Scrolling by page, get latest is page 1
// @Tags        GroupChat
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       community path    string true "community slug"
// @Param       pagenum   path    string true "page number to get (1-N)"
// @Success     200     {array}   LandingPageItems
// @Router      /v1/community_pagenum/{community}/{pagenum} [get]
func GetCommunityChatPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	community := vars["community"]

	pageNumber, _ := strconv.Atoi(vars["pagenum"])
	if pageNumber < 1 {
		pageNumber = 1
	}

	itemsPerPage := 100
	// Calculate the offset based on the page number
	offset := (pageNumber - 1) * itemsPerPage

	var messages []entity.Groupchatitem
	database.Connector.Where("nftaddr = ?", community).Order("id desc").Limit(itemsPerPage).Offset(offset).Find(&messages)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	json.NewEncoder(w).Encode(messages)
}

// IsOwner godoc
// @Summary     Check if given wallet address owns an NFT from given contract address
// @Description API user could check this directly via any third party service like NFTPort, Moralis as well
// @Tags        Common
// @Accept      json
// @Produce  json
// @Security BearerAuth
// @Param       contract path    string true "NFT Contract Address"
// @Param       wallet   path    string true "Wallet Address"
// @Success     200      {array} LandingPageItems
// @Router      /v1/is_owner/{contract}/{wallet} [get]
func IsOwner(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	contract := vars["contract"]
	wallet := vars["wallet"]

	result := false
	if strings.HasPrefix(wallet, "tz") {
		result = IsOwnerOfNFT(contract, wallet, "tezos")
	} else {
		result = IsOwnerOfNFT(contract, wallet, "ethereum")
		if !result {
			result = IsOwnerOfNFT(contract, wallet, "polygon")
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	json.NewEncoder(w).Encode(result)
}

// internal
func GetOwnerNFTs(walletAddr string, chain string) MoralisOwnerOf {

	//url := "https://eth-mainnet.alchemyapi.io/v2/${process.env.REACT_APP_ALCHEMY_API_KEY}/getOwnersForToken" + contractAddr
	//url := "https://api.nftport.xyz/v0/accounts/" + walletAddr + "?chain=" + chain
	url := "https://deep-index.moralis.io/api/v2.2/" + walletAddr + "/nft?chain=" + chain + "&format=decimal&normalizeMetadata=false"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Content-Type", "application/json")
	//req.Header.Add("Authorization", os.Getenv("NFTPORT_API_KEY"))
	req.Header.Add("X-API-Key", os.Getenv("MORALIS_NFT_API_KEY"))

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	var result MoralisOwnerOf
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON - GetOwnerNFTs", body)
	}

	//fmt.Printf("IsOwner: %#v\n", result.Total)

	return result
}

//internal use (NFTPORT version)
// func IsOwnerOfNFT(contractAddr string, walletAddr string, chain string) bool {
// 	//url := "https://eth-mainnet.alchemyapi.io/v2/${process.env.REACT_APP_ALCHEMY_API_KEY}/getOwnersForToken" + contractAddr
// 	url := "https://api.nftport.xyz/v0/accounts/" + walletAddr + "?chain=" + chain + "&contract_address=" + contractAddr

// 	req, _ := http.NewRequest("GET", url, nil)

// 	req.Header.Add("Content-Type", "application/json")
// 	req.Header.Add("Authorization", os.Getenv("NFTPORT_API_KEY"))

// 	// Send req using http Client
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Println("Error on response.\n[ERROR] -", err)
// 	}
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Println("Error while reading the response bytes:", err)
// 	}

// 	var result NFTPortOwnerOf
// 	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
// 		fmt.Println("Can not unmarshal JSON")
// 	}

// 	//fmt.Printf("IsOwner: %#v\n", result.Total)

// 	return result.Total > 0
// }

func IsOwnerOfNFT(contractAddr string, walletAddr string, chain string) bool {

	result := IsOwnerOfNftLocal(contractAddr, walletAddr, chain)
	//fmt.Println("IsOwnerOfNFT params / holder: ", contractAddr, walletAddr, chain, result)

	if !result && (chain == "ethereum" || chain == "polygon") {
		delegates := auth.GetDelegationsByDelegate(walletAddr)
		//fmt.Println("Wallet Delegates in OwnerOfNFT: ", delegates)
		for _, delegateWallet := range delegates {
			result = IsOwnerOfNftLocal(contractAddr, delegateWallet.Vault.Hex(), chain)
			if result {
				break
			} //if we find an NFT, can stop here
		}
	}

	return result
}

// internal - called from wrapper which checks DelegateCash as well
func IsOwnerOfNftLocal(contractAddr string, walletAddr string, chain string) bool {
	//For now if we use Moralis, ethereum needs to be "eth"
	if chain == "ethereum" {
		chain = "eth"
	}

	if chain == "near" {
		//fmt.Println("Near Chain ", walletAddr, contractAddr)
		url := "https://near-mainnet.api.pagoda.co/eapi/v1/accounts/" + walletAddr + "/NFT/" + contractAddr

		req, _ := http.NewRequest("GET", url, nil)

		req.Header.Add("accept", "application/json")
		req.Header.Add("X-API-Key", os.Getenv("PAGODA_NFT_API_KEY"))

		// Send req using http Client
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println("Error on response.\n[ERROR] -", err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error while reading the response bytes:", err)
		}

		var result NearOwnerOf
		if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
			fmt.Println("Can not unmarshal JSON - Near IsOwnerOfNFT", body)
		}
		//fmt.Printf("IsOwner: %#v\n", result.Nfts)

		return len(result.Nfts) > 0
	}
	if chain == "tezos" {
		//fmt.Println("Tezos Chain ", walletAddr, contractAddr)
		url := "https://api.tzkt.io/v1/tokens/balances?" + walletAddr + "&token.contract=" + contractAddr

		req, _ := http.NewRequest("GET", url, nil)

		req.Header.Add("accept", "application/json")

		// Send req using http Client
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println("Error on response.\n[ERROR] -", err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error while reading the response bytes:", err)
		}

		var result []TezosOwnerOf
		if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
			fmt.Println("Can not unmarshal JSON - Tezos IsOwnerOfNFT", body)
		}
		//fmt.Printf("IsOwner: %#v\n", result.Total)

		return len(result) > 0
	} else {
		url := "https://deep-index.moralis.io/api/v2.2/" + walletAddr + "/nft?chain=" + chain + "&format=decimal&token_addresses=" + contractAddr + "&normalizeMetadata=false"
		//url := "https://deep-index.moralis.io/api/v2.2/0x57ca1B13510D82a6286a225a217798e079BD0767/nft?chain=eth&format=decimal&token_addresses=0x34d85c9cdeb23fa97cb08333b511ac86e1c4e258&normalizeMetadata=false"

		req, _ := http.NewRequest("GET", url, nil)

		req.Header.Add("accept", "application/json")
		//req.Header.Add("Authorization", os.Getenv("NFTPORT_API_KEY"))
		req.Header.Add("X-API-Key", os.Getenv("MORALIS_NFT_API_KEY"))

		// Send req using http Client
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error on response.\n[ERROR] -", err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error while reading the response bytes:", err)
		}

		var result MoralisOwnerOf
		if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
			fmt.Println("Can not unmarshal JSON - IsOwnerOfNFT", body)
		}
		//fmt.Println("Moralis Returned Data - IsOwnerOfNFT", result)

		return len(result.Result) > 0
	}
}

func IsOwnerOfPOAP(eventId string, walletAddr string) bool {
	url := "https://api.poap.tech/actions/scan/" + walletAddr + "/" + eventId

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("X-API-Key", os.Getenv("POAP_API_KEY"))

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}

	return resp.StatusCode != 404
}

func IsOnChain(contractAddr string, chain string) bool {
	//For now if we use Moralis, ethereum needs to be "eth"
	if chain == "ethereum" {
		chain = "eth"
	}

	//url := "https://api.nftport.xyz/v0/nfts/" + contractAddr + "?chain=" + chain
	url := "https://deep-index.moralis.io/api/v2.2/nft/" + contractAddr + "?chain=" + chain + "&format=decimal&normalizeMetadata=false"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Content-Type", "application/json")
	//req.Header.Add("Authorization", os.Getenv("NFTPORT_API_KEY"))
	req.Header.Add("X-API-Key", os.Getenv("MORALIS_NFT_API_KEY"))

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	var result MoralisContractInfoNFT
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON - IsOnChain", body)
	}

	//fmt.Printf("Chain Response: %#v\n", result.Response)

	var returnVal = false
	if len(result.Result) > 0 {
		returnVal = true
	}
	return returnVal
}

// this was just used to fix up users info after adding new column
// not intended for extenal calls
func FixUpBookmarks(w http.ResponseWriter, r *http.Request) {
	var bookmarks []entity.Bookmarkitem
	database.Connector.Find(&bookmarks)

	for _, bookmark := range bookmarks {
		if strings.HasPrefix(bookmark.Nftaddr, "0x") {
			var result = IsOnChain(bookmark.Nftaddr, "ethereum")
			if result {
				database.Connector.Model(&entity.Bookmarkitem{}).Where("walletaddr = ?", bookmark.Walletaddr).Where("nftaddr = ?", bookmark.Nftaddr).Update("chain", "ethereum")
			} else {
				var result = IsOnChain(bookmark.Nftaddr, "polygon")
				if result {
					database.Connector.Model(&entity.Bookmarkitem{}).Where("walletaddr = ?", bookmark.Walletaddr).Where("nftaddr = ?", bookmark.Nftaddr).Update("chain", "polygon")
				}
			}
		}
	}
}

func AutoJoinCommunities(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	walletAddr := vars["wallet"]
	AutoJoinCommunitiesByChainWithDelegates(walletAddr, "ethereum")
	//AutoJoinCommunitiesByChainWithDelegates(walletAddr, "polygon")
	AutoJoinPoapChats(walletAddr)
}

// internal use only
func AutoJoinCommunitiesByChainWithDelegates(walletAddr string, chain string) {
	AutoJoinCommunitiesByChain(walletAddr, "", chain, walletAddr)

	//Check DelegateCash for NFTs owned
	delegates := auth.GetDelegationsByDelegate(walletAddr)
	for _, delegateWallet := range delegates {
		fmt.Println("Wallet Delegate Found: ", delegateWallet)
		//DelegateCash type 1 is a full wallet delegation
		//if so, lets allow delegate to to be part of all NFTs in Vault/Cold wallet
		if delegateWallet.Type == 1 {
			fmt.Println("Wallet Full Delegate: ", delegateWallet.Vault.Hex())
			AutoJoinCommunitiesByChain(delegateWallet.Vault.Hex(), "", chain, walletAddr)
		} else {
			AutoJoinCommunitiesByChain(delegateWallet.Vault.Hex(), delegateWallet.Contract.Hex(), chain, walletAddr)
		}
	}
}

// internal use only
// database.Connector.Where("walletaddr = ?", delegateAddr).Where("chain = ?", chain).Delete(&entity.Bookmarkitem{})
func AutoJoinCommunitiesByChain(walletAddr string, nftAddr string, chain string, delegateAddr string) {
	//For now if we use Moralis, ethereum needs to be "eth"
	if chain == "ethereum" {
		chain = "eth"
	}

	url := "https://deep-index.moralis.io/api/v2.2/" + walletAddr + "/nft?chain=" + chain + "&format=decimal&normalizeMetadata=false"
	if nftAddr != "" {
		fmt.Println("Auto join by Contract: ", nftAddr)
		url = "https://deep-index.moralis.io/api/v2.2/" + walletAddr + "/nft?chain=" + chain + "&format=decimal&normalizeMetadata=false&token_addresses%5B0%5D=" + nftAddr
	}

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Content-Type", "application/json")
	//req.Header.Add("Authorization", os.Getenv("NFTPORT_API_KEY"))
	req.Header.Add("X-API-Key", os.Getenv("MORALIS_NFT_API_KEY"))

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	var result MoralisOwnerOf
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON - AutoJoinCommunitiesByChain", body)
	}

	//fmt.Printf("IsOwner: %#v\n", result.Total)
	for _, nft := range result.Result {
		//TODO: could be optimized, good enough for now
		var bookmarkExists entity.Bookmarkitem
		var dbResult = database.Connector.Where("nftaddr = ?", nft.TokenAddress).Where("walletaddr = ?", delegateAddr).Find(&bookmarkExists)

		if dbResult.RowsAffected == 0 {
			//check if the user already manually unjoined, if so don't auto rejoin them
			var userUnjoined entity.Userunjoined
			var dbUnjoined = database.Connector.Where("nftaddr = ?", nft.TokenAddress).Where("walletaddr = ?", delegateAddr).Find(&userUnjoined)
			userAlreadyUnjoined := false
			if dbUnjoined.RowsAffected > 0 {
				userAlreadyUnjoined = userUnjoined.Unjoined
			}

			if !userAlreadyUnjoined {
				fmt.Println("Found new NFT: " + nft.TokenAddress)
				var bookmark entity.Bookmarkitem

				bookmark.Nftaddr = nft.TokenAddress
				bookmark.Walletaddr = delegateAddr //for normal cases delegate=walletAddr

				//For now if we use Moralis, ethereum needs to be "eth" (NFTPort on client side and LIT uses "ethereum")
				if chain == "eth" {
					chain = "ethereum"
				}
				bookmark.Chain = chain

				database.Connector.Create(&bookmark)
			}
		}
	}
}

// internal use only
func AutoJoinPoapChats(walletAddr string) {
	//https://documentation.poap.tech/reference/getactionsscan-5
	var poapInfo []POAPInfoByAddress = getPoapInfoByAddress(walletAddr)
	//fmt.Printf("AutoJoinPoapChats: %#v\n", poapInfo)
	for _, poap := range poapInfo {
		var bookmarkExists entity.Bookmarkitem

		var poapAddr = "poap_" + strconv.Itoa(poap.Event.ID)
		//fmt.Printf("POAP Event: %#v\n", poapAddr)
		var dbResult = database.Connector.Where("nftaddr = ?", poapAddr).Where("walletaddr = ?", walletAddr).Find(&bookmarkExists)

		if dbResult.RowsAffected == 0 {
			//check if the user already manually unjoined, if so don't auto rejoin them
			var userUnjoined entity.Userunjoined
			var dbUnjoined = database.Connector.Where("nftaddr = ?", poapAddr).Where("walletaddr = ?", walletAddr).Find(&userUnjoined)
			userAlreadyUnjoined := false
			if dbUnjoined.RowsAffected > 0 {
				userAlreadyUnjoined = userUnjoined.Unjoined
			}

			if !userAlreadyUnjoined {
				fmt.Printf("POAP is new for user: %#v\n", walletAddr)
				var bookmark entity.Bookmarkitem

				bookmark.Nftaddr = poapAddr
				bookmark.Walletaddr = walletAddr
				bookmark.Chain = poap.Chain

				database.Connector.Create(&bookmark)
			}
		}
	}
}

func GetPoapsByAddr(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	walletAddr := vars["wallet"]

	result := getPoapInfoByAddress(walletAddr)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	json.NewEncoder(w).Encode(result)
}

// internal use only
func getPoapInfoByAddress(walletAddr string) []POAPInfoByAddress {
	var result []POAPInfoByAddress
	url := "https://api.poap.tech/actions/scan/" + walletAddr

	if strings.HasPrefix(walletAddr, "tz") {
		return result
	}

	// Create a new request using http
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-API-KEY", os.Getenv("POAP_API_KEY"))

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	} else {
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error while reading the response bytes:", err)
		}

		if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
			fmt.Println("Can not unmarshal JSON - getPoapInfoByAddress", resp.Body, walletAddr)
		}
	}

	//fmt.Printf("returning: %#v\n", result)

	return result
}

type WalletGuardStruct struct {
	DomainName        string `json:"domainName"`
	RecommendedAction string `json:"recommendedAction"`
	RiskFactors       []struct {
		Type     string `json:"type"`
		Severity string `json:"severity"`
		Message  string `json:"message"`
		Value    string `json:"value,omitempty"`
	} `json:"riskFactors"`
	Verified bool   `json:"verified"`
	Status   string `json:"status"`
}

type WalletGuardResponse struct {
	RecommendedAction string `json:"recommendedAction"`
}

func WalletGuardCheck(w http.ResponseWriter, r *http.Request) {
	// Read the request body to get the list of URLs
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error requestBody: ", r.Body)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Parse the list of URLs from the request body
	var urls []string
	if err := json.Unmarshal(requestBody, &urls); err != nil {
		fmt.Println("Parse Error : ", urls)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Initialize a slice to store RecommendedActions
	recommendedActions := make([]string, len(urls))

	for i, url := range urls {
		// Create the URL for the WalletGuard API
		apiUrl := "https://api.walletguard.app/v1/scan?url=" + url

		// Create the HTTP request
		req, err := http.NewRequest("GET", apiUrl, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		req.Header.Add("accept", "application/json")
		req.Header.Add("X-API-KEY", os.Getenv("WALLET_GUARD_API_KEY"))

		// Send the HTTP request
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer res.Body.Close()

		// Read and parse the response body
		responseBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Unmarshal the JSON response to get the RecommendedAction
		var walletGuardResponse WalletGuardResponse
		if err := json.Unmarshal(responseBody, &walletGuardResponse); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Store the RecommendedAction in the slice
		recommendedActions[i] = walletGuardResponse.RecommendedAction
	}

	// Now, recommendedActions contains the RecommendedAction for each URL
	//fmt.Printf("Recommended Actions: %v\n", recommendedActions)

	// Respond with the results as JSON
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	json.NewEncoder(w).Encode(recommendedActions)
}

func RegisterOuraUser(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var newUserTemp entity.Ourausertemp
	json.Unmarshal(requestBody, &newUserTemp)

	//verify its a real and valid API key
	url := "https://api.ouraring.com/v2/usercollection/daily_activity"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	req.Header.Add("Authorization", "Bearer "+newUserTemp.Pac)

	res, err := client.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Println("Error making request:", err)
		return
	}
	defer res.Body.Close()

	// Check if the response status code is not 200
	fmt.Println("Attempt to register user: ", newUserTemp.Pac, newUserTemp.Nickname, newUserTemp.Referralcode)
	if res.StatusCode != http.StatusOK {
		fmt.Printf("RegisterOuraUser Unexpected response code: %d\n", res.StatusCode)
		w.WriteHeader(http.StatusForbidden)
		return
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err, string(body))
		return
	}
	//fmt.Println("Test Register User Check: ", string(body))

	// Authuser := auth.GetUserFromReqContext(r)
	// if strings.EqualFold(Authuser.Address, newUser.Wallet) {

	var existinguser entity.Ourauser
	var pacAlreadyExists = database.Connector.Where("pac = ?", newUserTemp.Pac).Find(&existinguser)
	if pacAlreadyExists.RowsAffected > 0 {
		fmt.Println("PAC already registered: ", newUserTemp.Wallet, newUserTemp.Pac)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	//Now actually register the user, redeem referral code and set nickname
	var newUser entity.Ourauser
	newUser.Pac = newUserTemp.Pac
	newUser.Wallet = newUserTemp.Wallet
	newUser.Signature = newUserTemp.Signature

	// Public key as PEM string
	publicKeyPEM := `-----BEGIN PUBLIC KEY-----
MIIBojANBgkqhkiG9w0BAQEFAAOCAY8AMIIBigKCAYEA4129oK+dUEalpqP5aT/M
A6yhFbAjNppOidQuVgeSgEPquXLlJrdLoomHGhzugbBYeKS6lceEDM3oygFdCGhT
sly26Ws8qyUIGlk0/JGf4mRHd9RMs0uOF50/mB4abNM/mA/k8cO46+UmXOK2rwEL
U2rPb5tWVzxjPqs8Aw9eT1n7UlvOXxFc4ChyIHX/plfbkKK1R1+PYhtBHeQT8aW1
o7wLsbbnkCGh2iahJaNacMWmUZ9YygdPg2DICQLK2KbZfZHhhylBjDzuBgjUzNai
ikVHzrR6f9eTihYjmpx8Br5Ubhj3lVt45nAXFidxMBe1e7IILNVl9C57sqV+nPFM
2s5ad/r3TDjOZ23e0FGBVsyG+lJwn9q/kx4kjSFsO8fNzJ7wUczVnfW+akox2rMX
rnvdxUhpAAEtJZme5+pnS6Fr4Zi8mUBPt9kC/mHTtbPQoLsX+FeBs/u+rpXe4xBr
+QhqShKWQ+4HzwQHCc5h9d4pqZEKK8UnpdeJ0c/QTqcVAgMBAAE=
-----END PUBLIC KEY-----`

	encryptedSecret, _ := vanaencrypt.EncryptSecretForProof(publicKeyPEM, []byte(newUser.Pac))
	newUser.Encryptedpac = encryptedSecret

	var existinguser2 entity.Ourauser
	var walletAlreadyExists = database.Connector.Where("wallet = ?", newUser.Wallet).Find(&existinguser2)
	if walletAlreadyExists.RowsAffected > 0 {
		fmt.Println("wallet updated: ", newUser.Wallet, newUser.Pac)
		database.Connector.Model(&entity.Ourauser{}).Where("wallet = ?", newUser.Wallet).Update("pac", newUser.Pac)
		database.Connector.Model(&entity.Ourauser{}).Where("wallet = ?", newUser.Wallet).Update("encrypted_pac", newUser.Encryptedpac)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newUser)
		return
	}

	//give new users 3 new referral codes
	referrals.CreateReferralCodeInternal(newUser.Wallet)

	if newUserTemp.Referralcode != "" {
		database.Connector.Model(&entity.Referralcode{}).
			Where("code = ?", newUserTemp.Referralcode).
			Update("redeemed", true)

		//set user as validated in the referral code table (used separate table in the case we drop this in future)
		var uservalid entity.Referraluser
		uservalid.Referralcode = newUserTemp.Referralcode
		uservalid.Walletaddr = newUserTemp.Wallet
		database.Connector.Create(&uservalid)
	}

	if newUserTemp.Nickname != "" {
		var addrnameDB entity.Addrnameitem
		var addrNameItem entity.Addrnameitem
		addrNameItem.Address = newUserTemp.Wallet
		addrNameItem.Name = newUserTemp.Nickname
		var dbQuery = database.Connector.Where("address = ?", newUserTemp.Wallet).Find(&addrnameDB)
		if dbQuery.RowsAffected == 0 {
			database.Connector.Create(&addrNameItem)
		}

		//this isn't great long-term, but for smaller table it allows new usernames to show up in table instead of waiting for once daily update
		referrals.GetOuraLeaderboardDataCronJob()
	}

	wc_analytics.SendCustomIntraEvent(newUser.Wallet, "NEW_OURA_REGISTRATION")

	database.Connector.Create(&newUser)
	fmt.Println("New PAC User: ", newUser.Wallet)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
	// } else {
	// 	w.WriteHeader(http.StatusForbidden)
	// }
}

var ouraEndpoints = []string{
	"daily_activity",
	"daily_cardiovascular_age",
	"daily_readiness",
	"daily_resilience",
	"daily_sleep",
	"daily_spo2",
	"daily_stress",
	// "heartrate",
	"rest_mode_period",
	"sleep_time",
	"vO2_max",
	"workout",
}

func addFileToZip(zipWriter *zip.Writer, fileName string, data []byte) error {
	// Create a zip file header
	header := &zip.FileHeader{
		Name:   fileName,
		Method: zip.Deflate, // Use Deflate compression
	}

	// Create the writer for the file in the zip
	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	// Write the raw JSON bytes (data) to the zip
	_, err = writer.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func FetchAndDecryptFile(fileUrl string, decryptionKey string) error {
	// Fetch the file from the URL
	resp, err := http.Get(fileUrl)
	if err != nil {
		return fmt.Errorf("failed to fetch file: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	encryptedData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// Decrypt the data using clientSideDecryption with the provided key
	decryptedData, err := vanaencrypt.ClientSideDecrypt(encryptedData, decryptionKey)
	if err != nil {
		return fmt.Errorf("failed to decrypt data: %w", err)
	}

	// Save the decrypted data to a file in the root directory
	err = ioutil.WriteFile("decrypted_file.zip", decryptedData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write decrypted file: %w", err)
	}

	fmt.Println("decrypted file: ", decryptedData)

	return nil
}

func FetchOuraData() {
	var ourausers []entity.Ourauser
	database.Connector.Find(&ourausers)

	for _, ourauser := range ourausers {
		//skip test users
		if len(ourauser.Signature) < 1 {
			continue
		}
		fmt.Println("Fetching Daily Data for: ", ourauser.Wallet)
		// Create a buffer to hold the zip data
		var zipFileBuf bytes.Buffer
		// Create a new zip writer
		zipWriter := zip.NewWriter(&zipFileBuf)

		//DLP public Key
		publicKeyDLP, err := vanatransact.GetDlpPublicKey()
		if err != nil {
			continue
		}
		fmt.Println("DLP encryption publicKey: ", publicKeyDLP)
		//encrypt data client using signature of a fixed message (tbd - how to do as proxy?)
		for _, endpoint := range ouraEndpoints {
			url := "https://api.ouraring.com/v2/usercollection/" + endpoint
			method := "GET"

			client := &http.Client{}
			req, err := http.NewRequest(method, url, nil)
			if err != nil {
				fmt.Println(err)
				break
			}
			req.Header.Add("Authorization", "Bearer "+ourauser.Pac)

			res, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
				break
			}
			defer res.Body.Close()

			body, err := io.ReadAll(res.Body)
			if err != nil {
				fmt.Println(err)
				break
			}

			// Decode the JSON body into a generic interface
			var jsonData interface{}
			if err := json.Unmarshal(body, &jsonData); err != nil {
				fmt.Println("Failed to unmarshal JSON:", err)
				break
			}

			// Marshal the JSON data with indentation
			formattedJSON, err := json.MarshalIndent(jsonData, "", "  ")
			if err != nil {
				fmt.Println("Failed to marshal JSON with indentation:", err)
				break
			}

			var currentData entity.Ouradata
			currentData.Endpoint = endpoint
			currentData.Wallet = ourauser.Wallet
			currentData.Jsondata = string(formattedJSON)
			database.Connector.Create(&currentData)

			// Add formatted JSON to the zip
			if err := addFileToZip(zipWriter, endpoint+".json", formattedJSON); err != nil {
				log.Fatalf("Failed to add "+endpoint+".json to zip: %v", err)
			}
		}

		// Close the zip writer
		if err := zipWriter.Close(); err != nil {
			log.Fatalf("Failed to close zip writer: %v", err)
		}

		//Vana stuff:
		//for now since we use PAC - we upload the whole dataset as one zip (cheaper to verify via TEE)
		//in the future if users choose to only share specific data - we need to upload each endpoint separately.

		// Open the zip file (for testing alg and formatting)
		// zipData, err := os.ReadFile("archive.zip")
		// if err != nil {
		// 	log.Fatalf("Failed to read zip file data: %v", err)
		// }

		//encrypt the file encryption key with the users signature (TODO fill this based on user signature)
		encryptedBytes, err := vanaencrypt.ClientSideEncrypt(zipFileBuf.Bytes(), ourauser.Signature)
		if err != nil {
			fmt.Println("error in ClientSideEncrypt", err)
		}

		// Upload the zip file to DigitalOcean Spaces
		fileUrl, err := SaveFileToSpaces(encryptedBytes, ourauser.Wallet+time.Now().Format("2006-01-02_15-04-05")+"_archive.zip")
		if err != nil {
			log.Fatalf("Failed to upload to DigitalOcean Spaces: %v", err)
		}
		fmt.Println("file stored at: ", fileUrl)

		//for test
		// err = FetchAndDecryptFile(fileUrl, ourauser.Signature)
		// if err != nil {
		// 	fmt.Println("error: ", err)
		// }

		// iv := []byte{
		// 	169, 138, 29, 49, 139, 11, 183, 51,
		// 	167, 5, 144, 163, 203, 214, 217, 224,
		// }
		// ephemPrivateKeyBytes := []byte{
		// 	147, 207, 81, 186, 169, 91, 245, 42,
		// 	148, 220, 122, 136, 222, 82, 10, 86,
		// 	230, 210, 241, 85, 15, 154, 77, 60,
		// 	38, 91, 211, 211, 243, 2, 214, 203,
		// }
		// Initialize the IV and ephemeral private key with random bytes if was not provided
		iv := make([]byte, 16) // 16 bytes for IV
		if _, err := rand.Read(iv); err != nil {
			continue
		}
		ephemPrivateKeyBytes := make([]byte, 32) // 32 bytes for ephemeral private key
		if _, err := rand.Read(ephemPrivateKeyBytes); err != nil {
			continue
		}

		//get EEK with EK
		vanaDlpEEK, _, _ := vanaencrypt.EncryptWithWalletPublicKey(ourauser.Signature, publicKeyDLP, iv, ephemPrivateKeyBytes)
		finalDlpEEK := append(vanaDlpEEK["iv"],
			append(vanaDlpEEK["ephemPublicKey"],
				append(vanaDlpEEK["ciphertext"], vanaDlpEEK["mac"]...)...)...)

		// Return the final result as a hex string
		hexDataDlpEEK := hex.EncodeToString(finalDlpEEK)
		fmt.Println("DLP EEK: ", hexDataDlpEEK)

		//function - addFileWithPermissions - blockchain RPC call
		//parameters - (publicly accessible link to encrypted data, "permissions" is the encrypted encryption key)
		// returns fileID (ex: file id is '601971')
		walletAddress := common.HexToAddress(ourauser.Wallet) //ourauser.Wallet)
		txHash, err := vanatransact.AddFileWithPermissions(walletAddress, fileUrl, hexDataDlpEEK)
		if err != nil {
			fmt.Println("Uploaded File  err: ", txHash, err)
			continue
		}
		fmt.Println("Uploaded File TX: ", txHash)
		var fileID = vanatransact.GetFileID(txHash)
		//fmt.Println("Uploaded File: ", fileID)

		// ** BEGIN TEST CODE ** test that the DLP can decrypt the file: TEST CODE ONLY
		// privKey, _ := hex.DecodeString(os.Getenv("VANA_SIGNER_PRIVATE_KEY"))
		// decryptedDlpEEK, err := vanaencrypt.DecryptWithPrivKey(privKey, vanaDlpEEK)
		// if err != nil {
		// 	fmt.Println("Error in testing decryption of EEK to get back user signature: ", err)
		// }
		// // Convert the ASCII byte array back to a string
		// asciiString := string(decryptedDlpEEK)
		// // Remove the "0x" prefix if present
		// if len(asciiString) > 2 && asciiString[:2] == "0x" {
		// 	asciiString = asciiString[2:]
		// }
		// // Decode the cleaned hex string back to bytes
		// decodedBytes, err := hex.DecodeString(asciiString)
		// if err != nil {
		// 	fmt.Println("Error decoding hex:", err)
		// 	return
		// }
		// // Convert the decoded bytes to the final hex representation
		// finalHex := fmt.Sprintf("0x%x", decodedBytes)
		// fmt.Println("decrypted EEK: ", finalHex)
		// fmt.Println("Orig user sig: ", ourauser.Signature)
		// err = FetchAndDecryptFile(fileUrl, finalHex)
		// if err != nil {
		// 	fmt.Println("error: ", err)
		// }
		//** END ** test that the DLP can decrypt the file: TEST CODE ONLY

		//now get proof from TEE the file is valid / authentic
		//teeFee = await teePoolContract.teeFee(); //get estimated required fee for proof?
		var teePrice = vanatransact.GetTeePrice()
		fmt.Println("TEE Price: ", teePrice)

		var contributionProofTx = vanatransact.GetTeeContributionProof(fileID)
		fmt.Println("TEE contribution proof tx: ", contributionProofTx)

		time.Sleep(45 * time.Second)
		// Call the API to fetch transaction logs
		url := "https://api.moksha.vanascan.io/api/v2/transactions/" + contributionProofTx + "/logs"
		resp, err := http.Get(url)
		if err != nil {
			log.Println("Error fetching transaction logs:", err)
			return // Handle the error appropriately
		}
		defer resp.Body.Close()

		// Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error reading response body:", err)
			return // Handle the error appropriately
		}
		fmt.Println(string(body))

		// Define a struct to match the JSON response
		type Response struct {
			Items []struct {
				Decoded struct {
					Parameters []struct {
						Name  string `json:"name"`
						Value string `json:"value"`
					} `json:"parameters"`
				} `json:"decoded"`
			} `json:"items"`
		}

		// Unmarshal the JSON response into the struct
		var responseData Response
		if err := json.Unmarshal(body, &responseData); err != nil {
			log.Println("Error unmarshalling JSON:", err)
			return // Handle the error appropriately
		}

		// Extract the jobId from the response
		if len(responseData.Items) > 0 {
			for _, param := range responseData.Items[0].Decoded.Parameters {
				if param.Name == "jobId" {
					log.Println("Job ID from TX log:", param.Value)
					break
				}
			}
		} else {
			log.Println("No items found in the response.")
		}

		//getJobId and teeDetails (tbd)
		fileIDBigInt := new(big.Int)
		fileIDBigInt.SetString(fileID[2:], 16) // Skip the "0x" prefix
		// Convert big.Int to string representation of the integer
		fileIDstr := fileIDBigInt.String()
		var jobIDS = vanatransact.GetFileJobIDs(fileIDstr)
		fmt.Println("`Latest JobIDs for FileID ", jobIDS, fileIDstr)

		if len(jobIDS) > 0 {
			latestJobId := jobIDS[len(jobIDS)-1]
			teeUrl, teePublicKey := vanatransact.GetTeeDetails(*latestJobId)

			//specific to the DLP proof code
			//envVars := map[string]string{}
			//need this for testing archive.zip manually!
			// envVars := map[string]string{
			// 	"USER_EMAIL": "user123@gmail.com", // Add USER_EMAIL to EnvVars
			// }
			envVars := map[string]string{}

			//secrets := map[string]string{}
			//encryptedSecret, _ := vanaencrypt.EncryptSecretForProof(publicKeyPEM, []byte("user123@gmail.com"))
			secrets := map[string]string{
				"USER_API_KEY": ourauser.Encryptedpac, // Add USER_EMAIL to EnvVars
			} //this would be API keys, etc needed in proof code

			//ask a specific TEE to run the proof of contribution
			//${jobDetails.teeUrl}/RunProof
			err := vanatransact.SendContributionProof(latestJobId, fileID, publicKeyDLP, envVars, secrets, teePublicKey, teeUrl, iv, ephemPrivateKeyBytes, ourauser.Signature)
			if err != nil {
				fmt.Println("Error in SendContributionProof", err)
				continue
			}

			//now request reward from DLP contract
			txHashReward, err := vanatransact.RequestRewardFromDLP(fileID)
			fmt.Println("Request Reward from DLP: ", txHashReward, err)
		}
	}
}

type POAPInfoByAddress struct {
	Event struct {
		ID          int    `json:"id"`
		FancyID     string `json:"fancy_id"`
		Name        string `json:"name"`
		EventURL    string `json:"event_url"`
		ImageURL    string `json:"image_url"`
		Country     string `json:"country"`
		City        string `json:"city"`
		Description string `json:"description"`
		Year        int    `json:"year"`
		StartDate   string `json:"start_date"`
		EndDate     string `json:"end_date"`
		ExpiryDate  string `json:"expiry_date"`
		Supply      int    `json:"supply"`
	} `json:"event"`
	TokenID string `json:"tokenId"`
	Owner   string `json:"owner"`
	Chain   string `json:"chain"`
	Created string `json:"created"`
}

type NFTPortOwnerOf struct {
	Response string `json:"response"`
	Nfts     []struct {
		ContractAddress string `json:"contract_address"`
		TokenID         string `json:"token_id"`
		CreatorAddress  string `json:"creator_address"`
	} `json:"nfts"`
	Total        int         `json:"total"`
	Continuation interface{} `json:"continuation"`
}

type MoralisOwnerOf struct {
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
	Cursor   interface{} `json:"cursor"`
	Result   []struct {
		TokenAddress      string    `json:"token_address"`
		TokenID           string    `json:"token_id"`
		OwnerOf           string    `json:"owner_of"`
		BlockNumber       string    `json:"block_number"`
		BlockNumberMinted string    `json:"block_number_minted"`
		TokenHash         string    `json:"token_hash"`
		Amount            string    `json:"amount"`
		ContractType      string    `json:"contract_type"`
		Name              string    `json:"name"`
		Symbol            string    `json:"symbol"`
		TokenURI          string    `json:"token_uri"`
		Metadata          string    `json:"metadata"`
		LastTokenURISync  time.Time `json:"last_token_uri_sync"`
		LastMetadataSync  time.Time `json:"last_metadata_sync"`
		MinterAddress     string    `json:"minter_address"`
	} `json:"result"`
	Status string `json:"status"`
}

type MoralisContractInfoNFT struct {
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	Cursor   string `json:"cursor"`
	Result   []struct {
		TokenAddress       string    `json:"token_address"`
		TokenID            string    `json:"token_id"`
		Amount             string    `json:"amount"`
		TokenHash          string    `json:"token_hash"`
		BlockNumberMinted  string    `json:"block_number_minted"`
		UpdatedAt          any       `json:"updated_at"`
		ContractType       any       `json:"contract_type"`
		Name               string    `json:"name"`
		Symbol             string    `json:"symbol"`
		TokenURI           any       `json:"token_uri"`
		Metadata           string    `json:"metadata"`
		LastTokenURISync   time.Time `json:"last_token_uri_sync"`
		LastMetadataSync   time.Time `json:"last_metadata_sync"`
		MinterAddress      any       `json:"minter_address"`
		PossibleSpam       bool      `json:"possible_spam"`
		VerifiedCollection bool      `json:"verified_collection"`
	} `json:"result"`
}

// Address -> Name lookup
type BtcStacksName struct {
	Names []string `json:"names"`
}

// Name -> address lookup
type BtcStacksAddress struct {
	Address      string `json:"address"`
	Blockchain   string `json:"blockchain"`
	ExpireBlock  int    `json:"expire_block"`
	LastTxid     string `json:"last_txid"`
	Status       string `json:"status"`
	Zonefile     string `json:"zonefile"`
	ZonefileHash string `json:"zonefile_hash"`
}

type TezosOwnerOf struct {
	ID      int64 `json:"id"`
	Account struct {
		Address string `json:"address"`
	} `json:"account"`
	Token struct {
		ID       int64 `json:"id"`
		Contract struct {
			Address string `json:"address"`
		} `json:"contract"`
		TokenID     string `json:"tokenId"`
		Standard    string `json:"standard"`
		TotalSupply string `json:"totalSupply"`
		Metadata    struct {
			Name      string        `json:"name"`
			Image     string        `json:"image"`
			Rights    string        `json:"rights"`
			Symbol    string        `json:"symbol"`
			Formats   []interface{} `json:"formats"`
			Creators  []string      `json:"creators"`
			Decimals  string        `json:"decimals"`
			Royalties struct {
				Shares struct {
					Address string `json:"address"`
				} `json:"shares"`
				Decimals string `json:"decimals"`
			} `json:"royalties"`
			Attributes         []interface{} `json:"attributes"`
			DisplayURI         string        `json:"displayUri"`
			ArtifactURI        string        `json:"artifactUri"`
			Description        string        `json:"description"`
			ThumbnailURI       string        `json:"thumbnailUri"`
			IsBooleanAmount    bool          `json:"isBooleanAmount"`
			ShouldPreferSymbol bool          `json:"shouldPreferSymbol"`
		} `json:"metadata"`
	} `json:"token"`
	Balance        string    `json:"balance"`
	TransfersCount int       `json:"transfersCount"`
	FirstLevel     int       `json:"firstLevel"`
	FirstTime      time.Time `json:"firstTime"`
	LastLevel      int       `json:"lastLevel"`
	LastTime       time.Time `json:"lastTime"`
}

type NearOwnerOf struct {
	Nfts []struct {
		TokenID        string `json:"token_id"`
		OwnerAccountID string `json:"owner_account_id"`
		Metadata       struct {
			Title         string      `json:"title"`
			Description   interface{} `json:"description"`
			Media         string      `json:"media"`
			MediaHash     interface{} `json:"media_hash"`
			Copies        int         `json:"copies"`
			Extra         interface{} `json:"extra"`
			Reference     string      `json:"reference"`
			ReferenceHash interface{} `json:"reference_hash"`
		} `json:"metadata"`
	} `json:"nfts"`
	ContractMetadata struct {
		Spec          string      `json:"spec"`
		Name          string      `json:"name"`
		Symbol        string      `json:"symbol"`
		Icon          string      `json:"icon"`
		BaseURI       string      `json:"base_uri"`
		Reference     interface{} `json:"reference"`
		ReferenceHash interface{} `json:"reference_hash"`
	} `json:"contract_metadata"`
	BlockTimestampNanos string `json:"block_timestamp_nanos"`
	BlockHeight         string `json:"block_height"`
}

type NFTPortNftContract struct {
	Response string `json:"response"`
	Nfts     []struct {
		Chain           string `json:"chain"`
		ContractAddress string `json:"contract_address"`
		TokenID         string `json:"token_id"`
	} `json:"nfts"`
	Contract struct {
		Name     string `json:"name"`
		Symbol   string `json:"symbol"`
		Type     string `json:"type"`
		Metadata struct {
			Description        string `json:"description"`
			ThumbnailURL       string `json:"thumbnail_url"`
			CachedThumbnailURL string `json:"cached_thumbnail_url"`
			BannerURL          string `json:"banner_url"`
			CachedBannerURL    string `json:"cached_banner_url"`
		} `json:"metadata"`
	} `json:"contract"`
	Total int `json:"total"`
}

type User struct {
	Username        string `json:"username"`
	ProfileImageURL string `json:"profile_image_url"`
	ID              string `json:"id"`
	Name            string `json:"name"`
}

type Attachments struct {
	MediaKeys []string `json:"media_keys"`
}

type TwitterTweetsData struct {
	Data []struct {
		Text        string `json:"text"`
		ID          string `json:"id"`
		Attachments struct {
			MediaKeys []string `json:"media_keys"`
		} `json:"attachments,omitempty"`
		AuthorID  string    `json:"author_id"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"data"`
	Includes struct {
		Media []struct {
			Type            string `json:"type"`
			Width           int    `json:"width"`
			PreviewImageURL string `json:"preview_image_url,omitempty"`
			Height          int    `json:"height"`
			MediaKey        string `json:"media_key"`
			URL             string `json:"url,omitempty"`
		} `json:"media"`
		Users []struct {
			Username        string `json:"username"`
			ProfileImageURL string `json:"profile_image_url"`
			ID              string `json:"id"`
			Name            string `json:"name"`
		} `json:"users"`
	} `json:"includes"`
	Meta struct {
		NextToken   string `json:"next_token"`
		ResultCount int    `json:"result_count"`
		NewestID    string `json:"newest_id"`
		OldestID    string `json:"oldest_id"`
	} `json:"meta"`
}

// formatted for use in client side per Mana
type TweetType struct {
	Text        string `json:"text"`
	ID          string `json:"id"`
	Attachments struct {
		MediaKeys []string `json:"media_keys"`
	} `json:"attachments"`
	AuthorID  string    `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
	User      struct {
		Username        string `json:"username"`
		ProfileImageURL string `json:"profile_image_url"`
		ID              string `json:"id"`
		Name            string `json:"name"`
	} `json:"user"`
	Media Attachments `json:"media"`
}

type TwitterIdResp struct {
	Data struct {
		Id       string `json:"id"`
		Name     string `json:"name"`
		Username string `json:"username"`
	} `json:"data"`
}

type Social struct {
	SocialMsg []string `json:"social"`
}

type SocialMsg struct {
	Type     string `json:"type"`
	Username string `json:"username"`
}

type CommunityMember struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Image   string `json:"image"`
	Admin   bool   `json:"admin"`
}
type LandingPageItems struct {
	Name        string                 `json:"name"`
	MemberCount int                    `json:"member_count"`
	Members     []CommunityMember      `json:"members"`
	Logo        string                 `json:"logo"`         // logo url, stored in backend
	Verified    bool                   `json:"is_verified"`  // is this group verified? WalletChat's group is verified by default
	Joined      bool                   `json:"joined"`       //number of members of the group
	Messaged    bool                   `json:"has_messaged"` // has user messaged in this group chat before? if not show "Say hi" button
	Messages    []entity.Groupchatitem `json:"messages"`
	Tweets      []TweetType            `json:"tweets"` // follow format of GET /get_twitter/{nftAddr}
	Social      []SocialMsg            `json:"social"`
}

type OpenseaData struct {
	Collection struct {
		BannerImageURL          string      `json:"banner_image_url"`
		ChatURL                 interface{} `json:"chat_url"`
		CreatedDate             string      `json:"created_date"`
		DefaultToFiat           bool        `json:"default_to_fiat"`
		Description             string      `json:"description"`
		DevBuyerFeeBasisPoints  string      `json:"dev_buyer_fee_basis_points"`
		DevSellerFeeBasisPoints string      `json:"dev_seller_fee_basis_points"`
		DiscordURL              string      `json:"discord_url"`
		DisplayData             struct {
			CardDisplayStyle string `json:"card_display_style"`
		} `json:"display_data"`
		ExternalURL                 string      `json:"external_url"`
		Featured                    bool        `json:"featured"`
		FeaturedImageURL            string      `json:"featured_image_url"`
		Hidden                      bool        `json:"hidden"`
		SafelistRequestStatus       string      `json:"safelist_request_status"`
		ImageURL                    string      `json:"image_url"`
		IsSubjectToWhitelist        bool        `json:"is_subject_to_whitelist"`
		LargeImageURL               string      `json:"large_image_url"`
		MediumUsername              string      `json:"medium_username"`
		Name                        string      `json:"name"`
		OnlyProxiedTransfers        bool        `json:"only_proxied_transfers"`
		OpenseaBuyerFeeBasisPoints  string      `json:"opensea_buyer_fee_basis_points"`
		OpenseaSellerFeeBasisPoints string      `json:"opensea_seller_fee_basis_points"`
		PayoutAddress               string      `json:"payout_address"`
		RequireEmail                bool        `json:"require_email"`
		ShortDescription            interface{} `json:"short_description"`
		Slug                        string      `json:"slug"`
		TelegramURL                 interface{} `json:"telegram_url"`
		TwitterUsername             string      `json:"twitter_username"`
		InstagramUsername           string      `json:"instagram_username"`
		WikiURL                     interface{} `json:"wiki_url"`
		IsNsfw                      bool        `json:"is_nsfw"`
	} `json:"collection"`
	Address                     string      `json:"address"`
	AssetContractType           string      `json:"asset_contract_type"`
	CreatedDate                 string      `json:"created_date"`
	Name                        string      `json:"name"`
	NftVersion                  string      `json:"nft_version"`
	OpenseaVersion              interface{} `json:"opensea_version"`
	Owner                       int         `json:"owner"`
	SchemaName                  string      `json:"schema_name"`
	Symbol                      string      `json:"symbol"`
	TotalSupply                 string      `json:"total_supply"`
	Description                 string      `json:"description"`
	ExternalLink                string      `json:"external_link"`
	ImageURL                    string      `json:"image_url"`
	DefaultToFiat               bool        `json:"default_to_fiat"`
	DevBuyerFeeBasisPoints      int         `json:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints     int         `json:"dev_seller_fee_basis_points"`
	OnlyProxiedTransfers        bool        `json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  int         `json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints int         `json:"opensea_seller_fee_basis_points"`
	BuyerFeeBasisPoints         int         `json:"buyer_fee_basis_points"`
	SellerFeeBasisPoints        int         `json:"seller_fee_basis_points"`
	PayoutAddress               string      `json:"payout_address"`
}

type TelegramUpdateNotifsData struct {
	Ok     bool `json:"ok"`
	Result []struct {
		UpdateID int `json:"update_id"`
		Message  struct {
			MessageID int `json:"message_id"`
			From      struct {
				ID           int64  `json:"id"`
				IsBot        bool   `json:"is_bot"`
				FirstName    string `json:"first_name"`
				LastName     string `json:"last_name"`
				Username     string `json:"username"`
				LanguageCode string `json:"language_code"`
			} `json:"from"`
			Chat struct {
				ID        int64  `json:"id"`
				FirstName string `json:"first_name"`
				LastName  string `json:"last_name"`
				Username  string `json:"username"`
				Type      string `json:"type"`
			} `json:"chat"`
			ReplyToMessage *struct {
				MessageID int `json:"message_id"`
				From      struct {
					ID        int64  `json:"id"`
					IsBot     bool   `json:"is_bot"`
					FirstName string `json:"first_name"`
					Username  string `json:"username"`
				} `json:"from"`
				Chat struct {
					ID                          int    `json:"id"`
					Title                       string `json:"title"`
					Type                        string `json:"type"`
					AllMembersAreAdministrators bool   `json:"all_members_are_administrators"`
				} `json:"chat"`
				Date int    `json:"date"`
				Text string `json:"text"`
			} `json:"reply_to_message,omitempty"`
			Date int    `json:"date"`
			Text string `json:"text"`
		} `json:"message"`
	} `json:"result"`
}
