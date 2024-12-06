package telegram

import (
	"chee-go-backend/internal/http/dto"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type TelegramClient interface {
	ValidateToken(token string) bool
	ValidateChatID(token string, chatID string) bool
	SendMessage(token string, chatID string, message string) error
	SendNotificationMessage(sendNotificationMessageDto dto.SendNotificationMessageDto) error
}

type telegramClient struct {
	ApiUrl string
}

func (c *telegramClient) SendMessage(token string, chatID string, message string) error {
	client := &http.Client{Timeout: 10 * time.Second}
	url := fmt.Sprintf("%s%s/sendMessage?chat_id=%s&text=%s", c.ApiUrl, token, url.QueryEscape(chatID), message)

	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var msgResp messageResponse
	if err := json.Unmarshal(body, &msgResp); err != nil {
		return err
	}

	if !msgResp.Ok {
		return err
	}

	return nil
}

func (c *telegramClient) SendNotificationMessage(sendNotificationMessageDto dto.SendNotificationMessageDto) error {
	// 메시지 텍스트 생성
	messageText := fmt.Sprintf("[취Go 알림]\n공지사항\n제목: %s\n링크: %s\n작성일: %s",
		sendNotificationMessageDto.Title,
		sendNotificationMessageDto.Url,
		sendNotificationMessageDto.Date.Format("2006-01-02"))

	// URL 인코딩
	encodedText := url.QueryEscape(messageText)

	// SendMessage 호출
	if err := c.SendMessage(
		sendNotificationMessageDto.Token,
		sendNotificationMessageDto.ChatID,
		encodedText,
	); err != nil {
		return fmt.Errorf("텔레그램 메시지 전송 실패: %v", err)
	}

	return nil
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
	testMessage := url.QueryEscape("[취Go 알림]\nChat ID 확인 중...")

	client := &http.Client{Timeout: 10 * time.Second}
	url := fmt.Sprintf("%s%s/sendMessage?chat_id=%s&text=%s", c.ApiUrl, token, url.QueryEscape(chatID), testMessage)

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
