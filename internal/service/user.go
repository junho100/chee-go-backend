package service

import (
	"chee-go-backend/internal/domain/repository"
	"chee-go-backend/internal/domain/service"
)

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) service.UserService {
	return &userService{
		userRepository: userRepository,
	}
}
