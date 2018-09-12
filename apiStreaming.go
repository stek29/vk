package vkCallbackApi

import (
	"encoding/json"
)

// APIStreaming implements VK API namespace `streaming`
type APIStreaming struct {
	API *API
}

// StreamingGetServerURLResponse is response for APIStreaming.GetServerURL
//easyjson:json
type StreamingGetServerURLResponse struct {
	// Server host
	Endpoint string `json:"endpoint,omitempty"`
	// Access key
	Key string `json:"key,omitempty"`
}

// GetServerURL Allows to receive data for the connection to Streaming API.
func (v APIStreaming) GetServerURL() (*StreamingGetServerURLResponse, error) {
	r, err := v.API.Request("streaming.getServerUrl", nil)
	if err != nil {
		return nil, err
	}

	var resp StreamingGetServerURLResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
