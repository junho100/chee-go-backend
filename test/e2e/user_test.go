package e2e

import (
	"bytes"
	"chee-go-backend/internal/http/dto"
	"chee-go-backend/internal/http/handler"
	"chee-go-backend/internal/http/router"
	"chee-go-backend/internal/repository"
	"chee-go-backend/internal/service"
	"chee-go-backend/test/util"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserAPI(t *testing.T) {
	// 테스트용 DB 설정
	db, cleanup := util.NewDB()
	defer cleanup()

	// 라우터 및 핸들러 설정
	r := router.NewRouter()
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	handler.NewUserHandler(r, userService)

	t.Run("회원가입 성공", func(t *testing.T) {
		// Given
		signUpRequest := dto.SignUpRequest{
			ID:       util.RandomString(10),
			Email:    util.RandomString(8) + "@test.com",
			Password: "password123!",
		}
		jsonData, _ := json.Marshal(signUpRequest)

		// When
		req := httptest.NewRequest("POST", "/api/users", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// Then
		assert.Equal(t, http.StatusCreated, w.Code)

		var response dto.SignUpResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, signUpRequest.ID, response.ID)
	})

	t.Run("ID 중복 체크", func(t *testing.T) {
		// Given
		userID := util.RandomString(10)
		signUpRequest := dto.SignUpRequest{
			ID:       userID,
			Email:    util.RandomString(8) + "@test.com",
			Password: "password123!",
		}
		jsonData, _ := json.Marshal(signUpRequest)

		// 사용자 생성
		req := httptest.NewRequest("POST", "/api/users", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// When: 중복 체크 요청
		req = httptest.NewRequest("GET", "/api/users/check-id?id="+userID, nil)
		w = httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// Then
		assert.Equal(t, http.StatusOK, w.Code)

		var response dto.CheckIDResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.True(t, response.IsExists)
	})

	t.Run("로그인 성공", func(t *testing.T) {
		// Given
		userID := util.RandomString(10)
		password := "password123!"

		// 사용자 생성
		signUpRequest := dto.SignUpRequest{
			ID:       userID,
			Email:    util.RandomString(8) + "@test.com",
			Password: password,
		}
		jsonData, _ := json.Marshal(signUpRequest)
		req := httptest.NewRequest("POST", "/api/users", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// When: 로그인 요청
		loginRequest := dto.LoginRequest{
			ID:       userID,
			Password: password,
		}
		jsonData, _ = json.Marshal(loginRequest)
		req = httptest.NewRequest("POST", "/api/users/login", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w = httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// Then
		assert.Equal(t, http.StatusCreated, w.Code)

		var loginResponse dto.LoginResponse
		err := json.Unmarshal(w.Body.Bytes(), &loginResponse)
		assert.NoError(t, err)
		assert.NotEmpty(t, loginResponse.Token)
	})

	t.Run("토큰 검증", func(t *testing.T) {
		// Given
		userID := util.RandomString(10)
		password := "password123!"

		// 사용자 생성 및 로그인
		signUpRequest := dto.SignUpRequest{
			ID:       userID,
			Email:    util.RandomString(8) + "@test.com",
			Password: password,
		}
		jsonData, _ := json.Marshal(signUpRequest)
		req := httptest.NewRequest("POST", "/api/users", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		loginRequest := dto.LoginRequest{
			ID:       userID,
			Password: password,
		}
		jsonData, _ = json.Marshal(loginRequest)
		req = httptest.NewRequest("POST", "/api/users/login", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w = httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var loginResponse dto.LoginResponse
		json.Unmarshal(w.Body.Bytes(), &loginResponse)

		// When: 토큰 검증 요청
		checkMeRequest := dto.CheckMeRequest(loginResponse)
		jsonData, _ = json.Marshal(checkMeRequest)
		req = httptest.NewRequest("POST", "/api/users/me", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w = httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// Then
		assert.Equal(t, http.StatusCreated, w.Code)

		var checkMeResponse dto.CheckMeResponse
		err := json.Unmarshal(w.Body.Bytes(), &checkMeResponse)
		assert.NoError(t, err)
		assert.Equal(t, userID, checkMeResponse.UserID)
	})
}
