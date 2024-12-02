package config

import (
	"chee-go-backend/internal/domain/entity"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DB                 *gorm.DB
	YoutubeService     *youtube.Service
	RedisClient        *redis.Client
	DeptGeneralURL     string
	DeptScholarshipURL string
	SchoolNoticeURL    string   // 학교 공지사항 URL
	DeptNoticeURLs     []string // 학과 공지사항 URL들
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

	redisClient := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
		DB:   0,
	})

	// Discord 봇 토큰 검증
	if os.Getenv("DISCORD_BOT_TOKEN") == "" {
		log.Fatal("DISCORD_BOT_TOKEN이 설정되지 않았습니다")
	}

	return &Config{
		DB:                 db,
		YoutubeService:     service,
		RedisClient:        redisClient,
		DeptGeneralURL:     os.Getenv("DEPT_GENERAL_URL"),
		DeptScholarshipURL: os.Getenv("DEPT_SCHOLARSHIP_URL"),
		SchoolNoticeURL:    os.Getenv("SCHOOL_NOTICE_URL"),
		DeptNoticeURLs: []string{
			os.Getenv("DEPT_NOTICE_URL_1"),
			os.Getenv("DEPT_NOTICE_URL_2"),
			os.Getenv("DEPT_NOTICE_URL_3"),
		},
	}
}
