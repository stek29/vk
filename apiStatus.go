package vk

import (
	"encoding/json"
)

// APIStatus implements VK API namespace `status`
type APIStatus struct {
	API *API
}

// StatusGetParams are params for APIStatus.Get
type StatusGetParams struct {
	// User ID or community ID. Use a negative value to designate a community ID.
	UserID  int `url:"user_id,omitempty"`
	GroupID int `url:"group_id,omitempty"`
}

// StatusGetResponse is response for APIStatus.Get
//easyjson:json
type StatusGetResponse struct {
	// Status text
	Text  string `json:"text,omitempty"`
	Audio Audio  `json:"audio,omitempty"`
}

// Get Returns data required to show the status of a user or community.
func (v APIStatus) Get(params StatusGetParams) (*StatusGetResponse, error) {
	r, err := v.API.Request("status.get", params)
	if err != nil {
		return nil, err
	}

	var resp StatusGetResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// StatusSetParams are params for APIStatus.Set
type StatusSetParams struct {
	// Text of the new status.
	Text string `url:"text,omitempty"`
	// Identifier of a community to set a status in. If left blank the status is set to current user.
	GroupID int `url:"group_id,omitempty"`
}

// Set Sets a new status for the current user.
func (v APIStatus) Set(params StatusSetParams) (bool, error) {
	r, err := v.API.Request("status.set", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}
