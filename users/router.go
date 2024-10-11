package users

import (
	"net/http"

	"chee-go-backend/common"

	"github.com/gin-gonic/gin"
)

func RegisterUsersRouters(router *gin.RouterGroup) {
	router.POST("/", SignUp)
	router.GET("/check-id", CheckID)
}

func SignUp(c *gin.Context) {
	var signUpRequest SignUpRequest

	if err := c.BindJSON(&signUpRequest); err != nil {
		response := &common.CommonErrorResponse{
			Message: "bad content.",
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	dto := &CreateUserDto{
		ID:       signUpRequest.ID,
		Email:    signUpRequest.Email,
		Password: signUpRequest.Password,
	}

	if err := CreateUser(dto); err != nil {
		response := &common.CommonErrorResponse{
			Message: "failed to create user.",
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := &SignUpResponse{
		ID: dto.ID,
	}
	c.JSON(http.StatusCreated, response)
}

func CheckID(c *gin.Context) {
	id := c.Query("id")

	response := &CheckIDResponse{
		IsExists: CheckUserByID(id),
	}
	c.JSON(http.StatusOK, response)
}
