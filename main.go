package main

import (
	"chee-go-backend/internal/config"
	"chee-go-backend/internal/http/handler"
	"chee-go-backend/internal/http/router"
	"chee-go-backend/internal/infrastructure/crawler"
	"chee-go-backend/internal/infrastructure/cron"
	"chee-go-backend/internal/infrastructure/redis"
	"chee-go-backend/internal/infrastructure/telegram"
	"chee-go-backend/internal/infrastructure/youtube"
	"chee-go-backend/internal/repository"
	"chee-go-backend/internal/service"
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

	handler.NewLectureHandler(router, lectureService, youtubeClient)
	handler.NewResumeHandler(router, resumeService, userService)
	handler.NewUserHandler(router, userService)
	handler.NewHealthCheck(router)
	handler.NewNotificationHandler(router, telegramClient, userService, notificationService)

	// Cron job 시작
	cronJob := cron.NewCronJob(
		notificationService,
		telegramClient,
		notificationStatus,
		crawler,
	)
	if cronJob != nil {
		cronJob.Start()
		defer cronJob.Stop()
	}

	router.Run(":8080")
}
