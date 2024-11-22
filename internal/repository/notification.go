package repository

import (
	"chee-go-backend/internal/domain/entity"
	"chee-go-backend/internal/domain/repository"

	"gorm.io/gorm"
)

type notificationRepository struct {
	db *gorm.DB
}

// CreateNotificationConfig implements repository.NotificationRepository.
func (r *notificationRepository) CreateNotificationConfig(tx *gorm.DB, notificationConfig *entity.NotificationConfig) error {
	if err := tx.Save(notificationConfig).Error; err != nil {
		return err
	}

	return nil
}

// FindNotificationConfigByUserID implements repository.NotificationRepository.
func (r *notificationRepository) FindNotificationConfigByUserID(notificationConfig *entity.NotificationConfig, userID string) error {
	if err := r.db.Where(&entity.NotificationConfig{
		UserID: userID,
	}).First(notificationConfig).Error; err != nil {
		return err
	}

	return nil
}

// StartTransaction implements repository.NotificationRepository.
func (r *notificationRepository) StartTransaction() (*gorm.DB, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	return tx, nil
}

// UpdateNotificationConfig implements repository.NotificationRepository.
func (r *notificationRepository) UpdateNotificationConfig(tx *gorm.DB, notificationConfig *entity.NotificationConfig) error {
	if err := tx.Save(notificationConfig).Error; err != nil {
		return err
	}

	return nil
}

// CreateKeywordByNotificationConfigID implements repository.NotificationRepository.
func (r *notificationRepository) CreateKeywordByNotificationConfigID(tx *gorm.DB, notificationConfigID uint, keywords []string) error {
	if err := tx.Where(&entity.NotificationConfigKeyword{
		NotificationConfigID: notificationConfigID,
	}).Delete(&entity.NotificationConfigKeyword{}).Error; err != nil {
		return err
	}

	for _, keyword := range keywords {
		var existingKeyword entity.NotificationKeyword
		err := tx.Where("name = ?", keyword).First(&existingKeyword).Error

		var keywordID uint
		if err == gorm.ErrRecordNotFound {
			newKeyword := entity.NotificationKeyword{
				Name: keyword,
			}
			if err := tx.Create(&newKeyword).Error; err != nil {
				return err
			}
			keywordID = newKeyword.ID
		} else if err != nil {
			return err
		} else {
			keywordID = existingKeyword.ID
		}

		configKeyword := entity.NotificationConfigKeyword{
			NotificationConfigID:  notificationConfigID,
			NotificationKeywordID: keywordID,
		}
		if err := tx.Create(&configKeyword).Error; err != nil {
			return err
		}
	}

	return nil
}

func (r *notificationRepository) FindKeywordsByNotificationID(notificationConfigID uint) []entity.NotificationKeyword {
	var notificationConfigKeywords []entity.NotificationConfigKeyword

	if err := r.db.Preload("NotificationKeyword").Where(&entity.NotificationConfigKeyword{
		NotificationConfigID: notificationConfigID,
	}).Find(&notificationConfigKeywords).Error; err != nil {
		return make([]entity.NotificationKeyword, 0)
	}

	keywords := make([]entity.NotificationKeyword, len(notificationConfigKeywords))
	for i, v := range notificationConfigKeywords {
		keywords[i] = v.NotificationKeyword
	}

	return keywords
}

func NewNotificationRepository(db *gorm.DB) repository.NotificationRepository {
	return &notificationRepository{
		db: db,
	}
}
