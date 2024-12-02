package discord

import (
	"fmt"
	"log"
	"os"
	"sync"

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

	// Discord 세션 생성
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, fmt.Errorf("Discord 세션 생성 실패: %v", err)
	}

	// 봇이 준비될 때까지 기다리기 위한 WaitGroup
	var wg sync.WaitGroup
	wg.Add(1)

	// Ready 이벤트 핸들러 등록
	discord.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Discord 봇 '%s' 준비 완료", s.State.User.Username)
		wg.Done()
	})

	// 봇이 메시지를 받을 수 있도록 Intents 설정
	discord.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsDirectMessages

	// 봇 연결
	if err := discord.Open(); err != nil {
		return nil, fmt.Errorf("Discord 연결 실패: %v", err)
	}

	// 봇이 준비될 때까지 대기
	wg.Wait()

	return &discordClient{
		session: discord,
	}, nil
}
