package service

import (
	"chee-go-backend/internal/domain/repository"
	"chee-go-backend/internal/domain/service"
)

type resumeService struct {
	resumeRepository repository.ResumeRepository
}

func NewResumeService(resumeRepository repository.ResumeRepository) service.ResumeService {
	return &resumeService{
		resumeRepository: resumeRepository,
	}
}
