package vkapi

import (
	"encoding/json"

	"github.com/stek29/vk"
)

// Status implements VK API namespace `status`
type Status struct {
	API vk.API
}

// StatusGetParams are params for Status.Get
type StatusGetParams struct {
	// User ID or community ID. Use a negative value to designate a community ID.
	UserID  int `url:"user_id,omitempty"`
	GroupID int `url:"group_id,omitempty"`
}

// StatusGetResponse is response for Status.Get
//easyjson:json
type StatusGetResponse struct {
	// Status text
	Text  string   `json:"text,omitempty"`
	Audio vk.Audio `json:"audio,omitempty"`
}

// Get Returns data required to show the status of a user or community.
func (v Status) Get(params StatusGetParams) (*StatusGetResponse, error) {
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

// StatusSetParams are params for Status.Set
type StatusSetParams struct {
	// Text of the new status.
	Text string `url:"text,omitempty"`
	// Identifier of a community to set a status in. If left blank the status is set to current user.
	GroupID int `url:"group_id,omitempty"`
}

// Set Sets a new status for the current user.
func (v Status) Set(params StatusSetParams) (bool, error) {
	r, err := v.API.Request("status.set", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}
