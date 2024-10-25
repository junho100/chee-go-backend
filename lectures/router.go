package lectures

import (
	"chee-go-backend/common"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterLecturesRouters(router *gin.RouterGroup) {
	router.POST("", RegisterLecture)
	router.GET("", GetLectures)
	router.GET(":id", GetLecture)
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
		log.Println(err)
		response := &common.CommonErrorResponse{
			Message: "failed to get play list.",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	youtubePlayListItems, err := common.ListVideosByPlayListID(registerLectureRequest.PlayListID)

	if err != nil {
		log.Println(err)
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

func GetLecture(c *gin.Context) {
	subjectID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		response := &common.CommonErrorResponse{
			Message: "bad path variable.",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	subject, err := GetSubjectByID(uint(subjectID))
	if err != nil {
		response := &common.CommonErrorResponse{
			Message: "failed to get subject.",
		}
		c.JSON(http.StatusNotFound, response)
		return
	}

	var response GetLectureResponse
	response.from(*subject)
	c.JSON(http.StatusOK, response)
}
