package vk

import (
	"encoding/json"
	"fmt"
)

// Attachment is a wrapper for attachments
// Photo, Video, Audio, Document, Link, Note,
// Poll, Page, Album, MarketItem, MarketAlbum,
// Sticker, Wall, WallReply, Gift
type Attachment struct {
	Val interface{}
}

// UnmarshalJSON implements json.Unmarshaler interface
func (a *Attachment) UnmarshalJSON(data []byte) error {
	raw := map[string]json.RawMessage{}
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	aTypeRaw, ok := raw["type"]
	if !ok {
		return fmt.Errorf("Attachment doesn't have type field")
	}

	var aType string
	err = json.Unmarshal(aTypeRaw, &aType)
	if err != nil {
		return err
	}

	v, ok := raw[aType]
	if !ok {
		return fmt.Errorf("Attachment type is %v, but it has no field with such keys", aType)
	}

	switch aType {
	case "photo":
		val := Photo{}
		err = json.Unmarshal(v, &val)
		a.Val = val
	case "video":
		val := Video{}
		err = json.Unmarshal(v, &val)
		a.Val = val
	case "audio":
		val := Audio{}
		err = json.Unmarshal(v, &val)
		a.Val = val
	case "doc":
		val := Document{}
		err = json.Unmarshal(v, &val)
		a.Val = val
	case "link":
		val := Link{}
		err = json.Unmarshal(v, &val)
		a.Val = val
	case "note":
		val := Note{}
		err = json.Unmarshal(v, &val)
		a.Val = val
	case "poll":
		val := Poll{}
		err = json.Unmarshal(v, &val)
		a.Val = val
	case "page":
		val := Page{}
		err = json.Unmarshal(v, &val)
		a.Val = val
	case "album":
		val := Album{}
		err = json.Unmarshal(v, &val)
		a.Val = val
	case "market":
		val := MarketItem{}
		err = json.Unmarshal(v, &val)
		a.Val = val
	case "market_album":
		val := MarketAlbum{}
		err = json.Unmarshal(v, &val)
		a.Val = val
	case "sticker":
		val := Sticker{}
		err = json.Unmarshal(v, &val)
		a.Val = val
	case "wall":
		val := Wall{}
		err = json.Unmarshal(v, &val)
		a.Val = val
	case "wall_reply":
		val := WallReply{}
		err = json.Unmarshal(v, &val)
		a.Val = val
	case "gift":
		val := Gift{}
		err = json.Unmarshal(v, &val)
		a.Val = val
	default:
		err = fmt.Errorf("Unknown Attachment type: %v", aType)
	}

	return err
}

//easyjson:json
type CropPhoto struct {
	Photo Photo `json:"photo"`
	Crop  struct {
		X  int `json:"x"`
		Y  int `json:"y"`
		X2 int `json:"x2"`
		Y2 int `json:"y2"`
	} `json:"crop"`
	Rect struct {
		X  int `json:"x"`
		Y  int `json:"y"`
		X2 int `json:"x2"`
		Y2 int `json:"y2"`
	} `json:"rect"`
}

//easyjson:json
type BaseObject struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

//easyjson:json
type DatabaseCity struct {
	BaseObject
	Area      string `json:"area"`
	Region    string `json:"region"`
	Important int    `json:"important"`
}

//easyjson:json
type BaseObjectWithName struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//easyjson:json
type Category struct {
	BaseObjectWithName
	Subcategories []BaseObjectWithName `json:"subcategories"`
	PageCount     int                  `json:"page_count,omitempty"`
	PagePreviews  []Group              `json:"page_previews,omitempty"`
}

//easyjson:json
type MarketCategory struct {
	BaseObjectWithName
	Section BaseObjectWithName `json:"section"`
}

//easyjson:json
type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Deactivated string `json:"deactivated"`

	Sex        int    `json:"sex"`
	Domain     string `json:"domain"`
	MaidenName string `json:"maiden_name"`
	Nickname   string `json:"nickname"`
	ScreenName string `json:"screen_name"`
	Status     string `json:"status"`
	BirthDate  string `json:"bdate"`
	About      string `json:"about"`
	Books      string `json:"books"`
	Activites  string `json:"activites"`
	Interests  string `json:"interests"`
	Games      string `json:"games"`
	HomeTown   string `json:"home_town"`
	Movies     string `json:"movies"`
	Music      string `json:"music"`
	Quotes     string `json:"quotes"`
	TV         string `json:"tv"`
	Site       string `json:"site"`
	Verified   int    `json:"verified"`

	LastSeen struct {
		Time     int `json:"time"`
		Platform int `json:"platform"`
	} `json:"last_seen"`

	Online       *int    `json:"online"`
	OnlineMobile BoolInt `json:"online_mobile"`
	OnlineApp    int     `json:"online_app"`

	PhotoID      string `json:"photo_id"`
	PhotoMax     string `json:"photo_max"`
	PhotoMaxOrig string `json:"photo_max_orig"`

	Career []struct {
		GroupID   int    `json:"group_id"`
		Company   string `json:"company"`
		CountryID int    `json:"country_id"`
		CityID    int    `json:"city_id"`
		CityName  string `json:"city_name"`
		From      int    `json:"from"`
		Until     int    `json:"until"`
		Position  string `json:"position"`
	} `json:"career"`
	City        *BaseObject `json:"city"`
	Country     *BaseObject `json:"country"`
	MobilePhone string      `json:"mobile_phone"`
	HomePhone   string      `json:"home_phone"`

	CropPhoto *CropPhoto `json:"crop_photo"`

	Counters *struct {
		Albums        int `json:"albums"`
		Videos        int `json:"videos"`
		Audios        int `json:"audios"`
		Photos        int `json:"photos"`
		Notes         int `json:"notes"`
		Friends       int `json:"friends"`
		Groups        int `json:"groups"`
		OnlineFriends int `json:"online_friends"`
		MutualFriends int `json:"mutual_friends"`
		UserVideos    int `json:"user_videos"`
		Followers     int `json:"followers"`
		Pages         int `json:"pages"`
	} `json:"counters"`

	CommonCount    int `json:"common_count"`
	FollowersCount int `json:"followers_count"`
	FriendStatus   int `json:"friend_status"`

	// TODO: education, exports, first_name_{case}, last_name_{case}, military, occupation, personal, relatives, relation, schools, universities, wall_default
	// TODO: Connections
	// XXX: Can easyjson handle map[string]string?

	CanPost                BoolInt `json:"can_post"`
	CanSeeAllPosts         BoolInt `json:"can_see_all_posts"`
	CanSeeAudio            BoolInt `json:"can_see_audio"`
	CanSendFriendRequest   BoolInt `json:"can_send_friend_request"`
	CanWritePrivateMessage BoolInt `json:"can_write_private_message"`

	HasMobile BoolInt `json:"has_mobile"`
	HasPhoto  BoolInt `json:"has_photo"`

	IsFavorite       BoolInt `json:"is_favorite"`
	IsFriend         BoolInt `json:"is_friend"`
	IsHiddenFromFeed BoolInt `json:"is_hidden_from_feed"`

	Blacklisted     BoolInt `json:"blacklisted"`
	BlacklistedByMe BoolInt `json:"blacklisted_by_me"`

	IsClosed        bool `json:"is_closed"`
	CanAccessClosed bool `json:"can_access_closed"`
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
	URL    string  `json:"url"`
	Width  IntFrac `json:"width"`
	Height IntFrac `json:"height"`
}

//easyjson:json
type Group struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	ScreenName  string  `json:"screen_name"`
	IsClosed    BoolInt `json:"is_closed"`
	Deactivated string  `json:"deactivated"`
	InvitedBy   int     `json:"invited_by"`
	Type        string  `json:"type"`
	Description string  `json:"description"`
	FixedPost   int     `json:"fixed_post"`

	MainAlbumID int `json:"main_album_id"`
	MainSection int `json:"main_section"`
	Market      *struct {
		Enabled     int `json:"enabled"`
		PriceMin    int `json:"price_min"`
		PriceMax    int `json:"price_max"`
		MainAlbumID int `json:"main_album_id"`
		ContactID   int `json:"contact_id"`
		Currency    struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"currency"`
		CurrencyText string `json:"currency_text"`
	} `json:"market"`

	Photo50  string `json:"photo_50"`
	Photo100 string `json:"photo_100"`
	Photo200 string `json:"photo_200"`

	Activity  string `json:"activity"`
	AgeLimits int    `json:"age_limits"`

	AdminLevel        int     `json:"admin_level"`
	IsAdmin           BoolInt `json:"is_admin"`
	IsMember          BoolInt `json:"is_member"`
	IsFavorite        BoolInt `json:"is_favorite"`
	IsHiddenFromFeed  BoolInt `json:"is_hidden_from_feed"`
	IsMessagesBlocked BoolInt `json:"is_messages_blocked"`
	CanCreateTopic    BoolInt `json:"can_create_topic"`
	CanMessage        BoolInt `json:"can_message"`
	CanPost           BoolInt `json:"can_post"`
	CanSeeAllPosts    BoolInt `json:"can_see_all_posts"`
	CanUploadDoc      BoolInt `json:"can_upload_doc"`
	CanUploadVideo    BoolInt `json:"can_upload_video"`
	HasPhoto          BoolInt `json:"has_photo"`

	BanInfo *struct {
		EndDate int    `json:"end_date"`
		Comment string `json:"comment"`
	} `json:"ban_info"`

	City    *BaseObject `json:"city"`
	Country *BaseObject `json:"country"`

	Cover *struct {
		Enabled int         `json:"enabled"`
		Images  []BaseImage `json:"images"`
	} `json:"cover"`

	CropPhoto *CropPhoto `json:"crop_photo"`

	Contacts []struct {
		UserID      int    `json:"user_id"`
		Description string `json:"desc"`
		Phone       string `json:"phone"`
		Email       string `json:"email"`
	} `json:"contacts"`

	Links []MiniLink `json:"links"`

	Counters *struct {
		Albums int `json:"albums"`
		Videos int `json:"videos"`
		Audios int `json:"audios"`
		Photos int `json:"photos"`
		Topics int `json:"topics"`
		Docs   int `json:"docs"`
		Market int `json:"market"`
	} `json:"counters"`
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
type GroupAddress struct {
	// Address id
	ID int `json:"id,omitempty"`
	// Title of the place (Zinger, etc)
	Title string `json:"title,omitempty"`
	// String address to the place (Nevsky, 28)
	Address string `json:"address,omitempty"`
	// Additional address to the place (6 floor, left door)
	AdditionalAddress string `json:"additional_address,omitempty"`
	// Country id of address
	CountryID int `json:"country_id,omitempty"`
	// City id of address
	CityID int `json:"city_id,omitempty"`
	// Metro id of address
	MetroStationID int `json:"metro_station_id,omitempty"`
	// Address latitude
	Latitude float32 `json:"latitude,omitempty"`
	// Address longitude
	Longitude float32 `json:"longitude,omitempty"`
	// Distance from the point
	Distance int `json:"distance,omitempty"`
	// Status of information about timetable
	WorkInfoStatus string `json:"work_info_status,omitempty"`
	// Week timetable for the address
	// Timetable genTODOType /* #/definitions/groups_address_timetable */ `json:"timetable,omitempty"`
	// Address phone
	Phone string `json:"phone,omitempty"`
	// Time offset int minutes from utc time
	TimeOffset int `json:"time_offset,omitempty"`
}

//easyjson:json
type Message struct {
	ID int `json:"id"`
	// Unique auto-incremented number for all messages with this peer
	ConversationID int          `json:"conversation_message_id"`
	Date           int          `json:"date"`
	PeerID         int          `json:"peer_id"`
	FromID         int          `json:"from_id"`
	Text           string       `json:"text"`
	RandomID       int          `json:"random_id"`
	Attachments    []Attachment `json:"attachments"`
	Important      bool         `json:"important"`
	// TODO: Geo
	Payload           string    `json:"payload"`
	ForwardedMessages []Message `json:"fwd_messages"`

	// TODO: Action types
	Action *struct {
		Type     string `json:"type"`
		MemberID int    `json:"member_id"`
		Text     string `json:"text"`
		Email    string `json:"email"`
		Photo    struct {
			Photo50  string `json:"photo_50"`
			Photo100 string `json:"photo_100"`
			Photo200 string `json:"photo_200"`
		} `json:"photo"`
	} `json:"action"`
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
		ID      int    `json:"id"`
		Type    string `json:"type"`
		LocalID int    `json:"local_id"`
	} `json:"peer"`

	InRead      int `json:"in_read"`
	OutRead     int `json:"out_read"`
	UnreadCount int `json:"unread_count"`

	Important    bool `json:"important"`
	Unanswered   bool `json:"unanswered"`
	PushSettings *struct {
		DisabledUntil   int  `json:"disabled_until"`
		DisabledForever bool `json:"disabled_forever"`
		NoSound         bool `json:"no_sound"`
	} `json:"push_settings"`

	CanWrite *struct {
		Allowed bool `json:"allowed"`
		Reason  int  `json:"reason"`
	} `json:"can_write"`

	ChatSettings *struct {
		MembersCount  int      `json:"members_count"`
		Title         string   `json:"title"`
		PinnedMessage *Message `json:"pinned_message"`
		State         string   `json:"state"`

		Photo struct {
			Photo50  string `json:"photo_50"`
			Photo100 string `json:"photo_100"`
			Photo200 string `json:"photo_200"`
		} `json:"photo"`

		ActiveIDs []int `json:"active_ids"`

		IsGroupChannel bool `json:"is_group_channel"`
	} `json:"chat_settings"`
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
	ID           int    `json:"id"`
	Type         string `json:"type"`
	Title        string `json:"title"`
	AdminID      int    `json:"admin_id"`
	Users        []int  `json:"users"`
	MembersCount int    `json:"members_count"`

	PushSettings struct {
		Sound         BoolInt `json:"sound"`
		DisabledUntil int     `json:"disabled_until"`
	} `json:"push_settings"`

	Photo50  string `json:"photo_50"`
	Photo100 string `json:"photo_100"`
	Photo200 string `json:"photo_200"`

	Left   BoolInt `json:"left"`
	Kicked BoolInt `json:"kicked"`
}

//easyjson:json
type Comment struct {
	ID             int          `json:"id"`
	FromID         int          `json:"from_id"`
	Date           int          `json:"date"`
	Text           string       `json:"text"`
	ReplyToUser    int          `json:"reply_to_user"`
	ReplyToComment int          `json:"reply_to_comment"`
	Attachments    []Attachment `json:"attachments"`
}

//easyjson:json
type VideoFiles struct {
	MP240    string `json:"mp4_240"`
	MP360    string `json:"mp4_360"`
	MP480    string `json:"mp4_480"`
	MP720    string `json:"mp4_720"`
	MP1080   string `json:"mp4_1080"`
	External string `json:"external"`
}

//easyjson:json
type Video struct {
	ID          int    `json:"id"`
	OwnerID     int    `json:"owner_id"`
	AccessKey   string `json:"access_key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Duration    int    `json:"duration"`

	Photo130 string `json:"photo_130"`
	Photo320 string `json:"photo_320"`
	Photo640 string `json:"photo_640"`
	Photo800 string `json:"photo_800"`

	FirstFrame130 string `json:"first_frame_130"`
	FirstFrame160 string `json:"first_frame_160"`
	FirstFrame320 string `json:"first_frame_320"`
	FirstFrame800 string `json:"first_frame_800"`

	Files *VideoFiles `json:"files"`

	Date       int    `json:"date"`
	AddingDate int    `json:"adding_date"`
	Views      int    `json:"views"`
	Comments   int    `json:"comments"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
	Player     string `json:"player"`
	Platfrom   string `json:"platfrom"`

	CanEdit    BoolInt `json:"can_edit"`
	CanComment BoolInt `json:"can_comment"`
	CanRepost  BoolInt `json:"can_repost"`
	CanAdd     BoolInt `json:"can_add"`
	IsPrivate  BoolInt `json:"is_private"`
	Processing BoolInt `json:"processing"`
	Live       BoolInt `json:"live"`
	Upcoming   BoolInt `json:"upcoming"`
	Repeat     BoolInt `json:"repeat"`

	Likes struct {
		UserLikes int `json:"user_likes"`
		Count     int `json:"count"`
	} `json:"likes"`
	Reposts struct {
		UserLikes int `json:"user_likes"`
		Count     int `json:"count"`
	} `json:"reposts"`
}

//easyjson:json
type Audio struct {
	ID        int    `json:"id"`
	OwnerID   int    `json:"owner_id"`
	AccessKey string `json:"access_key"`
	Artist    string `json:"artist"`
	Title     string `json:"title"`
	Duration  int    `json:"duration"`
	URL       string `json:"url"`
	LyricsID  int    `json:"lyrics_id"`
	AlbumID   int    `json:"album_id"`
	// TODO: Genre list
	GenreID      int     `json:"genre_id"`
	TrackGenreID int     `json:"track_genre_id"`
	Date         int     `json:"date"`
	NoSearch     BoolInt `json:"no_search"`
	IsLicensed   bool    `json:"is_licensed"`
	IsHQ         bool    `json:"is_hq"`
}

//easyjson:json
type PhotoSize struct {
	BaseImage
	Type string `json:"type"`
}

//easyjson:json
type Photo struct {
	ID        int         `json:"id"`
	AlbumID   int         `json:"album_id"`
	OwnerID   int         `json:"owner_id"`
	UserID    int         `json:"user_id"`
	AccessKey string      `json:"access_key"`
	Text      string      `json:"text"`
	Date      int         `json:"date"`
	Sizes     []PhotoSize `json:"sizes"`
	Width     int         `json:"width"`
	Height    int         `json:"height"`
}

//easyjson:json
type Document struct {
	ID        int             `json:"id"`
	OwnerID   int             `json:"owner_id"`
	AccessKey string          `json:"access_key"`
	Title     string          `json:"title"`
	Size      int             `json:"size"`
	Ext       string          `json:"ext"`
	URL       string          `json:"url"`
	Date      int             `json:"date"`
	Type      int             `json:"type"`
	Preview   DocumentPreview `json:"preview"`
}

//easyjson:json
type DocumentPreview struct {
	Photo        *DocumentPreviewPhoto        `json:"photo"`
	Video        *DocumentPreviewVideo        `json:"video"`
	AudioMessage *DocumentPreviewAudioMessage `json:"audio_message"`
	Graffiti     *DocumentPreviewGraffiti     `json:"graffiti"`
}

//easyjson:json
type DocumentPreviewPhoto struct {
	Sizes []PhotoSize `json:"sizes"`
}

//easyjson:json
type DocumentPreviewVideo struct {
	Src      string `json:"src"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Filesize int    `json:"filesize"`
}

//easyjson:json
type DocumentPreviewAudioMessage struct {
	Duration int    `json:"duration"`
	Waveform []int  `json:"waveform"`
	LinkOGG  string `json:"link_ogg"`
	LinkMP3  string `json:"link_mp3"`
}

//easyjson:json
type DocumentPreviewGraffiti struct {
	Src    string `json:"src"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
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
	URL         string `json:"url"`
	Title       string `json:"title"`
	Caption     string `json:"caption"`
	Description string `json:"description"`
	Photo       Photo  `json:"photo"`
	// TODO: Product
	// TODO: Button
	PreviewPage string `json:"preview_page"`
	PreviewURL  string `json:"preview_url"`
}

//easyjson:json
type MiniLink struct {
	ID          int    `json:"id"`
	URL         string `json:"url"`
	Name        string `json:"name"`
	Description string `json:"desc"`
	Photo50     string `json:"photo_50"`
	Photo100    string `json:"photo_100"`
	Photo200    string `json:"photo_200"`
}

//easyjson:json
type Note struct {
	ID           int    `json:"id"`
	OwnerID      int    `json:"owner_id"`
	Title        string `json:"title"`
	Text         string `json:"text"`
	Date         int    `json:"date"`
	Comments     int    `json:"comments"`
	ReadComments int    `json:"read_comments"`
	ViewURL      string `json:"view_url"`
}

//easyjson:json
type PollAnswer struct {
	ID    int     `json:"id"`
	Text  string  `json:"text"`
	Votes int     `json:"votes"`
	Rate  float32 `json:"rate"`
}

//easyjson:json
type Poll struct {
	ID        int          `json:"id"`
	OwnerID   int          `json:"owner_id"`
	AuthorID  int          `json:"author_id"`
	Created   int          `json:"created"`
	Question  string       `json:"question"`
	Votes     int          `json:"votes"`
	AnswerIDs []int        `json:"answer_ids"`
	Asnwers   []PollAnswer `json:"asnwers"`
	Anonymous BoolInt      `json:"anonymous"`
	Multiple  BoolInt      `json:"multiple"`
	EndDate   int          `json:"end_date"`

	Closed    BoolInt `json:"closed"`
	IsBoard   BoolInt `json:"is_board"`
	CanEdit   BoolInt `json:"can_edit"`
	CanVote   BoolInt `json:"can_vote"`
	CanReport BoolInt `json:"can_report"`
	CanShare  BoolInt `json:"can_share"`
	Photo     *Photo  `json:"photo"`
	// TODO: friends
	// TODO: background
}

//easyjson:json
type BoardTopicPoll struct {
	PollID   int          `json:"poll_id"`
	OwnerID  int          `json:"owner_id"`
	Created  int          `json:"created"`
	Question string       `json:"question"`
	Votes    int          `json:"votes"`
	AnswerID int          `json:"answer_id"`
	Asnwers  []PollAnswer `json:"asnwers"`
}

//easyjson:json
type Page struct {
	ID                       int     `json:"id"`
	GroupID                  int     `json:"group_id"`
	CreatorID                int     `json:"creator_id"`
	Title                    string  `json:"title"`
	CurrentUserCanEdit       BoolInt `json:"current_user_can_edit"`
	CurrentUserCanEditAccess BoolInt `json:"current_user_can_edit_access"`
	// TODO: enums
	WhoCanView int    `json:"who_can_view"`
	WhoCanEdit int    `json:"who_can_edit"`
	Edited     int    `json:"edited"`
	Created    int    `json:"created"`
	EditorID   int    `json:"editor_id"`
	Views      int    `json:"views"`
	Parent     string `json:"parent"`
	Parent2    string `json:"parent2"`
	Source     string `json:"source"`
	HTML       string `json:"html"`
	ViewURL    string `json:"view_url"`
}

//easyjson:json
type Album struct {
	ID          int    `json:"id"`
	Thumb       Photo  `json:"thumb"`
	OwnerID     int    `json:"owner_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Created     int    `json:"created"`
	Updated     int    `json:"updated"`
	Size        int    `json:"size"`
}

// TODO: PhotosList

//easyjson:json
type MarketItem struct {
	ID          int    `json:"id"`
	OwnerID     int    `json:"owner_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       struct {
		Amount   int `json:"amount"`
		Currency struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"currency"`
		Text string `json:"text"`
	} `json:"price"`
	Category   MarketCategory `json:"category"`
	ThumbPhoto string         `json:"thumb_photo"`
	Date       int            `json:"date"`
	// TODO: Availability enum/type
	Availability int `json:"availability"`

	Photos     []Photo `json:"photos"`
	CanComment BoolInt `json:"can_comment"`
	CanRepost  BoolInt `json:"can_repost"`
	Likes      struct {
		UserLikes int `json:"user_likes"`
		Count     int `json:"count"`
	} `json:"likes"`
}

//easyjson:json
type MarketAlbum struct {
	ID          int    `json:"id"`
	OwnerID     int    `json:"owner_id"`
	Title       string `json:"title"`
	Photo       Photo  `json:"photo"`
	Count       int    `json:"count"`
	UpdatedTime int    `json:"updated_time"`
}

//easyjson:json
type Sticker struct {
	ProductID int `json:"product_id"`
	StickerID int `json:"sticker_id"`
	Images    []struct {
		URL    string `json:"url"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
	} `json:"images"`
	ImagesWithBackground []struct {
		URL    string `json:"url"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
	} `json:"images_with_background"`
}

// TODO: PrettyCards

//easyjson:json
type WallReply struct {
	Comment
	PostID  int `json:"post_id"`
	OwnerID int `json:"owner_id"`
}

//easyjson:json
type Gift struct {
	ID       int    `json:"id"`
	Thumb256 string `json:"thumb_256"`
	Thumb96  string `json:"thumb_96"`
	Thumb48  string `json:"thumb_48"`
}

//easyjson:json
type Post struct {
	ID           int     `json:"id"`
	OwnerID      int     `json:"owner_id"`
	FromID       int     `json:"from_id"`
	AccessKey    string  `json:"access_key"`
	CreatedBy    int     `json:"created_by"`
	Date         int     `json:"date"`
	Text         string  `json:"text"`
	ReplyOwnerID int     `json:"reply_owner_id"`
	ReplyPostID  int     `json:"reply_post_id"`
	FriendsOnly  BoolInt `json:"friends_only"`
	Comments     *struct {
		Count         int     `json:"count"`
		CanPost       BoolInt `json:"can_post"`
		GroupsCanPost BoolInt `json:"groups_can_post"`
		CanClose      BoolInt `json:"can_close"`
		CanOpen       BoolInt `json:"can_open"`
	} `json:"comments"`
	Likes *struct {
		Count      int     `json:"count"`
		UserLikes  BoolInt `json:"user_likes"`
		CanLike    BoolInt `json:"can_like"`
		CanPublish BoolInt `json:"can_publish"`
	} `json:"likes"`
	Reposts *struct {
		Count        int     `json:"count"`
		UserReposted BoolInt `json:"user_reposted"`
	} `json:"reposts"`
	Views *struct {
		Count int `json:"count"`
	} `json:"views"`
	PostType string `json:"post_type"`
	// TODO: PostSource
	Attachments []Attachment `json:"attachments"`
	// TODO: Geo
	SignerID    int     `json:"signer_id"`
	CopyHistory []Post  `json:"copy_history"`
	CanPin      BoolInt `json:"can_pin"`
	CanDelete   BoolInt `json:"can_delete"`
	CanEdit     BoolInt `json:"can_edit"`
	IsPinned    BoolInt `json:"is_pinned"`
	MarkedAsAds BoolInt `json:"marked_as_ads"`
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
	ToID int `json:"to_id"`
}

//easyjson:json
type CommentBoard struct {
	ID          int          `json:"id"`
	FromID      int          `json:"from_id"`
	Date        int          `json:"date"`
	Text        string       `json:"text"`
	Attachments []Attachment `json:"attachments"`
	Likes       struct {
		Count     int `json:"count"`
		UserLikes int `json:"user_likes"`
		CanLike   int `json:"can_like"`
	} `json:"likes"`
	// XXX: Is RealOffset real?
}

// TODO: VideoCatElement

//easyjson:json
type BoardTopic struct {
	ID        int     `json:"id"`
	Title     string  `json:"title"`
	Created   int     `json:"created"`
	CreatedBy int     `json:"created_by"`
	Updated   int     `json:"updated"`
	UpdatedBy int     `json:"updated_by"`
	IsClosed  BoolInt `json:"is_closed"`
	IsFixed   BoolInt `json:"is_fixed"`
	Comments  int     `json:"comments"`
}

//easyjson:json
type Place struct {
	// Place ID
	ID int `json:"id"`
	// Place title
	Title string `json:"title"`
	// Place latitude
	Latitude float32 `json:"latitude"`
	// Place longitude
	Longitude float32 `json:"longitude"`
	// Date of the place creation in Unixtime
	Created int `json:"created"`
	// URL of the place's icon
	Icon string `json:"icon"`
	// Checkins number
	Checkins int `json:"checkins"`
	// Place type
	Type string `json:"type"`
	// Country ID
	Country int `json:"country"`
	// City ID
	City int `json:"city"`
	// Place address
	Place string `json:"place"`
	// Distance to the place
	Distance int `json:"distance"`
	// Community ID
	GroupID int `json:"group_id"`
	// URL of the community's photo
	GroupPhoto string `json:"group_photo"`
}

//easyjson:json
type WallpostStats struct {
	// Subscribers reach
	ReachSubscribers int `json:"reach_subscribers"`
	// Total reach
	ReachTotal int `json:"reach_total"`
	// Link clickthrough
	Links int `json:"links"`
	// Clickthrough to community
	ToGroup int `json:"to_group"`
	// People have joined the group
	JoinGroup int `json:"join_group"`
	// Reports number
	Report int `json:"report"`
	// Hidings number
	Hide int `json:"hide"`
	// Unsubscribed members
	Unsubscribe int `json:"unsubscribe"`
}

const (
	StoryTypePhoto string = "photo"
	StoryTypeVideo        = "video"
)

//easyjson:json
type StoryVideo struct {
	Video     `json:"video"`
	IsPrivate BoolInt `json:"is_private"`
}

//easyjson:json
type Story struct {
	// Story ID.
	ID int `json:"id"`
	// Story owner's ID.
	OwnerID int `json:"owner_id"`
	// Date when story has been added in Unixtime.
	Date int `json:"date"`
	// Information whether current user has seen the story or not (0 - no, 1 - yes).
	Seen  BoolInt     `json:"seen"`
	Type  string      `json:"type"`
	Photo *Photo      `json:"photo"`
	Video *StoryVideo `json:"video"`
	// Views number.
	Views int `json:"views"`
	// Information whether current user can see the story (0 - no, 1 - yes).
	CanSee BoolInt `json:"can_see"`
	// Information whether current user can reply to the story (0 - no, 1 - yes).
	CanReply BoolInt `json:"can_reply"`
	// Information whether current user can share the story (0 - no, 1 - yes).
	CanShare BoolInt `json:"can_share"`
	// Information whether current user can comment the story (0 - no, 1 - yes).
	CanComment BoolInt `json:"can_comment"`
	// Information whether the story is deleted (false - no, true - yes).
	IsDeleted bool `json:"is_deleted"`
	// Information whether the story is expired (false - no, true - yes).
	IsExpired bool `json:"is_expired"`
	// Access key for private object.
	AccessKey string `json:"access_key"`
	// Parent story owner's ID.
	ParentStoryOwnerID int `json:"parent_story_owner_id"`
	// Parent story ID.
	ParentStoryID int `json:"parent_story_id"`
	// Access key for private object.
	ParentStoryAccessKey string `json:"parent_story_access_key"`
	ParentStory          *Story `json:"parent_story"`
	Link                 struct {
		// Link text
		Text string `json:"text"`
		// Link URL
		URL string `json:"url"`
	} `json:"link"`
	// Replies to current story.
	Replies []struct {
		// Replies number.
		Count int `json:"count"`
		// New replies number.
		New int `json:"new"`
	} `json:"replies"`
}

//easyjson:json
type NewsfeedItem struct {
	// Item type
	Type string `json:"type,omitempty"`
	// Item source ID
	SourceID int `json:"source_id,omitempty"`
	// Date when item has been added in Unixtime
	Date int `json:"date,omitempty"`
}
