package vkCallbackApi

import (
	"github.com/mailru/easyjson"
)

// APIAuth implements VK API namespace `auth`
type APIAuth struct {
	API *API
}

// AuthCheckPhoneParams are params for APIAuth.CheckPhone
type AuthCheckPhoneParams struct {
	// Phone number.
	Phone string `url:"phone"`
	// User ID.
	ClientID     int    `url:"client_id,omitempty"`
	ClientSecret string `url:"client_secret"`
	AuthByPhone  bool   `url:"auth_by_phone,omitempty"`
}

// CheckPhone Checks a user's phone number for correctness.
func (v APIAuth) CheckPhone(params AuthCheckPhoneParams) (bool, error) {
	r, err := v.API.Request("auth.checkPhone", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// AuthSignupParams are params for APIAuth.Signup
type AuthSignupParams struct {
	// User's first name.
	FirstName string `url:"first_name"`
	// User's surname.
	LastName string `url:"last_name"`
	// User's birthday.
	Birthday string `url:"birthday"`
	// Your application ID.
	ClientID     int    `url:"client_id"`
	ClientSecret string `url:"client_secret"`
	// User's phone number. Can be pre-checked with the [vk.com/dev/auth.checkPhone|auth.checkPhone] method.
	Phone string `url:"phone"`
	// User's password (minimum of 6 characters). Can be specified later with the [vk.com/dev/auth.confirm|auth.confirm] method.
	Password string `url:"password,omitempty"`
	// '1' — test mode, in which the user will not be registered and the phone number will not be checked for availability, '0' — default mode (default)
	TestMode bool `url:"test_mode,omitempty"`
	// '1' — call the phone number and leave a voice message of the authorization code, '0' — send the code by SMS (default)
	Voice bool `url:"voice,omitempty"`
	// '1' — female, '2' — male
	Sex int `url:"sex,omitempty"`
	// Session ID required for method recall when SMS was not delivered.
	Sid string `url:"sid,omitempty"`
}

// AuthSignupResponse is response for APIAuth.Signup
//easyjson:json
type AuthSignupResponse struct {
	// Parameter to retry
	Sid string `json:"sid,omitempty"`
}

// Signup Registers a new user by phone number.
func (v APIAuth) Signup(params AuthSignupParams) (*AuthSignupResponse, error) {
	r, err := v.API.Request("auth.signup", params)
	if err != nil {
		return nil, err
	}

	var resp AuthSignupResponse
	err = easyjson.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AuthConfirmParams are params for APIAuth.Confirm
type AuthConfirmParams struct {
	ClientID     int    `url:"client_id"`
	ClientSecret string `url:"client_secret"`
	Phone        string `url:"phone"`
	Code         string `url:"code"`
	Password     string `url:"password,omitempty"`
	TestMode     bool   `url:"test_mode,omitempty"`
	Intro        int    `url:"intro,omitempty"`
}

// AuthConfirmResponse is response for APIAuth.Confirm
//easyjson:json
type AuthConfirmResponse struct {
	// 1 if success
	Success int `json:"success,omitempty"`
	// User ID
	UserID int `json:"user_id,omitempty"`
}

// Confirm Completes a user's registration (begun with the [vk.com/dev/auth.signup|auth.signup] method) using an authorization code.
func (v APIAuth) Confirm(params AuthConfirmParams) (*AuthConfirmResponse, error) {
	r, err := v.API.Request("auth.confirm", params)
	if err != nil {
		return nil, err
	}

	var resp AuthConfirmResponse
	err = easyjson.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AuthRestoreParams are params for APIAuth.Restore
type AuthRestoreParams struct {
	// User phone number.
	Phone string `url:"phone"`
	// User last name.
	LastName string `url:"last_name"`
}

// AuthRestoreResponse is response for APIAuth.Restore
//easyjson:json
type AuthRestoreResponse struct {
	// 1 if success
	Success int `json:"success,omitempty"`
	// Parameter needed to grant access by code
	Sid string `json:"sid,omitempty"`
}

// Restore Allows to restore account access using a code received via SMS. " This method is only available for apps with [vk.com/dev/auth_direct|Direct authorization] access. "
func (v APIAuth) Restore(params AuthRestoreParams) (*AuthRestoreResponse, error) {
	r, err := v.API.Request("auth.restore", params)
	if err != nil {
		return nil, err
	}

	var resp AuthRestoreResponse
	err = easyjson.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
