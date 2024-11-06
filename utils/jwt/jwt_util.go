package jwt

import (
	"fmt"
	"go-auth/domain/model"
	"go-auth/utils/helper"

	"time"

	"github.com/golang-jwt/jwt"
)

var jwtSecret = []byte(helper.GetENV("JWT_SECRET", "secret"))

func GenerateToken(user model.User) (string, error) {
	claims := jwt.MapClaims{
		"username": user.Username,
        "role":     user.Role,
        "exp":      time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseToken(tokenStr string) (model.User, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	fmt.Println(token)
	if err != nil || !token.Valid {
		return model.User{}, err
	}

	claims := token.Claims.(*jwt.MapClaims)

	fmt.Println(claims)

	return model.User{
		Username: (*claims)["username"].(string),
		Role:     (*claims)["role"].(string),
	}, nil
}