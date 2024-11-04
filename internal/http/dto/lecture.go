package dto

import "chee-go-backend/internal/domain/entity"

type RegisterLectureRequest struct {
	PlayListID string `json:"playlist_id"`
}

type RegisterLectureResponse struct {
	IsSuccess bool `json:"is_success"`
}

type GetLecturesResponse struct {
	Subjects []GetLecturesResponseSubject `json:"subjects"`
}

type GetLecturesResponseSubject struct {
	ID           uint   `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	ThumbnailURL string `json:"thumbnailUrl"`
	Instructor   string `json:"instructor"`
}

type GetLectureResponse struct {
	ID     uint                      `json:"id"`
	Title  string                    `json:"title"`
	Videos []GetLectureResponseVideo `json:"videos"`
}

type GetLectureResponseVideo struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	YoutubeID   string `json:"youtubeId"`
}

func (c *GetLectureResponse) From(subject entity.Subject) {
	c.ID = subject.ID
	c.Title = subject.SubjectName
	c.Videos = make([]GetLectureResponseVideo, len(subject.Lectures))

	for i, video := range subject.Lectures {
		c.Videos[i] = GetLectureResponseVideo{
			ID:          video.ID,
			Title:       video.Title,
			Description: video.Description,
			YoutubeID:   video.YoutubeID,
		}
	}
}
