package entity

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
	Name      string    `gorm:"column:name"`
	ResumeID  uint
	Summary   string `gorm:"column:summary"`
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
