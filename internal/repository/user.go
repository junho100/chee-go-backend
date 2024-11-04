package repository

import (
	"chee-go-backend/internal/domain/entity"
	"chee-go-backend/internal/domain/repository"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(tx *gorm.DB, user *entity.User) error {
	if err := tx.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) FindUserById(id string) (*entity.User, error) {
	var user entity.User

	if err := r.db.Where(entity.User{
		ID: id,
	}).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) StartTransaction() (*gorm.DB, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	return tx, nil
}
