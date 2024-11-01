package main

import (
	"chee-go-backend/internal/config"
	"chee-go-backend/internal/http"
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

	// http.NewLectureHandler(router, lectureService)
	// http.NewResumeHandler(router, resumeService)
	// http.NewUserHandler(router, userService)
	// http.NewHealthCheck(router)

	// serverRoute := r.Group("/api")
	// users.RegisterUsersRouters(serverRoute.Group("/users"), DB)
	// resumes.RegisterResumesRouters(serverRoute.Group("/resumes"), DB)
	// health.RegisterUsersRouters(serverRoute.Group("/health"))
	// lectures.RegisterLecturesRouters(serverRoute.Group("/lectures"), DB)

	router.Run(":8080")
}
