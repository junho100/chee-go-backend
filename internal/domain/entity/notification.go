package entity

import "time"

type NotificationConfig struct {
	ID             uint `gorm:"primary_key"`
	UserID         string
	User           User
	TelegramToken  string `gorm:"telegram_token"`
	TelegramChatID string `gorm:"telegram_chat_id"`
}

type NotificationKeyword struct {
	ID   uint `gorm:"primary_key"`
	Name string
}

type NotificationConfigKeyword struct {
	NotificationConfig    NotificationConfig
	NotificationConfigID  uint
	NotificationKeyword   NotificationKeyword
	NotificationKeywordID uint
}

type SchoolNotification struct {
	ID      uint `gorm:"primary_key"`
	Title   string
	Date    time.Time `gorm:"column:date;type:date"`
	Content string    `gorm:"column:date;type:text"`
	Url     string
}