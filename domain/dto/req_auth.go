package dto

type ReqRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type ReqLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}