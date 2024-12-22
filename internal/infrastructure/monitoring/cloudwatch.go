package monitoring

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
)

type CloudWatchClient struct {
	client *cloudwatch.Client
}

func NewCloudWatchClient() (*CloudWatchClient, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("ap-northeast-2"),
	)
	if err != nil {
		return nil, err
	}

	client := cloudwatch.NewFromConfig(cfg)
	return &CloudWatchClient{client: client}, nil
}

func (c *CloudWatchClient) PutMetric(namespace, metricName string, value float64, unit types.StandardUnit, dimensions []types.Dimension) error {
	if len(dimensions) > 0 {
		for _, dim := range dimensions {
			if dim.Name == nil || *dim.Name == "" {
				return fmt.Errorf("dimension name이 비어있습니다")
			}
			if dim.Value == nil || *dim.Value == "" {
				return fmt.Errorf("dimension value가 비어있습니다")
			}
		}
	}

	now := time.Now()

	input := &cloudwatch.PutMetricDataInput{
		Namespace: &namespace,
		MetricData: []types.MetricDatum{
			{
				MetricName: &metricName,
				Value:      &value,
				Unit:       unit,
				Timestamp:  &now,
				Dimensions: dimensions,
			},
		},
	}

	_, err := c.client.PutMetricData(context.TODO(), input)
	if err != nil {
		log.Printf("CloudWatch 메트릭 전송 실패: %v", err)
		return err
	}

	return nil
}
