package vkCallbackApi

import (
	"github.com/mailru/easyjson/jlexer"
)

// TODO: Stories

type Attachment struct {
	Val interface{}
}

func (v *Attachment) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	v.UnmarshalEasyJSON(&r)
	return r.Error()
}

func (v *Attachment) UnmarshalEasyJSON(in *jlexer.Lexer) {
	in.Delim('{')

	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()

		switch key {
		case "photo":
			tmp := Photo{}
			tmp.UnmarshalEasyJSON(in)
			v.Val = tmp
		case "video":
			tmp := Video{}
			tmp.UnmarshalEasyJSON(in)
			v.Val = tmp
		case "audio":
			tmp := Audio{}
			tmp.UnmarshalEasyJSON(in)
			v.Val = tmp
		case "doc":
			tmp := Document{}
			tmp.UnmarshalEasyJSON(in)
			v.Val = tmp
		case "link":
			tmp := Link{}
			tmp.UnmarshalEasyJSON(in)
			v.Val = tmp
		case "note":
			tmp := Note{}
			tmp.UnmarshalEasyJSON(in)
			v.Val = tmp
		case "poll":
			tmp := Poll{}
			tmp.UnmarshalEasyJSON(in)
			v.Val = tmp
		case "page":
			tmp := Page{}
			tmp.UnmarshalEasyJSON(in)
			v.Val = tmp
		case "album":
			tmp := Album{}
			tmp.UnmarshalEasyJSON(in)
			v.Val = tmp
		case "market":
			tmp := Market{}
			tmp.UnmarshalEasyJSON(in)
			v.Val = tmp
		case "market_album":
			tmp := MarketAlbum{}
			tmp.UnmarshalEasyJSON(in)
			v.Val = tmp
		case "sticker":
			tmp := Sticker{}
			tmp.UnmarshalEasyJSON(in)
			v.Val = tmp
		case "wall":
			tmp := Wall{}
			tmp.UnmarshalEasyJSON(in)
			v.Val = tmp
		case "wall_reply":
			tmp := WallReply{}
			tmp.UnmarshalEasyJSON(in)
			v.Val = tmp
		case "gift":
			tmp := Gift{}
			tmp.UnmarshalEasyJSON(in)
			v.Val = tmp
		default:
			in.SkipRecursive()
		}

		in.WantComma()
	}
	in.Delim('}')
}

//easyjson:json
type Message struct {
	ID          int
	Date        int
	PeerID      int
	FromID      int
	Text        string
	RandomID    int
	Attachments []Attachment
	Important   bool
	// TODO: Geo
	Payload     string
	FwdMessages []Message

	// TODO: Action types
	Action struct {
		Type     string
		MemberID int
		Text     string
		Email    string
		Photo    struct {
			Photo50  string
			Photo100 string
			Photo200 string
		}
	}
}

//easyjson:json
type Comment struct {
	ID             int
	FromID         int
	Date           int
	Text           string
	ReplyToUser    int
	ReplyToComment int
	Attachments    []Attachment
}

//easyjson:json
type VideoFiles struct {
	MP240    string `json:"mp4_240"`
	MP360    string `json:"mp4_360"`
	MP480    string `json:"mp4_480"`
	MP720    string `json:"mp4_720"`
	MP1080   string `json:"mp4_1080"`
	External string
}

//easyjson:json
type Video struct {
	ID          int
	OwnerID     int
	AccessKey   string
	Title       string
	Description string
	Duration    int

	Photo130 string `json:"photo_130"`
	Photo320 string `json:"photo_320"`
	Photo640 string `json:"photo_640"`
	Photo800 string `json:"photo_800"`

	FirstFrame130 string `json:"first_frame_130"`
	FirstFrame160 string `json:"first_frame_160"`
	FirstFrame320 string `json:"first_frame_320"`
	FirstFrame800 string `json:"first_frame_800"`

	Files VideoFiles

	Date       int
	AddingDate int
	Views      int
	Comments   int
	Player     string
	Platfrom   string
	CanEdit    int
	CanAdd     int
	IsPrivate  int
	Processing int
	Live       int
	Upcoming   int
	Repeat     int
}

//easyjson:json
type Audio struct {
	ID        int
	OwnerID   int
	AccessKey string
	Artist    string
	Title     string
	Duration  int
	Url       string
	LyricsID  int
	AlbumID   int
	// TODO: Genre list
	GenreID  int
	Date     int
	NoSearch int
	IsHq     int
}

//easyjson:json
type PhotoSize struct {
	Type   string
	Url    string
	Width  int
	Height int
}

//easyjson:json
type Photo struct {
	ID        int
	AlbumID   int
	OwnerID   int
	UserID    int
	AccessKey string
	Text      string
	Date      int
	Sizes     []PhotoSize
	Width     int
	Height    int
}

//easyjson:json
type Document struct {
	ID        int
	OwnerID   int
	AccessKey string
	Title     string
	Size      int
	Ext       string
	Url       string
	Date      int
	// TODO: as enum
	Type int
	// TODO: Preview
}

//easyjson:json
type Link struct {
	Url         string
	Title       string
	Caption     string
	Description string
	Photo       Photo
	// TODO: Product
	// TODO: Button
	PreviewPage string
	PreviewUrl  string
}

//easyjson:json
type Note struct {
	ID           int
	OwnerID      int
	Title        string
	Text         string
	Date         int
	Comments     int
	ReadComments int
	ViewUrl      string
}

//easyjson:json
type Poll struct {
	ID       int
	OWnerID  int
	Created  int
	Question string
	Votest   int
	AnswerID int
	Asnwers  []struct {
		ID    int
		Text  string
		Votes int
		Rate  float32
	}
	Anonymous int
}

//easyjson:json
type Page struct {
	ID                       int
	GroupID                  int
	CreatorID                int
	Title                    string
	CurrentUserCanEdit       int
	CurrentUserCanEditAccess int
	// TODO: enums
	WhoCanView int
	WhoCanEdit int
	Edited     int
	Created    int
	EditorID   int
	Views      int
	Parent     string
	Parent2    string
	Source     string
	Html       string
	ViewUrl    string
}

//easyjson:json
type Album struct {
	ID          int
	Thumb       Photo
	OwnerID     int
	Title       string
	Description string
	Created     int
	Updated     int
	Size        int
}

// TODO: PhotosList

//easyjson:json
type Market struct {
	ID          int
	OwnerID     int
	Title       string
	Description string
	Price       struct {
		Amount   int
		Currency struct {
			ID   int
			Name string
		}
		Text string
	}
	Category struct {
		ID      int
		Name    string
		Section struct {
			ID   int
			Name string
		}
	}
	ThumbPhoto string
	Date       int
	// TODO: Availability enum/type
	Availability int

	Photos     []Photo
	CanComment int
	CanRepost  int
	Likes      struct {
		UserLikes int
		Count     int
	}
}

//easyjson:json
type MarketAlbum struct {
	ID          int
	OwnerID     int
	Title       string
	Photo       Photo
	Count       int
	UpdatedTime int
}

//easyjson:json
type Sticker struct {
	ProductID int
	StickerID int
	Images    []struct {
		Url    string
		Width  int
		Height int
	}
	ImagesWithBackground []struct {
		Url    string
		Width  int
		Height int
	}
}

// TODO: PrettyCards

//easyjson:json
type WallReply struct {
	Comment
	PostID  int
	OwnerID int
}

//easyjson:json
type Gift struct {
	ID       int
	Thumb256 string
	Thumb96  string
	Thumb48  string
}

//easyjson:json
type Post struct {
	ID           int
	OwnerID      int
	FromID       int
	AccessKey    string
	CreatedBy    int
	Date         int
	Text         string
	ReplyOwnerID int
	ReplyPostID  int
	FriendsOnly  int
	Comments     struct {
		Count         int
		CanPost       int
		GroupsCanPost int
	}
	Likes struct {
		Count      int
		UserLikes  int
		CanLike    int
		CanPublish int
	}
	Reposts struct {
		Count        int
		UserReposted int
	}
	Views struct {
		Count int
	}
	PostType string
	// TODO: PostSource
	Attachments []Attachment
	// TODO: Geo
	SignerID    int
	CopyHistory []Post
	CanPin      int
	CanDelete   int
	CanEdit     int
	IsPinned    int
	MarkedAsAds int
}

const (
	PostTypePost string = "post"
	PostTypeCopy = "copy"
	PostTypeReply = "reply"
	PostTypePostpone = "postpone"
	PostTypeSuggest = "suggest"
)

//easyjson:json
type Wall struct {
	Post
	ToID int
}

//easyjson:json
type CommentBoard struct {
	ID          int
	FromID      int
	Date        int
	Text        string
	Attachments []Attachment
	Likes       struct {
		Count     int
		UserLikes int
		CanLike   int
	}
}
