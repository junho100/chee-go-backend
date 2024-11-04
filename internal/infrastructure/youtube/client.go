package youtube

import (
	"errors"

	"google.golang.org/api/youtube/v3"
)

type YoutubeClient interface {
	ListVideosByPlayListID(playListID string) ([]*youtube.PlaylistItem, error)
	GetPlayListByID(playListID string) (*youtube.PlaylistListResponse, error)
}

type youtubeClient struct {
	youtubeService *youtube.Service
}

func NewYoutubeClient(youtubeService *youtube.Service) YoutubeClient {
	return &youtubeClient{
		youtubeService: youtubeService,
	}
}

func (c *youtubeClient) ListVideosByPlayListID(playListID string) ([]*youtube.PlaylistItem, error) {
	call := c.youtubeService.PlaylistItems.List([]string{"snippet"}).PlaylistId(playListID)

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

func (c *youtubeClient) GetPlayListByID(playListID string) (*youtube.PlaylistListResponse, error) {
	call := c.youtubeService.Playlists.List([]string{"snippet"})
	response, err := call.Id(playListID).Do()

	if err != nil {
		return nil, err
	}

	if len(response.Items) == 0 {
		return nil, errors.New("no play list")
	}

	return response, nil
}
