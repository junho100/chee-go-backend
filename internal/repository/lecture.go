package repository

import (
	"chee-go-backend/internal/domain/entity"
	"chee-go-backend/internal/domain/repository"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type lectureRepository struct {
	db *gorm.DB
}

func (r *lectureRepository) CreateLecture(tx *gorm.DB, lecture *entity.Lecture) error {
	if err := tx.Save(lecture).Error; err != nil {
		return err
	}

	return nil
}

func (r *lectureRepository) CreateSubject(tx *gorm.DB, subject *entity.Subject) error {
	if err := tx.Save(subject).Error; err != nil {
		return err
	}

	return nil
}

func (r *lectureRepository) FindAllSubjects() []entity.Subject {
	var subjects []entity.Subject

	if err := r.db.Find(&subjects).Error; err != nil {
		return make([]entity.Subject, 0)
	}

	return subjects
}

func (r *lectureRepository) FindSubjectById(subjectId uint) (*entity.Subject, error) {
	var subject entity.Subject

	if err := r.db.Preload(clause.Associations).Where(&entity.Subject{
		ID: subjectId,
	}).First(&subject).Error; err != nil {
		return nil, err
	}

	return &subject, nil
}

func (r *lectureRepository) FindSubjectByYoutubePlayListId(playListId string) (*entity.Subject, error) {
	var subject entity.Subject

	if err := r.db.Where(&entity.Subject{
		YoutubePlayListID: playListId,
	}).First(&subject).Error; err != nil {
		return nil, err
	}

	return &subject, nil
}

func (r *lectureRepository) StartTransaction() (*gorm.DB, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	return tx, nil
}

func NewLectureRepository(db *gorm.DB) repository.LectureRepository {
	return &lectureRepository{
		db: db,
	}
}
