package service

import (
	"chee-go-backend/internal/domain/entity"
	"chee-go-backend/internal/http/dto"
)

type ResumeService interface {
	CreateResume(createResumeDto *dto.CreateResumeDTO) (uint, error)
	GetResumeByUserID(userID string) (*entity.Resume, error)
	GetKeywordsByResumeID(id uint) []string
	ConvertResumeToWanted(resume entity.Resume, keywords []string) dto.WantedResume
	ConvertResumeToProgrammers(resume entity.Resume) dto.ProgrammersResume
	ConvertResumeToLinkedin(resume entity.Resume, keywords []string) dto.LinkedinResume
}
