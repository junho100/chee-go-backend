package telegram

import "net/http"

type TelegramClient interface {
	ValidateToken(token string) bool
}

type telegramClient struct {
	ApiUrl string
}

func NewTelegramClient() TelegramClient {
	return &telegramClient{
		ApiUrl: "https://api.telegram.org/bot",
	}
}

func (c *telegramClient) ValidateToken(token string) bool {
	resp, err := http.Get(c.ApiUrl + token + "/getMe")
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}
