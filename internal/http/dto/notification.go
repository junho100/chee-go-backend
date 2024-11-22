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
