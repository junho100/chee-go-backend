package resumes

import (
	"chee-go-backend/common"
	"chee-go-backend/users"
	"time"
)

type Resume struct {
	ID              uint   `gorm:"primary_key"`
	Introduction    string `gorm:"column:introduction"`
	GithubURL       string `gorm:"column:github_url"`
	BlogURL         string `gorm:"column:blog_url"`
	UserID          string
	User            users.User
	Educations      []Education
	Projects        []Project
	Activities      []Activity
	Certificates    []Certificate
	WorkExperiences []WorkExperience
}

type Education struct {
	ID         uint      `gorm:"primary_key"`
	SchoolName string    `gorm:"column:school_name"`
	MajorName  string    `gorm:"column:major_name"`
	StartDate  time.Time `gorm:"column:start_date;type:date"`
	EndDate    time.Time `gorm:"column:end_date;type:date"`
	ResumeID   uint
}

type Project struct {
	ID        uint      `gorm:"primary_key"`
	StartDate time.Time `gorm:"column:start_date;type:date"`
	EndDate   time.Time `gorm:"column:end_date;type:date"`
	Content   string    `gorm:"column:content"`
	GithubURL string    `gorm:"column:github_url"`
	ResumeID  uint
}

type Keyword struct {
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"column:name"`
}

type KeywordResume struct {
	KeywordID uint    `gorm:"primary_key"`
	ResumeID  uint    `gorm:"primary_key"`
	Keyword   Keyword `gorm:"foreignKey:KeywordID"`
	Resume    Resume  `gorm:"foreignKey:ResumeID"`
}

type Activity struct {
	ID       uint   `gorm:"primary_key"`
	Name     string `gorm:"column:name"`
	Content  string `gorm:"column:content"`
	ResumeID uint
}

type Certificate struct {
	ID         uint      `gorm:"primary_key"`
	Name       string    `gorm:"column:name"`
	IssuedBy   string    `gorm:"column:issued_by"`
	IssuedDate time.Time `gorm:"column:issued_date;type:date"`
	ResumeID   uint
}

type WorkExperience struct {
	ID                    uint   `gorm:"primary_key"`
	CompanyName           string `gorm:"column:company_name"`
	Department            string `gorm:"column:department"`
	Position              string `gorm:"column:position"`
	Job                   string `gorm:"column:job"`
	ResumeID              uint
	WorkExperienceDetails []WorkExperienceDetail `gorm:"constraint:OnDelete:CASCADE;"`
}

type WorkExperienceDetail struct {
	ID               uint      `gorm:"primary_key"`
	Name             string    `gorm:"column:name"`
	StartDate        time.Time `gorm:"column:start_date;type:date"`
	EndDate          time.Time `gorm:"column:end_date;type:date"`
	Content          string    `gorm:"column:content"`
	WorkExperienceID uint
	WorkExperience   WorkExperience
}

type RegisterResumeRequest struct {
	Introduction    string
	GithubURL       string `json:"github_url"`
	BlogURL         string `json:"blog_url"`
	Educations      []RegisterResumeRequestReducation
	Projects        []RegisterResumeRequestProject
	Activities      []RegisterResumeRequestActivity
	Certificates    []RegisterResumeRequestCertificate
	WorkExperiences []RegisterResumeRequestWorkExperience
	Keywords        []string
}

type RegisterResumeResponse struct {
	ResumeID uint
}

type RegisterResumeRequestReducation struct {
	SchoolName string    `json:"school_name"`
	MajorName  string    `json:"major_name"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
}

type RegisterResumeRequestProject struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Content   string
	GithubURL string `json:"github_url"`
}

type RegisterResumeRequestActivity struct {
	Name    string
	Content string
}

type RegisterResumeRequestCertificate struct {
	Name       string
	IssuedBy   string    `json:"issued_by"`
	IssuedDate time.Time `json:"issued_date"`
}

type RegisterResumeRequestWorkExperience struct {
	CompanyName string `json:"company_name"`
	Department  string
	Position    string
	Job         string
	Details     []RegisterResumeRequestWorkExperienceDetail
}

type RegisterResumeRequestWorkExperienceDetail struct {
	Name      string
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Content   string
}

type CreateResumeDTO struct {
	Introduction    string
	GithubURL       string
	BlogURL         string
	Educations      []RegisterResumeRequestReducation
	Projects        []RegisterResumeRequestProject
	Activities      []RegisterResumeRequestActivity
	Certificates    []RegisterResumeRequestCertificate
	WorkExperiences []RegisterResumeRequestWorkExperience
	Keywords        []string
	UserID          string
}

func CreateResume(dto *CreateResumeDTO) (uint, error) {
	var resume Resume
	db := common.GetDB()
	tx := db.Begin()

	if err := db.Where(&Resume{
		UserID: dto.UserID,
	}).First(&resume).Error; err != nil {
		resume = Resume{
			Introduction: dto.Introduction,
			GithubURL:    dto.GithubURL,
			BlogURL:      dto.BlogURL,
			UserID:       dto.UserID,
		}
	}

	if err := tx.Save(&resume).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	if err := tx.Where(&Education{
		ResumeID: resume.ID,
	}).Delete(&Education{}).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	for _, education := range dto.Educations {
		savedEducation := &Education{
			SchoolName: education.SchoolName,
			MajorName:  education.MajorName,
			StartDate:  education.StartDate,
			EndDate:    education.EndDate,
			ResumeID:   resume.ID,
		}

		if err := tx.Save(&savedEducation).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	if err := tx.Where(&Project{
		ResumeID: resume.ID,
	}).Delete(&Project{}).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	for _, project := range dto.Projects {
		savedProject := &Project{
			StartDate: project.StartDate,
			EndDate:   project.EndDate,
			Content:   project.Content,
			GithubURL: project.GithubURL,
			ResumeID:  resume.ID,
		}

		if err := tx.Save(&savedProject).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	if err := tx.Where(&Activity{
		ResumeID: resume.ID,
	}).Delete(&Activity{}).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	for _, activity := range dto.Activities {
		savedActivity := &Activity{
			Name:    activity.Name,
			Content: activity.Content,
		}

		if err := tx.Save(&savedActivity).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	if err := tx.Where(&Certificate{
		ResumeID: resume.ID,
	}).Delete(&Certificate{}).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	for _, certificate := range dto.Certificates {
		savedCertificate := &Certificate{
			Name:       certificate.Name,
			IssuedBy:   certificate.IssuedBy,
			IssuedDate: certificate.IssuedDate,
		}

		if err := tx.Save(&savedCertificate).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	if err := tx.Where(&WorkExperience{
		ResumeID: resume.ID,
	}).Delete(&WorkExperience{}).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	for _, workExperience := range dto.WorkExperiences {
		savedWorkExperience := &WorkExperience{
			CompanyName: workExperience.CompanyName,
			Department:  workExperience.Department,
			Position:    workExperience.Position,
			Job:         workExperience.Job,
		}

		if err := tx.Save(&savedWorkExperience).Error; err != nil {
			tx.Rollback()
			return 0, err
		}

		for _, detail := range workExperience.Details {
			savedDetail := &WorkExperienceDetail{
				Name:             detail.Name,
				StartDate:        detail.StartDate,
				EndDate:          detail.EndDate,
				Content:          detail.Content,
				WorkExperienceID: savedWorkExperience.ID,
			}

			if err := tx.Save(&savedDetail).Error; err != nil {
				tx.Rollback()
				return 0, err
			}
		}
	}

	return resume.ID, tx.Commit().Error
}
