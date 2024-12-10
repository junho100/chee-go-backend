package mock

import (
	"fmt"
	"time"
)

type MockDiscordClient struct {
	ValidateClientIDFunc func(clientID string) bool
	SendMessageFunc      func(clientID string, message string) error
}

func (m *MockDiscordClient) ValidateClientID(clientID string) bool {
	return m.ValidateClientIDFunc(clientID)
}

func (m *MockDiscordClient) SendMessage(clientID string, message string) error {
	return m.SendMessageFunc(clientID, message)
}

func (m *MockDiscordClient) SendNotificationMessage(clientID string, title string, url string, date time.Time) error {
	message := fmt.Sprintf("[취Go 알림]\n공지사항\n제목: %s\n링크: %s\n작성일: %s",
		title,
		url,
		date.Format("2006-01-02"))
	return m.SendMessageFunc(clientID, message)
}
