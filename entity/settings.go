package entity

//Settings object for REST(CRUD)
type Settings struct {
	ID              int    `json:"id"`                             //AUTO-GENERATED (PRIMARY KEY)
	Walletaddr      string `json:"walletaddr" validate:"required"` //*** REQUIRED INPUT ***
	Telegramhandle  string `json:"telegramhandle"`                 //TELEGRAM @ handle - MAINLY USED AS DOUBLE CHECK DURING INITIAL SETUP
	Telegramid      string `json:"telegramid"`                     //TELEGRAM CHAT ID - REQUIRES MSG SENT TO WALLETCHAT BOT TO VERIFY
	Telegramcode    string `json:"telegramcode"`                   //TELEGRAM VERIFICATION CODE - REQUIRES THIS MSG SENT TO WALLETCHAT BOT
	Email           string `json:"email"`                          //EMAIL ADDRESS TO GET NOTIFICATIONS
	Verified        string `json:"verified"`                       //USER CONFIRMED EMAIL OR NOT (string value true/false)
	Notifydm        string `json:"notifydm"`                       //RECEIVE DAILY NOTIFICATION SUMMARY EMAIL (string value true/false)
	Notify24        string `json:"notify24"`                       //RECEIVE NOTIFICATION FOR EVERY DM RECEIVED (string value true/false)
	Signupsite      string `json:"signupsite"`                     //LATEST SITE WHERE NOTIFICATIONS EMAIL WAS ENTERED
	Domain          string `json:"domain"`                         //DOMAIN
	Installedsnap   string `json:"installedsnap"`                  //IS METAMASK SNAP INSTALLED
	Twitteruser     string `json:"twitteruser"`                    //TWITTER/X @user
	Twitterverified string `json:"twitterverified"`                //HAS USER VERIFIED @user with WALLETCHAT (not twitter blue checkmark)
	Twitterid       string `json:"twitterid"`                      //TWITTER USER ID - FUTURE USE IF USER CHANGES NAME?
}
