package handler

import (
	"chee-go-backend/internal/common"
	"chee-go-backend/internal/domain/service"
	"chee-go-backend/internal/http/dto"
	"chee-go-backend/internal/infrastructure/youtube"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LectureHandler struct {
	lectureService service.LectureService
	youtubeClient  youtube.YoutubeClient
}

func NewLectureHandler(router *gin.Engine, lectureService service.LectureService, youtubeClient youtube.YoutubeClient) {
	handler := &LectureHandler{
		lectureService: lectureService,
		youtubeClient:  youtubeClient,
	}

	router.POST("/api/lectures", handler.CreateLecture)
	router.GET("/api/lectures", handler.GetAllSubjects)
	router.GET("/api/lectures/:id", handler.GetSubjectById)
}

func (h *LectureHandler) CreateLecture(c *gin.Context) {
	var registerLectureRequest dto.RegisterLectureRequest
	var err error

	if err := c.BindJSON(&registerLectureRequest); err != nil {
		response := &common.CommonErrorResponse{
			Message: "bad content.",
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	youtubePlayListResponse, err := h.youtubeClient.GetPlayListByID(registerLectureRequest.PlayListID)

	if err != nil {
		response := &common.CommonErrorResponse{
			Message: "failed to get play list.",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	youtubePlayListItems, err := h.youtubeClient.ListVideosByPlayListID(registerLectureRequest.PlayListID)

	if err != nil {
		log.Println(err)
		response := &common.CommonErrorResponse{
			Message: "failed to get play list items.",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if err := h.lectureService.CreateSubjectWithLectures(youtubePlayListResponse, youtubePlayListItems); err != nil {
		response := &common.CommonErrorResponse{
			Message: "failed to create subject.",
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := &dto.RegisterLectureResponse{
		IsSuccess: true,
	}
	c.JSON(http.StatusCreated, response)
}

func (h *LectureHandler) GetAllSubjects(c *gin.Context) {
	var getLecturesResponse dto.GetLecturesResponse

	subjects := h.lectureService.GetAllSubjects()
	getLecturesResponse.Subjects = make([]dto.GetLecturesResponseSubject, len(subjects))
	for i, subject := range subjects {
		getLecturesResponse.Subjects[i] = dto.GetLecturesResponseSubject{
			ID:           subject.ID,
			Title:        subject.SubjectName,
			Description:  subject.Name,
			ThumbnailURL: subject.ThumbnailURL,
			Instructor:   subject.LecturerName,
		}
	}

	c.JSON(http.StatusOK, getLecturesResponse)
}

func (h *LectureHandler) GetSubjectById(c *gin.Context) {
	subjectID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		response := &common.CommonErrorResponse{
			Message: "bad path variable.",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	subject, err := h.lectureService.GetSubjectByID(uint(subjectID))
	if err != nil {
		response := &common.CommonErrorResponse{
			Message: "failed to get subject.",
		}
		c.JSON(http.StatusNotFound, response)
		return
	}

	var response dto.GetLectureResponse
	response.From(*subject)
	c.JSON(http.StatusOK, response)
}
