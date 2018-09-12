package vkCallbackApi

import (
	"encoding/json"
)

// APIGifts implements VK API namespace `gifts`
type APIGifts struct {
	API *API
}

// GiftsGetParams are params for APIGifts.Get
type GiftsGetParams struct {
	// User ID.
	UserID int `url:"user_id,omitempty"`
	// Number of gifts to return.
	Count int `url:"count,omitempty"`
	// Offset needed to return a specific subset of results.
	Offset int `url:"offset,omitempty"`
}

// GiftsGetResponse is response for APIGifts.Get
//easyjson:json
type GiftsGetResponse struct {
	// Total number
	Count int    `json:"count,omitempty"`
	Items []Gift `json:"items,omitempty"`
}

// Get Returns a list of user gifts.
func (v APIGifts) Get(params GiftsGetParams) (*GiftsGetResponse, error) {
	r, err := v.API.Request("gifts.get", params)
	if err != nil {
		return nil, err
	}

	var resp GiftsGetResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
