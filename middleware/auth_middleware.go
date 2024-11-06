package middleware

import (
	"go-auth/domain/dto"
	"go-auth/domain/model"
	"go-auth/utils/jwt"
	"net/http"
	"strings"

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

		token = strings.TrimPrefix(token, "Bearer ")

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

func AdminOnlyMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        user, exists := c.Get("user")
        if !exists {
			c.JSON(http.StatusUnauthorized, dto.Resp{
				Code:    http.StatusUnauthorized,
				Message: "User not authenticated",
				Data:    nil,
			})
            c.Abort()
            return
        }

		var u *model.User
        switch v := user.(type) {
        case *model.User:
            u = v
        case model.User:
            u = &v
        default:
			c.JSON(http.StatusUnauthorized, dto.Resp{
				Code:    http.StatusUnauthorized,
				Message: "Invalid user type",
				Data:    nil,
			})
            c.Abort()
            return
        }

        if u.Role != "super_admin" {
			c.JSON(http.StatusForbidden, dto.Resp{
				Code:    http.StatusForbidden,
				Message: "Super Admin only",
				Data:    nil,
			})
            c.Abort()
            return
        }

        c.Set("role", u.Role)

        c.Next()
    }
}

func AccessControlMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        user, exists := c.Get("user")
        if !exists {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
            c.Abort()
            return
        }

        currentUser := user.(model.User)

        c.Set("role", currentUser.Role)

        c.Next()
    }
}