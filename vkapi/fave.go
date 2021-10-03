package vkapi

import (
	"encoding/json"

	"github.com/stek29/vk"
)

// Fave implements VK API namespace `fave`
type Fave struct {
	API vk.API
}

// FaveGetUsersParams are params for Fave.GetUsers
type FaveGetUsersParams struct {
	// Offset needed to return a specific subset of users.
	Offset int `url:"offset,omitempty"`
	// Number of users to return.
	Count int `url:"count,omitempty"`
}

// FaveGetUsersResponse is response for Fave.GetUsers
//easyjson:json
type FaveGetUsersResponse struct {
	// Total number
	Count int       `json:"count,omitempty"`
	Items []vk.User `json:"items,omitempty"`
}

// GetUsers Returns a list of users whom the current user has bookmarked.
func (v Fave) GetUsers(params FaveGetUsersParams) (*FaveGetUsersResponse, error) {
	r, err := v.API.Request("fave.getUsers", params)
	if err != nil {
		return nil, err
	}

	var resp FaveGetUsersResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// FaveGetPhotosParams are params for Fave.GetPhotos
type FaveGetPhotosParams struct {
	// Offset needed to return a specific subset of photos.
	Offset int `url:"offset,omitempty"`
	// Number of photos to return.
	Count int `url:"count,omitempty"`
}

// FaveGetPhotosResponse is response for Fave.GetPhotos
//easyjson:json
type FaveGetPhotosResponse struct {
	// Total number
	Count int        `json:"count,omitempty"`
	Items []vk.Photo `json:"items,omitempty"`
}

// GetPhotos Returns a list of photos that the current user has liked.
func (v Fave) GetPhotos(params FaveGetPhotosParams) (*FaveGetPhotosResponse, error) {
	r, err := v.API.Request("fave.getPhotos", params)
	if err != nil {
		return nil, err
	}

	var resp FaveGetPhotosResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// FaveGetPostsParams are params for Fave.GetPosts
type FaveGetPostsParams struct {
	// Offset needed to return a specific subset of posts.
	Offset int `url:"offset,omitempty"`
	// Number of posts to return.
	Count int `url:"count,omitempty"`
	// '1' — to return additional 'wall', 'profiles', and 'groups' fields. By default: '0'.
	Extended bool `url:"extended,omitempty"`
}

// FaveGetPostsResponse is response for Fave.GetPosts
//easyjson:json
type FaveGetPostsResponse struct {
	// Total number
	Count int       `json:"count,omitempty"`
	Items []vk.Post `json:"items,omitempty"`
}

// GetPosts Returns a list of wall posts that the current user has liked.
func (v Fave) GetPosts(params FaveGetPostsParams) (*FaveGetPostsResponse, error) {
	r, err := v.API.Request("fave.getPosts", params)
	if err != nil {
		return nil, err
	}

	var resp FaveGetPostsResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// FaveGetVideosParams are params for Fave.GetVideos
type FaveGetVideosParams struct {
	// Offset needed to return a specific subset of videos.
	Offset int `url:"offset,omitempty"`
	// Number of videos to return.
	Count int `url:"count,omitempty"`
	// Return an additional information about videos. Also returns all owners profiles and groups.
	Extended bool `url:"extended,omitempty"`
}

// FaveGetVideosResponse is response for Fave.GetVideos
//easyjson:json
type FaveGetVideosResponse struct {
	// Total number
	Count int        `json:"count,omitempty"`
	Items []vk.Video `json:"items,omitempty"`
}

// GetVideos Returns a list of videos that the current user has liked.
func (v Fave) GetVideos(params FaveGetVideosParams) (*FaveGetVideosResponse, error) {
	r, err := v.API.Request("fave.getVideos", params)
	if err != nil {
		return nil, err
	}

	var resp FaveGetVideosResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// FaveGetLinksParams are params for Fave.GetLinks
type FaveGetLinksParams struct {
	// Offset needed to return a specific subset of users.
	Offset int `url:"offset,omitempty"`
	// Number of results to return.
	Count int `url:"count,omitempty"`
}

// FaveGetLinksResponse is response for Fave.GetLinks
//easyjson:json
type FaveGetLinksResponse struct {
	// Total number
	Count int           `json:"count,omitempty"`
	Items []vk.MiniLink `json:"items,omitempty"`
}

// GetLinks Returns a list of links that the current user has bookmarked.
func (v Fave) GetLinks(params FaveGetLinksParams) (*FaveGetLinksResponse, error) {
	r, err := v.API.Request("fave.getLinks", params)
	if err != nil {
		return nil, err
	}

	var resp FaveGetLinksResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// FaveGetMarketItemsParams are params for Fave.GetMarketItems
type FaveGetMarketItemsParams struct {
	// Number of results to return.
	Count  int `url:"count,omitempty"`
	Offset int `url:"offset,omitempty"`
	// '1' – to return additional fields 'likes, can_comment, can_repost, photos'. By default: '0'.
	Extended bool `url:"extended,omitempty"`
}

// FaveGetMarketItemsResponse is response for Fave.GetMarketItems
//easyjson:json
type FaveGetMarketItemsResponse struct {
	// Total number
	Count int             `json:"count,omitempty"`
	Items []vk.MarketItem `json:"items,omitempty"`
}

// GetMarketItems Returns market items bookmarked by current user.
func (v Fave) GetMarketItems(params FaveGetMarketItemsParams) (*FaveGetMarketItemsResponse, error) {
	r, err := v.API.Request("fave.getMarketItems", params)
	if err != nil {
		return nil, err
	}

	var resp FaveGetMarketItemsResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// FaveAddUserParams are params for Fave.AddUser
type FaveAddUserParams struct {
	// Profile ID.
	UserID int `url:"user_id"`
}

// AddUser Adds a profile to user faves.
func (v Fave) AddUser(params FaveAddUserParams) (bool, error) {
	r, err := v.API.Request("fave.addUser", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// FaveRemoveUserParams are params for Fave.RemoveUser
type FaveRemoveUserParams struct {
	// Profile ID.
	UserID int `url:"user_id"`
}

// RemoveUser Removes a profile from user faves.
func (v Fave) RemoveUser(params FaveRemoveUserParams) (bool, error) {
	r, err := v.API.Request("fave.removeUser", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// FaveAddGroupParams are params for Fave.AddGroup
type FaveAddGroupParams struct {
	// Community ID.
	GroupID int `url:"group_id"`
}

// AddGroup Adds a community to user faves.
func (v Fave) AddGroup(params FaveAddGroupParams) (bool, error) {
	r, err := v.API.Request("fave.addGroup", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// FaveRemoveGroupParams are params for Fave.RemoveGroup
type FaveRemoveGroupParams struct {
	// Community ID.
	GroupID int `url:"group_id"`
}

// RemoveGroup Removes a community from user faves.
func (v Fave) RemoveGroup(params FaveRemoveGroupParams) (bool, error) {
	r, err := v.API.Request("fave.removeGroup", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// FaveAddLinkParams are params for Fave.AddLink
type FaveAddLinkParams struct {
	// Link URL.
	Link string `url:"link"`
}

// AddLink Adds a link to user faves.
func (v Fave) AddLink(params FaveAddLinkParams) (bool, error) {
	r, err := v.API.Request("fave.addLink", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// FaveRemoveLinkParams are params for Fave.RemoveLink
type FaveRemoveLinkParams struct {
	// Link ID (can be obtained by [vk.com/dev/faves.getLinks|faves.getLinks] method).
	LinkID string `url:"link_id,omitempty"`
}

// RemoveLink Removes link from the user's faves.
func (v Fave) RemoveLink(params FaveRemoveLinkParams) (bool, error) {
	r, err := v.API.Request("fave.removeLink", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}
