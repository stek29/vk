package vkCallbackApi

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/mailru/easyjson"
)

// TODO: Generate Params/Response structs and Methods for all API calls

type VideoGetParams struct {
	OwnerID  *int
	Videos   []string
	AlbumID  *int
	Count    *int
	Offset   *int
	Extended bool
}

//easyjson:json
type VideoGetResponse struct {
	Count int
	Items []Video
}

func (vk *VKApi) VideoGet(params VideoGetParams) (*VideoGetResponse, error) {
	method := "video.get"
	query := url.Values{}

	if v := params.OwnerID; v != nil {
		query.Set("owner_id", strconv.Itoa(*v))
	}

	if v := params.Videos; len(v) != 0 {
		videos := strings.Join(v, ",")
		query.Set("videos", videos)
	}

	if v := params.AlbumID; v != nil {
		query.Set("album_id", strconv.Itoa(*v))
	}

	if v := params.Count; v != nil {
		query.Set("count", strconv.Itoa(*v))
	}

	if v := params.Offset; v != nil {
		query.Set("offset", strconv.Itoa(*v))
	}

	if params.Extended {
		query.Set("extended", "1")
	}

	r, err := vk.Request(method, query)
	if err != nil {
		return nil, err
	}

	resp := &VideoGetResponse{}
	err = easyjson.Unmarshal(r, resp)

	if err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}
