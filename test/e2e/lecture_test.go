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
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/api/youtube/v3"
)

func TestLectureAPI(t *testing.T) {
	// 테스트용 DB 설정
	db, cleanup := util.NewDB()
	defer cleanup()

	// Mock Youtube Client 설정
	mockYoutubeClient := &mock.MockYoutubeClient{
		GetPlayListByIDFunc: func(playlistId string) (*youtube.PlaylistListResponse, error) {
			return &youtube.PlaylistListResponse{
				Items: []*youtube.Playlist{
					{
						Id: "test_playlist_id",
						Snippet: &youtube.PlaylistSnippet{
							Title:        "테스트 강의",
							ChannelTitle: "테스트 강사",
							Description:  "테스트 강의 설명",
							Thumbnails: &youtube.ThumbnailDetails{
								Medium: &youtube.Thumbnail{
									Url: "https://test.com/thumbnail.jpg",
								},
							},
						},
					},
				},
			}, nil
		},
		ListVideosByPlayListIDFunc: func(playListID string) ([]*youtube.PlaylistItem, error) {
			return []*youtube.PlaylistItem{
				{
					Snippet: &youtube.PlaylistItemSnippet{
						Title:       "테스트 비디오 1",
						Description: "테스트 비디오 설명 1",
						ResourceId: &youtube.ResourceId{
							VideoId: "video1",
						},
					},
				},
				{
					Snippet: &youtube.PlaylistItemSnippet{
						Title:       "테스트 비디오 2",
						Description: "테스트 비디오 설명 2",
						ResourceId: &youtube.ResourceId{
							VideoId: "video2",
						},
					},
				},
			}, nil
		},
	}

	// 라우터 및 핸들러 설정
	r := router.NewRouter()
	lectureRepository := repository.NewLectureRepository(db)
	lectureService := service.NewLectureService(lectureRepository)
	handler.NewLectureHandler(r, lectureService, mockYoutubeClient)

	t.Run("강의 등록 성공", func(t *testing.T) {
		// Given
		registerLectureRequest := dto.RegisterLectureRequest{
			PlayListID: "test_playlist_id",
		}
		jsonData, _ := json.Marshal(registerLectureRequest)

		// When
		req := httptest.NewRequest("POST", "/api/lectures", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// Then
		assert.Equal(t, http.StatusCreated, w.Code)

		var response dto.RegisterLectureResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.True(t, response.IsSuccess)
	})

	t.Run("전체 강의 목록 조회 성공", func(t *testing.T) {
		// When
		req := httptest.NewRequest("GET", "/api/lectures", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// Then
		assert.Equal(t, http.StatusOK, w.Code)

		var response dto.GetLecturesResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Len(t, response.Subjects, 1)

		subject := response.Subjects[0]
		assert.Equal(t, "", subject.Title)
		assert.Equal(t, "테스트 강사", subject.Instructor)
		assert.Equal(t, "테스트 강의", subject.Description)
		assert.Equal(t, "https://test.com/thumbnail.jpg", subject.ThumbnailURL)
	})

	t.Run("강의 상세 조회 성공", func(t *testing.T) {
		// Given
		// 먼저 전체 목록을 조회하여 강의 ID를 얻음
		req := httptest.NewRequest("GET", "/api/lectures", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var lecturesResponse dto.GetLecturesResponse
		json.Unmarshal(w.Body.Bytes(), &lecturesResponse)
		subjectID := lecturesResponse.Subjects[0].ID

		// When
		req = httptest.NewRequest("GET", "/api/lectures/"+strconv.FormatUint(uint64(subjectID), 10), nil)
		w = httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// Then
		assert.Equal(t, http.StatusOK, w.Code)

		var response dto.GetLectureResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "", response.Title)
		assert.Len(t, response.Videos, 2)
		assert.Equal(t, "테스트 비디오 1", response.Videos[0].Title)
		assert.Equal(t, "테스트 비디오 2", response.Videos[1].Title)
	})

	t.Run("존재하지 않는 강의 조회 실패", func(t *testing.T) {
		// When
		req := httptest.NewRequest("GET", "/api/lectures/999999", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// Then
		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("잘못된 플레이리스트 ID로 강의 등록 실패", func(t *testing.T) {
		// Given
		registerLectureRequest := dto.RegisterLectureRequest{
			PlayListID: "invalid_playlist_id",
		}
		jsonData, _ := json.Marshal(registerLectureRequest)

		// Mock 클라이언트가 에러를 반환하도록 설정
		mockYoutubeClient.GetPlayListByIDFunc = func(playlistId string) (*youtube.PlaylistListResponse, error) {
			return nil, errors.New("playlist not found")
		}

		// When
		req := httptest.NewRequest("POST", "/api/lectures", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// Then
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
