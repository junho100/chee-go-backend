package resumes

import (
	"chee-go-backend/common"
	"chee-go-backend/users"
	"time"

	"gorm.io/gorm/clause"
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
	Name      string    `gorm:"column:name"`
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
	ID        uint      `gorm:"primary_key"`
	Name      string    `gorm:"column:name"`
	Content   string    `gorm:"column:content"`
	StartDate time.Time `gorm:"column:start_date;type:date"`
	EndDate   time.Time `gorm:"column:end_date;type:date"`
	ResumeID  uint
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
	StartDate             time.Time              `gorm:"column:start_date;type:date"`
	EndDate               time.Time              `gorm:"column:end_date;type:date"`
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
	Educations      []RegisterResumeRequestEducation
	Projects        []RegisterResumeRequestProject
	Activities      []RegisterResumeRequestActivity
	Certificates    []RegisterResumeRequestCertificate
	WorkExperiences []RegisterResumeRequestWorkExperience `json:"work_experiences"`
	Keywords        []string                              `json:"keywords"`
}

type RegisterResumeResponse struct {
	ResumeID uint
}

type RegisterResumeRequestEducation struct {
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
	Name      string `json:"name"`
}

type RegisterResumeRequestActivity struct {
	Name      string
	Content   string
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
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
	Details     []RegisterResumeRequestWorkExperienceDetail `json:"details"`
	StartDate   time.Time                                   `json:"start_date"`
	EndDate     time.Time                                   `json:"end_date"`
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
	Educations      []RegisterResumeRequestEducation
	Projects        []RegisterResumeRequestProject
	Activities      []RegisterResumeRequestActivity
	Certificates    []RegisterResumeRequestCertificate
	WorkExperiences []RegisterResumeRequestWorkExperience
	Keywords        []string
	UserID          string
}

type GetResumeResponse struct {
	ID              uint                              `json:"id"`
	Introduction    string                            `json:"introduction"`
	GithubURL       string                            `json:"github_url"`
	BlogURL         string                            `json:"blog_url"`
	Educations      []GetResumeResponseEducation      `json:"educations"`
	Projects        []GetResumeResponseProject        `json:"projects"`
	Activities      []GetResumeResponseActivity       `json:"activities"`
	Certificates    []GetResumeResponseCertificate    `json:"certificates"`
	WorkExperiences []GetResumeResponseWorkExperience `json:"work_experiences"`
	Keywords        []string                          `json:"keywords"`
}

type GetResumeResponseEducation struct {
	ID         uint      `json:"id"`
	SchoolName string    `json:"school_name"`
	MajorName  string    `json:"major_name"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
}

type GetResumeResponseProject struct {
	ID        uint      `json:"id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Content   string    `json:"content"`
	GithubURL string    `json:"github_url"`
	Name      string    `json:"name"`
}

type GetResumeResponseActivity struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Content   string    `json:"content"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type GetResumeResponseCertificate struct {
	ID         uint      `json:"id"`
	Name       string    `json:"name"`
	IssuedBy   string    `json:"issued_by"`
	IssuedDate time.Time `json:"issued_date"`
}

type GetResumeResponseWorkExperience struct {
	ID                    uint                                    `json:"id"`
	CompanyName           string                                  `json:"company_name"`
	Department            string                                  `json:"department"`
	Position              string                                  `json:"position"`
	Job                   string                                  `json:"job"`
	WorkExperienceDetails []GetResumeResponseWorkExperienceDetail `json:"work_experience_details"`
	StartDate             time.Time                               `json:"start_date"`
	EndDate               time.Time                               `json:"end_date"`
}

type GetResumeResponseWorkExperienceDetail struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Content   string    `json:"content"`
}

type WantedResume struct {
	Introduction    string
	WorkExperiences []WantedResumeWorkExperience
	Educations      []WantedResumeEducation
	Skills          []string
	// 수상 및 기타 항목에 프로젝트, 자격증, 활동 기입
	Certificates []WantedResumeCertificate
}

type WantedResumeWorkExperience struct {
	CompanyName string
	Position    string
	StartDate   time.Time
	EndDate     time.Time
	Details     []WantedResumeWorkExperienceDetail
}

type WantedResumeWorkExperienceDetail struct {
	Name      string
	StartDate time.Time
	EndDate   time.Time
	Content   string
}

type WantedResumeEducation struct {
	SchoolName string
	MajorName  string
	StartDate  time.Time
	EndDate    time.Time
}

type WantedResumeCertificate struct {
	Name      string
	Content   string
	StartDate time.Time
}

type GetWantedResumeResponse struct {
	Introduction    string
	WorkExperiences []GetWantedResumeResponseWorkExperience
	Educations      []GetWantedResumeResponseEducation
	Skills          []string
	// 수상 및 기타 항목에 프로젝트, 자격증, 활동 기입
	Certificates []GetWantedResumeResponseCertificate
}

type GetWantedResumeResponseWorkExperience struct {
	CompanyName string
	Position    string
	StartDate   time.Time
	EndDate     time.Time
	Details     []GetWantedResumeResponseWorkExperienceDetail
}

type GetWantedResumeResponseWorkExperienceDetail struct {
	Name      string
	StartDate time.Time
	EndDate   time.Time
	Content   string
}

type GetWantedResumeResponseEducation struct {
	SchoolName string
	MajorName  string
	StartDate  time.Time
	EndDate    time.Time
}

type GetWantedResumeResponseCertificate struct {
	Name      string
	Content   string
	StartDate time.Time
}

func (r *GetResumeResponse) from(resume Resume, keywords []string) *GetResumeResponse {
	r.ID = resume.ID
	r.Introduction = resume.Introduction
	r.GithubURL = resume.GithubURL
	r.BlogURL = resume.BlogURL
	r.Keywords = keywords

	r.Educations = make([]GetResumeResponseEducation, len(resume.Educations))
	for i, education := range resume.Educations {
		r.Educations[i] = GetResumeResponseEducation{
			ID:         education.ID,
			SchoolName: education.SchoolName,
			MajorName:  education.MajorName,
			StartDate:  education.StartDate,
			EndDate:    education.EndDate,
		}
	}

	r.Projects = make([]GetResumeResponseProject, len(resume.Projects))
	for i, project := range resume.Projects {
		r.Projects[i] = GetResumeResponseProject{
			ID:        project.ID,
			StartDate: project.StartDate,
			EndDate:   project.EndDate,
			Content:   project.Content,
			GithubURL: project.GithubURL,
			Name:      project.Name,
		}
	}

	r.Activities = make([]GetResumeResponseActivity, len(resume.Activities))
	for i, activity := range resume.Activities {
		r.Activities[i] = GetResumeResponseActivity{
			ID:        activity.ID,
			Name:      activity.Name,
			Content:   activity.Content,
			StartDate: activity.StartDate,
			EndDate:   activity.EndDate,
		}
	}

	r.Certificates = make([]GetResumeResponseCertificate, len(resume.Certificates))
	for i, certificate := range resume.Certificates {
		r.Certificates[i] = GetResumeResponseCertificate{
			ID:         certificate.ID,
			Name:       certificate.Name,
			IssuedBy:   certificate.IssuedBy,
			IssuedDate: certificate.IssuedDate,
		}
	}

	r.WorkExperiences = make([]GetResumeResponseWorkExperience, len(resume.WorkExperiences))
	for i, workExperience := range resume.WorkExperiences {
		r.WorkExperiences[i] = GetResumeResponseWorkExperience{
			ID:          workExperience.ID,
			CompanyName: workExperience.CompanyName,
			Department:  workExperience.Department,
			Position:    workExperience.Position,
			Job:         workExperience.Job,
			StartDate:   workExperience.StartDate,
			EndDate:     workExperience.EndDate,
		}

		r.WorkExperiences[i].WorkExperienceDetails = make([]GetResumeResponseWorkExperienceDetail, len(workExperience.WorkExperienceDetails))
		for j, workExperienceDetail := range workExperience.WorkExperienceDetails {
			r.WorkExperiences[i].WorkExperienceDetails[j] = GetResumeResponseWorkExperienceDetail{
				ID:        workExperienceDetail.ID,
				Name:      workExperienceDetail.Name,
				StartDate: workExperienceDetail.StartDate,
				EndDate:   workExperienceDetail.EndDate,
				Content:   workExperienceDetail.Content,
			}
		}
	}

	return r
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
	} else {
		resume.Introduction = dto.Introduction
		resume.GithubURL = dto.GithubURL
		resume.BlogURL = dto.BlogURL
		resume.UserID = dto.UserID
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
			Name:      project.Name,
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
			Name:      activity.Name,
			Content:   activity.Content,
			ResumeID:  resume.ID,
			StartDate: activity.StartDate,
			EndDate:   activity.EndDate,
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
			ResumeID:   resume.ID,
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
			ResumeID:    resume.ID,
			StartDate:   workExperience.StartDate,
			EndDate:     workExperience.EndDate,
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

	if err := tx.Where(&KeywordResume{
		ResumeID: resume.ID,
	}).Delete(&KeywordResume{}).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	for _, keywordName := range dto.Keywords {
		var savedKeyword Keyword
		if err := db.Where(&Keyword{
			Name: keywordName,
		}).First(&savedKeyword).Error; err != nil {
			savedKeyword = Keyword{
				Name: keywordName,
			}

			if err := tx.Save(&savedKeyword).Error; err != nil {
				tx.Rollback()
				return 0, err
			}
		}

		savedKeywordResume := &KeywordResume{
			ResumeID:  resume.ID,
			KeywordID: savedKeyword.ID,
		}

		if err := tx.Save(&savedKeywordResume).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	return resume.ID, tx.Commit().Error
}

func GetResumeByUserID(userID string) (*Resume, error) {
	var resume Resume
	db := common.GetDB()

	if err := db.Preload("WorkExperiences.WorkExperienceDetails").
		Preload(clause.Associations).Where(Resume{
		UserID: userID,
	}).First(&resume).Error; err != nil {
		return nil, err
	}

	return &resume, nil
}

func GetKeywordsByResumeID(id uint) []string {
	var keywordResumes []KeywordResume
	db := common.GetDB()

	if err := db.Preload("Keyword").Where(&KeywordResume{
		ResumeID: id,
	}).Find(&keywordResumes).Error; err != nil {
		return make([]string, 0)
	}

	result := make([]string, len(keywordResumes))

	for i, keywordResume := range keywordResumes {
		result[i] = keywordResume.Keyword.Name
	}

	return result
}

func ConvertResumeToWanted(resume Resume, keywords []string) WantedResume {
	wantedResume := &WantedResume{
		Introduction:    resume.Introduction,
		WorkExperiences: make([]WantedResumeWorkExperience, len(resume.WorkExperiences)),
		Educations:      make([]WantedResumeEducation, len(resume.Educations)),
		Certificates:    make([]WantedResumeCertificate, len(resume.Certificates)+len(resume.Projects)+len(resume.Activities)),
		Skills:          keywords,
	}

	for i, workExperience := range resume.WorkExperiences {
		wantedResume.WorkExperiences[i] = WantedResumeWorkExperience{
			CompanyName: workExperience.CompanyName,
			Position:    workExperience.Position,
			StartDate:   workExperience.StartDate,
			EndDate:     workExperience.EndDate,
		}

		wantedResume.WorkExperiences[i].Details = make([]WantedResumeWorkExperienceDetail, len(resume.WorkExperiences[i].WorkExperienceDetails))
		for j, detail := range resume.WorkExperiences[i].WorkExperienceDetails {
			wantedResume.WorkExperiences[i].Details[j] = WantedResumeWorkExperienceDetail{
				Name:      detail.Name,
				StartDate: detail.StartDate,
				EndDate:   detail.EndDate,
				Content:   detail.Content,
			}
		}
	}

	for i, education := range resume.Educations {
		wantedResume.Educations[i] = WantedResumeEducation{
			SchoolName: education.SchoolName,
			MajorName:  education.MajorName,
			StartDate:  education.StartDate,
			EndDate:    education.EndDate,
		}
	}

	wantedResume.Certificates = make([]WantedResumeCertificate, len(resume.Certificates)+len(resume.Projects)+len(resume.Activities))
	idx := 0
	for _, certificate := range resume.Certificates {
		wantedResume.Certificates[idx].Name = certificate.Name
		wantedResume.Certificates[idx].StartDate = certificate.IssuedDate
		wantedResume.Certificates[idx].Content = certificate.IssuedBy
		idx++
	}

	for _, project := range resume.Projects {
		wantedResume.Certificates[idx].Name = project.Name
		wantedResume.Certificates[idx].StartDate = project.StartDate
		wantedResume.Certificates[idx].Content = project.Content
		idx++
	}

	for _, activity := range resume.Activities {
		wantedResume.Certificates[idx].Name = activity.Name
		wantedResume.Certificates[idx].StartDate = activity.StartDate
		wantedResume.Certificates[idx].Content = activity.Content
	}

	return *wantedResume
}

func (r *GetWantedResumeResponse) from(wantedResume WantedResume) *GetWantedResumeResponse {
	r.Introduction = wantedResume.Introduction
	r.Skills = wantedResume.Skills

	r.Educations = make([]GetWantedResumeResponseEducation, len(wantedResume.Educations))
	for i, education := range wantedResume.Educations {
		r.Educations[i] = GetWantedResumeResponseEducation{
			SchoolName: education.SchoolName,
			MajorName:  education.MajorName,
			StartDate:  education.StartDate,
			EndDate:    education.EndDate,
		}
	}

	r.WorkExperiences = make([]GetWantedResumeResponseWorkExperience, len(wantedResume.WorkExperiences))
	for i, workExperience := range wantedResume.WorkExperiences {
		r.WorkExperiences[i].CompanyName = workExperience.CompanyName
		r.WorkExperiences[i].Position = workExperience.CompanyName
		r.WorkExperiences[i].StartDate = workExperience.StartDate
		r.WorkExperiences[i].EndDate = workExperience.EndDate

		r.WorkExperiences[i].Details = make([]GetWantedResumeResponseWorkExperienceDetail, len(workExperience.Details))
		for j, detail := range workExperience.Details {
			r.WorkExperiences[i].Details[j].Name = detail.Name
			r.WorkExperiences[i].Details[j].StartDate = detail.StartDate
			r.WorkExperiences[i].Details[j].EndDate = detail.EndDate
			r.WorkExperiences[i].Details[j].Content = detail.Content
		}
	}

	r.Certificates = make([]GetWantedResumeResponseCertificate, len(wantedResume.Certificates))
	for i, certificates := range wantedResume.Certificates {
		r.Certificates[i].Name = certificates.Name
		r.Certificates[i].Content = certificates.Content
		r.Certificates[i].StartDate = certificates.StartDate
	}

	return r
}
