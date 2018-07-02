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
	Id          int
	Date        int
	PeerId      int
	FromId      int
	Text        string
	RandomId    int
	Attachments []Attachment
	Important   bool
	// TODO: Geo
	Payload     string
	FwdMessages []Message

	// TODO: Action types
	Action struct {
		Type     string
		MemberId int
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
	Id             int
	FromId         int
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
	Id          int
	OwnerId     int
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
	Id        int
	OwnerId   int
	AccessKey string
	Artist    string
	Title     string
	Duration  int
	Url       string
	LyricsId  int
	AlbumId   int
	// TODO: Genre list
	GenreId  int
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
	Id        int
	AlbumId   int
	OwnerId   int
	UserId    int
	AccessKey string
	Text      string
	Date      int
	Sizes     []PhotoSize
	Width     int
	Height    int
}

//easyjson:json
type Document struct {
	Id        int
	OwnerId   int
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
	Id           int
	OwnerId      int
	Title        string
	Text         string
	Date         int
	Comments     int
	ReadComments int
	ViewUrl      string
}

//easyjson:json
type Poll struct {
	Id       int
	OWnerId  int
	Created  int
	Question string
	Votest   int
	AnswerId int
	Asnwers  []struct {
		Id    int
		Text  string
		Votes int
		Rate  float32
	}
	Anonymous int
}

//easyjson:json
type Page struct {
	Id                       int
	GroupId                  int
	CreatorId                int
	Title                    string
	CurrentUserCanEdit       int
	CurrentUserCanEditAccess int
	// TODO: enums
	WhoCanView int
	WhoCanEdit int
	Edited     int
	Created    int
	EditorId   int
	Views      int
	Parent     string
	Parent2    string
	Source     string
	Html       string
	ViewUrl    string
}

//easyjson:json
type Album struct {
	Id          int
	Thumb       Photo
	OwnerId     int
	Title       string
	Description string
	Created     int
	Updated     int
	Size        int
}

// TODO: PhotosList

//easyjson:json
type Market struct {
	Id          int
	OwnerId     int
	Title       string
	Description string
	Price       struct {
		Amount   int
		Currency struct {
			Id   int
			Name string
		}
		Text string
	}
	Category struct {
		Id      int
		Name    string
		Section struct {
			Id   int
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
	Id          int
	OwnerId     int
	Title       string
	Photo       Photo
	Count       int
	UpdatedTime int
}

//easyjson:json
type Sticker struct {
	ProductId int
	StickerId int
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
	PostId  int
	OwnerId int
}

//easyjson:json
type Gift struct {
	Id       int
	Thumb256 string
	Thumb96  string
	Thumb48  string
}

//easyjson:json
type Post struct {
	Id           int
	OwnerId      int
	FromId       int
	AccessKey    string
	CreatedBy    int
	Date         int
	Text         string
	ReplyOwnerId int
	ReplyPostId  int
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
	// TODO: PostType as enum/type
	PostType string
	// TODO: PostSource
	Attachments []Attachment
	// TODO: Geo
	SignerId    int
	CopyHistory []Post
	CanPin      int
	CanDelete   int
	CanEdit     int
	IsPinned    int
	MarkedAsAds int
}

//easyjson:json
type Wall struct {
	Post
	ToId int
}

//easyjson:json
type CommentBoard struct {
	Id          int
	FromId      int
	Date        int
	Text        string
	Attachments []Attachment
	Likes       struct {
		Count     int
		UserLikes int
		CanLike   int
	}
}
