package e2e

import (
	"bytes"
	"chee-go-backend/internal/http/dto"
	"chee-go-backend/internal/http/handler"
	"chee-go-backend/internal/http/router"
	"chee-go-backend/internal/repository"
	"chee-go-backend/internal/service"
	"chee-go-backend/test/mock"
	"chee-go-backend/test/util"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestDiscordNotificationAPI(t *testing.T) {
	// 테스트용 DB 설정
	db, cleanup := util.NewDB()
	defer cleanup()

	// Mock Discord 클라이언트 설정
	mockDiscordClient := &mock.MockDiscordClient{
		ValidateClientIDFunc: func(clientID string) bool {
			return clientID == "valid_client_id"
		},
		SendMessageFunc: func(clientID string, message string) error {
			return nil
		},
	}

	// 라우터 및 핸들러 설정
	r := router.NewRouter()
	userRepository := repository.NewUserRepository(db)
	notificationRepository := repository.NewNotificationRepository(db)
	userService := service.NewUserService(userRepository)
	notificationService := service.NewNotificationService(notificationRepository, nil)
	handler.NewUserHandler(r, userService)
	handler.NewNotificationHandler(r, nil, mockDiscordClient, userService, notificationService)

	// 테스트용 사용자 생성 및 토큰 발급
	userID := util.RandomString(10)
	password := "password123!"
	token := setupTestUser(t, r, userID, password)

	t.Run("Discord Client ID 검증 성공", func(t *testing.T) {
		validateRequest := dto.ValidateDiscordClientIDRequest{
			ClientID: "valid_client_id",
		}
		jsonData, _ := json.Marshal(validateRequest)

		req := httptest.NewRequest("POST", "/api/notifications/validate-discord-client-id", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		var response dto.ValidateDiscordClientIDResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.True(t, response.IsValid)
	})

	t.Run("Discord Client ID 검증 실패", func(t *testing.T) {
		validateRequest := dto.ValidateDiscordClientIDRequest{
			ClientID: "invalid_client_id",
		}
		jsonData, _ := json.Marshal(validateRequest)

		req := httptest.NewRequest("POST", "/api/notifications/validate-discord-client-id", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		var response dto.ValidateDiscordClientIDResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.False(t, response.IsValid)
	})

	t.Run("알림 설정 생성 성공", func(t *testing.T) {
		createRequest := dto.CreateNotificationConfigRequest{
			DiscordClientID: "valid_client_id",
			Keywords:        []string{"공지", "시험"},
		}
		jsonData, _ := json.Marshal(createRequest)

		req := httptest.NewRequest("POST", "/api/notifications/config", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		var response dto.CreateNotificationConfigResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.NotZero(t, response.ConfigID)
	})

	t.Run("알림 설정 조회 성공", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/notifications/config", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response dto.GetNotificationConfigResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "valid_client_id", response.DiscordClientID)
		assert.Equal(t, 2, len(response.Keywords))
	})

	t.Run("인증되지 않은 사용자의 알림 설정 조회 실패", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/notifications/config", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusForbidden, w.Code)
	})
}

// 테스트용 사용자 생성 및 토큰 발급 헬퍼 함수
func setupTestUser(t *testing.T, r *gin.Engine, userID string, password string) string {
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
	return loginResponse.Token
}
