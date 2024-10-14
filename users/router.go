package users

import (
	"net/http"

	"chee-go-backend/common"

	"github.com/gin-gonic/gin"
)

func RegisterUsersRouters(router *gin.RouterGroup) {
	router.POST("", SignUp)
	router.GET("/check-id", CheckID)
	router.POST("/login", Login)
	router.POST("/me", CheckMe)
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

func Login(c *gin.Context) {
	var loginRequest LoginRequest

	if err := c.BindJSON(&loginRequest); err != nil {
		response := &common.CommonErrorResponse{
			Message: "bad content.",
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user, err := GetUserByID(loginRequest.ID)

	if err != nil {
		response := &common.CommonErrorResponse{
			Message: "login failed.",
		}
		c.JSON(http.StatusForbidden, response)
		return
	}

	if CheckPassword(loginRequest.Password, user.HashedPassword) != nil {
		response := &common.CommonErrorResponse{
			Message: "login failed.",
		}
		c.JSON(http.StatusForbidden, response)
		return
	}

	token, err := CreateToken(loginRequest.ID)
	if err != nil {
		response := &common.CommonErrorResponse{
			Message: "login failed.",
		}
		c.JSON(http.StatusForbidden, response)
		return
	}

	response := &LoginResponse{
		Token: token,
	}
	c.JSON(http.StatusCreated, response)
}

func CheckMe(c *gin.Context) {
	var checkMeRequest CheckMeRequest

	if err := c.BindJSON(&checkMeRequest); err != nil {
		response := &common.CommonErrorResponse{
			Message: "bad content.",
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var userID string
	var err error
	if userID, err = GetUserIDFromToken(checkMeRequest.Token); err != nil {
		response := &common.CommonErrorResponse{
			Message: "invalid token.",
		}
		c.JSON(http.StatusForbidden, response)
		return
	}

	response := &CheckMeResponse{
		UserID: userID,
	}
	c.JSON(http.StatusCreated, response)
}
