package vkapi

import (
	"encoding/json"
	"strconv"

	"github.com/stek29/vk"
)

// Market implements VK API namespace `market`
type Market struct {
	API vk.API
}

// MarketGetParams are params for Market.Get
type MarketGetParams struct {
	// ID of an item owner community, "Note that community id in the 'owner_id' parameter should be negative number. For example 'owner_id'=-1 matches the [vk.com/apiclub|VK API] community "
	OwnerID int `url:"owner_id"`
	AlbumID int `url:"album_id,omitempty"`
	// Number of items to return.
	Count int `url:"count,omitempty"`
	// Offset needed to return a specific subset of results.
	Offset int `url:"offset,omitempty"`
	// '1' – method will return additional fields: 'likes, can_comment, car_repost, photos'. These parameters are not returned by default.
	Extended bool `url:"extended,omitempty"`
}

// MarketGetResponse is response for Market.Get
// Either MarketGetResponseNormal or MarketGetResponseExtended, depending on Extended flag
type MarketGetResponse interface {
	isMarketGet()
}

// MarketGetResponseNormal is non-extended version of MarketGetResponse
//easyjson:json
type MarketGetResponseNormal struct {
	// Total number
	Count int             `json:"count,omitempty"`
	Items []vk.MarketItem `json:"items,omitempty"`
}

func (MarketGetResponseNormal) isMarketGet() {}

// MarketGetResponseExtended is extended version of MarketGetResponse
//easyjson:json
type MarketGetResponseExtended struct {
	// Total number
	Count int             `json:"count,omitempty"`
	Items []vk.MarketItem `json:"items,omitempty"`
}

func (MarketGetResponseExtended) isMarketGet() {}

// Get Returns items list for a community.
func (v Market) Get(params MarketGetParams) (MarketGetResponse, error) {
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

// MarketGetByIDParams are params for Market.GetByID
type MarketGetByIDParams struct {
	// Comma-separated ids list: {user id}_{item id}. If an item belongs to a community -{community id} is used. " 'Videos' value example: , '-4363_136089719,13245770_137352259'"
	ItemIDs CSVStringSlice `url:"item_ids"`
	// '1' – to return additional fields: 'likes, can_comment, car_repost, photos'. By default: '0'.
	Extended bool `url:"extended,omitempty"`
}

// MarketGetByIDResponse is response for Market.GetByID
// Either MarketGetByIDResponseNormal or MarketGetByIDResponseExtended, depending on Extended flag
type MarketGetByIDResponse interface {
	isMarketGetByID()
}

// MarketGetByIDResponseNormal is non-extended version of MarketGetByIDResponse
//easyjson:json
type MarketGetByIDResponseNormal struct {
	// Total number
	Count int             `json:"count,omitempty"`
	Items []vk.MarketItem `json:"items,omitempty"`
}

func (MarketGetByIDResponseNormal) isMarketGetByID() {}

// MarketGetByIDResponseExtended is extended version of MarketGetByIDResponse
//easyjson:json
type MarketGetByIDResponseExtended struct {
	// Total number
	Count int             `json:"count,omitempty"`
	Items []vk.MarketItem `json:"items,omitempty"`
}

func (MarketGetByIDResponseExtended) isMarketGetByID() {}

// GetByID Returns information about market items by their ids.
func (v Market) GetByID(params MarketGetByIDParams) (MarketGetByIDResponse, error) {
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

// MarketSearchParams are params for Market.Search
type MarketSearchParams struct {
	// ID of an items owner community.
	OwnerID int `url:"owner_id"`
	AlbumID int `url:"album_id,omitempty"`
	// Search query, for example "pink slippers".
	Q string `url:"q,omitempty"`
	// Minimum item price value.
	PriceFrom int `url:"price_from,omitempty"`
	// Maximum item price value.
	PriceTo int `url:"price_to,omitempty"`
	// Comma-separated tag IDs list.
	Tags CSVIntSlice `url:"tags,omitempty"`
	Sort int         `url:"sort,omitempty"`
	// '0' — do not use reverse order, '1' — use reverse order
	Rev int `url:"rev,omitempty"`
	// Offset needed to return a specific subset of results.
	Offset int `url:"offset,omitempty"`
	// Number of items to return.
	Count int `url:"count,omitempty"`
	// '1' – to return additional fields: 'likes, can_comment, car_repost, photos'. By default: '0'.
	Extended bool `url:"extended,omitempty"`
	Status   int  `url:"status,omitempty"`
}

// MarketSearchResponse is response for Market.Search
// Either MarketSearchResponseNormal or MarketSearchResponseExtended, depending on Extended flag
type MarketSearchResponse interface {
	isMarketSearch()
}

// MarketSearchResponseNormal is non-extended version of MarketSearchResponse
//easyjson:json
type MarketSearchResponseNormal struct {
	// Total number
	Count int             `json:"count,omitempty"`
	Items []vk.MarketItem `json:"items,omitempty"`
}

func (MarketSearchResponseNormal) isMarketSearch() {}

// MarketSearchResponseExtended is extended version of MarketSearchResponse
//easyjson:json
type MarketSearchResponseExtended struct {
	// Total number
	Count int             `json:"count,omitempty"`
	Items []vk.MarketItem `json:"items,omitempty"`
}

func (MarketSearchResponseExtended) isMarketSearch() {}

// Search Searches market items in a community's catalog
func (v Market) Search(params MarketSearchParams) (MarketSearchResponse, error) {
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

// MarketGetAlbumsParams are params for Market.GetAlbums
type MarketGetAlbumsParams struct {
	// ID of an items owner community.
	OwnerID int `url:"owner_id"`
	// Offset needed to return a specific subset of results.
	Offset int `url:"offset,omitempty"`
	// Number of items to return.
	Count int `url:"count,omitempty"`
}

// MarketGetAlbumsResponse is response for Market.GetAlbums
//easyjson:json
type MarketGetAlbumsResponse struct {
	// Total number
	Count int              `json:"count,omitempty"`
	Items []vk.MarketAlbum `json:"items,omitempty"`
}

// GetAlbums Returns community's collections list.
func (v Market) GetAlbums(params MarketGetAlbumsParams) (*MarketGetAlbumsResponse, error) {
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

// MarketGetAlbumByIDParams are params for Market.GetAlbumByID
type MarketGetAlbumByIDParams struct {
	// identifier of an album owner community, "Note that community id in the 'owner_id' parameter should be negative number. For example 'owner_id'=-1 matches the [vk.com/apiclub|VK API] community "
	OwnerID int `url:"owner_id"`
	// collections identifiers to obtain data from
	AlbumIDs CSVIntSlice `url:"album_ids"`
}

// MarketGetAlbumByIDResponse is response for Market.GetAlbumByID
//easyjson:json
type MarketGetAlbumByIDResponse struct {
	// Total number
	Count int              `json:"count,omitempty"`
	Items []vk.MarketAlbum `json:"items,omitempty"`
}

// GetAlbumByID Returns items album's data
func (v Market) GetAlbumByID(params MarketGetAlbumByIDParams) (*MarketGetAlbumByIDResponse, error) {
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

// MarketCreateCommentParams are params for Market.CreateComment
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

// MarketCreateCommentResponse is response for Market.CreateComment
// Comment ID
type MarketCreateCommentResponse int

// CreateComment Creates a new comment for an item.
func (v Market) CreateComment(params MarketCreateCommentParams) (MarketCreateCommentResponse, error) {
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

// MarketGetCommentsParams are params for Market.GetComments
type MarketGetCommentsParams struct {
	// ID of an item owner community
	OwnerID int `url:"owner_id"`
	// Item ID.
	ItemID int `url:"item_id"`
	// '1' — to return likes info.
	NeedLikes bool `url:"need_likes,omitempty"`
	// ID of a comment to start a list from (details below).
	StartCommentID int `url:"start_comment_id,omitempty"`
	Offset         int `url:"offset,omitempty"`
	// Number of results to return.
	Count int `url:"count,omitempty"`
	// Sort order ('asc' — from old to new, 'desc' — from new to old)
	Sort string `url:"sort,omitempty"`
	// '1' — comments will be returned as numbered objects, in addition lists of 'profiles' and 'groups' objects will be returned.
	Extended bool `url:"extended,omitempty"`
	// List of additional profile fields to return. See the [vk.com/dev/fields|details]
	Fields CSVStringSlice `url:"fields,omitempty"`
}

// MarketGetCommentsResponse is response for Market.GetComments
//easyjson:json
type MarketGetCommentsResponse struct {
	// Total number
	Count int          `json:"count,omitempty"`
	Items []vk.Comment `json:"items,omitempty"`
}

// GetComments Returns comments list for an item.
func (v Market) GetComments(params MarketGetCommentsParams) (*MarketGetCommentsResponse, error) {
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

// MarketDeleteCommentParams are params for Market.DeleteComment
type MarketDeleteCommentParams struct {
	// identifier of an item owner community, "Note that community id in the 'owner_id' parameter should be negative number. For example 'owner_id'=-1 matches the [vk.com/apiclub|VK API] community "
	OwnerID int `url:"owner_id"`
	// comment id
	CommentID int `url:"comment_id"`
}

// DeleteComment Deletes an item's comment
func (v Market) DeleteComment(params MarketDeleteCommentParams) (bool, error) {
	r, err := v.API.Request("market.deleteComment", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// MarketRestoreCommentParams are params for Market.RestoreComment
type MarketRestoreCommentParams struct {
	// identifier of an item owner community, "Note that community id in the 'owner_id' parameter should be negative number. For example 'owner_id'=-1 matches the [vk.com/apiclub|VK API] community "
	OwnerID int `url:"owner_id"`
	// deleted comment id
	CommentID int `url:"comment_id"`
}

// RestoreComment Restores a recently deleted comment
func (v Market) RestoreComment(params MarketRestoreCommentParams) (bool, error) {
	r, err := v.API.Request("market.restoreComment", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// MarketEditCommentParams are params for Market.EditComment
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
func (v Market) EditComment(params MarketEditCommentParams) (bool, error) {
	r, err := v.API.Request("market.editComment", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// MarketReportCommentParams are params for Market.ReportComment
type MarketReportCommentParams struct {
	// ID of an item owner community.
	OwnerID int `url:"owner_id"`
	// Comment ID.
	CommentID int `url:"comment_id"`
	// Complaint reason. Possible values: *'0' — spam,, *'1' — child porn,, *'2' — extremism,, *'3' — violence,, *'4' — drugs propaganda,, *'5' — adult materials,, *'6' — insult.
	Reason int `url:"reason"`
}

// ReportComment Sends a complaint to the item's comment.
func (v Market) ReportComment(params MarketReportCommentParams) (bool, error) {
	r, err := v.API.Request("market.reportComment", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// MarketGetCategoriesParams are params for Market.GetCategories
type MarketGetCategoriesParams struct {
	// Number of results to return.
	Count int `url:"count,omitempty"`
	// Offset needed to return a specific subset of results.
	Offset int `url:"offset,omitempty"`
}

// MarketGetCategoriesResponse is response for Market.GetCategories
//easyjson:json
type MarketGetCategoriesResponse struct {
	// Total number
	Count int                 `json:"count,omitempty"`
	Items []vk.MarketCategory `json:"items,omitempty"`
}

// GetCategories Returns a list of market categories.
func (v Market) GetCategories(params MarketGetCategoriesParams) (*MarketGetCategoriesResponse, error) {
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

// MarketReportParams are params for Market.Report
type MarketReportParams struct {
	// ID of an item owner community.
	OwnerID int `url:"owner_id"`
	// Item ID.
	ItemID int `url:"item_id"`
	// Complaint reason. Possible values: *'0' — spam,, *'1' — child porn,, *'2' — extremism,, *'3' — violence,, *'4' — drugs propaganda,, *'5' — adult materials,, *'6' — insult.
	Reason int `url:"reason,omitempty"`
}

// Report Sends a complaint to the item.
func (v Market) Report(params MarketReportParams) (bool, error) {
	r, err := v.API.Request("market.report", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// MarketAddParams are params for Market.Add
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
	// Url for button in market item.
	URL string `url:"url,omitempty"`
}

// MarketAddResponse is response for Market.Add
//easyjson:json
type MarketAddResponse struct {
	// Item ID
	MarketItemID int `json:"market_item_id,omitempty"`
}

// Add Ads a new item to the market.
func (v Market) Add(params MarketAddParams) (*MarketAddResponse, error) {
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

// MarketEditParams are params for Market.Edit
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
	// Url for button in market item.
	URL string `url:"url,omitempty"`
}

// Edit Edits an item.
func (v Market) Edit(params MarketEditParams) (bool, error) {
	r, err := v.API.Request("market.edit", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// MarketDeleteParams are params for Market.Delete
type MarketDeleteParams struct {
	// ID of an item owner community.
	OwnerID int `url:"owner_id"`
	// Item ID.
	ItemID int `url:"item_id"`
}

// Delete Deletes an item.
func (v Market) Delete(params MarketDeleteParams) (bool, error) {
	r, err := v.API.Request("market.delete", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// MarketRestoreParams are params for Market.Restore
type MarketRestoreParams struct {
	// ID of an item owner community.
	OwnerID int `url:"owner_id"`
	// Deleted item ID.
	ItemID int `url:"item_id"`
}

// Restore Restores recently deleted item
func (v Market) Restore(params MarketRestoreParams) (bool, error) {
	r, err := v.API.Request("market.restore", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// MarketReorderItemsParams are params for Market.ReorderItems
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
func (v Market) ReorderItems(params MarketReorderItemsParams) (bool, error) {
	r, err := v.API.Request("market.reorderItems", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// MarketReorderAlbumsParams are params for Market.ReorderAlbums
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
func (v Market) ReorderAlbums(params MarketReorderAlbumsParams) (bool, error) {
	r, err := v.API.Request("market.reorderAlbums", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// MarketAddAlbumParams are params for Market.AddAlbum
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

// MarketAddAlbumResponse is response for Market.AddAlbum
//easyjson:json
type MarketAddAlbumResponse struct {
	// Album ID
	MarketAlbumID int `json:"market_album_id,omitempty"`
}

// AddAlbum Creates new collection of items
func (v Market) AddAlbum(params MarketAddAlbumParams) (*MarketAddAlbumResponse, error) {
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

// MarketEditAlbumParams are params for Market.EditAlbum
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
func (v Market) EditAlbum(params MarketEditAlbumParams) (bool, error) {
	r, err := v.API.Request("market.editAlbum", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// MarketDeleteAlbumParams are params for Market.DeleteAlbum
type MarketDeleteAlbumParams struct {
	// ID of an collection owner community.
	OwnerID int `url:"owner_id"`
	// Collection ID.
	AlbumID int `url:"album_id"`
}

// DeleteAlbum Deletes a collection of items.
func (v Market) DeleteAlbum(params MarketDeleteAlbumParams) (bool, error) {
	r, err := v.API.Request("market.deleteAlbum", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// MarketRemoveFromAlbumParams are params for Market.RemoveFromAlbum
type MarketRemoveFromAlbumParams struct {
	// ID of an item owner community.
	OwnerID int `url:"owner_id"`
	// Item ID.
	ItemID int `url:"item_id"`
	// Collections IDs to remove item from.
	AlbumIDs CSVIntSlice `url:"album_ids"`
}

// RemoveFromAlbum Removes an item from one or multiple collections.
func (v Market) RemoveFromAlbum(params MarketRemoveFromAlbumParams) (bool, error) {
	r, err := v.API.Request("market.removeFromAlbum", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// MarketAddToAlbumParams are params for Market.AddToAlbum
type MarketAddToAlbumParams struct {
	// ID of an item owner community.
	OwnerID int `url:"owner_id"`
	// Item ID.
	ItemID int `url:"item_id"`
	// Collections IDs to add item to.
	AlbumIDs CSVIntSlice `url:"album_ids"`
}

// AddToAlbum Adds an item to one or multiple collections.
func (v Market) AddToAlbum(params MarketAddToAlbumParams) (bool, error) {
	r, err := v.API.Request("market.addToAlbum", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}
