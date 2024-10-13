package main

import (
	"chee-go-backend/common"
	"chee-go-backend/resumes"
	"chee-go-backend/users"
	"log"

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

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}

	r.Use(cors.New(config))

	r.SetTrustedProxies(nil)

	serverRoute := r.Group("/api")
	users.RegisterUsersRouters(serverRoute.Group("/users"))
	resumes.RegisterResumesRouters(serverRoute.Group("/resumes"))

	r.Run(":8080")
}
