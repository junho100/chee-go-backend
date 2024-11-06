package repository

import (
	"chee-go-backend/internal/domain/entity"

	"gorm.io/gorm"
)

type ResumeRepository interface {
	StartTransaction() (*gorm.DB, error)
	FindResumeByUserId(userId string) (*entity.Resume, error)
	CreateResume(tx *gorm.DB, resume *entity.Resume) error
	DeleteEducationsInResume(tx *gorm.DB, resume *entity.Resume) error
	CreateEducation(tx *gorm.DB, education *entity.Education) error
	DeleteProjectsInResume(tx *gorm.DB, resume *entity.Resume) error
	CreateProject(tx *gorm.DB, project *entity.Project) error
	DeleteActivitiesInResume(tx *gorm.DB, resume *entity.Resume) error
	CreateActivity(tx *gorm.DB, activity *entity.Activity) error
	DeleteWorkExperiencesInResume(tx *gorm.DB, resume *entity.Resume) error
	CreateWorkExperience(tx *gorm.DB, workExperience *entity.WorkExperience) error
	CreateWorkExperienceDetail(tx *gorm.DB, workExperienceDetail *entity.WorkExperienceDetail) error
	DeleteKeywordsInResume(tx *gorm.DB, resume *entity.Resume) error
	FindKeywordByName(keywordName string, keyword *entity.Keyword) error
	CreateKeyword(tx *gorm.DB, keyword *entity.Keyword) error
	CreateKeywordResume(tx *gorm.DB, keywordResume *entity.KeywordResume) error
	FindKeywordsByResumeId(resumeId uint) ([]entity.Keyword, error)
	CreateCertificate(tx *gorm.DB, certificate *entity.Certificate) error
	DeleteCertificatesInResume(tx *gorm.DB, resume *entity.Resume) error
}
