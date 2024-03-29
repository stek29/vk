package vkapi

import (
	"encoding/json"

	"github.com/stek29/vk"
)

// Stories implements VK API namespace `stories`
type Stories struct {
	API vk.API
}

// StoriesBanOwnerParams are params for Stories.BanOwner
type StoriesBanOwnerParams struct {
	// List of sources IDs
	OwnersIDs CSVIntSlice `url:"owners_ids"`
}

// BanOwner Allows to hide stories from chosen sources from current user's feed.
func (v Stories) BanOwner(params StoriesBanOwnerParams) (bool, error) {
	r, err := v.API.Request("stories.banOwner", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// StoriesDeleteParams are params for Stories.Delete
type StoriesDeleteParams struct {
	// Story owner's ID. Current user id is used by default.
	OwnerID int `url:"owner_id"`
	// Story ID.
	StoryID int `url:"story_id"`
}

// Delete Allows to delete story.
func (v Stories) Delete(params StoriesDeleteParams) (bool, error) {
	r, err := v.API.Request("stories.delete", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// StoriesGetParams are params for Stories.Get
type StoriesGetParams struct {
	// Owner ID.
	OwnerID int `url:"owner_id,omitempty"`
	// '1' — to return additional fields for users and communities. Default value is 0.
	Extended bool `url:"extended,omitempty"`
}

// StoriesGetResponse is response for Stories.Get
// Either StoriesGetResponseNormal or StoriesGetResponseExtended, depending on Extended flag
type StoriesGetResponse interface {
	isStoriesGet()
}

// StoriesGetResponseNormal is non-extended version of StoriesGetResponse
//easyjson:json
type StoriesGetResponseNormal struct {
	// Stories count
	Count int          `json:"count,omitempty"`
	Items [][]vk.Story `json:"items,omitempty"`
}

func (StoriesGetResponseNormal) isStoriesGet() {}

// StoriesGetResponseExtended is extended version of StoriesGetResponse
//easyjson:json
type StoriesGetResponseExtended struct {
	// Stories count
	Count    int          `json:"count,omitempty"`
	Items    [][]vk.Story `json:"items,omitempty"`
	Profiles []vk.User    `json:"profiles,omitempty"`
	Groups   []vk.Group   `json:"groups,omitempty"`
}

func (StoriesGetResponseExtended) isStoriesGet() {}

// Get Returns stories available for current user.
func (v Stories) Get(params StoriesGetParams) (StoriesGetResponse, error) {
	r, err := v.API.Request("stories.get", params)
	if err != nil {
		return nil, err
	}

	var resp StoriesGetResponse
	if params.Extended {
		var tmp StoriesGetResponseExtended
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	} else {
		var tmp StoriesGetResponseNormal
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// StoriesGetBannedParams are params for Stories.GetBanned
type StoriesGetBannedParams struct {
	// '1' — to return additional fields for users and communities. Default value is 0.
	Extended bool `url:"extended,omitempty"`
	// Additional fields to return
	Fields CSVStringSlice `url:"fields,omitempty"`
}

// StoriesGetBannedResponse is response for Stories.GetBanned
// Either StoriesGetBannedResponseNormal or StoriesGetBannedResponseExtended, depending on Extended flag
type StoriesGetBannedResponse interface {
	isStoriesGetBanned()
}

// StoriesGetBannedResponseNormal is non-extended version of StoriesGetBannedResponse
//easyjson:json
type StoriesGetBannedResponseNormal struct {
	// Stories count
	Count int `json:"count,omitempty"`
	// Owner ID
	Items []int `json:"items,omitempty"`
}

func (StoriesGetBannedResponseNormal) isStoriesGetBanned() {}

// StoriesGetBannedResponseExtended is extended version of StoriesGetBannedResponse
//easyjson:json
type StoriesGetBannedResponseExtended struct {
	// Stories count
	Count int `json:"count,omitempty"`
	// Owner ID
	Items    []int      `json:"items,omitempty"`
	Profiles []vk.User  `json:"profiles,omitempty"`
	Groups   []vk.Group `json:"groups,omitempty"`
}

func (StoriesGetBannedResponseExtended) isStoriesGetBanned() {}

// GetBanned Returns list of sources hidden from current user's feed.
func (v Stories) GetBanned(params StoriesGetBannedParams) (StoriesGetBannedResponse, error) {
	r, err := v.API.Request("stories.getBanned", params)
	if err != nil {
		return nil, err
	}

	var resp StoriesGetBannedResponse
	if params.Extended {
		var tmp StoriesGetBannedResponseExtended
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	} else {
		var tmp StoriesGetBannedResponseNormal
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// StoriesGetByIDParams are params for Stories.GetByID
type StoriesGetByIDParams struct {
	// Stories IDs separated by commas. Use format {owner_id}+'_'+{story_id}, for example, 12345_54331.
	Stories CSVStringSlice `url:"stories"`
	// '1' — to return additional fields for users and communities. Default value is 0.
	Extended bool `url:"extended,omitempty"`
	// Additional fields to return
	Fields CSVStringSlice `url:"fields,omitempty"`
}

// StoriesGetByIDResponse is response for Stories.GetByID
// Either StoriesGetByIDResponseNormal or StoriesGetByIDResponseExtended, depending on Extended flag
type StoriesGetByIDResponse interface {
	isStoriesGetByID()
}

// StoriesGetByIDResponseNormal is non-extended version of StoriesGetByIDResponse
//easyjson:json
type StoriesGetByIDResponseNormal struct {
	// Stories count
	Count int        `json:"count,omitempty"`
	Items []vk.Story `json:"items,omitempty"`
}

func (StoriesGetByIDResponseNormal) isStoriesGetByID() {}

// StoriesGetByIDResponseExtended is extended version of StoriesGetByIDResponse
//easyjson:json
type StoriesGetByIDResponseExtended struct {
	// Stories count
	Count    int        `json:"count,omitempty"`
	Items    []vk.Story `json:"items,omitempty"`
	Profiles []vk.User  `json:"profiles,omitempty"`
	Groups   []vk.Group `json:"groups,omitempty"`
}

func (StoriesGetByIDResponseExtended) isStoriesGetByID() {}

// GetByID Returns story by its ID.
func (v Stories) GetByID(params StoriesGetByIDParams) (StoriesGetByIDResponse, error) {
	r, err := v.API.Request("stories.getById", params)
	if err != nil {
		return nil, err
	}

	var resp StoriesGetByIDResponse
	if params.Extended {
		var tmp StoriesGetByIDResponseExtended
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	} else {
		var tmp StoriesGetByIDResponseNormal
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// StoriesGetPhotoUploadServerParams are params for Stories.GetPhotoUploadServer
type StoriesGetPhotoUploadServerParams struct {
	// 1 — to add the story to friend's feed.
	AddToNews bool `url:"add_to_news,omitempty"`
	// List of users IDs who can see the story.
	UserIDs CSVIntSlice `url:"user_ids,omitempty"`
	// ID of the story to reply with the current.
	ReplyToStory string `url:"reply_to_story,omitempty"`
	// Link text (for community's stories only).
	LinkText string `url:"link_text,omitempty"`
	// Link URL. Internal links on https://vk.com only.
	LinkURL string `url:"link_url,omitempty"`
	// ID of the community to upload the story (should be verified or with the "fire" icon).
	GroupID int `url:"group_id,omitempty"`
}

// StoriesGetPhotoUploadServerResponse is response for Stories.GetPhotoUploadServer
//easyjson:json
type StoriesGetPhotoUploadServerResponse struct {
	// Upload URL
	UploadURL string `json:"upload_url,omitempty"`
	// Users ID who can to see story.
	UserIDs []int `json:"user_ids,omitempty"`
}

// GetPhotoUploadServer Returns URL for uploading a story with photo.
func (v Stories) GetPhotoUploadServer(params StoriesGetPhotoUploadServerParams) (*StoriesGetPhotoUploadServerResponse, error) {
	r, err := v.API.Request("stories.getPhotoUploadServer", params)
	if err != nil {
		return nil, err
	}

	var resp StoriesGetPhotoUploadServerResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// StoriesGetRepliesParams are params for Stories.GetReplies
type StoriesGetRepliesParams struct {
	// Story owner ID.
	OwnerID int `url:"owner_id"`
	// Story ID.
	StoryID int `url:"story_id"`
	// Access key for the private object.
	AccessKey string `url:"access_key,omitempty"`
	// '1' — to return additional fields for users and communities. Default value is 0.
	Extended bool `url:"extended,omitempty"`
	// Additional fields to return
	Fields CSVStringSlice `url:"fields,omitempty"`
}

// StoriesGetRepliesResponse is response for Stories.GetReplies
// Either StoriesGetRepliesResponseNormal or StoriesGetRepliesResponseExtended, depending on Extended flag
type StoriesGetRepliesResponse interface {
	isStoriesGetReplies()
}

// StoriesGetRepliesResponseNormal is non-extended version of StoriesGetRepliesResponse
//easyjson:json
type StoriesGetRepliesResponseNormal struct {
	// Stories count
	Count int          `json:"count,omitempty"`
	Items [][]vk.Story `json:"items,omitempty"`
}

func (StoriesGetRepliesResponseNormal) isStoriesGetReplies() {}

// StoriesGetRepliesResponseExtended is extended version of StoriesGetRepliesResponse
//easyjson:json
type StoriesGetRepliesResponseExtended struct {
	// Stories count
	Count    int          `json:"count,omitempty"`
	Items    [][]vk.Story `json:"items,omitempty"`
	Profiles []vk.User    `json:"profiles,omitempty"`
	Groups   []vk.Group   `json:"groups,omitempty"`
}

func (StoriesGetRepliesResponseExtended) isStoriesGetReplies() {}

// GetReplies Returns replies to the story.
func (v Stories) GetReplies(params StoriesGetRepliesParams) (StoriesGetRepliesResponse, error) {
	r, err := v.API.Request("stories.getReplies", params)
	if err != nil {
		return nil, err
	}

	var resp StoriesGetRepliesResponse
	if params.Extended {
		var tmp StoriesGetRepliesResponseExtended
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	} else {
		var tmp StoriesGetRepliesResponseNormal
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// StoriesGetStatsParams are params for Stories.GetStats
type StoriesGetStatsParams struct {
	// Story owner ID.
	OwnerID int `url:"owner_id"`
	// Story ID.
	StoryID int `url:"story_id"`
}

// StoriesGetStatsResponse is response for Stories.GetStats
//easyjson:json
type StoriesGetStatsResponse struct {
	Views struct {
		// Statistic state
		State string `json:"state,omitempty"`
		// Stat value
		Count int `json:"count,omitempty"`
	} `json:"views,omitempty"`
	Replies struct {
		// Statistic state
		State string `json:"state,omitempty"`
		// Stat value
		Count int `json:"count,omitempty"`
	} `json:"replies,omitempty"`
	Answer struct {
		// Statistic state
		State string `json:"state,omitempty"`
		// Stat value
		Count int `json:"count,omitempty"`
	} `json:"answer,omitempty"`
	Shares struct {
		// Statistic state
		State string `json:"state,omitempty"`
		// Stat value
		Count int `json:"count,omitempty"`
	} `json:"shares,omitempty"`
	Subscribers struct {
		// Statistic state
		State string `json:"state,omitempty"`
		// Stat value
		Count int `json:"count,omitempty"`
	} `json:"subscribers,omitempty"`
	Bans struct {
		// Statistic state
		State string `json:"state,omitempty"`
		// Stat value
		Count int `json:"count,omitempty"`
	} `json:"bans,omitempty"`
	OpenLink struct {
		// Statistic state
		State string `json:"state,omitempty"`
		// Stat value
		Count int `json:"count,omitempty"`
	} `json:"open_link,omitempty"`
}

// GetStats Returns stories available for current user.
func (v Stories) GetStats(params StoriesGetStatsParams) (*StoriesGetStatsResponse, error) {
	r, err := v.API.Request("stories.getStats", params)
	if err != nil {
		return nil, err
	}

	var resp StoriesGetStatsResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// StoriesGetVideoUploadServerParams are params for Stories.GetVideoUploadServer
type StoriesGetVideoUploadServerParams struct {
	// 1 — to add the story to friend's feed.
	AddToNews bool `url:"add_to_news,omitempty"`
	// List of users IDs who can see the story.
	UserIDs CSVIntSlice `url:"user_ids,omitempty"`
	// ID of the story to reply with the current.
	ReplyToStory string `url:"reply_to_story,omitempty"`
	// Link text (for community's stories only).
	LinkText string `url:"link_text,omitempty"`
	// Link URL. Internal links on https://vk.com only.
	LinkURL string `url:"link_url,omitempty"`
	// ID of the community to upload the story (should be verified or with the "fire" icon).
	GroupID int `url:"group_id,omitempty"`
}

// StoriesGetVideoUploadServerResponse is response for Stories.GetVideoUploadServer
//easyjson:json
type StoriesGetVideoUploadServerResponse struct {
	// Upload URL
	UploadURL string `json:"upload_url,omitempty"`
	// Users ID who can to see story.
	UserIDs []int `json:"user_ids,omitempty"`
}

// GetVideoUploadServer Allows to receive URL for uploading story with video.
func (v Stories) GetVideoUploadServer(params StoriesGetVideoUploadServerParams) (*StoriesGetVideoUploadServerResponse, error) {
	r, err := v.API.Request("stories.getVideoUploadServer", params)
	if err != nil {
		return nil, err
	}

	var resp StoriesGetVideoUploadServerResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// StoriesGetViewersParams are params for Stories.GetViewers
type StoriesGetViewersParams struct {
	// Story owner ID.
	OwnerID int `url:"owner_id"`
	// Story ID.
	StoryID int `url:"story_id"`
	// Maximum number of results.
	Count int `url:"count,omitempty"`
	// Offset needed to return a specific subset of results.
	Offset int `url:"offset,omitempty"`
	// '1' — to return detailed information about photos
	Extended bool `url:"extended,omitempty"`
}

// StoriesGetViewersResponse is response for Stories.GetViewers
// Either StoriesGetViewersResponseNormal or StoriesGetViewersResponseExtended, depending on Extended flag
type StoriesGetViewersResponse interface {
	isStoriesGetViewers()
}

// StoriesGetViewersResponseNormal is non-extended version of StoriesGetViewersResponse
//easyjson:json
type StoriesGetViewersResponseNormal struct {
	// Viewers count
	Count int   `json:"count,omitempty"`
	Items []int `json:"items,omitempty"`
}

func (StoriesGetViewersResponseNormal) isStoriesGetViewers() {}

// StoriesGetViewersResponseExtended is extended version of StoriesGetViewersResponse
//easyjson:json
type StoriesGetViewersResponseExtended struct {
	// Viewers count
	Count int       `json:"count,omitempty"`
	Items []vk.User `json:"items,omitempty"`
}

func (StoriesGetViewersResponseExtended) isStoriesGetViewers() {}

// GetViewers Returns a list of story viewers.
func (v Stories) GetViewers(params StoriesGetViewersParams) (StoriesGetViewersResponse, error) {
	r, err := v.API.Request("stories.getViewers", params)
	if err != nil {
		return nil, err
	}

	var resp StoriesGetViewersResponse
	if params.Extended {
		var tmp StoriesGetViewersResponseExtended
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	} else {
		var tmp StoriesGetViewersResponseNormal
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// StoriesHideAllRepliesParams are params for Stories.HideAllReplies
type StoriesHideAllRepliesParams struct {
	// ID of the user whose replies should be hidden.
	OwnerID int `url:"owner_id"`
	GroupID int `url:"group_id,omitempty"`
}

// HideAllReplies Hides all replies in the last 24 hours from the user to current user's stories.
func (v Stories) HideAllReplies(params StoriesHideAllRepliesParams) (bool, error) {
	r, err := v.API.Request("stories.hideAllReplies", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// StoriesHideReplyParams are params for Stories.HideReply
type StoriesHideReplyParams struct {
	// ID of the user whose replies should be hidden.
	OwnerID int `url:"owner_id"`
	// Story ID.
	StoryID int `url:"story_id"`
	// Access key for the private object.
	AccessKey string `url:"access_key,omitempty"`
}

// HideReply Hides the reply to the current user's story.
func (v Stories) HideReply(params StoriesHideReplyParams) (bool, error) {
	r, err := v.API.Request("stories.hideReply", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// StoriesUnbanOwnerParams are params for Stories.UnbanOwner
type StoriesUnbanOwnerParams struct {
	// List of hidden sources to show stories from.
	OwnersIDs CSVIntSlice `url:"owners_ids"`
}

// UnbanOwner Allows to show stories from hidden sources in current user's feed.
func (v Stories) UnbanOwner(params StoriesUnbanOwnerParams) (bool, error) {
	r, err := v.API.Request("stories.unbanOwner", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}
