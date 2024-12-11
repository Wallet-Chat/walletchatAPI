package entity

type Ourausertemp struct {
	Id           int    `gorm:"primaryKey;autoIncrement"`
	Wallet       string `json:"wallet"`
	Pac          string `json:"pac"`
	Signature    string `json:"signature"`
	Referralcode string `json:"referralcode"`
	Nickname     string `json:"nickname"`
}

type Ourauser struct {
	Id           int    `gorm:"primaryKey;autoIncrement"`
	Wallet       string `json:"wallet"`
	Oauth        string `json:"oauth"`
	Pac          string `json:"pac"`
	Encryptedpac string `json:"encrypted_pac"` //for use in secrets for /runProof
	Signature    string `json:"signature"`
}

type Ouradata struct {
	Id       int    `gorm:"primaryKey;autoIncrement"`
	Wallet   string `json:"wallet"`
	Endpoint string `json:"oauth"`
	Jsondata string `json:"pac"`
}
