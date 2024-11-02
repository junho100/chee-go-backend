package handler

import (
	"chee-go-backend/internal/domain/service"

	"github.com/gin-gonic/gin"
)

type ResumeHandler struct {
	resumeService service.ResumeService
}

func NewResumeHandler(router *gin.Engine, resumeService service.ResumeService) {
	handler := &ResumeHandler{
		resumeService: resumeService,
	}

	router.POST("/api/resumes", handler.CreateResume)
	router.GET("/api/resumes", handler.GetResume)
	router.GET("/api/resumes/wanted", handler.GetWantedResume)
	router.GET("/api/resumes/programmers", handler.GetProgrammersResume)
	router.GET("/api/resumes/linkedin", handler.GetLinkedinResume)
}

func (h *ResumeHandler) CreateResume(c *gin.Context) {

}

func (h *ResumeHandler) GetResume(c *gin.Context) {

}

func (h *ResumeHandler) GetWantedResume(c *gin.Context) {

}

func (h *ResumeHandler) GetProgrammersResume(c *gin.Context) {

}

func (h *ResumeHandler) GetLinkedinResume(c *gin.Context) {

}
