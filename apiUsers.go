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
	UserIDs  CSVStringSlice `url:"user_ids,omitempty"`
	Fields   CSVStringSlice `url:"fields,omitempty"`
	NameCase string         `url:"name_case,omitempty"`
}

// UsersGetResponse is response for Users.Get
//easyjson:json
type UsersGetResponse []User

// Get is users.get
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
