package lectures

import (
	"chee-go-backend/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterLecturesRouters(router *gin.RouterGroup) {
	router.POST("", RegisterLecture)
	router.GET("", GetLectures)
}

func RegisterLecture(c *gin.Context) {
	var registerLectureRequest RegisterLectureRequest
	var err error

	if err := c.BindJSON(&registerLectureRequest); err != nil {
		response := &common.CommonErrorResponse{
			Message: "bad content.",
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	youtubePlayListResponse, err := common.GetPlayListByID(registerLectureRequest.PlayListID)

	if err != nil {
		response := &common.CommonErrorResponse{
			Message: "failed to get play list.",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	youtubePlayListItems, err := common.ListVideosByPlayListID(registerLectureRequest.PlayListID)

	if err != nil {
		response := &common.CommonErrorResponse{
			Message: "failed to get play list items.",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if err := CreateSubjectWithLectures(youtubePlayListResponse, youtubePlayListItems); err != nil {
		response := &common.CommonErrorResponse{
			Message: "failed to create subject.",
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := &RegisterLectureResponse{
		IsSuccess: true,
	}
	c.JSON(http.StatusCreated, response)
}

func GetLectures(c *gin.Context) {
	var getLecturesResponse GetLecturesResponse

	subjects := GetAllSubjects()
	getLecturesResponse.Subjects = make([]GetLecturesResponseSubject, len(subjects))
	for i, subject := range subjects {
		getLecturesResponse.Subjects[i] = GetLecturesResponseSubject{
			ID:           subject.ID,
			Title:        subject.SubjectName,
			Description:  subject.Name,
			ThumbnailURL: subject.ThumbnailURL,
			Instructor:   subject.LecturerName,
		}
	}

	c.JSON(http.StatusOK, getLecturesResponse)
}
