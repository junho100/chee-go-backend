package router

import (
	"log"
	"time"

	"chee-go-backend/internal/infrastructure/monitoring"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// CloudWatch 클라이언트 초기화
	cloudwatch, err := monitoring.NewCloudWatchClient()
	if err != nil {
		log.Fatalf("CloudWatch 클라이언트 초기화 실패: %v", err)
	}

	// 메트릭스 미들웨어 적용
	metricsMiddleware := monitoring.NewMetricsMiddleware(cloudwatch)
	r.Use(metricsMiddleware.Metrics())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://chee-go.com", "https://www.chee-go.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.SetTrustedProxies(nil)

	return r
}
