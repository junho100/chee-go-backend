package entity

type User struct {
	ID             string `gorm:"type:varchar(255);primaryKey"`
	Email          string `gorm:"column:email"`
	HashedPassword string `gorm:"column:hashed_password;not null"`
}
