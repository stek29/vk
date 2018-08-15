package vkCallbackApi

import (
	"github.com/mailru/easyjson"
)

// APIGroups implements VK API namespace `groups`
type APIGroups struct {
	API *API
}

// GroupsGetByIDParams are params for Users.Get
// I don't know why VK keeps both group_ids and group_id
type GroupsGetByIDParams struct {
	// Up to 500 IDs or screen names of communities.
	GroupIDs CSVStringSlice `url:"group_ids,omitempty"`
	// ID or screen name of the community.
	GroupID string `url:"group_id,omitempty"`
	// Additional group fields to return. See Group type.
	Fields CSVStringSlice `url:"fields,omitempty"`
}

// GroupsGetByIDResponse is response for Users.Get
//easyjson:json
type GroupsGetByIDResponse []Group

// GetByID is groups.getById
// Returns information about communities by their IDs.
func (v APIGroups) GetByID(params GroupsGetByIDParams) (GroupsGetByIDResponse, error) {
	r, err := v.API.Request("groups.getById", params)
	if err != nil {
		return nil, err
	}

	resp := GroupsGetByIDResponse{}
	err = easyjson.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
