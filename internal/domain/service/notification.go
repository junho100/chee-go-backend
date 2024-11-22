package service

import "chee-go-backend/internal/http/dto"

type NotificationService interface {
	CreateNotificationConfig(createNotificationDto dto.CreateNotificationConfigDto) (uint, error)
}
