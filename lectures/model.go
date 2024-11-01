// package lectures

// import (
// 	"errors"

// 	"google.golang.org/api/youtube/v3"
// 	"gorm.io/gorm/clause"
// )

// type RegisterLectureRequest struct {
// 	PlayListID string `json:"playlist_id"`
// }

// type RegisterLectureResponse struct {
// 	IsSuccess bool `json:"is_success"`
// }

// type GetLecturesResponse struct {
// 	Subjects []GetLecturesResponseSubject `json:"subjects"`
// }

// type GetLecturesResponseSubject struct {
// 	ID           uint   `json:"id"`
// 	Title        string `json:"title"`
// 	Description  string `json:"description"`
// 	ThumbnailURL string `json:"thumbnailUrl"`
// 	Instructor   string `json:"instructor"`
// }

// type GetLectureResponse struct {
// 	ID     uint                      `json:"id"`
// 	Title  string                    `json:"title"`
// 	Videos []GetLectureResponseVideo `json:"videos"`
// }

// type GetLectureResponseVideo struct {
// 	ID          uint   `json:"id"`
// 	Title       string `json:"title"`
// 	Description string `json:"description"`
// 	YoutubeID   string `json:"youtubeId"`
// }

// func CreateSubjectWithLectures(playList *youtube.PlaylistListResponse, playListItems []*youtube.PlaylistItem) error {
// 	var subject *Subject

// 	db := DB
// 	tx := db.Begin()

// 	if err := db.Where(&Subject{
// 		YoutubePlayListID: playList.Items[0].Id,
// 	}).First(&subject).Error; err == nil {
// 		return errors.New("subject already exists")
// 	}

// 	subject = &Subject{
// 		Name:              playList.Items[0].Snippet.Title,
// 		LecturerName:      playList.Items[0].Snippet.ChannelTitle,
// 		ThumbnailURL:      playList.Items[0].Snippet.Thumbnails.Medium.Url,
// 		YoutubePlayListID: playList.Items[0].Id,
// 	}

// 	if err := tx.Create(&subject).Error; err != nil {
// 		tx.Rollback()
// 		return err
// 	}

// 	for _, playListItem := range playListItems {
// 		lecture := &Lecture{
// 			Title:       playListItem.Snippet.Title,
// 			Description: playListItem.Snippet.Description,
// 			YoutubeID:   playListItem.Snippet.ResourceId.VideoId,
// 			SubjectID:   subject.ID,
// 		}

// 		if err := tx.Save(&lecture).Error; err != nil {
// 			tx.Rollback()
// 			return err
// 		}
// 	}

// 	return tx.Commit().Error
// }

// func GetAllSubjects() []Subject {
// 	var subjects []Subject

// 	db := DB

// 	if err := db.Find(&subjects).Error; err != nil {
// 		return make([]Subject, 0)
// 	}

// 	return subjects
// }

// func (c *GetLectureResponse) from(subject Subject) {
// 	c.ID = subject.ID
// 	c.Title = subject.SubjectName
// 	c.Videos = make([]GetLectureResponseVideo, len(subject.Lectures))

// 	for i, video := range subject.Lectures {
// 		c.Videos[i] = GetLectureResponseVideo{
// 			ID:          video.ID,
// 			Title:       video.Title,
// 			Description: video.Description,
// 			YoutubeID:   video.YoutubeID,
// 		}
// 	}
// }

// func GetSubjectByID(subjectID uint) (*Subject, error) {
// 	db := DB
// 	var subject Subject

// 	if err := db.Preload(clause.Associations).Where(&Subject{
// 		ID: subjectID,
// 	}).First(&subject).Error; err != nil {
// 		return nil, err
// 	}

// 	return &subject, nil
// }
