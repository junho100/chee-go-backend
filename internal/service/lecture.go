package service

import (
	"chee-go-backend/internal/domain/entity"
	"chee-go-backend/internal/domain/repository"
	"chee-go-backend/internal/domain/service"
	"errors"

	"google.golang.org/api/youtube/v3"
)

type lectureService struct {
	lectureRepository repository.LectureRepository
}

func (s *lectureService) CreateSubjectWithLectures(playList *youtube.PlaylistListResponse, playListItems []*youtube.PlaylistItem) error {
	tx, err := s.lectureRepository.StartTransaction()
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if _, err := s.lectureRepository.FindSubjectByYoutubePlayListId(playList.Items[0].Id); err == nil {
		return errors.New("subject already exists")
	}

	subject := &entity.Subject{
		Name:              playList.Items[0].Snippet.Title,
		LecturerName:      playList.Items[0].Snippet.ChannelTitle,
		ThumbnailURL:      playList.Items[0].Snippet.Thumbnails.Medium.Url,
		YoutubePlayListID: playList.Items[0].Id,
	}

	if err := s.lectureRepository.CreateSubject(tx, subject); err != nil {
		tx.Rollback()
		return err
	}

	for _, playListItem := range playListItems {
		lecture := &entity.Lecture{
			Title:       playListItem.Snippet.Title,
			Description: playListItem.Snippet.Description,
			YoutubeID:   playListItem.Snippet.ResourceId.VideoId,
			SubjectID:   subject.ID,
		}

		if err := s.lectureRepository.CreateLecture(tx, lecture); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (s *lectureService) GetAllSubjects() []entity.Subject {
	subjects := s.lectureRepository.FindAllSubjects()

	return subjects
}

func (s *lectureService) GetSubjectByID(subjectID uint) (*entity.Subject, error) {
	subject, err := s.lectureRepository.FindSubjectById(subjectID)

	if err != nil {
		return nil, err
	}

	return subject, nil
}

func NewLectureService(lectureRepository repository.LectureRepository) service.LectureService {
	return &lectureService{
		lectureRepository: lectureRepository,
	}
}
