package repository

import (
	"chee-go-backend/internal/domain/repository"

	"gorm.io/gorm"
)

type resumeRepository struct {
	db *gorm.DB
}

func NewResumeRepository(db *gorm.DB) repository.ResumeRepository {
	return &resumeRepository{
		db: db,
	}
}
