package vk

import (
	"encoding/json"
	"strconv"
)

// APIApps implements VK API namespace `apps`
type APIApps struct {
	API *API
}

// AppsGetCatalogParams are params for APIApps.GetCatalog
type AppsGetCatalogParams struct {
	// Sort order: 'popular_today' — popular for one day (default), 'visitors' — by visitors number , 'create_date' — by creation date, 'growth_rate' — by growth rate, 'popular_week' — popular for one week
	Sort string `url:"sort,omitempty"`
	// Offset required to return a specific subset of apps.
	Offset int `url:"offset,omitempty"`
	// Number of apps to return.
	Count    int    `url:"count"`
	Platform string `url:"platform,omitempty"`
	// '1' — to return additional fields 'screenshots', 'MAU', 'catalog_position', and 'international'. If set, 'count' must be less than or equal to '100'. '0' — not to return additional fields (default).
	Extended      bool           `url:"extended,omitempty"`
	ReturnFriends bool           `url:"return_friends,omitempty"`
	Fields        CSVStringSlice `url:"fields,omitempty"`
	NameCase      string         `url:"name_case,omitempty"`
	// Search query string.
	Q       string `url:"q,omitempty"`
	GenreID int    `url:"genre_id,omitempty"`
	// 'installed' — to return list of installed apps (only for mobile platform).
	Filter string `url:"filter,omitempty"`
}

// AppsGetCatalogResponse is response for APIApps.GetCatalog
//easyjson:json
type AppsGetCatalogResponse struct {
	// Total number
	Count int `json:"count,omitempty"`
	Items []genTODOType/* objects.json#/definitions/apps_app */ `json:"items,omitempty"`
}

// GetCatalog Returns a list of applications (apps) available to users in the App Catalog.
func (v APIApps) GetCatalog(params AppsGetCatalogParams) (*AppsGetCatalogResponse, error) {
	r, err := v.API.Request("apps.getCatalog", params)
	if err != nil {
		return nil, err
	}

	var resp AppsGetCatalogResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AppsGetParams are params for APIApps.Get
type AppsGetParams struct {
	// Application ID
	AppID int `url:"app_id,omitempty"`
	// List of application ID
	AppIDs CSVStringSlice `url:"app_ids,omitempty"`
	// platform. Possible values: *'ios' — iOS,, *'android' — Android,, *'winphone' — Windows Phone,, *'web' — приложения на vk.com. By default: 'web'.
	Platform string `url:"platform,omitempty"`
	// Profile fields to return. Sample values: 'nickname', 'screen_name', 'sex', 'bdate' (birthdate), 'city', 'country', 'timezone', 'photo', 'photo_medium', 'photo_big', 'has_mobile', 'contacts', 'education', 'online', 'counters', 'relation', 'last_seen', 'activity', 'can_write_private_message', 'can_see_all_posts', 'can_post', 'universities', (only if return_friends - 1)
	Fields CSVStringSlice `url:"fields,omitempty"`
	// Case for declension of user name and surname: 'nom' — nominative (default),, 'gen' — genitive,, 'dat' — dative,, 'acc' — accusative,, 'ins' — instrumental,, 'abl' — prepositional. (only if 'return_friends' = '1')
	NameCase string `url:"name_case,omitempty"`
}

// AppsGetResponse is response for APIApps.Get
//easyjson:json
type AppsGetResponse struct {
	// Total number
	Count int `json:"count,omitempty"`
	Items []genTODOType/* objects.json#/definitions/apps_app */ `json:"items,omitempty"`
}

// Get Returns applications data.
func (v APIApps) Get(params AppsGetParams) (*AppsGetResponse, error) {
	r, err := v.API.Request("apps.get", params)
	if err != nil {
		return nil, err
	}

	var resp AppsGetResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AppsSendRequestParams are params for APIApps.SendRequest
type AppsSendRequestParams struct {
	// id of the user to send a request
	UserID int `url:"user_id"`
	// request text
	Text string `url:"text,omitempty"`
	// request type. Values: 'invite' – if the request is sent to a user who does not have the app installed,, 'request' – if a user has already installed the app
	Type string `url:"type,omitempty"`
	Name string `url:"name,omitempty"`
	// special string key to be sent with the request
	Key      string `url:"key,omitempty"`
	Separate bool   `url:"separate,omitempty"`
}

// AppsSendRequestResponse is response for APIApps.SendRequest
// Request ID
type AppsSendRequestResponse int

// SendRequest Sends a request to another user in an app that uses VK authorization.
func (v APIApps) SendRequest(params AppsSendRequestParams) (AppsSendRequestResponse, error) {
	r, err := v.API.Request("apps.sendRequest", params)
	if err != nil {
		return 0, err
	}

	var resp AppsSendRequestResponse

	var cnv int
	cnv, err = strconv.Atoi(string(r))
	resp = AppsSendRequestResponse(cnv)

	if err != nil {
		return 0, err
	}
	return resp, nil
}

// DeleteAppRequests Deletes all request notifications from the current app.
func (v APIApps) DeleteAppRequests() (bool, error) {
	r, err := v.API.Request("apps.deleteAppRequests", nil)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// AppsGetFriendsListParams are params for APIApps.GetFriendsList
type AppsGetFriendsListParams struct {
	// List size.
	Count int `url:"count,omitempty"`
	// List type. Possible values: * 'invite' — available for invites (don't play the game),, * 'request' — available for request (play the game). By default: 'invite'.
	Type string `url:"type,omitempty"`
	// Additional profile fields, see [vk.com/dev/fields|description].
	Fields CSVStringSlice `url:"fields,omitempty"`
}

// AppsGetFriendsListResponse is response for APIApps.GetFriendsList
//easyjson:json
type AppsGetFriendsListResponse struct {
	// Total number
	Count int    `json:"count,omitempty"`
	Items []User `json:"items,omitempty"`
}

// GetFriendsList Creates friends list for requests and invites in current app.
func (v APIApps) GetFriendsList(params AppsGetFriendsListParams) (*AppsGetFriendsListResponse, error) {
	r, err := v.API.Request("apps.getFriendsList", params)
	if err != nil {
		return nil, err
	}

	var resp AppsGetFriendsListResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AppsGetLeaderboardParams are params for APIApps.GetLeaderboard
type AppsGetLeaderboardParams struct {
	// Leaderboard type. Possible values: *'level' — by level,, *'points' — by mission points,, *'score' — by score ().
	Type string `url:"type"`
	// Rating type. Possible values: *'1' — global rating among all players,, *'0' — rating among user friends.
	Global bool `url:"global,omitempty"`
	// 1 — to return additional info about users
	Extended bool `url:"extended,omitempty"`
}

// AppsGetLeaderboardResponse is response for APIApps.GetLeaderboard
// Either AppsGetLeaderboardResponseNormal or AppsGetLeaderboardResponseExtended, depending on Extended flag
type AppsGetLeaderboardResponse interface {
	isAppsGetLeaderboard()
}

// AppsGetLeaderboardResponseNormal is non-extended version of AppsGetLeaderboardResponse
//easyjson:json
type AppsGetLeaderboardResponseNormal struct {
	// Total number
	Count int `json:"count,omitempty"`
	Items []struct {
		// Score number
		Score int `json:"score,omitempty"`
		// Level
		Level int `json:"level,omitempty"`
		// Points number
		Points int `json:"points,omitempty"`
		// User ID
		UserID int `json:"user_id,omitempty"`
	} `json:"items,omitempty"`
}

func (AppsGetLeaderboardResponseNormal) isAppsGetLeaderboard() {}

// AppsGetLeaderboardResponseExtended is extended version of AppsGetLeaderboardResponse
//easyjson:json
type AppsGetLeaderboardResponseExtended struct {
	// Total number
	Count int `json:"count,omitempty"`
	Items []struct {
		// Score number
		Score int `json:"score,omitempty"`
		// Level
		Level int `json:"level,omitempty"`
		// Points number
		Points int `json:"points,omitempty"`
		// User ID
		UserID int `json:"user_id,omitempty"`
	} `json:"items,omitempty"`
	Profiles []User `json:"profiles,omitempty"`
}

func (AppsGetLeaderboardResponseExtended) isAppsGetLeaderboard() {}

// GetLeaderboard Returns players rating in the game.
func (v APIApps) GetLeaderboard(params AppsGetLeaderboardParams) (AppsGetLeaderboardResponse, error) {
	r, err := v.API.Request("apps.getLeaderboard", params)
	if err != nil {
		return nil, err
	}

	var resp AppsGetLeaderboardResponse
	if params.Extended {
		var tmp AppsGetLeaderboardResponseExtended
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	} else {
		var tmp AppsGetLeaderboardResponseNormal
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// AppsGetScoreParams are params for APIApps.GetScore
type AppsGetScoreParams struct {
	UserID int `url:"user_id"`
}

// AppsGetScoreResponse is response for APIApps.GetScore
// Score number
type AppsGetScoreResponse int

// GetScore Returns user score in app
func (v APIApps) GetScore(params AppsGetScoreParams) (AppsGetScoreResponse, error) {
	r, err := v.API.Request("apps.getScore", params)
	if err != nil {
		return 0, err
	}

	var resp AppsGetScoreResponse

	var cnv int
	cnv, err = strconv.Atoi(string(r))
	resp = AppsGetScoreResponse(cnv)

	if err != nil {
		return 0, err
	}
	return resp, nil
}
