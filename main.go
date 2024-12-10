package main

import (
	"chee-go-backend/internal/config"
	"chee-go-backend/internal/http/handler"
	"chee-go-backend/internal/http/router"
	"chee-go-backend/internal/infrastructure/crawler"
	"chee-go-backend/internal/infrastructure/cron"
	"chee-go-backend/internal/infrastructure/discord"
	"chee-go-backend/internal/infrastructure/monitoring"
	"chee-go-backend/internal/infrastructure/redis"
	"chee-go-backend/internal/infrastructure/telegram"
	"chee-go-backend/internal/infrastructure/youtube"
	"chee-go-backend/internal/repository"
	"chee-go-backend/internal/service"
	"log"
	"os"
)

func main() {
	cfg := config.NewConfig()
	router := router.NewRouter()
	notificationStatus := redis.NewNotificationStatus(cfg.RedisClient)

	lectureRepository := repository.NewLectureRepository(cfg.DB)
	resumeRepository := repository.NewResumeRepository(cfg.DB)
	userRepository := repository.NewUserRepository(cfg.DB)
	notificationRepository := repository.NewNotificationRepository(cfg.DB)

	crawler := crawler.NewCrawler(
		cfg.SchoolNoticeURL,
		cfg.DeptNoticeURLs,
	)

	lectureService := service.NewLectureService(lectureRepository)
	resumeService := service.NewResumeService(resumeRepository)
	userService := service.NewUserService(userRepository)
	notificationService := service.NewNotificationService(notificationRepository, crawler)

	youtubeClient := youtube.NewYoutubeClient(cfg.YoutubeService)
	telegramClient := telegram.NewTelegramClient()
	discordClient, err := discord.NewDiscordClient()
	if err != nil {
		log.Fatalf("Discord 클라이언트 초기화 실패: %v", err)
	}

	handler.NewLectureHandler(router, lectureService, youtubeClient)
	handler.NewResumeHandler(router, resumeService, userService)
	handler.NewUserHandler(router, userService)
	handler.NewHealthCheck(router)
	handler.NewNotificationHandler(router, telegramClient, discordClient, userService, notificationService)

	var batchMetrics *monitoring.BatchMetrics

	// 로컬 환경이 아닐 때만 모니터링 활성화
	if os.Getenv("GO_ENV") != "local" {
		cloudwatch, err := monitoring.NewCloudWatchClient()
		if err != nil {
			log.Fatalf("CloudWatch 클라이언트 초기화 실패: %v", err)
		}
		batchMetrics = monitoring.NewBatchMetrics(cloudwatch)
	}

	// Cron job 시작
	cronJob := cron.NewCronJob(
		notificationService,
		telegramClient,
		discordClient,
		notificationStatus,
		crawler,
	)

	if cronJob != nil {
		go func() {
			if batchMetrics != nil {
				// 프로덕션 환경: 각 크론 작업 실행을 모니터링
				cronJob.SetJobWrapper(func(job func()) func() {
					return func() {
						err := batchMetrics.TrackBatchJob("DailyNotificationDelivery", func() error {
							job()
							return nil
						})
						if err != nil {
							log.Printf("배치 작업 메트릭 기록 실패: %v", err)
						}
					}
				})
			}
			// 크론 작업 시작
			cronJob.Start()
		}()
		defer cronJob.Stop()
	}

	router.Run(":8080")
}
