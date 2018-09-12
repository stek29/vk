package vkCallbackApi

import (
	"encoding/json"
	"strconv"
)

// APIAccount implements VK API namespace `account`
type APIAccount struct {
	API *API
}

// AccountGetCountersParams are params for APIAccount.GetCounters
type AccountGetCountersParams struct {
	// Counters to be returned.
	Filter CSVStringSlice `url:"filter,omitempty"`
}

// AccountGetCountersResponse is response for APIAccount.GetCounters
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
func (v APIAccount) GetCounters(params AccountGetCountersParams) (*AccountGetCountersResponse, error) {
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

// AccountSetNameInMenuParams are params for APIAccount.SetNameInMenu
type AccountSetNameInMenuParams struct {
	// User ID.
	UserID int `url:"user_id"`
	// Application screen name.
	Name string `url:"name,omitempty"`
}

// SetNameInMenu Sets an application screen name (up to 17 characters), that is shown to the user in the left menu.
func (v APIAccount) SetNameInMenu(params AccountSetNameInMenuParams) (bool, error) {
	r, err := v.API.Request("account.setNameInMenu", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// AccountSetOnlineParams are params for APIAccount.SetOnline
type AccountSetOnlineParams struct {
	// '1' if videocalls are available for current device.
	Voip bool `url:"voip,omitempty"`
}

// SetOnline Marks the current user as online for 15 minutes.
func (v APIAccount) SetOnline(params AccountSetOnlineParams) (bool, error) {
	r, err := v.API.Request("account.setOnline", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// SetOffline Marks a current user as offline.
func (v APIAccount) SetOffline() (bool, error) {
	r, err := v.API.Request("account.setOffline", nil)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// AccountLookupContactsParams are params for APIAccount.LookupContacts
type AccountLookupContactsParams struct {
	// List of contacts separated with commas
	Contacts CSVStringSlice `url:"contacts,omitempty"`
	// String identifier of a service which contacts are used for searching. Possible values: , * email, * phone, * twitter, * facebook, * odnoklassniki, * instagram, * google
	Service string `url:"service"`
	// Contact of a current user on a specified service
	Mycontact string `url:"mycontact,omitempty"`
	// '1' – also return contacts found using this service before, '0' – return only contacts found using 'contacts' field.
	ReturnAll bool `url:"return_all,omitempty"`
	// Profile fields to return. Possible values: 'nickname, domain, sex, bdate, city, country, timezone, photo_50, photo_100, photo_200_orig, has_mobile, contacts, education, online, relation, last_seen, status, can_write_private_message, can_see_all_posts, can_post, universities'.
	Fields CSVStringSlice `url:"fields,omitempty"`
}

// AccountLookupContactsResponse is response for APIAccount.LookupContacts
//easyjson:json
type AccountLookupContactsResponse struct {
	Found []struct {
		User

		Contact     string `json:"contact,omitempty"`
		RequestSent int    `json:"request_sent,omitempty"`
		SortNum     int    `json:"sort_num,omitempty"`
	} `json:"found,omitempty"`
	Other []struct {
		// Contact
		Contact string `json:"contact,omitempty"`
		// Mutual friends count
		CommonCount int `json:"common_count,omitempty"`
	} `json:"other,omitempty"`
}

// LookupContacts Allows to search the VK users using phone numbers, e-mail addresses and user IDs on other services.
func (v APIAccount) LookupContacts(params AccountLookupContactsParams) (*AccountLookupContactsResponse, error) {
	r, err := v.API.Request("account.lookupContacts", params)
	if err != nil {
		return nil, err
	}

	var resp AccountLookupContactsResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AccountRegisterDeviceParams are params for APIAccount.RegisterDevice
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
}

// RegisterDevice Subscribes an iOS/Android/Windows Phone-based device to receive push notifications
func (v APIAccount) RegisterDevice(params AccountRegisterDeviceParams) (bool, error) {
	r, err := v.API.Request("account.registerDevice", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// AccountUnregisterDeviceParams are params for APIAccount.UnregisterDevice
type AccountUnregisterDeviceParams struct {
	// Unique device ID.
	DeviceID string `url:"device_id,omitempty"`
}

// UnregisterDevice Unsubscribes a device from push notifications.
func (v APIAccount) UnregisterDevice(params AccountUnregisterDeviceParams) (bool, error) {
	r, err := v.API.Request("account.unregisterDevice", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// AccountSetSilenceModeParams are params for APIAccount.SetSilenceMode
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
func (v APIAccount) SetSilenceMode(params AccountSetSilenceModeParams) (bool, error) {
	r, err := v.API.Request("account.setSilenceMode", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// AccountGetPushSettingsParams are params for APIAccount.GetPushSettings
type AccountGetPushSettingsParams struct {
	// Unique device ID.
	DeviceID string `url:"device_id,omitempty"`
}

// AccountGetPushSettingsResponse is response for APIAccount.GetPushSettings
//easyjson:json
type AccountGetPushSettingsResponse struct {
	// Information whether notifications are disabled
	Disabled BoolInt `json:"disabled,omitempty"`
	// Time until that notifications are disabled in Unixtime
	DisabledUntil int `json:"disabled_until,omitempty"`
	Conversations struct {
		// Items count
		Count int `json:"count,omitempty"`
		Items []struct {
			// Peer ID
			PeerID int `json:"peer_id,omitempty"`
			// Information whether the sound are enabled
			Sound BoolInt `json:"sound,omitempty"`
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
func (v APIAccount) GetPushSettings(params AccountGetPushSettingsParams) (*AccountGetPushSettingsResponse, error) {
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

// AccountSetPushSettingsParams are params for APIAccount.SetPushSettings
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
func (v APIAccount) SetPushSettings(params AccountSetPushSettingsParams) (bool, error) {
	r, err := v.API.Request("account.setPushSettings", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// AccountGetAppPermissionsParams are params for APIAccount.GetAppPermissions
type AccountGetAppPermissionsParams struct {
	// User ID whose settings information shall be got. By default: current user.
	UserID int `url:"user_id"`
}

// AccountGetAppPermissionsResponse is response for APIAccount.GetAppPermissions
// Permissions mask
type AccountGetAppPermissionsResponse int

// GetAppPermissions Gets settings of the user in this application.
func (v APIAccount) GetAppPermissions(params AccountGetAppPermissionsParams) (AccountGetAppPermissionsResponse, error) {
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

// AccountGetActiveOffersParams are params for APIAccount.GetActiveOffers
type AccountGetActiveOffersParams struct {
	// Number of results to return.
	Count int `url:"count,omitempty"`
}

// AccountGetActiveOffersResponse is response for APIAccount.GetActiveOffers
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
func (v APIAccount) GetActiveOffers(params AccountGetActiveOffersParams) (*AccountGetActiveOffersResponse, error) {
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

// AccountBanParams are params for APIAccount.Ban
type AccountBanParams struct {
	// User/Group ID to be banned.
	OwnerID int `url:"owner_id"`
}

// Ban Adds user or group to the banlist.
func (v APIAccount) Ban(params AccountBanParams) (bool, error) {
	r, err := v.API.Request("account.ban", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// AccountUnbanParams are params for APIAccount.Unban
type AccountUnbanParams struct {
	// User/Group ID to be unbanned.
	UserID int `url:"user_id"`
}

// Unban Deletes user or group from the blacklist.
func (v APIAccount) Unban(params AccountUnbanParams) (bool, error) {
	r, err := v.API.Request("account.unban", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// AccountGetBannedParams are params for APIAccount.GetBanned
type AccountGetBannedParams struct {
	// Offset needed to return a specific subset of results.
	Offset int `url:"offset,omitempty"`
	// Number of results to return.
	Count int `url:"count,omitempty"`
}

// AccountGetBannedResponse is response for APIAccount.GetBanned
//easyjson:json
type AccountGetBannedResponse struct {
	// Total number
	Count int    `json:"count,omitempty"`
	Items []User `json:"items,omitempty"`
}

// GetBanned Returns a user's blacklist.
func (v APIAccount) GetBanned(params AccountGetBannedParams) (*AccountGetBannedResponse, error) {
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

// AccountGetInfoParams are params for APIAccount.GetInfo
type AccountGetInfoParams struct {
	// Fields to return. Possible values: *'country' — user country,, *'https_required' — is "HTTPS only" option enabled,, *'own_posts_default' — is "Show my posts only" option is enabled,, *'no_wall_replies' — are wall replies disabled or not,, *'intro' — is intro passed by user or not,, *'lang' — user language. By default: all.
	Fields CSVStringSlice `url:"fields,omitempty"`
}

// AccountGetInfoResponse is response for APIAccount.GetInfo
//easyjson:json
type AccountGetInfoResponse struct {
	// Country code
	Country string `json:"country,omitempty"`
	// Information whether HTTPS-only is enabled
	HTTPSRequired BoolInt `json:"https_required,omitempty"`
	// Information whether only owners posts should be shown
	OwnPostsDefault BoolInt `json:"own_posts_default,omitempty"`
	// Information whether wall comments should be hidden
	NoWallReplies BoolInt `json:"no_wall_replies,omitempty"`
	// Information whether user has been processed intro
	Intro BoolInt `json:"intro,omitempty"`
	// Language ID
	Lang int `json:"lang,omitempty"`
	// Two factor authentication is enabled
	X2FaRequired BoolInt `json:"2fa_required,omitempty"`
}

// GetInfo Returns current account info.
func (v APIAccount) GetInfo(params AccountGetInfoParams) (*AccountGetInfoResponse, error) {
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

// AccountSetInfoParams are params for APIAccount.SetInfo
type AccountSetInfoParams struct {
	// Setting name.
	Name string `url:"name,omitempty"`
	// Setting value.
	Value string `url:"value,omitempty"`
}

// SetInfo Allows to edit the current account info.
func (v APIAccount) SetInfo(params AccountSetInfoParams) (bool, error) {
	r, err := v.API.Request("account.setInfo", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// AccountChangePasswordParams are params for APIAccount.ChangePassword
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

// AccountChangePasswordResponse is response for APIAccount.ChangePassword
//easyjson:json
type AccountChangePasswordResponse struct {
	// New token
	Token string `json:"token,omitempty"`
	// New secret
	Secret string `json:"secret,omitempty"`
}

// ChangePassword Changes a user password after access is successfully restored with the [vk.com/dev/auth.restore|auth.restore] method.
func (v APIAccount) ChangePassword(params AccountChangePasswordParams) (*AccountChangePasswordResponse, error) {
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

// AccountGetProfileInfoResponse is response for APIAccount.GetProfileInfo
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
	Relation        int  `json:"relation,omitempty"`
	RelationPartner User `json:"relation_partner,omitempty"`
	// Information whether relation status is pending
	RelationPending  int    `json:"relation_pending,omitempty"`
	RelationRequests []User `json:"relation_requests,omitempty"`
	// User's date of birth
	Bdate string `json:"bdate,omitempty"`
	// Information whether user's birthdate are hidden
	BdateVisibility int `json:"bdate_visibility,omitempty"`
	// User's hometown
	HomeTown string     `json:"home_town,omitempty"`
	Country  BaseObject `json:"country,omitempty"`
	City     BaseObject `json:"city,omitempty"`
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
func (v APIAccount) GetProfileInfo() (*AccountGetProfileInfoResponse, error) {
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

// AccountSaveProfileInfoParams are params for APIAccount.SaveProfileInfo
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

// AccountSaveProfileInfoResponse is response for APIAccount.SaveProfileInfo
//easyjson:json
type AccountSaveProfileInfoResponse struct {
	// 1 if changes has been processed
	Changed     BoolInt `json:"changed,omitempty"`
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
func (v APIAccount) SaveProfileInfo(params AccountSaveProfileInfoParams) (*AccountSaveProfileInfoResponse, error) {
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
