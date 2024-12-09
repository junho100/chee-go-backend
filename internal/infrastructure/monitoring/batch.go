package monitoring

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
)

type BatchMetrics struct {
	cloudwatch *CloudWatchClient
}

func NewBatchMetrics(cloudwatch *CloudWatchClient) *BatchMetrics {
	return &BatchMetrics{
		cloudwatch: cloudwatch,
	}
}

func (b *BatchMetrics) TrackBatchJob(jobName string, fn func() error) error {
	start := time.Now()

	err := fn()

	// 실행 시간 메트릭
	duration := time.Since(start).Seconds()
	b.cloudwatch.PutMetric(
		"CheeGo/Batch",
		"ExecutionTime",
		duration,
		types.StandardUnitSeconds,
		[]types.Dimension{
			{Name: stringPtr("JobName"), Value: &jobName},
		},
	)

	return err
}
