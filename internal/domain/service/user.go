package service

import (
	"chee-go-backend/internal/domain/entity"
	"chee-go-backend/internal/http/dto"
)

type UserService interface {
	CreateUser(createUserDto *dto.CreateUserDto) error
	CheckUserByID(id string) bool
	GetUserByID(id string) (entity.User, error)
	CheckPassword(password string, hashedPassword string) error
	CreateToken(id string) (string, error)
	GetUserIDFromToken(token string) (string, error)
}
