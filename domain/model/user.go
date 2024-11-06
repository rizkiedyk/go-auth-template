package model

type User struct {
	Id string `json:"id" bson:"id"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Role string `json:"role" bson:"role"`
	Email string `json:"email" bson:"email"`
}

const (
    RoleAdmin = "admin"
    RoleUser  = "user"
)