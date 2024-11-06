package middleware

import (
	"go-auth/domain/dto"
	"go-auth/utils/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, dto.Resp{
				Code:    http.StatusUnauthorized,
				Message: "Unauthorized",
				Data:    nil,
			})
			c.Abort()
			return
		}

		user, err := jwt.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, dto.Resp{
				Code:    http.StatusUnauthorized,
				Message: "Unauthorized",
				Data:    nil,
			})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}