package entity

type Ourauser struct {
	ID     int
	Wallet string `json:"wallet"`
	Oauth  string `json:"oauth"`
	Pac    string `json:"pac"`
}
