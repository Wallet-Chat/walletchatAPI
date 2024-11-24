package entity

type Ourauser struct {
	Id        int    `gorm:"primaryKey;autoIncrement"`
	Wallet    string `json:"wallet"`
	Oauth     string `json:"oauth"`
	Pac       string `json:"pac"`
	Signature string `json:"signature"`
}

type Ouradata struct {
	Id       int    `gorm:"primaryKey;autoIncrement"`
	Wallet   string `json:"wallet"`
	Endpoint string `json:"oauth"`
	Jsondata string `json:"pac"`
}
