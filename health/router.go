package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUsersRouters(router *gin.RouterGroup) {
	router.GET("", HealthCheck)
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
