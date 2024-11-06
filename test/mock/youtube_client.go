package mock

import (
	"google.golang.org/api/youtube/v3"
)

type MockYoutubeClient struct {
	ListVideosByPlayListIDFunc func(playListID string) ([]*youtube.PlaylistItem, error)
	GetPlayListByIDFunc        func(playListID string) (*youtube.PlaylistListResponse, error)
}

func (m *MockYoutubeClient) ListVideosByPlayListID(playListID string) ([]*youtube.PlaylistItem, error) {
	return m.ListVideosByPlayListIDFunc(playListID)
}

func (m *MockYoutubeClient) GetPlayListByID(playListID string) (*youtube.PlaylistListResponse, error) {
	return m.GetPlayListByIDFunc(playListID)
}
