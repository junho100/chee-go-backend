package users

import (
	"chee-go-backend/common"
	"errors"
	"fmt"
	"strings"
	"time"

	"os"

	"github.com/golang-jwt/jwt/v5"
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

type LoginResponse struct {
	Token string `json:"token"`
}

type LoginRequest struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

type User struct {
	ID             string `gorm:"type:varchar(255);primaryKey"`
	Email          string `gorm:"column:email"`
	HashedPassword string `gorm:"column:hashed_password;not null"`
}

type CheckMeRequest struct {
	Token string `json:"token"`
}

type CheckMeResponse struct {
	UserID string `json:"user_id"`
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

func GetUserByID(id string) (*User, error) {
	db := common.GetDB()
	var user User
	if err := db.Where(User{
		ID: id,
	}).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func CheckPassword(password string, hashedPassword string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(hashedPassword)
	fmt.Println(bytePassword)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

func CreateToken(id string) (string, error) {
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

func ExtractToken(rawTokenHeaderValue string) (string, error) {
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

func GetUserIDFromToken(tokenValue string) (string, error) {
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
