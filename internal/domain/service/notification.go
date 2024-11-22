package service

import (
	"chee-go-backend/internal/domain/entity"
	"chee-go-backend/internal/http/dto"
)

type NotificationService interface {
	CreateNotificationConfig(createNotificationDto dto.CreateNotificationConfigDto) (uint, error)
	GetNotificationConfigByUserID(userID string) (*entity.NotificationConfig, error)
	GetKeywordsByNotificationID(notificationConfigID uint) []string
}
