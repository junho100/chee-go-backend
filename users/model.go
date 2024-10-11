package users

import (
	"chee-go-backend/common"

	"golang.org/x/crypto/bcrypt"
)

type SignUpRequest struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	ID string
}

type CreateUserDto struct {
	ID       string
	Email    string
	Password string
}

type CheckIDResponse struct {
	IsExists bool
}

type User struct {
	ID             string `gorm:"primary_key"`
	Email          string `gorm:"column:email"`
	HashedPassword string `gorm:"column:hashed_password;not null"`
}

func CreateUser(dto *CreateUserDto) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	user := &User{
		ID:             dto.ID,
		Email:          dto.Email,
		HashedPassword: string(hashedPassword),
	}

	db := common.GetDB()
	tx := db.Begin()

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func CheckUserByID(id string) bool {
	db := common.GetDB()
	var user User
	if err := db.Where(User{
		ID: id,
	}).First(&user).Error; err != nil {
		return false
	}

	return true
}
