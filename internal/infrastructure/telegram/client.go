package telegram

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type TelegramClient interface {
	ValidateToken(token string) bool
	ValidateChatID(token string, chatID string) bool
}

type telegramClient struct {
	ApiUrl string
}

func NewTelegramClient() TelegramClient {
	return &telegramClient{
		ApiUrl: "https://api.telegram.org/bot",
	}
}

type messageResponse struct {
	Ok     bool `json:"ok"`
	Result struct {
		MessageID int `json:"message_id"`
		Chat      struct {
			ID int64 `json:"id"`
		} `json:"chat"`
	} `json:"result"`
}

func (c *telegramClient) ValidateToken(token string) bool {
	resp, err := http.Get(c.ApiUrl + token + "/getMe")
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}

func (c *telegramClient) ValidateChatID(token string, chatID string) bool {
	testMessage := "Chat ID 확인 중..."

	client := &http.Client{Timeout: 10 * time.Second}
	url := fmt.Sprintf("%s%s/sendMessage?chat_id=%s&text=%s", c.ApiUrl, token, chatID, testMessage)

	resp, err := client.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}

	var msgResp messageResponse
	if err := json.Unmarshal(body, &msgResp); err != nil {
		return false
	}

	if !msgResp.Ok {
		return false
	}

	return true
}
