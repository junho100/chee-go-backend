package service

import (
	"chee-go-backend/internal/domain/entity"

	"google.golang.org/api/youtube/v3"
)

type LectureService interface {
	CreateSubjectWithLectures(playList *youtube.PlaylistListResponse, playListItems []*youtube.PlaylistItem) error
	GetAllSubjects() []entity.Subject
	GetSubjectByID(subjectID uint) (*entity.Subject, error)
}
