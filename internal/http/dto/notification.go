package dto

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
