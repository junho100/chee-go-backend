package resumes

import "github.com/gin-gonic/gin"

func RegisterResumesRouters(router *gin.RouterGroup) {
	router.POST("/", RegisterResume)
}

func RegisterResume(c *gin.Context) {

}
