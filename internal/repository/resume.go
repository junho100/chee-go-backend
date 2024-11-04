package repository

import (
	"chee-go-backend/internal/domain/entity"
	"chee-go-backend/internal/domain/repository"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type resumeRepository struct {
	db *gorm.DB
}

func (r *resumeRepository) CreateActivity(tx *gorm.DB, activity *entity.Activity) error {
	if err := tx.Save(activity).Error; err != nil {
		return err
	}

	return nil
}

func (r *resumeRepository) CreateEducation(tx *gorm.DB, education *entity.Education) error {
	if err := tx.Save(education).Error; err != nil {
		return err
	}

	return nil
}

func (r *resumeRepository) CreateKeyword(tx *gorm.DB, keyword *entity.Keyword) error {
	if err := tx.Save(keyword).Error; err != nil {
		return err
	}

	return nil
}

func (r *resumeRepository) CreateKeywordResume(tx *gorm.DB, keywordResume *entity.KeywordResume) error {
	if err := tx.Save(keywordResume).Error; err != nil {
		return err
	}

	return nil
}

func (r *resumeRepository) CreateProject(tx *gorm.DB, project *entity.Project) error {
	if err := tx.Save(project).Error; err != nil {
		return err
	}

	return nil
}

func (r *resumeRepository) CreateResume(tx *gorm.DB, resume *entity.Resume) error {
	if err := tx.Save(resume).Error; err != nil {
		return err
	}

	return nil
}

func (r *resumeRepository) CreateWorkExperience(tx *gorm.DB, workExperience *entity.WorkExperience) error {
	if err := tx.Save(workExperience).Error; err != nil {
		return err
	}

	return nil
}

func (r *resumeRepository) CreateWorkExperienceDetail(tx *gorm.DB, workExperienceDetail *entity.WorkExperienceDetail) error {
	if err := tx.Save(workExperienceDetail).Error; err != nil {
		return err
	}

	return nil
}

func (r *resumeRepository) DeleteActivitiesInResume(tx *gorm.DB, resume *entity.Resume) error {
	if err := tx.Where(&entity.Activity{
		ResumeID: resume.ID,
	}).Delete(&entity.Activity{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *resumeRepository) DeleteEducationsInResume(tx *gorm.DB, resume *entity.Resume) error {
	if err := tx.Where(&entity.Education{
		ResumeID: resume.ID,
	}).Delete(&entity.Education{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r *resumeRepository) DeleteKeywordsInResume(tx *gorm.DB, resume *entity.Resume) error {
	if err := tx.Where(&entity.KeywordResume{
		ResumeID: resume.ID,
	}).Delete(&entity.KeywordResume{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r *resumeRepository) DeleteProjectsInResume(tx *gorm.DB, resume *entity.Resume) error {
	if err := tx.Where(&entity.Project{
		ResumeID: resume.ID,
	}).Delete(&entity.Project{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r *resumeRepository) DeleteWorkExperiencesInResume(tx *gorm.DB, resume *entity.Resume) error {
	if err := tx.Where(&entity.WorkExperience{
		ResumeID: resume.ID,
	}).Delete(&entity.WorkExperience{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r *resumeRepository) FindKeywordByName(keywordName string, keyword *entity.Keyword) error {
	if err := r.db.Where(&entity.Keyword{
		Name: keywordName,
	}).First(keyword).Error; err != nil {
		return err
	}

	return nil
}

func (r *resumeRepository) FindResumeByUserId(userId string) (*entity.Resume, error) {
	var resume entity.Resume

	if err := r.db.Preload("WorkExperiences.WorkExperienceDetails").
		Preload(clause.Associations).Where(&entity.Resume{
		UserID: userId,
	}).First(&resume).Error; err != nil {
		return nil, err
	}

	return &resume, nil
}

func (r *resumeRepository) StartTransaction() (*gorm.DB, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	return tx, nil
}

func (r *resumeRepository) FindKeywordsByResumeId(resumeId uint) ([]entity.Keyword, error) {
	var keywordResumes []entity.KeywordResume

	if err := r.db.Preload("Keyword").Where(&entity.KeywordResume{
		ResumeID: resumeId,
	}).Find(&keywordResumes).Error; err != nil {
		return nil, err
	}

	keywords := make([]entity.Keyword, len(keywordResumes))
	for i, keywordResume := range keywordResumes {
		keywords[i] = keywordResume.Keyword
	}

	return keywords, nil
}

func NewResumeRepository(db *gorm.DB) repository.ResumeRepository {
	return &resumeRepository{
		db: db,
	}
}
