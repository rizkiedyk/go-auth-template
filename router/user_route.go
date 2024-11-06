package router

import (
	"go-auth/handler"
	"go-auth/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.RouterGroup, userHandler *handler.UserHandler) {
	userRoutes := r.Group("/user", middleware.AuthMiddleware())
	{
		userRoutes.GET("/", middleware.AccessControlMiddleware(), userHandler.GetUsers)
		userRoutes.GET("", middleware.AccessControlMiddleware(), userHandler.GetUserByID)
		userRoutes.POST("/set-role", middleware.AccessControlMiddleware(), userHandler.SetRole)
	}
}