package vkCallbackApi

import (
	"net/url"
	"strings"

	"github.com/mailru/easyjson"
)

type UsersGetParams struct {
	UserIDs  []string
	Fields   []string
	NameCase string
}

//easyjson:json
type UsersGetResponse []User

func (vk *VKApi) UsersGet(params UsersGetParams) (UsersGetResponse, error) {
	method := "users.get"
	query := url.Values{}

	if v := params.UserIDs; len(v) != 0 {
		query.Set("user_ids", strings.Join(v, ","))
	}

	if v := params.Fields; len(v) != 0 {
		query.Set("fields", strings.Join(v, ","))
	}

	if v := params.NameCase; len(v) != 0 {
		query.Set("name_case", v)
	}

	r, err := vk.Request(method, query)
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
