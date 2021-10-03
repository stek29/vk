package vkapi

import (
	"encoding/json"

	"github.com/stek29/vk"
)

// Search implements VK API namespace `search`
type Search struct {
	API vk.API
}

// SearchGetHintsParams are params for Search.GetHints
type SearchGetHintsParams struct {
	// Search query string.
	Q string `url:"q,omitempty"`
	// Offset for querying specific result subset
	Offset int `url:"offset,omitempty"`
	// Maximum number of results to return.
	Limit        int            `url:"limit,omitempty"`
	Filters      CSVStringSlice `url:"filters,omitempty"`
	Fields       CSVStringSlice `url:"fields,omitempty"`
	SearchGlobal bool           `url:"search_global,omitempty"`
}

// SearchGetHintsResponse is response for Search.GetHints
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
		Global  vk.BoolInt `json:"global,omitempty"`
		Group   vk.Group   `json:"group,omitempty"`
		Profile vk.User    `json:"profile,omitempty"`
	} `json:"items,omitempty"`
	SuggestedQueries []string `json:"suggested_queries,omitempty"`
}

// GetHints Allows the programmer to do a quick search for any substring.
func (v Search) GetHints(params SearchGetHintsParams) (*SearchGetHintsResponse, error) {
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
