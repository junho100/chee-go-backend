package util

import (
	"chee-go-backend/internal/domain/entity"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, func()) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", "test", "test", "localhost", "cheego_test")
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

	// cleanup 함수 정의
	cleanup := func() {
		// 모든 테이블의 데이터 삭제
		tables := []interface{}{
			&entity.Lecture{},
			&entity.Subject{},
			&entity.WorkExperienceDetail{},
			&entity.WorkExperience{},
			&entity.Certificate{},
			&entity.Activity{},
			&entity.KeywordResume{},
			&entity.Keyword{},
			&entity.Project{},
			&entity.Education{},
			&entity.Resume{},
			&entity.User{},
		}

		// Foreign key checks 비활성화
		db.Exec("SET FOREIGN_KEY_CHECKS = 0")

		// 각 테이블의 데이터 삭제
		for _, table := range tables {
			if err := db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(table).Error; err != nil {
				log.Fatalf("Error cleaning up table: %v", err)
			}
		}

		// Foreign key checks 활성화
		db.Exec("SET FOREIGN_KEY_CHECKS = 1")
	}

	return db, cleanup
}
