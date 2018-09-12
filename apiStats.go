package vkCallbackApi

import (
	"encoding/json"
)

// APIStats implements VK API namespace `stats`
type APIStats struct {
	API *API
}

// StatsGetParams are params for APIStats.Get
type StatsGetParams struct {
	// Community ID.
	GroupID int `url:"group_id,omitempty"`
	// Application ID.
	AppID int `url:"app_id,omitempty"`
	// Latest datestamp (in Unix time) of statistics to return.
	DateFrom string `url:"date_from,omitempty"`
	// End datestamp (in Unix time) of statistics to return.
	DateTo string `url:"date_to,omitempty"`
}

// StatsGetResponse is response for APIStats.Get
//easyjson:json
type StatsGetResponse []genTODOType /* objects.json#/definitions/stats_period */
// Get Returns statistics of a community or an application.
func (v APIStats) Get(params StatsGetParams) (StatsGetResponse, error) {
	r, err := v.API.Request("stats.get", params)
	if err != nil {
		return nil, err
	}

	var resp StatsGetResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// TrackVisitor does stats.trackVisitor
func (v APIStats) TrackVisitor() (bool, error) {
	r, err := v.API.Request("stats.trackVisitor", nil)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// StatsGetPostReachParams are params for APIStats.GetPostReach
type StatsGetPostReachParams struct {
	// post owner community id. Specify with "-" sign.
	OwnerID int `url:"owner_id"`
	// wall post id. Note that stats are available only for '300' last (newest) posts on a community wall.
	PostID int `url:"post_id"`
}

// StatsGetPostReachResponse is response for APIStats.GetPostReach
//easyjson:json
type StatsGetPostReachResponse []WallpostStats

// GetPostReach Returns stats for a wall post.
func (v APIStats) GetPostReach(params StatsGetPostReachParams) (StatsGetPostReachResponse, error) {
	r, err := v.API.Request("stats.getPostReach", params)
	if err != nil {
		return nil, err
	}

	var resp StatsGetPostReachResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
