package vkCallbackApi

import (
	"encoding/json"
	"strconv"
)

// APIBoard implements VK API namespace `board`
type APIBoard struct {
	API *API
}

// BoardGetTopicsParams are params for APIBoard.GetTopics
type BoardGetTopicsParams struct {
	// ID of the community that owns the discussion board.
	GroupID int `url:"group_id"`
	// IDs of topics to be returned (100 maximum). By default, all topics are returned. If this parameter is set, the 'order', 'offset', and 'count' parameters are ignored.
	TopicIDs CSVIntSlice `url:"topic_ids,omitempty"`
	// Sort order: '1' — by date updated in reverse chronological order. '2' — by date created in reverse chronological order. '-1' — by date updated in chronological order. '-2' — by date created in chronological order. If no sort order is specified, topics are returned in the order specified by the group administrator. Pinned topics are returned first, regardless of the sorting.
	Order int `url:"order,omitempty"`
	// Offset needed to return a specific subset of topics.
	Offset int `url:"offset,omitempty"`
	// Number of topics to return.
	Count int `url:"count,omitempty"`
	// '1' — to return information about users who created topics or who posted there last, '0' — to return no additional fields (default)
	Extended bool `url:"extended,omitempty"`
	// '1' — to return the first comment in each topic,, '2' — to return the last comment in each topic,, '0' — to return no comments. By default: '0'.
	Preview int `url:"preview,omitempty"`
	// Number of characters after which to truncate the previewed comment. To preview the full comment, specify '0'.
	PreviewLength int `url:"preview_length,omitempty"`
}

// BoardGetTopicsResponse is response for APIBoard.GetTopics
// Either BoardGetTopicsResponseNormal or BoardGetTopicsResponseExtended, depending on Extended flag
type BoardGetTopicsResponse interface {
	isBoardGetTopics()
}

// BoardGetTopicsResponseNormal is non-extended version of BoardGetTopicsResponse
//easyjson:json
type BoardGetTopicsResponseNormal struct {
	// Total number
	Count int          `json:"count,omitempty"`
	Items []BoardTopic `json:"items,omitempty"`
	// Sort type
	DefaultOrder int `json:"default_order,omitempty"`
	// Information whether current user can add topic
	CanAddTopics BoolInt `json:"can_add_topics,omitempty"`
}

func (BoardGetTopicsResponseNormal) isBoardGetTopics() {}

// BoardGetTopicsResponseExtended is extended version of BoardGetTopicsResponse
//easyjson:json
type BoardGetTopicsResponseExtended struct {
	// Total number
	Count int          `json:"count,omitempty"`
	Items []BoardTopic `json:"items,omitempty"`
	// Sort type
	DefaultOrder int `json:"default_order,omitempty"`
	// Information whether current user can add topic
	CanAddTopics BoolInt `json:"can_add_topics,omitempty"`
	Profiles     []User  `json:"profiles,omitempty"`
}

func (BoardGetTopicsResponseExtended) isBoardGetTopics() {}

// GetTopics Returns a list of topics on a community's discussion board.
func (v APIBoard) GetTopics(params BoardGetTopicsParams) (BoardGetTopicsResponse, error) {
	r, err := v.API.Request("board.getTopics", params)
	if err != nil {
		return nil, err
	}

	var resp BoardGetTopicsResponse
	if params.Extended {
		var tmp BoardGetTopicsResponseExtended
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	} else {
		var tmp BoardGetTopicsResponseNormal
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// BoardGetCommentsParams are params for APIBoard.GetComments
type BoardGetCommentsParams struct {
	// ID of the community that owns the discussion board.
	GroupID int `url:"group_id"`
	// Topic ID.
	TopicID int `url:"topic_id"`
	// '1' — to return the 'likes' field, '0' — not to return the 'likes' field (default)
	NeedLikes      bool `url:"need_likes,omitempty"`
	StartCommentID int  `url:"start_comment_id,omitempty"`
	// Offset needed to return a specific subset of comments.
	Offset int `url:"offset,omitempty"`
	// Number of comments to return.
	Count int `url:"count,omitempty"`
	// '1' — to return information about users who posted comments, '0' — to return no additional fields (default)
	Extended bool `url:"extended,omitempty"`
	// Sort order: 'asc' — by creation date in chronological order, 'desc' — by creation date in reverse chronological order,
	Sort string `url:"sort,omitempty"`
}

// BoardGetCommentsResponse is response for APIBoard.GetComments
// Either BoardGetCommentsResponseNormal or BoardGetCommentsResponseExtended, depending on Extended flag
type BoardGetCommentsResponse interface {
	isBoardGetComments()
}

// BoardGetCommentsResponseNormal is non-extended version of BoardGetCommentsResponse
//easyjson:json
type BoardGetCommentsResponseNormal struct {
	// Total number
	Count int            `json:"count,omitempty"`
	Items []CommentBoard `json:"items,omitempty"`
	Poll  Poll           `json:"poll,omitempty"`
}

func (BoardGetCommentsResponseNormal) isBoardGetComments() {}

// BoardGetCommentsResponseExtended is extended version of BoardGetCommentsResponse
//easyjson:json
type BoardGetCommentsResponseExtended struct {
	// Total number
	Count    int            `json:"count,omitempty"`
	Items    []CommentBoard `json:"items,omitempty"`
	Poll     Poll           `json:"poll,omitempty"`
	Profiles []User         `json:"profiles,omitempty"`
	Groups   []Group        `json:"groups,omitempty"`
}

func (BoardGetCommentsResponseExtended) isBoardGetComments() {}

// GetComments Returns a list of comments on a topic on a community's discussion board.
func (v APIBoard) GetComments(params BoardGetCommentsParams) (BoardGetCommentsResponse, error) {
	r, err := v.API.Request("board.getComments", params)
	if err != nil {
		return nil, err
	}

	var resp BoardGetCommentsResponse
	if params.Extended {
		var tmp BoardGetCommentsResponseExtended
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	} else {
		var tmp BoardGetCommentsResponseNormal
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// BoardAddTopicParams are params for APIBoard.AddTopic
type BoardAddTopicParams struct {
	// ID of the community that owns the discussion board.
	GroupID int `url:"group_id"`
	// Topic title.
	Title string `url:"title"`
	// Text of the topic.
	Text string `url:"text,omitempty"`
	// For a community: '1' — to post the topic as by the community, '0' — to post the topic as by the user (default)
	FromGroup bool `url:"from_group,omitempty"`
	// List of media objects attached to the topic, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media object: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, '<owner_id>' — ID of the media owner. '<media_id>' — Media ID. Example: "photo100172_166443618,photo66748_265827614", , "NOTE: If you try to attach more than one reference, an error will be thrown.",
	Attachments CSVStringSlice `url:"attachments,omitempty"`
}

// BoardAddTopicResponse is response for APIBoard.AddTopic
// Topic ID
type BoardAddTopicResponse int

// AddTopic Creates a new topic on a community's discussion board.
func (v APIBoard) AddTopic(params BoardAddTopicParams) (BoardAddTopicResponse, error) {
	r, err := v.API.Request("board.addTopic", params)
	if err != nil {
		return 0, err
	}

	var resp BoardAddTopicResponse

	var cnv int
	cnv, err = strconv.Atoi(string(r))
	resp = BoardAddTopicResponse(cnv)

	if err != nil {
		return 0, err
	}
	return resp, nil
}

// BoardCreateCommentParams are params for APIBoard.CreateComment
type BoardCreateCommentParams struct {
	// ID of the community that owns the discussion board.
	GroupID int `url:"group_id"`
	// ID of the topic to be commented on.
	TopicID int `url:"topic_id"`
	// (Required if 'attachments' is not set.) Text of the comment.
	Message string `url:"message,omitempty"`
	// (Required if 'text' is not set.) List of media objects attached to the comment, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media object: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, '<owner_id>' — ID of the media owner. '<media_id>' — Media ID.
	Attachments CSVStringSlice `url:"attachments,omitempty"`
	// '1' — to post the comment as by the community, '0' — to post the comment as by the user (default)
	FromGroup bool `url:"from_group,omitempty"`
	// Sticker ID.
	StickerID int `url:"sticker_id,omitempty"`
	// Unique identifier to avoid repeated comments.
	GUID string `url:"guid,omitempty"`
}

// BoardCreateCommentResponse is response for APIBoard.CreateComment
// Comment ID
type BoardCreateCommentResponse int

// CreateComment Adds a comment on a topic on a community's discussion board.
func (v APIBoard) CreateComment(params BoardCreateCommentParams) (BoardCreateCommentResponse, error) {
	r, err := v.API.Request("board.createComment", params)
	if err != nil {
		return 0, err
	}

	var resp BoardCreateCommentResponse

	var cnv int
	cnv, err = strconv.Atoi(string(r))
	resp = BoardCreateCommentResponse(cnv)

	if err != nil {
		return 0, err
	}
	return resp, nil
}

// BoardDeleteTopicParams are params for APIBoard.DeleteTopic
type BoardDeleteTopicParams struct {
	// ID of the community that owns the discussion board.
	GroupID int `url:"group_id"`
	// Topic ID.
	TopicID int `url:"topic_id"`
}

// DeleteTopic Deletes a topic from a community's discussion board.
func (v APIBoard) DeleteTopic(params BoardDeleteTopicParams) (bool, error) {
	r, err := v.API.Request("board.deleteTopic", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// BoardEditTopicParams are params for APIBoard.EditTopic
type BoardEditTopicParams struct {
	// ID of the community that owns the discussion board.
	GroupID int `url:"group_id"`
	// Topic ID.
	TopicID int `url:"topic_id"`
	// New title of the topic.
	Title string `url:"title"`
}

// EditTopic Edits the title of a topic on a community's discussion board.
func (v APIBoard) EditTopic(params BoardEditTopicParams) (bool, error) {
	r, err := v.API.Request("board.editTopic", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// BoardEditCommentParams are params for APIBoard.EditComment
type BoardEditCommentParams struct {
	// ID of the community that owns the discussion board.
	GroupID int `url:"group_id"`
	// Topic ID.
	TopicID int `url:"topic_id"`
	// ID of the comment on the topic.
	CommentID int `url:"comment_id"`
	// (Required if 'attachments' is not set). New comment text.
	Message string `url:"message,omitempty"`
	// (Required if 'message' is not set.) List of media objects attached to the comment, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media object: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, '<owner_id>' — ID of the media owner. '<media_id>' — Media ID. Example: "photo100172_166443618,photo66748_265827614"
	Attachments CSVStringSlice `url:"attachments,omitempty"`
}

// EditComment Edits a comment on a topic on a community's discussion board.
func (v APIBoard) EditComment(params BoardEditCommentParams) (bool, error) {
	r, err := v.API.Request("board.editComment", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// BoardRestoreCommentParams are params for APIBoard.RestoreComment
type BoardRestoreCommentParams struct {
	// ID of the community that owns the discussion board.
	GroupID int `url:"group_id"`
	// Topic ID.
	TopicID int `url:"topic_id"`
	// Comment ID.
	CommentID int `url:"comment_id"`
}

// RestoreComment Restores a comment deleted from a topic on a community's discussion board.
func (v APIBoard) RestoreComment(params BoardRestoreCommentParams) (bool, error) {
	r, err := v.API.Request("board.restoreComment", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// BoardDeleteCommentParams are params for APIBoard.DeleteComment
type BoardDeleteCommentParams struct {
	// ID of the community that owns the discussion board.
	GroupID int `url:"group_id"`
	// Topic ID.
	TopicID int `url:"topic_id"`
	// Comment ID.
	CommentID int `url:"comment_id"`
}

// DeleteComment Deletes a comment on a topic on a community's discussion board.
func (v APIBoard) DeleteComment(params BoardDeleteCommentParams) (bool, error) {
	r, err := v.API.Request("board.deleteComment", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// BoardOpenTopicParams are params for APIBoard.OpenTopic
type BoardOpenTopicParams struct {
	// ID of the community that owns the discussion board.
	GroupID int `url:"group_id"`
	// Topic ID.
	TopicID int `url:"topic_id"`
}

// OpenTopic Re-opens a previously closed topic on a community's discussion board.
func (v APIBoard) OpenTopic(params BoardOpenTopicParams) (bool, error) {
	r, err := v.API.Request("board.openTopic", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// BoardCloseTopicParams are params for APIBoard.CloseTopic
type BoardCloseTopicParams struct {
	// ID of the community that owns the discussion board.
	GroupID int `url:"group_id"`
	// Topic ID.
	TopicID int `url:"topic_id"`
}

// CloseTopic Closes a topic on a community's discussion board so that comments cannot be posted.
func (v APIBoard) CloseTopic(params BoardCloseTopicParams) (bool, error) {
	r, err := v.API.Request("board.closeTopic", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// BoardFixTopicParams are params for APIBoard.FixTopic
type BoardFixTopicParams struct {
	// ID of the community that owns the discussion board.
	GroupID int `url:"group_id"`
	// Topic ID.
	TopicID int `url:"topic_id"`
}

// FixTopic Pins a topic (fixes its place) to the top of a community's discussion board.
func (v APIBoard) FixTopic(params BoardFixTopicParams) (bool, error) {
	r, err := v.API.Request("board.fixTopic", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// BoardUnfixTopicParams are params for APIBoard.UnfixTopic
type BoardUnfixTopicParams struct {
	// ID of the community that owns the discussion board.
	GroupID int `url:"group_id"`
	// Topic ID.
	TopicID int `url:"topic_id"`
}

// UnfixTopic Unpins a pinned topic from the top of a community's discussion board.
func (v APIBoard) UnfixTopic(params BoardUnfixTopicParams) (bool, error) {
	r, err := v.API.Request("board.unfixTopic", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}
