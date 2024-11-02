package handler

import (
	"chee-go-backend/internal/domain/service"

	"github.com/gin-gonic/gin"
)

type LectureHandler struct {
	lectureService service.LectureService
}

func NewLectureHandler(router *gin.Engine, lectureService service.LectureService) {
	handler := &LectureHandler{
		lectureService: lectureService,
	}

	router.POST("/api/lectures", handler.CreateLecture)
	router.GET("/api/lectures", handler.GetAllSubjects)
	router.GET("/api/lectures/:id", handler.GetSubjectById)
}

func (h *LectureHandler) CreateLecture(c *gin.Context) {

}

func (h *LectureHandler) GetAllSubjects(c *gin.Context) {

}

func (h *LectureHandler) GetSubjectById(c *gin.Context) {

}
