package vkCallbackApi

import (
	"github.com/mailru/easyjson"
)

// APIWall implements VK API namespace `wall`
type APIWall struct {
	API *API
}

// WallGetParams are params for APIWall.Get
type WallGetParams struct {
	// ID of the user or community that owns the wall. By default, current user ID. Use a negative value to designate a community ID.
	OwnerID int `url:"owner_id,omitempty"`
	// User or community short address.
	Domain string `url:"domain,omitempty"`
	// Offset needed to return a specific subset of posts.
	Offset int `url:"offset,omitempty"`
	// Number of posts to return (maximum 100).
	Count int `url:"count,omitempty"`
	// Filter to apply: 'owner' — posts by the wall owner, 'others' — posts by someone else, 'all' — posts by the wall owner and others (default), 'postponed' — timed posts (only available for calls with an 'access_token'), 'suggests' — suggested posts on a community wall
	Filter string `url:"filter,omitempty"`
	// '1' — to return 'wall', 'profiles', and 'groups' fields, '0' — to return no additional fields (default)
	Extended bool           `url:"extended,omitempty"`
	Fields   CSVStringSlice `url:"fields,omitempty"`
}

// WallGetResponse is response for APIWall.Get
// Either WallGetResponseNormal or WallGetResponseExtended, depending on Extended flag
type WallGetResponse interface {
	isWallGet()
}

// WallGetResponseNormal is non-extended version of WallGetResponse
//easyjson:json
type WallGetResponseNormal struct {
	// Total number
	Count int    `json:"count,omitempty"`
	Items []Post `json:"items,omitempty"`
}

func (WallGetResponseNormal) isWallGet() {}

// WallGetResponseExtended is extended version of WallGetResponse
//easyjson:json
type WallGetResponseExtended struct {
	// Total number
	Count    int     `json:"count,omitempty"`
	Items    []Post  `json:"items,omitempty"`
	Profiles []User  `json:"profiles,omitempty"`
	Groups   []Group `json:"groups,omitempty"`
}

func (WallGetResponseExtended) isWallGet() {}

// Get Returns a list of posts on a user wall or community wall.
func (v APIWall) Get(params WallGetParams) (WallGetResponse, error) {
	r, err := v.API.Request("wall.get", params)
	if err != nil {
		return nil, err
	}

	var resp WallGetResponse
	if params.Extended {
		var tmp WallGetResponseExtended
		err = easyjson.Unmarshal(r, &tmp)
		resp = &tmp
	} else {
		var tmp WallGetResponseNormal
		err = easyjson.Unmarshal(r, &tmp)
		resp = &tmp
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// WallSearchParams are params for APIWall.Search
type WallSearchParams struct {
	// user or community id. "Remember that for a community 'owner_id' must be negative."
	OwnerID int `url:"owner_id,omitempty"`
	// user or community screen name.
	Domain string `url:"domain,omitempty"`
	// search query string.
	Query string `url:"query,omitempty"`
	// '1' – returns only page owner's posts.
	OwnersOnly bool `url:"owners_only,omitempty"`
	// count of posts to return.
	Count int `url:"count,omitempty"`
	// Offset needed to return a specific subset of posts.
	Offset int `url:"offset,omitempty"`
	// show extended post info.
	Extended bool           `url:"extended,omitempty"`
	Fields   CSVStringSlice `url:"fields,omitempty"`
}

// WallSearchResponse is response for APIWall.Search
// Either WallSearchResponseNormal or WallSearchResponseExtended, depending on Extended flag
type WallSearchResponse interface {
	isWallSearch()
}

// WallSearchResponseNormal is non-extended version of WallSearchResponse
//easyjson:json
type WallSearchResponseNormal struct {
	// Total number
	Count int    `json:"count,omitempty"`
	Items []Post `json:"items,omitempty"`
}

func (WallSearchResponseNormal) isWallSearch() {}

// WallSearchResponseExtended is extended version of WallSearchResponse
//easyjson:json
type WallSearchResponseExtended struct {
	// Total number
	Count    int     `json:"count,omitempty"`
	Items    []Post  `json:"items,omitempty"`
	Profiles []User  `json:"profiles,omitempty"`
	Groups   []Group `json:"groups,omitempty"`
}

func (WallSearchResponseExtended) isWallSearch() {}

// Search Allows to search posts on user or community walls.
func (v APIWall) Search(params WallSearchParams) (WallSearchResponse, error) {
	r, err := v.API.Request("wall.search", params)
	if err != nil {
		return nil, err
	}

	var resp WallSearchResponse
	if params.Extended {
		var tmp WallSearchResponseExtended
		err = easyjson.Unmarshal(r, &tmp)
		resp = &tmp
	} else {
		var tmp WallSearchResponseNormal
		err = easyjson.Unmarshal(r, &tmp)
		resp = &tmp
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// WallGetByIDParams are params for APIWall.GetByID
type WallGetByIDParams struct {
	// User or community IDs and post IDs, separated by underscores. Use a negative value to designate a community ID. Example: "93388_21539,93388_20904,2943_4276,-1_1"
	Posts CSVStringSlice `url:"posts"`
	// '1' — to return user and community objects needed to display posts, '0' — no additional fields are returned (default)
	Extended bool `url:"extended,omitempty"`
	// Sets the number of parent elements to include in the array 'copy_history' that is returned if the post is a repost from another wall.
	CopyHistoryDepth int            `url:"copy_history_depth,omitempty"`
	Fields           CSVStringSlice `url:"fields,omitempty"`
}

// WallGetByIDResponse is response for APIWall.GetByID
// Either WallGetByIDResponseNormal or WallGetByIDResponseExtended, depending on Extended flag
type WallGetByIDResponse interface {
	isWallGetByID()
}

// WallGetByIDResponseNormal is non-extended version of WallGetByIDResponse
//easyjson:json
type WallGetByIDResponseNormal []Post

func (WallGetByIDResponseNormal) isWallGetByID() {}

// WallGetByIDResponseExtended is extended version of WallGetByIDResponse
//easyjson:json
type WallGetByIDResponseExtended struct {
	Items    []Post  `json:"items,omitempty"`
	Profiles []User  `json:"profiles,omitempty"`
	Groups   []Group `json:"groups,omitempty"`
}

func (WallGetByIDResponseExtended) isWallGetByID() {}

// GetByID Returns a list of posts from user or community walls by their IDs.
func (v APIWall) GetByID(params WallGetByIDParams) (WallGetByIDResponse, error) {
	r, err := v.API.Request("wall.getById", params)
	if err != nil {
		return nil, err
	}

	var resp WallGetByIDResponse
	if params.Extended {
		var tmp WallGetByIDResponseExtended
		err = easyjson.Unmarshal(r, &tmp)
		resp = &tmp
	} else {
		var tmp WallGetByIDResponseNormal
		err = easyjson.Unmarshal(r, &tmp)
		resp = &tmp
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// WallPostParams are params for APIWall.Post
type WallPostParams struct {
	// User ID or community ID. Use a negative value to designate a community ID.
	OwnerID int `url:"owner_id,omitempty"`
	// '1' — post will be available to friends only, '0' — post will be available to all users (default)
	FriendsOnly bool `url:"friends_only,omitempty"`
	// For a community: '1' — post will be published by the community, '0' — post will be published by the user (default)
	FromGroup bool `url:"from_group,omitempty"`
	// (Required if 'attachments' is not set.) Text of the post.
	Message string `url:"message,omitempty"`
	// (Required if 'message' is not set.) List of objects attached to the post, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, 'page' — wiki-page, 'note' — note, 'poll' — poll, 'album' — photo album, '<owner_id>' — ID of the media application owner. '<media_id>' — Media application ID. Example: "photo100172_166443618,photo66748_265827614", May contain a link to an external page to include in the post. Example: "photo66748_265827614,http://habrahabr.ru", "NOTE: If more than one link is being attached, an error will be thrown."
	Attachments CSVStringSlice `url:"attachments,omitempty"`
	// List of services or websites the update will be exported to, if the user has so requested. Sample values: 'twitter', 'facebook'.
	Services string `url:"services,omitempty"`
	// Only for posts in communities with 'from_group' set to '1': '1' — post will be signed with the name of the posting user, '0' — post will not be signed (default)
	Signed bool `url:"signed,omitempty"`
	// Publication date (in Unix time). If used, posting will be delayed until the set time.
	PublishDate int `url:"publish_date,omitempty"`
	// Geographical latitude of a check-in, in degrees (from -90 to 90).
	Lat float32 `url:"lat,omitempty"`
	// Geographical longitude of a check-in, in degrees (from -180 to 180).
	Long float32 `url:"long,omitempty"`
	// ID of the location where the user was tagged.
	PlaceID int `url:"place_id,omitempty"`
	// Post ID. Used for publishing of scheduled and suggested posts.
	PostID    int    `url:"post_id,omitempty"`
	GUID      string `url:"guid,omitempty"`
	MarkAsAds bool   `url:"mark_as_ads,omitempty"`
}

// WallPostResponse is response for APIWall.Post
//easyjson:json
type WallPostResponse struct {
	// Created post ID
	PostID int `json:"post_id,omitempty"`
}

// Post Adds a new post on a user wall or community wall. Can also be used to publish suggested or scheduled posts.
func (v APIWall) Post(params WallPostParams) (*WallPostResponse, error) {
	r, err := v.API.Request("wall.post", params)
	if err != nil {
		return nil, err
	}

	var resp WallPostResponse
	err = easyjson.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// WallOpenCommentsParams are params for APIWall.OpenComments
type WallOpenCommentsParams struct {
	// User ID or community ID.
	OwnerID int `url:"owner_id"`
	// Post ID.
	PostID int `url:"post_id"`
}

// OpenComments Enables comments on post.
func (v APIWall) OpenComments(params WallOpenCommentsParams) (bool, error) {
	r, err := v.API.Request("wall.openComments", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// WallCloseCommentsParams are params for APIWall.CloseComments
type WallCloseCommentsParams struct {
	// User ID or community ID.
	OwnerID int `url:"owner_id"`
	// Post ID.
	PostID int `url:"post_id"`
}

// CloseComments Disables comments on post.
func (v APIWall) CloseComments(params WallCloseCommentsParams) (bool, error) {
	r, err := v.API.Request("wall.closeComments", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// WallPostAdsStealthParams are params for APIWall.PostAdsStealth
type WallPostAdsStealthParams struct {
	// User ID or community ID. Use a negative value to designate a community ID.
	OwnerID int `url:"owner_id"`
	// (Required if 'attachments' is not set.) Text of the post.
	Message string `url:"message,omitempty"`
	// (Required if 'message' is not set.) List of objects attached to the post, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, 'page' — wiki-page, 'note' — note, 'poll' — poll, 'album' — photo album, '<owner_id>' — ID of the media application owner. '<media_id>' — Media application ID. Example: "photo100172_166443618,photo66748_265827614", May contain a link to an external page to include in the post. Example: "photo66748_265827614,http://habrahabr.ru", "NOTE: If more than one link is being attached, an error will be thrown."
	Attachments CSVStringSlice `url:"attachments,omitempty"`
	// Only for posts in communities with 'from_group' set to '1': '1' — post will be signed with the name of the posting user, '0' — post will not be signed (default)
	Signed bool `url:"signed,omitempty"`
	// Geographical latitude of a check-in, in degrees (from -90 to 90).
	Lat float32 `url:"lat,omitempty"`
	// Geographical longitude of a check-in, in degrees (from -180 to 180).
	Long float32 `url:"long,omitempty"`
	// ID of the location where the user was tagged.
	PlaceID int `url:"place_id,omitempty"`
	// Post ID. Used for publishing of scheduled and suggested posts.
	PostID int `url:"post_id,omitempty"`
	// Unique identifier to avoid duplication the same post.
	GUID string `url:"guid,omitempty"`
	// Link button ID
	LinkButton string `url:"link_button,omitempty"`
	// Link title
	LinkTitle string `url:"link_title,omitempty"`
	// Link image url
	LinkImage string `url:"link_image,omitempty"`
}

// WallPostAdsStealthResponse is response for APIWall.PostAdsStealth
//easyjson:json
type WallPostAdsStealthResponse struct {
	// Created post ID
	PostID int `json:"post_id,omitempty"`
}

// PostAdsStealth Allows to create hidden post which will not be shown on the community's wall and can be used for creating an ad with type "Community post".
func (v APIWall) PostAdsStealth(params WallPostAdsStealthParams) (*WallPostAdsStealthResponse, error) {
	r, err := v.API.Request("wall.postAdsStealth", params)
	if err != nil {
		return nil, err
	}

	var resp WallPostAdsStealthResponse
	err = easyjson.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// WallRepostParams are params for APIWall.Repost
type WallRepostParams struct {
	// ID of the object to be reposted on the wall. Example: "wall66748_3675"
	Object string `url:"object"`
	// Comment to be added along with the reposted object.
	Message string `url:"message,omitempty"`
	// Target community ID when reposting to a community.
	GroupID   int  `url:"group_id,omitempty"`
	MarkAsAds bool `url:"mark_as_ads,omitempty"`
}

// WallRepostResponse is response for APIWall.Repost
//easyjson:json
type WallRepostResponse struct {
	Success BoolInt `json:"success,omitempty"`
	// Created post ID
	PostID int `json:"post_id,omitempty"`
	// Reposts number
	RepostsCount int `json:"reposts_count,omitempty"`
	// Reposts number
	LikesCount int `json:"likes_count,omitempty"`
}

// Repost Reposts (copies) an object to a user wall or community wall.
func (v APIWall) Repost(params WallRepostParams) (*WallRepostResponse, error) {
	r, err := v.API.Request("wall.repost", params)
	if err != nil {
		return nil, err
	}

	var resp WallRepostResponse
	err = easyjson.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// WallGetRepostsParams are params for APIWall.GetReposts
type WallGetRepostsParams struct {
	// User ID or community ID. By default, current user ID. Use a negative value to designate a community ID.
	OwnerID int `url:"owner_id,omitempty"`
	// Post ID.
	PostID int `url:"post_id,omitempty"`
	// Offset needed to return a specific subset of reposts.
	Offset int `url:"offset,omitempty"`
	// Number of reposts to return.
	Count int `url:"count,omitempty"`
}

// WallGetRepostsResponse is response for APIWall.GetReposts
//easyjson:json
type WallGetRepostsResponse struct {
	Items    []Post  `json:"items,omitempty"`
	Profiles []User  `json:"profiles,omitempty"`
	Groups   []Group `json:"groups,omitempty"`
}

// GetReposts Returns information about reposts of a post on user wall or community wall.
func (v APIWall) GetReposts(params WallGetRepostsParams) (*WallGetRepostsResponse, error) {
	r, err := v.API.Request("wall.getReposts", params)
	if err != nil {
		return nil, err
	}

	var resp WallGetRepostsResponse
	err = easyjson.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// WallEditParams are params for APIWall.Edit
type WallEditParams struct {
	// User ID or community ID. Use a negative value to designate a community ID.
	OwnerID int `url:"owner_id,omitempty"`
	// Post ID.
	PostID int `url:"post_id"`
	// (Applies only when editing a scheduled post.), '1' — post will be available to friends only, '0' — post will be available to all users (default)
	FriendsOnly bool `url:"friends_only,omitempty"`
	// (Required if 'attachments' is not set.) Text of the post.
	Message string `url:"message,omitempty"`
	// (Required if 'message' is not set.) List of objects attached to the post, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, '<owner_id>' — ID of the media application owner. '<media_id>' — Media application ID. Example: "photo100172_166443618,photo66748_265827614", May contain a link to an external page to include in the post. Example: "photo66748_265827614,http://habrahabr.ru", "NOTE: If more than one link is being attached, an error is thrown."
	Attachments CSVStringSlice `url:"attachments,omitempty"`
	// (Applies only to a scheduled post.) List of services or websites where status will be updated, if the user has so requested. Sample values: 'twitter', 'facebook'.
	Services string `url:"services,omitempty"`
	// (Applies only to a post that was created "as community" on a community wall.), '1' — to add the signature of the user who created the post
	Signed bool `url:"signed,omitempty"`
	// (Applies only to a scheduled post.) Publication date (in Unix time). If used, posting will be delayed until the set time.
	PublishDate int `url:"publish_date,omitempty"`
	// Geographical latitude of the check-in, in degrees (from -90 to 90).
	Lat float32 `url:"lat,omitempty"`
	// Geographical longitude of the check-in, in degrees (from -180 to 180).
	Long float32 `url:"long,omitempty"`
	// ID of the location where the user was tagged.
	PlaceID   int  `url:"place_id,omitempty"`
	MarkAsAds bool `url:"mark_as_ads,omitempty"`
}

// Edit Edits a post on a user wall or community wall.
func (v APIWall) Edit(params WallEditParams) (bool, error) {
	r, err := v.API.Request("wall.edit", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// WallEditAdsStealthParams are params for APIWall.EditAdsStealth
type WallEditAdsStealthParams struct {
	// User ID or community ID. Use a negative value to designate a community ID.
	OwnerID int `url:"owner_id,omitempty"`
	// Post ID
	PostID int `url:"post_id"`
	// (Required if 'attachments' is not set.) Text of the post.
	Message string `url:"message,omitempty"`
	// (Required if 'message' is not set.) List of objects attached to the post, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, 'page' — wiki-page, 'note' — note, 'poll' — poll, 'album' — photo album, '<owner_id>' — ID of the media application owner. '<media_id>' — Media application ID. Example: "photo100172_166443618,photo66748_265827614", May contain a link to an external page to include in the post. Example: "photo66748_265827614,http://habrahabr.ru", "NOTE: If more than one link is being attached, an error will be thrown."
	Attachments CSVStringSlice `url:"attachments,omitempty"`
	// Only for posts in communities with 'from_group' set to '1': '1' — post will be signed with the name of the posting user, '0' — post will not be signed (default)
	Signed bool `url:"signed,omitempty"`
	// Geographical latitude of a check-in, in degrees (from -90 to 90).
	Lat float32 `url:"lat,omitempty"`
	// Geographical longitude of a check-in, in degrees (from -180 to 180).
	Long float32 `url:"long,omitempty"`
	// ID of the location where the user was tagged.
	PlaceID int `url:"place_id,omitempty"`
	// Link button ID
	LinkButton string `url:"link_button,omitempty"`
	// Link title
	LinkTitle string `url:"link_title,omitempty"`
	// Link image url
	LinkImage string `url:"link_image,omitempty"`
}

// EditAdsStealth Allows to edit hidden post.
func (v APIWall) EditAdsStealth(params WallEditAdsStealthParams) (bool, error) {
	r, err := v.API.Request("wall.editAdsStealth", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// WallDeleteParams are params for APIWall.Delete
type WallDeleteParams struct {
	// User ID or community ID. Use a negative value to designate a community ID.
	OwnerID int `url:"owner_id,omitempty"`
	// ID of the post to be deleted.
	PostID int `url:"post_id,omitempty"`
}

// Delete Deletes a post from a user wall or community wall.
func (v APIWall) Delete(params WallDeleteParams) (bool, error) {
	r, err := v.API.Request("wall.delete", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// WallRestoreParams are params for APIWall.Restore
type WallRestoreParams struct {
	// User ID or community ID from whose wall the post was deleted. Use a negative value to designate a community ID.
	OwnerID int `url:"owner_id,omitempty"`
	// ID of the post to be restored.
	PostID int `url:"post_id,omitempty"`
}

// Restore Restores a post deleted from a user wall or community wall.
func (v APIWall) Restore(params WallRestoreParams) (bool, error) {
	r, err := v.API.Request("wall.restore", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// WallPinParams are params for APIWall.Pin
type WallPinParams struct {
	// ID of the user or community that owns the wall. By default, current user ID. Use a negative value to designate a community ID.
	OwnerID int `url:"owner_id,omitempty"`
	// Post ID.
	PostID int `url:"post_id"`
}

// Pin Pins the post on wall.
func (v APIWall) Pin(params WallPinParams) (bool, error) {
	r, err := v.API.Request("wall.pin", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// WallUnpinParams are params for APIWall.Unpin
type WallUnpinParams struct {
	// ID of the user or community that owns the wall. By default, current user ID. Use a negative value to designate a community ID.
	OwnerID int `url:"owner_id,omitempty"`
	// Post ID.
	PostID int `url:"post_id"`
}

// Unpin Unpins the post on wall.
func (v APIWall) Unpin(params WallUnpinParams) (bool, error) {
	r, err := v.API.Request("wall.unpin", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// WallGetCommentsParams are params for APIWall.GetComments
type WallGetCommentsParams struct {
	// User ID or community ID. Use a negative value to designate a community ID.
	OwnerID int `url:"owner_id,omitempty"`
	// Post ID.
	PostID int `url:"post_id"`
	// '1' — to return the 'likes' field, '0' — not to return the 'likes' field (default)
	NeedLikes      bool `url:"need_likes,omitempty"`
	StartCommentID int  `url:"start_comment_id,omitempty"`
	// Offset needed to return a specific subset of comments.
	Offset int `url:"offset,omitempty"`
	// Number of comments to return (maximum 100).
	Count int `url:"count,omitempty"`
	// Sort order: 'asc' — chronological, 'desc' — reverse chronological
	Sort string `url:"sort,omitempty"`
	// Number of characters at which to truncate comments when previewed. By default, '90'. Specify '0' if you do not want to truncate comments.
	PreviewLength int  `url:"preview_length,omitempty"`
	Extended      bool `url:"extended,omitempty"`
}

// WallGetCommentsResponse is response for APIWall.GetComments
// Either WallGetCommentsResponseNormal or WallGetCommentsResponseExtended, depending on Extended flag
type WallGetCommentsResponse interface {
	isWallGetComments()
}

// WallGetCommentsResponseNormal is non-extended version of WallGetCommentsResponse
//easyjson:json
type WallGetCommentsResponseNormal struct {
	// Total number
	Count int       `json:"count,omitempty"`
	Items []Comment `json:"items,omitempty"`
}

func (WallGetCommentsResponseNormal) isWallGetComments() {}

// WallGetCommentsResponseExtended is extended version of WallGetCommentsResponse
//easyjson:json
type WallGetCommentsResponseExtended struct {
	// Total number
	Count    int       `json:"count,omitempty"`
	Items    []Comment `json:"items,omitempty"`
	Profiles []User    `json:"profiles,omitempty"`
	Groups   []Group   `json:"groups,omitempty"`
}

func (WallGetCommentsResponseExtended) isWallGetComments() {}

// GetComments Returns a list of comments on a post on a user wall or community wall.
func (v APIWall) GetComments(params WallGetCommentsParams) (WallGetCommentsResponse, error) {
	r, err := v.API.Request("wall.getComments", params)
	if err != nil {
		return nil, err
	}

	var resp WallGetCommentsResponse
	if params.Extended {
		var tmp WallGetCommentsResponseExtended
		err = easyjson.Unmarshal(r, &tmp)
		resp = &tmp
	} else {
		var tmp WallGetCommentsResponseNormal
		err = easyjson.Unmarshal(r, &tmp)
		resp = &tmp
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// WallCreateCommentParams are params for APIWall.CreateComment
type WallCreateCommentParams struct {
	// User ID or community ID. Use a negative value to designate a community ID.
	OwnerID int `url:"owner_id,omitempty"`
	// Post ID.
	PostID int `url:"post_id"`
	// Group ID.
	FromGroup int `url:"from_group,omitempty"`
	// (Required if 'attachments' is not set.) Text of the comment.
	Message string `url:"message,omitempty"`
	// ID of comment to reply.
	ReplyToComment int `url:"reply_to_comment,omitempty"`
	// (Required if 'message' is not set.) List of media objects attached to the comment, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media ojbect: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, '<owner_id>' — ID of the media owner. '<media_id>' — Media ID. For example: "photo100172_166443618,photo66748_265827614"
	Attachments CSVStringSlice `url:"attachments,omitempty"`
	// Sticker ID.
	StickerID int `url:"sticker_id,omitempty"`
	// Unique identifier to avoid repeated comments.
	GUID string `url:"guid,omitempty"`
}

// WallCreateCommentResponse is response for APIWall.CreateComment
//easyjson:json
type WallCreateCommentResponse struct {
	// Created comment ID
	CommentID int `json:"comment_id,omitempty"`
}

// CreateComment Adds a comment to a post on a user wall or community wall.
func (v APIWall) CreateComment(params WallCreateCommentParams) (*WallCreateCommentResponse, error) {
	r, err := v.API.Request("wall.createComment", params)
	if err != nil {
		return nil, err
	}

	var resp WallCreateCommentResponse
	err = easyjson.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// WallEditCommentParams are params for APIWall.EditComment
type WallEditCommentParams struct {
	// User ID or community ID. Use a negative value to designate a community ID.
	OwnerID int `url:"owner_id,omitempty"`
	// Comment ID.
	CommentID int `url:"comment_id"`
	// New comment text.
	Message string `url:"message,omitempty"`
	// List of objects attached to the comment, in the following format: , "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, '<owner_id>' — ID of the media attachment owner. '<media_id>' — Media attachment ID. For example: "photo100172_166443618,photo66748_265827614"
	Attachments CSVStringSlice `url:"attachments,omitempty"`
}

// EditComment Edits a comment on a user wall or community wall.
func (v APIWall) EditComment(params WallEditCommentParams) (bool, error) {
	r, err := v.API.Request("wall.editComment", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// WallDeleteCommentParams are params for APIWall.DeleteComment
type WallDeleteCommentParams struct {
	// User ID or community ID. Use a negative value to designate a community ID.
	OwnerID int `url:"owner_id,omitempty"`
	// Comment ID.
	CommentID int `url:"comment_id"`
}

// DeleteComment Deletes a comment on a post on a user wall or community wall.
func (v APIWall) DeleteComment(params WallDeleteCommentParams) (bool, error) {
	r, err := v.API.Request("wall.deleteComment", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// WallRestoreCommentParams are params for APIWall.RestoreComment
type WallRestoreCommentParams struct {
	// User ID or community ID. Use a negative value to designate a community ID.
	OwnerID int `url:"owner_id,omitempty"`
	// Comment ID.
	CommentID int `url:"comment_id"`
}

// RestoreComment Restores a comment deleted from a user wall or community wall.
func (v APIWall) RestoreComment(params WallRestoreCommentParams) (bool, error) {
	r, err := v.API.Request("wall.restoreComment", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// WallReportPostParams are params for APIWall.ReportPost
type WallReportPostParams struct {
	// ID of the user or community that owns the wall.
	OwnerID int `url:"owner_id"`
	// Post ID.
	PostID int `url:"post_id"`
	// Reason for the complaint: '0' – spam, '1' – child pornography, '2' – extremism, '3' – violence, '4' – drug propaganda, '5' – adult material, '6' – insult, abuse
	Reason int `url:"reason,omitempty"`
}

// ReportPost Reports (submits a complaint about) a post on a user wall or community wall.
func (v APIWall) ReportPost(params WallReportPostParams) (bool, error) {
	r, err := v.API.Request("wall.reportPost", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// WallReportCommentParams are params for APIWall.ReportComment
type WallReportCommentParams struct {
	// ID of the user or community that owns the wall.
	OwnerID int `url:"owner_id"`
	// Comment ID.
	CommentID int `url:"comment_id"`
	// Reason for the complaint: '0' – spam, '1' – child pornography, '2' – extremism, '3' – violence, '4' – drug propaganda, '5' – adult material, '6' – insult, abuse
	Reason int `url:"reason,omitempty"`
}

// ReportComment Reports (submits a complaint about) a comment on a post on a user wall or community wall.
func (v APIWall) ReportComment(params WallReportCommentParams) (bool, error) {
	r, err := v.API.Request("wall.reportComment", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}
