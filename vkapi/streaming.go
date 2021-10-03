package vkapi

import (
	"encoding/json"

	"github.com/stek29/vk"
)

// Streaming implements VK API namespace `streaming`
type Streaming struct {
	API vk.API
}

// StreamingGetServerURLResponse is response for Streaming.GetServerURL
//easyjson:json
type StreamingGetServerURLResponse struct {
	// Server host
	Endpoint string `json:"endpoint,omitempty"`
	// Access key
	Key string `json:"key,omitempty"`
}

// GetServerURL Allows to receive data for the connection to Streaming API.
func (v Streaming) GetServerURL() (*StreamingGetServerURLResponse, error) {
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

// StreamingSetSettingsParams are params for Streaming.SetSettings
type StreamingSetSettingsParams struct {
	MonthlyTier string `url:"monthly_tier,omitempty"`
}

// SetSettings does streaming.setSettings
func (v Streaming) SetSettings(params StreamingSetSettingsParams) (bool, error) {
	r, err := v.API.Request("streaming.setSettings", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}
