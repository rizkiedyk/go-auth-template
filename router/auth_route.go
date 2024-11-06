package router

import (
	"go-auth/handler"

	"github.com/gin-gonic/gin"
)

func AuthRoute(r *gin.RouterGroup, authHandler *handler.AuthHandler) {
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", authHandler.Register)
		authRoutes.POST("/login", authHandler.Login)
	}
}