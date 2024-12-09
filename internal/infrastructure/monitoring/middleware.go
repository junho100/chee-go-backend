package monitoring

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/gin-gonic/gin"
)

type MetricsMiddleware struct {
	cloudwatch *CloudWatchClient
}

func NewMetricsMiddleware(cloudwatch *CloudWatchClient) *MetricsMiddleware {
	return &MetricsMiddleware{
		cloudwatch: cloudwatch,
	}
}

func (m *MetricsMiddleware) Metrics() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// 요청 처리
		c.Next()

		// 메트릭 수집
		duration := time.Since(start).Milliseconds()
		path := c.FullPath()
		method := c.Request.Method
		status := c.Writer.Status()

		// API 응답 시간 메트릭
		m.cloudwatch.PutMetric(
			"CheeGo/API",
			"ResponseTime",
			float64(duration),
			types.StandardUnitMilliseconds,
			[]types.Dimension{
				{Name: stringPtr("Path"), Value: &path},
				{Name: stringPtr("Method"), Value: &method},
			},
		)

		// API 상태 코드 메트릭
		m.cloudwatch.PutMetric(
			"CheeGo/API",
			"StatusCode",
			float64(status),
			types.StandardUnitCount,
			[]types.Dimension{
				{Name: stringPtr("Path"), Value: &path},
				{Name: stringPtr("Method"), Value: &method},
			},
		)
	}
}

func stringPtr(s string) *string {
	return &s
}
