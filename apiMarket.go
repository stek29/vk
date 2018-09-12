package vkCallbackApi

import (
	"encoding/json"
	"strconv"
)

// APIMarket implements VK API namespace `market`
type APIMarket struct {
	API *API
}

// MarketGetParams are params for APIMarket.Get
type MarketGetParams struct {
	// ID of an item owner community, "Note that community id in the 'owner_id' parameter should be negative number. For example 'owner_id'=-1 matches the [vk.com/apiclub|VK API] community "
	OwnerID int `url:"owner_id"`
	// Number of items to return.
	Count int `url:"count,omitempty"`
	// Offset needed to return a specific subset of results.
	Offset int `url:"offset,omitempty"`
	// '1' – method will return additional fields: 'likes, can_comment, car_repost, photos'. These parameters are not returned by default.
	Extended bool `url:"extended,omitempty"`
}

// MarketGetResponse is response for APIMarket.Get
// Either MarketGetResponseNormal or MarketGetResponseExtended, depending on Extended flag
type MarketGetResponse interface {
	isMarketGet()
}

// MarketGetResponseNormal is non-extended version of MarketGetResponse
//easyjson:json
type MarketGetResponseNormal struct {
	// Total number
	Count int          `json:"count,omitempty"`
	Items []MarketItem `json:"items,omitempty"`
}

func (MarketGetResponseNormal) isMarketGet() {}

// MarketGetResponseExtended is extended version of MarketGetResponse
//easyjson:json
type MarketGetResponseExtended struct {
	// Total number
	Count int          `json:"count,omitempty"`
	Items []MarketItem `json:"items,omitempty"`
}

func (MarketGetResponseExtended) isMarketGet() {}

// Get Returns items list for a community.
func (v APIMarket) Get(params MarketGetParams) (MarketGetResponse, error) {
	r, err := v.API.Request("market.get", params)
	if err != nil {
		return nil, err
	}

	var resp MarketGetResponse
	if params.Extended {
		var tmp MarketGetResponseExtended
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	} else {
		var tmp MarketGetResponseNormal
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// MarketGetByIDParams are params for APIMarket.GetByID
type MarketGetByIDParams struct {
	// Comma-separated ids list: {user id}_{item id}. If an item belongs to a community -{community id} is used. " 'Videos' value example: , '-4363_136089719,13245770_137352259'"
	ItemIDs CSVStringSlice `url:"item_ids"`
	// '1' – to return additional fields: 'likes, can_comment, car_repost, photos'. By default: '0'.
	Extended bool `url:"extended,omitempty"`
}

// MarketGetByIDResponse is response for APIMarket.GetByID
// Either MarketGetByIDResponseNormal or MarketGetByIDResponseExtended, depending on Extended flag
type MarketGetByIDResponse interface {
	isMarketGetByID()
}

// MarketGetByIDResponseNormal is non-extended version of MarketGetByIDResponse
//easyjson:json
type MarketGetByIDResponseNormal struct {
	// Total number
	Count int          `json:"count,omitempty"`
	Items []MarketItem `json:"items,omitempty"`
}

func (MarketGetByIDResponseNormal) isMarketGetByID() {}

// MarketGetByIDResponseExtended is extended version of MarketGetByIDResponse
//easyjson:json
type MarketGetByIDResponseExtended struct {
	// Total number
	Count int          `json:"count,omitempty"`
	Items []MarketItem `json:"items,omitempty"`
}

func (MarketGetByIDResponseExtended) isMarketGetByID() {}

// GetByID Returns information about market items by their ids.
func (v APIMarket) GetByID(params MarketGetByIDParams) (MarketGetByIDResponse, error) {
	r, err := v.API.Request("market.getById", params)
	if err != nil {
		return nil, err
	}

	var resp MarketGetByIDResponse
	if params.Extended {
		var tmp MarketGetByIDResponseExtended
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	} else {
		var tmp MarketGetByIDResponseNormal
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// MarketSearchParams are params for APIMarket.Search
type MarketSearchParams struct {
	// ID of an items owner community.
	OwnerID int `url:"owner_id"`
	// Search query, for example "pink slippers".
	Q string `url:"q,omitempty"`
	// Minimum item price value.
	PriceFrom int `url:"price_from,omitempty"`
	// Maximum item price value.
	PriceTo int `url:"price_to,omitempty"`
	// Comma-separated tag IDs list.
	Tags CSVIntSlice `url:"tags,omitempty"`
	// '0' — do not use reverse order, '1' — use reverse order
	Rev int `url:"rev,omitempty"`
	// Offset needed to return a specific subset of results.
	Offset int `url:"offset,omitempty"`
	// Number of items to return.
	Count int `url:"count,omitempty"`
	// '1' – to return additional fields: 'likes, can_comment, car_repost, photos'. By default: '0'.
	Extended bool `url:"extended,omitempty"`
}

// MarketSearchResponse is response for APIMarket.Search
// Either MarketSearchResponseNormal or MarketSearchResponseExtended, depending on Extended flag
type MarketSearchResponse interface {
	isMarketSearch()
}

// MarketSearchResponseNormal is non-extended version of MarketSearchResponse
//easyjson:json
type MarketSearchResponseNormal struct {
	// Total number
	Count int          `json:"count,omitempty"`
	Items []MarketItem `json:"items,omitempty"`
}

func (MarketSearchResponseNormal) isMarketSearch() {}

// MarketSearchResponseExtended is extended version of MarketSearchResponse
//easyjson:json
type MarketSearchResponseExtended struct {
	// Total number
	Count int          `json:"count,omitempty"`
	Items []MarketItem `json:"items,omitempty"`
}

func (MarketSearchResponseExtended) isMarketSearch() {}

// Search Searches market items in a community's catalog
func (v APIMarket) Search(params MarketSearchParams) (MarketSearchResponse, error) {
	r, err := v.API.Request("market.search", params)
	if err != nil {
		return nil, err
	}

	var resp MarketSearchResponse
	if params.Extended {
		var tmp MarketSearchResponseExtended
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	} else {
		var tmp MarketSearchResponseNormal
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// MarketGetAlbumsParams are params for APIMarket.GetAlbums
type MarketGetAlbumsParams struct {
	// ID of an items owner community.
	OwnerID int `url:"owner_id"`
	// Offset needed to return a specific subset of results.
	Offset int `url:"offset,omitempty"`
	// Number of items to return.
	Count int `url:"count,omitempty"`
}

// MarketGetAlbumsResponse is response for APIMarket.GetAlbums
//easyjson:json
type MarketGetAlbumsResponse struct {
	// Total number
	Count int           `json:"count,omitempty"`
	Items []MarketAlbum `json:"items,omitempty"`
}

// GetAlbums Returns community's collections list.
func (v APIMarket) GetAlbums(params MarketGetAlbumsParams) (*MarketGetAlbumsResponse, error) {
	r, err := v.API.Request("market.getAlbums", params)
	if err != nil {
		return nil, err
	}

	var resp MarketGetAlbumsResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// MarketGetAlbumByIDParams are params for APIMarket.GetAlbumByID
type MarketGetAlbumByIDParams struct {
	// identifier of an album owner community, "Note that community id in the 'owner_id' parameter should be negative number. For example 'owner_id'=-1 matches the [vk.com/apiclub|VK API] community "
	OwnerID int `url:"owner_id"`
	// collections identifiers to obtain data from
	AlbumIDs CSVIntSlice `url:"album_ids"`
}

// MarketGetAlbumByIDResponse is response for APIMarket.GetAlbumByID
//easyjson:json
type MarketGetAlbumByIDResponse struct {
	// Total number
	Count int           `json:"count,omitempty"`
	Items []MarketAlbum `json:"items,omitempty"`
}

// GetAlbumByID Returns items album's data
func (v APIMarket) GetAlbumByID(params MarketGetAlbumByIDParams) (*MarketGetAlbumByIDResponse, error) {
	r, err := v.API.Request("market.getAlbumById", params)
	if err != nil {
		return nil, err
	}

	var resp MarketGetAlbumByIDResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// MarketCreateCommentParams are params for APIMarket.CreateComment
type MarketCreateCommentParams struct {
	// ID of an item owner community.
	OwnerID int `url:"owner_id"`
	// Item ID.
	ItemID int `url:"item_id"`
	// Comment text (required if 'attachments' parameter is not specified)
	Message string `url:"message,omitempty"`
	// Comma-separated list of objects attached to a comment. The field is submitted the following way: , "'<owner_id>_<media_id>,<owner_id>_<media_id>'", , '' - media attachment type: "'photo' - photo, 'video' - video, 'audio' - audio, 'doc' - document", , '<owner_id>' - media owner id, '<media_id>' - media attachment id, , For example: "photo100172_166443618,photo66748_265827614",
	Attachments CSVStringSlice `url:"attachments,omitempty"`
	// '1' - comment will be published on behalf of a community, '0' - on behalf of a user (by default).
	FromGroup bool `url:"from_group,omitempty"`
	// ID of a comment to reply with current comment to.
	ReplyToComment int `url:"reply_to_comment,omitempty"`
	// Sticker ID.
	StickerID int `url:"sticker_id,omitempty"`
	// Random value to avoid resending one comment.
	GUID string `url:"guid,omitempty"`
}

// MarketCreateCommentResponse is response for APIMarket.CreateComment
// Comment ID
type MarketCreateCommentResponse int

// CreateComment Creates a new comment for an item.
func (v APIMarket) CreateComment(params MarketCreateCommentParams) (MarketCreateCommentResponse, error) {
	r, err := v.API.Request("market.createComment", params)
	if err != nil {
		return 0, err
	}

	var resp MarketCreateCommentResponse

	var cnv int
	cnv, err = strconv.Atoi(string(r))
	resp = MarketCreateCommentResponse(cnv)

	if err != nil {
		return 0, err
	}
	return resp, nil
}

// MarketGetCommentsParams are params for APIMarket.GetComments
type MarketGetCommentsParams struct {
	// ID of an item owner community
	OwnerID int `url:"owner_id"`
	// Item ID.
	ItemID int `url:"item_id"`
	// '1' — to return likes info.
	NeedLikes bool `url:"need_likes,omitempty"`
	// ID of a comment to start a list from (details below).
	StartCommentID int `url:"start_comment_id,omitempty"`
	// Number of results to return.
	Count int `url:"count,omitempty"`
	// Sort order ('asc' — from old to new, 'desc' — from new to old)
	Sort string `url:"sort,omitempty"`
	// '1' — comments will be returned as numbered objects, in addition lists of 'profiles' and 'groups' objects will be returned.
	Extended bool `url:"extended,omitempty"`
	// List of additional profile fields to return. See the [vk.com/dev/fields|details]
	Fields CSVStringSlice `url:"fields,omitempty"`
}

// MarketGetCommentsResponse is response for APIMarket.GetComments
//easyjson:json
type MarketGetCommentsResponse struct {
	// Total number
	Count int       `json:"count,omitempty"`
	Items []Comment `json:"items,omitempty"`
}

// GetComments Returns comments list for an item.
func (v APIMarket) GetComments(params MarketGetCommentsParams) (*MarketGetCommentsResponse, error) {
	r, err := v.API.Request("market.getComments", params)
	if err != nil {
		return nil, err
	}

	var resp MarketGetCommentsResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// MarketDeleteCommentParams are params for APIMarket.DeleteComment
type MarketDeleteCommentParams struct {
	// identifier of an item owner community, "Note that community id in the 'owner_id' parameter should be negative number. For example 'owner_id'=-1 matches the [vk.com/apiclub|VK API] community "
	OwnerID int `url:"owner_id"`
	// comment id
	CommentID int `url:"comment_id"`
}

// DeleteComment Deletes an item's comment
func (v APIMarket) DeleteComment(params MarketDeleteCommentParams) (bool, error) {
	r, err := v.API.Request("market.deleteComment", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// MarketRestoreCommentParams are params for APIMarket.RestoreComment
type MarketRestoreCommentParams struct {
	// identifier of an item owner community, "Note that community id in the 'owner_id' parameter should be negative number. For example 'owner_id'=-1 matches the [vk.com/apiclub|VK API] community "
	OwnerID int `url:"owner_id"`
	// deleted comment id
	CommentID int `url:"comment_id"`
}

// RestoreComment Restores a recently deleted comment
func (v APIMarket) RestoreComment(params MarketRestoreCommentParams) (bool, error) {
	r, err := v.API.Request("market.restoreComment", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// MarketEditCommentParams are params for APIMarket.EditComment
type MarketEditCommentParams struct {
	// ID of an item owner community.
	OwnerID int `url:"owner_id"`
	// Comment ID.
	CommentID int `url:"comment_id"`
	// New comment text (required if 'attachments' are not specified), , 2048 symbols maximum.
	Message string `url:"message,omitempty"`
	// Comma-separated list of objects attached to a comment. The field is submitted the following way: , "'<owner_id>_<media_id>,<owner_id>_<media_id>'", , '' - media attachment type: "'photo' - photo, 'video' - video, 'audio' - audio, 'doc' - document", , '<owner_id>' - media owner id, '<media_id>' - media attachment id, , For example: "photo100172_166443618,photo66748_265827614",
	Attachments CSVStringSlice `url:"attachments,omitempty"`
}

// EditComment Chages item comment's text
func (v APIMarket) EditComment(params MarketEditCommentParams) (bool, error) {
	r, err := v.API.Request("market.editComment", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// MarketReportCommentParams are params for APIMarket.ReportComment
type MarketReportCommentParams struct {
	// ID of an item owner community.
	OwnerID int `url:"owner_id"`
	// Comment ID.
	CommentID int `url:"comment_id"`
	// Complaint reason. Possible values: *'0' — spam,, *'1' — child porn,, *'2' — extremism,, *'3' — violence,, *'4' — drugs propaganda,, *'5' — adult materials,, *'6' — insult.
	Reason int `url:"reason"`
}

// ReportComment Sends a complaint to the item's comment.
func (v APIMarket) ReportComment(params MarketReportCommentParams) (bool, error) {
	r, err := v.API.Request("market.reportComment", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// MarketGetCategoriesParams are params for APIMarket.GetCategories
type MarketGetCategoriesParams struct {
	// Number of results to return.
	Count int `url:"count,omitempty"`
	// Offset needed to return a specific subset of results.
	Offset int `url:"offset,omitempty"`
}

// MarketGetCategoriesResponse is response for APIMarket.GetCategories
//easyjson:json
type MarketGetCategoriesResponse struct {
	// Total number
	Count int              `json:"count,omitempty"`
	Items []MarketCategory `json:"items,omitempty"`
}

// GetCategories Returns a list of market categories.
func (v APIMarket) GetCategories(params MarketGetCategoriesParams) (*MarketGetCategoriesResponse, error) {
	r, err := v.API.Request("market.getCategories", params)
	if err != nil {
		return nil, err
	}

	var resp MarketGetCategoriesResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// MarketReportParams are params for APIMarket.Report
type MarketReportParams struct {
	// ID of an item owner community.
	OwnerID int `url:"owner_id"`
	// Item ID.
	ItemID int `url:"item_id"`
	// Complaint reason. Possible values: *'0' — spam,, *'1' — child porn,, *'2' — extremism,, *'3' — violence,, *'4' — drugs propaganda,, *'5' — adult materials,, *'6' — insult.
	Reason int `url:"reason"`
}

// Report Sends a complaint to the item.
func (v APIMarket) Report(params MarketReportParams) (bool, error) {
	r, err := v.API.Request("market.report", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// MarketAddParams are params for APIMarket.Add
type MarketAddParams struct {
	// ID of an item owner community.
	OwnerID int `url:"owner_id"`
	// Item name.
	Name string `url:"name"`
	// Item description.
	Description string `url:"description"`
	// Item category ID.
	CategoryID int `url:"category_id"`
	// Item price.
	Price float32 `url:"price"`
	// Item status ('1' — deleted, '0' — not deleted).
	Deleted bool `url:"deleted,omitempty"`
	// Cover photo ID.
	MainPhotoID int `url:"main_photo_id"`
	// IDs of additional photos.
	PhotoIDs CSVIntSlice `url:"photo_ids,omitempty"`
}

// MarketAddResponse is response for APIMarket.Add
//easyjson:json
type MarketAddResponse struct {
	// Item ID
	MarketItemID int `json:"market_item_id,omitempty"`
}

// Add Ads a new item to the market.
func (v APIMarket) Add(params MarketAddParams) (*MarketAddResponse, error) {
	r, err := v.API.Request("market.add", params)
	if err != nil {
		return nil, err
	}

	var resp MarketAddResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// MarketEditParams are params for APIMarket.Edit
type MarketEditParams struct {
	// ID of an item owner community.
	OwnerID int `url:"owner_id"`
	// Item ID.
	ItemID int `url:"item_id"`
	// Item name.
	Name string `url:"name"`
	// Item description.
	Description string `url:"description"`
	// Item category ID.
	CategoryID int `url:"category_id"`
	// Item price.
	Price float32 `url:"price"`
	// Item status ('1' — deleted, '0' — not deleted).
	Deleted bool `url:"deleted,omitempty"`
	// Cover photo ID.
	MainPhotoID int `url:"main_photo_id"`
	// IDs of additional photos.
	PhotoIDs CSVIntSlice `url:"photo_ids,omitempty"`
}

// Edit Edits an item.
func (v APIMarket) Edit(params MarketEditParams) (bool, error) {
	r, err := v.API.Request("market.edit", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// MarketDeleteParams are params for APIMarket.Delete
type MarketDeleteParams struct {
	// ID of an item owner community.
	OwnerID int `url:"owner_id"`
	// Item ID.
	ItemID int `url:"item_id"`
}

// Delete Deletes an item.
func (v APIMarket) Delete(params MarketDeleteParams) (bool, error) {
	r, err := v.API.Request("market.delete", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// MarketRestoreParams are params for APIMarket.Restore
type MarketRestoreParams struct {
	// ID of an item owner community.
	OwnerID int `url:"owner_id"`
	// Deleted item ID.
	ItemID int `url:"item_id"`
}

// Restore Restores recently deleted item
func (v APIMarket) Restore(params MarketRestoreParams) (bool, error) {
	r, err := v.API.Request("market.restore", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// MarketReorderItemsParams are params for APIMarket.ReorderItems
type MarketReorderItemsParams struct {
	// ID of an item owner community.
	OwnerID int `url:"owner_id"`
	// ID of a collection to reorder items in. Set 0 to reorder full items list.
	AlbumID int `url:"album_id,omitempty"`
	// Item ID.
	ItemID int `url:"item_id"`
	// ID of an item to place current item before it.
	Before int `url:"before,omitempty"`
	// ID of an item to place current item after it.
	After int `url:"after,omitempty"`
}

// ReorderItems Changes item place in a collection.
func (v APIMarket) ReorderItems(params MarketReorderItemsParams) (bool, error) {
	r, err := v.API.Request("market.reorderItems", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// MarketReorderAlbumsParams are params for APIMarket.ReorderAlbums
type MarketReorderAlbumsParams struct {
	// ID of an item owner community.
	OwnerID int `url:"owner_id"`
	// Collection ID.
	AlbumID int `url:"album_id"`
	// ID of a collection to place current collection before it.
	Before int `url:"before,omitempty"`
	// ID of a collection to place current collection after it.
	After int `url:"after,omitempty"`
}

// ReorderAlbums Reorders the collections list.
func (v APIMarket) ReorderAlbums(params MarketReorderAlbumsParams) (bool, error) {
	r, err := v.API.Request("market.reorderAlbums", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// MarketAddAlbumParams are params for APIMarket.AddAlbum
type MarketAddAlbumParams struct {
	// ID of an item owner community.
	OwnerID int `url:"owner_id"`
	// Collection title.
	Title string `url:"title"`
	// Cover photo ID.
	PhotoID int `url:"photo_id,omitempty"`
	// Set as main ('1' – set, '0' – no).
	MainAlbum bool `url:"main_album,omitempty"`
}

// MarketAddAlbumResponse is response for APIMarket.AddAlbum
//easyjson:json
type MarketAddAlbumResponse struct {
	// Album ID
	MarketAlbumID int `json:"market_album_id,omitempty"`
}

// AddAlbum Creates new collection of items
func (v APIMarket) AddAlbum(params MarketAddAlbumParams) (*MarketAddAlbumResponse, error) {
	r, err := v.API.Request("market.addAlbum", params)
	if err != nil {
		return nil, err
	}

	var resp MarketAddAlbumResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// MarketEditAlbumParams are params for APIMarket.EditAlbum
type MarketEditAlbumParams struct {
	// ID of an collection owner community.
	OwnerID int `url:"owner_id"`
	// Collection ID.
	AlbumID int `url:"album_id"`
	// Collection title.
	Title string `url:"title"`
	// Cover photo id
	PhotoID int `url:"photo_id,omitempty"`
	// Set as main ('1' – set, '0' – no).
	MainAlbum bool `url:"main_album,omitempty"`
}

// EditAlbum Edits a collection of items
func (v APIMarket) EditAlbum(params MarketEditAlbumParams) (bool, error) {
	r, err := v.API.Request("market.editAlbum", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// MarketDeleteAlbumParams are params for APIMarket.DeleteAlbum
type MarketDeleteAlbumParams struct {
	// ID of an collection owner community.
	OwnerID int `url:"owner_id"`
	// Collection ID.
	AlbumID int `url:"album_id"`
}

// DeleteAlbum Deletes a collection of items.
func (v APIMarket) DeleteAlbum(params MarketDeleteAlbumParams) (bool, error) {
	r, err := v.API.Request("market.deleteAlbum", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// MarketRemoveFromAlbumParams are params for APIMarket.RemoveFromAlbum
type MarketRemoveFromAlbumParams struct {
	// ID of an item owner community.
	OwnerID int `url:"owner_id"`
	// Item ID.
	ItemID int `url:"item_id"`
	// Collections IDs to remove item from.
	AlbumIDs CSVIntSlice `url:"album_ids"`
}

// RemoveFromAlbum Removes an item from one or multiple collections.
func (v APIMarket) RemoveFromAlbum(params MarketRemoveFromAlbumParams) (bool, error) {
	r, err := v.API.Request("market.removeFromAlbum", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// MarketAddToAlbumParams are params for APIMarket.AddToAlbum
type MarketAddToAlbumParams struct {
	// ID of an item owner community.
	OwnerID int `url:"owner_id"`
	// Item ID.
	ItemID int `url:"item_id"`
	// Collections IDs to add item to.
	AlbumIDs CSVIntSlice `url:"album_ids"`
}

// AddToAlbum Adds an item to one or multiple collections.
func (v APIMarket) AddToAlbum(params MarketAddToAlbumParams) (bool, error) {
	r, err := v.API.Request("market.addToAlbum", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}
