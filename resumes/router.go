package resumes

import (
	"chee-go-backend/common"
	"chee-go-backend/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterResumesRouters(router *gin.RouterGroup) {
	router.POST("", RegisterResume)
	router.GET("", GetResume)
	router.GET("/wanted", GetWantedResume)
}

func RegisterResume(c *gin.Context) {
	var token string
	var err error
	var userID string
	var registerResumeRequest RegisterResumeRequest

	if err := c.BindJSON(&registerResumeRequest); err != nil {
		response := &common.CommonErrorResponse{
			Message: "bad content.",
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if token, err = users.ExtractToken(c.GetHeader("Authorization")); err != nil {
		response := &common.CommonErrorResponse{
			Message: "failed to authorization.",
		}
		c.JSON(http.StatusForbidden, response)
		return
	}

	if userID, err = users.GetUserIDFromToken(token); err != nil {
		response := &common.CommonErrorResponse{
			Message: "failed to authorization.",
		}
		c.JSON(http.StatusForbidden, response)
		return
	}

	// TODO: Mapping correctly
	dto := &CreateResumeDTO{
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
	if resumeID, err = CreateResume(dto); err != nil {
		response := &common.CommonErrorResponse{
			Message: "failed to create resume.",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := &RegisterResumeResponse{
		ResumeID: resumeID,
	}
	c.JSON(http.StatusCreated, response)
}

func GetResume(c *gin.Context) {
	var token string
	var err error
	var userID string

	if token, err = users.ExtractToken(c.GetHeader("Authorization")); err != nil {
		response := &common.CommonErrorResponse{
			Message: "failed to authorization.",
		}
		c.JSON(http.StatusForbidden, response)
		return
	}

	if userID, err = users.GetUserIDFromToken(token); err != nil {
		response := &common.CommonErrorResponse{
			Message: "failed to authorization.",
		}
		c.JSON(http.StatusForbidden, response)
		return
	}

	resume, err := GetResumeByUserID(userID)

	if err != nil {
		response := &common.CommonErrorResponse{
			Message: "failed to get resume.",
		}
		c.JSON(http.StatusNotFound, response)
		return
	}

	keywords := GetKeywordsByResumeID(resume.ID)

	var response GetResumeResponse
	response.from(*resume, keywords)
	c.JSON(http.StatusOK, response)
}

func GetWantedResume(c *gin.Context) {
	var token string
	var err error
	var userID string

	if token, err = users.ExtractToken(c.GetHeader("Authorization")); err != nil {
		response := &common.CommonErrorResponse{
			Message: "failed to authorization.",
		}
		c.JSON(http.StatusForbidden, response)
		return
	}

	if userID, err = users.GetUserIDFromToken(token); err != nil {
		response := &common.CommonErrorResponse{
			Message: "failed to authorization.",
		}
		c.JSON(http.StatusForbidden, response)
		return
	}

	resume, err := GetResumeByUserID(userID)
	keywords := GetKeywordsByResumeID(resume.ID)

	wantedResume := ConvertResumeToWanted(*resume, keywords)

	var response GetWantedResumeResponse
	response.from(wantedResume)
	c.JSON(http.StatusOK, response)
}
