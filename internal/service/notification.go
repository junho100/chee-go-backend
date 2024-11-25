package service

import (
	"chee-go-backend/internal/domain/entity"
	"chee-go-backend/internal/domain/repository"
	"chee-go-backend/internal/domain/service"
	"chee-go-backend/internal/http/dto"
	"chee-go-backend/internal/infrastructure/crawler"
)

type notificationService struct {
	notificationRepository repository.NotificationRepository
	crawler                crawler.Crawler
}

// FindAllNotificationConfigs implements service.NotificationService.
func (s *notificationService) FindAllNotificationConfigs(configs *[]entity.NotificationConfig) error {
	if err := s.notificationRepository.FindAllNotificationConfigs(configs); err != nil {
		return err
	}

	return nil
}

func (s *notificationService) FindKeywordsByConfigID(configID uint) []string {
	notificationKeywords := s.notificationRepository.FindKeywordsByNotificationID(configID)
	keywords := make([]string, len(notificationKeywords))

	for i, v := range notificationKeywords {
		keywords[i] = v.Name
	}

	return keywords
}

func NewNotificationService(notificationRepository repository.NotificationRepository, crawler crawler.Crawler) service.NotificationService {
	return &notificationService{
		notificationRepository: notificationRepository,
		crawler:                crawler,
	}
}

// CreateNotificationConfig implements service.NotificationService.
func (s *notificationService) CreateNotificationConfig(createNotificationDto dto.CreateNotificationConfigDto) (uint, error) {
	tx, err := s.notificationRepository.StartTransaction()
	if err != nil {
		return 0, err
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	notificationConfig := entity.NotificationConfig{}
	if err := s.notificationRepository.FindNotificationConfigByUserID(&notificationConfig, createNotificationDto.UserID); err != nil {
		notificationConfig.UserID = createNotificationDto.UserID
		notificationConfig.TelegramToken = createNotificationDto.Token
		notificationConfig.TelegramChatID = createNotificationDto.ChatID

		if err := s.notificationRepository.CreateNotificationConfig(tx, &notificationConfig); err != nil {
			tx.Rollback()
			return 0, err
		}
	} else {
		notificationConfig.TelegramToken = createNotificationDto.Token
		notificationConfig.TelegramChatID = createNotificationDto.ChatID

		if err := s.notificationRepository.UpdateNotificationConfig(tx, &notificationConfig); err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	if err := s.notificationRepository.CreateKeywordByNotificationConfigID(tx, notificationConfig.ID, createNotificationDto.Keywords); err != nil {
		tx.Rollback()
		return 0, err
	}

	return notificationConfig.ID, tx.Commit().Error
}

func (s *notificationService) GetNotificationConfigByUserID(userID string) (*entity.NotificationConfig, error) {
	notificationConfig := entity.NotificationConfig{}

	if err := s.notificationRepository.FindNotificationConfigByUserID(&notificationConfig, userID); err != nil {
		return nil, err
	}

	return &notificationConfig, nil
}

func (s *notificationService) GetKeywordsByNotificationID(notificationConfigID uint) []string {
	notificationKeywords := s.notificationRepository.FindKeywordsByNotificationID(notificationConfigID)

	keywords := make([]string, len(notificationKeywords))
	for i, v := range notificationKeywords {
		keywords[i] = v.Name
	}

	return keywords
}

func (s *notificationService) SaveTodayNotifications(notifications []entity.SchoolNotification) error {
	return s.notificationRepository.SaveNotifications(notifications)
}
