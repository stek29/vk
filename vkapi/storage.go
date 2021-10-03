package vkapi

import (
	"encoding/json"

	"github.com/stek29/vk"
)

// Storage implements VK API namespace `storage`
type Storage struct {
	API vk.API
}

// StorageGetParams are params for Storage.Get
type StorageGetParams struct {
	Key    string         `url:"key,omitempty"`
	Keys   CSVStringSlice `url:"keys,omitempty"`
	UserID int            `url:"user_id,omitempty"`
	Global bool           `url:"global,omitempty"`
}

// StorageGetResponse is response for Storage.Get
// Key value
type StorageGetResponse string

// Get Returns a value of variable with the name set by key parameter.
func (v Storage) Get(params StorageGetParams) (StorageGetResponse, error) {
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

// StorageSetParams are params for Storage.Set
type StorageSetParams struct {
	Key    string `url:"key"`
	Value  string `url:"value,omitempty"`
	UserID int    `url:"user_id,omitempty"`
	Global bool   `url:"global,omitempty"`
}

// Set Saves a value of variable with the name set by 'key' parameter.
func (v Storage) Set(params StorageSetParams) (bool, error) {
	r, err := v.API.Request("storage.set", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// StorageGetKeysParams are params for Storage.GetKeys
type StorageGetKeysParams struct {
	// user id, whose variables names are returned if they were requested with a server method.
	UserID int  `url:"user_id,omitempty"`
	Global bool `url:"global,omitempty"`
	Offset int  `url:"offset,omitempty"`
	// amount of variable names the info needs to be collected from.
	Count int `url:"count,omitempty"`
}

// StorageGetKeysResponse is response for Storage.GetKeys
// Key name
//easyjson:json
type StorageGetKeysResponse []string

// GetKeys Returns the names of all variables.
func (v Storage) GetKeys(params StorageGetKeysParams) (StorageGetKeysResponse, error) {
	r, err := v.API.Request("storage.getKeys", params)
	if err != nil {
		return nil, err
	}

	var resp StorageGetKeysResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
