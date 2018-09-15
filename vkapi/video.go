package vkapi

import (
	"encoding/json"
	"strconv"

	"github.com/stek29/vk"
)

// Video implements VK API namespace `video`
type Video struct {
	API vk.API
}

// VideoGetParams are params for Video.Get
type VideoGetParams struct {
	// ID of the user or community that owns the video(s).
	OwnerID int `url:"owner_id,omitempty"`
	// Video IDs, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", Use a negative value to designate a community ID. Example: "-4363_136089719,13245770_137352259"
	Videos CSVStringSlice `url:"videos,omitempty"`
	// ID of the album containing the video(s).
	AlbumID int `url:"album_id,omitempty"`
	// Number of videos to return.
	Count int `url:"count,omitempty"`
	// Offset needed to return a specific subset of videos.
	Offset int `url:"offset,omitempty"`
	// '1' — to return an extended response with additional fields
	Extended bool `url:"extended,omitempty"`
}

// VideoGetResponse is response for Video.Get
// Either VideoGetResponseNormal or VideoGetResponseExtended, depending on Extended flag
type VideoGetResponse interface {
	isVideoGet()
}

// VideoGetResponseNormal is non-extended version of VideoGetResponse
//easyjson:json
type VideoGetResponseNormal struct {
	// Total number
	Count int        `json:"count,omitempty"`
	Items []vk.Video `json:"items,omitempty"`
}

func (VideoGetResponseNormal) isVideoGet() {}

// VideoGetResponseExtended is extended version of VideoGetResponse
//easyjson:json
type VideoGetResponseExtended struct {
	// Total number
	Count    int        `json:"count,omitempty"`
	Items    []vk.Video `json:"items,omitempty"`
	Profiles []vk.User  `json:"profiles,omitempty"`
	Groups   []vk.Group `json:"groups,omitempty"`
}

func (VideoGetResponseExtended) isVideoGet() {}

// Get Returns detailed information about videos.
func (v Video) Get(params VideoGetParams) (VideoGetResponse, error) {
	r, err := v.API.Request("video.get", params)
	if err != nil {
		return nil, err
	}

	var resp VideoGetResponse
	if params.Extended {
		var tmp VideoGetResponseExtended
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	} else {
		var tmp VideoGetResponseNormal
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// VideoEditParams are params for Video.Edit
type VideoEditParams struct {
	// ID of the user or community that owns the video.
	OwnerID int `url:"owner_id,omitempty"`
	// Video ID.
	VideoID int `url:"video_id"`
	// New video title.
	Name string `url:"name,omitempty"`
	// New video description.
	Desc string `url:"desc,omitempty"`
	// Privacy settings in a [vk.com/dev/privacy_setting|special format]. Privacy setting is available for videos uploaded to own profile by user.
	PrivacyView CSVStringSlice `url:"privacy_view,omitempty"`
	// Privacy settings for comments in a [vk.com/dev/privacy_setting|special format].
	PrivacyComment CSVStringSlice `url:"privacy_comment,omitempty"`
	// Disable comments for the group video.
	NoComments bool `url:"no_comments,omitempty"`
	// '1' — to repeat the playback of the video, '0' — to play the video once,
	Repeat bool `url:"repeat,omitempty"`
}

// Edit Edits information about a video on a user or community page.
func (v Video) Edit(params VideoEditParams) (bool, error) {
	r, err := v.API.Request("video.edit", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// VideoAddParams are params for Video.Add
type VideoAddParams struct {
	// identifier of a user or community to add a video to. Use a negative value to designate a community ID.
	TargetID int `url:"target_id,omitempty"`
	// Video ID.
	VideoID int `url:"video_id"`
	// ID of the user or community that owns the video. Use a negative value to designate a community ID.
	OwnerID int `url:"owner_id"`
}

// Add Adds a video to a user or community page.
func (v Video) Add(params VideoAddParams) (bool, error) {
	r, err := v.API.Request("video.add", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// VideoSaveParams are params for Video.Save
type VideoSaveParams struct {
	// Name of the video.
	Name string `url:"name,omitempty"`
	// Description of the video.
	Description string `url:"description,omitempty"`
	// '1' — to designate the video as private (send it via a private message), the video will not appear on the user's video list and will not be available by ID for other users, '0' — not to designate the video as private
	IsPrivate bool `url:"is_private,omitempty"`
	// '1' — to post the saved video on a user's wall, '0' — not to post the saved video on a user's wall
	Wallpost bool `url:"wallpost,omitempty"`
	// URL for embedding the video from an external website.
	Link string `url:"link,omitempty"`
	// ID of the community in which the video will be saved. By default, the current user's page.
	GroupID int `url:"group_id,omitempty"`
	// ID of the album to which the saved video will be added.
	AlbumID        int            `url:"album_id,omitempty"`
	PrivacyView    CSVStringSlice `url:"privacy_view,omitempty"`
	PrivacyComment CSVStringSlice `url:"privacy_comment,omitempty"`
	NoComments     bool           `url:"no_comments,omitempty"`
	// '1' — to repeat the playback of the video, '0' — to play the video once,
	Repeat bool `url:"repeat,omitempty"`
}

// VideoSaveResponse is response for Video.Save
//easyjson:json
type VideoSaveResponse struct {
	// URL for the video uploading
	UploadURL string `json:"upload_url,omitempty"`
	// Video ID
	VideoID int `json:"video_id,omitempty"`
	// Video owner ID
	OwnerID int `json:"owner_id,omitempty"`
	// Video title
	Title string `json:"title,omitempty"`
	// Video description
	Description string `json:"description,omitempty"`
}

// Save Returns a server address (required for upload) and video data.
func (v Video) Save(params VideoSaveParams) (*VideoSaveResponse, error) {
	r, err := v.API.Request("video.save", params)
	if err != nil {
		return nil, err
	}

	var resp VideoSaveResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// VideoDeleteParams are params for Video.Delete
type VideoDeleteParams struct {
	// Video ID.
	VideoID int `url:"video_id"`
	// ID of the user or community that owns the video.
	OwnerID  int `url:"owner_id,omitempty"`
	TargetID int `url:"target_id,omitempty"`
}

// Delete Deletes a video from a user or community page.
func (v Video) Delete(params VideoDeleteParams) (bool, error) {
	r, err := v.API.Request("video.delete", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// VideoRestoreParams are params for Video.Restore
type VideoRestoreParams struct {
	// Video ID.
	VideoID int `url:"video_id"`
	// ID of the user or community that owns the video.
	OwnerID int `url:"owner_id,omitempty"`
}

// Restore Restores a previously deleted video.
func (v Video) Restore(params VideoRestoreParams) (bool, error) {
	r, err := v.API.Request("video.restore", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// VideoSearchParams are params for Video.Search
type VideoSearchParams struct {
	// Search query string (e.g., 'The Beatles').
	Q string `url:"q"`
	// Sort order: '1' — by duration, '2' — by relevance, '0' — by date added
	Sort int `url:"sort,omitempty"`
	// If not null, only searches for high-definition videos.
	Hd int `url:"hd,omitempty"`
	// '1' — to disable the Safe Search filter, '0' — to enable the Safe Search filter
	Adult bool `url:"adult,omitempty"`
	// Filters to apply: 'youtube' — return YouTube videos only, 'vimeo' — return Vimeo videos only, 'short' — return short videos only, 'long' — return long videos only
	Filters CSVStringSlice `url:"filters,omitempty"`
	//
	SearchOwn bool `url:"search_own,omitempty"`
	// Offset needed to return a specific subset of videos.
	Offset  int `url:"offset,omitempty"`
	Longer  int `url:"longer,omitempty"`
	Shorter int `url:"shorter,omitempty"`
	// Number of videos to return.
	Count int `url:"count,omitempty"`
	//
	Extended bool `url:"extended,omitempty"`
}

// VideoSearchResponse is response for Video.Search
// Either VideoSearchResponseNormal or VideoSearchResponseExtended, depending on Extended flag
type VideoSearchResponse interface {
	isVideoSearch()
}

// VideoSearchResponseNormal is non-extended version of VideoSearchResponse
//easyjson:json
type VideoSearchResponseNormal struct {
	// Total number
	Count int        `json:"count,omitempty"`
	Items []vk.Video `json:"items,omitempty"`
}

func (VideoSearchResponseNormal) isVideoSearch() {}

// VideoSearchResponseExtended is extended version of VideoSearchResponse
//easyjson:json
type VideoSearchResponseExtended struct {
	// Total number
	Count    int        `json:"count,omitempty"`
	Items    []vk.Video `json:"items,omitempty"`
	Profiles []vk.User  `json:"profiles,omitempty"`
	Groups   []vk.Group `json:"groups,omitempty"`
}

func (VideoSearchResponseExtended) isVideoSearch() {}

// Search Returns a list of videos under the set search criterion.
func (v Video) Search(params VideoSearchParams) (VideoSearchResponse, error) {
	r, err := v.API.Request("video.search", params)
	if err != nil {
		return nil, err
	}

	var resp VideoSearchResponse
	if params.Extended {
		var tmp VideoSearchResponseExtended
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	} else {
		var tmp VideoSearchResponseNormal
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// VideoGetUserVideosParams are params for Video.GetUserVideos
type VideoGetUserVideosParams struct {
	// User ID.
	UserID int `url:"user_id,omitempty"`
	// Offset needed to return a specific subset of videos.
	Offset int `url:"offset,omitempty"`
	// Number of videos to return.
	Count int `url:"count,omitempty"`
	//
	Extended bool `url:"extended,omitempty"`
}

// VideoGetUserVideosResponse is response for Video.GetUserVideos
// Either VideoGetUserVideosResponseNormal or VideoGetUserVideosResponseExtended, depending on Extended flag
type VideoGetUserVideosResponse interface {
	isVideoGetUserVideos()
}

// VideoGetUserVideosResponseNormal is non-extended version of VideoGetUserVideosResponse
//easyjson:json
type VideoGetUserVideosResponseNormal struct {
	// Total number
	Count int        `json:"count,omitempty"`
	Items []vk.Video `json:"items,omitempty"`
}

func (VideoGetUserVideosResponseNormal) isVideoGetUserVideos() {}

// VideoGetUserVideosResponseExtended is extended version of VideoGetUserVideosResponse
//easyjson:json
type VideoGetUserVideosResponseExtended struct {
	// Total number
	Count    int        `json:"count,omitempty"`
	Items    []vk.Video `json:"items,omitempty"`
	Profiles []vk.User  `json:"profiles,omitempty"`
	Groups   []vk.Group `json:"groups,omitempty"`
}

func (VideoGetUserVideosResponseExtended) isVideoGetUserVideos() {}

// GetUserVideos Returns list of videos in which the user is tagged.
func (v Video) GetUserVideos(params VideoGetUserVideosParams) (VideoGetUserVideosResponse, error) {
	r, err := v.API.Request("video.getUserVideos", params)
	if err != nil {
		return nil, err
	}

	var resp VideoGetUserVideosResponse
	if params.Extended {
		var tmp VideoGetUserVideosResponseExtended
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	} else {
		var tmp VideoGetUserVideosResponseNormal
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// VideoGetAlbumsParams are params for Video.GetAlbums
type VideoGetAlbumsParams struct {
	// ID of the user or community that owns the video album(s).
	OwnerID int `url:"owner_id,omitempty"`
	// Offset needed to return a specific subset of video albums.
	Offset int `url:"offset,omitempty"`
	// Number of video albums to return.
	Count int `url:"count,omitempty"`
	// '1' — to return additional information about album privacy settings for the current user
	Extended bool `url:"extended,omitempty"`
}

// VideoGetAlbumsResponse is response for Video.GetAlbums
// Either VideoGetAlbumsResponseNormal or VideoGetAlbumsResponseExtended, depending on Extended flag
type VideoGetAlbumsResponse interface {
	isVideoGetAlbums()
}

// VideoGetAlbumsResponseNormal is non-extended version of VideoGetAlbumsResponse
//easyjson:json
type VideoGetAlbumsResponseNormal struct {
	// Total number
	Count int `json:"count,omitempty"`
	Items []genTODOType/* objects.json#/definitions/video_video_album_full */ `json:"items,omitempty"`
}

func (VideoGetAlbumsResponseNormal) isVideoGetAlbums() {}

// VideoGetAlbumsResponseExtended is extended version of VideoGetAlbumsResponse
//easyjson:json
type VideoGetAlbumsResponseExtended struct {
	// Total number
	Count int `json:"count,omitempty"`
	Items []genTODOType/* objects.json#/definitions/video_video_album_full */ `json:"items,omitempty"`
}

func (VideoGetAlbumsResponseExtended) isVideoGetAlbums() {}

// GetAlbums Returns a list of video albums owned by a user or community.
func (v Video) GetAlbums(params VideoGetAlbumsParams) (VideoGetAlbumsResponse, error) {
	r, err := v.API.Request("video.getAlbums", params)
	if err != nil {
		return nil, err
	}

	var resp VideoGetAlbumsResponse
	if params.Extended {
		var tmp VideoGetAlbumsResponseExtended
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	} else {
		var tmp VideoGetAlbumsResponseNormal
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// VideoGetAlbumByIDParams are params for Video.GetAlbumByID
type VideoGetAlbumByIDParams struct {
	// identifier of a user or community to add a video to. Use a negative value to designate a community ID.
	OwnerID int `url:"owner_id,omitempty"`
	// Album ID.
	AlbumID int `url:"album_id"`
}

// VideoGetAlbumByIDResponse is response for Video.GetAlbumByID
//easyjson:json
type VideoGetAlbumByIDResponse genTODOType /* objects.json#/definitions/video_video_album_full */
// GetAlbumByID Returns video album info
func (v Video) GetAlbumByID(params VideoGetAlbumByIDParams) (*VideoGetAlbumByIDResponse, error) {
	r, err := v.API.Request("video.getAlbumById", params)
	if err != nil {
		return nil, err
	}

	var resp VideoGetAlbumByIDResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// VideoAddAlbumParams are params for Video.AddAlbum
type VideoAddAlbumParams struct {
	// Community ID (if the album will be created in a community).
	GroupID int `url:"group_id,omitempty"`
	// Album title.
	Title string `url:"title,omitempty"`
	// new access permissions for the album. Possible values: , *'0' – all users,, *'1' – friends only,, *'2' – friends and friends of friends,, *'3' – "only me".
	Privacy CSVStringSlice `url:"privacy,omitempty"`
}

// VideoAddAlbumResponse is response for Video.AddAlbum
//easyjson:json
type VideoAddAlbumResponse struct {
	// Created album ID
	AlbumID int `json:"album_id,omitempty"`
}

// AddAlbum Creates an empty album for videos.
func (v Video) AddAlbum(params VideoAddAlbumParams) (*VideoAddAlbumResponse, error) {
	r, err := v.API.Request("video.addAlbum", params)
	if err != nil {
		return nil, err
	}

	var resp VideoAddAlbumResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// VideoEditAlbumParams are params for Video.EditAlbum
type VideoEditAlbumParams struct {
	// Community ID (if the album edited is owned by a community).
	GroupID int `url:"group_id,omitempty"`
	// Album ID.
	AlbumID int `url:"album_id"`
	// New album title.
	Title string `url:"title"`
	// new access permissions for the album. Possible values: , *'0' – all users,, *'1' – friends only,, *'2' – friends and friends of friends,, *'3' – "only me".
	Privacy CSVStringSlice `url:"privacy,omitempty"`
}

// EditAlbum Edits the title of a video album.
func (v Video) EditAlbum(params VideoEditAlbumParams) (bool, error) {
	r, err := v.API.Request("video.editAlbum", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// VideoDeleteAlbumParams are params for Video.DeleteAlbum
type VideoDeleteAlbumParams struct {
	// Community ID (if the album is owned by a community).
	GroupID int `url:"group_id,omitempty"`
	// Album ID.
	AlbumID int `url:"album_id"`
}

// DeleteAlbum Deletes a video album.
func (v Video) DeleteAlbum(params VideoDeleteAlbumParams) (bool, error) {
	r, err := v.API.Request("video.deleteAlbum", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// VideoReorderAlbumsParams are params for Video.ReorderAlbums
type VideoReorderAlbumsParams struct {
	// ID of the user or community that owns the albums..
	OwnerID int `url:"owner_id,omitempty"`
	// Album ID.
	AlbumID int `url:"album_id"`
	// ID of the album before which the album in question shall be placed.
	Before int `url:"before,omitempty"`
	// ID of the album after which the album in question shall be placed.
	After int `url:"after,omitempty"`
}

// ReorderAlbums Reorders the album in the list of user video albums.
func (v Video) ReorderAlbums(params VideoReorderAlbumsParams) (bool, error) {
	r, err := v.API.Request("video.reorderAlbums", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// VideoReorderVideosParams are params for Video.ReorderVideos
type VideoReorderVideosParams struct {
	// ID of the user or community that owns the album with videos.
	TargetID int `url:"target_id,omitempty"`
	// ID of the video album.
	AlbumID int `url:"album_id,omitempty"`
	// ID of the user or community that owns the video.
	OwnerID int `url:"owner_id"`
	// ID of the video.
	VideoID int `url:"video_id"`
	// ID of the user or community that owns the video before which the video in question shall be placed.
	BeforeOwnerID int `url:"before_owner_id,omitempty"`
	// ID of the video before which the video in question shall be placed.
	BeforeVideoID int `url:"before_video_id,omitempty"`
	// ID of the user or community that owns the video after which the photo in question shall be placed.
	AfterOwnerID int `url:"after_owner_id,omitempty"`
	// ID of the video after which the photo in question shall be placed.
	AfterVideoID int `url:"after_video_id,omitempty"`
}

// ReorderVideos Reorders the video in the video album.
func (v Video) ReorderVideos(params VideoReorderVideosParams) (bool, error) {
	r, err := v.API.Request("video.reorderVideos", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// VideoAddToAlbumParams are params for Video.AddToAlbum
type VideoAddToAlbumParams struct {
	TargetID int         `url:"target_id,omitempty"`
	AlbumID  int         `url:"album_id,omitempty"`
	AlbumIDs CSVIntSlice `url:"album_ids,omitempty"`
	OwnerID  int         `url:"owner_id"`
	VideoID  int         `url:"video_id"`
}

// AddToAlbum does video.addToAlbum
func (v Video) AddToAlbum(params VideoAddToAlbumParams) (bool, error) {
	r, err := v.API.Request("video.addToAlbum", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// VideoRemoveFromAlbumParams are params for Video.RemoveFromAlbum
type VideoRemoveFromAlbumParams struct {
	TargetID int         `url:"target_id,omitempty"`
	AlbumID  int         `url:"album_id,omitempty"`
	AlbumIDs CSVIntSlice `url:"album_ids,omitempty"`
	OwnerID  int         `url:"owner_id"`
	VideoID  int         `url:"video_id"`
}

// RemoveFromAlbum does video.removeFromAlbum
func (v Video) RemoveFromAlbum(params VideoRemoveFromAlbumParams) (bool, error) {
	r, err := v.API.Request("video.removeFromAlbum", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// VideoGetAlbumsByVideoParams are params for Video.GetAlbumsByVideo
type VideoGetAlbumsByVideoParams struct {
	TargetID int `url:"target_id,omitempty"`
	OwnerID  int `url:"owner_id"`
	VideoID  int `url:"video_id"`
	//
	Extended bool `url:"extended,omitempty"`
}

// VideoGetAlbumsByVideoResponse is response for Video.GetAlbumsByVideo
// Either VideoGetAlbumsByVideoResponseNormal or VideoGetAlbumsByVideoResponseExtended, depending on Extended flag
type VideoGetAlbumsByVideoResponse interface {
	isVideoGetAlbumsByVideo()
}

// VideoGetAlbumsByVideoResponseNormal is non-extended version of VideoGetAlbumsByVideoResponse
// Album ID
//easyjson:json
type VideoGetAlbumsByVideoResponseNormal []int

func (VideoGetAlbumsByVideoResponseNormal) isVideoGetAlbumsByVideo() {}

// VideoGetAlbumsByVideoResponseExtended is extended version of VideoGetAlbumsByVideoResponse
//easyjson:json
type VideoGetAlbumsByVideoResponseExtended struct {
	// Total number
	Count int `json:"count,omitempty"`
	Items []genTODOType/* objects.json#/definitions/video_video_album_full */ `json:"items,omitempty"`
}

func (VideoGetAlbumsByVideoResponseExtended) isVideoGetAlbumsByVideo() {}

// GetAlbumsByVideo does video.getAlbumsByVideo
func (v Video) GetAlbumsByVideo(params VideoGetAlbumsByVideoParams) (VideoGetAlbumsByVideoResponse, error) {
	r, err := v.API.Request("video.getAlbumsByVideo", params)
	if err != nil {
		return nil, err
	}

	var resp VideoGetAlbumsByVideoResponse
	if params.Extended {
		var tmp VideoGetAlbumsByVideoResponseExtended
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	} else {
		var tmp VideoGetAlbumsByVideoResponseNormal
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// VideoGetCommentsParams are params for Video.GetComments
type VideoGetCommentsParams struct {
	// ID of the user or community that owns the video.
	OwnerID int `url:"owner_id,omitempty"`
	// Video ID.
	VideoID int `url:"video_id"`
	// '1' — to return an additional 'likes' field
	NeedLikes      bool `url:"need_likes,omitempty"`
	StartCommentID int  `url:"start_comment_id,omitempty"`
	// Offset needed to return a specific subset of comments.
	Offset int `url:"offset,omitempty"`
	// Number of comments to return.
	Count int `url:"count,omitempty"`
	// Sort order: 'asc' — oldest comment first, 'desc' — newest comment first
	Sort     string `url:"sort,omitempty"`
	Extended bool   `url:"extended,omitempty"`
}

// VideoGetCommentsResponse is response for Video.GetComments
// Either VideoGetCommentsResponseNormal or VideoGetCommentsResponseExtended, depending on Extended flag
type VideoGetCommentsResponse interface {
	isVideoGetComments()
}

// VideoGetCommentsResponseNormal is non-extended version of VideoGetCommentsResponse
//easyjson:json
type VideoGetCommentsResponseNormal struct {
	// Total number
	Count    int          `json:"count,omitempty"`
	Items    []vk.Comment `json:"items,omitempty"`
	Profiles []vk.User    `json:"profiles,omitempty"`
	Groups   []vk.Group   `json:"groups,omitempty"`
}

func (VideoGetCommentsResponseNormal) isVideoGetComments() {}

// VideoGetCommentsResponseExtended is extended version of VideoGetCommentsResponse
//easyjson:json
type VideoGetCommentsResponseExtended struct {
	// Total number
	Count int          `json:"count,omitempty"`
	Items []vk.Comment `json:"items,omitempty"`
}

func (VideoGetCommentsResponseExtended) isVideoGetComments() {}

// GetComments Returns a list of comments on a video.
func (v Video) GetComments(params VideoGetCommentsParams) (VideoGetCommentsResponse, error) {
	r, err := v.API.Request("video.getComments", params)
	if err != nil {
		return nil, err
	}

	var resp VideoGetCommentsResponse
	if params.Extended {
		var tmp VideoGetCommentsResponseExtended
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	} else {
		var tmp VideoGetCommentsResponseNormal
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// VideoCreateCommentParams are params for Video.CreateComment
type VideoCreateCommentParams struct {
	// ID of the user or community that owns the video.
	OwnerID int `url:"owner_id,omitempty"`
	// Video ID.
	VideoID int `url:"video_id"`
	// New comment text.
	Message string `url:"message,omitempty"`
	// List of objects attached to the comment, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, '<owner_id>' — ID of the media attachment owner. '<media_id>' — Media attachment ID. Example: "photo100172_166443618,photo66748_265827614"
	Attachments CSVStringSlice `url:"attachments,omitempty"`
	// '1' — to post the comment from a community name (only if 'owner_id'<0)
	FromGroup bool `url:"from_group,omitempty"`
	//
	ReplyToComment int    `url:"reply_to_comment,omitempty"`
	StickerID      int    `url:"sticker_id,omitempty"`
	GUID           string `url:"guid,omitempty"`
}

// VideoCreateCommentResponse is response for Video.CreateComment
// Created comment ID
type VideoCreateCommentResponse int

// CreateComment Adds a new comment on a video.
func (v Video) CreateComment(params VideoCreateCommentParams) (VideoCreateCommentResponse, error) {
	r, err := v.API.Request("video.createComment", params)
	if err != nil {
		return 0, err
	}

	var resp VideoCreateCommentResponse

	var cnv int
	cnv, err = strconv.Atoi(string(r))
	resp = VideoCreateCommentResponse(cnv)

	if err != nil {
		return 0, err
	}
	return resp, nil
}

// VideoDeleteCommentParams are params for Video.DeleteComment
type VideoDeleteCommentParams struct {
	// ID of the user or community that owns the video.
	OwnerID int `url:"owner_id,omitempty"`
	// ID of the comment to be deleted.
	CommentID int `url:"comment_id"`
}

// DeleteComment Deletes a comment on a video.
func (v Video) DeleteComment(params VideoDeleteCommentParams) (bool, error) {
	r, err := v.API.Request("video.deleteComment", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// VideoRestoreCommentParams are params for Video.RestoreComment
type VideoRestoreCommentParams struct {
	// ID of the user or community that owns the video.
	OwnerID int `url:"owner_id,omitempty"`
	// ID of the deleted comment.
	CommentID int `url:"comment_id"`
}

// RestoreComment Restores a previously deleted comment on a video.
func (v Video) RestoreComment(params VideoRestoreCommentParams) (bool, error) {
	r, err := v.API.Request("video.restoreComment", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// VideoEditCommentParams are params for Video.EditComment
type VideoEditCommentParams struct {
	// ID of the user or community that owns the video.
	OwnerID int `url:"owner_id,omitempty"`
	// Comment ID.
	CommentID int `url:"comment_id"`
	// New comment text.
	Message string `url:"message,omitempty"`
	// List of objects attached to the comment, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, '<owner_id>' — ID of the media attachment owner. '<media_id>' — Media attachment ID. Example: "photo100172_166443618,photo66748_265827614"
	Attachments CSVStringSlice `url:"attachments,omitempty"`
}

// EditComment Edits the text of a comment on a video.
func (v Video) EditComment(params VideoEditCommentParams) (bool, error) {
	r, err := v.API.Request("video.editComment", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// VideoGetTagsParams are params for Video.GetTags
type VideoGetTagsParams struct {
	// ID of the user or community that owns the video.
	OwnerID int `url:"owner_id,omitempty"`
	// Video ID.
	VideoID int `url:"video_id"`
}

// VideoGetTagsResponse is response for Video.GetTags
//easyjson:json
type VideoGetTagsResponse []genTODOType /* objects.json#/definitions/video_video_tag */
// GetTags Returns a list of tags on a video.
func (v Video) GetTags(params VideoGetTagsParams) (VideoGetTagsResponse, error) {
	r, err := v.API.Request("video.getTags", params)
	if err != nil {
		return nil, err
	}

	var resp VideoGetTagsResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// VideoPutTagParams are params for Video.PutTag
type VideoPutTagParams struct {
	// ID of the user to be tagged.
	UserID int `url:"user_id"`
	// ID of the user or community that owns the video.
	OwnerID int `url:"owner_id,omitempty"`
	// Video ID.
	VideoID int `url:"video_id"`
	// Tag text.
	TaggedName string `url:"tagged_name,omitempty"`
}

// VideoPutTagResponse is response for Video.PutTag
// Created tag ID
type VideoPutTagResponse int

// PutTag Adds a tag on a video.
func (v Video) PutTag(params VideoPutTagParams) (VideoPutTagResponse, error) {
	r, err := v.API.Request("video.putTag", params)
	if err != nil {
		return 0, err
	}

	var resp VideoPutTagResponse

	var cnv int
	cnv, err = strconv.Atoi(string(r))
	resp = VideoPutTagResponse(cnv)

	if err != nil {
		return 0, err
	}
	return resp, nil
}

// VideoRemoveTagParams are params for Video.RemoveTag
type VideoRemoveTagParams struct {
	// Tag ID.
	TagID int `url:"tag_id"`
	// ID of the user or community that owns the video.
	OwnerID int `url:"owner_id,omitempty"`
	// Video ID.
	VideoID int `url:"video_id"`
}

// RemoveTag Removes a tag from a video.
func (v Video) RemoveTag(params VideoRemoveTagParams) (bool, error) {
	r, err := v.API.Request("video.removeTag", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// VideoGetNewTagsParams are params for Video.GetNewTags
type VideoGetNewTagsParams struct {
	// Offset needed to return a specific subset of videos.
	Offset int `url:"offset,omitempty"`
	// Number of videos to return.
	Count int `url:"count,omitempty"`
}

// VideoGetNewTagsResponse is response for Video.GetNewTags
//easyjson:json
type VideoGetNewTagsResponse struct {
	// Total number
	Count int        `json:"count,omitempty"`
	Items []vk.Video `json:"items,omitempty"`
}

// GetNewTags Returns a list of videos with tags that have not been viewed.
func (v Video) GetNewTags(params VideoGetNewTagsParams) (*VideoGetNewTagsResponse, error) {
	r, err := v.API.Request("video.getNewTags", params)
	if err != nil {
		return nil, err
	}

	var resp VideoGetNewTagsResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// VideoReportParams are params for Video.Report
type VideoReportParams struct {
	// ID of the user or community that owns the video.
	OwnerID int `url:"owner_id"`
	// Video ID.
	VideoID int `url:"video_id"`
	// Reason for the complaint: '0' – spam, '1' – child pornography, '2' – extremism, '3' – violence, '4' – drug propaganda, '5' – adult material, '6' – insult, abuse
	Reason int `url:"reason,omitempty"`
	// Comment describing the complaint.
	Comment string `url:"comment,omitempty"`
	// (If the video was found in search results.) Search query string.
	SearchQuery string `url:"search_query,omitempty"`
}

// Report Reports (submits a complaint about) a video.
func (v Video) Report(params VideoReportParams) (bool, error) {
	r, err := v.API.Request("video.report", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// VideoReportCommentParams are params for Video.ReportComment
type VideoReportCommentParams struct {
	// ID of the user or community that owns the video.
	OwnerID int `url:"owner_id"`
	// ID of the comment being reported.
	CommentID int `url:"comment_id"`
	// Reason for the complaint: , 0 – spam , 1 – child pornography , 2 – extremism , 3 – violence , 4 – drug propaganda , 5 – adult material , 6 – insult, abuse
	Reason int `url:"reason,omitempty"`
}

// ReportComment Reports (submits a complaint about) a comment on a video.
func (v Video) ReportComment(params VideoReportCommentParams) (bool, error) {
	r, err := v.API.Request("video.reportComment", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// VideoGetCatalogParams are params for Video.GetCatalog
type VideoGetCatalogParams struct {
	// number of catalog blocks to return.
	Count int `url:"count,omitempty"`
	// number of videos in each block.
	ItemsCount int `url:"items_count,omitempty"`
	// parameter for requesting the next results page. Value for transmitting here is returned in the 'next' field in a reply.
	From string `url:"from,omitempty"`
	// list of requested catalog sections
	Filters CSVStringSlice `url:"filters,omitempty"`
	// 1 – return additional infor about users and communities in profiles and groups fields.
	Extended bool `url:"extended,omitempty"`
}

// VideoGetCatalogResponse is response for Video.GetCatalog
// Either VideoGetCatalogResponseNormal or VideoGetCatalogResponseExtended, depending on Extended flag
type VideoGetCatalogResponse interface {
	isVideoGetCatalog()
}

// VideoGetCatalogResponseNormal is non-extended version of VideoGetCatalogResponse
//easyjson:json
type VideoGetCatalogResponseNormal struct {
	Items []genTODOType/* objects.json#/definitions/video_cat_block */ `json:"items,omitempty"`
	// New value for _from_ parameter
	Next string `json:"next,omitempty"`
}

func (VideoGetCatalogResponseNormal) isVideoGetCatalog() {}

// VideoGetCatalogResponseExtended is extended version of VideoGetCatalogResponse
//easyjson:json
type VideoGetCatalogResponseExtended struct {
	Items    []genTODOType/* objects.json#/definitions/video_cat_block */ `json:"items,omitempty"`
	Profiles []vk.User  `json:"profiles,omitempty"`
	Groups   []vk.Group `json:"groups,omitempty"`
	// New value for _from_ parameter
	Next string `json:"next,omitempty"`
}

func (VideoGetCatalogResponseExtended) isVideoGetCatalog() {}

// GetCatalog Returns video catalog
func (v Video) GetCatalog(params VideoGetCatalogParams) (VideoGetCatalogResponse, error) {
	r, err := v.API.Request("video.getCatalog", params)
	if err != nil {
		return nil, err
	}

	var resp VideoGetCatalogResponse
	if params.Extended {
		var tmp VideoGetCatalogResponseExtended
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	} else {
		var tmp VideoGetCatalogResponseNormal
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// VideoGetCatalogSectionParams are params for Video.GetCatalogSection
type VideoGetCatalogSectionParams struct {
	// 'id' value returned with a block by the '' method.
	SectionID string `url:"section_id"`
	// 'next' value returned with a block by the '' method.
	From string `url:"from"`
	// number of blocks to return.
	Count int `url:"count,omitempty"`
	// 1 – return additional infor about users and communities in profiles and groups fields.
	Extended bool `url:"extended,omitempty"`
}

// VideoGetCatalogSectionResponse is response for Video.GetCatalogSection
// Either VideoGetCatalogSectionResponseNormal or VideoGetCatalogSectionResponseExtended, depending on Extended flag
type VideoGetCatalogSectionResponse interface {
	isVideoGetCatalogSection()
}

// VideoGetCatalogSectionResponseNormal is non-extended version of VideoGetCatalogSectionResponse
//easyjson:json
type VideoGetCatalogSectionResponseNormal struct {
	Items []genTODOType/* objects.json#/definitions/video_cat_element */ `json:"items,omitempty"`
	// New value for _from_ parameter
	Next string `json:"next,omitempty"`
}

func (VideoGetCatalogSectionResponseNormal) isVideoGetCatalogSection() {}

// VideoGetCatalogSectionResponseExtended is extended version of VideoGetCatalogSectionResponse
//easyjson:json
type VideoGetCatalogSectionResponseExtended struct {
	Items    []genTODOType/* objects.json#/definitions/video_cat_element */ `json:"items,omitempty"`
	Profiles []vk.User  `json:"profiles,omitempty"`
	Groups   []vk.Group `json:"groups,omitempty"`
	// New value for _from_ parameter
	Next string `json:"next,omitempty"`
}

func (VideoGetCatalogSectionResponseExtended) isVideoGetCatalogSection() {}

// GetCatalogSection Returns a separate catalog section
func (v Video) GetCatalogSection(params VideoGetCatalogSectionParams) (VideoGetCatalogSectionResponse, error) {
	r, err := v.API.Request("video.getCatalogSection", params)
	if err != nil {
		return nil, err
	}

	var resp VideoGetCatalogSectionResponse
	if params.Extended {
		var tmp VideoGetCatalogSectionResponseExtended
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	} else {
		var tmp VideoGetCatalogSectionResponseNormal
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// VideoHideCatalogSectionParams are params for Video.HideCatalogSection
type VideoHideCatalogSectionParams struct {
	// 'id' value returned with a block to hide by the '' method.
	SectionID int `url:"section_id"`
}

// HideCatalogSection Hides a video catalog section from a user.
func (v Video) HideCatalogSection(params VideoHideCatalogSectionParams) (bool, error) {
	r, err := v.API.Request("video.hideCatalogSection", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}
