package vkapi

import (
	"encoding/json"

	"github.com/stek29/vk"
)

// Auth implements VK API namespace `auth`
type Auth struct {
	API vk.API
}

// AuthCheckPhoneParams are params for Auth.CheckPhone
type AuthCheckPhoneParams struct {
	// Phone number.
	Phone string `url:"phone"`
	// User ID.
	ClientID     int    `url:"client_id,omitempty"`
	ClientSecret string `url:"client_secret"`
	AuthByPhone  bool   `url:"auth_by_phone,omitempty"`
}

// CheckPhone Checks a user's phone number for correctness.
func (v Auth) CheckPhone(params AuthCheckPhoneParams) (bool, error) {
	r, err := v.API.Request("auth.checkPhone", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// AuthRestoreParams are params for Auth.Restore
type AuthRestoreParams struct {
	// User phone number.
	Phone string `url:"phone"`
	// User last name.
	LastName string `url:"last_name"`
}

// AuthRestoreResponse is response for Auth.Restore
//easyjson:json
type AuthRestoreResponse struct {
	// 1 if success
	Success int `json:"success,omitempty"`
	// Parameter needed to grant access by code
	Sid string `json:"sid,omitempty"`
}

// Restore Allows to restore account access using a code received via SMS. " This method is only available for apps with [vk.com/dev/auth_direct|Direct authorization] access. "
func (v Auth) Restore(params AuthRestoreParams) (*AuthRestoreResponse, error) {
	r, err := v.API.Request("auth.restore", params)
	if err != nil {
		return nil, err
	}

	var resp AuthRestoreResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
