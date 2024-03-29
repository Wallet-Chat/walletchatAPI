package entity

import "time"

//*** VARIABLE NAMING INFO DUE TO GORM FUNCTIONALITY ***
//CamelCase is not used for variable names due to GORM auto functionality.
//A struct named ChatItem would go into the database as chat_items - the case auto-adds _ chars
//table names within a database have an 's' auto-added to them
//...(Chatitem struct here will have a table named Chatitems in the DB)

//rename the type in GET /inbox to context_type: [nft, community, dm] and
//retain variable name type in message objects in communities to be [welcome, message] instead of [communitymsg, communitywelcome]
// string mapping
const ( //context_type mapping just for bookkeeping(golang sucks for enums as well...)
	Nft       string = "nft"
	Community string = "community"
	DM        string = "dm"
	All       string = "all"
)
const ( //type mapping just for bookkeeping(golang sucks for enums as well...)
	Welcome string = "welcome"
	Message string = "message"
)

type Unreadcountitem struct {
	//Id       int    `gorm:"primaryKey;autoIncrement"`
	//Walletaddr string `json:"walletaddr"`
	Nft       int `json:"nft"`
	Dm        int `json:"dm"`
	Community int `json:"community"`
}

type Chatitem struct {
	Id            int       `gorm:"primary_key"`                 //AUTO-GENERATED (PRIMARY KEY)
	Fromaddr      string    `json:"fromaddr" binding:"required"` //*** REQUIRED INPUT ***
	Toaddr        string    `json:"toaddr" validate:"required"`  //*** REQUIRED INPUT ***
	Timestamp     string    `json:"timestamp"`                   //AUTO-SET BY REST API
	Timestamp_dtm time.Time `json:"timestamp_dtm"`               //USED FOR SORTING WHEN TIME FORMAT NEEDED
	Msgread       bool      `json:"read"`                        //DEFAULT FALSE
	Message       string    `json:"message" validate:"required"` //*** REQUIRED INPUT ***
	Nftaddr       string    `json:"nftaddr"`                     //ONLY USED FOR NFT DM CONTEXT
	Nftid         string    `json:"nftid"`                       //ONLY USED FOR NFT DM CONTEXT
	Name          string    `json:"sender_name"`                 //AUTO-SET BY BACKED FOR RETURN VALUE
	Encryptsymkey string    `json:"encrypted_sym_lit_key"`       //USE IF USING LIT ENCRYPTION
	Litaccesscond string    `json:"lit_access_conditions"`
}

//for olivers view function
type V_chatitem struct {
	Id            int       `gorm:"primaryKey"`
	Fromaddr      string    `json:"fromaddr"`
	Toaddr        string    `json:"toaddr"`
	Timestamp     string    `json:"timestamp"`
	Timestamp_dtm time.Time `json:"timestamp_dtm"`
	Msgread       bool      `json:"read"`
	Message       string    `json:"message"`
	Nftaddr       string    `json:"nftaddr"`
	NftId         string    `json:"nftid"`
	Name          string    `json:"sender_name"`
	Encryptsymkey string    `json:"encrypted_sym_lit_key"` //USE IF USING LIT ENCRYPTION
	Litaccesscond string    `json:"lit_access_conditions"`
}

//changing case causes _ in Golang table name calls....thats why its all lower case after first char
type Groupchatitem struct {
	Id            int       `gorm:"primary_key"`
	Fromaddr      string    `json:"fromaddr"`
	Timestamp     string    `json:"timestamp"`
	Timestamp_dtm time.Time `json:"timestamp_dtm"`
	Message       string    `json:"message"`
	Nftaddr       string    `json:"nftaddr"`
	Type          string    `json:"type"`
	Contexttype   string    `json:"context_type"`
	Name          string    `json:"sender_name"`
}

//secondary table to help only load new messages for each user (not reload whole chat history)
type Groupchatreadtime struct {
	Id                int       `gorm:"primaryKey;autoIncrement"`
	Fromaddr          string    `json:"fromaddr"`
	Readtimestamp_dtm time.Time `json:"readtimestamp_dtm"`
	Nftaddr           string    `json:"nftaddr"`
}

type Blockeduser struct {
	Id             int    `gorm:"primary_key"`                        //AUTO-GENERATED (PRIMARY KEY)
	Owneraddress   string `json:"owneraddress" binding:"required"`    //*** REQUIRED INPUT ***
	Blockedaddress string `json:"blockedaddress" validate:"required"` //*** REQUIRED INPUT ***
}

type Moderator struct {
	Id      int    `gorm:"primary_key"`                 //AUTO-GENERATED (PRIMARY KEY)
	Address string `json:"address" binding:"required"`  //*** REQUIRED INPUT ***
	Company string `json:"company" validate:"required"` //*** REQUIRED INPUT ***
}

//potentially use this to keep track of user logins for DAU metrics
type Logintime struct {
	Id        int       `gorm:"primaryKey;autoIncrement"`
	Address   string    `json:"address"`
	Timestamp time.Time `json:"timestamp"`
}

type Addrnamesignupitem struct {
	Id         int    `gorm:"primaryKey;autoIncrement"`
	Address    string `json:"address"`    //ADDRESS
	Name       string `json:"name"`       //NAME
	Signupsite string `json:"signupsite"` //SITE USER SIGNED UP FROM
	Domain     string `json:"domain"`     //DOMAIN
	Email      string `json:"email"`      //ADMIN ONLY USAGE
}

type Addrnameitem struct {
	Id      int    `gorm:"primaryKey;autoIncrement"`
	Address string `json:"address"`
	Name    string `json:"name"`
}

type Imageitem struct {
	Id         int    `gorm:"primaryKey;autoIncrement"`
	Base64data string `json:"base64data"`
	Addr       string `json:"addr"`
}

type ImageitemPhoto struct {
	Id         int    `gorm:"primaryKey;autoIncrement"`
	Base64data string `json:"base64data"`
	Addr       string `json:"addr"`
	Imageid    string `json:"imageid"`
}

//we have to keep track if a user has manually unjoined, if they did don't auto-join them again
type Userunjoined struct {
	Id         int    `gorm:"primaryKey;autoIncrement"`
	Walletaddr string `json:"walletaddr"`
	Nftaddr    string `json:"nftaddr"`
	Unjoined   bool   `json:"unjoined"`
}

type Bookmarkitem struct {
	Id         int    `gorm:"primaryKey;autoIncrement"`
	Walletaddr string `json:"walletaddr"`
	Nftaddr    string `json:"nftaddr"`
	Chain      string `json:"chain"`
}

type Communitysocial struct {
	Id        int    `gorm:"primaryKey;autoIncrement"`
	Community string `json:"community"`
	Type      string `json:"type"`
	Name      string `json:"name"`
}

type CommunitySocialStruct struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

//cant use camel notation because gorm adds "_" for any subsequent capital letter
type Communityaccesscondition struct {
	Id      int    `gorm:"primaryKey;autoIncrement"`
	Slug    string `json:"slug"` //change community to slug in DB table too
	Nftaddr string `json:"address"`
	Count   string `json:"count"`
}
type Createcommunityitem struct {
	Id     int                     `gorm:"primaryKey;autoIncrement"`
	Name   string                  `json:"name"`
	Slug   string                  `json:"slug"`
	Image  string                  `json:"image"` //base64
	Social []CommunitySocialStruct `json:"social"`
}

type Communityadmin struct {
	Id          int    `gorm:"primaryKey;autoIncrement"`
	Slug        string `json:"slug"`
	Adminaddr   string `json:"adminaddr"`   //wallet addr
	Accesslevel string `json:"accesslevel"` //admin, moderator, ?
}

type BookmarkReturnItem struct {
	Id                int       `gorm:"primaryKey;autoIncrement"`
	Walletaddr        string    `json:"walletaddr"`
	Nftaddr           string    `json:"nftaddr"`
	Lastmsg           string    `json:"lastmsg"`
	Lasttimestamp     string    `json:"lasttimestamp"`
	Lasttimestamp_dtm time.Time `json:"lasttimestamp_dtm"`
	Unreadcnt         int       `json:"unreadcnt"`
}

type Nftsidebar struct {
	Id       int    `gorm:"primaryKey;autoIncrement"`
	Fromaddr string `json:"fromaddr"`
	Unread   int    `json:"unread"`
	Nftaddr  string `json:"nftaddr"`
	Nftid    string `json:"nftid"`
}

// Chatiteminbox entity info
// @Description Used as Return Data Struct Only
type Chatiteminbox struct {
	Id            int       `gorm:"primaryKey;autoIncrement"`
	Fromaddr      string    `json:"fromaddr"`
	Toaddr        string    `json:"toaddr"`
	Timestamp     string    `json:"timestamp"`
	Timestamp_dtm time.Time `json:"timestamp_dtm"`
	Msgread       bool      `json:"read"`
	Message       string    `json:"message"`
	Nftaddr       string    `json:"nftaddr"`
	Nftid         string    `json:"nftid"`
	Unreadcnt     int       `json:"unread"`
	Type          string    `json:"type"`
	Contexttype   string    `json:"context_type"`
	Sendername    string    `json:"sender_name"`
	Name          string    `json:"name"`
	LogoData      string    `json:"logo"`
	Chain         string    `json:"chain"`
	Encryptsymkey string    `json:"encrypted_sym_lit_key"` //USE IF USING LIT ENCRYPTION
	Litaccesscond string    `json:"lit_access_conditions"`
}

type Chatiteminboxconvos struct {
	Address string `json:"address"`
}
