package router

import (
	"go-auth/handler"
	"go-auth/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.RouterGroup, userHandler *handler.UserHandler) {
	userRoutes := r.Group("/user", middleware.AuthMiddleware())
	{
		userRoutes.POST("/set-role", middleware.AdminOnlyMiddleware(), userHandler.SetRole)
	}
}