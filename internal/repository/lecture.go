package repository

import (
	"chee-go-backend/internal/domain/repository"

	"gorm.io/gorm"
)

type lectureRepository struct {
	db *gorm.DB
}

func NewLectureRepository(db *gorm.DB) repository.LectureRepository {
	return &lectureRepository{
		db: db,
	}
}
