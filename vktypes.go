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
type CropPhoto struct {
	Photo Photo
	Crop  struct {
		X  int
		Y  int
		X2 int
		Y2 int
	}
	Rect struct {
		X  int
		Y  int
		X2 int
		Y2 int
	}
}

//easyjson:json
type User struct {
	ID          int
	FirstName   string
	LastName    string
	Deactivated string

	Sex        int
	Domain     string
	MaidenName string
	Nickname   string
	ScreenName string
	Status     string
	BirthDate  string `json:"bdate"`
	About      string
	Books      string
	Activites  string
	Interests  string
	Games      string
	HomeTown   string
	Movies     string
	Music      string
	Quotes     string
	TV         string
	Site       string
	Verified   int

	LastSeen struct {
		Time     int
		Platform int
	}

	Online       *int
	OnlineMobile int
	OnlineApp    int

	PhotoID      string
	PhotoMax     string
	PhotoMaxOrig string

	Career []struct {
		GroupID   int
		Company   string
		CountryID int
		CityID    int
		CityName  string
		From      int
		Until     int
		Position  string
	}
	City *struct {
		ID    int
		Title string
	}
	Country *struct {
		ID    int
		Title string
	}
	Contacts *struct {
		MobilePhone string
		HomePhone   string
	}

	CropPhoto *CropPhoto

	Counters *struct {
		Albums        int
		Videos        int
		Audios        int
		Photos        int
		Notes         int
		Friends       int
		Groups        int
		OnlineFriends int
		MutualFriends int
		UserVideos    int
		Followers     int
		Pages         int
	}

	CommonCount    int
	FollowersCount int
	FriendStatus   int

	// TODO: education, exports, first_name_{case}, last_name_{case}, military, occupation, personal, relatives, relation, schools, universities, wall_default
	// TODO: Connections
	// XXX: Can easyjson handle map[string]string?

	CanPost                int
	CanSeeAllPosts         int
	CanSeeAudio            int
	CanSendFriendRequest   int
	CanWritePrivateMessage int

	HasMobile int
	HasPhoto  int

	IsFavorite       int
	IsFriend         int
	IsHiddenFromFeed int

	Blacklisted     int
	BlacklistedByMe int
}

const (
	UserFriendStatusNone            int = 0
	UserFriendStatusRequestSent         = 1
	UserFriendStatusIncomingRequest     = 2
	UserFriendStatusFriends             = 3
)

const (
	UserSexUnknown int = 0
	UserSexFemale      = 1
	UserSexMale        = 2
)

//easyjson:json
type Group struct {
	ID          int
	Name        string
	ScreenName  string
	IsClosed    int
	Deactivated string
	InvitedBy   int
	Type        string
	Description string
	FixedPost   int

	MainAlbumID int
	MainSection int
	Market      *struct {
		Enabled     int
		PriceMin    int
		PriceMax    int
		MainAlbumID int
		ContactID   int
		Currency    struct {
			ID   int
			Name string
		}
		CurrencyText string
	}

	Photo50  string
	Photo100 string
	Photo200 string

	Activity  string
	AgeLimits int

	AdminLevel        int
	IsAdmin           int
	IsMember          int
	IsFavorite        int
	IsHiddenFromFeed  int
	IsMessagesBlocked int
	CanCreateTopic    int
	CanMessage        int
	CanPost           int
	CanSeeAllPosts    int
	CanUploadDoc      int
	CanUploadVideo    int
	HasPhoto          int

	BanInfo *struct {
		EndDate int
		Comment string
	}

	City *struct {
		ID    int
		Title string
	}
	Country *struct {
		ID    int
		Title string
	}

	CropPhoto *CropPhoto

	Contacts []struct {
		UserID      int
		Description string `json:"desc"`
		Phone       string
		Email       string
	}

	Links []struct {
		ID          int
		URL         string
		Name        string
		Description string `json:"desc"`
		Photo50     string
		Photo100    string
	}

	Counters *struct {
		Albums int
		Videos int
		Audios int
		Photos int
		Topics int
		Docs   int
		Market int
	}
}

const (
	GroupIsOpen    int = 0
	GroupIsClosed      = 1
	GroupIsPrivate     = 2

	GroupAdminLevelNotAdmin      int = 0
	GroupAdminLevelModerator         = 1
	GroupAdminLevelEditor            = 1
	GroupAdminLevelAdministrator     = 1

	GroupTypeGroup string = "group"
	GroupTypePage         = "page"
	GroupTypeEvent        = "event"

	GroupAgeLimitUnknown int = 0
	GroupAgeLimitNone        = 1
	GroupAgeLimit16          = 2
	GroupAgeLimit18          = 3

	GroupMainSectionAbsent int = 0
	GroupMainSectionPhotos     = 1
	GroupMainSectionTopics     = 2
	GroupMainSectionAudio      = 3
	GroupMainSectionVideo      = 4
	GroupMainSectionMarket     = 5
)

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
	Payload           string
	ForwardedMessages []Message `json:"fwd_messages"`

	// TODO: Action types
	Action *struct {
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

const (
	MessageActionTypeChatPhotoUpdate      string = "chat_photo_update"
	MessageActionTypeChatPhotoRemove             = "chat_photo_remove"
	MessageActionTypeChatCreate                  = "chat_create"
	MessageActionTypeChatTitleUpdate             = "chat_title_update"
	MessageActionTypeChatInviteUser              = "chat_invite_user"
	MessageActionTypeChatKickUser                = "chat_kick_user"
	MessageActionTypeChatPinMessage              = "chat_pin_message"
	MessageActionTypeChatUnpinMessage            = "chat_unpin_message"
	MessageActionTypeChatInviteUserByLink        = "chat_invite_user_by_link"
)

//easyjson:json
type Conversation struct {
	Peer struct {
		ID      int
		Type    string
		LocalID int
	}

	InRead      int
	OutRead     int
	UnreadCount int

	Important    bool
	Unanswered   bool
	PushSettings *struct {
		DisabledUntil   int
		DisabledForever bool
		NoSound         bool
	}

	CanWrite *struct {
		Allowed bool
		Reason  int
	}

	ChatSettings *struct {
		MembersCount  int
		Title         string
		PinnedMessage *Message
		State         string

		Photo struct {
			Photo50  string
			Photo100 string
			Photo200 string
		}

		ActiveIDs []int
	}
}

const (
	PeerTypeUser  string = "user"
	PeerTypeChat         = "chat"
	PeerTypeGroup        = "group"
	PeerTypeEmail        = "email"

	CantWriteReasonBlockedOrDeleted      int = 18
	CantWriteReasonIsBlacklisted             = 900
	CantWriteReasonPeerBlockedGroup          = 901
	CantWriteReasonPrivacySettings           = 902
	CantWriteReasonGroupMessagesDisabled     = 915
	CantWriteReasonGroupMessagesBlocked      = 916
	CantWriteReasonChatNotAccessible         = 917
	CantWriteReasonEmailNotAccessible        = 918
	CantWriteReasonGroupNotAccessible        = 203

	ChatStateIn     string = "in"
	ChatStateKicked        = "kicked"
	ChatStateLeft          = "left"
)

//easyjson:json
type Chat struct {
	ID           int
	Type         string
	Title        string
	AdminID      int
	Users        []int
	MembersCount int

	PushSettings struct {
		Sound         int
		DisabledUntil int
	}

	Photo50  string
	Photo100 string
	Photo200 string

	Left   int
	Kicked int
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

	Files *VideoFiles

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
	URL       string
	LyricsID  int
	AlbumID   int
	// TODO: Genre list
	GenreID  int
	Date     int
	NoSearch int
	IsHQ     bool
}

//easyjson:json
type PhotoSize struct {
	Type   string
	URL    string
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
	URL       string
	Date      int
	Type      int
	Preview   DocumentPreview
}

type DocumentPreview struct {
	Photo *DocumentPreviewPhoto
	Video *DocumentPreviewVideo
}

type DocumentPreviewPhoto struct {
	Sizes []PhotoSize
}

type DocumentPreviewVideo struct {
	Src      string
	Width    int
	Height   int
	Filesize int
}

const (
	DocumentTypeText    int = 1
	DocumentTypeArchive     = 2
	DocumentTypeGIF         = 3
	DocumentTypeImage       = 4
	DocumentTypeAudio       = 5
	DocumentTypeVideo       = 6
	DocumentTypeEBook       = 7
	DocumentTypeUnknown     = 8
)

//easyjson:json
type Link struct {
	URL         string
	Title       string
	Caption     string
	Description string
	Photo       Photo
	// TODO: Product
	// TODO: Button
	PreviewPage string
	PreviewURL  string
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
	ViewURL      string
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
	HTML       string
	ViewURL    string
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
		URL    string
		Width  int
		Height int
	}
	ImagesWithBackground []struct {
		URL    string
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
	Comments     *struct {
		Count         int
		CanPost       int
		GroupsCanPost int
	}
	Likes *struct {
		Count      int
		UserLikes  int
		CanLike    int
		CanPublish int
	}
	Reposts *struct {
		Count        int
		UserReposted int
	}
	Views *struct {
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
	PostTypePost     string = "post"
	PostTypeCopy            = "copy"
	PostTypeReply           = "reply"
	PostTypePostpone        = "postpone"
	PostTypeSuggest         = "suggest"
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
