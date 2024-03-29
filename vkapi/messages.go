package vkapi

import (
	"encoding/json"
	"strconv"

	"github.com/stek29/vk"
)

// Messages implements VK API namespace `messages`
type Messages struct {
	API vk.API
}

// MessagesJoinChatByInviteLinkParams are params for Messages.JoinChatByInviteLink
type MessagesJoinChatByInviteLinkParams struct {
	// Invitation link.
	Link string `url:"link"`
}

// MessagesJoinChatByInviteLinkResponse is response for Messages.JoinChatByInviteLink
//easyjson:json
type MessagesJoinChatByInviteLinkResponse struct {
	ChatID int `json:"chat_id,omitempty"`
}

// JoinChatByInviteLink does messages.joinChatByInviteLink
func (v Messages) JoinChatByInviteLink(params MessagesJoinChatByInviteLinkParams) (*MessagesJoinChatByInviteLinkResponse, error) {
	r, err := v.API.Request("messages.joinChatByInviteLink", params)
	if err != nil {
		return nil, err
	}

	var resp MessagesJoinChatByInviteLinkResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// MessagesGetInviteLinkParams are params for Messages.GetInviteLink
type MessagesGetInviteLinkParams struct {
	// Destination ID.
	PeerID int `url:"peer_id"`
	// 1 — to generate new link (revoke previous), 0 — to return previous link.
	Reset bool `url:"reset,omitempty"`
	// Group ID
	GroupID int `url:"group_id,omitempty"`
}

// MessagesGetInviteLinkResponse is response for Messages.GetInviteLink
//easyjson:json
type MessagesGetInviteLinkResponse struct {
	Link string `json:"link,omitempty"`
}

// GetInviteLink does messages.getInviteLink
func (v Messages) GetInviteLink(params MessagesGetInviteLinkParams) (*MessagesGetInviteLinkResponse, error) {
	r, err := v.API.Request("messages.getInviteLink", params)
	if err != nil {
		return nil, err
	}

	var resp MessagesGetInviteLinkResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// MessagesGetConversationsParams are params for Messages.GetConversations
type MessagesGetConversationsParams struct {
	// Offset needed to return a specific subset of conversations.
	Offset int `url:"offset,omitempty"`
	// Number of conversations to return.
	Count int `url:"count,omitempty"`
	// Filter to apply: 'all' — all conversations, 'unread' — conversations with unread messages, 'important' — conversations, marked as important (only for community messages), 'unanswered' — conversations, marked as unanswered (only for community messages)
	Filter string `url:"filter,omitempty"`
	// '1' — return extra information about users and communities
	Extended bool `url:"extended,omitempty"`
	// ID of the message from what to return dialogs.
	StartMessageID int `url:"start_message_id,omitempty"`
	// Profile and communities fields to return.
	Fields CSVStringSlice `url:"fields,omitempty"`
	// Group ID (for group messages with group access token)
	GroupID int `url:"group_id,omitempty"`
}

// MessagesGetConversationsResponse is response for Messages.GetConversations
//easyjson:json
type MessagesGetConversationsResponse struct {
	// Total number
	Count int `json:"count,omitempty"`
	// Unread dialogs number
	UnreadCount int `json:"unread_count,omitempty"`
	Items       []struct {
		Conversation vk.Conversation `json:"conversation,omitempty"`
		LastMessage  vk.Message      `json:"last_message,omitempty"`
	} `json:"items,omitempty"`
	Profiles []vk.User  `json:"profiles,omitempty"`
	Groups   []vk.Group `json:"groups,omitempty"`
}

// GetConversations Returns a list of the current user's conversations.
func (v Messages) GetConversations(params MessagesGetConversationsParams) (*MessagesGetConversationsResponse, error) {
	r, err := v.API.Request("messages.getConversations", params)
	if err != nil {
		return nil, err
	}

	var resp MessagesGetConversationsResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// MessagesGetConversationsByIDParams are params for Messages.GetConversationsByID
type MessagesGetConversationsByIDParams struct {
	// Destination IDs. "For user: 'User ID', e.g. '12345'. For chat: '2000000000' + 'chat_id', e.g. '2000000001'. For community: '- community ID', e.g. '-12345'. "
	PeerIDs CSVIntSlice `url:"peer_ids"`
	// Return extended properties
	Extended bool `url:"extended,omitempty"`
	// Profile and communities fields to return.
	Fields CSVStringSlice `url:"fields,omitempty"`
	// Group ID (for group messages with group access token)
	GroupID int `url:"group_id,omitempty"`
}

// MessagesGetConversationsByIDResponse is response for Messages.GetConversationsByID
//easyjson:json
type MessagesGetConversationsByIDResponse struct {
	// Total number
	Count int `json:"count,omitempty"`
	Items []struct {
		Conversation vk.Conversation `json:"conversation,omitempty"`
		LastMessage  vk.Message      `json:"last_message,omitempty"`
	} `json:"items,omitempty"`
}

// GetConversationsByID Returns conversations by their IDs
func (v Messages) GetConversationsByID(params MessagesGetConversationsByIDParams) (*MessagesGetConversationsByIDResponse, error) {
	r, err := v.API.Request("messages.getConversationsById", params)
	if err != nil {
		return nil, err
	}

	var resp MessagesGetConversationsByIDResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// MessagesGetByIDParams are params for Messages.GetByID
type MessagesGetByIDParams struct {
	// Message IDs.
	MessageIDs CSVIntSlice `url:"message_ids"`
	// Number of characters after which to truncate a previewed message. To preview the full message, specify '0'. "NOTE: Messages are not truncated by default. Messages are truncated by words."
	PreviewLength int `url:"preview_length,omitempty"`
	// Information whether the response should be extended
	Extended bool `url:"extended,omitempty"`
	// Profile fields to return.
	Fields CSVStringSlice `url:"fields,omitempty"`
	// Group ID (for group messages with group access token)
	GroupID int `url:"group_id,omitempty"`
}

// MessagesGetByIDResponse is response for Messages.GetByID
//easyjson:json
type MessagesGetByIDResponse struct {
	// Total number
	Count int          `json:"count,omitempty"`
	Items []vk.Message `json:"items,omitempty"`
}

// GetByID Returns messages by their IDs.
func (v Messages) GetByID(params MessagesGetByIDParams) (*MessagesGetByIDResponse, error) {
	r, err := v.API.Request("messages.getById", params)
	if err != nil {
		return nil, err
	}

	var resp MessagesGetByIDResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// MessagesGetByConversationMessageIDParams are params for Messages.GetByConversationMessageID
type MessagesGetByConversationMessageIDParams struct {
	// Destination ID. "For user: 'User ID', e.g. '12345'. For chat: '2000000000' + 'chat_id', e.g. '2000000001'. For community: '- community ID', e.g. '-12345'. "
	PeerID int `url:"peer_id"`
	// Conversation message IDs.
	ConversationMessageIDs CSVIntSlice `url:"conversation_message_ids"`
	// Information whether the response should be extended
	Extended bool `url:"extended,omitempty"`
	// Profile fields to return.
	Fields CSVStringSlice `url:"fields,omitempty"`
	// Group ID (for group messages with group access token)
	GroupID int `url:"group_id,omitempty"`
}

// MessagesGetByConversationMessageIDResponse is response for Messages.GetByConversationMessageID
//easyjson:json
type MessagesGetByConversationMessageIDResponse struct {
	// Total number
	Count int          `json:"count,omitempty"`
	Items []vk.Message `json:"items,omitempty"`
}

// GetByConversationMessageID Returns messages by their IDs within the conversation.
func (v Messages) GetByConversationMessageID(params MessagesGetByConversationMessageIDParams) (*MessagesGetByConversationMessageIDResponse, error) {
	r, err := v.API.Request("messages.getByConversationMessageId", params)
	if err != nil {
		return nil, err
	}

	var resp MessagesGetByConversationMessageIDResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// MessagesSearchParams are params for Messages.Search
type MessagesSearchParams struct {
	// Search query string.
	Q string `url:"q,omitempty"`
	// Destination ID. "For user: 'User ID', e.g. '12345'. For chat: '2000000000' + 'chat_id', e.g. '2000000001'. For community: '- community ID', e.g. '-12345'. "
	PeerID int `url:"peer_id,omitempty"`
	// Date to search message before in Unixtime.
	Date int `url:"date,omitempty"`
	// Number of characters after which to truncate a previewed message. To preview the full message, specify '0'. "NOTE: Messages are not truncated by default. Messages are truncated by words."
	PreviewLength int `url:"preview_length,omitempty"`
	// Offset needed to return a specific subset of messages.
	Offset int `url:"offset,omitempty"`
	// Number of messages to return.
	Count    int            `url:"count,omitempty"`
	Extended bool           `url:"extended,omitempty"`
	Fields   CSVStringSlice `url:"fields,omitempty"`
	// Group ID (for group messages with group access token)
	GroupID int `url:"group_id,omitempty"`
}

// MessagesSearchResponse is response for Messages.Search
//easyjson:json
type MessagesSearchResponse struct {
	// Total number
	Count int          `json:"count,omitempty"`
	Items []vk.Message `json:"items,omitempty"`
}

// Search Returns a list of the current user's private messages that match search criteria.
func (v Messages) Search(params MessagesSearchParams) (*MessagesSearchResponse, error) {
	r, err := v.API.Request("messages.search", params)
	if err != nil {
		return nil, err
	}

	var resp MessagesSearchResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// MessagesGetHistoryParams are params for Messages.GetHistory
type MessagesGetHistoryParams struct {
	// Offset needed to return a specific subset of messages.
	Offset int `url:"offset,omitempty"`
	// Number of messages to return.
	Count int `url:"count,omitempty"`
	// ID of the user whose message history you want to return.
	UserID int `url:"user_id,omitempty"`
	PeerID int `url:"peer_id,omitempty"`
	// Starting message ID from which to return history.
	StartMessageID int `url:"start_message_id,omitempty"`
	// Sort order: '1' — return messages in chronological order. '0' — return messages in reverse chronological order.
	Rev int `url:"rev,omitempty"`
	// Information whether the response should be extended
	Extended bool `url:"extended,omitempty"`
	// Profile fields to return.
	Fields CSVStringSlice `url:"fields,omitempty"`
	// Group ID (for group messages with group access token)
	GroupID int `url:"group_id,omitempty"`
}

// MessagesGetHistoryResponse is response for Messages.GetHistory
//easyjson:json
type MessagesGetHistoryResponse struct {
	// Total number
	Count    int          `json:"count,omitempty"`
	Items    []vk.Message `json:"items,omitempty"`
	Profiles []vk.User    `json:"profiles,omitempty"`
	Groups   []vk.Group   `json:"groups,omitempty"`
}

// GetHistory Returns message history for the specified user or group chat.
func (v Messages) GetHistory(params MessagesGetHistoryParams) (*MessagesGetHistoryResponse, error) {
	r, err := v.API.Request("messages.getHistory", params)
	if err != nil {
		return nil, err
	}

	var resp MessagesGetHistoryResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// MessagesGetHistoryAttachmentsParams are params for Messages.GetHistoryAttachments
type MessagesGetHistoryAttachmentsParams struct {
	// Peer ID. ", For group chat: '2000000000 + chat ID' , , For community: '-community ID'"
	PeerID int `url:"peer_id"`
	// Type of media files to return: *'photo',, *'video',, *'audio',, *'doc',, *'link'.,*'market'.,*'wall'.,*'share'
	MediaType string `url:"media_type,omitempty"`
	// Message ID to start return results from.
	StartFrom string `url:"start_from,omitempty"`
	// Number of objects to return.
	Count int `url:"count,omitempty"`
	// '1' — to return photo sizes in a
	PhotoSizes bool `url:"photo_sizes,omitempty"`
	// Additional profile [vk.com/dev/fields|fields] to return.
	Fields CSVStringSlice `url:"fields,omitempty"`
	// Group ID (for group messages with group access token)
	GroupID int `url:"group_id,omitempty"`
}

// MessagesGetHistoryAttachmentsResponse is response for Messages.GetHistoryAttachments
//easyjson:json
type MessagesGetHistoryAttachmentsResponse struct {
	Items []struct {
		// Message ID
		MessageID  int           `json:"message_id,omitempty"`
		Attachment vk.Attachment `json:"attachment,omitempty"`
	} `json:"items,omitempty"`
	// Value for pagination
	NextFrom string `json:"next_from,omitempty"`
}

// GetHistoryAttachments Returns media files from the dialog or group chat.
func (v Messages) GetHistoryAttachments(params MessagesGetHistoryAttachmentsParams) (*MessagesGetHistoryAttachmentsResponse, error) {
	r, err := v.API.Request("messages.getHistoryAttachments", params)
	if err != nil {
		return nil, err
	}

	var resp MessagesGetHistoryAttachmentsResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// MessagesSendParams are params for Messages.Send
type MessagesSendParams struct {
	// User ID (by default — current user).
	UserID int `url:"user_id,omitempty"`
	// Unique identifier to avoid resending the message.
	RandomID int `url:"random_id,omitempty"`
	// Destination ID. "For user: 'User ID', e.g. '12345'. For chat: '2000000000' + 'chat_id', e.g. '2000000001'. For community: '- community ID', e.g. '-12345'. "
	PeerID int `url:"peer_id,omitempty"`
	// User's short address (for example, 'illarionov').
	Domain string `url:"domain,omitempty"`
	// ID of conversation the message will relate to.
	ChatID int `url:"chat_id,omitempty"`
	// IDs of message recipients (if new conversation shall be started).
	UserIDs CSVIntSlice `url:"user_ids,omitempty"`
	// (Required if 'attachments' is not set.) Text of the message.
	Message string `url:"message,omitempty"`
	// Geographical latitude of a check-in, in degrees (from -90 to 90).
	Lat float32 `url:"lat,omitempty"`
	// Geographical longitude of a check-in, in degrees (from -180 to 180).
	Long float32 `url:"long,omitempty"`
	// (Required if 'message' is not set.) List of objects attached to the message, separated by commas, in the following format: "<owner_id>_<media_id>", '' — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, 'wall' — wall post, '<owner_id>' — ID of the media attachment owner. '<media_id>' — media attachment ID. Example: "photo100172_166443618"
	Attachment CSVStringSlice `url:"attachment,omitempty"`
	ReplyTo    int            `url:"reply_to,omitempty"`
	// ID of forwarded messages, separated with a comma. Listed messages of the sender will be shown in the message body at the recipient's. Example: "123,431,544"
	ForwardMessages CSVIntSlice `url:"forward_messages,omitempty"`
	// Sticker id.
	StickerID int `url:"sticker_id,omitempty"`
	// Group ID (for group messages with group access token)
	GroupID        int    `url:"group_id,omitempty"`
	Keyboard       string `url:"keyboard,omitempty"`
	Payload        string `url:"payload,omitempty"`
	DontParseLinks bool   `url:"dont_parse_links,omitempty"`
}

// MessagesSendResponse is response for Messages.Send
// Message ID
type MessagesSendResponse int

// Send Sends a message.
func (v Messages) Send(params MessagesSendParams) (MessagesSendResponse, error) {
	r, err := v.API.Request("messages.send", params)
	if err != nil {
		return 0, err
	}

	var resp MessagesSendResponse

	var cnv int
	cnv, err = strconv.Atoi(string(r))
	resp = MessagesSendResponse(cnv)

	if err != nil {
		return 0, err
	}
	return resp, nil
}

// MessagesEditParams are params for Messages.Edit
type MessagesEditParams struct {
	// Destination ID. "For user: 'User ID', e.g. '12345'. For chat: '2000000000' + 'chat_id', e.g. '2000000001'. For community: '- community ID', e.g. '-12345'. "
	PeerID int `url:"peer_id"`
	// (Required if 'attachments' is not set.) Text of the message.
	Message   string `url:"message,omitempty"`
	MessageID int    `url:"message_id"`
	// Geographical latitude of a check-in, in degrees (from -90 to 90).
	Lat float32 `url:"lat,omitempty"`
	// Geographical longitude of a check-in, in degrees (from -180 to 180).
	Long float32 `url:"long,omitempty"`
	// (Required if 'message' is not set.) List of objects attached to the message, separated by commas, in the following format: "<owner_id>_<media_id>", '' — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, 'wall' — wall post, '<owner_id>' — ID of the media attachment owner. '<media_id>' — media attachment ID. Example: "photo100172_166443618"
	Attachment CSVStringSlice `url:"attachment,omitempty"`
	// '1' — to keep forwarded, messages.
	KeepForwardMessages bool `url:"keep_forward_messages,omitempty"`
	// '1' — to keep attached snippets.
	KeepSnippets bool `url:"keep_snippets,omitempty"`
	// Group ID (for group messages with user access token)
	GroupID        int  `url:"group_id,omitempty"`
	DontParseLinks bool `url:"dont_parse_links,omitempty"`
}

// Edit Edits the message.
func (v Messages) Edit(params MessagesEditParams) (bool, error) {
	r, err := v.API.Request("messages.edit", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// MessagesDeleteParams are params for Messages.Delete
type MessagesDeleteParams struct {
	// Message IDs.
	MessageIDs CSVIntSlice `url:"message_ids,omitempty"`
	// '1' — to mark message as spam.
	Spam bool `url:"spam,omitempty"`
	// Group ID (for group messages with user access token)
	GroupID int `url:"group_id,omitempty"`
	// '1' — delete message for for all.
	DeleteForAll bool `url:"delete_for_all,omitempty"`
}

// MessagesDeleteResponse is response for Messages.Delete
//easyjson:json
type MessagesDeleteResponse ArrayOnMeth

// Delete Deletes one or more messages.
func (v Messages) Delete(params MessagesDeleteParams) (*MessagesDeleteResponse, error) {
	r, err := v.API.Request("messages.delete", params)
	if err != nil {
		return nil, err
	}

	var resp MessagesDeleteResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// MessagesDeleteConversationParams are params for Messages.DeleteConversation
type MessagesDeleteConversationParams struct {
	// User ID. To clear a chat history use 'chat_id'
	UserID int `url:"user_id,omitempty"`
	// Destination ID. "For user: 'User ID', e.g. '12345'. For chat: '2000000000' + 'chat_id', e.g. '2000000001'. For community: '- community ID', e.g. '-12345'. "
	PeerID int `url:"peer_id,omitempty"`
	// Offset needed to delete a specific subset of conversations.
	Offset int `url:"offset,omitempty"`
	// Number of conversations to delete. "NOTE: If the number of messages exceeds the maximum, the method shall be called several times."
	Count int `url:"count,omitempty"`
	// Group ID (for group messages with user access token)
	GroupID int `url:"group_id,omitempty"`
}

// MessagesDeleteConversationResponse is response for Messages.DeleteConversation
//easyjson:json
type MessagesDeleteConversationResponse struct {
	// Id of the last message, that was deleted
	LastDeletedID int `json:"last_deleted_id,omitempty"`
}

// DeleteConversation Deletes all private messages in a conversation.
func (v Messages) DeleteConversation(params MessagesDeleteConversationParams) (*MessagesDeleteConversationResponse, error) {
	r, err := v.API.Request("messages.deleteConversation", params)
	if err != nil {
		return nil, err
	}

	var resp MessagesDeleteConversationResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// MessagesPinParams are params for Messages.Pin
type MessagesPinParams struct {
	// Destination ID. "For user: 'User ID', e.g. '12345'. For chat: '2000000000' + 'Chat ID', e.g. '2000000001'. For community: '- Community ID', e.g. '-12345'. "
	PeerID    int `url:"peer_id"`
	MessageID int `url:"message_id,omitempty"`
}

// MessagesPinResponse is response for Messages.Pin
//easyjson:json
type MessagesPinResponse vk.Message

// Pin Pin a message.
func (v Messages) Pin(params MessagesPinParams) (*MessagesPinResponse, error) {
	r, err := v.API.Request("messages.pin", params)
	if err != nil {
		return nil, err
	}

	var resp MessagesPinResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// MessagesRestoreParams are params for Messages.Restore
type MessagesRestoreParams struct {
	// ID of a previously-deleted message to restore.
	MessageID int `url:"message_id"`
	// Group ID (for group messages with user access token)
	GroupID int `url:"group_id,omitempty"`
}

// Restore Restores a deleted message.
func (v Messages) Restore(params MessagesRestoreParams) (bool, error) {
	r, err := v.API.Request("messages.restore", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// MessagesMarkAsReadParams are params for Messages.MarkAsRead
type MessagesMarkAsReadParams struct {
	// Destination ID. "For user: 'User ID', e.g. '12345'. For chat: '2000000000' + 'chat_id', e.g. '2000000001'. For community: '- community ID', e.g. '-12345'. "
	PeerID int `url:"peer_id,omitempty"`
	// Message ID to start from.
	StartMessageID int `url:"start_message_id,omitempty"`
	// Group ID (for group messages with user access token)
	GroupID int `url:"group_id,omitempty"`
}

// MarkAsRead Marks messages as read.
func (v Messages) MarkAsRead(params MessagesMarkAsReadParams) (bool, error) {
	r, err := v.API.Request("messages.markAsRead", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// MessagesMarkAsImportantParams are params for Messages.MarkAsImportant
type MessagesMarkAsImportantParams struct {
	// IDs of messages to mark as important.
	MessageIDs CSVIntSlice `url:"message_ids,omitempty"`
	// '1' — to add a star (mark as important), '0' — to remove the star
	Important int `url:"important,omitempty"`
}

// MessagesMarkAsImportantResponse is response for Messages.MarkAsImportant
// Messages IDs
//easyjson:json
type MessagesMarkAsImportantResponse []int

// MarkAsImportant Marks and unmarks messages as important (starred).
func (v Messages) MarkAsImportant(params MessagesMarkAsImportantParams) (MessagesMarkAsImportantResponse, error) {
	r, err := v.API.Request("messages.markAsImportant", params)
	if err != nil {
		return nil, err
	}

	var resp MessagesMarkAsImportantResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// MessagesMarkAsImportantConversationParams are params for Messages.MarkAsImportantConversation
type MessagesMarkAsImportantConversationParams struct {
	// ID of conversation to mark as important.
	PeerID int `url:"peer_id"`
	// '1' — to add a star (mark as important), '0' — to remove the star
	Important bool `url:"important,omitempty"`
	// Group ID (for group messages with group access token)
	GroupID int `url:"group_id,omitempty"`
}

// MarkAsImportantConversation Marks and unmarks conversations as important.
func (v Messages) MarkAsImportantConversation(params MessagesMarkAsImportantConversationParams) (bool, error) {
	r, err := v.API.Request("messages.markAsImportantConversation", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// MessagesMarkAsAnsweredConversationParams are params for Messages.MarkAsAnsweredConversation
type MessagesMarkAsAnsweredConversationParams struct {
	// ID of conversation to mark as important.
	PeerID int `url:"peer_id"`
	// '1' — to mark as answered, '0' — to remove the mark
	Answered bool `url:"answered,omitempty"`
	// Group ID (for group messages with group access token)
	GroupID int `url:"group_id,omitempty"`
}

// MarkAsAnsweredConversation Marks and unmarks conversations as unanswered.
func (v Messages) MarkAsAnsweredConversation(params MessagesMarkAsAnsweredConversationParams) (bool, error) {
	r, err := v.API.Request("messages.markAsAnsweredConversation", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// MessagesGetLongPollServerParams are params for Messages.GetLongPollServer
type MessagesGetLongPollServerParams struct {
	// '1' — to return the 'pts' field, needed for the [vk.com/dev/messages.getLongPollHistory|messages.getLongPollHistory] method.
	NeedPts bool `url:"need_pts,omitempty"`
	// Group ID (for group messages with user access token)
	GroupID int `url:"group_id,omitempty"`
	// Long poll version
	LpVersion int `url:"lp_version,omitempty"`
}

// MessagesGetLongPollServerResponse is response for Messages.GetLongPollServer
//easyjson:json
type MessagesGetLongPollServerResponse struct {
	// Key
	Key string `json:"key,omitempty"`
	// Server URL
	Server string `json:"server,omitempty"`
	// Timestamp
	TS int `json:"ts,omitempty"`
	// Persistent timestamp
	Pts int `json:"pts,omitempty"`
}

// GetLongPollServer Returns data required for connection to a Long Poll server.
func (v Messages) GetLongPollServer(params MessagesGetLongPollServerParams) (*MessagesGetLongPollServerResponse, error) {
	r, err := v.API.Request("messages.getLongPollServer", params)
	if err != nil {
		return nil, err
	}

	var resp MessagesGetLongPollServerResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// MessagesGetLongPollHistoryParams are params for Messages.GetLongPollHistory
type MessagesGetLongPollHistoryParams struct {
	// Last value of the 'ts' parameter returned from the Long Poll server or by using [vk.com/dev/messages.getLongPollHistory|messages.getLongPollHistory] method.
	TS int `url:"ts,omitempty"`
	// Lsat value of 'pts' parameter returned from the Long Poll server or by using [vk.com/dev/messages.getLongPollHistory|messages.getLongPollHistory] method.
	Pts int `url:"pts,omitempty"`
	// Number of characters after which to truncate a previewed message. To preview the full message, specify '0'. "NOTE: Messages are not truncated by default. Messages are truncated by words."
	PreviewLength int `url:"preview_length,omitempty"`
	// '1' — to return history with online users only.
	Onlines bool `url:"onlines,omitempty"`
	// Additional profile [vk.com/dev/fields|fields] to return.
	Fields CSVStringSlice `url:"fields,omitempty"`
	// Maximum number of events to return.
	EventsLimit int `url:"events_limit,omitempty"`
	// Maximum number of messages to return.
	MsgsLimit int `url:"msgs_limit,omitempty"`
	// Maximum ID of the message among existing ones in the local copy. Both messages received with API methods (for example, , ), and data received from a Long Poll server (events with code 4) are taken into account.
	MaxMsgID int `url:"max_msg_id,omitempty"`
	// Group ID (for group messages with user access token)
	GroupID   int `url:"group_id,omitempty"`
	LpVersion int `url:"lp_version,omitempty"`
}

// MessagesGetLongPollHistoryResponse is response for Messages.GetLongPollHistory
//easyjson:json
type MessagesGetLongPollHistoryResponse struct {
	// Longpoll event value
	History  [][]int    `json:"history,omitempty"`
	Groups   []vk.Group `json:"groups,omitempty"`
	Messages struct {
		// Total number
		Count int          `json:"count,omitempty"`
		Items []vk.Message `json:"items,omitempty"`
	} `json:"messages,omitempty"`
	Profiles []vk.User `json:"profiles,omitempty"`
	Chats    []vk.Chat `json:"chats,omitempty"`
	// Persistence timestamp
	NewPts int `json:"new_pts,omitempty"`
	// Has more
	More vk.BoolInt `json:"more,omitempty"`
}

// GetLongPollHistory Returns updates in user's private messages.
func (v Messages) GetLongPollHistory(params MessagesGetLongPollHistoryParams) (*MessagesGetLongPollHistoryResponse, error) {
	r, err := v.API.Request("messages.getLongPollHistory", params)
	if err != nil {
		return nil, err
	}

	var resp MessagesGetLongPollHistoryResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// MessagesGetChatPreviewParams are params for Messages.GetChatPreview
type MessagesGetChatPreviewParams struct {
	// Invitation link.
	Link string `url:"link"`
	// Profile fields to return.
	Fields CSVStringSlice `url:"fields,omitempty"`
}

// MessagesGetChatPreviewResponse is response for Messages.GetChatPreview
//easyjson:json
type MessagesGetChatPreviewResponse struct {
	Preview struct {
		AdminID      int    `json:"admin_id,omitempty"`
		MembersCount int    `json:"members_count,omitempty"`
		Members      []int  `json:"members,omitempty"`
		Title        string `json:"title,omitempty"`
		LocalID      int    `json:"local_id,omitempty"`
		Joined       bool   `json:"joined,omitempty"`
	} `json:"preview,omitempty"`
	Profiles []vk.User `json:"profiles,omitempty"`
}

// GetChatPreview does messages.getChatPreview
func (v Messages) GetChatPreview(params MessagesGetChatPreviewParams) (*MessagesGetChatPreviewResponse, error) {
	r, err := v.API.Request("messages.getChatPreview", params)
	if err != nil {
		return nil, err
	}

	var resp MessagesGetChatPreviewResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// MessagesCreateChatParams are params for Messages.CreateChat
type MessagesCreateChatParams struct {
	// IDs of the users to be added to the chat.
	UserIDs CSVIntSlice `url:"user_ids,omitempty"`
	// Chat title.
	Title string `url:"title,omitempty"`
}

// MessagesCreateChatResponse is response for Messages.CreateChat
// Chat ID
type MessagesCreateChatResponse int

// CreateChat Creates a chat with several participants.
func (v Messages) CreateChat(params MessagesCreateChatParams) (MessagesCreateChatResponse, error) {
	r, err := v.API.Request("messages.createChat", params)
	if err != nil {
		return 0, err
	}

	var resp MessagesCreateChatResponse

	var cnv int
	cnv, err = strconv.Atoi(string(r))
	resp = MessagesCreateChatResponse(cnv)

	if err != nil {
		return 0, err
	}
	return resp, nil
}

// MessagesEditChatParams are params for Messages.EditChat
type MessagesEditChatParams struct {
	// Chat ID.
	ChatID int `url:"chat_id"`
	// New title of the chat.
	Title string `url:"title"`
}

// EditChat Edits the title of a chat.
func (v Messages) EditChat(params MessagesEditChatParams) (bool, error) {
	r, err := v.API.Request("messages.editChat", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// MessagesGetConversationMembersParams are params for Messages.GetConversationMembers
type MessagesGetConversationMembersParams struct {
	// Peer ID.
	PeerID int `url:"peer_id"`
	// Profile fields to return.
	Fields CSVStringSlice `url:"fields,omitempty"`
	// Group ID (for group messages with group access token)
	GroupID int `url:"group_id,omitempty"`
}

// MessagesGetConversationMembersResponse is response for Messages.GetConversationMembers
//easyjson:json
type MessagesGetConversationMembersResponse struct {
	// Chat members count
	Count int `json:"count,omitempty"`
	Items []struct {
		MemberID  int  `json:"member_id,omitempty"`
		JoinDate  int  `json:"join_date,omitempty"`
		InvitedBy int  `json:"invited_by,omitempty"`
		IsOwner   bool `json:"is_owner,omitempty"`
		IsAdmin   bool `json:"is_admin,omitempty"`
		// Is it possible for user to kick this member
		CanKick bool `json:"can_kick,omitempty"`
	} `json:"items,omitempty"`
	ChatRestrictions struct {
		// Only admins can invite users to this chat
		OnlyAdminsInvite bool `json:"only_admins_invite,omitempty"`
		// Only admins can kick users from this chat
		OnlyAdminsKick bool `json:"only_admins_kick,omitempty"`
		// Only admins can change chat info
		OnlyAdminsEditInfo bool `json:"only_admins_edit_info,omitempty"`
		// Only admins can edit pinned message
		OnlyAdminsEditPin bool `json:"only_admins_edit_pin,omitempty"`
		// Only admins can promote users to admins
		AdminsPromoteUsers bool `json:"admins_promote_users,omitempty"`
	} `json:"chat_restrictions,omitempty"`
	Profiles []vk.User  `json:"profiles,omitempty"`
	Groups   []vk.Group `json:"groups,omitempty"`
}

// GetConversationMembers Returns a list of IDs of users participating in a chat.
func (v Messages) GetConversationMembers(params MessagesGetConversationMembersParams) (*MessagesGetConversationMembersResponse, error) {
	r, err := v.API.Request("messages.getConversationMembers", params)
	if err != nil {
		return nil, err
	}

	var resp MessagesGetConversationMembersResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// MessagesSetActivityParams are params for Messages.SetActivity
type MessagesSetActivityParams struct {
	// User ID.
	UserID int `url:"user_id,omitempty"`
	// 'typing' — user has started to type.
	Type string `url:"type,omitempty"`
	// Destination ID. "For user: 'User ID', e.g. '12345'. For chat: '2000000000' + 'chat_id', e.g. '2000000001'. For community: '- community ID', e.g. '-12345'. "
	PeerID int `url:"peer_id,omitempty"`
	// Group ID (for group messages with group access token)
	GroupID int `url:"group_id,omitempty"`
}

// SetActivity Changes the status of a user as typing in a conversation.
func (v Messages) SetActivity(params MessagesSetActivityParams) (bool, error) {
	r, err := v.API.Request("messages.setActivity", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// MessagesSearchConversationsParams are params for Messages.SearchConversations
type MessagesSearchConversationsParams struct {
	// Search query string.
	Q string `url:"q,omitempty"`
	// Maximum number of results.
	Count int `url:"count,omitempty"`
	// '1' — return extra information about users and communities
	Extended bool `url:"extended,omitempty"`
	// Profile fields to return.
	Fields CSVStringSlice `url:"fields,omitempty"`
	// Group ID (for group messages with user access token)
	GroupID int `url:"group_id,omitempty"`
}

// MessagesSearchConversationsResponse is response for Messages.SearchConversations
//easyjson:json
type MessagesSearchConversationsResponse struct {
	// Total results number
	Count    int               `json:"count,omitempty"`
	Items    []vk.Conversation `json:"items,omitempty"`
	Profiles []vk.User         `json:"profiles,omitempty"`
	Groups   []vk.Group        `json:"groups,omitempty"`
}

// SearchConversations Returns a list of the current user's conversations that match search criteria.
func (v Messages) SearchConversations(params MessagesSearchConversationsParams) (*MessagesSearchConversationsResponse, error) {
	r, err := v.API.Request("messages.searchConversations", params)
	if err != nil {
		return nil, err
	}

	var resp MessagesSearchConversationsResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// MessagesAddChatUserParams are params for Messages.AddChatUser
type MessagesAddChatUserParams struct {
	// Chat ID.
	ChatID int `url:"chat_id"`
	// ID of the user to be added to the chat.
	UserID int `url:"user_id"`
}

// AddChatUser Adds a new user to a chat.
func (v Messages) AddChatUser(params MessagesAddChatUserParams) (bool, error) {
	r, err := v.API.Request("messages.addChatUser", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// MessagesRemoveChatUserParams are params for Messages.RemoveChatUser
type MessagesRemoveChatUserParams struct {
	// Chat ID.
	ChatID int `url:"chat_id"`
	// ID of the user to be removed from the chat.
	UserID   int `url:"user_id,omitempty"`
	MemberID int `url:"member_id,omitempty"`
}

// RemoveChatUser Allows the current user to leave a chat or, if the current user started the chat, allows the user to remove another user from the chat.
func (v Messages) RemoveChatUser(params MessagesRemoveChatUserParams) (bool, error) {
	r, err := v.API.Request("messages.removeChatUser", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// MessagesGetLastActivityParams are params for Messages.GetLastActivity
type MessagesGetLastActivityParams struct {
	// User ID.
	UserID int `url:"user_id"`
}

// MessagesGetLastActivityResponse is response for Messages.GetLastActivity
//easyjson:json
type MessagesGetLastActivityResponse struct {
	// Information whether user is online
	Online vk.BoolInt `json:"online,omitempty"`
	// Time when user was online in Unixtime
	Time int `json:"time,omitempty"`
}

// GetLastActivity Returns a user's current status and date of last activity.
func (v Messages) GetLastActivity(params MessagesGetLastActivityParams) (*MessagesGetLastActivityResponse, error) {
	r, err := v.API.Request("messages.getLastActivity", params)
	if err != nil {
		return nil, err
	}

	var resp MessagesGetLastActivityResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// MessagesSetChatPhotoParams are params for Messages.SetChatPhoto
type MessagesSetChatPhotoParams struct {
	// Upload URL from the 'response' field returned by the [vk.com/dev/photos.getChatUploadServer|photos.getChatUploadServer] method upon successfully uploading an image.
	File string `url:"file"`
}

// MessagesSetChatPhotoResponse is response for Messages.SetChatPhoto
//easyjson:json
type MessagesSetChatPhotoResponse struct {
	// Service message ID
	MessageID int     `json:"message_id,omitempty"`
	Chat      vk.Chat `json:"chat,omitempty"`
}

// SetChatPhoto Sets a previously-uploaded picture as the cover picture of a chat.
func (v Messages) SetChatPhoto(params MessagesSetChatPhotoParams) (*MessagesSetChatPhotoResponse, error) {
	r, err := v.API.Request("messages.setChatPhoto", params)
	if err != nil {
		return nil, err
	}

	var resp MessagesSetChatPhotoResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// MessagesDeleteChatPhotoParams are params for Messages.DeleteChatPhoto
type MessagesDeleteChatPhotoParams struct {
	// Chat ID.
	ChatID  int `url:"chat_id"`
	GroupID int `url:"group_id,omitempty"`
}

// MessagesDeleteChatPhotoResponse is response for Messages.DeleteChatPhoto
//easyjson:json
type MessagesDeleteChatPhotoResponse struct {
	// Service message ID
	MessageID int     `json:"message_id,omitempty"`
	Chat      vk.Chat `json:"chat,omitempty"`
}

// DeleteChatPhoto Deletes a chat's cover picture.
func (v Messages) DeleteChatPhoto(params MessagesDeleteChatPhotoParams) (*MessagesDeleteChatPhotoResponse, error) {
	r, err := v.API.Request("messages.deleteChatPhoto", params)
	if err != nil {
		return nil, err
	}

	var resp MessagesDeleteChatPhotoResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// MessagesDenyMessagesFromGroupParams are params for Messages.DenyMessagesFromGroup
type MessagesDenyMessagesFromGroupParams struct {
	// Group ID.
	GroupID int `url:"group_id"`
}

// DenyMessagesFromGroup Denies sending message from community to the current user.
func (v Messages) DenyMessagesFromGroup(params MessagesDenyMessagesFromGroupParams) (bool, error) {
	r, err := v.API.Request("messages.denyMessagesFromGroup", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// MessagesAllowMessagesFromGroupParams are params for Messages.AllowMessagesFromGroup
type MessagesAllowMessagesFromGroupParams struct {
	// Group ID.
	GroupID int    `url:"group_id"`
	Key     string `url:"key,omitempty"`
}

// AllowMessagesFromGroup Allows sending messages from community to the current user.
func (v Messages) AllowMessagesFromGroup(params MessagesAllowMessagesFromGroupParams) (bool, error) {
	r, err := v.API.Request("messages.allowMessagesFromGroup", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// MessagesIsMessagesFromGroupAllowedParams are params for Messages.IsMessagesFromGroupAllowed
type MessagesIsMessagesFromGroupAllowedParams struct {
	// Group ID.
	GroupID int `url:"group_id"`
	// User ID.
	UserID int `url:"user_id"`
}

// MessagesIsMessagesFromGroupAllowedResponse is response for Messages.IsMessagesFromGroupAllowed
//easyjson:json
type MessagesIsMessagesFromGroupAllowedResponse struct {
	IsAllowed vk.BoolInt `json:"is_allowed,omitempty"`
}

// IsMessagesFromGroupAllowed Returns information whether sending messages from the community to current user is allowed.
func (v Messages) IsMessagesFromGroupAllowed(params MessagesIsMessagesFromGroupAllowedParams) (*MessagesIsMessagesFromGroupAllowedResponse, error) {
	r, err := v.API.Request("messages.isMessagesFromGroupAllowed", params)
	if err != nil {
		return nil, err
	}

	var resp MessagesIsMessagesFromGroupAllowedResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// MessagesUnpinParams are params for Messages.Unpin
type MessagesUnpinParams struct {
	PeerID  int `url:"peer_id"`
	GroupID int `url:"group_id,omitempty"`
}

// Unpin does messages.unpin
func (v Messages) Unpin(params MessagesUnpinParams) (bool, error) {
	r, err := v.API.Request("messages.unpin", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}
