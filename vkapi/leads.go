package vkapi

import (
	"encoding/json"

	"github.com/stek29/vk"
)

// Leads implements VK API namespace `leads`
type Leads struct {
	API vk.API
}

// LeadsCompleteParams are params for Leads.Complete
type LeadsCompleteParams struct {
	// Session obtained as GET parameter when session started.
	VkSid string `url:"vk_sid"`
	// Secret key from the lead testing interface.
	Secret string `url:"secret"`
	// Comment text.
	Comment string `url:"comment,omitempty"`
}

// LeadsCompleteResponse is response for Leads.Complete
//easyjson:json
type LeadsCompleteResponse struct {
	// Offer limit
	Limit int `json:"limit,omitempty"`
	// Amount of spent votes
	Spent int `json:"spent,omitempty"`
	// Offer cost
	Cost int `json:"cost,omitempty"`
	// Information whether test mode is enabled
	TestMode vk.BoolInt `json:"test_mode,omitempty"`
	Success  vk.BoolInt `json:"success,omitempty"`
}

// Complete Completes the lead started by user.
func (v Leads) Complete(params LeadsCompleteParams) (*LeadsCompleteResponse, error) {
	r, err := v.API.Request("leads.complete", params)
	if err != nil {
		return nil, err
	}

	var resp LeadsCompleteResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// LeadsStartParams are params for Leads.Start
type LeadsStartParams struct {
	// Lead ID.
	LeadID int `url:"lead_id"`
	// Secret key from the lead testing interface.
	Secret   string `url:"secret"`
	UID      int    `url:"uid,omitempty"`
	Aid      int    `url:"aid,omitempty"`
	TestMode bool   `url:"test_mode,omitempty"`
	Force    bool   `url:"force,omitempty"`
}

// LeadsStartResponse is response for Leads.Start
//easyjson:json
type LeadsStartResponse struct {
	// Information whether test mode is enabled
	TestMode vk.BoolInt `json:"test_mode,omitempty"`
	// Session data
	VkSid string `json:"vk_sid,omitempty"`
}

// Start Creates new session for the user passing the offer.
func (v Leads) Start(params LeadsStartParams) (*LeadsStartResponse, error) {
	r, err := v.API.Request("leads.start", params)
	if err != nil {
		return nil, err
	}

	var resp LeadsStartResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// LeadsGetStatsParams are params for Leads.GetStats
type LeadsGetStatsParams struct {
	// Lead ID.
	LeadID int `url:"lead_id"`
	// Secret key obtained from the lead testing interface.
	Secret string `url:"secret,omitempty"`
	// Day to start stats from (YYYY_MM_DD, e.g.2011-09-17).
	DateStart string `url:"date_start,omitempty"`
	// Day to finish stats (YYYY_MM_DD, e.g.2011-09-17).
	DateEnd string `url:"date_end,omitempty"`
}

// LeadsGetStatsResponse is response for Leads.GetStats
//easyjson:json
type LeadsGetStatsResponse struct {
	// Lead limit
	Limit int `json:"limit,omitempty"`
	// Amount of spent votes
	Spent int `json:"spent,omitempty"`
	// Offer cost
	Cost int `json:"cost,omitempty"`
	// Impressions number
	Impressions int `json:"impressions,omitempty"`
	// Started offers number
	Started int `json:"started,omitempty"`
	// Completed offers number
	Completed int `json:"completed,omitempty"`
	Days      struct {
		// Impressions number
		Impressions int `json:"impressions,omitempty"`
		// Started offers number
		Started int `json:"started,omitempty"`
		// Completed offers number
		Completed int `json:"completed,omitempty"`
		// Amount of spent votes
		Spent int `json:"spent,omitempty"`
	} `json:"days,omitempty"`
}

// GetStats Returns lead stats data.
func (v Leads) GetStats(params LeadsGetStatsParams) (*LeadsGetStatsResponse, error) {
	r, err := v.API.Request("leads.getStats", params)
	if err != nil {
		return nil, err
	}

	var resp LeadsGetStatsResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// LeadsGetUsersParams are params for Leads.GetUsers
type LeadsGetUsersParams struct {
	// Offer ID.
	OfferID int `url:"offer_id"`
	// Secret key obtained in the lead testing interface.
	Secret string `url:"secret"`
	// Offset needed to return a specific subset of results.
	Offset int `url:"offset,omitempty"`
	// Number of results to return.
	Count int `url:"count,omitempty"`
	// Action type. Possible values: *'0' — start,, *'1' — finish,, *'2' — blocking users,, *'3' — start in a test mode,, *'4' — finish in a test mode.
	Status int `url:"status,omitempty"`
	// Sort order. Possible values: *'1' — chronological,, *'0' — reverse chronological.
	Reverse bool `url:"reverse,omitempty"`
}

// LeadsGetUsersResponse is response for Leads.GetUsers
//easyjson:json
type LeadsGetUsersResponse []struct {
	// User ID
	UID int `json:"uid,omitempty"`
	// Application ID
	Aid int `json:"aid,omitempty"`
	// Session string ID
	Sid string `json:"sid,omitempty"`
	// Date when the action has been started in Unixtime
	Date int `json:"date,omitempty"`
	// Action type
	Status int `json:"status,omitempty"`
	// Information whether test mode is enabled
	TestMode vk.BoolInt `json:"test_mode,omitempty"`
	// Start date in Unixtime (for status=2)
	StartDate int `json:"start_date,omitempty"`
	// Comment text
	Comment string `json:"comment,omitempty"`
}

// GetUsers Returns a list of last user actions for the offer.
func (v Leads) GetUsers(params LeadsGetUsersParams) (LeadsGetUsersResponse, error) {
	r, err := v.API.Request("leads.getUsers", params)
	if err != nil {
		return nil, err
	}

	var resp LeadsGetUsersResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// LeadsCheckUserParams are params for Leads.CheckUser
type LeadsCheckUserParams struct {
	// Lead ID.
	LeadID int `url:"lead_id"`
	// Value to be return in 'result' field when test mode is used.
	TestResult int  `url:"test_result,omitempty"`
	TestMode   bool `url:"test_mode,omitempty"`
	AutoStart  bool `url:"auto_start,omitempty"`
	// User age.
	Age int `url:"age,omitempty"`
	// User country code.
	Country string `url:"country,omitempty"`
}

// LeadsCheckUserResponse is response for Leads.CheckUser
//easyjson:json
type LeadsCheckUserResponse struct {
	// Information whether user can start the lead
	Result string `json:"result,omitempty"`
	// Reason why user can't start the lead
	Reason string `json:"reason,omitempty"`
	// URL user should open to start the lead
	StartLink string `json:"start_link,omitempty"`
	// Session ID
	Sid string `json:"sid,omitempty"`
}

// CheckUser Checks if the user can start the lead.
func (v Leads) CheckUser(params LeadsCheckUserParams) (*LeadsCheckUserResponse, error) {
	r, err := v.API.Request("leads.checkUser", params)
	if err != nil {
		return nil, err
	}

	var resp LeadsCheckUserResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// LeadsMetricHitParams are params for Leads.MetricHit
type LeadsMetricHitParams struct {
	// Metric data obtained in the lead interface.
	Data string `url:"data"`
}

// LeadsMetricHitResponse is response for Leads.MetricHit
//easyjson:json
type LeadsMetricHitResponse struct {
	// Information whether request has been processed successfully
	Result bool `json:"result,omitempty"`
	// Redirect link
	RedirectLink string `json:"redirect_link,omitempty"`
}

// MetricHit Counts the metric event.
func (v Leads) MetricHit(params LeadsMetricHitParams) (*LeadsMetricHitResponse, error) {
	r, err := v.API.Request("leads.metricHit", params)
	if err != nil {
		return nil, err
	}

	var resp LeadsMetricHitResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
