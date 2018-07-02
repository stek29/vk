package vkCallbackApi

import (
	"fmt"
	"github.com/mailru/easyjson"
	"net/http"
	"net/url"
)

// Inspired by https://github.com/Vorkytaka/easyvk-go

const (
	apiBaseUrl = "https://api.vk.com/method/"
)

type VKApi struct {
	AccessToken string
	Version     string
}

func ApiWithAccessToken(token string) (vk VKApi) {
	vk.AccessToken = token
	vk.Version = ApiVersion
	return
}

type ApiError struct {
	Code          int `json:"error_code"`
	Message       int `json:"error_msg"`
	RequestParams []struct {
		Key   string
		Value string
	}
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("ApiError %d: %s", e.Code, e.Message)
}

//easyjson:json
type VKApiResponse struct {
	Error    *ApiError
	Response easyjson.RawMessage
}

func (vk *VKApi) Request(method string, query url.Values) (easyjson.RawMessage, error) {
	u, err := url.Parse(apiBaseUrl + method)
	if err != nil {
		return nil, err
	}

	query.Set("access_token", vk.AccessToken)
	query.Set("v", vk.Version)
	u.RawQuery = query.Encode()

	r, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	resp := VKApiResponse{}
	easyjson.UnmarshalFromReader(r.Body, &resp)

	if resp.Error != nil {
		return nil, resp.Error
	}

	return resp.Response, nil
}
