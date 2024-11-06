package model

type User struct {
	Id string `json:"id" bson:"id"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Role string `json:"role" bson:"role"`
	Email string `json:"email" bson:"email"`
	CreatedAt int `json:"created_at" bson:"created_at"`
	UpdatedAt int `json:"updated_at" bson:"updated_at"`
}

const (
    RoleAdmin = "admin"
    RoleUser  = "user"
)