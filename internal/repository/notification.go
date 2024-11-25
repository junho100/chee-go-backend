package repository

import (
	"chee-go-backend/internal/domain/entity"
	"chee-go-backend/internal/domain/repository"

	"gorm.io/gorm"
)

type notificationRepository struct {
	db *gorm.DB
}

func (r *notificationRepository) FindAllNotificationConfigs(configs *[]entity.NotificationConfig) error {
	if err := r.db.Find(configs).Error; err != nil {
		return err
	}

	return nil
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

func (r *notificationRepository) SaveNotifications(notifications []entity.SchoolNotification) error {
	if len(notifications) == 0 {
		return nil
	}

	// 중복 체크를 위한 기존 ID 조회
	var existingIDs []string
	for _, notice := range notifications {
		var exists bool
		err := r.db.Model(&entity.SchoolNotification{}).
			Select("count(*) > 0").
			Where("id = ?", notice.ID).
			Find(&exists).Error
		if err != nil {
			return err
		}
		if exists {
			existingIDs = append(existingIDs, notice.ID)
		}
	}

	// 중복되지 않은 공지사항만 필터링
	var newNotifications []entity.SchoolNotification
	for _, notice := range notifications {
		if !contains(existingIDs, notice.ID) {
			newNotifications = append(newNotifications, notice)
		}
	}

	if len(newNotifications) == 0 {
		return nil
	}

	return r.db.Create(&newNotifications).Error
}

// 문자열 슬라이스에 특정 값이 포함되어 있는지 확인하는 헬퍼 함수
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func NewNotificationRepository(db *gorm.DB) repository.NotificationRepository {
	return &notificationRepository{
		db: db,
	}
}
