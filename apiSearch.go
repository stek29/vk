package vkCallbackApi

import (
	"encoding/json"
)

// APISearch implements VK API namespace `search`
type APISearch struct {
	API *API
}

// SearchGetHintsParams are params for APISearch.GetHints
type SearchGetHintsParams struct {
	// Search query string.
	Q string `url:"q,omitempty"`
	// Offset for querying specific result subset
	Offset int `url:"offset,omitempty"`
	// Maximum number of results to return.
	Limit int `url:"limit,omitempty"`
	//
	Filters CSVStringSlice `url:"filters,omitempty"`
	//
	SearchGlobal bool `url:"search_global,omitempty"`
}

// SearchGetHintsResponse is response for APISearch.GetHints
//easyjson:json
type SearchGetHintsResponse struct {
	Items []struct {
		// Object type
		Type string `json:"type,omitempty"`
		// Section title
		Section string `json:"section,omitempty"`
		// Object description
		Description string `json:"description,omitempty"`
		// Information whether the object has been found globally
		Global  BoolInt `json:"global,omitempty"`
		Group   Group   `json:"group,omitempty"`
		Profile User    `json:"profile,omitempty"`
	} `json:"items,omitempty"`
	SuggestedQueries []string `json:"suggested_queries,omitempty"`
}

// GetHints Allows the programmer to do a quick search for any substring.
func (v APISearch) GetHints(params SearchGetHintsParams) (*SearchGetHintsResponse, error) {
	r, err := v.API.Request("search.getHints", params)
	if err != nil {
		return nil, err
	}

	var resp SearchGetHintsResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
