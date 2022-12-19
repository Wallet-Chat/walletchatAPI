package entity

//Settings object for REST(CRUD)
type Settings struct {
	ID         int    `json:"id"`                             //AUTO-GENERATED (PRIMARY KEY)
	Walletaddr string `json:"walletaddr" validate:"required"` //*** REQUIRED INPUT ***
	//Publickey  string `json:"publickey"`
	Email    string `json:"email"`    //EMAIL ADDRESS TO GET NOTIFICATIONS
	Verified string `json:"verified"` //USER CONFIRMED EMAIL OR NOT (string value true/false)
	Notifydm string `json:"notifydm"` //RECEIVE DAILY NOTIFICATION SUMMARY EMAIL (string value true/false)
	Notify24 string `json:"notify24"` //RECEIVE NOTIFICATION FOR EVERY DM RECEIVED (string value true/false)
	Signupsite string `json:"signupsite"` //LATEST SITE WHERE NOTIFICATIONS EMAIL WAS ENTERED
	Domain     string `json:"domain"`     //DOMAIN
}
