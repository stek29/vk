package vkapi

import (
	"encoding/json"

	"github.com/stek29/vk"
)

// Stats implements VK API namespace `stats`
type Stats struct {
	API vk.API
}

// StatsGetParams are params for Stats.Get
type StatsGetParams struct {
	// Community ID.
	GroupID int `url:"group_id,omitempty"`
	// Application ID.
	AppID          int            `url:"app_id,omitempty"`
	TimestampFrom  int            `url:"timestamp_from,omitempty"`
	TimestampTo    int            `url:"timestamp_to,omitempty"`
	Interval       string         `url:"interval,omitempty"`
	IntervalsCount int            `url:"intervals_count,omitempty"`
	Filters        CSVStringSlice `url:"filters,omitempty"`
	StatsGroups    CSVStringSlice `url:"stats_groups,omitempty"`
	Extended       bool           `url:"extended,omitempty"`
}

// StatsGetResponse is response for Stats.Get
//easyjson:json
type StatsGetResponse []genTODOType /* objects.json#/definitions/stats_period */
// Get Returns statistics of a community or an application.
func (v Stats) Get(params StatsGetParams) (StatsGetResponse, error) {
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

// StatsTrackVisitorParams are params for Stats.TrackVisitor
type StatsTrackVisitorParams struct {
	ID string `url:"id"`
}

// TrackVisitor does stats.trackVisitor
func (v Stats) TrackVisitor(params StatsTrackVisitorParams) (bool, error) {
	r, err := v.API.Request("stats.trackVisitor", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// StatsGetPostReachParams are params for Stats.GetPostReach
type StatsGetPostReachParams struct {
	// post owner community id. Specify with "-" sign.
	OwnerID int `url:"owner_id"`
	// wall post id. Note that stats are available only for '300' last (newest) posts on a community wall.
	PostID int `url:"post_id"`
}

// StatsGetPostReachResponse is response for Stats.GetPostReach
//easyjson:json
type StatsGetPostReachResponse []vk.WallpostStats

// GetPostReach Returns stats for a wall post.
func (v Stats) GetPostReach(params StatsGetPostReachParams) (StatsGetPostReachResponse, error) {
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
