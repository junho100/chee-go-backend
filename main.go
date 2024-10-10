package main

import (
	"chee-go-backend/common"
	"chee-go-backend/users"

	"github.com/gin-gonic/gin"
)

func main() {
	DB := common.Init()
	DB.AutoMigrate(&users.User{})

	r := gin.Default()

	r.SetTrustedProxies(nil)

	serverRoute := r.Group("/api")
	users.RegisterUsersRouters(serverRoute.Group("/users"))

	r.Run(":8080")
}
