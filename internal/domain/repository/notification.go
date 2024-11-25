package repository

import (
	"chee-go-backend/internal/domain/entity"

	"gorm.io/gorm"
)

type NotificationRepository interface {
	StartTransaction() (*gorm.DB, error)
	FindNotificationConfigByUserID(notificationConfig *entity.NotificationConfig, userID string) error
	CreateNotificationConfig(tx *gorm.DB, notificationConfig *entity.NotificationConfig) error
	UpdateNotificationConfig(tx *gorm.DB, notificationConfig *entity.NotificationConfig) error
	CreateKeywordByNotificationConfigID(tx *gorm.DB, notificationConfigID uint, keywords []string) error
	FindKeywordsByNotificationID(notificationConfigID uint) []entity.NotificationKeyword
	FindAllNotificationConfigs(configs *[]entity.NotificationConfig) error
	SaveNotifications(notifications []entity.SchoolNotification) error
	FindNotificationByID(notification *entity.SchoolNotification, id string) error
}
