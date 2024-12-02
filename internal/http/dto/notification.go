package dto

import "time"

type ValidateTokenRequest struct {
	Token string `json:"token"`
}

type ValidateTokenResponse struct {
	IsValid bool `json:"is_valid"`
}

type ValidateChatIDRequest struct {
	Token  string `json:"token"`
	ChatID string `json:"chat_id"`
}

type ValidateChatIDResponse struct {
	IsValid bool `json:"is_valid"`
}

type CreateNotificationConfigRequest struct {
	Token    string   `json:"token"`
	ChatID   string   `json:"chat_id"`
	Keywords []string `json:"keywords"`
}

type CreateNotificationConfigDto struct {
	UserID   string
	Token    string
	ChatID   string
	Keywords []string
}

type CreateNotificationConfigResponse struct {
	ConfigID uint `json:"config_id"`
}

type GetNotificationConfigResponse struct {
	Token           string   `json:"token"`
	ChatID          string   `json:"chat_id"`
	Keywords        []string `json:"keywords"`
	DiscordClientID string   `json:"discord_client_id"`
}

type SendNotificationMessageDto struct {
	Title  string
	Date   time.Time
	Url    string
	Token  string
	ChatID string
}

type FetchSchoolNoticeDto struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	URL      string    `json:"url"`
	PostDate time.Time `json:"post_date"`
	Content  string    `json:"content"`
}

type GetNotificationByIDResponse struct {
	ID      string    `json:"id"`
	Title   string    `json:"title"`
	Date    time.Time `json:"date"`
	Content string    `json:"content"`
	Url     string    `json:"url"`
}

type ValidateDiscordClientIDRequest struct {
	ClientID string `json:"client_id"`
}

type ValidateDiscordClientIDResponse struct {
	IsValid bool `json:"is_valid"`
}
