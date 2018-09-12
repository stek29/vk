package vkCallbackApi

import (
	"github.com/mailru/easyjson"
)

// APIStorage implements VK API namespace `storage`
type APIStorage struct {
	API *API
}

// StorageGetParams are params for APIStorage.Get
type StorageGetParams struct {
	Key    string         `url:"key,omitempty"`
	Keys   CSVStringSlice `url:"keys,omitempty"`
	UserID int            `url:"user_id,omitempty"`
}

// StorageGetResponse is response for APIStorage.Get
// Key value
type StorageGetResponse string

// Get Returns a value of variable with the name set by key parameter.
func (v APIStorage) Get(params StorageGetParams) (StorageGetResponse, error) {
	r, err := v.API.Request("storage.get", params)
	if err != nil {
		return "", err
	}

	var resp StorageGetResponse

	resp = StorageGetResponse(string(r))

	if err != nil {
		return "", err
	}
	return resp, nil
}

// StorageSetParams are params for APIStorage.Set
type StorageSetParams struct {
	Key    string `url:"key"`
	Value  string `url:"value,omitempty"`
	UserID int    `url:"user_id,omitempty"`
}

// Set Saves a value of variable with the name set by 'key' parameter.
func (v APIStorage) Set(params StorageSetParams) (bool, error) {
	r, err := v.API.Request("storage.set", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// StorageGetKeysParams are params for APIStorage.GetKeys
type StorageGetKeysParams struct {
	// user id, whose variables names are returned if they were requested with a server method.
	UserID int `url:"user_id,omitempty"`
	// amount of variable names the info needs to be collected from.
	Count int `url:"count,omitempty"`
}

// StorageGetKeysResponse is response for APIStorage.GetKeys
// Key name
//easyjson:json
type StorageGetKeysResponse []string

// GetKeys Returns the names of all variables.
func (v APIStorage) GetKeys(params StorageGetKeysParams) (StorageGetKeysResponse, error) {
	r, err := v.API.Request("storage.getKeys", params)
	if err != nil {
		return nil, err
	}

	var resp StorageGetKeysResponse
	err = easyjson.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
