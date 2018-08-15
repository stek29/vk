package vkCallbackApi

import (
	"github.com/mailru/easyjson"
)

// APIVideo implements VK API namespace `video`
type APIVideo struct {
	api *APIBase
}

// VideoGetParams are params for Video.Get
type VideoGetParams struct {
	OwnerID  int            `url:"owner_id,omitempty"`
	Videos   CSVStringSlice `url:"videos,omitempty"`
	AlbumID  int            `url:"album_id,omitempty"`
	Count    *int           `url:"count,omitempty"`
	Offset   int            `url:"offset,omitemtpy"`
	Extended bool           `url:"extended,omitempty"`
}

// VideoGetResponse is response for Video.Get
//easyjson:json
type VideoGetResponse struct {
	Count int
	Items []Video
}

// Get is video.get
func (v APIVideo) Get(params VideoGetParams) (*VideoGetResponse, error) {
	r, err := v.api.Request("video.get", params)
	if err != nil {
		return nil, err
	}

	resp := &VideoGetResponse{}
	err = easyjson.Unmarshal(r, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
