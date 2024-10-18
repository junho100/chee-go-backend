package main

import (
	"chee-go-backend/common"
	"chee-go-backend/health"
	"chee-go-backend/resumes"
	"chee-go-backend/users"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	DB := common.Init()
	DB.AutoMigrate(&users.User{}, &resumes.Resume{}, &resumes.Education{}, &resumes.Project{}, &resumes.Keyword{}, &resumes.KeywordResume{}, &resumes.Activity{}, &resumes.Certificate{}, &resumes.WorkExperience{}, &resumes.WorkExperienceDetail{})

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.SetTrustedProxies(nil)

	serverRoute := r.Group("/api")
	users.RegisterUsersRouters(serverRoute.Group("/users"))
	resumes.RegisterResumesRouters(serverRoute.Group("/resumes"))
	health.RegisterUsersRouters(serverRoute.Group("/health"))

	r.Run(":8080")
}
