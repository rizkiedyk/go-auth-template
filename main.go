package main

import (
	"fmt"
	"go-auth/config"
	"go-auth/router"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	config.ConfigureLogger()
}

func main() {
	router := router.SetupRouter()

	router.Run(":1010")
}