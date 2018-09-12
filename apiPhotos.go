package vkCallbackApi

import (
	"encoding/json"
	"strconv"
)

// APIPhotos implements VK API namespace `photos`
type APIPhotos struct {
	API *API
}

// PhotosCreateAlbumParams are params for APIPhotos.CreateAlbum
type PhotosCreateAlbumParams struct {
	// Album title.
	Title string `url:"title"`
	// ID of the community in which the album will be created.
	GroupID int `url:"group_id,omitempty"`
	// Album description.
	Description        string         `url:"description,omitempty"`
	PrivacyView        CSVStringSlice `url:"privacy_view,omitempty"`
	PrivacyComment     CSVStringSlice `url:"privacy_comment,omitempty"`
	UploadByAdminsOnly bool           `url:"upload_by_admins_only,omitempty"`
	CommentsDisabled   bool           `url:"comments_disabled,omitempty"`
}

// PhotosCreateAlbumResponse is response for APIPhotos.CreateAlbum
//easyjson:json
type PhotosCreateAlbumResponse genTODOType /* objects.json#/definitions/photos_photo_album_full */
// CreateAlbum Creates an empty photo album.
func (v APIPhotos) CreateAlbum(params PhotosCreateAlbumParams) (*PhotosCreateAlbumResponse, error) {
	r, err := v.API.Request("photos.createAlbum", params)
	if err != nil {
		return nil, err
	}

	var resp PhotosCreateAlbumResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PhotosEditAlbumParams are params for APIPhotos.EditAlbum
type PhotosEditAlbumParams struct {
	// ID of the photo album to be edited.
	AlbumID int `url:"album_id"`
	// New album title.
	Title string `url:"title,omitempty"`
	// New album description.
	Description string `url:"description,omitempty"`
	// ID of the user or community that owns the album.
	OwnerID            int            `url:"owner_id,omitempty"`
	PrivacyView        CSVStringSlice `url:"privacy_view,omitempty"`
	PrivacyComment     CSVStringSlice `url:"privacy_comment,omitempty"`
	UploadByAdminsOnly bool           `url:"upload_by_admins_only,omitempty"`
	CommentsDisabled   bool           `url:"comments_disabled,omitempty"`
}

// EditAlbum Edits information about a photo album.
func (v APIPhotos) EditAlbum(params PhotosEditAlbumParams) (bool, error) {
	r, err := v.API.Request("photos.editAlbum", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// PhotosGetAlbumsParams are params for APIPhotos.GetAlbums
type PhotosGetAlbumsParams struct {
	// ID of the user or community that owns the albums.
	OwnerID int `url:"owner_id,omitempty"`
	// Album IDs.
	AlbumIDs CSVIntSlice `url:"album_ids,omitempty"`
	// Offset needed to return a specific subset of albums.
	Offset int `url:"offset,omitempty"`
	// Number of albums to return.
	Count int `url:"count,omitempty"`
	// '1' — to return system albums with negative IDs
	NeedSystem bool `url:"need_system,omitempty"`
	// '1' — to return an additional 'thumb_src' field, '0' — (default)
	NeedCovers bool `url:"need_covers,omitempty"`
	// '1' — to return photo sizes in a
	PhotoSizes bool `url:"photo_sizes,omitempty"`
}

// PhotosGetAlbumsResponse is response for APIPhotos.GetAlbums
//easyjson:json
type PhotosGetAlbumsResponse struct {
	// Total number
	Count int `json:"count,omitempty"`
	Items []genTODOType/* objects.json#/definitions/photos_photo_album_full */ `json:"items,omitempty"`
}

// GetAlbums Returns a list of a user's or community's photo albums.
func (v APIPhotos) GetAlbums(params PhotosGetAlbumsParams) (*PhotosGetAlbumsResponse, error) {
	r, err := v.API.Request("photos.getAlbums", params)
	if err != nil {
		return nil, err
	}

	var resp PhotosGetAlbumsResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PhotosGetParams are params for APIPhotos.Get
type PhotosGetParams struct {
	// ID of the user or community that owns the photos. Use a negative value to designate a community ID.
	OwnerID int `url:"owner_id,omitempty"`
	// Photo album ID. To return information about photos from service albums, use the following string values: 'profile, wall, saved'.
	AlbumID string `url:"album_id,omitempty"`
	// Photo IDs.
	PhotoIDs CSVStringSlice `url:"photo_ids,omitempty"`
	// Sort order: '1' — reverse chronological, '0' — chronological
	Rev bool `url:"rev,omitempty"`
	// '1' — to return additional 'likes', 'comments', and 'tags' fields, '0' — (default)
	Extended bool `url:"extended,omitempty"`
	// Type of feed obtained in 'feed' field of the method.
	FeedType string `url:"feed_type,omitempty"`
	// unixtime, that can be obtained with [vk.com/dev/newsfeed.get|newsfeed.get] method in date field to get all photos uploaded by the user on a specific day, or photos the user has been tagged on. Also, 'uid' parameter of the user the event happened with shall be specified.
	Feed int `url:"feed,omitempty"`
	// '1' — to return photo sizes in a [vk.com/dev/photo_sizes|special format]
	PhotoSizes bool `url:"photo_sizes,omitempty"`
	Offset     int  `url:"offset,omitempty"`
	Count      int  `url:"count,omitempty"`
}

// PhotosGetResponse is response for APIPhotos.Get
// Either PhotosGetResponseNormal or PhotosGetResponseExtended, depending on Extended flag
type PhotosGetResponse interface {
	isPhotosGet()
}

// PhotosGetResponseNormal is non-extended version of PhotosGetResponse
//easyjson:json
type PhotosGetResponseNormal struct {
	// Total number
	Count int     `json:"count,omitempty"`
	Items []Photo `json:"items,omitempty"`
}

func (PhotosGetResponseNormal) isPhotosGet() {}

// PhotosGetResponseExtended is extended version of PhotosGetResponse
//easyjson:json
type PhotosGetResponseExtended struct {
	// Total number
	Count int     `json:"count,omitempty"`
	Items []Photo `json:"items,omitempty"`
}

func (PhotosGetResponseExtended) isPhotosGet() {}

// Get Returns a list of a user's or community's photos.
func (v APIPhotos) Get(params PhotosGetParams) (PhotosGetResponse, error) {
	r, err := v.API.Request("photos.get", params)
	if err != nil {
		return nil, err
	}

	var resp PhotosGetResponse
	if params.Extended {
		var tmp PhotosGetResponseExtended
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	} else {
		var tmp PhotosGetResponseNormal
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// PhotosGetAlbumsCountParams are params for APIPhotos.GetAlbumsCount
type PhotosGetAlbumsCountParams struct {
	// User ID.
	UserID int `url:"user_id,omitempty"`
	// Community ID.
	GroupID int `url:"group_id,omitempty"`
}

// PhotosGetAlbumsCountResponse is response for APIPhotos.GetAlbumsCount
// Albums number
type PhotosGetAlbumsCountResponse int

// GetAlbumsCount Returns the number of photo albums belonging to a user or community.
func (v APIPhotos) GetAlbumsCount(params PhotosGetAlbumsCountParams) (PhotosGetAlbumsCountResponse, error) {
	r, err := v.API.Request("photos.getAlbumsCount", params)
	if err != nil {
		return 0, err
	}

	var resp PhotosGetAlbumsCountResponse

	var cnv int
	cnv, err = strconv.Atoi(string(r))
	resp = PhotosGetAlbumsCountResponse(cnv)

	if err != nil {
		return 0, err
	}
	return resp, nil
}

// PhotosGetByIDParams are params for APIPhotos.GetByID
type PhotosGetByIDParams struct {
	// IDs separated with a comma, that are IDs of users who posted photos and IDs of photos themselves with an underscore character between such IDs. To get information about a photo in the group album, you shall specify group ID instead of user ID. Example: "1_129207899,6492_135055734, , -20629724_271945303"
	Photos CSVStringSlice `url:"photos"`
	// '1' — to return additional fields, '0' — (default)
	Extended bool `url:"extended,omitempty"`
	// '1' — to return photo sizes in a
	PhotoSizes bool `url:"photo_sizes,omitempty"`
}

// PhotosGetByIDResponse is response for APIPhotos.GetByID
// Either PhotosGetByIDResponseNormal or PhotosGetByIDResponseExtended, depending on Extended flag
type PhotosGetByIDResponse interface {
	isPhotosGetByID()
}

// PhotosGetByIDResponseNormal is non-extended version of PhotosGetByIDResponse
//easyjson:json
type PhotosGetByIDResponseNormal []Photo

func (PhotosGetByIDResponseNormal) isPhotosGetByID() {}

// PhotosGetByIDResponseExtended is extended version of PhotosGetByIDResponse
//easyjson:json
type PhotosGetByIDResponseExtended []Photo

func (PhotosGetByIDResponseExtended) isPhotosGetByID() {}

// GetByID Returns information about photos by their IDs.
func (v APIPhotos) GetByID(params PhotosGetByIDParams) (PhotosGetByIDResponse, error) {
	r, err := v.API.Request("photos.getById", params)
	if err != nil {
		return nil, err
	}

	var resp PhotosGetByIDResponse
	if params.Extended {
		var tmp PhotosGetByIDResponseExtended
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	} else {
		var tmp PhotosGetByIDResponseNormal
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// PhotosGetUploadServerParams are params for APIPhotos.GetUploadServer
type PhotosGetUploadServerParams struct {
	// Album ID.
	AlbumID int `url:"album_id,omitempty"`
	// ID of community that owns the album (if the photo will be uploaded to a community album).
	GroupID int `url:"group_id,omitempty"`
}

// PhotosGetUploadServerResponse is response for APIPhotos.GetUploadServer
//easyjson:json
type PhotosGetUploadServerResponse struct {
	// URL to upload photo
	UploadURL string `json:"upload_url,omitempty"`
	// Album ID
	AlbumID int `json:"album_id,omitempty"`
	// User ID
	UserID int `json:"user_id,omitempty"`
}

// GetUploadServer Returns the server address for photo upload.
func (v APIPhotos) GetUploadServer(params PhotosGetUploadServerParams) (*PhotosGetUploadServerResponse, error) {
	r, err := v.API.Request("photos.getUploadServer", params)
	if err != nil {
		return nil, err
	}

	var resp PhotosGetUploadServerResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PhotosGetOwnerCoverPhotoUploadServerParams are params for APIPhotos.GetOwnerCoverPhotoUploadServer
type PhotosGetOwnerCoverPhotoUploadServerParams struct {
	// ID of community that owns the album (if the photo will be uploaded to a community album).
	GroupID int `url:"group_id,omitempty"`
	// X coordinate of the left-upper corner
	CropX int `url:"crop_x,omitempty"`
	// Y coordinate of the left-upper corner
	CropY int `url:"crop_y,omitempty"`
	// X coordinate of the right-bottom corner
	CropX2 int `url:"crop_x2,omitempty"`
	// Y coordinate of the right-bottom corner
	CropY2 int `url:"crop_y2,omitempty"`
}

// PhotosGetOwnerCoverPhotoUploadServerResponse is response for APIPhotos.GetOwnerCoverPhotoUploadServer
//easyjson:json
type PhotosGetOwnerCoverPhotoUploadServerResponse struct {
	// Upload URL
	UploadURL string `json:"upload_url,omitempty"`
}

// GetOwnerCoverPhotoUploadServer Returns the server address for owner cover upload.
func (v APIPhotos) GetOwnerCoverPhotoUploadServer(params PhotosGetOwnerCoverPhotoUploadServerParams) (*PhotosGetOwnerCoverPhotoUploadServerResponse, error) {
	r, err := v.API.Request("photos.getOwnerCoverPhotoUploadServer", params)
	if err != nil {
		return nil, err
	}

	var resp PhotosGetOwnerCoverPhotoUploadServerResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PhotosGetOwnerPhotoUploadServerParams are params for APIPhotos.GetOwnerPhotoUploadServer
type PhotosGetOwnerPhotoUploadServerParams struct {
	// identifier of a community or current user. "Note that community id must be negative. 'owner_id=1' – user, 'owner_id=-1' – community, "
	OwnerID int `url:"owner_id,omitempty"`
}

// PhotosGetOwnerPhotoUploadServerResponse is response for APIPhotos.GetOwnerPhotoUploadServer
//easyjson:json
type PhotosGetOwnerPhotoUploadServerResponse struct {
	// Upload URL
	UploadURL string `json:"upload_url,omitempty"`
}

// GetOwnerPhotoUploadServer Returns an upload server address for a profile or community photo.
func (v APIPhotos) GetOwnerPhotoUploadServer(params PhotosGetOwnerPhotoUploadServerParams) (*PhotosGetOwnerPhotoUploadServerResponse, error) {
	r, err := v.API.Request("photos.getOwnerPhotoUploadServer", params)
	if err != nil {
		return nil, err
	}

	var resp PhotosGetOwnerPhotoUploadServerResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PhotosGetChatUploadServerParams are params for APIPhotos.GetChatUploadServer
type PhotosGetChatUploadServerParams struct {
	// ID of the chat for which you want to upload a cover photo.
	ChatID int `url:"chat_id"`
	//
	CropX int `url:"crop_x,omitempty"`
	//
	CropY int `url:"crop_y,omitempty"`
	// Width (in pixels) of the photo after cropping.
	CropWidth int `url:"crop_width,omitempty"`
}

// PhotosGetChatUploadServerResponse is response for APIPhotos.GetChatUploadServer
//easyjson:json
type PhotosGetChatUploadServerResponse struct {
	// Upload URL
	UploadURL string `json:"upload_url,omitempty"`
}

// GetChatUploadServer Returns an upload link for chat cover pictures.
func (v APIPhotos) GetChatUploadServer(params PhotosGetChatUploadServerParams) (*PhotosGetChatUploadServerResponse, error) {
	r, err := v.API.Request("photos.getChatUploadServer", params)
	if err != nil {
		return nil, err
	}

	var resp PhotosGetChatUploadServerResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PhotosGetMarketUploadServerParams are params for APIPhotos.GetMarketUploadServer
type PhotosGetMarketUploadServerParams struct {
	// Community ID.
	GroupID int `url:"group_id"`
	// '1' if you want to upload the main item photo.
	MainPhoto bool `url:"main_photo,omitempty"`
	// X coordinate of the crop left upper corner.
	CropX int `url:"crop_x,omitempty"`
	// Y coordinate of the crop left upper corner.
	CropY int `url:"crop_y,omitempty"`
	// Width of the cropped photo in px.
	CropWidth int `url:"crop_width,omitempty"`
}

// PhotosGetMarketUploadServerResponse is response for APIPhotos.GetMarketUploadServer
//easyjson:json
type PhotosGetMarketUploadServerResponse struct {
	// Upload URL
	UploadURL string `json:"upload_url,omitempty"`
}

// GetMarketUploadServer Returns the server address for market photo upload.
func (v APIPhotos) GetMarketUploadServer(params PhotosGetMarketUploadServerParams) (*PhotosGetMarketUploadServerResponse, error) {
	r, err := v.API.Request("photos.getMarketUploadServer", params)
	if err != nil {
		return nil, err
	}

	var resp PhotosGetMarketUploadServerResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PhotosGetMarketAlbumUploadServerParams are params for APIPhotos.GetMarketAlbumUploadServer
type PhotosGetMarketAlbumUploadServerParams struct {
	// Community ID.
	GroupID int `url:"group_id"`
}

// PhotosGetMarketAlbumUploadServerResponse is response for APIPhotos.GetMarketAlbumUploadServer
//easyjson:json
type PhotosGetMarketAlbumUploadServerResponse struct {
	// Upload URL
	UploadURL string `json:"upload_url,omitempty"`
}

// GetMarketAlbumUploadServer Returns the server address for market album photo upload.
func (v APIPhotos) GetMarketAlbumUploadServer(params PhotosGetMarketAlbumUploadServerParams) (*PhotosGetMarketAlbumUploadServerResponse, error) {
	r, err := v.API.Request("photos.getMarketAlbumUploadServer", params)
	if err != nil {
		return nil, err
	}

	var resp PhotosGetMarketAlbumUploadServerResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PhotosSaveMarketPhotoParams are params for APIPhotos.SaveMarketPhoto
type PhotosSaveMarketPhotoParams struct {
	// Community ID.
	GroupID int `url:"group_id,omitempty"`
	// Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
	Photo string `url:"photo"`
	// Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
	Server int `url:"server"`
	// Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
	Hash string `url:"hash"`
	// Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
	CropData string `url:"crop_data,omitempty"`
	// Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
	CropHash string `url:"crop_hash,omitempty"`
}

// PhotosSaveMarketPhotoResponse is response for APIPhotos.SaveMarketPhoto
//easyjson:json
type PhotosSaveMarketPhotoResponse []Photo

// SaveMarketPhoto Saves market photos after successful uploading.
func (v APIPhotos) SaveMarketPhoto(params PhotosSaveMarketPhotoParams) (PhotosSaveMarketPhotoResponse, error) {
	r, err := v.API.Request("photos.saveMarketPhoto", params)
	if err != nil {
		return nil, err
	}

	var resp PhotosSaveMarketPhotoResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// PhotosSaveOwnerCoverPhotoParams are params for APIPhotos.SaveOwnerCoverPhoto
type PhotosSaveOwnerCoverPhotoParams struct {
	// Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
	Photo string `url:"photo"`
	// Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
	Hash string `url:"hash"`
}

// PhotosSaveOwnerCoverPhotoResponse is response for APIPhotos.SaveOwnerCoverPhoto
//easyjson:json
type PhotosSaveOwnerCoverPhotoResponse []BaseImage

// SaveOwnerCoverPhoto Saves cover photo after successful uploading.
func (v APIPhotos) SaveOwnerCoverPhoto(params PhotosSaveOwnerCoverPhotoParams) (PhotosSaveOwnerCoverPhotoResponse, error) {
	r, err := v.API.Request("photos.saveOwnerCoverPhoto", params)
	if err != nil {
		return nil, err
	}

	var resp PhotosSaveOwnerCoverPhotoResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// PhotosSaveMarketAlbumPhotoParams are params for APIPhotos.SaveMarketAlbumPhoto
type PhotosSaveMarketAlbumPhotoParams struct {
	// Community ID.
	GroupID int `url:"group_id"`
	// Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
	Photo string `url:"photo"`
	// Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
	Server int `url:"server"`
	// Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
	Hash string `url:"hash"`
}

// PhotosSaveMarketAlbumPhotoResponse is response for APIPhotos.SaveMarketAlbumPhoto
//easyjson:json
type PhotosSaveMarketAlbumPhotoResponse []Photo

// SaveMarketAlbumPhoto Saves market album photos after successful uploading.
func (v APIPhotos) SaveMarketAlbumPhoto(params PhotosSaveMarketAlbumPhotoParams) (PhotosSaveMarketAlbumPhotoResponse, error) {
	r, err := v.API.Request("photos.saveMarketAlbumPhoto", params)
	if err != nil {
		return nil, err
	}

	var resp PhotosSaveMarketAlbumPhotoResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// PhotosSaveOwnerPhotoParams are params for APIPhotos.SaveOwnerPhoto
type PhotosSaveOwnerPhotoParams struct {
	// parameter returned after [vk.com/dev/upload_files|photo upload].
	Server string `url:"server,omitempty"`
	// parameter returned after [vk.com/dev/upload_files|photo upload].
	Hash string `url:"hash,omitempty"`
	// parameter returned after [vk.com/dev/upload_files|photo upload].
	Photo string `url:"photo,omitempty"`
}

// PhotosSaveOwnerPhotoResponse is response for APIPhotos.SaveOwnerPhoto
//easyjson:json
type PhotosSaveOwnerPhotoResponse struct {
	// Parameter for saveProfilePhoto method
	PhotoHash string `json:"photo_hash,omitempty"`
	// Uploaded image url
	PhotoSrc string `json:"photo_src,omitempty"`
}

// SaveOwnerPhoto Saves a profile or community photo. Upload URL can be got with the [vk.com/dev/photos.getOwnerPhotoUploadServer|photos.getOwnerPhotoUploadServer] method.
func (v APIPhotos) SaveOwnerPhoto(params PhotosSaveOwnerPhotoParams) (*PhotosSaveOwnerPhotoResponse, error) {
	r, err := v.API.Request("photos.saveOwnerPhoto", params)
	if err != nil {
		return nil, err
	}

	var resp PhotosSaveOwnerPhotoResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PhotosSaveWallPhotoParams are params for APIPhotos.SaveWallPhoto
type PhotosSaveWallPhotoParams struct {
	// ID of the user on whose wall the photo will be saved.
	UserID int `url:"user_id,omitempty"`
	// ID of community on whose wall the photo will be saved.
	GroupID int `url:"group_id,omitempty"`
	// Parameter returned when the the photo is [vk.com/dev/upload_files|uploaded to the server].
	Photo  string `url:"photo"`
	Server int    `url:"server,omitempty"`
	Hash   string `url:"hash,omitempty"`
	// Geographical latitude, in degrees (from '-90' to '90').
	Latitude float32 `url:"latitude,omitempty"`
	// Geographical longitude, in degrees (from '-180' to '180').
	Longitude float32 `url:"longitude,omitempty"`
	// Text describing the photo. 2048 digits max.
	Caption string `url:"caption,omitempty"`
}

// PhotosSaveWallPhotoResponse is response for APIPhotos.SaveWallPhoto
//easyjson:json
type PhotosSaveWallPhotoResponse []Photo

// SaveWallPhoto Saves a photo to a user's or community's wall after being uploaded.
func (v APIPhotos) SaveWallPhoto(params PhotosSaveWallPhotoParams) (PhotosSaveWallPhotoResponse, error) {
	r, err := v.API.Request("photos.saveWallPhoto", params)
	if err != nil {
		return nil, err
	}

	var resp PhotosSaveWallPhotoResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// PhotosGetWallUploadServerParams are params for APIPhotos.GetWallUploadServer
type PhotosGetWallUploadServerParams struct {
	// ID of community to whose wall the photo will be uploaded.
	GroupID int `url:"group_id,omitempty"`
}

// PhotosGetWallUploadServerResponse is response for APIPhotos.GetWallUploadServer
//easyjson:json
type PhotosGetWallUploadServerResponse struct {
	// URL to upload photo
	UploadURL string `json:"upload_url,omitempty"`
	// Album ID
	AlbumID int `json:"album_id,omitempty"`
	// User ID
	UserID int `json:"user_id,omitempty"`
}

// GetWallUploadServer Returns the server address for photo upload onto a user's wall.
func (v APIPhotos) GetWallUploadServer(params PhotosGetWallUploadServerParams) (*PhotosGetWallUploadServerResponse, error) {
	r, err := v.API.Request("photos.getWallUploadServer", params)
	if err != nil {
		return nil, err
	}

	var resp PhotosGetWallUploadServerResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PhotosGetMessagesUploadServerParams are params for APIPhotos.GetMessagesUploadServer
type PhotosGetMessagesUploadServerParams struct {
	// Destination ID. "For user: 'User ID', e.g. '12345'. For chat: '2000000000' + 'Chat ID', e.g. '2000000001'. For community: '- Community ID', e.g. '-12345'. "
	PeerID int `url:"peer_id,omitempty"`
}

// PhotosGetMessagesUploadServerResponse is response for APIPhotos.GetMessagesUploadServer
//easyjson:json
type PhotosGetMessagesUploadServerResponse struct {
	// URL to upload photo
	UploadURL string `json:"upload_url,omitempty"`
	// Album ID
	AlbumID int `json:"album_id,omitempty"`
	// User ID
	UserID int `json:"user_id,omitempty"`
}

// GetMessagesUploadServer Returns the server address for photo upload in a private message for a user.
func (v APIPhotos) GetMessagesUploadServer(params PhotosGetMessagesUploadServerParams) (*PhotosGetMessagesUploadServerResponse, error) {
	r, err := v.API.Request("photos.getMessagesUploadServer", params)
	if err != nil {
		return nil, err
	}

	var resp PhotosGetMessagesUploadServerResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PhotosSaveMessagesPhotoParams are params for APIPhotos.SaveMessagesPhoto
type PhotosSaveMessagesPhotoParams struct {
	// Parameter returned when the photo is [vk.com/dev/upload_files|uploaded to the server].
	Photo  string `url:"photo"`
	Server int    `url:"server,omitempty"`
	Hash   string `url:"hash,omitempty"`
}

// PhotosSaveMessagesPhotoResponse is response for APIPhotos.SaveMessagesPhoto
//easyjson:json
type PhotosSaveMessagesPhotoResponse []Photo

// SaveMessagesPhoto Saves a photo after being successfully uploaded. URL obtained with [vk.com/dev/photos.getMessagesUploadServer|photos.getMessagesUploadServer] method.
func (v APIPhotos) SaveMessagesPhoto(params PhotosSaveMessagesPhotoParams) (PhotosSaveMessagesPhotoResponse, error) {
	r, err := v.API.Request("photos.saveMessagesPhoto", params)
	if err != nil {
		return nil, err
	}

	var resp PhotosSaveMessagesPhotoResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// PhotosReportParams are params for APIPhotos.Report
type PhotosReportParams struct {
	// ID of the user or community that owns the photo.
	OwnerID int `url:"owner_id"`
	// Photo ID.
	PhotoID int `url:"photo_id"`
	// Reason for the complaint: '0' – spam, '1' – child pornography, '2' – extremism, '3' – violence, '4' – drug propaganda, '5' – adult material, '6' – insult, abuse
	Reason int `url:"reason,omitempty"`
}

// Report Reports (submits a complaint about) a photo.
func (v APIPhotos) Report(params PhotosReportParams) (bool, error) {
	r, err := v.API.Request("photos.report", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// PhotosReportCommentParams are params for APIPhotos.ReportComment
type PhotosReportCommentParams struct {
	// ID of the user or community that owns the photo.
	OwnerID int `url:"owner_id"`
	// ID of the comment being reported.
	CommentID int `url:"comment_id"`
	// Reason for the complaint: '0' – spam, '1' – child pornography, '2' – extremism, '3' – violence, '4' – drug propaganda, '5' – adult material, '6' – insult, abuse
	Reason int `url:"reason,omitempty"`
}

// ReportComment Reports (submits a complaint about) a comment on a photo.
func (v APIPhotos) ReportComment(params PhotosReportCommentParams) (bool, error) {
	r, err := v.API.Request("photos.reportComment", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// PhotosSearchParams are params for APIPhotos.Search
type PhotosSearchParams struct {
	// Search query string.
	Q string `url:"q,omitempty"`
	// Geographical latitude, in degrees (from '-90' to '90').
	Lat float32 `url:"lat,omitempty"`
	// Geographical longitude, in degrees (from '-180' to '180').
	Long float32 `url:"long,omitempty"`
	//
	StartTime int `url:"start_time,omitempty"`
	//
	EndTime int `url:"end_time,omitempty"`
	// Sort order:
	Sort int `url:"sort,omitempty"`
	// Offset needed to return a specific subset of photos.
	Offset int `url:"offset,omitempty"`
	// Number of photos to return.
	Count int `url:"count,omitempty"`
	// Radius of search in meters (works very approximately). Available values: '10', '100', '800', '6000', '50000'.
	Radius int `url:"radius,omitempty"`
}

// PhotosSearchResponse is response for APIPhotos.Search
//easyjson:json
type PhotosSearchResponse struct {
	// Total number
	Count int     `json:"count,omitempty"`
	Items []Photo `json:"items,omitempty"`
}

// Search Returns a list of photos.
func (v APIPhotos) Search(params PhotosSearchParams) (*PhotosSearchResponse, error) {
	r, err := v.API.Request("photos.search", params)
	if err != nil {
		return nil, err
	}

	var resp PhotosSearchResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PhotosSaveParams are params for APIPhotos.Save
type PhotosSaveParams struct {
	// ID of the album to save photos to.
	AlbumID int `url:"album_id,omitempty"`
	// ID of the community to save photos to.
	GroupID int `url:"group_id,omitempty"`
	// Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
	Server int `url:"server,omitempty"`
	// Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
	PhotosList string `url:"photos_list,omitempty"`
	// Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
	Hash string `url:"hash,omitempty"`
	// Geographical latitude, in degrees (from '-90' to '90').
	Latitude float32 `url:"latitude,omitempty"`
	// Geographical longitude, in degrees (from '-180' to '180').
	Longitude float32 `url:"longitude,omitempty"`
	// Text describing the photo. 2048 digits max.
	Caption string `url:"caption,omitempty"`
}

// PhotosSaveResponse is response for APIPhotos.Save
//easyjson:json
type PhotosSaveResponse []Photo

// Save Saves photos after successful uploading.
func (v APIPhotos) Save(params PhotosSaveParams) (PhotosSaveResponse, error) {
	r, err := v.API.Request("photos.save", params)
	if err != nil {
		return nil, err
	}

	var resp PhotosSaveResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// PhotosCopyParams are params for APIPhotos.Copy
type PhotosCopyParams struct {
	// photo's owner ID
	OwnerID int `url:"owner_id"`
	// photo ID
	PhotoID int `url:"photo_id"`
	// for private photos
	AccessKey string `url:"access_key,omitempty"`
}

// PhotosCopyResponse is response for APIPhotos.Copy
// Photo ID
type PhotosCopyResponse int

// Copy Allows to copy a photo to the "Saved photos" album
func (v APIPhotos) Copy(params PhotosCopyParams) (PhotosCopyResponse, error) {
	r, err := v.API.Request("photos.copy", params)
	if err != nil {
		return 0, err
	}

	var resp PhotosCopyResponse

	var cnv int
	cnv, err = strconv.Atoi(string(r))
	resp = PhotosCopyResponse(cnv)

	if err != nil {
		return 0, err
	}
	return resp, nil
}

// PhotosEditParams are params for APIPhotos.Edit
type PhotosEditParams struct {
	// ID of the user or community that owns the photo.
	OwnerID int `url:"owner_id,omitempty"`
	// Photo ID.
	PhotoID int `url:"photo_id"`
	// New caption for the photo. If this parameter is not set, it is considered to be equal to an empty string.
	Caption      string  `url:"caption,omitempty"`
	Latitude     float32 `url:"latitude,omitempty"`
	Longitude    float32 `url:"longitude,omitempty"`
	PlaceStr     string  `url:"place_str,omitempty"`
	FoursquareID string  `url:"foursquare_id,omitempty"`
	DeletePlace  bool    `url:"delete_place,omitempty"`
}

// Edit Edits the caption of a photo.
func (v APIPhotos) Edit(params PhotosEditParams) (bool, error) {
	r, err := v.API.Request("photos.edit", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// PhotosMoveParams are params for APIPhotos.Move
type PhotosMoveParams struct {
	// ID of the user or community that owns the photo.
	OwnerID int `url:"owner_id,omitempty"`
	// ID of the album to which the photo will be moved.
	TargetAlbumID int `url:"target_album_id"`
	// Photo ID.
	PhotoID int `url:"photo_id"`
}

// Move Moves a photo from one album to another.
func (v APIPhotos) Move(params PhotosMoveParams) (bool, error) {
	r, err := v.API.Request("photos.move", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// PhotosMakeCoverParams are params for APIPhotos.MakeCover
type PhotosMakeCoverParams struct {
	// ID of the user or community that owns the photo.
	OwnerID int `url:"owner_id,omitempty"`
	// Photo ID.
	PhotoID int `url:"photo_id"`
	// Album ID.
	AlbumID int `url:"album_id,omitempty"`
}

// MakeCover Makes a photo into an album cover.
func (v APIPhotos) MakeCover(params PhotosMakeCoverParams) (bool, error) {
	r, err := v.API.Request("photos.makeCover", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// PhotosReorderAlbumsParams are params for APIPhotos.ReorderAlbums
type PhotosReorderAlbumsParams struct {
	// ID of the user or community that owns the album.
	OwnerID int `url:"owner_id,omitempty"`
	// Album ID.
	AlbumID int `url:"album_id"`
	// ID of the album before which the album in question shall be placed.
	Before int `url:"before,omitempty"`
	// ID of the album after which the album in question shall be placed.
	After int `url:"after,omitempty"`
}

// ReorderAlbums Reorders the album in the list of user albums.
func (v APIPhotos) ReorderAlbums(params PhotosReorderAlbumsParams) (bool, error) {
	r, err := v.API.Request("photos.reorderAlbums", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// PhotosReorderPhotosParams are params for APIPhotos.ReorderPhotos
type PhotosReorderPhotosParams struct {
	// ID of the user or community that owns the photo.
	OwnerID int `url:"owner_id,omitempty"`
	// Photo ID.
	PhotoID int `url:"photo_id"`
	// ID of the photo before which the photo in question shall be placed.
	Before int `url:"before,omitempty"`
	// ID of the photo after which the photo in question shall be placed.
	After int `url:"after,omitempty"`
}

// ReorderPhotos Reorders the photo in the list of photos of the user album.
func (v APIPhotos) ReorderPhotos(params PhotosReorderPhotosParams) (bool, error) {
	r, err := v.API.Request("photos.reorderPhotos", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// PhotosGetAllParams are params for APIPhotos.GetAll
type PhotosGetAllParams struct {
	// ID of a user or community that owns the photos. Use a negative value to designate a community ID.
	OwnerID int `url:"owner_id,omitempty"`
	// '1' — to return detailed information about photos
	Extended bool `url:"extended,omitempty"`
	// Offset needed to return a specific subset of photos. By default, '0'.
	Offset int `url:"offset,omitempty"`
	// Number of photos to return.
	Count int `url:"count,omitempty"`
	// '1' – to return image sizes in [vk.com/dev/photo_sizes|special format].
	PhotoSizes bool `url:"photo_sizes,omitempty"`
	// '1' – to return photos only from standard albums, '0' – to return all photos including those in service albums, e.g., 'My wall photos' (default)
	NoServiceAlbums bool `url:"no_service_albums,omitempty"`
	// '1' – to show information about photos being hidden from the block above the wall.
	NeedHidden bool `url:"need_hidden,omitempty"`
	// '1' – not to return photos being hidden from the block above the wall. Works only with owner_id>0, no_service_albums is ignored.
	SkipHidden bool `url:"skip_hidden,omitempty"`
}

// PhotosGetAllResponse is response for APIPhotos.GetAll
// Either PhotosGetAllResponseNormal or PhotosGetAllResponseExtended, depending on Extended flag
type PhotosGetAllResponse interface {
	isPhotosGetAll()
}

// PhotosGetAllResponseNormal is non-extended version of PhotosGetAllResponse
//easyjson:json
type PhotosGetAllResponseNormal struct {
	// Total number
	Count int     `json:"count,omitempty"`
	Items []Photo `json:"items,omitempty"`
	// Information whether next page is presented
	More BoolInt `json:"more,omitempty"`
}

func (PhotosGetAllResponseNormal) isPhotosGetAll() {}

// PhotosGetAllResponseExtended is extended version of PhotosGetAllResponse
//easyjson:json
type PhotosGetAllResponseExtended struct {
	// Total number
	Count int     `json:"count,omitempty"`
	Items []Photo `json:"items,omitempty"`
	// Information whether next page is presented
	More BoolInt `json:"more,omitempty"`
}

func (PhotosGetAllResponseExtended) isPhotosGetAll() {}

// GetAll Returns a list of photos belonging to a user or community, in reverse chronological order.
func (v APIPhotos) GetAll(params PhotosGetAllParams) (PhotosGetAllResponse, error) {
	r, err := v.API.Request("photos.getAll", params)
	if err != nil {
		return nil, err
	}

	var resp PhotosGetAllResponse
	if params.Extended {
		var tmp PhotosGetAllResponseExtended
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	} else {
		var tmp PhotosGetAllResponseNormal
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// PhotosGetUserPhotosParams are params for APIPhotos.GetUserPhotos
type PhotosGetUserPhotosParams struct {
	// User ID.
	UserID int `url:"user_id,omitempty"`
	// Offset needed to return a specific subset of photos. By default, '0'.
	Offset int `url:"offset,omitempty"`
	// Number of photos to return. Maximum value is 1000.
	Count int `url:"count,omitempty"`
	// '1' — to return an additional 'likes' field, '0' — (default)
	Extended bool `url:"extended,omitempty"`
	// Sort order: '1' — by date the tag was added in ascending order, '0' — by date the tag was added in descending order
	Sort string `url:"sort,omitempty"`
}

// PhotosGetUserPhotosResponse is response for APIPhotos.GetUserPhotos
//easyjson:json
type PhotosGetUserPhotosResponse struct {
	// Total number
	Count int     `json:"count,omitempty"`
	Items []Photo `json:"items,omitempty"`
}

// GetUserPhotos Returns a list of photos in which a user is tagged.
func (v APIPhotos) GetUserPhotos(params PhotosGetUserPhotosParams) (*PhotosGetUserPhotosResponse, error) {
	r, err := v.API.Request("photos.getUserPhotos", params)
	if err != nil {
		return nil, err
	}

	var resp PhotosGetUserPhotosResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PhotosDeleteAlbumParams are params for APIPhotos.DeleteAlbum
type PhotosDeleteAlbumParams struct {
	// Album ID.
	AlbumID int `url:"album_id"`
	// ID of the community that owns the album.
	GroupID int `url:"group_id,omitempty"`
}

// DeleteAlbum Deletes a photo album belonging to the current user.
func (v APIPhotos) DeleteAlbum(params PhotosDeleteAlbumParams) (bool, error) {
	r, err := v.API.Request("photos.deleteAlbum", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// PhotosDeleteParams are params for APIPhotos.Delete
type PhotosDeleteParams struct {
	// ID of the user or community that owns the photo.
	OwnerID int `url:"owner_id,omitempty"`
	// Photo ID.
	PhotoID int `url:"photo_id"`
}

// Delete Deletes a photo.
func (v APIPhotos) Delete(params PhotosDeleteParams) (bool, error) {
	r, err := v.API.Request("photos.delete", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// PhotosRestoreParams are params for APIPhotos.Restore
type PhotosRestoreParams struct {
	// ID of the user or community that owns the photo.
	OwnerID int `url:"owner_id,omitempty"`
	// Photo ID.
	PhotoID int `url:"photo_id"`
}

// Restore Restores a deleted photo.
func (v APIPhotos) Restore(params PhotosRestoreParams) (bool, error) {
	r, err := v.API.Request("photos.restore", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// PhotosConfirmTagParams are params for APIPhotos.ConfirmTag
type PhotosConfirmTagParams struct {
	// ID of the user or community that owns the photo.
	OwnerID int `url:"owner_id,omitempty"`
	// Photo ID.
	PhotoID string `url:"photo_id"`
	// Tag ID.
	TagID int `url:"tag_id"`
}

// ConfirmTag Confirms a tag on a photo.
func (v APIPhotos) ConfirmTag(params PhotosConfirmTagParams) (bool, error) {
	r, err := v.API.Request("photos.confirmTag", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// PhotosGetCommentsParams are params for APIPhotos.GetComments
type PhotosGetCommentsParams struct {
	// ID of the user or community that owns the photo.
	OwnerID int `url:"owner_id,omitempty"`
	// Photo ID.
	PhotoID int `url:"photo_id"`
	// '1' — to return an additional 'likes' field, '0' — (default)
	NeedLikes      bool `url:"need_likes,omitempty"`
	StartCommentID int  `url:"start_comment_id,omitempty"`
	// Offset needed to return a specific subset of comments. By default, '0'.
	Offset int `url:"offset,omitempty"`
	// Number of comments to return.
	Count int `url:"count,omitempty"`
	// Sort order: 'asc' — old first, 'desc' — new first
	Sort      string         `url:"sort,omitempty"`
	AccessKey string         `url:"access_key,omitempty"`
	Extended  bool           `url:"extended,omitempty"`
	Fields    CSVStringSlice `url:"fields,omitempty"`
}

// PhotosGetCommentsResponse is response for APIPhotos.GetComments
// Either PhotosGetCommentsResponseNormal or PhotosGetCommentsResponseExtended, depending on Extended flag
type PhotosGetCommentsResponse interface {
	isPhotosGetComments()
}

// PhotosGetCommentsResponseNormal is non-extended version of PhotosGetCommentsResponse
//easyjson:json
type PhotosGetCommentsResponseNormal struct {
	// Total number
	Count int `json:"count,omitempty"`
	// Real offset of the comments
	RealOffset int       `json:"real_offset,omitempty"`
	Items      []Comment `json:"items,omitempty"`
}

func (PhotosGetCommentsResponseNormal) isPhotosGetComments() {}

// PhotosGetCommentsResponseExtended is extended version of PhotosGetCommentsResponse
//easyjson:json
type PhotosGetCommentsResponseExtended struct {
	// Total number
	Count int `json:"count,omitempty"`
	// Real offset of the comments
	RealOffset int       `json:"real_offset,omitempty"`
	Items      []Comment `json:"items,omitempty"`
	Profiles   []User    `json:"profiles,omitempty"`
	Groups     []Group   `json:"groups,omitempty"`
}

func (PhotosGetCommentsResponseExtended) isPhotosGetComments() {}

// GetComments Returns a list of comments on a photo.
func (v APIPhotos) GetComments(params PhotosGetCommentsParams) (PhotosGetCommentsResponse, error) {
	r, err := v.API.Request("photos.getComments", params)
	if err != nil {
		return nil, err
	}

	var resp PhotosGetCommentsResponse
	if params.Extended {
		var tmp PhotosGetCommentsResponseExtended
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	} else {
		var tmp PhotosGetCommentsResponseNormal
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// PhotosGetAllCommentsParams are params for APIPhotos.GetAllComments
type PhotosGetAllCommentsParams struct {
	// ID of the user or community that owns the album(s).
	OwnerID int `url:"owner_id,omitempty"`
	// Album ID. If the parameter is not set, comments on all of the user's albums will be returned.
	AlbumID int `url:"album_id,omitempty"`
	// '1' — to return an additional 'likes' field, '0' — (default)
	NeedLikes bool `url:"need_likes,omitempty"`
	// Offset needed to return a specific subset of comments. By default, '0'.
	Offset int `url:"offset,omitempty"`
	// Number of comments to return. By default, '20'. Maximum value, '100'.
	Count int `url:"count,omitempty"`
}

// PhotosGetAllCommentsResponse is response for APIPhotos.GetAllComments
//easyjson:json
type PhotosGetAllCommentsResponse struct {
	// Total number
	Count int       `json:"count,omitempty"`
	Items []Comment `json:"items,omitempty"`
}

// GetAllComments Returns a list of comments on a specific photo album or all albums of the user sorted in reverse chronological order.
func (v APIPhotos) GetAllComments(params PhotosGetAllCommentsParams) (*PhotosGetAllCommentsResponse, error) {
	r, err := v.API.Request("photos.getAllComments", params)
	if err != nil {
		return nil, err
	}

	var resp PhotosGetAllCommentsResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PhotosCreateCommentParams are params for APIPhotos.CreateComment
type PhotosCreateCommentParams struct {
	// ID of the user or community that owns the photo.
	OwnerID int `url:"owner_id,omitempty"`
	// Photo ID.
	PhotoID int `url:"photo_id"`
	// Comment text.
	Message string `url:"message,omitempty"`
	// (Required if 'message' is not set.) List of objects attached to the post, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, '<owner_id>' — Media attachment owner ID. '<media_id>' — Media attachment ID. Example: "photo100172_166443618,photo66748_265827614"
	Attachments CSVStringSlice `url:"attachments,omitempty"`
	// '1' — to post a comment from the community
	FromGroup bool `url:"from_group,omitempty"`
	//
	ReplyToComment int    `url:"reply_to_comment,omitempty"`
	StickerID      int    `url:"sticker_id,omitempty"`
	AccessKey      string `url:"access_key,omitempty"`
	GUID           string `url:"guid,omitempty"`
}

// PhotosCreateCommentResponse is response for APIPhotos.CreateComment
// Created comment ID
type PhotosCreateCommentResponse int

// CreateComment Adds a new comment on the photo.
func (v APIPhotos) CreateComment(params PhotosCreateCommentParams) (PhotosCreateCommentResponse, error) {
	r, err := v.API.Request("photos.createComment", params)
	if err != nil {
		return 0, err
	}

	var resp PhotosCreateCommentResponse

	var cnv int
	cnv, err = strconv.Atoi(string(r))
	resp = PhotosCreateCommentResponse(cnv)

	if err != nil {
		return 0, err
	}
	return resp, nil
}

// PhotosDeleteCommentParams are params for APIPhotos.DeleteComment
type PhotosDeleteCommentParams struct {
	// ID of the user or community that owns the photo.
	OwnerID int `url:"owner_id,omitempty"`
	// Comment ID.
	CommentID int `url:"comment_id"`
}

// DeleteComment Deletes a comment on the photo.
func (v APIPhotos) DeleteComment(params PhotosDeleteCommentParams) (bool, error) {
	r, err := v.API.Request("photos.deleteComment", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// PhotosRestoreCommentParams are params for APIPhotos.RestoreComment
type PhotosRestoreCommentParams struct {
	// ID of the user or community that owns the photo.
	OwnerID int `url:"owner_id,omitempty"`
	// ID of the deleted comment.
	CommentID int `url:"comment_id"`
}

// RestoreComment Restores a deleted comment on a photo.
func (v APIPhotos) RestoreComment(params PhotosRestoreCommentParams) (bool, error) {
	r, err := v.API.Request("photos.restoreComment", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// PhotosEditCommentParams are params for APIPhotos.EditComment
type PhotosEditCommentParams struct {
	// ID of the user or community that owns the photo.
	OwnerID int `url:"owner_id,omitempty"`
	// Comment ID.
	CommentID int `url:"comment_id"`
	// New text of the comment.
	Message string `url:"message,omitempty"`
	// (Required if 'message' is not set.) List of objects attached to the post, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, '<owner_id>' — Media attachment owner ID. '<media_id>' — Media attachment ID. Example: "photo100172_166443618,photo66748_265827614"
	Attachments CSVStringSlice `url:"attachments,omitempty"`
}

// EditComment Edits a comment on a photo.
func (v APIPhotos) EditComment(params PhotosEditCommentParams) (bool, error) {
	r, err := v.API.Request("photos.editComment", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// PhotosGetTagsParams are params for APIPhotos.GetTags
type PhotosGetTagsParams struct {
	// ID of the user or community that owns the photo.
	OwnerID int `url:"owner_id,omitempty"`
	// Photo ID.
	PhotoID   int    `url:"photo_id"`
	AccessKey string `url:"access_key,omitempty"`
}

// PhotosGetTagsResponse is response for APIPhotos.GetTags
//easyjson:json
type PhotosGetTagsResponse []genTODOType /* objects.json#/definitions/photos_photo_tag */
// GetTags Returns a list of tags on a photo.
func (v APIPhotos) GetTags(params PhotosGetTagsParams) (PhotosGetTagsResponse, error) {
	r, err := v.API.Request("photos.getTags", params)
	if err != nil {
		return nil, err
	}

	var resp PhotosGetTagsResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// PhotosPutTagParams are params for APIPhotos.PutTag
type PhotosPutTagParams struct {
	// ID of the user or community that owns the photo.
	OwnerID int `url:"owner_id,omitempty"`
	// Photo ID.
	PhotoID int `url:"photo_id"`
	// ID of the user to be tagged.
	UserID int `url:"user_id"`
	// Upper left-corner coordinate of the tagged area (as a percentage of the photo's width).
	X float32 `url:"x,omitempty"`
	// Upper left-corner coordinate of the tagged area (as a percentage of the photo's height).
	Y float32 `url:"y,omitempty"`
	// Lower right-corner coordinate of the tagged area (as a percentage of the photo's width).
	X2 float32 `url:"x2,omitempty"`
	// Lower right-corner coordinate of the tagged area (as a percentage of the photo's height).
	Y2 float32 `url:"y2,omitempty"`
}

// PhotosPutTagResponse is response for APIPhotos.PutTag
// Created tag ID
type PhotosPutTagResponse int

// PutTag Adds a tag on the photo.
func (v APIPhotos) PutTag(params PhotosPutTagParams) (PhotosPutTagResponse, error) {
	r, err := v.API.Request("photos.putTag", params)
	if err != nil {
		return 0, err
	}

	var resp PhotosPutTagResponse

	var cnv int
	cnv, err = strconv.Atoi(string(r))
	resp = PhotosPutTagResponse(cnv)

	if err != nil {
		return 0, err
	}
	return resp, nil
}

// PhotosRemoveTagParams are params for APIPhotos.RemoveTag
type PhotosRemoveTagParams struct {
	// ID of the user or community that owns the photo.
	OwnerID int `url:"owner_id,omitempty"`
	// Photo ID.
	PhotoID int `url:"photo_id"`
	// Tag ID.
	TagID int `url:"tag_id"`
}

// RemoveTag Removes a tag from a photo.
func (v APIPhotos) RemoveTag(params PhotosRemoveTagParams) (bool, error) {
	r, err := v.API.Request("photos.removeTag", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// PhotosGetNewTagsParams are params for APIPhotos.GetNewTags
type PhotosGetNewTagsParams struct {
	// Offset needed to return a specific subset of photos.
	Offset int `url:"offset,omitempty"`
	// Number of photos to return.
	Count int `url:"count,omitempty"`
}

// PhotosGetNewTagsResponse is response for APIPhotos.GetNewTags
//easyjson:json
type PhotosGetNewTagsResponse struct {
	// Total number
	Count int     `json:"count,omitempty"`
	Items []Photo `json:"items,omitempty"`
}

// GetNewTags Returns a list of photos with tags that have not been viewed.
func (v APIPhotos) GetNewTags(params PhotosGetNewTagsParams) (*PhotosGetNewTagsResponse, error) {
	r, err := v.API.Request("photos.getNewTags", params)
	if err != nil {
		return nil, err
	}

	var resp PhotosGetNewTagsResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
