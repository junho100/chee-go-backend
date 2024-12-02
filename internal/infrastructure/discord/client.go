package discord

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
)

type DiscordClient interface {
	ValidateClientID(clientID string) bool
}

type discordClient struct {
	session *discordgo.Session
}

func (c *discordClient) ValidateClientID(clientID string) bool {
	// DM 채널 생성 시도
	channel, err := c.session.UserChannelCreate(clientID)
	if err != nil {
		log.Printf("DM 채널 생성 실패: %v", err)
		return false
	}

	// 테스트 메시지 전송
	testMessage := "[취Go 알림]\nDiscord Client ID 확인 중..."
	_, err = c.session.ChannelMessageSend(channel.ID, testMessage)
	if err != nil {
		log.Printf("메시지 전송 실패: %v", err)
		return false
	}

	return true
}

func NewDiscordClient() (DiscordClient, error) {
	token := os.Getenv("DISCORD_BOT_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("DISCORD_BOT_TOKEN이 설정되지 않았습니다")
	}

	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, fmt.Errorf("Discord 세션 생성 실패: %v", err)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	discord.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Discord 봇 '%s' 준비 완료", s.State.User.Username)
		wg.Done()
	})

	discord.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsDirectMessages

	if err := discord.Open(); err != nil {
		return nil, fmt.Errorf("Discord 연결 실패: %v", err)
	}

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-ctx.Done():
		discord.Close()
		return nil, fmt.Errorf("Discord 봇 준비 시간 초과 (30초)")
	case <-done:
		return &discordClient{
			session: discord,
		}, nil
	}
}
