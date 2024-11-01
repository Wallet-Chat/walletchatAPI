package entity

type Oura struct {
	ID     int
	Wallet string `json:"wallet"`
	Oauth  string `json:"oauth"`
	Pac    string `json:"pac"`
}
