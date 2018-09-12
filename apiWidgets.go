package vkCallbackApi

import (
	"github.com/mailru/easyjson"
)

// APIWidgets implements VK API namespace `widgets`
type APIWidgets struct {
	API *API
}

// WidgetsGetCommentsParams are params for APIWidgets.GetComments
type WidgetsGetCommentsParams struct {
	WidgetAPIID int            `url:"widget_api_id,omitempty"`
	URL         string         `url:"url,omitempty"`
	PageID      string         `url:"page_id,omitempty"`
	Order       string         `url:"order,omitempty"`
	Fields      CSVStringSlice `url:"fields,omitempty"`
	Count       int            `url:"count,omitempty"`
}

// WidgetsGetCommentsResponse is response for APIWidgets.GetComments
//easyjson:json
type WidgetsGetCommentsResponse struct {
	// Total number
	Count int `json:"count,omitempty"`
	Posts []genTODOType/* objects.json#/definitions/widgets_widget_comment */ `json:"posts,omitempty"`
}

// GetComments Gets a list of comments for the page added through the [vk.com/dev/Comments|Comments widget].
func (v APIWidgets) GetComments(params WidgetsGetCommentsParams) (*WidgetsGetCommentsResponse, error) {
	r, err := v.API.Request("widgets.getComments", params)
	if err != nil {
		return nil, err
	}

	var resp WidgetsGetCommentsResponse
	err = easyjson.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// WidgetsGetPagesParams are params for APIWidgets.GetPages
type WidgetsGetPagesParams struct {
	WidgetAPIID int    `url:"widget_api_id,omitempty"`
	Order       string `url:"order,omitempty"`
	Period      string `url:"period,omitempty"`
	Count       int    `url:"count,omitempty"`
}

// WidgetsGetPagesResponse is response for APIWidgets.GetPages
//easyjson:json
type WidgetsGetPagesResponse struct {
	// Total number
	Count int `json:"count,omitempty"`
	Pages []genTODOType/* objects.json#/definitions/widgets_widget_page */ `json:"pages,omitempty"`
}

// GetPages Gets a list of application/site pages where the [vk.com/dev/Comments|Comments widget] or [vk.com/dev/Like|Like widget] is installed.
func (v APIWidgets) GetPages(params WidgetsGetPagesParams) (*WidgetsGetPagesResponse, error) {
	r, err := v.API.Request("widgets.getPages", params)
	if err != nil {
		return nil, err
	}

	var resp WidgetsGetPagesResponse
	err = easyjson.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
