package vkCallbackApi

import (
	"encoding/json"
	"strconv"
)

// APINewsfeed implements VK API namespace `newsfeed`
type APINewsfeed struct {
	API *API
}

// NewsfeedGetParams are params for APINewsfeed.Get
type NewsfeedGetParams struct {
	// Filters to apply: 'post' — new wall posts, 'photo' — new photos, 'photo_tag' — new photo tags, 'wall_photo' — new wall photos, 'friend' — new friends, 'note' — new notes
	Filters CSVStringSlice `url:"filters,omitempty"`
	// '1' — to return news items from banned sources
	ReturnBanned bool `url:"return_banned,omitempty"`
	// Earliest timestamp (in Unix time) of a news item to return. By default, 24 hours ago.
	StartTime int `url:"start_time,omitempty"`
	// Latest timestamp (in Unix time) of a news item to return. By default, the current time.
	EndTime int `url:"end_time,omitempty"`
	// Maximum number of photos to return. By default, '5'.
	MaxPhotos int `url:"max_photos,omitempty"`
	// Sources to obtain news from, separated by commas. User IDs can be specified in formats '' or 'u' , where '' is the user's friend ID. Community IDs can be specified in formats '-' or 'g' , where '' is the community ID. If the parameter is not set, all of the user's friends and communities are returned, except for banned sources, which can be obtained with the [vk.com/dev/newsfeed.getBanned|newsfeed.getBanned] method.
	SourceIDs CSVStringSlice `url:"source_ids,omitempty"`
	// identifier required to get the next page of results. Value for this parameter is returned in 'next_from' field in a reply.
	StartFrom string `url:"start_from,omitempty"`
	// Number of news items to return (default 50, maximum 100). For auto feed, you can use the 'new_offset' parameter returned by this method.
	Count int `url:"count,omitempty"`
	// Additional fields of [vk.com/dev/fields|profiles] and [vk.com/dev/fields_groups|communities] to return.
	Fields CSVStringSlice `url:"fields,omitempty"`
}

// NewsfeedGetResponse is response for APINewsfeed.Get
//easyjson:json
type NewsfeedGetResponse struct {
	Items    []genTODOType/* objects.json#/definitions/newsfeed_newsfeed_item */ `json:"items,omitempty"`
	Profiles []User  `json:"profiles,omitempty"`
	Groups   []Group `json:"groups,omitempty"`
}

// Get Returns data required to show newsfeed for the current user.
func (v APINewsfeed) Get(params NewsfeedGetParams) (*NewsfeedGetResponse, error) {
	r, err := v.API.Request("newsfeed.get", params)
	if err != nil {
		return nil, err
	}

	var resp NewsfeedGetResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// NewsfeedGetRecommendedParams are params for APINewsfeed.GetRecommended
type NewsfeedGetRecommendedParams struct {
	// Earliest timestamp (in Unix time) of a news item to return. By default, 24 hours ago.
	StartTime int `url:"start_time,omitempty"`
	// Latest timestamp (in Unix time) of a news item to return. By default, the current time.
	EndTime int `url:"end_time,omitempty"`
	// Maximum number of photos to return. By default, '5'.
	MaxPhotos int `url:"max_photos,omitempty"`
	// 'new_from' value obtained in previous call.
	StartFrom string `url:"start_from,omitempty"`
	// Number of news items to return.
	Count int `url:"count,omitempty"`
	// Additional fields of [vk.com/dev/fields|profiles] and [vk.com/dev/fields_groups|communities] to return.
	Fields CSVStringSlice `url:"fields,omitempty"`
}

// NewsfeedGetRecommendedResponse is response for APINewsfeed.GetRecommended
//easyjson:json
type NewsfeedGetRecommendedResponse struct {
	Items    []genTODOType/* objects.json#/definitions/newsfeed_newsfeed_item */ `json:"items,omitempty"`
	Profiles []User  `json:"profiles,omitempty"`
	Groups   []Group `json:"groups,omitempty"`
	// New offset value
	NewOffset string `json:"new_offset,omitempty"`
	// New from value
	NewFrom string `json:"new_from,omitempty"`
}

// GetRecommended , Returns a list of newsfeeds recommended to the current user.
func (v APINewsfeed) GetRecommended(params NewsfeedGetRecommendedParams) (*NewsfeedGetRecommendedResponse, error) {
	r, err := v.API.Request("newsfeed.getRecommended", params)
	if err != nil {
		return nil, err
	}

	var resp NewsfeedGetRecommendedResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// NewsfeedGetCommentsParams are params for APINewsfeed.GetComments
type NewsfeedGetCommentsParams struct {
	// Number of comments to return. For auto feed, you can use the 'new_offset' parameter returned by this method.
	Count int `url:"count,omitempty"`
	// Filters to apply: 'post' — new comments on wall posts, 'photo' — new comments on photos, 'video' — new comments on videos, 'topic' — new comments on discussions, 'note' — new comments on notes,
	Filters CSVStringSlice `url:"filters,omitempty"`
	// Object ID, comments on repost of which shall be returned, e.g. 'wall1_45486'. (If the parameter is set, the 'filters' parameter is optional.),
	Reposts string `url:"reposts,omitempty"`
	// Earliest timestamp (in Unix time) of a comment to return. By default, 24 hours ago.
	StartTime int `url:"start_time,omitempty"`
	// Latest timestamp (in Unix time) of a comment to return. By default, the current time.
	EndTime int `url:"end_time,omitempty"`
	// Identificator needed to return the next page with results. Value for this parameter returns in 'next_from' field.
	StartFrom string `url:"start_from,omitempty"`
	// Additional fields of [vk.com/dev/fields|profiles] and [vk.com/dev/fields_groups|communities] to return.
	Fields CSVStringSlice `url:"fields,omitempty"`
}

// NewsfeedGetCommentsResponse is response for APINewsfeed.GetComments
//easyjson:json
type NewsfeedGetCommentsResponse struct {
	Items    []genTODOType/* objects.json#/definitions/newsfeed_newsfeed_item */ `json:"items,omitempty"`
	Profiles []User  `json:"profiles,omitempty"`
	Groups   []Group `json:"groups,omitempty"`
	// New from value
	NextFrom string `json:"next_from,omitempty"`
}

// GetComments Returns a list of comments in the current user's newsfeed.
func (v APINewsfeed) GetComments(params NewsfeedGetCommentsParams) (*NewsfeedGetCommentsResponse, error) {
	r, err := v.API.Request("newsfeed.getComments", params)
	if err != nil {
		return nil, err
	}

	var resp NewsfeedGetCommentsResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// NewsfeedGetMentionsParams are params for APINewsfeed.GetMentions
type NewsfeedGetMentionsParams struct {
	// Owner ID.
	OwnerID int `url:"owner_id,omitempty"`
	// Earliest timestamp (in Unix time) of a post to return. By default, 24 hours ago.
	StartTime int `url:"start_time,omitempty"`
	// Latest timestamp (in Unix time) of a post to return. By default, the current time.
	EndTime int `url:"end_time,omitempty"`
	// Offset needed to return a specific subset of posts.
	Offset int `url:"offset,omitempty"`
	// Number of posts to return.
	Count int `url:"count,omitempty"`
}

// NewsfeedGetMentionsResponse is response for APINewsfeed.GetMentions
//easyjson:json
type NewsfeedGetMentionsResponse struct {
	// Total number
	Count int    `json:"count,omitempty"`
	Items []Post `json:"items,omitempty"`
}

// GetMentions Returns a list of posts on user walls in which the current user is mentioned.
func (v APINewsfeed) GetMentions(params NewsfeedGetMentionsParams) (*NewsfeedGetMentionsResponse, error) {
	r, err := v.API.Request("newsfeed.getMentions", params)
	if err != nil {
		return nil, err
	}

	var resp NewsfeedGetMentionsResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// NewsfeedGetBannedParams are params for APINewsfeed.GetBanned
type NewsfeedGetBannedParams struct {
	// '1' — return extra information about users and communities
	Extended bool `url:"extended,omitempty"`
	// Profile fields to return.
	Fields CSVStringSlice `url:"fields,omitempty"`
	// Case for declension of user name and surname: 'nom' — nominative (default), 'gen' — genitive , 'dat' — dative, 'acc' — accusative , 'ins' — instrumental , 'abl' — prepositional
	NameCase string `url:"name_case,omitempty"`
}

// NewsfeedGetBannedResponse is response for APINewsfeed.GetBanned
// Either NewsfeedGetBannedResponseNormal or NewsfeedGetBannedResponseExtended, depending on Extended flag
type NewsfeedGetBannedResponse interface {
	isNewsfeedGetBanned()
}

// NewsfeedGetBannedResponseNormal is non-extended version of NewsfeedGetBannedResponse
//easyjson:json
type NewsfeedGetBannedResponseNormal struct {
	// Community ID
	Groups []int `json:"groups,omitempty"`
	// User ID
	Members []int `json:"members,omitempty"`
}

func (NewsfeedGetBannedResponseNormal) isNewsfeedGetBanned() {}

// NewsfeedGetBannedResponseExtended is extended version of NewsfeedGetBannedResponse
//easyjson:json
type NewsfeedGetBannedResponseExtended struct {
	Groups  []User  `json:"groups,omitempty"`
	Members []Group `json:"members,omitempty"`
}

func (NewsfeedGetBannedResponseExtended) isNewsfeedGetBanned() {}

// GetBanned Returns a list of users and communities banned from the current user's newsfeed.
func (v APINewsfeed) GetBanned(params NewsfeedGetBannedParams) (NewsfeedGetBannedResponse, error) {
	r, err := v.API.Request("newsfeed.getBanned", params)
	if err != nil {
		return nil, err
	}

	var resp NewsfeedGetBannedResponse
	if params.Extended {
		var tmp NewsfeedGetBannedResponseExtended
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	} else {
		var tmp NewsfeedGetBannedResponseNormal
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// NewsfeedAddBanParams are params for APINewsfeed.AddBan
type NewsfeedAddBanParams struct {
	UserIDs  CSVIntSlice `url:"user_ids,omitempty"`
	GroupIDs CSVIntSlice `url:"group_ids,omitempty"`
}

// AddBan Prevents news from specified users and communities from appearing in the current user's newsfeed.
func (v APINewsfeed) AddBan(params NewsfeedAddBanParams) (bool, error) {
	r, err := v.API.Request("newsfeed.addBan", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// NewsfeedDeleteBanParams are params for APINewsfeed.DeleteBan
type NewsfeedDeleteBanParams struct {
	UserIDs  CSVIntSlice `url:"user_ids,omitempty"`
	GroupIDs CSVIntSlice `url:"group_ids,omitempty"`
}

// DeleteBan Allows news from previously banned users and communities to be shown in the current user's newsfeed.
func (v APINewsfeed) DeleteBan(params NewsfeedDeleteBanParams) (bool, error) {
	r, err := v.API.Request("newsfeed.deleteBan", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// NewsfeedIgnoreItemParams are params for APINewsfeed.IgnoreItem
type NewsfeedIgnoreItemParams struct {
	// Item type. Possible values: *'wall' – post on the wall,, *'tag' – tag on a photo,, *'profilephoto' – profile photo,, *'video' – video,, *'audio' – audio.
	Type string `url:"type"`
	// Item owner's identifier (user or community), "Note that community id must be negative. 'owner_id=1' – user , 'owner_id=-1' – community "
	OwnerID int `url:"owner_id"`
	// Item identifier
	ItemID int `url:"item_id"`
}

// IgnoreItem Hides an item from the newsfeed.
func (v APINewsfeed) IgnoreItem(params NewsfeedIgnoreItemParams) (bool, error) {
	r, err := v.API.Request("newsfeed.ignoreItem", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// NewsfeedUnignoreItemParams are params for APINewsfeed.UnignoreItem
type NewsfeedUnignoreItemParams struct {
	// Item type. Possible values: *'wall' – post on the wall,, *'tag' – tag on a photo,, *'profilephoto' – profile photo,, *'video' – video,, *'audio' – audio.
	Type string `url:"type"`
	// Item owner's identifier (user or community), "Note that community id must be negative. 'owner_id=1' – user , 'owner_id=-1' – community "
	OwnerID int `url:"owner_id"`
	// Item identifier
	ItemID int `url:"item_id"`
}

// UnignoreItem Returns a hidden item to the newsfeed.
func (v APINewsfeed) UnignoreItem(params NewsfeedUnignoreItemParams) (bool, error) {
	r, err := v.API.Request("newsfeed.unignoreItem", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// NewsfeedSearchParams are params for APINewsfeed.Search
type NewsfeedSearchParams struct {
	// Search query string (e.g., 'New Year').
	Q string `url:"q,omitempty"`
	// '1' — to return additional information about the user or community that placed the post.
	Extended bool `url:"extended,omitempty"`
	// Number of posts to return.
	Count int `url:"count,omitempty"`
	// Geographical latitude point (in degrees, -90 to 90) within which to search.
	Latitude float32 `url:"latitude,omitempty"`
	// Geographical longitude point (in degrees, -180 to 180) within which to search.
	Longitude float32 `url:"longitude,omitempty"`
	// Earliest timestamp (in Unix time) of a news item to return. By default, 24 hours ago.
	StartTime int `url:"start_time,omitempty"`
	// Latest timestamp (in Unix time) of a news item to return. By default, the current time.
	EndTime   int    `url:"end_time,omitempty"`
	StartFrom string `url:"start_from,omitempty"`
	// Additional fields of [vk.com/dev/fields|profiles] and [vk.com/dev/fields_groups|communities] to return.
	Fields CSVStringSlice `url:"fields,omitempty"`
}

// NewsfeedSearchResponse is response for APINewsfeed.Search
// Either NewsfeedSearchResponseNormal or NewsfeedSearchResponseExtended, depending on Extended flag
type NewsfeedSearchResponse interface {
	isNewsfeedSearch()
}

// NewsfeedSearchResponseNormal is non-extended version of NewsfeedSearchResponse
//easyjson:json
type NewsfeedSearchResponseNormal struct {
	Items            []Post   `json:"items,omitempty"`
	SuggestedQueries []string `json:"suggested_queries,omitempty"`
}

func (NewsfeedSearchResponseNormal) isNewsfeedSearch() {}

// NewsfeedSearchResponseExtended is extended version of NewsfeedSearchResponse
//easyjson:json
type NewsfeedSearchResponseExtended struct {
	Items    []Post  `json:"items,omitempty"`
	Profiles []User  `json:"profiles,omitempty"`
	Groups   []Group `json:"groups,omitempty"`
}

func (NewsfeedSearchResponseExtended) isNewsfeedSearch() {}

// Search Returns search results by statuses.
func (v APINewsfeed) Search(params NewsfeedSearchParams) (NewsfeedSearchResponse, error) {
	r, err := v.API.Request("newsfeed.search", params)
	if err != nil {
		return nil, err
	}

	var resp NewsfeedSearchResponse
	if params.Extended {
		var tmp NewsfeedSearchResponseExtended
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	} else {
		var tmp NewsfeedSearchResponseNormal
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// NewsfeedGetListsParams are params for APINewsfeed.GetLists
type NewsfeedGetListsParams struct {
	// numeric list identifiers.
	ListIDs CSVIntSlice `url:"list_ids,omitempty"`
	// Return additional list info
	Extended bool `url:"extended,omitempty"`
}

// NewsfeedGetListsResponse is response for APINewsfeed.GetLists
// Either NewsfeedGetListsResponseNormal or NewsfeedGetListsResponseExtended, depending on Extended flag
type NewsfeedGetListsResponse interface {
	isNewsfeedGetLists()
}

// NewsfeedGetListsResponseNormal is non-extended version of NewsfeedGetListsResponse
//easyjson:json
type NewsfeedGetListsResponseNormal struct {
	// Total number
	Count int `json:"count,omitempty"`
	Items []genTODOType/* objects.json#/definitions/newsfeed_list */ `json:"items,omitempty"`
}

func (NewsfeedGetListsResponseNormal) isNewsfeedGetLists() {}

// NewsfeedGetListsResponseExtended is extended version of NewsfeedGetListsResponse
//easyjson:json
type NewsfeedGetListsResponseExtended struct {
	// Total number
	Count int `json:"count,omitempty"`
	Items []genTODOType/* objects.json#/definitions/newsfeed_list_full */ `json:"items,omitempty"`
}

func (NewsfeedGetListsResponseExtended) isNewsfeedGetLists() {}

// GetLists Returns a list of newsfeeds followed by the current user.
func (v APINewsfeed) GetLists(params NewsfeedGetListsParams) (NewsfeedGetListsResponse, error) {
	r, err := v.API.Request("newsfeed.getLists", params)
	if err != nil {
		return nil, err
	}

	var resp NewsfeedGetListsResponse
	if params.Extended {
		var tmp NewsfeedGetListsResponseExtended
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	} else {
		var tmp NewsfeedGetListsResponseNormal
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// NewsfeedSaveListParams are params for APINewsfeed.SaveList
type NewsfeedSaveListParams struct {
	// numeric list identifier (if not sent, will be set automatically).
	ListID int `url:"list_id,omitempty"`
	// list name.
	Title string `url:"title"`
	// users and communities identifiers to be added to the list. Community identifiers must be negative numbers.
	SourceIDs CSVIntSlice `url:"source_ids,omitempty"`
	// reposts display on and off ('1' is for off).
	NoReposts bool `url:"no_reposts,omitempty"`
}

// NewsfeedSaveListResponse is response for APINewsfeed.SaveList
// List ID
type NewsfeedSaveListResponse int

// SaveList Creates and edits user newsfeed lists
func (v APINewsfeed) SaveList(params NewsfeedSaveListParams) (NewsfeedSaveListResponse, error) {
	r, err := v.API.Request("newsfeed.saveList", params)
	if err != nil {
		return 0, err
	}

	var resp NewsfeedSaveListResponse

	var cnv int
	cnv, err = strconv.Atoi(string(r))
	resp = NewsfeedSaveListResponse(cnv)

	if err != nil {
		return 0, err
	}
	return resp, nil
}

// NewsfeedDeleteListParams are params for APINewsfeed.DeleteList
type NewsfeedDeleteListParams struct {
	ListID int `url:"list_id"`
}

// DeleteList does newsfeed.deleteList
func (v APINewsfeed) DeleteList(params NewsfeedDeleteListParams) (bool, error) {
	r, err := v.API.Request("newsfeed.deleteList", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// NewsfeedUnsubscribeParams are params for APINewsfeed.Unsubscribe
type NewsfeedUnsubscribeParams struct {
	// Type of object from which to unsubscribe: 'note' — note, 'photo' — photo, 'post' — post on user wall or community wall, 'topic' — topic, 'video' — video
	Type string `url:"type"`
	// Object owner ID.
	OwnerID int `url:"owner_id,omitempty"`
	// Object ID.
	ItemID int `url:"item_id"`
}

// Unsubscribe Unsubscribes the current user from specified newsfeeds.
func (v APINewsfeed) Unsubscribe(params NewsfeedUnsubscribeParams) (bool, error) {
	r, err := v.API.Request("newsfeed.unsubscribe", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// NewsfeedGetSuggestedSourcesParams are params for APINewsfeed.GetSuggestedSources
type NewsfeedGetSuggestedSourcesParams struct {
	// offset required to choose a particular subset of communities or users.
	Offset int `url:"offset,omitempty"`
	// amount of communities or users to return.
	Count int `url:"count,omitempty"`
	// shuffle the returned list or not.
	Shuffle bool `url:"shuffle,omitempty"`
	// list of extra fields to be returned. See available fields for [vk.com/dev/fields|users] and [vk.com/dev/fields_groups|communities].
	Fields CSVStringSlice `url:"fields,omitempty"`
}

// NewsfeedGetSuggestedSourcesResponse is response for APINewsfeed.GetSuggestedSources
//easyjson:json
type NewsfeedGetSuggestedSourcesResponse struct {
	// Total number
	Count int         `json:"count,omitempty"`
	Items genTODOType `json:"items,omitempty"`
}

// GetSuggestedSources Returns communities and users that current user is suggested to follow.
func (v APINewsfeed) GetSuggestedSources(params NewsfeedGetSuggestedSourcesParams) (*NewsfeedGetSuggestedSourcesResponse, error) {
	r, err := v.API.Request("newsfeed.getSuggestedSources", params)
	if err != nil {
		return nil, err
	}

	var resp NewsfeedGetSuggestedSourcesResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
