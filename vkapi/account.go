package vkapi

import (
	"encoding/json"
	"strconv"

	"github.com/stek29/vk"
)

// Account implements VK API namespace `account`
type Account struct {
	API vk.API
}

// AccountGetCountersParams are params for Account.GetCounters
type AccountGetCountersParams struct {
	// Counters to be returned.
	Filter CSVStringSlice `url:"filter,omitempty"`
}

// AccountGetCountersResponse is response for Account.GetCounters
//easyjson:json
type AccountGetCountersResponse struct {
	// New friends requests number
	Friends int `json:"friends,omitempty"`
	// New messages number
	Messages int `json:"messages,omitempty"`
	// New photo tags number
	Photos int `json:"photos,omitempty"`
	// New video tags number
	Videos int `json:"videos,omitempty"`
	// New gifts number
	Gifts int `json:"gifts,omitempty"`
	// New events number
	Events int `json:"events,omitempty"`
	// New groups number
	Groups int `json:"groups,omitempty"`
	// New notifications number
	Notifications int `json:"notifications,omitempty"`
	// New app requests number
	AppRequests int `json:"app_requests,omitempty"`
	// New friends suggestions number
	FriendsSuggestions int `json:"friends_suggestions,omitempty"`
}

// GetCounters Returns non-null values of user counters.
func (v Account) GetCounters(params AccountGetCountersParams) (*AccountGetCountersResponse, error) {
	r, err := v.API.Request("account.getCounters", params)
	if err != nil {
		return nil, err
	}

	var resp AccountGetCountersResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AccountSetNameInMenuParams are params for Account.SetNameInMenu
type AccountSetNameInMenuParams struct {
	// User ID.
	UserID int `url:"user_id"`
	// Application screen name.
	Name string `url:"name,omitempty"`
}

// SetNameInMenu Sets an application screen name (up to 17 characters), that is shown to the user in the left menu.
func (v Account) SetNameInMenu(params AccountSetNameInMenuParams) (bool, error) {
	r, err := v.API.Request("account.setNameInMenu", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// AccountSetOnlineParams are params for Account.SetOnline
type AccountSetOnlineParams struct {
	// '1' if videocalls are available for current device.
	Voip bool `url:"voip,omitempty"`
}

// SetOnline Marks the current user as online for 15 minutes.
func (v Account) SetOnline(params AccountSetOnlineParams) (bool, error) {
	r, err := v.API.Request("account.setOnline", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// SetOffline Marks a current user as offline.
func (v Account) SetOffline() (bool, error) {
	r, err := v.API.Request("account.setOffline", nil)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// AccountRegisterDeviceParams are params for Account.RegisterDevice
type AccountRegisterDeviceParams struct {
	// Device token used to send notifications. (for mpns, the token shall be URL for sending of notifications)
	Token string `url:"token"`
	// String name of device model.
	DeviceModel string `url:"device_model,omitempty"`
	// Device year.
	DeviceYear int `url:"device_year,omitempty"`
	// Unique device ID.
	DeviceID string `url:"device_id"`
	// String version of device operating system.
	SystemVersion string `url:"system_version,omitempty"`
	// Push settings in a [vk.com/dev/push_settings|special format].
	Settings string `url:"settings,omitempty"`
	Sandbox  bool   `url:"sandbox,omitempty"`
}

// RegisterDevice Subscribes an iOS/Android/Windows Phone-based device to receive push notifications
func (v Account) RegisterDevice(params AccountRegisterDeviceParams) (bool, error) {
	r, err := v.API.Request("account.registerDevice", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// AccountUnregisterDeviceParams are params for Account.UnregisterDevice
type AccountUnregisterDeviceParams struct {
	// Unique device ID.
	DeviceID string `url:"device_id,omitempty"`
	Sandbox  bool   `url:"sandbox,omitempty"`
}

// UnregisterDevice Unsubscribes a device from push notifications.
func (v Account) UnregisterDevice(params AccountUnregisterDeviceParams) (bool, error) {
	r, err := v.API.Request("account.unregisterDevice", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// AccountSetSilenceModeParams are params for Account.SetSilenceMode
type AccountSetSilenceModeParams struct {
	// Unique device ID.
	DeviceID string `url:"device_id,omitempty"`
	// Time in seconds for what notifications should be disabled. '-1' to disable forever.
	Time int `url:"time,omitempty"`
	// Destination ID. "For user: 'User ID', e.g. '12345'. For chat: '2000000000' + 'Chat ID', e.g. '2000000001'. For community: '- Community ID', e.g. '-12345'. "
	PeerID int `url:"peer_id,omitempty"`
	// '1' — to enable sound in this dialog, '0' — to disable sound. Only if 'peer_id' contains user or community ID.
	Sound int `url:"sound,omitempty"`
}

// SetSilenceMode Mutes push notifications for the set period of time.
func (v Account) SetSilenceMode(params AccountSetSilenceModeParams) (bool, error) {
	r, err := v.API.Request("account.setSilenceMode", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// AccountGetPushSettingsParams are params for Account.GetPushSettings
type AccountGetPushSettingsParams struct {
	// Unique device ID.
	DeviceID string `url:"device_id,omitempty"`
}

// AccountGetPushSettingsResponse is response for Account.GetPushSettings
//easyjson:json
type AccountGetPushSettingsResponse struct {
	// Information whether notifications are disabled
	Disabled vk.BoolInt `json:"disabled,omitempty"`
	// Time until that notifications are disabled in Unixtime
	DisabledUntil int `json:"disabled_until,omitempty"`
	Conversations struct {
		// Items count
		Count int `json:"count,omitempty"`
		Items []struct {
			// Peer ID
			PeerID int `json:"peer_id,omitempty"`
			// Information whether the sound are enabled
			Sound vk.BoolInt `json:"sound,omitempty"`
			// Time until that notifications are disabled in seconds
			DisabledUntil int `json:"disabled_until,omitempty"`
		} `json:"items,omitempty"`
	} `json:"conversations,omitempty"`
	Settings struct {
		// Settings parameters
		Msg []string `json:"msg,omitempty"`
		// Settings parameters
		Chat []string `json:"chat,omitempty"`
		// Settings parameters
		Friend []string `json:"friend,omitempty"`
		// Settings parameters
		FriendFound []string `json:"friend_found,omitempty"`
		// Settings parameters
		FriendAccepted []string `json:"friend_accepted,omitempty"`
		// Settings parameters
		Reply []string `json:"reply,omitempty"`
		// Settings parameters
		Comment []string `json:"comment,omitempty"`
		// Settings parameters
		Mention []string `json:"mention,omitempty"`
		// Settings parameters
		Like []string `json:"like,omitempty"`
		// Settings parameters
		Repost []string `json:"repost,omitempty"`
		// Settings parameters
		WallPost []string `json:"wall_post,omitempty"`
		// Settings parameters
		WallPublish []string `json:"wall_publish,omitempty"`
		// Settings parameters
		GroupInvite []string `json:"group_invite,omitempty"`
		// Settings parameters
		GroupAccepted []string `json:"group_accepted,omitempty"`
		// Settings parameters
		EventSoon []string `json:"event_soon,omitempty"`
		// Settings parameters
		PhotosTag []string `json:"photos_tag,omitempty"`
		// Settings parameters
		AppRequest []string `json:"app_request,omitempty"`
		// Settings parameters
		SdkOpen []string `json:"sdk_open,omitempty"`
		// Settings parameters
		NewPost []string `json:"new_post,omitempty"`
		// Settings parameters
		Birthday []string `json:"birthday,omitempty"`
	} `json:"settings,omitempty"`
}

// GetPushSettings Gets settings of push notifications.
func (v Account) GetPushSettings(params AccountGetPushSettingsParams) (*AccountGetPushSettingsResponse, error) {
	r, err := v.API.Request("account.getPushSettings", params)
	if err != nil {
		return nil, err
	}

	var resp AccountGetPushSettingsResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AccountSetPushSettingsParams are params for Account.SetPushSettings
type AccountSetPushSettingsParams struct {
	// Unique device ID.
	DeviceID string `url:"device_id"`
	// Push settings in a [vk.com/dev/push_settings|special format].
	Settings string `url:"settings,omitempty"`
	// Notification key.
	Key string `url:"key,omitempty"`
	// New value for the key in a [vk.com/dev/push_settings|special format].
	Value CSVStringSlice `url:"value,omitempty"`
}

// SetPushSettings Change push settings.
func (v Account) SetPushSettings(params AccountSetPushSettingsParams) (bool, error) {
	r, err := v.API.Request("account.setPushSettings", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// AccountGetAppPermissionsParams are params for Account.GetAppPermissions
type AccountGetAppPermissionsParams struct {
	// User ID whose settings information shall be got. By default: current user.
	UserID int `url:"user_id"`
}

// AccountGetAppPermissionsResponse is response for Account.GetAppPermissions
// Permissions mask
type AccountGetAppPermissionsResponse int

// GetAppPermissions Gets settings of the user in this application.
func (v Account) GetAppPermissions(params AccountGetAppPermissionsParams) (AccountGetAppPermissionsResponse, error) {
	r, err := v.API.Request("account.getAppPermissions", params)
	if err != nil {
		return 0, err
	}

	var resp AccountGetAppPermissionsResponse

	var cnv int
	cnv, err = strconv.Atoi(string(r))
	resp = AccountGetAppPermissionsResponse(cnv)

	if err != nil {
		return 0, err
	}
	return resp, nil
}

// AccountGetActiveOffersParams are params for Account.GetActiveOffers
type AccountGetActiveOffersParams struct {
	Offset int `url:"offset,omitempty"`
	// Number of results to return.
	Count int `url:"count,omitempty"`
}

// AccountGetActiveOffersResponse is response for Account.GetActiveOffers
//easyjson:json
type AccountGetActiveOffersResponse struct {
	// Total number
	Count int `json:"count,omitempty"`
	Items []struct {
		// Offer ID
		ID int `json:"id,omitempty"`
		// Offer title
		Title string `json:"title,omitempty"`
		// Instruction how to process the offer
		Instruction string `json:"instruction,omitempty"`
		// Instruction how to process the offer (HTML format)
		InstructionHTML string `json:"instruction_html,omitempty"`
		// Offer short description
		ShortDescription string `json:"short_description,omitempty"`
		// Offer description
		Description string `json:"description,omitempty"`
		// URL of the preview image
		Img string `json:"img,omitempty"`
		// Offer tag
		Tag string `json:"tag,omitempty"`
		// Offer price
		Price int `json:"price,omitempty"`
	} `json:"items,omitempty"`
}

// GetActiveOffers Returns a list of active ads (offers) which executed by the user will bring him/her respective number of votes to his balance in the application.
func (v Account) GetActiveOffers(params AccountGetActiveOffersParams) (*AccountGetActiveOffersResponse, error) {
	r, err := v.API.Request("account.getActiveOffers", params)
	if err != nil {
		return nil, err
	}

	var resp AccountGetActiveOffersResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AccountGetBannedParams are params for Account.GetBanned
type AccountGetBannedParams struct {
	// Offset needed to return a specific subset of results.
	Offset int `url:"offset,omitempty"`
	// Number of results to return.
	Count int `url:"count,omitempty"`
}

// AccountGetBannedResponse is response for Account.GetBanned
//easyjson:json
type AccountGetBannedResponse struct {
	// Total number
	Count int       `json:"count,omitempty"`
	Items []vk.User `json:"items,omitempty"`
}

// GetBanned Returns a user's blacklist.
func (v Account) GetBanned(params AccountGetBannedParams) (*AccountGetBannedResponse, error) {
	r, err := v.API.Request("account.getBanned", params)
	if err != nil {
		return nil, err
	}

	var resp AccountGetBannedResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AccountGetInfoParams are params for Account.GetInfo
type AccountGetInfoParams struct {
	// Fields to return. Possible values: *'country' — user country,, *'https_required' — is "HTTPS only" option enabled,, *'own_posts_default' — is "Show my posts only" option is enabled,, *'no_wall_replies' — are wall replies disabled or not,, *'intro' — is intro passed by user or not,, *'lang' — user language. By default: all.
	Fields CSVStringSlice `url:"fields,omitempty"`
}

// AccountGetInfoResponse is response for Account.GetInfo
//easyjson:json
type AccountGetInfoResponse struct {
	// Country code
	Country string `json:"country,omitempty"`
	// Information whether HTTPS-only is enabled
	HTTPSRequired vk.BoolInt `json:"https_required,omitempty"`
	// Information whether only owners posts should be shown
	OwnPostsDefault vk.BoolInt `json:"own_posts_default,omitempty"`
	// Information whether wall comments should be hidden
	NoWallReplies vk.BoolInt `json:"no_wall_replies,omitempty"`
	// Information whether user has been processed intro
	Intro vk.BoolInt `json:"intro,omitempty"`
	// Language ID
	Lang int `json:"lang,omitempty"`
	// Two factor authentication is enabled
	X2FaRequired vk.BoolInt `json:"2fa_required,omitempty"`
}

// GetInfo Returns current account info.
func (v Account) GetInfo(params AccountGetInfoParams) (*AccountGetInfoResponse, error) {
	r, err := v.API.Request("account.getInfo", params)
	if err != nil {
		return nil, err
	}

	var resp AccountGetInfoResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AccountSetInfoParams are params for Account.SetInfo
type AccountSetInfoParams struct {
	// Setting name.
	Name string `url:"name,omitempty"`
	// Setting value.
	Value string `url:"value,omitempty"`
}

// SetInfo Allows to edit the current account info.
func (v Account) SetInfo(params AccountSetInfoParams) (bool, error) {
	r, err := v.API.Request("account.setInfo", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// AccountChangePasswordParams are params for Account.ChangePassword
type AccountChangePasswordParams struct {
	// Session id received after the [vk.com/dev/auth.restore|auth.restore] method is executed. (If the password is changed right after the access was restored)
	RestoreSid string `url:"restore_sid,omitempty"`
	// Hash received after a successful OAuth authorization with a code got by SMS. (If the password is changed right after the access was restored)
	ChangePasswordHash string `url:"change_password_hash,omitempty"`
	// Current user password.
	OldPassword string `url:"old_password,omitempty"`
	// New password that will be set as a current
	NewPassword string `url:"new_password"`
}

// AccountChangePasswordResponse is response for Account.ChangePassword
//easyjson:json
type AccountChangePasswordResponse struct {
	// New token
	Token string `json:"token,omitempty"`
	// New secret
	Secret string `json:"secret,omitempty"`
}

// ChangePassword Changes a user password after access is successfully restored with the [vk.com/dev/auth.restore|auth.restore] method.
func (v Account) ChangePassword(params AccountChangePasswordParams) (*AccountChangePasswordResponse, error) {
	r, err := v.API.Request("account.changePassword", params)
	if err != nil {
		return nil, err
	}

	var resp AccountChangePasswordResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AccountGetProfileInfoResponse is response for Account.GetProfileInfo
//easyjson:json
type AccountGetProfileInfoResponse struct {
	// User first name
	FirstName string `json:"first_name,omitempty"`
	// User last name
	LastName string `json:"last_name,omitempty"`
	// User maiden name
	MaidenName string `json:"maiden_name,omitempty"`
	// Domain name of the user's page
	ScreenName string `json:"screen_name,omitempty"`
	// User sex
	Sex int `json:"sex,omitempty"`
	// User relationship status
	Relation        int     `json:"relation,omitempty"`
	RelationPartner vk.User `json:"relation_partner,omitempty"`
	// Information whether relation status is pending
	RelationPending  int       `json:"relation_pending,omitempty"`
	RelationRequests []vk.User `json:"relation_requests,omitempty"`
	// User's date of birth
	Bdate string `json:"bdate,omitempty"`
	// Information whether user's birthdate are hidden
	BdateVisibility int `json:"bdate_visibility,omitempty"`
	// User's hometown
	HomeTown string        `json:"home_town,omitempty"`
	Country  vk.BaseObject `json:"country,omitempty"`
	City     vk.BaseObject `json:"city,omitempty"`
	// User status
	Status string `json:"status,omitempty"`
	// User phone number with some hidden digits
	Phone       string `json:"phone,omitempty"`
	NameRequest struct {
		// Request ID needed to cancel the request
		ID int `json:"id,omitempty"`
		// Request status
		Status string `json:"status,omitempty"`
		// First name in request
		FirstName string `json:"first_name,omitempty"`
		// Last name in request
		LastName string `json:"last_name,omitempty"`
	} `json:"name_request,omitempty"`
}

// GetProfileInfo Returns the current account info.
func (v Account) GetProfileInfo() (*AccountGetProfileInfoResponse, error) {
	r, err := v.API.Request("account.getProfileInfo", nil)
	if err != nil {
		return nil, err
	}

	var resp AccountGetProfileInfoResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AccountSaveProfileInfoParams are params for Account.SaveProfileInfo
type AccountSaveProfileInfoParams struct {
	// User first name.
	FirstName string `url:"first_name,omitempty"`
	// User last name.
	LastName string `url:"last_name,omitempty"`
	// User maiden name (female only)
	MaidenName string `url:"maiden_name,omitempty"`
	// User screen name.
	ScreenName string `url:"screen_name,omitempty"`
	// ID of the name change request to be canceled. If this parameter is sent, all the others are ignored.
	CancelRequestID int `url:"cancel_request_id,omitempty"`
	// User sex. Possible values: , * '1' – female,, * '2' – male.
	Sex int `url:"sex,omitempty"`
	// User relationship status. Possible values: , * '1' – single,, * '2' – in a relationship,, * '3' – engaged,, * '4' – married,, * '5' – it's complicated,, * '6' – actively searching,, * '7' – in love,, * '0' – not specified.
	Relation int `url:"relation,omitempty"`
	// ID of the relationship partner.
	RelationPartnerID int `url:"relation_partner_id,omitempty"`
	// User birth date, format: DD.MM.YYYY.
	Bdate string `url:"bdate,omitempty"`
	// Birth date visibility. Returned values: , * '1' – show birth date,, * '2' – show only month and day,, * '0' – hide birth date.
	BdateVisibility int `url:"bdate_visibility,omitempty"`
	// User home town.
	HomeTown string `url:"home_town,omitempty"`
	// User country.
	CountryID int `url:"country_id,omitempty"`
	// User city.
	CityID int `url:"city_id,omitempty"`
	// Status text.
	Status string `url:"status,omitempty"`
}

// AccountSaveProfileInfoResponse is response for Account.SaveProfileInfo
//easyjson:json
type AccountSaveProfileInfoResponse struct {
	// 1 if changes has been processed
	Changed     vk.BoolInt `json:"changed,omitempty"`
	NameRequest struct {
		// Request ID needed to cancel the request
		ID int `json:"id,omitempty"`
		// Request status
		Status string `json:"status,omitempty"`
		// First name in request
		FirstName string `json:"first_name,omitempty"`
		// Last name in request
		LastName string `json:"last_name,omitempty"`
	} `json:"name_request,omitempty"`
}

// SaveProfileInfo Edits current profile info.
func (v Account) SaveProfileInfo(params AccountSaveProfileInfoParams) (*AccountSaveProfileInfoResponse, error) {
	r, err := v.API.Request("account.saveProfileInfo", params)
	if err != nil {
		return nil, err
	}

	var resp AccountSaveProfileInfoResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AccountBanParams are params for Account.Ban
type AccountBanParams struct {
	OwnerID int `url:"owner_id,omitempty"`
}

// Ban does account.ban
func (v Account) Ban(params AccountBanParams) (bool, error) {
	r, err := v.API.Request("account.ban", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// AccountUnbanParams are params for Account.Unban
type AccountUnbanParams struct {
	OwnerID int `url:"owner_id,omitempty"`
}

// Unban does account.unban
func (v Account) Unban(params AccountUnbanParams) (bool, error) {
	r, err := v.API.Request("account.unban", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}
