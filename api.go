package vk

import (
	"encoding/json"
	"fmt"
	"net/url"

	"net/http"

	"github.com/google/go-querystring/query"
)

const (
	apiBaseURL = "https://api.vk.com/method/"
)

type API interface {
	Request(method string, params interface{}) (json.RawMessage, error)
}

// BaseAPI is a helper type used for making requests
type BaseAPI struct {
	BaseURL     string `url:"-"`
	AccessToken string `url:"access_token,omitempty"`
	Version     string `url:"v,omitempty"`
	Language    string `url:"lang,omitempty"`
}

// BaseAPIWithAccessToken creates and initializes a new API instance
func BaseAPIWithAccessToken(token string) API {
	return &BaseAPI{
		AccessToken: token,
		Version:     APIVersion,
		BaseURL:     apiBaseURL,
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
	Response json.RawMessage
}

// Request performs an API request
// method is method name
// params should be nil, url.Values or an url tagged
// struct (https://godoc.org/github.com/google/go-querystring/query)
func (vk *BaseAPI) Request(method string, params interface{}) (json.RawMessage, error) {
	u, err := url.Parse(vk.BaseURL + method)
	if err != nil {
		return nil, err
	}

	var q url.Values

	switch v := params.(type) {
	case nil:
		q = make(url.Values)
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

	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&resp); err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, resp.Error
	}

	return resp.Response, nil
}
