package entity

import "time"

type Referralcode struct {
	Id         int       `gorm:"primary_key"` //AUTO-GENERATED (PRIMARY KEY)
	Walletaddr string    `json:"walletaddr"`  //
	Code       string    `json:"code"`        //
	Date       time.Time `json:"date"`        //
	Redeemed   bool      `json:"redeemed"`
	Twitterid  string    `json:"twitterid"`
}

//used to track if the user is valid or not (to skip prompting for a new code upon login - similar to checking for name being set)
type Referraluser struct {
	Id           int    `gorm:"primary_key"`  //AUTO-GENERATED (PRIMARY KEY)
	Walletaddr   string `json:"walletaddr"`   //
	Referralcode string `json:"referralcode"` //
}
