package entity

type Subject struct {
	ID                uint   `gorm:"primary_key"`
	Name              string `gorm:"column:name"`
	LecturerName      string `gorm:"column:lecturer_name"`
	ThumbnailURL      string `gorm:"column:thumbnail_url;type:text"`
	YoutubePlayListID string `gorm:"column:youtube_playlist_id"`
	SubjectName       string `gorm:"subject_name"`
	IsForSchool       bool   `gorm:"column:is_for_school;default:false"`
	TargetGrade       uint   `gorm:"column:target_grade"`
	Lectures          []Lecture
}

type Lecture struct {
	ID          uint   `gorm:"primay_key"`
	Title       string `gorm:"column:title"`
	Description string `gorm:"column:description;type:text"`
	YoutubeID   string `gorm:"column:youtube_id"`
	SubjectID   uint
}
