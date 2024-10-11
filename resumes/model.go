package resumes

import (
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
	ID          uint   `gorm:"primary_key"`
	CompanyName string `gorm:"column:company_name"`
	Department  string `gorm:"column:department"`
	Position    string `gorm:"column:position"`
	Job         string `gorm:"column:job"`
	ResumeID    uint
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
	GithubURL       string
	BlogURL         string
	Educations      []RegisterResumeRequestReducation
	Projects        []RegisterResumeRequestProject
	Activities      []RegisterResumeRequestActivity
	Certificates    []RegisterResumeRequestCertificate
	WorkExperiences []RegisterResumeRequestWorkExperience
	Keywords        []string
}

type RegisterResumeRequestReducation struct {
	SchoolName string
	MajorName  string
	StartDate  time.Time
	EndDate    time.Time
}

type RegisterResumeRequestProject struct {
	StartDate time.Time
	EndDate   time.Time
	Content   string
	GithubURL string
}

type RegisterResumeRequestActivity struct {
	Name    string
	Content string
}

type RegisterResumeRequestCertificate struct {
	Name       string
	IssuedBy   string
	IssuedDate time.Time
}

type RegisterResumeRequestWorkExperience struct {
	CompanyName string
	Department  string
	Position    string
	Job         string
	Details     []RegisterResumeRequestWorkExperienceDetail
}

type RegisterResumeRequestWorkExperienceDetail struct {
	Name      string
	StartDate time.Time
	EndDate   time.Time
	Content   string
}
