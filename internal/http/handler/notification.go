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
