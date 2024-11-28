package cron

import (
	"chee-go-backend/internal/domain/entity"
	"chee-go-backend/internal/domain/service"
	"chee-go-backend/internal/http/dto"
	"chee-go-backend/internal/infrastructure/crawler"
	"chee-go-backend/internal/infrastructure/redis"
	"chee-go-backend/internal/infrastructure/telegram"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/robfig/cron/v3"
)

// user -> key -> notification 순서
type CronJob struct {
	cron                *cron.Cron
	notificationService service.NotificationService
	telegramClient      telegram.TelegramClient
	notificationStatus  redis.NotificationStatus
	crawler             crawler.Crawler
}

func NewCronJob(notificationService service.NotificationService, telegramClient telegram.TelegramClient, notificationStatus redis.NotificationStatus, crawler crawler.Crawler) *CronJob {
	// UTC+9 시간으로 고정 설정
	loc := time.FixedZone("KST", 9*60*60) // UTC+9 시간

	c := cron.New(cron.WithLocation(loc))
	return &CronJob{
		cron:                c,
		notificationService: notificationService,
		telegramClient:      telegramClient,
		notificationStatus:  notificationStatus,
		crawler:             crawler,
	}
}

func (c *CronJob) Start() {
	// 매일 오전 11시에 실행 (KST)
	c.cron.AddFunc("0 11 * * *", func() {
		// 1. 오늘 올라온 모든 공지사항 크롤링 및 DB 저장
		deptNotices, err := c.crawler.FetchDepartmentNotices()
		if err != nil {
			log.Println("학과 공지사항 크롤링 실패:", err)
			return
		}

		generalNotices, err := c.crawler.FetchSchoolNotices()
		if err != nil {
			log.Println("일반 공지사항 크롤링 실패:", err)
			return
		}

		allNotices := append(deptNotices, generalNotices...)
		if err := c.notificationService.SaveTodayNotifications(allNotices); err != nil {
			log.Println("공지사항 저장 실패:", err)
			return
		}

		// 2. 모든 알림 설정 및 키워드 조회
		var notificationConfigs []entity.NotificationConfig
		if err := c.notificationService.FindAllNotificationConfigs(&notificationConfigs); err != nil {
			log.Println("알림 설정 조회 실패:", err)
			return
		}

		// 3. 각 사용자의 키워드에 매칭되는 공지사항 필터링 및 알림 전송
		for _, config := range notificationConfigs {
			keywords := c.notificationService.FindKeywordsByConfigID(config.ID)

			// 키워드에 매칭되는 공지사항 필터링
			matchedNotices := []entity.SchoolNotification{}
			for _, notice := range allNotices {
				for _, keyword := range keywords {
					if strings.Contains(notice.Title, keyword) {
						// 4. 이미 처리된 알림인지 확인
						processed, err := c.notificationStatus.IsProcessed(config.UserID, notice.ID)
						if err != nil {
							log.Printf("알림 상태 확인 실패 (사용자: %s, 공지: %s): %v", config.UserID, notice.ID, err)
							continue
						}
						if !processed {
							matchedNotices = append(matchedNotices, notice)
						}
						break
					}
				}
			}

			// 매칭된 공지사항이 없으면 다음 사용자로
			if len(matchedNotices) == 0 {
				now := time.Now().In(c.cron.Location())
				checkTime := now.Format("2006-01-02 15:04")
				message := fmt.Sprintf("[취Go 알림]\n발견된 공지사항이 없습니다.\n확인 시각: %s", checkTime)
				encodedText := url.QueryEscape(message)
				if err := c.telegramClient.SendMessage(config.TelegramToken, config.TelegramChatID, encodedText); err != nil {
					log.Printf("알림 전송 실패 (사용자: %s): %v", config.UserID, err)
				}
				continue
			}

			// 매칭된 공지사항에 대해 알림 전송
			for _, notice := range matchedNotices {
				var url string
				if os.Getenv("CLIENT_BASE_URL") == "" {
					url = "https://chee-go.com"
				} else {
					url = os.Getenv("CLIENT_BASE_URL")
				}

				sendNotificationMessageDto := dto.SendNotificationMessageDto{
					Title:  notice.Title,
					Date:   notice.Date,
					Url:    fmt.Sprintf("%s/notification/%s", url, notice.ID),
					Token:  config.TelegramToken,
					ChatID: config.TelegramChatID,
				}

				if err := c.telegramClient.SendNotificationMessage(sendNotificationMessageDto); err != nil {
					log.Printf("알림 전송 실패 (사용자: %s, 공지: %s): %v", config.UserID, notice.ID, err)
					continue
				}

				// 5. 알림 처리 상태를 Redis에 저장
				if err := c.notificationStatus.MarkAsProcessed(config.UserID, notice.ID); err != nil {
					log.Printf("알림 상태 저장 실패 (사용자: %s, 공지: %s): %v", config.UserID, notice.ID, err)
				}
			}
		}
	})

	c.cron.Start()
	log.Println("Cron job started (KST - runs at 11:00 AM)")
}

func (c *CronJob) Stop() {
	if c.cron != nil {
		c.cron.Stop()
		log.Println("Cron job stopped")
	}
}
