package common

import (
	"context"
	"errors"
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

func ListVideosByPlayListID(playListID string) ([]*youtube.PlaylistItem, error) {
	service := GetService()
	call := service.PlaylistItems.List([]string{"snippet"}).PlaylistId(playListID)

	response, err := call.Do()
	if err != nil {
		return nil, err
	}

	result := make([]*youtube.PlaylistItem, response.PageInfo.TotalResults)
	idx := 0
	for _, video := range response.Items {
		result[idx] = video
		idx++
	}
	pageToken := response.NextPageToken

	for {
		if pageToken == "" {
			break
		}

		call.PageToken(pageToken)
		response, err := call.Do()
		if err != nil {
			return nil, err
		}

		for _, video := range response.Items {
			result[idx] = video
			idx++
		}
		pageToken = response.NextPageToken
	}

	return result, nil
}

func GetPlayListByID(playListID string) (*youtube.PlaylistListResponse, error) {
	service := GetService()
	call := service.Playlists.List([]string{"snippet"})
	response, err := call.Id(playListID).Do()

	if err != nil {
		return nil, err
	}

	if len(response.Items) == 0 {
		return nil, errors.New("no play list")
	}

	return response, nil
}
