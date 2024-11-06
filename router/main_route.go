package router

import (
	"go-auth/config"
	"go-auth/handler"
	"go-auth/repository"
	"go-auth/service"

	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
)

var logger = logging.MustGetLogger("main")

func SetupRouter() *gin.Engine {
	route := gin.Default()

	apiV1 := route.Group("/api/v1")

	db:= config.ConnectDatabase()

	// index repo
	indexRepo := repository.NewIndexRepo(db)

	// auth route
	authRepo := repository.NewAuthRepository(db, indexRepo)
	authService := service.NewAuthService(authRepo)
	authHandler := handler.NewAuthHandler(authService)
	AuthRoute(apiV1, authHandler)

	return route
}