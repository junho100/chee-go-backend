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
	"time"

	"github.com/stretchr/testify/assert"
)

func TestResumeAPI(t *testing.T) {
	// 테스트용 DB 설정
	db, cleanup := util.NewDB()
	defer cleanup()

	// 라우터 및 핸들러 설정
	r := router.NewRouter()
	userRepository := repository.NewUserRepository(db)
	resumeRepository := repository.NewResumeRepository(db)
	userService := service.NewUserService(userRepository)
	resumeService := service.NewResumeService(resumeRepository)
	handler.NewUserHandler(r, userService)
	handler.NewResumeHandler(r, resumeService, userService)

	// 테스트용 사용자 생성 및 토큰 발급
	userID := util.RandomString(10)
	password := "password123!"
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
	token := loginResponse.Token

	t.Run("이력서 생성 성공", func(t *testing.T) {
		// Given
		now := time.Now()
		registerResumeRequest := dto.RegisterResumeRequest{
			Introduction: "안녕하세요",
			GithubURL:    "https://github.com/test",
			BlogURL:      "https://blog.test.com",
			Educations: []dto.RegisterResumeRequestEducation{
				{
					SchoolName: "테스트대학교",
					MajorName:  "컴퓨터공학과",
					StartDate:  now.AddDate(-4, 0, 0),
					EndDate:    now,
				},
			},
			Projects: []dto.RegisterResumeRequestProject{
				{
					Name:      "테스트 프로젝트",
					StartDate: now.AddDate(-1, 0, 0),
					EndDate:   now,
					Content:   "테스트 프로젝트입니다",
					GithubURL: "https://github.com/test/project",
					Summary:   "테스트 프로젝트 요약",
				},
			},
			Activities: []dto.RegisterResumeRequestActivity{
				{
					Name:      "테스트 활동",
					Content:   "테스트 활동입니다",
					StartDate: now.AddDate(-2, 0, 0),
					EndDate:   now,
				},
			},
			Certificates: []dto.RegisterResumeRequestCertificate{
				{
					Name:       "정보처리기사",
					IssuedBy:   "한국산업인력공단",
					IssuedDate: now.AddDate(-1, 0, 0),
				},
			},
			WorkExperiences: []dto.RegisterResumeRequestWorkExperience{
				{
					CompanyName: "테스트회사",
					Department:  "개발팀",
					Position:    "백엔드 개발자",
					Job:         "백엔드 개발",
					StartDate:   now.AddDate(-2, 0, 0),
					EndDate:     now,
					Details: []dto.RegisterResumeRequestWorkExperienceDetail{
						{
							Name:      "테스트 프로젝트",
							StartDate: now.AddDate(-1, 0, 0),
							EndDate:   now,
							Content:   "테스트 프로젝트를 진행했습니다",
						},
					},
				},
			},
			Keywords: []string{"Go", "MySQL", "Docker"},
		}
		jsonData, _ := json.Marshal(registerResumeRequest)

		// When
		req := httptest.NewRequest("POST", "/api/resumes", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// Then
		assert.Equal(t, http.StatusCreated, w.Code)

		var response dto.RegisterResumeResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.NotZero(t, response.ResumeID)
	})

	t.Run("이력서 조회 성공", func(t *testing.T) {
		// When
		req := httptest.NewRequest("GET", "/api/resumes", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// Then
		assert.Equal(t, http.StatusOK, w.Code)

		var response dto.GetResumeResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)

		assert.NoError(t, err)
		assert.Equal(t, "안녕하세요", response.Introduction)
		assert.Equal(t, "https://github.com/test", response.GithubURL)
		assert.Equal(t, "https://blog.test.com", response.BlogURL)
		assert.Len(t, response.Educations, 1)
		assert.Len(t, response.Projects, 1)
		assert.Len(t, response.Activities, 1)
		assert.Len(t, response.Certificates, 1)
		assert.Len(t, response.WorkExperiences, 1)
		assert.Len(t, response.Keywords, 3)
	})

	t.Run("원티드 이력서 변환 조회 성공", func(t *testing.T) {
		// When
		req := httptest.NewRequest("GET", "/api/resumes/wanted", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// Then
		assert.Equal(t, http.StatusOK, w.Code)

		var response dto.GetWantedResumeResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "안녕하세요", response.Introduction)
		assert.Len(t, response.WorkExperiences, 1)
		assert.Len(t, response.Educations, 1)
		assert.Len(t, response.Skills, 3)
		assert.Len(t, response.Certificates, 3) // 프로젝트, 자격증, 활동이 모두 포함됨
	})

	t.Run("프로그래머스 이력서 변환 조회 성공", func(t *testing.T) {
		// When
		req := httptest.NewRequest("GET", "/api/resumes/programmers", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// Then
		assert.Equal(t, http.StatusOK, w.Code)

		var response dto.GetProgrammersResumeResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Len(t, response.WorkExperiences, 1)
		assert.Len(t, response.Educations, 1)
		assert.Len(t, response.Projects, 1)
		assert.Len(t, response.Certificates, 1)
		assert.Len(t, response.Activities, 1)
	})

	t.Run("링크드인 이력서 변환 조회 성공", func(t *testing.T) {
		// When
		req := httptest.NewRequest("GET", "/api/resumes/linkedin", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// Then
		assert.Equal(t, http.StatusOK, w.Code)

		var response dto.GetLinkedinResumeResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "안녕하세요", response.Introduction)
		assert.Len(t, response.WorkExperiences, 1)
		assert.Len(t, response.Educations, 1)
		assert.Len(t, response.Certificates, 1)
		assert.Len(t, response.Projects, 1)
		assert.Len(t, response.Skills, 3)
	})
}
