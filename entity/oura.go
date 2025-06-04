package entity

import "time"

type Ourausertemp struct {
	Id           int    `gorm:"primaryKey;autoIncrement"`
	Wallet       string `json:"wallet"`
	Pac          string `json:"pac"`
	Signature    string `json:"signature"`
	Referralcode string `json:"referralcode"`
	Nickname     string `json:"nickname"`
}

type Ourauser struct {
	Id           int       `gorm:"primaryKey;autoIncrement"`
	Wallet       string    `json:"wallet"`
	Uuid         string    `json:"uuid"`
	Pac          string    `json:"pac"`
	Encryptedpac string    `json:"encrypted_pac"` //for use in secrets for /runProof when using Oura API directly
	Signature    string    `json:"signature"`
	Deviceid     string    `json:"deviceid"` //used for mobile device ID for notifications later on
	Lastcheckin  time.Time `json:"lastcheckin"`
	Numcheckins  int       `json:"numcheckins"`
}

type Ouradata struct {
	Id       int    `gorm:"primaryKey;autoIncrement"`
	Wallet   string `json:"wallet"`
	Endpoint string `json:"oauth"`
	Jsondata string `json:"pac"`
}
