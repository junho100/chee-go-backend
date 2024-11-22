package handler

import (
	"chee-go-backend/internal/common"
	"chee-go-backend/internal/http/dto"
	"chee-go-backend/internal/infrastructure/telegram"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
	telegramClient telegram.TelegramClient
}

func NewNotificationHandler(router *gin.Engine, telegramClient telegram.TelegramClient) {
	handler := &NotificationHandler{
		telegramClient: telegramClient,
	}

	router.POST("/api/notifications/validate-token", handler.ValidateToken)
	router.POST("/api/notifications/validate-chat-id", handler.ValidateChatID)
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
