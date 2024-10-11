package main

import (
	"chee-go-backend/common"
	"chee-go-backend/users"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	DB := common.Init()
	DB.AutoMigrate(&users.User{})

	r := gin.Default()

	r.SetTrustedProxies(nil)

	serverRoute := r.Group("/api")
	users.RegisterUsersRouters(serverRoute.Group("/users"))

	r.Run(":8080")
}
