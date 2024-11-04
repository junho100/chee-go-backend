package repository

import (
	"chee-go-backend/internal/domain/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	StartTransaction() (*gorm.DB, error)
	CreateUser(tx *gorm.DB, user *entity.User) error
	FindUserById(id string) (*entity.User, error)
}
