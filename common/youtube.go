package common

import (
	"context"
	"log"
	"os"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

var SERVICE *youtube.Service

func GetService() *youtube.Service {
	return SERVICE
}

func InitYoutube() {
	ctx := context.Background()
	apiKey := os.Getenv("YOUTUBE_API_KEY")

	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))

	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
	}

	SERVICE = service
}

func ListVideosByPlayListID(playListID string) (*youtube.PlaylistItemListResponse, error) {
	service := GetService()
	call := service.PlaylistItems.List([]string{"snippet"})
	response, err := call.PlaylistId(playListID).Do()

	if err != nil {
		log.Fatal(err)
	}

	return response, nil
}
