package vkCallbackApi

import (
	"github.com/mailru/easyjson"
)

// APILikes implements VK API namespace `likes`
type APILikes struct {
	API *API
}

// LikesGetListParams are params for APILikes.GetList
type LikesGetListParams struct {
	// , Object type: 'post' — post on user or community wall, 'comment' — comment on a wall post, 'photo' — photo, 'audio' — audio, 'video' — video, 'note' — note, 'photo_comment' — comment on the photo, 'video_comment' — comment on the video, 'topic_comment' — comment in the discussion, 'sitepage' — page of the site where the [vk.com/dev/Like|Like widget] is installed
	Type string `url:"type"`
	// ID of the user, community, or application that owns the object. If the 'type' parameter is set as 'sitepage', the application ID is passed as 'owner_id'. Use negative value for a community id. If the 'type' parameter is not set, the 'owner_id' is assumed to be either the current user or the same application ID as if the 'type' parameter was set to 'sitepage'.
	OwnerID int `url:"owner_id,omitempty"`
	// Object ID. If 'type' is set as 'sitepage', 'item_id' can include the 'page_id' parameter value used during initialization of the [vk.com/dev/Like|Like widget].
	ItemID int `url:"item_id,omitempty"`
	// URL of the page where the [vk.com/dev/Like|Like widget] is installed. Used instead of the 'item_id' parameter.
	PageURL string `url:"page_url,omitempty"`
	// Filters to apply: 'likes' — returns information about all users who liked the object (default), 'copies' — returns information only about users who told their friends about the object
	Filter string `url:"filter,omitempty"`
	// Specifies which users are returned: '1' — to return only the current user's friends, '0' — to return all users (default)
	FriendsOnly bool `url:"friends_only,omitempty"`
	// Specifies whether extended information will be returned. '1' — to return extended information about users and communities from the 'Likes' list, '0' — to return no additional information (default)
	Extended bool `url:"extended,omitempty"`
	// Offset needed to select a specific subset of users.
	Offset int `url:"offset,omitempty"`
	// Number of user IDs to return (maximum '1000'). Default is '100' if 'friends_only' is set to '0', otherwise, the default is '10' if 'friends_only' is set to '1'.
	Count   int  `url:"count,omitempty"`
	SkipOwn bool `url:"skip_own,omitempty"`
}

// LikesGetListResponse is response for APILikes.GetList
// Either LikesGetListResponseNormal or LikesGetListResponseExtended, depending on Extended flag
type LikesGetListResponse interface {
	isLikesGetList()
}

// LikesGetListResponseNormal is non-extended version of LikesGetListResponse
//easyjson:json
type LikesGetListResponseNormal struct {
	// Total number
	Count int `json:"count,omitempty"`
	// User ID
	Items []int `json:"items,omitempty"`
}

func (LikesGetListResponseNormal) isLikesGetList() {}

// LikesGetListResponseExtended is extended version of LikesGetListResponse
//easyjson:json
type LikesGetListResponseExtended struct {
	// Total number
	Count int    `json:"count,omitempty"`
	Items []User `json:"items,omitempty"`
}

func (LikesGetListResponseExtended) isLikesGetList() {}

// GetList Returns a list of IDs of users who added the specified object to their 'Likes' list.
func (v APILikes) GetList(params LikesGetListParams) (LikesGetListResponse, error) {
	r, err := v.API.Request("likes.getList", params)
	if err != nil {
		return nil, err
	}

	var resp LikesGetListResponse
	if params.Extended {
		var tmp LikesGetListResponseExtended
		err = easyjson.Unmarshal(r, &tmp)
		resp = &tmp
	} else {
		var tmp LikesGetListResponseNormal
		err = easyjson.Unmarshal(r, &tmp)
		resp = &tmp
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// LikesAddParams are params for APILikes.Add
type LikesAddParams struct {
	// Object type: 'post' — post on user or community wall, 'comment' — comment on a wall post, 'photo' — photo, 'audio' — audio, 'video' — video, 'note' — note, 'photo_comment' — comment on the photo, 'video_comment' — comment on the video, 'topic_comment' — comment in the discussion, 'sitepage' — page of the site where the [vk.com/dev/Like|Like widget] is installed
	Type string `url:"type"`
	// ID of the user or community that owns the object.
	OwnerID int `url:"owner_id,omitempty"`
	// Object ID.
	ItemID int `url:"item_id"`
	// Access key required for an object owned by a private entity.
	AccessKey string `url:"access_key,omitempty"`
}

// LikesAddResponse is response for APILikes.Add
//easyjson:json
type LikesAddResponse struct {
	// Total likes number
	Likes int `json:"likes,omitempty"`
}

// Add Adds the specified object to the 'Likes' list of the current user.
func (v APILikes) Add(params LikesAddParams) (*LikesAddResponse, error) {
	r, err := v.API.Request("likes.add", params)
	if err != nil {
		return nil, err
	}

	var resp LikesAddResponse
	err = easyjson.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// LikesDeleteParams are params for APILikes.Delete
type LikesDeleteParams struct {
	// Object type: 'post' — post on user or community wall, 'comment' — comment on a wall post, 'photo' — photo, 'audio' — audio, 'video' — video, 'note' — note, 'photo_comment' — comment on the photo, 'video_comment' — comment on the video, 'topic_comment' — comment in the discussion, 'sitepage' — page of the site where the [vk.com/dev/Like|Like widget] is installed
	Type string `url:"type"`
	// ID of the user or community that owns the object.
	OwnerID int `url:"owner_id,omitempty"`
	// Object ID.
	ItemID int `url:"item_id"`
}

// LikesDeleteResponse is response for APILikes.Delete
//easyjson:json
type LikesDeleteResponse struct {
	// Total likes number
	Likes int `json:"likes,omitempty"`
}

// Delete Deletes the specified object from the 'Likes' list of the current user.
func (v APILikes) Delete(params LikesDeleteParams) (*LikesDeleteResponse, error) {
	r, err := v.API.Request("likes.delete", params)
	if err != nil {
		return nil, err
	}

	var resp LikesDeleteResponse
	err = easyjson.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// LikesIsLikedParams are params for APILikes.IsLiked
type LikesIsLikedParams struct {
	// User ID.
	UserID int `url:"user_id,omitempty"`
	// Object type: 'post' — post on user or community wall, 'comment' — comment on a wall post, 'photo' — photo, 'audio' — audio, 'video' — video, 'note' — note, 'photo_comment' — comment on the photo, 'video_comment' — comment on the video, 'topic_comment' — comment in the discussion
	Type string `url:"type"`
	// ID of the user or community that owns the object.
	OwnerID int `url:"owner_id,omitempty"`
	// Object ID.
	ItemID int `url:"item_id"`
}

// LikesIsLikedResponse is response for APILikes.IsLiked
//easyjson:json
type LikesIsLikedResponse struct {
	// Information whether user liked the object
	Liked BoolInt `json:"liked,omitempty"`
	// Information whether user reposted the object
	Copied BoolInt `json:"copied,omitempty"`
}

// IsLiked Checks for the object in the 'Likes' list of the specified user.
func (v APILikes) IsLiked(params LikesIsLikedParams) (*LikesIsLikedResponse, error) {
	r, err := v.API.Request("likes.isLiked", params)
	if err != nil {
		return nil, err
	}

	var resp LikesIsLikedResponse
	err = easyjson.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
