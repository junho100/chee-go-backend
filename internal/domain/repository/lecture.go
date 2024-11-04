package repository

import (
	"chee-go-backend/internal/domain/entity"

	"gorm.io/gorm"
)

type LectureRepository interface {
	StartTransaction() (*gorm.DB, error)
	FindSubjectByYoutubePlayListId(playListId string) (*entity.Subject, error)
	CreateSubject(tx *gorm.DB, subject *entity.Subject) error
	CreateLecture(tx *gorm.DB, lecture *entity.Lecture) error
	FindAllSubjects() []entity.Subject
	FindSubjectById(subjectId uint) (*entity.Subject, error)
}
