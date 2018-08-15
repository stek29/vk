package vkCallbackApi

import (
	"fmt"
	"net/url"

	"net/http"

	"github.com/google/go-querystring/query"
	"github.com/mailru/easyjson"
)

const (
	apiBaseURL = "https://api.vk.com/method/"
)

// TODO: Add Lang

// APIBase is minimal struct holding parameters required for requests
type APIBase struct {
	BaseURL     string `url:"-"`
	AccessToken string `url:"access_token,omitempty"`
	Version     string `url:"v,omitempty"`
}

// API is a helper type used for making requests
type API struct {
	APIBase *APIBase

	Video APIVideo
	Wall  APIWall
	Users APIUsers
}

// APIWithAccessToken creates and initializes a new API instance
func APIWithAccessToken(token string) *API {
	base := APIBase{
		AccessToken: token,
		Version:     APIVersion,
		BaseURL:     apiBaseURL,
	}

	return &API{
		APIBase: &base,
		Video:   APIVideo{&base},
		Wall:    APIWall{&base},
		Users:   APIUsers{&base},
	}
}

// APIError is a type representing errors returned by VK API
//easyjson:json
type APIError struct {
	Code          int    `json:"error_code"`
	Message       string `json:"error_msg"`
	RequestParams []struct {
		Key   string
		Value string
	}
}

// Error implements error interface
func (e *APIError) Error() string {
	return fmt.Sprintf("vk.APIError %d: %s", e.Code, e.Message)
}

// APIResponse is a type representing general response returned by VK API
//easyjson:json
type APIResponse struct {
	Error    *APIError
	Response easyjson.RawMessage
}

// Request performs an API request
// method is method name
// params should be a url.Values or url tagged
// struct (https://godoc.org/github.com/google/go-querystring/query)
func (vk *APIBase) Request(method string, params interface{}) (easyjson.RawMessage, error) {
	u, err := url.Parse(vk.BaseURL + method)
	if err != nil {
		return nil, err
	}

	var q url.Values

	switch v := params.(type) {
	case url.Values:
		q = v
	default:
		q, err = query.Values(params)
		if err != nil {
			return nil, err
		}
	}

	if baseParams, err := query.Values(vk); err == nil {
		urlValuesMerge(q, baseParams)
	} else {
		return nil, err
	}

	u.RawQuery = q.Encode()

	r, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	resp := APIResponse{}
	easyjson.UnmarshalFromReader(r.Body, &resp)

	if resp.Error != nil {
		return nil, resp.Error
	}

	return resp.Response, nil
}
