package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type NotificationStatus interface {
	MarkAsProcessed(userID string, notificationID string) error
	IsProcessed(userID string, notificationID string) (bool, error)
}

type notificationStatus struct {
	client *redis.Client
}

func NewNotificationStatus(client *redis.Client) NotificationStatus {
	return &notificationStatus{
		client: client,
	}
}

func (n *notificationStatus) MarkAsProcessed(userID string, notificationID string) error {
	key := fmt.Sprintf("notification:%s:%s", userID, notificationID)
	return n.client.Set(context.Background(), key, "1", 30*24*time.Hour).Err()
}

func (n *notificationStatus) IsProcessed(userID string, notificationID string) (bool, error) {
	key := fmt.Sprintf("notification:%s:%s", userID, notificationID)
	exists, err := n.client.Exists(context.Background(), key).Result()
	if err != nil {
		return false, err
	}
	return exists == 1, nil
}
