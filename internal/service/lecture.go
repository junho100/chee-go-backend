package service

import (
	"chee-go-backend/internal/domain/repository"
	"chee-go-backend/internal/domain/service"
)

type lectureService struct {
	lectureRepository repository.LectureRepository
}

func NewLectureService(lectureRepository repository.LectureRepository) service.LectureService {
	return &lectureService{
		lectureRepository: lectureRepository,
	}
}
