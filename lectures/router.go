package lectures

import "github.com/gin-gonic/gin"

func RegisterLecturesRouters(router *gin.RouterGroup) {
	router.POST("", RegisterLecture)
	router.GET("", GetLectures)
}

func RegisterLecture(c *gin.Context) {

}

func GetLectures(c *gin.Context) {

}
