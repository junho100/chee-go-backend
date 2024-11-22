package main

import (
	"chee-go-backend/internal/config"
	"chee-go-backend/internal/http/handler"
	"chee-go-backend/internal/http/router"
	"chee-go-backend/internal/infrastructure/cron"
	"chee-go-backend/internal/infrastructure/telegram"
	"chee-go-backend/internal/infrastructure/youtube"
	"chee-go-backend/internal/repository"
	"chee-go-backend/internal/service"
)

func main() {
	config := config.NewConfig()
	router := router.NewRouter()

	lectureRepository := repository.NewLectureRepository(config.DB)
	resumeRepository := repository.NewResumeRepository(config.DB)
	userRepository := repository.NewUserRepository(config.DB)
	notificationRepository := repository.NewNotificationRepository(config.DB)

	lectureService := service.NewLectureService(lectureRepository)
	resumeService := service.NewResumeService(resumeRepository)
	userService := service.NewUserService(userRepository)
	notificationService := service.NewNotificationService(notificationRepository)

	youtubeClient := youtube.NewYoutubeClient(config.YoutubeService)
	telegramClient := telegram.NewTelegramClient()

	handler.NewLectureHandler(router, lectureService, youtubeClient)
	handler.NewResumeHandler(router, resumeService, userService)
	handler.NewUserHandler(router, userService)
	handler.NewHealthCheck(router)
	handler.NewNotificationHandler(router, telegramClient, userService, notificationService)

	// Cron job 시작
	cronJob := cron.NewCronJob()
	if cronJob != nil {
		cronJob.Start()
		defer cronJob.Stop()
	}

	router.Run(":8080")
}
