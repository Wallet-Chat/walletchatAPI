package entity

import "time"

type Referralcode struct {
	Id         int       `gorm:"primary_key"` //AUTO-GENERATED (PRIMARY KEY)
	Walletaddr string    `json:"walletaddr"`  //
	Code       string    `json:"code"`        //
	Date       time.Time `json:"date"`        //
	Redeemed   bool      `json:"redeemed"`
}
