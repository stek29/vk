package vkCallbackApi

import (
	"github.com/mailru/easyjson"
)

// APIUsers implements VK API namespace `users`
type APIUsers struct {
	API *API
}

// UsersGetParams are params for Users.Get
type UsersGetParams struct {
	// User IDs or screen names. Current user ID if empty.
	UserIDs CSVStringSlice `url:"user_ids,omitempty"`
	// Additional user fields to return. See User type.
	Fields   CSVStringSlice `url:"fields,omitempty"`
	NameCase string         `url:"name_case,omitempty"`
}

// UsersGetResponse is response for Users.Get
//easyjson:json
type UsersGetResponse []User

// Get is users.get
// Returns detailed information on users.
func (v APIUsers) Get(params UsersGetParams) (UsersGetResponse, error) {
	r, err := v.API.Request("users.get", params)
	if err != nil {
		return nil, err
	}

	resp := UsersGetResponse{}
	err = easyjson.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
