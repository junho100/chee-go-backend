package handler

import (
	"chee-go-backend/internal/common"
	"chee-go-backend/internal/domain/service"
	"chee-go-backend/internal/http/dto"
	"chee-go-backend/internal/infrastructure/telegram"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
	telegramClient      telegram.TelegramClient
	userService         service.UserService
	notificationService service.NotificationService
}

func NewNotificationHandler(router *gin.Engine, telegramClient telegram.TelegramClient, userService service.UserService, notificationService service.NotificationService) {
	handler := &NotificationHandler{
		telegramClient:      telegramClient,
		userService:         userService,
		notificationService: notificationService,
	}

	router.POST("/api/notifications/validate-token", handler.ValidateToken)
	router.POST("/api/notifications/validate-chat-id", handler.ValidateChatID)
	router.POST("/api/notifications/config", handler.CreateNotificationConfig)
	router.GET("/api/notifications/config", handler.GetNotificationConfig)
	router.GET("/api/notifications/:id", handler.GetNotificationByID)
}

func (h *NotificationHandler) ValidateToken(c *gin.Context) {
	var validateTokenRequest dto.ValidateTokenRequest

	if err := c.BindJSON(&validateTokenRequest); err != nil {
		response := &common.CommonErrorResponse{
			Message: "bad content.",
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	validateTokenResponse := dto.ValidateTokenResponse{
		IsValid: h.telegramClient.ValidateToken(validateTokenRequest.Token),
	}
	c.JSON(http.StatusCreated, validateTokenResponse)
}

func (h *NotificationHandler) ValidateChatID(c *gin.Context) {
	var validateChatIDRequest dto.ValidateChatIDRequest

	if err := c.BindJSON(&validateChatIDRequest); err != nil {
		response := &common.CommonErrorResponse{
			Message: "bad content.",
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	validateChatIDResponse := dto.ValidateChatIDResponse{
		IsValid: h.telegramClient.ValidateChatID(validateChatIDRequest.Token, validateChatIDRequest.ChatID),
	}
	c.JSON(http.StatusCreated, validateChatIDResponse)
}

func (h *NotificationHandler) CreateNotificationConfig(c *gin.Context) {
	var createNotificationConfigRequest dto.CreateNotificationConfigRequest
	var token string
	var err error
	var userID string
	var configID uint

	if err := c.BindJSON(&createNotificationConfigRequest); err != nil {
		response := &common.CommonErrorResponse{
			Message: "bad content.",
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if token, err = h.userService.ExtractToken(c.GetHeader("Authorization")); err != nil {
		response := &common.CommonErrorResponse{
			Message: "failed to authorization.",
		}
		c.JSON(http.StatusForbidden, response)
		return
	}

	if userID, err = h.userService.GetUserIDFromToken(token); err != nil {
		response := &common.CommonErrorResponse{
			Message: "failed to authorization.",
		}
		c.JSON(http.StatusForbidden, response)
		return
	}

	createNotificationDto := dto.CreateNotificationConfigDto{
		UserID:   userID,
		Token:    createNotificationConfigRequest.Token,
		ChatID:   createNotificationConfigRequest.ChatID,
		Keywords: createNotificationConfigRequest.Keywords,
	}
	if configID, err = h.notificationService.CreateNotificationConfig(createNotificationDto); err != nil {
		response := &common.CommonErrorResponse{
			Message: "failed to create config.",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	createNotificationConfigResponse := dto.CreateNotificationConfigResponse{
		ConfigID: configID,
	}
	c.JSON(http.StatusCreated, createNotificationConfigResponse)
}

func (h *NotificationHandler) GetNotificationConfig(c *gin.Context) {
	var token string
	var err error
	var userID string

	if token, err = h.userService.ExtractToken(c.GetHeader("Authorization")); err != nil {
		response := &common.CommonErrorResponse{
			Message: "failed to authorization.",
		}
		c.JSON(http.StatusForbidden, response)
		return
	}

	if userID, err = h.userService.GetUserIDFromToken(token); err != nil {
		response := &common.CommonErrorResponse{
			Message: "failed to authorization.",
		}
		c.JSON(http.StatusForbidden, response)
		return
	}

	notificationConfig, err := h.notificationService.GetNotificationConfigByUserID(userID)
	if err != nil {
		response := &common.CommonErrorResponse{
			Message: "no config.",
		}
		c.JSON(http.StatusNotFound, response)
		return
	}
	keywords := h.notificationService.GetKeywordsByNotificationID(notificationConfig.ID)

	getNotificationConfigResponse := dto.GetNotificationConfigResponse{
		Token:    notificationConfig.TelegramToken,
		ChatID:   notificationConfig.TelegramChatID,
		Keywords: keywords,
	}
	c.JSON(http.StatusOK, getNotificationConfigResponse)
}

func (h *NotificationHandler) GetNotificationByID(c *gin.Context) {
	notificationID := c.Param("id")

	if notificationID == "" {
		response := &common.CommonErrorResponse{
			Message: "bad path variable.",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	notification, err := h.notificationService.GetNotificationByID(notificationID)
	if err != nil {
		response := &common.CommonErrorResponse{
			Message: "no notification.",
		}
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := dto.GetNotificationByIDResponse(*notification)
	c.JSON(http.StatusOK, response)
}
