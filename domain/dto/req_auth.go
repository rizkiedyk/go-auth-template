package dto

type ReqRegister struct {
		Username string `json:"username" validate:"required,min=3,max=20"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=6"`
}

type ReqLogin struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SetRoleRequest struct {
    UserID string   `json:"user_id" validate:"required"`
    Role   string `json:"role" validate:"required,oneof=admin user guest"` // Valid roles
}