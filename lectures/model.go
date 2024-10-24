package lectures

type Subject struct {
	ID           uint   `gorm:"primary_key"`
	Name         string `gorm:"column:name"`
	LecturerName string `gorm:"column:lecturer_name"`
	ThumbnailURL string `gorm:"column:thumbnail_url;type:text"`
	Lectures     []Lecture
}

type Lecture struct {
	ID          uint   `gorm:"primay_key"`
	Title       string `gorm:"column:title"`
	Description string `gorm:"column:description;type:text"`
	YoutubeID   string `gorm:"column:youtube_id"`
	SubjectID   uint
}
