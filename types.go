package vk

import (
	"github.com/mailru/easyjson/jlexer"
)

// Attachment is a wrapper for attachments
// Photo, Video, Audio, Document, Link, Note,
// Poll, Page, Album, MarketItem, MarketAlbum,
// Sticker, Wall, WallReply, Gift
type Attachment struct {
	Val interface{}
}

// UnmarshalJSON implements json.Unmarshaler interface
func (v *Attachment) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	v.UnmarshalEasyJSON(&r)
	return r.Error()
}

// UnmarshalEasyJSON implements easyjson.Unmarshaler interface
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
			tmp := MarketItem{}
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
type BaseObject struct {
	ID    int
	Title string
}

//easyjson:json
type DatabaseCity struct {
	BaseObject
	Area      string
	Region    string
	Important int
}

//easyjson:json
type BaseObjectWithName struct {
	ID   int
	Name string
}

//easyjson:json
type Category struct {
	BaseObjectWithName
	Subcategories []BaseObjectWithName
}

//easyjson:json
type MarketCategory struct {
	BaseObjectWithName
	Section BaseObjectWithName
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
	OnlineMobile BoolInt
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
	City        *BaseObject
	Country     *BaseObject
	MobilePhone string
	HomePhone   string

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

	CanPost                BoolInt
	CanSeeAllPosts         BoolInt
	CanSeeAudio            BoolInt
	CanSendFriendRequest   BoolInt
	CanWritePrivateMessage BoolInt

	HasMobile BoolInt
	HasPhoto  BoolInt

	IsFavorite       BoolInt
	IsFriend         BoolInt
	IsHiddenFromFeed BoolInt

	Blacklisted     BoolInt
	BlacklistedByMe BoolInt
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
type BaseImage struct {
	URL    string
	Width  int
	Height int
}

//easyjson:json
type Group struct {
	ID          int
	Name        string
	ScreenName  string
	IsClosed    BoolInt
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
	IsAdmin           BoolInt
	IsMember          BoolInt
	IsFavorite        BoolInt
	IsHiddenFromFeed  BoolInt
	IsMessagesBlocked BoolInt
	CanCreateTopic    BoolInt
	CanMessage        BoolInt
	CanPost           BoolInt
	CanSeeAllPosts    BoolInt
	CanUploadDoc      BoolInt
	CanUploadVideo    BoolInt
	HasPhoto          BoolInt

	BanInfo *struct {
		EndDate int
		Comment string
	}

	City    *BaseObject
	Country *BaseObject

	Cover *struct {
		Enabled int
		Images  []BaseImage
	}

	CropPhoto *CropPhoto

	Contacts []struct {
		UserID      int
		Description string `json:"desc"`
		Phone       string
		Email       string
	}

	Links []MiniLink

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
		Sound         BoolInt
		DisabledUntil int
	}

	Photo50  string
	Photo100 string
	Photo200 string

	Left   BoolInt
	Kicked BoolInt
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
	Width      int
	Height     int
	Player     string
	Platfrom   string

	CanEdit    BoolInt
	CanComment BoolInt
	CanRepost  BoolInt
	CanAdd     BoolInt
	IsPrivate  BoolInt
	Processing BoolInt
	Live       BoolInt
	Upcoming   BoolInt
	Repeat     BoolInt

	Likes struct {
		UserLikes int
		Count     int
	}
	Reposts struct {
		UserLikes int
		Count     int
	}
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
	GenreID      int
	TrackGenreID int
	Date         int
	NoSearch     BoolInt
	IsLicensed   bool
	IsHQ         bool
}

//easyjson:json
type PhotoSize struct {
	BaseImage
	Type string
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

//easyjson:json
type DocumentPreview struct {
	Photo *DocumentPreviewPhoto
	Video *DocumentPreviewVideo
}

//easyjson:json
type DocumentPreviewPhoto struct {
	Sizes []PhotoSize
}

//easyjson:json
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
type MiniLink struct {
	ID          int
	URL         string
	Name        string
	Description string `json:"desc"`
	Photo50     string
	Photo100    string
	Photo200    string
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
type PollAnswer struct {
	ID    int
	Text  string
	Votes int
	Rate  float32
}

//easyjson:json
type Poll struct {
	ID        int
	OwnerID   int
	Created   int
	Question  string
	Votes     int
	AnswerID  int
	Asnwers   []PollAnswer
	Anonymous BoolInt
}

//easyjson:json
type BoardTopicPoll struct {
	PollID   int
	OwnerID  int
	Created  int
	Question string
	Votes    int
	AnswerID int
	Asnwers  []PollAnswer
}

//easyjson:json
type Page struct {
	ID                       int
	GroupID                  int
	CreatorID                int
	Title                    string
	CurrentUserCanEdit       BoolInt
	CurrentUserCanEditAccess BoolInt
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
type MarketItem struct {
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
	Category   MarketCategory
	ThumbPhoto string
	Date       int
	// TODO: Availability enum/type
	Availability int

	Photos     []Photo
	CanComment BoolInt
	CanRepost  BoolInt
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
	FriendsOnly  BoolInt
	Comments     *struct {
		Count         int
		CanPost       BoolInt
		GroupsCanPost BoolInt
	}
	Likes *struct {
		Count      int
		UserLikes  BoolInt
		CanLike    BoolInt
		CanPublish BoolInt
	}
	Reposts *struct {
		Count        int
		UserReposted BoolInt
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
	CanPin      BoolInt
	CanDelete   BoolInt
	CanEdit     BoolInt
	IsPinned    BoolInt
	MarkedAsAds BoolInt
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
	// XXX: Is RealOffset real?
}

// TODO: VideoCatElement

//easyjson:json
type BoardTopic struct {
	ID        int
	Title     string
	Created   int
	CreatedBy int
	Updated   int
	UpdatedBy int
	IsClosed  BoolInt
	IsFixed   BoolInt
	Comments  int
}

//easyjson:json
type Place struct {
	// Place ID
	ID int
	// Place title
	Title string
	// Place latitude
	Latitude float32
	// Place longitude
	Longitude float32
	// Date of the place creation in Unixtime
	Created int
	// URL of the place's icon
	Icon string
	// Checkins number
	Checkins int
	// Place type
	Type string
	// Country ID
	Country int
	// City ID
	City int
	// Place address
	Place string
	// Distance to the place
	Distance int
	// Community ID
	GroupID int
	// URL of the community's photo
	GroupPhoto string
}

//easyjson:json
type WallpostStats struct {
	// Subscribers reach
	ReachSubscribers int
	// Total reach
	ReachTotal int
	// Link clickthrough
	Links int
	// Clickthrough to community
	ToGroup int
	// People have joined the group
	JoinGroup int
	// Reports number
	Report int
	// Hidings number
	Hide int
	// Unsubscribed members
	Unsubscribe int
}

const (
	StoryTypePhoto string = "photo"
	StoryTypeVideo        = "video"
)

//easyjson:json
type StoryVideo struct {
	Video
	IsPrivate BoolInt
}

//easyjson:json
type Story struct {
	// Story ID.
	ID int
	// Story owner's ID.
	OwnerID int
	// Date when story has been added in Unixtime.
	Date int
	// Information whether current user has seen the story or not (0 - no, 1 - yes).
	Seen  BoolInt
	Type  string
	Photo *Photo
	Video *StoryVideo
	// Views number.
	Views int
	// Information whether current user can see the story (0 - no, 1 - yes).
	CanSee BoolInt
	// Information whether current user can reply to the story (0 - no, 1 - yes).
	CanReply BoolInt
	// Information whether current user can share the story (0 - no, 1 - yes).
	CanShare BoolInt
	// Information whether current user can comment the story (0 - no, 1 - yes).
	CanComment BoolInt
	// Information whether the story is deleted (false - no, true - yes).
	IsDeleted bool
	// Information whether the story is expired (false - no, true - yes).
	IsExpired bool
	// Access key for private object.
	AccessKey string
	// Parent story owner's ID.
	ParentStoryOwnerID int
	// Parent story ID.
	ParentStoryID int
	// Access key for private object.
	ParentStoryAccessKey string
	ParentStory          *Story
	Link                 struct {
		// Link text
		Text string
		// Link URL
		URL string
	}
	// Replies to current story.
	Replies []struct {
		// Replies number.
		Count int
		// New replies number.
		New int
	}
}
