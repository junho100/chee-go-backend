package service

import (
	"chee-go-backend/internal/domain/entity"
	"chee-go-backend/internal/domain/repository"
	"chee-go-backend/internal/domain/service"
	"chee-go-backend/internal/http/dto"
	"errors"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) service.UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) CreateUser(createUserDto *dto.CreateUserDto) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(createUserDto.Password), bcrypt.DefaultCost)
	user := &entity.User{
		ID:             createUserDto.ID,
		Email:          createUserDto.Email,
		HashedPassword: string(hashedPassword),
	}

	tx, err := s.userRepository.StartTransaction()
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := s.userRepository.CreateUser(tx, user); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (s *userService) CheckPassword(password string, hashedPassword string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(hashedPassword)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

func (s *userService) CheckUserByID(id string) bool {
	if _, err := s.userRepository.FindUserById(id); err != nil {
		return false
	}

	return true
}

func (s *userService) CreateToken(id string) (string, error) {
	var err error
	secret := os.Getenv("JWT_SECRET")
	tokenClaim := jwt.MapClaims{}
	tokenClaim["authorized"] = true
	tokenClaim["user_id"] = id
	tokenClaim["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaim)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s *userService) GetUserByID(id string) (*entity.User, error) {
	user, err := s.userRepository.FindUserById(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetUserIDFromToken(tokenValue string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	token, err := jwt.Parse(tokenValue, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if userID, ok := claims["user_id"].(string); ok {
			return userID, nil
		}
	}
	return "", errors.New("failed to extract token")
}

func (s *userService) ExtractToken(rawTokenHeaderValue string) (string, error) {
	tokenStringArray := strings.Split(rawTokenHeaderValue, " ")
	if len(tokenStringArray) <= 1 {
		return "", errors.New("no token")
	}

	token := tokenStringArray[1]

	if token == "" {
		return "", errors.New("no token")
	}

	return token, nil
}
