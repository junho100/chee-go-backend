package handler

import (
	"chee-go-backend/internal/common"
	"chee-go-backend/internal/domain/service"
	"chee-go-backend/internal/http/dto"
	"net/http"

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
	var signUpRequest dto.SignUpRequest

	if err := c.BindJSON(&signUpRequest); err != nil {
		response := &common.CommonErrorResponse{
			Message: "bad content.",
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	createUserDto := &dto.CreateUserDto{
		ID:       signUpRequest.ID,
		Email:    signUpRequest.Email,
		Password: signUpRequest.Password,
	}

	if err := h.userService.CreateUser(createUserDto); err != nil {
		response := &common.CommonErrorResponse{
			Message: "failed to create user.",
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := &dto.SignUpResponse{
		ID: createUserDto.ID,
	}

	c.JSON(http.StatusCreated, response)
}

func (h *UserHandler) CheckIdExists(c *gin.Context) {
	id := c.Query("id")

	response := &dto.CheckIDResponse{
		IsExists: h.userService.CheckUserByID(id),
	}

	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) Login(c *gin.Context) {
	var loginRequest dto.LoginRequest

	if err := c.BindJSON(&loginRequest); err != nil {
		response := &common.CommonErrorResponse{
			Message: "bad content.",
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user, err := h.userService.GetUserByID(loginRequest.ID)

	if err != nil {
		response := &common.CommonErrorResponse{
			Message: "login failed.",
		}
		c.JSON(http.StatusForbidden, response)
		return
	}

	if h.userService.CheckPassword(loginRequest.Password, user.HashedPassword) != nil {
		response := &common.CommonErrorResponse{
			Message: "login failed.",
		}
		c.JSON(http.StatusForbidden, response)
		return
	}

	token, err := h.userService.CreateToken(loginRequest.ID)
	if err != nil {
		response := &common.CommonErrorResponse{
			Message: "login failed.",
		}
		c.JSON(http.StatusForbidden, response)
		return
	}

	response := &dto.LoginResponse{
		Token: token,
	}
	c.JSON(http.StatusCreated, response)
}

func (h *UserHandler) CheckMe(c *gin.Context) {
	var checkMeRequest dto.CheckMeRequest

	if err := c.BindJSON(&checkMeRequest); err != nil {
		response := &common.CommonErrorResponse{
			Message: "bad content.",
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var userID string
	var err error
	if userID, err = h.userService.GetUserIDFromToken(checkMeRequest.Token); err != nil {
		response := &common.CommonErrorResponse{
			Message: "invalid token.",
		}
		c.JSON(http.StatusForbidden, response)
		return
	}

	response := &dto.CheckMeResponse{
		UserID: userID,
	}
	c.JSON(http.StatusCreated, response)
}
