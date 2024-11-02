package handler

import (
	"chee-go-backend/internal/domain/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(router *gin.Engine, userService service.UserService) {
	handler := &UserHandler{
		userService: userService,
	}

	router.POST("/api/users", handler.SignUp)
	router.GET("/api/users/check-id", handler.CheckIdExists)
	router.POST("/api/users/login", handler.Login)
	router.POST("/api/users/me", handler.CheckMe)
}

func (h *UserHandler) SignUp(c *gin.Context) {

}

func (h *UserHandler) CheckIdExists(c *gin.Context) {

}

func (h *UserHandler) Login(c *gin.Context) {

}

func (h *UserHandler) CheckMe(c *gin.Context) {

}
