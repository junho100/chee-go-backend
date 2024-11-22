package config

import (
	"chee-go-backend/internal/domain/entity"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DB             *gorm.DB
	YoutubeService *youtube.Service
}

func NewConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	db_username := os.Getenv("DB_USERNAME")
	db_password := os.Getenv("DB_PASSWORD")
	db_host := os.Getenv("DB_HOST")
	db_name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_username, db_password, db_host, db_name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error initialize database: %s", err)
	}

	db.AutoMigrate(
		&entity.User{},
		&entity.Resume{},
		&entity.Education{},
		&entity.Project{},
		&entity.Keyword{},
		&entity.KeywordResume{},
		&entity.Activity{},
		&entity.Certificate{},
		&entity.WorkExperience{},
		&entity.WorkExperienceDetail{},
		&entity.Subject{},
		&entity.Lecture{},
		&entity.NotificationConfig{},
		&entity.NotificationKeyword{},
		&entity.NotificationConfigKeyword{},
		&entity.SchoolNotification{},
	)

	ctx := context.Background()
	apiKey := os.Getenv("YOUTUBE_API_KEY")

	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))

	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
	}

	return &Config{
		DB:             db,
		YoutubeService: service,
	}
}
