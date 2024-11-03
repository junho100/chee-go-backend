package dto

import (
	"chee-go-backend/internal/domain/entity"
	"time"
)

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
	Summary   string `json:"summary"`
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
	Summary   string    `json:"summary"`
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

type GetProgrammersResumeResponse struct {
	WorkExperiences []GetProgrammersResumeResponseWorkExperience `json:"work_experiences"`
	Educations      []GetProgrammersResumeResponseEducation      `json:"educations"`
	Projects        []GetProgrammersResumeResponseProject        `json:"projects"`
	Certificates    []GetProgrammersResumeResponseCertificates   `json:"certificates"`
	Activities      []GetProgrammersResumeResponseActivity       `json:"activities"`
}

type GetProgrammersResumeResponseWorkExperience struct {
	CompanyName string                                             `json:"company_name"`
	Position    string                                             `json:"position"`
	StartDate   time.Time                                          `json:"start_date"`
	Details     []GetProgrammersResumeResponseWorkExperienceDetail `json:"details"`
}

type GetProgrammersResumeResponseWorkExperienceDetail struct {
	Name      string    `json:"name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Content   string    `json:"content"`
}

type GetProgrammersResumeResponseEducation struct {
	SchoolName string    `json:"school_name"`
	MajorName  string    `json:"major_name"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
}

type GetProgrammersResumeResponseProject struct {
	Name      string    `json:"name"`
	StartDate time.Time `json:"start_date"`
	Summary   string    `json:"summary"`
	Content   string    `json:"content"`
	GithubURL string    `json:"github_url"`
}

type GetProgrammersResumeResponseCertificates struct {
	Name       string    `json:"name"`
	IssuedBy   string    `json:"issued_by"`
	IssuedDate time.Time `json:"issued_date"`
}

type GetProgrammersResumeResponseActivity struct {
	Name      string    `json:"name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Content   string    `json:"content"`
}

type ProgrammersResume struct {
	WorkExperiences []ProgrammersResumeWorkExperience
	Educations      []ProgrammersResumeEducation
	Projects        []ProgrammersResumeProject
	Certificates    []ProgrammersResumeCertificates
	Activities      []ProgrammersResumeActivity
}

type ProgrammersResumeWorkExperience struct {
	CompanyName string
	Position    string
	StartDate   time.Time
	Details     []ProgrammersResumeWorkExperienceDetail
}

type ProgrammersResumeWorkExperienceDetail struct {
	Name      string
	StartDate time.Time
	EndDate   time.Time
	Content   string
}

type ProgrammersResumeEducation struct {
	SchoolName string
	MajorName  string
	StartDate  time.Time
	EndDate    time.Time
}

type ProgrammersResumeProject struct {
	Name      string
	StartDate time.Time
	Summary   string
	Content   string
	GithubURL string
}

type ProgrammersResumeCertificates struct {
	Name       string
	IssuedBy   string
	IssuedDate time.Time
}

type ProgrammersResumeActivity struct {
	Name      string
	StartDate time.Time
	EndDate   time.Time
	Content   string
}

type LinkedinResume struct {
	Introduction    string
	WorkExperiences []LinkedinResumeWorkExperience
	Educations      []LinkedinResumeEducation
	Certificates    []LinkedinResumeCertificate
	Projects        []LinkedinResumeProject
	Skills          []string
}

type LinkedinResumeWorkExperience struct {
	Position    string
	CompanyName string
	StartDate   time.Time
	EndDate     time.Time
	Content     string
}

type LinkedinResumeEducation struct {
	SchoolName string
	MajorName  string
	StartDate  time.Time
	EndDate    time.Time
}

type LinkedinResumeCertificate struct {
	Name       string
	IssuedBy   string
	IssuedDate time.Time
}

type LinkedinResumeProject struct {
	Name      string
	StartDate time.Time
	EndDate   time.Time
	Content   string
}

type GetLinkedinResumeResponse struct {
	Introduction    string                                    `json:"introduction"`
	WorkExperiences []GetLinkedinResumeResponseWorkExperience `json:"work_experiences"`
	Educations      []GetLinkedinResumeResponseEducation      `json:"educations"`
	Certificates    []GetLinkedinResumeResponseCertificate    `json:"certificates"`
	Projects        []GetLinkedinResumeResponseProject        `json:"projects"`
	Skills          []string                                  `json:"skills"`
}

type GetLinkedinResumeResponseWorkExperience struct {
	Position    string    `json:"position"`
	CompanyName string    `json:"company_name"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Content     string    `json:"content"`
}

type GetLinkedinResumeResponseEducation struct {
	SchoolName string    `json:"school_name"`
	MajorName  string    `json:"major_name"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
}

type GetLinkedinResumeResponseCertificate struct {
	Name       string    `json:"name"`
	IssuedBy   string    `json:"issued_by"`
	IssuedDate time.Time `json:"issued_date"`
}

type GetLinkedinResumeResponseProject struct {
	Name      string    `json:"name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Content   string    `json:"content"`
}

func (r *GetResumeResponse) From(resume entity.Resume, keywords []string) *GetResumeResponse {
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
			Summary:   project.Summary,
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

func (r *GetWantedResumeResponse) From(wantedResume WantedResume) *GetWantedResumeResponse {
	r.Introduction = wantedResume.Introduction
	r.Skills = wantedResume.Skills

	r.Educations = make([]GetWantedResumeResponseEducation, len(wantedResume.Educations))
	for i, education := range wantedResume.Educations {
		r.Educations[i] = GetWantedResumeResponseEducation(education)
	}

	r.WorkExperiences = make([]GetWantedResumeResponseWorkExperience, len(wantedResume.WorkExperiences))
	for i, workExperience := range wantedResume.WorkExperiences {
		r.WorkExperiences[i].CompanyName = workExperience.CompanyName
		r.WorkExperiences[i].Position = workExperience.Position
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

func (r *GetProgrammersResumeResponse) From(resume ProgrammersResume) *GetProgrammersResumeResponse {
	r.WorkExperiences = make([]GetProgrammersResumeResponseWorkExperience, len(resume.WorkExperiences))
	for i, workExperience := range resume.WorkExperiences {
		r.WorkExperiences[i].CompanyName = workExperience.CompanyName
		r.WorkExperiences[i].Position = workExperience.Position
		r.WorkExperiences[i].StartDate = workExperience.StartDate

		r.WorkExperiences[i].Details = make([]GetProgrammersResumeResponseWorkExperienceDetail, len(workExperience.Details))
		for j, detail := range workExperience.Details {
			r.WorkExperiences[i].Details[j].Name = detail.Name
			r.WorkExperiences[i].Details[j].StartDate = detail.StartDate
			r.WorkExperiences[i].Details[j].EndDate = detail.EndDate
			r.WorkExperiences[i].Details[j].Content = detail.Content
		}
	}

	r.Educations = make([]GetProgrammersResumeResponseEducation, len(resume.Educations))
	for i, education := range resume.Educations {
		r.Educations[i].SchoolName = education.SchoolName
		r.Educations[i].MajorName = education.MajorName
		r.Educations[i].StartDate = education.StartDate
		r.Educations[i].EndDate = education.EndDate
	}

	r.Projects = make([]GetProgrammersResumeResponseProject, len(resume.Projects))
	for i, project := range resume.Projects {
		r.Projects[i].Name = project.Name
		r.Projects[i].StartDate = project.StartDate
		r.Projects[i].Summary = project.Summary
		r.Projects[i].Content = project.Content
		r.Projects[i].GithubURL = project.GithubURL
	}

	r.Certificates = make([]GetProgrammersResumeResponseCertificates, len(resume.Certificates))
	for i, certificate := range resume.Certificates {
		r.Certificates[i].Name = certificate.Name
		r.Certificates[i].IssuedBy = certificate.IssuedBy
		r.Certificates[i].IssuedDate = certificate.IssuedDate
	}

	r.Activities = make([]GetProgrammersResumeResponseActivity, len(resume.Activities))
	for i, activity := range resume.Activities {
		r.Activities[i].Name = activity.Name
		r.Activities[i].StartDate = activity.StartDate
		r.Activities[i].EndDate = activity.EndDate
		r.Activities[i].Content = activity.Content
	}

	return r
}

func (r *GetLinkedinResumeResponse) From(resume LinkedinResume) *GetLinkedinResumeResponse {
	r.Introduction = resume.Introduction
	r.Skills = resume.Skills

	r.WorkExperiences = make([]GetLinkedinResumeResponseWorkExperience, len(resume.WorkExperiences))
	for i, workExperience := range resume.WorkExperiences {
		r.WorkExperiences[i] = GetLinkedinResumeResponseWorkExperience(workExperience)
	}

	r.Educations = make([]GetLinkedinResumeResponseEducation, len(resume.Educations))
	for i, education := range resume.Educations {
		r.Educations[i] = GetLinkedinResumeResponseEducation(education)
	}

	r.Certificates = make([]GetLinkedinResumeResponseCertificate, len(resume.Certificates))
	for i, certificate := range resume.Certificates {
		r.Certificates[i] = GetLinkedinResumeResponseCertificate(certificate)
	}

	r.Projects = make([]GetLinkedinResumeResponseProject, len(resume.Projects))
	for i, project := range resume.Projects {
		r.Projects[i] = GetLinkedinResumeResponseProject(project)
	}

	return r
}
