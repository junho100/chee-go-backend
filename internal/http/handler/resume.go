package handler

import (
	"chee-go-backend/internal/common"
	"chee-go-backend/internal/domain/service"
	"chee-go-backend/internal/http/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResumeHandler struct {
	resumeService service.ResumeService
	userService   service.UserService
}

func NewResumeHandler(router *gin.Engine, resumeService service.ResumeService, userService service.UserService) {
	handler := &ResumeHandler{
		resumeService: resumeService,
		userService:   userService,
	}

	router.POST("/api/resumes", handler.CreateResume)
	router.GET("/api/resumes", handler.GetResume)
	router.GET("/api/resumes/wanted", handler.GetWantedResume)
	router.GET("/api/resumes/programmers", handler.GetProgrammersResume)
	router.GET("/api/resumes/linkedin", handler.GetLinkedinResume)
}

func (h *ResumeHandler) CreateResume(c *gin.Context) {
	var token string
	var err error
	var userID string
	var registerResumeRequest dto.RegisterResumeRequest

	if err := c.BindJSON(&registerResumeRequest); err != nil {
		response := &common.CommonErrorResponse{
			Message: "bad content.",
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if token, err = h.userService.ExtractToken(c.GetHeader("Authorization")); err != nil {
		response := &common.CommonErrorResponse{
			Message: "failed to authorization.",
		}
		c.JSON(http.StatusForbidden, response)
		return
	}

	if userID, err = h.userService.GetUserIDFromToken(token); err != nil {
		response := &common.CommonErrorResponse{
			Message: "failed to authorization.",
		}
		c.JSON(http.StatusForbidden, response)
		return
	}

	// TODO: Mapping correctly
	createResumeDto := &dto.CreateResumeDTO{
		Introduction:    registerResumeRequest.Introduction,
		GithubURL:       registerResumeRequest.GithubURL,
		BlogURL:         registerResumeRequest.BlogURL,
		Educations:      registerResumeRequest.Educations,
		Projects:        registerResumeRequest.Projects,
		Activities:      registerResumeRequest.Activities,
		Certificates:    registerResumeRequest.Certificates,
		WorkExperiences: registerResumeRequest.WorkExperiences,
		Keywords:        registerResumeRequest.Keywords,
		UserID:          userID,
	}

	var resumeID uint
	if resumeID, err = h.resumeService.CreateResume(createResumeDto); err != nil {
		response := &common.CommonErrorResponse{
			Message: "failed to create resume.",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := &dto.RegisterResumeResponse{
		ResumeID: resumeID,
	}
	c.JSON(http.StatusCreated, response)
}

func (h *ResumeHandler) GetResume(c *gin.Context) {
	var token string
	var err error
	var userID string

	if token, err = h.userService.ExtractToken(c.GetHeader("Authorization")); err != nil {
		response := &common.CommonErrorResponse{
			Message: "failed to authorization.",
		}
		c.JSON(http.StatusForbidden, response)
		return
	}

	if userID, err = h.userService.GetUserIDFromToken(token); err != nil {
		response := &common.CommonErrorResponse{
			Message: "failed to authorization.",
		}
		c.JSON(http.StatusForbidden, response)
		return
	}

	resume, err := h.resumeService.GetResumeByUserID(userID)

	if err != nil {
		response := &common.CommonErrorResponse{
			Message: "failed to get resume.",
		}
		c.JSON(http.StatusNotFound, response)
		return
	}

	keywords := h.resumeService.GetKeywordsByResumeID(resume.ID)

	var response dto.GetResumeResponse
	response.From(*resume, keywords)
	c.JSON(http.StatusOK, response)
}

func (h *ResumeHandler) GetWantedResume(c *gin.Context) {
	var token string
	var err error
	var userID string

	if token, err = h.userService.ExtractToken(c.GetHeader("Authorization")); err != nil {
		response := &common.CommonErrorResponse{
			Message: "failed to authorization.",
		}
		c.JSON(http.StatusForbidden, response)
		return
	}

	if userID, err = h.userService.GetUserIDFromToken(token); err != nil {
		response := &common.CommonErrorResponse{
			Message: "failed to authorization.",
		}
		c.JSON(http.StatusForbidden, response)
		return
	}

	resume, err := h.resumeService.GetResumeByUserID(userID)

	if err != nil {
		response := &common.CommonErrorResponse{
			Message: "failed to get resume.",
		}
		c.JSON(http.StatusNotFound, response)
		return
	}

	keywords := h.resumeService.GetKeywordsByResumeID(resume.ID)

	wantedResume := h.resumeService.ConvertResumeToWanted(*resume, keywords)

	var response dto.GetWantedResumeResponse
	response.From(wantedResume)
	c.JSON(http.StatusOK, response)
}

func (h *ResumeHandler) GetProgrammersResume(c *gin.Context) {
	var token string
	var err error
	var userID string

	if token, err = h.userService.ExtractToken(c.GetHeader("Authorization")); err != nil {
		response := &common.CommonErrorResponse{
			Message: "failed to authorization.",
		}
		c.JSON(http.StatusForbidden, response)
		return
	}

	if userID, err = h.userService.GetUserIDFromToken(token); err != nil {
		response := &common.CommonErrorResponse{
			Message: "failed to authorization.",
		}
		c.JSON(http.StatusForbidden, response)
		return
	}

	resume, err := h.resumeService.GetResumeByUserID(userID)

	if err != nil {
		response := &common.CommonErrorResponse{
			Message: "failed to get resume.",
		}
		c.JSON(http.StatusNotFound, response)
		return
	}

	programmersResume := h.resumeService.ConvertResumeToProgrammers(*resume)

	var response dto.GetProgrammersResumeResponse
	response.From(programmersResume)
	c.JSON(http.StatusOK, response)
}

func (h *ResumeHandler) GetLinkedinResume(c *gin.Context) {
	var token string
	var err error
	var userID string

	if token, err = h.userService.ExtractToken(c.GetHeader("Authorization")); err != nil {
		response := &common.CommonErrorResponse{
			Message: "failed to authorization.",
		}
		c.JSON(http.StatusForbidden, response)
		return
	}

	if userID, err = h.userService.GetUserIDFromToken(token); err != nil {
		response := &common.CommonErrorResponse{
			Message: "failed to authorization.",
		}
		c.JSON(http.StatusForbidden, response)
		return
	}

	resume, err := h.resumeService.GetResumeByUserID(userID)

	if err != nil {
		response := &common.CommonErrorResponse{
			Message: "failed to get resume.",
		}
		c.JSON(http.StatusNotFound, response)
		return
	}

	keywords := h.resumeService.GetKeywordsByResumeID(resume.ID)

	linkedinResume := h.resumeService.ConvertResumeToLinkedin(*resume, keywords)

	var response dto.GetLinkedinResumeResponse
	response.From(linkedinResume)
	c.JSON(http.StatusOK, response)
}
