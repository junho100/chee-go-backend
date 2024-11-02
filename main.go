package main

import (
	"chee-go-backend/internal/config"
	"chee-go-backend/internal/http"
	"chee-go-backend/internal/http/handler"
	"chee-go-backend/internal/repository"
	"chee-go-backend/internal/service"
)

func main() {
	config := config.NewConfig()
	router := http.NewRouter()

	lectureRepository := repository.NewLectureRepository(config.DB)
	resumeRepository := repository.NewResumeRepository(config.DB)
	userRepository := repository.NewUserRepository(config.DB)

	lectureService := service.NewLectureService(lectureRepository)
	resumeService := service.NewResumeService(resumeRepository)
	userService := service.NewUserService(userRepository)

	handler.NewLectureHandler(router, lectureService)
	handler.NewResumeHandler(router, resumeService)
	handler.NewUserHandler(router, userService)
	handler.NewHealthCheck(router)

	router.Run(":8080")
}
