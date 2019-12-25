package vk

import (
	"encoding/json"
	"fmt"
)

// CallbackEvent is base event
type CallbackEvent struct {
	// ID of group this event occured in
	GroupID int
	// Secret for Callback API
	Secret string
	// Event itself
	//
	// One of Confirmation, MessageNew, MessageReply, MessageEdit,
	// MessageAllow, MessageDeny, MessageTypingState, PhotoNew,
	// PhotoCommentNew, PhotoCommentEdit, PhotoCommentRestore,
	// PhotoCommentDelete, AudioNew, VideoNew, VideoCommentNew,
	// VideoCommentEdit, VideoCommentRestore, VideoCommentDelete,
	// WallPostNew, WallRepost, WallReplyNew, WallReplyEdit,
	// WallReplyRestore, WallReplyDelete, BoardPostNew, BoardPostEdit,
	// BoardPostRestore, BoardPostDelete, MarketCommentNew,
	// MarketCommentEdit, MarketCommentRestore, MarketCommentDelete,
	// GroupLeave, GroupJoin, UserBlock, UserUnblock, PollVoteNew,
	// GroupOfficersEdit, GroupChangeSettings, GroupChangePhoto,
	// LeadFormsNew, NewVKPayTransaction.
	Event interface{}
}

// UnmarshalJSON implements json.Unmarshaler interface
func (e *CallbackEvent) UnmarshalJSON(data []byte) error {
	var rawEvent struct {
		GroupID int             `json:"group_id"`
		Secret  string          `json:"secret"`
		Type    string          `json:"type"`
		Object  json.RawMessage `json:"object"`
	}

	err := json.Unmarshal(data, &rawEvent)
	if err != nil {
		return err
	}

	e.GroupID = rawEvent.GroupID
	e.Secret = rawEvent.Secret

	switch rawEvent.Type {
	case "confirmation":
		// confirmation has no object
		e.Event = Confirmation{}
		err = nil
	case "message_new":
		evt := MessageNew{}
		err = json.Unmarshal(rawEvent.Object, &evt)
		e.Event = evt
	case "message_reply":
		evt := MessageReply{}
		err = json.Unmarshal(rawEvent.Object, &evt)
		e.Event = evt
	case "message_edit":
		evt := MessageEdit{}
		err = json.Unmarshal(rawEvent.Object, &evt)
		e.Event = evt
	case "message_typing_state":
		evt := MessageTypingState{}
		err = json.Unmarshal(rawEvent.Object, &evt)
		e.Event = evt
	case "message_allow":
		evt := MessageAllow{}
		err = json.Unmarshal(rawEvent.Object, &evt)
		e.Event = evt
	case "message_deny":
		evt := MessageDeny{}
		err = json.Unmarshal(rawEvent.Object, &evt)
		e.Event = evt
	case "photo_new":
		evt := PhotoNew{}
		err = json.Unmarshal(rawEvent.Object, &evt)
		e.Event = evt
	case "photo_comment_new":
		evt := PhotoCommentNew{}
		err = json.Unmarshal(rawEvent.Object, &evt)
		e.Event = evt
	case "photo_comment_edit":
		evt := PhotoCommentEdit{}
		err = json.Unmarshal(rawEvent.Object, &evt)
		e.Event = evt
	case "photo_comment_restore":
		evt := PhotoCommentRestore{}
		err = json.Unmarshal(rawEvent.Object, &evt)
		e.Event = evt
	case "photo_comment_delete":
		evt := PhotoCommentDelete{}
		err = json.Unmarshal(rawEvent.Object, &evt)
		e.Event = evt
	case "audio_new":
		evt := AudioNew{}
		err = json.Unmarshal(rawEvent.Object, &evt)
		e.Event = evt
	case "video_new":
		evt := VideoNew{}
		err = json.Unmarshal(rawEvent.Object, &evt)
		e.Event = evt
	case "video_comment_new":
		evt := VideoCommentNew{}
		err = json.Unmarshal(rawEvent.Object, &evt)
		e.Event = evt
	case "video_comment_edit":
		evt := VideoCommentEdit{}
		err = json.Unmarshal(rawEvent.Object, &evt)
		e.Event = evt
	case "video_comment_restore":
		evt := VideoCommentRestore{}
		err = json.Unmarshal(rawEvent.Object, &evt)
		e.Event = evt
	case "video_comment_delete":
		evt := VideoCommentDelete{}
		err = json.Unmarshal(rawEvent.Object, &evt)
		e.Event = evt
	case "wall_post_new":
		evt := WallPostNew{}
		err = json.Unmarshal(rawEvent.Object, &evt)
		e.Event = evt
	case "wall_repost":
		evt := WallRepost{}
		err = json.Unmarshal(rawEvent.Object, &evt)
		e.Event = evt
	case "wall_reply_new":
		evt := WallReplyNew{}
		err = json.Unmarshal(rawEvent.Object, &evt)
		e.Event = evt
	case "wall_reply_edit":
		evt := WallReplyEdit{}
		err = json.Unmarshal(rawEvent.Object, &evt)
		e.Event = evt
	case "wall_reply_restore":
		evt := WallReplyRestore{}
		err = json.Unmarshal(rawEvent.Object, &evt)
		e.Event = evt
	case "wall_reply_delete":
		evt := WallReplyDelete{}
		err = json.Unmarshal(rawEvent.Object, &evt)
		e.Event = evt
	case "board_post_new":
		evt := BoardPostNew{}
		err = json.Unmarshal(rawEvent.Object, &evt)
		e.Event = evt
	case "board_post_edit":
		evt := BoardPostEdit{}
		err = json.Unmarshal(rawEvent.Object, &evt)
		e.Event = evt
	case "board_post_restore":
		evt := BoardPostRestore{}
		err = json.Unmarshal(rawEvent.Object, &evt)
		e.Event = evt
	case "board_post_delete":
		evt := BoardPostDelete{}
		err = json.Unmarshal(rawEvent.Object, &evt)
		e.Event = evt
	case "market_comment_new":
		evt := MarketCommentNew{}
		err = json.Unmarshal(rawEvent.Object, &evt)
		e.Event = evt
	case "market_comment_edit":
		evt := MarketCommentEdit{}
		err = json.Unmarshal(rawEvent.Object, &evt)
		e.Event = evt
	case "market_comment_restore":
		evt := MarketCommentRestore{}
		err = json.Unmarshal(rawEvent.Object, &evt)
		e.Event = evt
	case "market_comment_delete":
		evt := MarketCommentDelete{}
		err = json.Unmarshal(rawEvent.Object, &evt)
		e.Event = evt
	case "group_leave":
		evt := GroupLeave{}
		err = json.Unmarshal(rawEvent.Object, &evt)
		e.Event = evt
	case "group_join":
		evt := GroupJoin{}
		err = json.Unmarshal(rawEvent.Object, &evt)
		e.Event = evt
	case "user_block":
		evt := UserBlock{}
		err = json.Unmarshal(rawEvent.Object, &evt)
		e.Event = evt
	case "user_unblock":
		evt := UserUnblock{}
		err = json.Unmarshal(rawEvent.Object, &evt)
		e.Event = evt
	case "poll_vote_new":
		evt := PollVoteNew{}
		err = json.Unmarshal(rawEvent.Object, &evt)
		e.Event = evt
	case "group_officers_edit":
		evt := GroupOfficersEdit{}
		err = json.Unmarshal(rawEvent.Object, &evt)
		e.Event = evt
	case "group_change_settings":
		evt := GroupChangeSettings{}
		err = json.Unmarshal(rawEvent.Object, &evt)
		e.Event = evt
	case "group_change_photo":
		evt := GroupChangePhoto{}
		err = json.Unmarshal(rawEvent.Object, &evt)
		e.Event = evt
	case "lead_forms_new":
		evt := LeadFormsNew{}
		err = json.Unmarshal(rawEvent.Object, &evt)
		e.Event = evt
	case "vkpay_transaction":
		evt := NewVKPayTransaction{}
		err = json.Unmarshal(rawEvent.Object, &evt)
		e.Event = evt
	default:
		return fmt.Errorf("Unknown event type: %v", rawEvent.Type)
	}

	return err
}

// Confirmation is used in Callback API.
// It requires listener to reply with Confirmation token instead of normal "ok".
//
//easyjson:json
type Confirmation struct{}

// MessageNew -- new message is recieved
//
//easyjson:json
type MessageNew struct {
	Message
}

// MessageReply -- new message is sent
//
//easyjson:json
type MessageReply struct {
	Message
}

// MessageEdit -- a message is edited
//
//easyjson:json
type MessageEdit struct {
	Message
}

// MessageAllow -- new user consent to messages sending
//
//easyjson:json
type MessageAllow struct {
	UserID int `json:"user_id"`
	// Key is parameter from messages.allowMessagesFromGroup
	Key string `json:"key"`
}

// MessageDeny -- new user prohibition to messages sending
//
//easyjson:json
type MessageDeny struct {
	UserID int `json:"user_id"`
}

// MessageTypingState -- new message typing state
//
//easyjson:json
type MessageTypingState struct {
	// State is always "typing" (XXX: Ask VK devs)
	State string `json:"state"`
	// FromID of peer who's typing
	FromID int `json:"from_id"`
	// ToID of other peer
	ToID int `json:"to_id"`
}

// PhotoNew -- new photo in community
//
//easyjson:json
type PhotoNew struct {
	Photo
}

// PhotoCommentNew -- new photo comment
//
//easyjson:json
type PhotoCommentNew struct {
	Comment

	PhotoID      int `json:"photo_id"`
	PhotoOwnerID int `json:"photo_owner_id"`
}

// PhotoCommentEdit -- photo comment edited
//
//easyjson:json
type PhotoCommentEdit struct {
	Comment

	PhotoID      int `json:"photo_id"`
	PhotoOwnerID int `json:"photo_owner_id"`
}

// PhotoCommentRestore -- photo comment restored
//
//easyjson:json
type PhotoCommentRestore struct {
	Comment

	PhotoID      int `json:"photo_id"`
	PhotoOwnerID int `json:"photo_owner_id"`
}

// PhotoCommentDelete -- photo comment deleted
//
//easyjson:json
type PhotoCommentDelete struct {
	OwnerID   int `json:"owner_id"`
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	DeleterID int `json:"deleter_id"`
	PhotoID   int `json:"photo_id"`
}

// AudioNew -- new audio in community
//
//easyjson:json
type AudioNew struct {
	Audio
}

// VideoNew -- new video in community
//
//easyjson:json
type VideoNew struct {
	Video
}

// VideoCommentNew -- new video comment
//
//easyjson:json
type VideoCommentNew struct {
	Comment

	VideoID      int `json:"video_id"`
	VideoOwnerID int `json:"video_owner_id"`
}

// VideoCommentEdit -- video comment edited
//
//easyjson:json
type VideoCommentEdit struct {
	Comment

	VideoID      int `json:"video_id"`
	VideoOwnerID int `json:"video_owner_id"`
}

// VideoCommentRestore -- video comment restored
//
//easyjson:json
type VideoCommentRestore struct {
	Comment

	VideoID      int `json:"video_id"`
	VideoOwnerID int `json:"video_owner_id"`
}

// VideoCommentDelete -- video comment deleted
//
//easyjson:json
type VideoCommentDelete struct {
	OwnerID   int `json:"owner_id"`
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	DeleterID int `json:"deleter_id"`
	VideoID   int `json:"video_id"`
}

// WallPostNew -- new post on community wall
//
//easyjson:json
type WallPostNew struct {
	Post

	PostponedID int `json:"postponed_id"`
}

// WallRepost -- new repost
// (XXX: repost of community post or repost on community wall?!)
//
//easyjson:json
type WallRepost struct {
	Post

	PostponedID int `json:"postponed_id"`
}

// WallReplyNew -- new wall comment
//
//easyjson:json
type WallReplyNew struct {
	Comment

	PostID      int `json:"post_id"`
	PostOwnerID int `json:"post_owner_id"`
}

// WallReplyEdit -- wall comment edited
//
//easyjson:json
type WallReplyEdit struct {
	Comment

	PostID      int `json:"post_id"`
	PostOwnerID int `json:"post_owner_id"`
}

// WallReplyRestore -- wall comment restored
//
//easyjson:json
type WallReplyRestore struct {
	Comment

	PostID      int `json:"post_id"`
	PostOwnerID int `json:"post_owner_id"`
}

// WallReplyDelete -- wall comment deleted
//
//easyjson:json
type WallReplyDelete struct {
	OwnerID   int `json:"owner_id"`
	ID        int `json:"id"`
	DeleterID int `json:"deleter_id"`
	PostID    int `json:"post_id"`
}

// BoardPostNew -- new board comment
//
//easyjson:json
type BoardPostNew struct {
	CommentBoard

	TopicID      int `json:"topic_id"`
	TopicOwnerID int `json:"topic_owner_id"`
}

// BoardPostEdit -- board comment edited
//
//easyjson:json
type BoardPostEdit struct {
	CommentBoard

	TopicID      int `json:"topic_id"`
	TopicOwnerID int `json:"topic_owner_id"`
}

// BoardPostRestore -- board comment restored
//
//easyjson:json
type BoardPostRestore struct {
	CommentBoard

	TopicID      int `json:"topic_id"`
	TopicOwnerID int `json:"topic_owner_id"`
}

// BoardPostDelete -- board comment deleted
//
//easyjson:json
type BoardPostDelete struct {
	TopicOwnerID int `json:"topic_owner_id"`
	TopicID      int `json:"topic_id"`
	ID           int `json:"id"`
}

// MarketCommentNew -- new market comment
//
//easyjson:json
type MarketCommentNew struct {
	Comment

	MarketOwnerID int `json:"market_owner_id"`
	ItemID        int `json:"item_id"`
}

// MarketCommentEdit -- market comment edited
//
//easyjson:json
type MarketCommentEdit struct {
	Comment

	MarketOwnerID int `json:"market_owner_id"`
	ItemID        int `json:"item_id"`
}

// MarketCommentRestore -- marked comment restored
//
//easyjson:json
type MarketCommentRestore struct {
	Comment

	MarketOwnerID int `json:"market_owner_id"`
	ItemID        int `json:"item_id"`
}

// MarketCommentDelete -- market comment deleted
//
//easyjson:json
type MarketCommentDelete struct {
	OwnerID   int `json:"owner_id"`
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	DeleterID int `json:"deleter_id"`
	ItemID    int `json:"item_id"`
}

// GroupLeave -- member removed from community
//
//easyjson:json
type GroupLeave struct {
	// UserID of user who has left
	UserID int `json:"user_id"`
	// Self is true if user has left on its own and false if user was kicked
	Self BoolInt `json:"self"`
}

// GroupJoinType is type for GroupJoin event
type GroupJoinType string

const (
	// GroupJoinTypeJoin -- User joined a group or event (subscribed a public page)
	GroupJoinTypeJoin GroupJoinType = "join"
	// GroupJoinTypeUnsure -- For events: user has chosen "I may attend"
	GroupJoinTypeUnsure = "unsure"
	// GroupJoinTypeAccepted -- User approved an invitation to a group or event
	GroupJoinTypeAccepted = "accepted"
	// GroupJoinTypeApproved -- Join request was approved by community administrators
	GroupJoinTypeApproved = "approved"
	// GroupJoinTypeRequest -- User sent a join request
	GroupJoinTypeRequest = "request"
)

// GroupJoin -- member added to community
//
//easyjson:json
type GroupJoin struct {
	// UserID who has joined
	UserID int `json:"user_id"`
	// JoinType is how user got into group
	JoinType GroupJoinType `json:"join_type"`
}

// UserBlockReason is the reason why user was blocked
type UserBlockReason int

const (
	// UserBlockReasonOther is default reason
	UserBlockReasonOther UserBlockReason = 0
	// UserBlockReasonSpam -- for spam
	UserBlockReasonSpam = 1
	// UserBlockReasonVerbalAbuse -- for verbal abuse
	UserBlockReasonVerbalAbuse = 2
	// UserBlockReasonStrongLanguage -- for strong language
	UserBlockReasonStrongLanguage = 3
	// UserBlockReasonIrrelevantMessages -- for irrelevant messages
	UserBlockReasonIrrelevantMessages = 4
)

// UserBlock -- new user in blacklist
//
//easyjson:json
type UserBlock struct {
	// AdminID of admin who has blocked user
	AdminID int `json:"admin_id"`
	// UserID who was unlocked
	UserID int `json:"user_id"`
	// UnblockDate when the user will be unblocked
	UnblockDate int `json:"unblock_date"`
	// Reason of block
	Reason UserBlockReason `json:"reason"`
	// Comment attached to block
	Comment string `json:"comment"`
}

// UserUnblock -- user has been removed from the blacklist
//
//easyjson:json
type UserUnblock struct {
	// AdminID of admin who has unblocked user
	AdminID int `json:"admin_id"`
	// UserID who was unlocked
	UserID int `json:"user_id"`
	// ByEndDate is true if the ban has expired
	ByEndDate BoolInt `json:"by_end_date"`
}

// PollVoteNew -- new vote in a public poll
//
//easyjson:json
type PollVoteNew struct {
	// OwnerID of poll
	OwnerID int `json:"owner_id"`
	// PollID of poll
	PollID int `json:"poll_id"`
	// OptionID of option in poll
	OptionID int `json:"option_id"`
	// UserID of user who has voted
	UserID int `json:"user_id"`
}

// GroupOfficerRole is role of group admin
type GroupOfficerRole int

const (
	// GroupOfficerRoleNone -- No role = normal user
	GroupOfficerRoleNone GroupOfficerRole = 0
	// GroupOfficerRoleModerator is moderator
	GroupOfficerRoleModerator = 1
	// GroupOfficerRoleEditor is editor
	GroupOfficerRoleEditor = 2
	// GroupOfficerRoleAdministrator is administrator
	GroupOfficerRoleAdministrator = 3
)

// GroupOfficersEdit -- changes in the administrators list
//
//easyjson:json
type GroupOfficersEdit struct {
	// AdminID of administrator who made changes
	AdminID int `json:"admin_id"`
	// UserID of whose role was changed
	UserID int `json:"user_id"`
	// LevelOld is old role
	LevelOld GroupOfficerRole `json:"level_old"`
	// LevelNew is new role
	LevelNew GroupOfficerRole `json:"level_new"`
}

// ChangedStringValue -- holds old and new string values
//
//easyjson:json
type ChangedStringValue struct {
	OldValue string `json:"old_value"`
	NewValue string `json:"new_value"`
}

// ChangedIntValue -- holds old and new int values
//
//easyjson:json
type ChangedIntValue struct {
	OldValue int `json:"old_value"`
	NewValue int `json:"new_value"`
}

// GroupChangeSettings -- changes in community settings
//
//easyjson:json
type GroupChangeSettings struct {
	// UserID of user who made changes
	UserID  int `json:"user_id"`
	Changes struct {
		Title             *ChangedStringValue `json:"title"`
		Description       *ChangedStringValue `json:"description"`
		Access            *ChangedStringValue `json:"access"`
		ScreenName        *ChangedStringValue `json:"screen_name"`
		PublicCategory    ChangedIntValue     `json:"public_category"`
		PublicSubcategory ChangedIntValue     `json:"public_subcategory"`
		Website           *ChangedStringValue `json:"website"`

		// 0=None, 1=0-16, 2=16+, 3=18+
		AgeLimits ChangedIntValue `json:"age_limits"`
		// 0=No one/Disabled, 1=All members/Everyone, 2=Community only
		Audio  ChangedIntValue `json:"audio"`
		Photo  ChangedIntValue `json:"photo"`
		Video  ChangedIntValue `json:"video"`
		Market ChangedIntValue `json:"market"`
		Docs   ChangedIntValue `json:"docs"`
		// Comments on wall
		Replies ChangedIntValue `json:"replies"`
		// Wall posts?..
		StatusDefault ChangedIntValue `json:"status_default"`
	} `json:"changes"`
}

// GroupChangePhoto -- changes of community main photo
//
//easyjson:json
type GroupChangePhoto struct {
	// UserID of user who changed photo
	UserID int `json:"user_id"`
	// Photo new photo
	Photo Photo `json:"photo"`
}

// LeadFormsNew -- new lead forms filled
// TODO: Find definition
//
//easyjson:json
type LeadFormsNew struct {
}

// NewVKPayTransaction -- new VKPay transaction
// TODO: Find definition
//
// Starts with New to silence golint:
//    type name will be used as vk.VKPayTransaction by other packages,
//    and that stutters; consider calling this PayTransaction
//
//easyjson:json
type NewVKPayTransaction struct {
}
