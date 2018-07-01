package vkCallbackApi

import (
	"github.com/mailru/easyjson/jlexer"
)

type CallbackEvent struct {
	GroupId int
	Secret  string
	Event   interface{}
}

func (v *CallbackEvent) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	v.UnmarshalEasyJSON(&r)
	return r.Error()
}

func (v *CallbackEvent) UnmarshalEasyJSON(in *jlexer.Lexer) {
	in.Delim('{')
	var vType string

	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()

		switch key {
		case "group_id":
			v.GroupId = in.Int()
		case "secret":
			v.Secret = in.String()
		case "type":
			vType = in.UnsafeString()
		case "object":
			switch vType {
			// there's no object in "confirmation"
			case "message_new":
				tmp := MessageNew{}
				tmp.UnmarshalEasyJSON(in)
				v.Event = tmp

			case "message_reply":
				tmp := MessageReply{}
				tmp.UnmarshalEasyJSON(in)
				v.Event = tmp

			case "message_edit":
				tmp := MessageEdit{}
				tmp.UnmarshalEasyJSON(in)
				v.Event = tmp

			case "message_allow":
				tmp := MessageAllow{}
				tmp.UnmarshalEasyJSON(in)
				v.Event = tmp

			case "message_deny":
				tmp := MessageDeny{}
				tmp.UnmarshalEasyJSON(in)
				v.Event = tmp

			case "photo_new":
				tmp := PhotoNew{}
				tmp.UnmarshalEasyJSON(in)
				v.Event = tmp

			case "photo_comment_new":
				tmp := PhotoCommentNew{}
				tmp.UnmarshalEasyJSON(in)
				v.Event = tmp

			case "photo_comment_edit":
				tmp := PhotoCommentEdit{}
				tmp.UnmarshalEasyJSON(in)
				v.Event = tmp

			case "photo_comment_restore":
				tmp := PhotoCommentRestore{}
				tmp.UnmarshalEasyJSON(in)
				v.Event = tmp

			case "photo_comment_delete":
				tmp := PhotoCommentDelete{}
				tmp.UnmarshalEasyJSON(in)
				v.Event = tmp

			case "audio_new":
				tmp := AudioNew{}
				tmp.UnmarshalEasyJSON(in)
				v.Event = tmp

			case "video_new":
				tmp := VideoNew{}
				tmp.UnmarshalEasyJSON(in)
				v.Event = tmp

			case "video_comment_new":
				tmp := VideoCommentNew{}
				tmp.UnmarshalEasyJSON(in)
				v.Event = tmp

			case "video_comment_edit":
				tmp := VideoCommentEdit{}
				tmp.UnmarshalEasyJSON(in)
				v.Event = tmp

			case "video_comment_restore":
				tmp := VideoCommentRestore{}
				tmp.UnmarshalEasyJSON(in)
				v.Event = tmp

			case "video_comment_delete":
				tmp := VideoCommentDelete{}
				tmp.UnmarshalEasyJSON(in)
				v.Event = tmp

			case "wall_post_new":
				tmp := WallPostNew{}
				tmp.UnmarshalEasyJSON(in)
				v.Event = tmp

			case "wall_repost":
				tmp := WallRepost{}
				tmp.UnmarshalEasyJSON(in)
				v.Event = tmp

			case "wall_reply_new":
				tmp := WallReplyNew{}
				tmp.UnmarshalEasyJSON(in)
				v.Event = tmp

			case "wall_reply_edit":
				tmp := WallReplyEdit{}
				tmp.UnmarshalEasyJSON(in)
				v.Event = tmp

			case "wall_reply_restore":
				tmp := WallReplyRestore{}
				tmp.UnmarshalEasyJSON(in)
				v.Event = tmp

			case "wall_reply_delete":
				tmp := WallReplyDelete{}
				tmp.UnmarshalEasyJSON(in)
				v.Event = tmp

			case "board_post_new":
				tmp := BoardPostNew{}
				tmp.UnmarshalEasyJSON(in)
				v.Event = tmp

			case "board_post_edit":
				tmp := BoardPostEdit{}
				tmp.UnmarshalEasyJSON(in)
				v.Event = tmp

			case "board_post_restore":
				tmp := BoardPostRestore{}
				tmp.UnmarshalEasyJSON(in)
				v.Event = tmp

			case "board_post_delete":
				tmp := BoardPostDelete{}
				tmp.UnmarshalEasyJSON(in)
				v.Event = tmp

			case "market_comment_new":
				tmp := MarketCommentNew{}
				tmp.UnmarshalEasyJSON(in)
				v.Event = tmp

			case "market_comment_edit":
				tmp := MarketCommentEdit{}
				tmp.UnmarshalEasyJSON(in)
				v.Event = tmp

			case "market_comment_restore":
				tmp := MarketCommentRestore{}
				tmp.UnmarshalEasyJSON(in)
				v.Event = tmp

			case "market_comment_delete":
				tmp := MarketCommentDelete{}
				tmp.UnmarshalEasyJSON(in)
				v.Event = tmp

			case "group_leave":
				tmp := GroupLeave{}
				tmp.UnmarshalEasyJSON(in)
				v.Event = tmp

			case "group_join":
				tmp := GroupJoin{}
				tmp.UnmarshalEasyJSON(in)
				v.Event = tmp

			case "user_block":
				tmp := UserBlock{}
				tmp.UnmarshalEasyJSON(in)
				v.Event = tmp

			case "user_unblock":
				tmp := UserUnblock{}
				tmp.UnmarshalEasyJSON(in)
				v.Event = tmp

			case "poll_vote_new":
				tmp := PollVoteNew{}
				tmp.UnmarshalEasyJSON(in)
				v.Event = tmp

			case "group_officers_edit":
				tmp := GroupOfficersEdit{}
				tmp.UnmarshalEasyJSON(in)
				v.Event = tmp

			case "group_change_settings":
				tmp := GroupChangeSettings{}
				tmp.UnmarshalEasyJSON(in)
				v.Event = tmp

			case "group_change_photo":
				tmp := GroupChangePhoto{}
				tmp.UnmarshalEasyJSON(in)
				v.Event = tmp

			default:
				in.SkipRecursive()
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}

	if vType == "confirmation" {
		v.Event = Confirmation{}
	}

	in.Delim('}')
}

//easyjson:json
type Confirmation struct {
}

//easyjson:json
type MessageNew struct {
	Message
}

//easyjson:json
type MessageReply struct {
	Message
}

//easyjson:json
type MessageEdit struct {
	Message
}

//easyjson:json
type MessageAllow struct {
	UserId int
	Key    string
}

//easyjson:json
type MessageDeny struct {
	UserId int
}

//easyjson:json
type PhotoNew struct {
	Photo
}

//easyjson:json
type PhotoCommentNew struct {
	Comment

	PhotoId      int
	PhotoOwnerId int
}

//easyjson:json
type PhotoCommentEdit struct {
	Comment

	PhotoId      int
	PhotoOwnerId int
}

//easyjson:json
type PhotoCommentRestore struct {
	Comment

	PhotoId      int
	PhotoOwnerId int
}

//easyjson:json
type PhotoCommentDelete struct {
	OwnerId   int
	Id        int
	UserId    int
	DeleterId int
	PhotoId   int
}

//easyjson:json
type AudioNew struct {
	Audio
}

//easyjson:json
type VideoNew struct {
	Video
}

//easyjson:json
type VideoCommentNew struct {
	Comment

	VideoId      int
	VideoOwnerId int
}

//easyjson:json
type VideoCommentEdit struct {
	Comment

	VideoId      int
	VideoOwnerId int
}

//easyjson:json
type VideoCommentRestore struct {
	Comment

	VideoId      int
	VideoOwnerId int
}

//easyjson:json
type VideoCommentDelete struct {
	OwnerId   int
	Id        int
	UserId    int
	DeleterId int
	VideoId   int
}

//easyjson:json
type WallPostNew struct {
	Post

	PostponedId int
}

//easyjson:json
type WallRepost struct {
	Post

	PostponedId int
}

//easyjson:json
type WallReplyNew struct {
	Comment

	PostId      int
	PostOwnerId int
}

//easyjson:json
type WallReplyEdit struct {
	Comment

	PostId      int
	PostOwnerId int
}

//easyjson:json
type WallReplyRestore struct {
	Comment

	PostId      int
	PostOwnerId int
}

//easyjson:json
type WallReplyDelete struct {
	OwnerId   int
	Id        int
	DeleterId int
	PostId    int
}

//easyjson:json
type BoardPostNew struct {
	CommentBoard

	TopicId      int
	TopicOwnerId int
}

//easyjson:json
type BoardPostEdit struct {
	CommentBoard

	TopicId      int
	TopicOwnerId int
}

//easyjson:json
type BoardPostRestore struct {
	CommentBoard

	TopicId      int
	TopicOwnerId int
}

//easyjson:json
type BoardPostDelete struct {
	TopicOwnerId int
	TopicId      int
	Id           int
}

//easyjson:json
type MarketCommentNew struct {
	Comment

	MarketOwnerId int
	ItemId        int
}

//easyjson:json
type MarketCommentEdit struct {
	Comment

	MarketOwnerId int
	ItemId        int
}

//easyjson:json
type MarketCommentRestore struct {
	Comment

	MarketOwnerId int
	ItemId        int
}

//easyjson:json
type MarketCommentDelete struct {
	OwnerId   int
	Id        int
	UserId    int
	DeleterId int
	ItemId    int
}

//easyjson:json
type GroupLeave struct {
	UserId int
	Self   int
}

// TODO: JoinType as enum/type
//easyjson:json
type GroupJoin struct {
	UserId   int
	JoinType string
}

// TODO: Reason as enum/type
//easyjson:json
type UserBlock struct {
	AdminId     int
	UserId      int
	UnblockDate int
	Reason      int
	Comment     string
}

//easyjson:json
type UserUnblock struct {
	AdminId   int
	UserId    int
	ByEndDate int
}

//easyjson:json
type PollVoteNew struct {
	OwnerId  int
	PollId   int
	OptionId int
	UserId   int
}

// TODO: Levels as enum/type
//easyjson:json
type GroupOfficersEdit struct {
	AdminId  int
	UserId   int
	LevelOld int
	LevelNew int
}

//easyjson:json
type GroupChangeSettings struct {
	UserId int
	// TODO: Changes
}

//easyjson:json
type GroupChangePhoto struct {
	UserId int
	Photo  Photo
}
