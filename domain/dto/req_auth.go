package dto

type ReqRegister struct {
		Username string `json:"username" validate:"required,min=3,max=20"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
}

type ReqLogin struct {
	Username string `json:"username" validate:"required,username"`
	Password string `json:"password" validate:"required"`
}