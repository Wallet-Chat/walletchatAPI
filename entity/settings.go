package entity

//Settings object for REST(CRUD)
type Settings struct {
	ID         int    `json:"id"`
	Walletaddr string `json:"walletaddr"`
	//Publickey  string `json:"publickey"`
	Email    string `json:"email"`
	Notifydm bool   `json:"notifydm"`
	Notify24 bool   `json:"notify24"`
}
