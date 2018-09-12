package vkCallbackApi

import (
	"encoding/json"
)

// APINotifications implements VK API namespace `notifications`
type APINotifications struct {
	API *API
}

// NotificationsGetParams are params for APINotifications.Get
type NotificationsGetParams struct {
	// Number of notifications to return.
	Count     int    `url:"count,omitempty"`
	StartFrom string `url:"start_from,omitempty"`
	// Type of notifications to return: 'wall' — wall posts, 'mentions' — mentions in wall posts, comments, or topics, 'comments' — comments to wall posts, photos, and videos, 'likes' — likes, 'reposted' — wall posts that are copied from the current user's wall, 'followers' — new followers, 'friends' — accepted friend requests
	Filters CSVStringSlice `url:"filters,omitempty"`
	// Earliest timestamp (in Unix time) of a notification to return. By default, 24 hours ago.
	StartTime int `url:"start_time,omitempty"`
	// Latest timestamp (in Unix time) of a notification to return. By default, the current time.
	EndTime int `url:"end_time,omitempty"`
}

// NotificationsGetResponse is response for APINotifications.Get
//easyjson:json
type NotificationsGetResponse struct {
	// Total number
	Count    int `json:"count,omitempty"`
	Items    []genTODOType/* objects.json#/definitions/notifications_notification */ `json:"items,omitempty"`
	Profiles []User  `json:"profiles,omitempty"`
	Groups   []Group `json:"groups,omitempty"`
	// Time when user has been checked notifications last time
	LastViewed int `json:"last_viewed,omitempty"`
}

// Get Returns a list of notifications about other users' feedback to the current user's wall posts.
func (v APINotifications) Get(params NotificationsGetParams) (*NotificationsGetResponse, error) {
	r, err := v.API.Request("notifications.get", params)
	if err != nil {
		return nil, err
	}

	var resp NotificationsGetResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// MarkAsViewed Resets the counter of new notifications about other users' feedback to the current user's wall posts.
func (v APINotifications) MarkAsViewed() (bool, error) {
	r, err := v.API.Request("notifications.markAsViewed", nil)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}
