package vkCallbackApi

import (
	"strconv"

	"github.com/mailru/easyjson"
)

// APIGroups implements VK API namespace `groups`
type APIGroups struct {
	API *API
}

// GroupsIsMemberParams are params for APIGroups.IsMember
type GroupsIsMemberParams struct {
	// ID or screen name of the community.
	GroupID string `url:"group_id"`
	// User ID.
	UserID int `url:"user_id,omitempty"`
	// User IDs.
	UserIDs CSVIntSlice `url:"user_ids,omitempty"`
	// '1' — to return an extended response with additional fields. By default: '0'.
	Extended bool `url:"extended,omitempty"`
}

// GroupsIsMemberResponse is response for APIGroups.IsMember
// Either GroupsIsMemberResponseNormal or GroupsIsMemberResponseExtended, depending on Extended flag
type GroupsIsMemberResponse interface {
	isGroupsIsMember()
}

// GroupsIsMemberResponseNormal is non-extended version of GroupsIsMemberResponse
// Information whether user is a member of the group
type GroupsIsMemberResponseNormal int

func (GroupsIsMemberResponseNormal) isGroupsIsMember() {}

// GroupsIsMemberResponseExtended is extended version of GroupsIsMemberResponse
//easyjson:json
type GroupsIsMemberResponseExtended struct {
	// Information whether user is a member of the group
	Member BoolInt `json:"member,omitempty"`
	// Information whether user has been invited to the group
	Invitation BoolInt `json:"invitation,omitempty"`
	// Information whether user has sent request to the group
	Request BoolInt `json:"request,omitempty"`
}

func (GroupsIsMemberResponseExtended) isGroupsIsMember() {}

// IsMember Returns information specifying whether a user is a member of a community.
func (v APIGroups) IsMember(params GroupsIsMemberParams) (GroupsIsMemberResponse, error) {
	r, err := v.API.Request("groups.isMember", params)
	if err != nil {
		return nil, err
	}

	var resp GroupsIsMemberResponse
	if params.Extended {
		var tmp GroupsIsMemberResponseExtended
		err = easyjson.Unmarshal(r, &tmp)
		resp = &tmp
	} else {
		var tmp GroupsIsMemberResponseNormal

		var cnv int
		cnv, err = strconv.Atoi(string(r))
		tmp = GroupsIsMemberResponseNormal(cnv)

		resp = &tmp
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GroupsGetByIDParams are params for APIGroups.GetByID
type GroupsGetByIDParams struct {
	// IDs or screen names of communities.
	GroupIDs CSVStringSlice `url:"group_ids,omitempty"`
	// ID or screen name of the community.
	GroupID string `url:"group_id,omitempty"`
	// Group fields to return.
	Fields CSVStringSlice `url:"fields,omitempty"`
}

// GroupsGetByIDResponse is response for APIGroups.GetByID
//easyjson:json
type GroupsGetByIDResponse []Group

// GetByID Returns information about communities by their IDs.
func (v APIGroups) GetByID(params GroupsGetByIDParams) (GroupsGetByIDResponse, error) {
	r, err := v.API.Request("groups.getById", params)
	if err != nil {
		return nil, err
	}

	var resp GroupsGetByIDResponse
	err = easyjson.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GroupsGetParams are params for APIGroups.Get
type GroupsGetParams struct {
	// User ID.
	UserID int `url:"user_id,omitempty"`
	// '1' — to return complete information about a user's communities, '0' — to return a list of community IDs without any additional fields (default),
	Extended bool `url:"extended,omitempty"`
	// Types of communities to return: 'admin' — to return communities administered by the user , 'editor' — to return communities where the user is an administrator or editor, 'moder' — to return communities where the user is an administrator, editor, or moderator, 'groups' — to return only groups, 'publics' — to return only public pages, 'events' — to return only events
	Filter CSVStringSlice `url:"filter,omitempty"`
	// Profile fields to return.
	Fields CSVStringSlice `url:"fields,omitempty"`
	// Offset needed to return a specific subset of communities.
	Offset int `url:"offset,omitempty"`
	// Number of communities to return.
	Count int `url:"count,omitempty"`
}

// GroupsGetResponse is response for APIGroups.Get
// Either GroupsGetResponseNormal or GroupsGetResponseExtended, depending on Extended flag
type GroupsGetResponse interface {
	isGroupsGet()
}

// GroupsGetResponseNormal is non-extended version of GroupsGetResponse
//easyjson:json
type GroupsGetResponseNormal struct {
	// Total communities number
	Count int `json:"count,omitempty"`
	// Community ID
	Items []int `json:"items,omitempty"`
}

func (GroupsGetResponseNormal) isGroupsGet() {}

// GroupsGetResponseExtended is extended version of GroupsGetResponse
//easyjson:json
type GroupsGetResponseExtended struct {
	// Total communities number
	Count int     `json:"count,omitempty"`
	Items []Group `json:"items,omitempty"`
}

func (GroupsGetResponseExtended) isGroupsGet() {}

// Get Returns a list of the communities to which a user belongs.
func (v APIGroups) Get(params GroupsGetParams) (GroupsGetResponse, error) {
	r, err := v.API.Request("groups.get", params)
	if err != nil {
		return nil, err
	}

	var resp GroupsGetResponse
	if params.Extended {
		var tmp GroupsGetResponseExtended
		err = easyjson.Unmarshal(r, &tmp)
		resp = &tmp
	} else {
		var tmp GroupsGetResponseNormal
		err = easyjson.Unmarshal(r, &tmp)
		resp = &tmp
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GroupsGetMembersParams are params for APIGroups.GetMembers
type GroupsGetMembersParams struct {
	// ID or screen name of the community.
	GroupID string `url:"group_id,omitempty"`
	// Sort order. Available values: 'id_asc', 'id_desc', 'time_asc', 'time_desc'. 'time_asc' and 'time_desc' are availavle only if the method is called by the group's 'moderator'.
	Sort string `url:"sort,omitempty"`
	// Offset needed to return a specific subset of community members.
	Offset int `url:"offset,omitempty"`
	// Number of community members to return.
	Count int `url:"count,omitempty"`
	// List of additional fields to be returned. Available values: 'sex, bdate, city, country, photo_50, photo_100, photo_200_orig, photo_200, photo_400_orig, photo_max, photo_max_orig, online, online_mobile, lists, domain, has_mobile, contacts, connections, site, education, universities, schools, can_post, can_see_all_posts, can_see_audio, can_write_private_message, status, last_seen, common_count, relation, relatives, counters'.
	Fields CSVStringSlice `url:"fields,omitempty"`
	// *'friends' – only friends in this community will be returned,, *'unsure' – only those who pressed 'I may attend' will be returned (if it's an event).
	Filter string `url:"filter,omitempty"`
}

// GroupsGetMembersResponse is response for APIGroups.GetMembers
//easyjson:json
type GroupsGetMembersResponse struct {
	// Total members number
	Count int `json:"count,omitempty"`
	// User ID
	Items []int `json:"items,omitempty"`
}

// GetMembers Returns a list of community members.
func (v APIGroups) GetMembers(params GroupsGetMembersParams) (*GroupsGetMembersResponse, error) {
	r, err := v.API.Request("groups.getMembers", params)
	if err != nil {
		return nil, err
	}

	var resp GroupsGetMembersResponse
	err = easyjson.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GroupsJoinParams are params for APIGroups.Join
type GroupsJoinParams struct {
	// ID or screen name of the community.
	GroupID int `url:"group_id,omitempty"`
	// Optional parameter which is taken into account when 'gid' belongs to the event: '1' — Perhaps I will attend, '0' — I will be there for sure (default), ,
	NotSure string `url:"not_sure,omitempty"`
}

// Join With this method you can join the group or public page, and also confirm your participation in an event.
func (v APIGroups) Join(params GroupsJoinParams) (bool, error) {
	r, err := v.API.Request("groups.join", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// GroupsLeaveParams are params for APIGroups.Leave
type GroupsLeaveParams struct {
	// ID or screen name of the community.
	GroupID int `url:"group_id"`
}

// Leave With this method you can leave a group, public page, or event.
func (v APIGroups) Leave(params GroupsLeaveParams) (bool, error) {
	r, err := v.API.Request("groups.leave", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// GroupsSearchParams are params for APIGroups.Search
type GroupsSearchParams struct {
	// Search query string.
	Q string `url:"q"`
	// Community type. Possible values: 'group, page, event.'
	Type string `url:"type,omitempty"`
	// Country ID.
	CountryID int `url:"country_id,omitempty"`
	// City ID. If this parameter is transmitted, country_id is ignored.
	CityID int `url:"city_id,omitempty"`
	// '1' — to return only upcoming events. Works with the 'type' = 'event' only.
	Future bool `url:"future,omitempty"`
	// '1' — to return communities with enabled market only.
	Market bool `url:"market,omitempty"`
	// Sort order. Possible values: *'0' — default sorting (similar the full version of the site),, *'1' — by growth speed,, *'2'— by the "day attendance/members number" ratio,, *'3' — by the "Likes number/members number" ratio,, *'4' — by the "comments number/members number" ratio,, *'5' — by the "boards entries number/members number" ratio.
	Sort int `url:"sort,omitempty"`
	// Offset needed to return a specific subset of results.
	Offset int `url:"offset,omitempty"`
	// Number of communities to return. "Note that you can not receive more than first thousand of results, regardless of 'count' and 'offset' values."
	Count int `url:"count,omitempty"`
}

// GroupsSearchResponse is response for APIGroups.Search
//easyjson:json
type GroupsSearchResponse struct {
	// Total communities number
	Count int     `json:"count,omitempty"`
	Items []Group `json:"items,omitempty"`
}

// Search Returns a list of communities matching the search criteria.
func (v APIGroups) Search(params GroupsSearchParams) (*GroupsSearchResponse, error) {
	r, err := v.API.Request("groups.search", params)
	if err != nil {
		return nil, err
	}

	var resp GroupsSearchResponse
	err = easyjson.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GroupsGetCatalogParams are params for APIGroups.GetCatalog
type GroupsGetCatalogParams struct {
	// Category id received from [vk.com/dev/groups.getCatalogInfo|groups.getCatalogInfo].
	CategoryID int `url:"category_id,omitempty"`
	// Subcategory id received from [vk.com/dev/groups.getCatalogInfo|groups.getCatalogInfo].
	SubcategoryID int `url:"subcategory_id,omitempty"`
}

// GroupsGetCatalogResponse is response for APIGroups.GetCatalog
//easyjson:json
type GroupsGetCatalogResponse struct {
	// Total communities number
	Count int     `json:"count,omitempty"`
	Items []Group `json:"items,omitempty"`
}

// GetCatalog Returns communities list for a catalog category.
func (v APIGroups) GetCatalog(params GroupsGetCatalogParams) (*GroupsGetCatalogResponse, error) {
	r, err := v.API.Request("groups.getCatalog", params)
	if err != nil {
		return nil, err
	}

	var resp GroupsGetCatalogResponse
	err = easyjson.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GroupsGetCatalogInfoParams are params for APIGroups.GetCatalogInfo
type GroupsGetCatalogInfoParams struct {
	// 1 – to return communities count and three communities for preview. By default: 0.
	Extended bool `url:"extended,omitempty"`
	// 1 – to return subcategories info. By default: 0.
	Subcategories bool `url:"subcategories,omitempty"`
}

// GroupsGetCatalogInfoResponse is response for APIGroups.GetCatalogInfo
// Either GroupsGetCatalogInfoResponseNormal or GroupsGetCatalogInfoResponseExtended, depending on Extended flag
type GroupsGetCatalogInfoResponse interface {
	isGroupsGetCatalogInfo()
}

// GroupsGetCatalogInfoResponseNormal is non-extended version of GroupsGetCatalogInfoResponse
//easyjson:json
type GroupsGetCatalogInfoResponseNormal struct {
	// Information whether catalog is enabled for current user
	Enabled    int        `json:"enabled,omitempty"`
	Categories []Category `json:"categories,omitempty"`
}

func (GroupsGetCatalogInfoResponseNormal) isGroupsGetCatalogInfo() {}

// GroupsGetCatalogInfoResponseExtended is extended version of GroupsGetCatalogInfoResponse
//easyjson:json
type GroupsGetCatalogInfoResponseExtended struct {
	// Information whether catalog is enabled for current user
	Enabled    int `json:"enabled,omitempty"`
	Categories []struct {
		Category

		// Pages number
		PageCount    int     `json:"page_count,omitempty"`
		PagePreviews []Group `json:"page_previews,omitempty"`
	} `json:"categories,omitempty"`
}

func (GroupsGetCatalogInfoResponseExtended) isGroupsGetCatalogInfo() {}

// GetCatalogInfo Returns categories list for communities catalog
func (v APIGroups) GetCatalogInfo(params GroupsGetCatalogInfoParams) (GroupsGetCatalogInfoResponse, error) {
	r, err := v.API.Request("groups.getCatalogInfo", params)
	if err != nil {
		return nil, err
	}

	var resp GroupsGetCatalogInfoResponse
	if params.Extended {
		var tmp GroupsGetCatalogInfoResponseExtended
		err = easyjson.Unmarshal(r, &tmp)
		resp = &tmp
	} else {
		var tmp GroupsGetCatalogInfoResponseNormal
		err = easyjson.Unmarshal(r, &tmp)
		resp = &tmp
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GroupsGetInvitesParams are params for APIGroups.GetInvites
type GroupsGetInvitesParams struct {
	// Offset needed to return a specific subset of invitations.
	Offset int `url:"offset,omitempty"`
	// Number of invitations to return.
	Count int `url:"count,omitempty"`
	// '1' — to return additional [vk.com/dev/fields_groups|fields] for communities..
	Extended bool `url:"extended,omitempty"`
}

// GroupsGetInvitesResponse is response for APIGroups.GetInvites
// Either GroupsGetInvitesResponseNormal or GroupsGetInvitesResponseExtended, depending on Extended flag
type GroupsGetInvitesResponse interface {
	isGroupsGetInvites()
}

// GroupsGetInvitesResponseNormal is non-extended version of GroupsGetInvitesResponse
//easyjson:json
type GroupsGetInvitesResponseNormal struct {
	// Total communities number
	Count int     `json:"count,omitempty"`
	Items []Group `json:"items,omitempty"`
}

func (GroupsGetInvitesResponseNormal) isGroupsGetInvites() {}

// GroupsGetInvitesResponseExtended is extended version of GroupsGetInvitesResponse
//easyjson:json
type GroupsGetInvitesResponseExtended struct {
	// Total communities number
	Count    int     `json:"count,omitempty"`
	Items    []Group `json:"items,omitempty"`
	Profiles []User  `json:"profiles,omitempty"`
}

func (GroupsGetInvitesResponseExtended) isGroupsGetInvites() {}

// GetInvites Returns a list of invitations to join communities and events.
func (v APIGroups) GetInvites(params GroupsGetInvitesParams) (GroupsGetInvitesResponse, error) {
	r, err := v.API.Request("groups.getInvites", params)
	if err != nil {
		return nil, err
	}

	var resp GroupsGetInvitesResponse
	if params.Extended {
		var tmp GroupsGetInvitesResponseExtended
		err = easyjson.Unmarshal(r, &tmp)
		resp = &tmp
	} else {
		var tmp GroupsGetInvitesResponseNormal
		err = easyjson.Unmarshal(r, &tmp)
		resp = &tmp
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GroupsGetInvitedUsersParams are params for APIGroups.GetInvitedUsers
type GroupsGetInvitedUsersParams struct {
	// Group ID to return invited users for.
	GroupID int `url:"group_id"`
	// Offset needed to return a specific subset of results.
	Offset int `url:"offset,omitempty"`
	// Number of results to return.
	Count int `url:"count,omitempty"`
	// List of additional fields to be returned. Available values: 'sex, bdate, city, country, photo_50, photo_100, photo_200_orig, photo_200, photo_400_orig, photo_max, photo_max_orig, online, online_mobile, lists, domain, has_mobile, contacts, connections, site, education, universities, schools, can_post, can_see_all_posts, can_see_audio, can_write_private_message, status, last_seen, common_count, relation, relatives, counters'.
	Fields CSVStringSlice `url:"fields,omitempty"`
	// Case for declension of user name and surname. Possible values: *'nom' — nominative (default),, *'gen' — genitive,, *'dat' — dative,, *'acc' — accusative, , *'ins' — instrumental,, *'abl' — prepositional.
	NameCase string `url:"name_case,omitempty"`
}

// GroupsGetInvitedUsersResponse is response for APIGroups.GetInvitedUsers
//easyjson:json
type GroupsGetInvitedUsersResponse struct {
	// Total communities number
	Count int    `json:"count,omitempty"`
	Items []User `json:"items,omitempty"`
}

// GetInvitedUsers Returns invited users list of a community
func (v APIGroups) GetInvitedUsers(params GroupsGetInvitedUsersParams) (*GroupsGetInvitedUsersResponse, error) {
	r, err := v.API.Request("groups.getInvitedUsers", params)
	if err != nil {
		return nil, err
	}

	var resp GroupsGetInvitedUsersResponse
	err = easyjson.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GroupsBanParams are params for APIGroups.Ban
type GroupsBanParams struct {
	// Community ID.
	GroupID int `url:"group_id"`
	// User/Group ID to be banned.
	OwnerID int `url:"owner_id"`
	// Date (in Unix time) when the user will be removed from the blacklist.
	EndDate int `url:"end_date,omitempty"`
	// Reason for ban: '1' — spam, '2' — verbal abuse, '3' — strong language, '4' — irrelevant messages, '0' — other (default)
	Reason int `url:"reason,omitempty"`
	// Text of comment to ban.
	Comment string `url:"comment,omitempty"`
	// '1' — text of comment will be visible to the user,, '0' — text of comment will be invisible to the user. By default: '0'.
	CommentVisible bool `url:"comment_visible,omitempty"`
}

// Ban Adds a user or group to a community blacklist.
func (v APIGroups) Ban(params GroupsBanParams) (bool, error) {
	r, err := v.API.Request("groups.ban", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// GroupsUnbanParams are params for APIGroups.Unban
type GroupsUnbanParams struct {
	// Community ID.
	GroupID int `url:"group_id"`
	// User/Group ID to be banned.
	UserID int `url:"user_id"`
}

// Unban Removes a user or group from a community blacklist.
func (v APIGroups) Unban(params GroupsUnbanParams) (bool, error) {
	r, err := v.API.Request("groups.unban", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// GroupsGetBannedParams are params for APIGroups.GetBanned
type GroupsGetBannedParams struct {
	// Community ID.
	GroupID int `url:"group_id"`
	// Offset needed to return a specific subset of users.
	Offset int `url:"offset,omitempty"`
	// Number of users to return.
	Count  int            `url:"count,omitempty"`
	Fields CSVStringSlice `url:"fields,omitempty"`
	UserID int            `url:"user_id,omitempty"`
}

// GroupsGetBannedResponse is response for APIGroups.GetBanned
//easyjson:json
type GroupsGetBannedResponse struct {
	// Total users number
	Count int `json:"count,omitempty"`
	Items []struct {
		// Owner type
		Type string `json:"type,omitempty"`
		// Information about group if type = group
		Group Group `json:"group,omitempty"`
		// Information about group if type = profile
		Profile User `json:"profile,omitempty"`
		BanInfo struct {
			// Administrator ID
			AdminID int `json:"admin_id,omitempty"`
			// Date when user has been added to blacklist in Unixtime
			Date int `json:"date,omitempty"`
			// Ban reason
			Reason int `json:"reason,omitempty"`
			// Comment for a ban
			Comment string `json:"comment,omitempty"`
			// Date when user will be removed from blacklist in Unixtime
			EndDate int `json:"end_date,omitempty"`
		} `json:"ban_info,omitempty"`
	} `json:"items,omitempty"`
}

// GetBanned Returns a list of users on a community blacklist.
func (v APIGroups) GetBanned(params GroupsGetBannedParams) (*GroupsGetBannedResponse, error) {
	r, err := v.API.Request("groups.getBanned", params)
	if err != nil {
		return nil, err
	}

	var resp GroupsGetBannedResponse
	err = easyjson.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GroupsCreateParams are params for APIGroups.Create
type GroupsCreateParams struct {
	// Community title.
	Title string `url:"title"`
	// Community description (ignored for 'type' = 'public').
	Description string `url:"description,omitempty"`
	// Community type. Possible values: *'group' – group,, *'event' – event,, *'public' – public page
	Type string `url:"type,omitempty"`
	// Category ID (for 'type' = 'public' only).
	PublicCategory int `url:"public_category,omitempty"`
	// Public page subtype. Possible values: *'1' – place or small business,, *'2' – company, organization or website,, *'3' – famous person or group of people,, *'4' – product or work of art.
	Subtype int `url:"subtype,omitempty"`
}

// GroupsCreateResponse is response for APIGroups.Create
//easyjson:json
type GroupsCreateResponse Group

// Create Creates a new community.
func (v APIGroups) Create(params GroupsCreateParams) (*GroupsCreateResponse, error) {
	r, err := v.API.Request("groups.create", params)
	if err != nil {
		return nil, err
	}

	var resp GroupsCreateResponse
	err = easyjson.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GroupsEditParams are params for APIGroups.Edit
type GroupsEditParams struct {
	// Community ID.
	GroupID int `url:"group_id"`
	// Community title.
	Title string `url:"title,omitempty"`
	// Community description.
	Description string `url:"description,omitempty"`
	// Community screen name.
	ScreenName string `url:"screen_name,omitempty"`
	// Community type. Possible values: *'0' – open,, *'1' – closed,, *'2' – private.
	Access int `url:"access,omitempty"`
	// Website that will be displayed in the community information field.
	Website string `url:"website,omitempty"`
	// Community subject. Possible values: , *'1' – auto/moto,, *'2' – activity holidays,, *'3' – business,, *'4' – pets,, *'5' – health,, *'6' – dating and communication, , *'7' – games,, *'8' – IT (computers and software),, *'9' – cinema,, *'10' – beauty and fashion,, *'11' – cooking,, *'12' – art and culture,, *'13' – literature,, *'14' – mobile services and internet,, *'15' – music,, *'16' – science and technology,, *'17' – real estate,, *'18' – news and media,, *'19' – security,, *'20' – education,, *'21' – home and renovations,, *'22' – politics,, *'23' – food,, *'24' – industry,, *'25' – travel,, *'26' – work,, *'27' – entertainment,, *'28' – religion,, *'29' – family,, *'30' – sports,, *'31' – insurance,, *'32' – television,, *'33' – goods and services,, *'34' – hobbies,, *'35' – finance,, *'36' – photo,, *'37' – esoterics,, *'38' – electronics and appliances,, *'39' – erotic,, *'40' – humor,, *'41' – society, humanities,, *'42' – design and graphics.
	Subject string `url:"subject,omitempty"`
	// Organizer email (for events).
	Email string `url:"email,omitempty"`
	// Organizer phone number (for events).
	Phone string `url:"phone,omitempty"`
	// RSS feed address for import (available only to communities with special permission. Contact vk.com/support to get it.
	Rss string `url:"rss,omitempty"`
	// Event start date in Unixtime format.
	EventStartDate int `url:"event_start_date,omitempty"`
	// Event finish date in Unixtime format.
	EventFinishDate int `url:"event_finish_date,omitempty"`
	// Organizer community ID (for events only).
	EventGroupID int `url:"event_group_id,omitempty"`
	// Public page category ID.
	PublicCategory int `url:"public_category,omitempty"`
	// Public page subcategory ID.
	PublicSubcategory int `url:"public_subcategory,omitempty"`
	// Founding date of a company or organization owning the community in "dd.mm.YYYY" format.
	PublicDate string `url:"public_date,omitempty"`
	// Wall settings. Possible values: *'0' – disabled,, *'1' – open,, *'2' – limited (groups and events only),, *'3' – closed (groups and events only).
	Wall int `url:"wall,omitempty"`
	// Board topics settings. Possbile values: , *'0' – disabled,, *'1' – open,, *'2' – limited (for groups and events only).
	Topics int `url:"topics,omitempty"`
	// Photos settings. Possible values: *'0' – disabled,, *'1' – open,, *'2' – limited (for groups and events only).
	Photos int `url:"photos,omitempty"`
	// Video settings. Possible values: *'0' – disabled,, *'1' – open,, *'2' – limited (for groups and events only).
	Video int `url:"video,omitempty"`
	// Audio settings. Possible values: *'0' – disabled,, *'1' – open,, *'2' – limited (for groups and events only).
	Audio int `url:"audio,omitempty"`
	// Links settings (for public pages only). Possible values: *'0' – disabled,, *'1' – enabled.
	Links bool `url:"links,omitempty"`
	// Events settings (for public pages only). Possible values: *'0' – disabled,, *'1' – enabled.
	Events bool `url:"events,omitempty"`
	// Places settings (for public pages only). Possible values: *'0' – disabled,, *'1' – enabled.
	Places bool `url:"places,omitempty"`
	// Contacts settings (for public pages only). Possible values: *'0' – disabled,, *'1' – enabled.
	Contacts bool `url:"contacts,omitempty"`
	// Documents settings. Possible values: *'0' – disabled,, *'1' – open,, *'2' – limited (for groups and events only).
	Docs int `url:"docs,omitempty"`
	// Wiki pages settings. Possible values: *'0' – disabled,, *'1' – open,, *'2' – limited (for groups and events only).
	Wiki int `url:"wiki,omitempty"`
	// Community messages. Possible values: *'0' — disabled,, *'1' — enabled.
	Messages bool `url:"messages,omitempty"`
	// Community age limits. Possible values: *'1' — no limits,, *'2' — 16+,, *'3' — 18+.
	AgeLimits int `url:"age_limits,omitempty"`
	// Market settings. Possible values: *'0' – disabled,, *'1' – enabled.
	Market bool `url:"market,omitempty"`
	// market comments settings. Possible values: *'0' – disabled,, *'1' – enabled.
	MarketComments bool `url:"market_comments,omitempty"`
	// Market delivery countries.
	MarketCountry CSVIntSlice `url:"market_country,omitempty"`
	// Market delivery cities (if only one country is specified).
	MarketCity CSVIntSlice `url:"market_city,omitempty"`
	// Market currency settings. Possbile values: , *'643' – Russian rubles,, *'980' – Ukrainian hryvnia,, *'398' – Kazakh tenge,, *'978' – Euro,, *'840' – US dollars
	MarketCurrency int `url:"market_currency,omitempty"`
	// Seller contact for market. Set '0' for community messages.
	MarketContact int `url:"market_contact,omitempty"`
	// ID of a wiki page with market description.
	MarketWiki int `url:"market_wiki,omitempty"`
	// Obscene expressions filter in comments. Possible values: , *'0' – disabled,, *'1' – enabled.
	ObsceneFilter bool `url:"obscene_filter,omitempty"`
	// Stopwords filter in comments. Possible values: , *'0' – disabled,, *'1' – enabled.
	ObsceneStopwords bool `url:"obscene_stopwords,omitempty"`
	// Keywords for stopwords filter.
	ObsceneWords CSVStringSlice `url:"obscene_words,omitempty"`
}

// Edit Edits a community.
func (v APIGroups) Edit(params GroupsEditParams) (bool, error) {
	r, err := v.API.Request("groups.edit", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// GroupsEditPlaceParams are params for APIGroups.EditPlace
type GroupsEditPlaceParams struct {
	// Community ID.
	GroupID int `url:"group_id"`
	// Place title.
	Title string `url:"title,omitempty"`
	// Place address.
	Address string `url:"address,omitempty"`
	// Country ID.
	CountryID int `url:"country_id,omitempty"`
	// City ID.
	CityID int `url:"city_id,omitempty"`
	// Geographical latitude.
	Latitude float32 `url:"latitude,omitempty"`
	// Geographical longitude.
	Longitude float32 `url:"longitude,omitempty"`
}

// GroupsEditPlaceResponse is response for APIGroups.EditPlace
//easyjson:json
type GroupsEditPlaceResponse struct {
	Success BoolInt `json:"success,omitempty"`
	// Place address
	Address string `json:"address,omitempty"`
}

// EditPlace Edits the place in community.
func (v APIGroups) EditPlace(params GroupsEditPlaceParams) (*GroupsEditPlaceResponse, error) {
	r, err := v.API.Request("groups.editPlace", params)
	if err != nil {
		return nil, err
	}

	var resp GroupsEditPlaceResponse
	err = easyjson.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GroupsGetSettingsParams are params for APIGroups.GetSettings
type GroupsGetSettingsParams struct {
	// Community ID.
	GroupID int `url:"group_id"`
}

// GroupsGetSettingsResponse is response for APIGroups.GetSettings
//easyjson:json
type GroupsGetSettingsResponse struct {
	// Community title
	Title string `json:"title,omitempty"`
	// Community description
	Description string `json:"description,omitempty"`
	// Community's page domain
	Address string `json:"address,omitempty"`
	Place   Place  `json:"place,omitempty"`
	// Wall settings
	Wall int `json:"wall,omitempty"`
	// Photos settings
	Photos int `json:"photos,omitempty"`
	// Video settings
	Video int `json:"video,omitempty"`
	// Audio settings
	Audio int `json:"audio,omitempty"`
	// Docs settings
	Docs int `json:"docs,omitempty"`
	// Topics settings
	Topics int `json:"topics,omitempty"`
	// Wiki settings
	Wiki int `json:"wiki,omitempty"`
	// Information whether the obscene filter is enabled
	ObsceneFilter BoolInt `json:"obscene_filter,omitempty"`
	// Information about the group category
	PublicCategory int `json:"public_category,omitempty"`
	// Information about the group subcategory
	PublicSubcategory  int        `json:"public_subcategory,omitempty"`
	PublicCategoryList []Category `json:"public_category_list,omitempty"`
	// Information whether the stopwords filter is enabled
	ObsceneStopwords BoolInt `json:"obscene_stopwords,omitempty"`
	// The list of stop words
	ObsceneWords string `json:"obscene_words,omitempty"`
	// Community access settings
	Access int `json:"access,omitempty"`
	// Community subject ID
	Subject     int                  `json:"subject,omitempty"`
	SubjectList []BaseObjectWithName `json:"subject_list,omitempty"`
	// URL of the RSS feed
	Rss string `json:"rss,omitempty"`
	// Community website
	Website string `json:"website,omitempty"`
}

// GetSettings Returns community settings.
func (v APIGroups) GetSettings(params GroupsGetSettingsParams) (*GroupsGetSettingsResponse, error) {
	r, err := v.API.Request("groups.getSettings", params)
	if err != nil {
		return nil, err
	}

	var resp GroupsGetSettingsResponse
	err = easyjson.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GroupsGetRequestsParams are params for APIGroups.GetRequests
type GroupsGetRequestsParams struct {
	// Community ID.
	GroupID int `url:"group_id"`
	// Offset needed to return a specific subset of results.
	Offset int `url:"offset,omitempty"`
	// Number of results to return.
	Count int `url:"count,omitempty"`
	// Profile fields to return.
	Fields CSVStringSlice `url:"fields,omitempty"`
}

// GroupsGetRequestsResponse is response for APIGroups.GetRequests
//easyjson:json
type GroupsGetRequestsResponse struct {
	// Total communities number
	Count int `json:"count,omitempty"`
	// User ID
	Items []int `json:"items,omitempty"`
}

// GetRequests Returns a list of requests to the community.
func (v APIGroups) GetRequests(params GroupsGetRequestsParams) (*GroupsGetRequestsResponse, error) {
	r, err := v.API.Request("groups.getRequests", params)
	if err != nil {
		return nil, err
	}

	var resp GroupsGetRequestsResponse
	err = easyjson.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GroupsEditManagerParams are params for APIGroups.EditManager
type GroupsEditManagerParams struct {
	// Community ID.
	GroupID int `url:"group_id"`
	// User ID.
	UserID int `url:"user_id"`
	// Manager role. Possible values: *'moderator',, *'editor',, *'administrator'.
	Role string `url:"role,omitempty"`
	// '1' — to show the manager in Contacts block of the community.
	IsContact bool `url:"is_contact,omitempty"`
	// Position to show in Contacts block.
	ContactPosition string `url:"contact_position,omitempty"`
	// Contact phone.
	ContactPhone string `url:"contact_phone,omitempty"`
	// Contact e-mail.
	ContactEmail string `url:"contact_email,omitempty"`
}

// EditManager Allows to add, remove or edit the community manager.
func (v APIGroups) EditManager(params GroupsEditManagerParams) (bool, error) {
	r, err := v.API.Request("groups.editManager", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// GroupsInviteParams are params for APIGroups.Invite
type GroupsInviteParams struct {
	// Community ID.
	GroupID int `url:"group_id"`
	// User ID.
	UserID int `url:"user_id"`
}

// Invite Allows to invite friends to the community.
func (v APIGroups) Invite(params GroupsInviteParams) (bool, error) {
	r, err := v.API.Request("groups.invite", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// GroupsAddLinkParams are params for APIGroups.AddLink
type GroupsAddLinkParams struct {
	// Community ID.
	GroupID int `url:"group_id"`
	// Link URL.
	Link string `url:"link"`
	// Description text for the link.
	Text string `url:"text,omitempty"`
}

// AddLink Allows to add a link to the community.
func (v APIGroups) AddLink(params GroupsAddLinkParams) (bool, error) {
	r, err := v.API.Request("groups.addLink", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// GroupsDeleteLinkParams are params for APIGroups.DeleteLink
type GroupsDeleteLinkParams struct {
	// Community ID.
	GroupID int `url:"group_id"`
	// Link ID.
	LinkID int `url:"link_id"`
}

// DeleteLink Allows to delete a link from the community.
func (v APIGroups) DeleteLink(params GroupsDeleteLinkParams) (bool, error) {
	r, err := v.API.Request("groups.deleteLink", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// GroupsEditLinkParams are params for APIGroups.EditLink
type GroupsEditLinkParams struct {
	// Community ID.
	GroupID int `url:"group_id"`
	// Link ID.
	LinkID int `url:"link_id"`
	// New description text for the link.
	Text string `url:"text,omitempty"`
}

// EditLink Allows to edit a link in the community.
func (v APIGroups) EditLink(params GroupsEditLinkParams) (bool, error) {
	r, err := v.API.Request("groups.editLink", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// GroupsReorderLinkParams are params for APIGroups.ReorderLink
type GroupsReorderLinkParams struct {
	// Community ID.
	GroupID int `url:"group_id"`
	// Link ID.
	LinkID int `url:"link_id"`
	// ID of the link after which to place the link with 'link_id'.
	After int `url:"after,omitempty"`
}

// ReorderLink Allows to reorder links in the community.
func (v APIGroups) ReorderLink(params GroupsReorderLinkParams) (bool, error) {
	r, err := v.API.Request("groups.reorderLink", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// GroupsRemoveUserParams are params for APIGroups.RemoveUser
type GroupsRemoveUserParams struct {
	// Community ID.
	GroupID int `url:"group_id"`
	// User ID.
	UserID int `url:"user_id"`
}

// RemoveUser Removes a user from the community.
func (v APIGroups) RemoveUser(params GroupsRemoveUserParams) (bool, error) {
	r, err := v.API.Request("groups.removeUser", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// GroupsApproveRequestParams are params for APIGroups.ApproveRequest
type GroupsApproveRequestParams struct {
	// Community ID.
	GroupID int `url:"group_id"`
	// User ID.
	UserID int `url:"user_id"`
}

// ApproveRequest Allows to approve join request to the community.
func (v APIGroups) ApproveRequest(params GroupsApproveRequestParams) (bool, error) {
	r, err := v.API.Request("groups.approveRequest", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// GroupsGetCallbackConfirmationCodeParams are params for APIGroups.GetCallbackConfirmationCode
type GroupsGetCallbackConfirmationCodeParams struct {
	// Community ID.
	GroupID int `url:"group_id"`
}

// GroupsGetCallbackConfirmationCodeResponse is response for APIGroups.GetCallbackConfirmationCode
//easyjson:json
type GroupsGetCallbackConfirmationCodeResponse struct {
	// Confirmation code
	Code string `json:"code,omitempty"`
}

// GetCallbackConfirmationCode Returns Callback API confirmation code for the community.
func (v APIGroups) GetCallbackConfirmationCode(params GroupsGetCallbackConfirmationCodeParams) (*GroupsGetCallbackConfirmationCodeResponse, error) {
	r, err := v.API.Request("groups.getCallbackConfirmationCode", params)
	if err != nil {
		return nil, err
	}

	var resp GroupsGetCallbackConfirmationCodeResponse
	err = easyjson.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GroupsGetCallbackSettingsParams are params for APIGroups.GetCallbackSettings
type GroupsGetCallbackSettingsParams struct {
	// Community ID.
	GroupID int `url:"group_id"`
	// Server ID.
	ServerID int `url:"server_id,omitempty"`
}

// GroupsGetCallbackSettingsResponse is response for APIGroups.GetCallbackSettings
//easyjson:json
type GroupsGetCallbackSettingsResponse struct {
	// API version used for the events
	APIVersion string `json:"api_version,omitempty"`
	Events     struct {
		MessageNew           BoolInt `json:"message_new,omitempty"`
		MessageReply         BoolInt `json:"message_reply,omitempty"`
		PhotoNew             BoolInt `json:"photo_new,omitempty"`
		AudioNew             BoolInt `json:"audio_new,omitempty"`
		VideoNew             BoolInt `json:"video_new,omitempty"`
		WallReplyNew         BoolInt `json:"wall_reply_new,omitempty"`
		WallReplyEdit        BoolInt `json:"wall_reply_edit,omitempty"`
		WallReplyDelete      BoolInt `json:"wall_reply_delete,omitempty"`
		WallReplyRestore     BoolInt `json:"wall_reply_restore,omitempty"`
		WallPostNew          BoolInt `json:"wall_post_new,omitempty"`
		BoardPostNew         BoolInt `json:"board_post_new,omitempty"`
		BoardPostEdit        BoolInt `json:"board_post_edit,omitempty"`
		BoardPostRestore     BoolInt `json:"board_post_restore,omitempty"`
		BoardPostDelete      BoolInt `json:"board_post_delete,omitempty"`
		PhotoCommentNew      BoolInt `json:"photo_comment_new,omitempty"`
		PhotoCommentEdit     BoolInt `json:"photo_comment_edit,omitempty"`
		PhotoCommentDelete   BoolInt `json:"photo_comment_delete,omitempty"`
		PhotoCommentRestore  BoolInt `json:"photo_comment_restore,omitempty"`
		VideoCommentNew      BoolInt `json:"video_comment_new,omitempty"`
		VideoCommentEdit     BoolInt `json:"video_comment_edit,omitempty"`
		VideoCommentDelete   BoolInt `json:"video_comment_delete,omitempty"`
		VideoCommentRestore  BoolInt `json:"video_comment_restore,omitempty"`
		MarketCommentNew     BoolInt `json:"market_comment_new,omitempty"`
		MarketCommentEdit    BoolInt `json:"market_comment_edit,omitempty"`
		MarketCommentDelete  BoolInt `json:"market_comment_delete,omitempty"`
		MarketCommentRestore BoolInt `json:"market_comment_restore,omitempty"`
		PollVoteNew          BoolInt `json:"poll_vote_new,omitempty"`
		GroupJoin            BoolInt `json:"group_join,omitempty"`
		GroupLeave           BoolInt `json:"group_leave,omitempty"`
		GroupChangeSettings  BoolInt `json:"group_change_settings,omitempty"`
		GroupChangePhoto     BoolInt `json:"group_change_photo,omitempty"`
		GroupOfficersEdit    BoolInt `json:"group_officers_edit,omitempty"`
		MessageAllow         BoolInt `json:"message_allow,omitempty"`
		MessageDeny          BoolInt `json:"message_deny,omitempty"`
		WallRepost           BoolInt `json:"wall_repost,omitempty"`
		UserBlock            BoolInt `json:"user_block,omitempty"`
		UserUnblock          BoolInt `json:"user_unblock,omitempty"`
		MessagesEdit         BoolInt `json:"messages_edit,omitempty"`
		MessageTypingState   BoolInt `json:"message_typing_state,omitempty"`
		LeadFormsNew         BoolInt `json:"lead_forms_new,omitempty"`
	} `json:"events,omitempty"`
}

// GetCallbackSettings Returns [vk.com/dev/callback_api|Callback API] notifications settings.
func (v APIGroups) GetCallbackSettings(params GroupsGetCallbackSettingsParams) (*GroupsGetCallbackSettingsResponse, error) {
	r, err := v.API.Request("groups.getCallbackSettings", params)
	if err != nil {
		return nil, err
	}

	var resp GroupsGetCallbackSettingsResponse
	err = easyjson.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GroupsSetCallbackSettingsParams are params for APIGroups.SetCallbackSettings
type GroupsSetCallbackSettingsParams struct {
	// Community ID.
	GroupID int `url:"group_id"`
	// Server ID.
	ServerID int `url:"server_id,omitempty"`
	// A new incoming message has been received ('0' — disabled, '1' — enabled).
	MessageNew bool `url:"message_new,omitempty"`
	// A new outcoming message has been received ('0' — disabled, '1' — enabled).
	MessageReply bool `url:"message_reply,omitempty"`
	// Allowed messages notifications ('0' — disabled, '1' — enabled).
	MessageAllow bool `url:"message_allow,omitempty"`
	// Denied messages notifications ('0' — disabled, '1' — enabled).
	MessageDeny bool `url:"message_deny,omitempty"`
	// New photos notifications ('0' — disabled, '1' — enabled).
	PhotoNew bool `url:"photo_new,omitempty"`
	// New audios notifications ('0' — disabled, '1' — enabled).
	AudioNew bool `url:"audio_new,omitempty"`
	// New videos notifications ('0' — disabled, '1' — enabled).
	VideoNew bool `url:"video_new,omitempty"`
	// New wall replies notifications ('0' — disabled, '1' — enabled).
	WallReplyNew bool `url:"wall_reply_new,omitempty"`
	// Wall replies edited notifications ('0' — disabled, '1' — enabled).
	WallReplyEdit bool `url:"wall_reply_edit,omitempty"`
	// A wall comment has been deleted ('0' — disabled, '1' — enabled).
	WallReplyDelete bool `url:"wall_reply_delete,omitempty"`
	// A wall comment has been restored ('0' — disabled, '1' — enabled).
	WallReplyRestore bool `url:"wall_reply_restore,omitempty"`
	// New wall posts notifications ('0' — disabled, '1' — enabled).
	WallPostNew bool `url:"wall_post_new,omitempty"`
	// New wall posts notifications ('0' — disabled, '1' — enabled).
	WallRepost bool `url:"wall_repost,omitempty"`
	// New board posts notifications ('0' — disabled, '1' — enabled).
	BoardPostNew bool `url:"board_post_new,omitempty"`
	// Board posts edited notifications ('0' — disabled, '1' — enabled).
	BoardPostEdit bool `url:"board_post_edit,omitempty"`
	// Board posts restored notifications ('0' — disabled, '1' — enabled).
	BoardPostRestore bool `url:"board_post_restore,omitempty"`
	// Board posts deleted notifications ('0' — disabled, '1' — enabled).
	BoardPostDelete bool `url:"board_post_delete,omitempty"`
	// New comment to photo notifications ('0' — disabled, '1' — enabled).
	PhotoCommentNew bool `url:"photo_comment_new,omitempty"`
	// A photo comment has been edited ('0' — disabled, '1' — enabled).
	PhotoCommentEdit bool `url:"photo_comment_edit,omitempty"`
	// A photo comment has been deleted ('0' — disabled, '1' — enabled).
	PhotoCommentDelete bool `url:"photo_comment_delete,omitempty"`
	// A photo comment has been restored ('0' — disabled, '1' — enabled).
	PhotoCommentRestore bool `url:"photo_comment_restore,omitempty"`
	// New comment to video notifications ('0' — disabled, '1' — enabled).
	VideoCommentNew bool `url:"video_comment_new,omitempty"`
	// A video comment has been edited ('0' — disabled, '1' — enabled).
	VideoCommentEdit bool `url:"video_comment_edit,omitempty"`
	// A video comment has been deleted ('0' — disabled, '1' — enabled).
	VideoCommentDelete bool `url:"video_comment_delete,omitempty"`
	// A video comment has been restored ('0' — disabled, '1' — enabled).
	VideoCommentRestore bool `url:"video_comment_restore,omitempty"`
	// New comment to market item notifications ('0' — disabled, '1' — enabled).
	MarketCommentNew bool `url:"market_comment_new,omitempty"`
	// A market comment has been edited ('0' — disabled, '1' — enabled).
	MarketCommentEdit bool `url:"market_comment_edit,omitempty"`
	// A market comment has been deleted ('0' — disabled, '1' — enabled).
	MarketCommentDelete bool `url:"market_comment_delete,omitempty"`
	// A market comment has been restored ('0' — disabled, '1' — enabled).
	MarketCommentRestore bool `url:"market_comment_restore,omitempty"`
	// A vote in a public poll has been added ('0' — disabled, '1' — enabled).
	PollVoteNew bool `url:"poll_vote_new,omitempty"`
	// Joined community notifications ('0' — disabled, '1' — enabled).
	GroupJoin bool `url:"group_join,omitempty"`
	// Left community notifications ('0' — disabled, '1' — enabled).
	GroupLeave bool `url:"group_leave,omitempty"`
	// User added to community blacklist
	UserBlock bool `url:"user_block,omitempty"`
	// User removed from community blacklist
	UserUnblock bool `url:"user_unblock,omitempty"`
	// New form in lead forms
	LeadFormsNew bool `url:"lead_forms_new,omitempty"`
}

// SetCallbackSettings Allow to set notifications settings for group.
func (v APIGroups) SetCallbackSettings(params GroupsSetCallbackSettingsParams) (bool, error) {
	r, err := v.API.Request("groups.setCallbackSettings", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// GroupsGetLongPollServerParams are params for APIGroups.GetLongPollServer
type GroupsGetLongPollServerParams struct {
	// Community ID
	GroupID int `url:"group_id"`
}

// GroupsGetLongPollServerResponse is response for APIGroups.GetLongPollServer
//easyjson:json
type GroupsGetLongPollServerResponse struct {
	// Long Poll key
	Key string `json:"key,omitempty"`
	// Long Poll server address
	Server string `json:"server,omitempty"`
	// Number of the last event
	Ts int `json:"ts,omitempty"`
}

// GetLongPollServer Returns the data needed to query a Long Poll server for events
func (v APIGroups) GetLongPollServer(params GroupsGetLongPollServerParams) (*GroupsGetLongPollServerResponse, error) {
	r, err := v.API.Request("groups.getLongPollServer", params)
	if err != nil {
		return nil, err
	}

	var resp GroupsGetLongPollServerResponse
	err = easyjson.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GroupsGetLongPollSettingsParams are params for APIGroups.GetLongPollSettings
type GroupsGetLongPollSettingsParams struct {
	// Community ID.
	GroupID int `url:"group_id"`
}

// GroupsGetLongPollSettingsResponse is response for APIGroups.GetLongPollSettings
//easyjson:json
type GroupsGetLongPollSettingsResponse struct {
	// Shows whether Long Poll is enabled
	IsEnabled bool `json:"is_enabled,omitempty"`
	// API version used for the events
	APIVersion string `json:"api_version,omitempty"`
	Events     struct {
		MessageNew           BoolInt `json:"message_new,omitempty"`
		MessageReply         BoolInt `json:"message_reply,omitempty"`
		PhotoNew             BoolInt `json:"photo_new,omitempty"`
		AudioNew             BoolInt `json:"audio_new,omitempty"`
		VideoNew             BoolInt `json:"video_new,omitempty"`
		WallReplyNew         BoolInt `json:"wall_reply_new,omitempty"`
		WallReplyEdit        BoolInt `json:"wall_reply_edit,omitempty"`
		WallReplyDelete      BoolInt `json:"wall_reply_delete,omitempty"`
		WallReplyRestore     BoolInt `json:"wall_reply_restore,omitempty"`
		WallPostNew          BoolInt `json:"wall_post_new,omitempty"`
		BoardPostNew         BoolInt `json:"board_post_new,omitempty"`
		BoardPostEdit        BoolInt `json:"board_post_edit,omitempty"`
		BoardPostRestore     BoolInt `json:"board_post_restore,omitempty"`
		BoardPostDelete      BoolInt `json:"board_post_delete,omitempty"`
		PhotoCommentNew      BoolInt `json:"photo_comment_new,omitempty"`
		PhotoCommentEdit     BoolInt `json:"photo_comment_edit,omitempty"`
		PhotoCommentDelete   BoolInt `json:"photo_comment_delete,omitempty"`
		PhotoCommentRestore  BoolInt `json:"photo_comment_restore,omitempty"`
		VideoCommentNew      BoolInt `json:"video_comment_new,omitempty"`
		VideoCommentEdit     BoolInt `json:"video_comment_edit,omitempty"`
		VideoCommentDelete   BoolInt `json:"video_comment_delete,omitempty"`
		VideoCommentRestore  BoolInt `json:"video_comment_restore,omitempty"`
		MarketCommentNew     BoolInt `json:"market_comment_new,omitempty"`
		MarketCommentEdit    BoolInt `json:"market_comment_edit,omitempty"`
		MarketCommentDelete  BoolInt `json:"market_comment_delete,omitempty"`
		MarketCommentRestore BoolInt `json:"market_comment_restore,omitempty"`
		PollVoteNew          BoolInt `json:"poll_vote_new,omitempty"`
		GroupJoin            BoolInt `json:"group_join,omitempty"`
		GroupLeave           BoolInt `json:"group_leave,omitempty"`
		GroupChangeSettings  BoolInt `json:"group_change_settings,omitempty"`
		GroupChangePhoto     BoolInt `json:"group_change_photo,omitempty"`
		GroupOfficersEdit    BoolInt `json:"group_officers_edit,omitempty"`
		MessageAllow         BoolInt `json:"message_allow,omitempty"`
		MessageDeny          BoolInt `json:"message_deny,omitempty"`
		WallRepost           BoolInt `json:"wall_repost,omitempty"`
		UserBlock            BoolInt `json:"user_block,omitempty"`
		UserUnblock          BoolInt `json:"user_unblock,omitempty"`
		MessagesEdit         BoolInt `json:"messages_edit,omitempty"`
		MessageTypingState   BoolInt `json:"message_typing_state,omitempty"`
		LeadFormsNew         BoolInt `json:"lead_forms_new,omitempty"`
	} `json:"events,omitempty"`
}

// GetLongPollSettings Returns Long Poll notification settings
func (v APIGroups) GetLongPollSettings(params GroupsGetLongPollSettingsParams) (*GroupsGetLongPollSettingsResponse, error) {
	r, err := v.API.Request("groups.getLongPollSettings", params)
	if err != nil {
		return nil, err
	}

	var resp GroupsGetLongPollSettingsResponse
	err = easyjson.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GroupsSetLongPollSettingsParams are params for APIGroups.SetLongPollSettings
type GroupsSetLongPollSettingsParams struct {
	// Community ID.
	GroupID int `url:"group_id"`
	// Sets whether Long Poll is enabled ('0' — disabled, '1' — enabled).
	Enabled bool `url:"enabled,omitempty"`
	// A new incoming message has been received ('0' — disabled, '1' — enabled).
	MessageNew bool `url:"message_new,omitempty"`
	// A new outcoming message has been received ('0' — disabled, '1' — enabled).
	MessageReply bool `url:"message_reply,omitempty"`
	// A message has been edited ('0' — disabled, '1' — enabled).
	MessageEdit bool `url:"message_edit,omitempty"`
	// Allowed messages notifications ('0' — disabled, '1' — enabled).
	MessageAllow bool `url:"message_allow,omitempty"`
	// Denied messages notifications ('0' — disabled, '1' — enabled).
	MessageDeny bool `url:"message_deny,omitempty"`
	// New photos notifications ('0' — disabled, '1' — enabled).
	PhotoNew bool `url:"photo_new,omitempty"`
	// New audios notifications ('0' — disabled, '1' — enabled).
	AudioNew bool `url:"audio_new,omitempty"`
	// New videos notifications ('0' — disabled, '1' — enabled).
	VideoNew bool `url:"video_new,omitempty"`
	// New wall replies notifications ('0' — disabled, '1' — enabled).
	WallReplyNew bool `url:"wall_reply_new,omitempty"`
	// Wall replies edited notifications ('0' — disabled, '1' — enabled).
	WallReplyEdit bool `url:"wall_reply_edit,omitempty"`
	// A wall comment has been deleted ('0' — disabled, '1' — enabled).
	WallReplyDelete bool `url:"wall_reply_delete,omitempty"`
	// A wall comment has been restored ('0' — disabled, '1' — enabled).
	WallReplyRestore bool `url:"wall_reply_restore,omitempty"`
	// New wall posts notifications ('0' — disabled, '1' — enabled).
	WallPostNew bool `url:"wall_post_new,omitempty"`
	// New wall posts notifications ('0' — disabled, '1' — enabled).
	WallRepost bool `url:"wall_repost,omitempty"`
	// New board posts notifications ('0' — disabled, '1' — enabled).
	BoardPostNew bool `url:"board_post_new,omitempty"`
	// Board posts edited notifications ('0' — disabled, '1' — enabled).
	BoardPostEdit bool `url:"board_post_edit,omitempty"`
	// Board posts restored notifications ('0' — disabled, '1' — enabled).
	BoardPostRestore bool `url:"board_post_restore,omitempty"`
	// Board posts deleted notifications ('0' — disabled, '1' — enabled).
	BoardPostDelete bool `url:"board_post_delete,omitempty"`
	// New comment to photo notifications ('0' — disabled, '1' — enabled).
	PhotoCommentNew bool `url:"photo_comment_new,omitempty"`
	// A photo comment has been edited ('0' — disabled, '1' — enabled).
	PhotoCommentEdit bool `url:"photo_comment_edit,omitempty"`
	// A photo comment has been deleted ('0' — disabled, '1' — enabled).
	PhotoCommentDelete bool `url:"photo_comment_delete,omitempty"`
	// A photo comment has been restored ('0' — disabled, '1' — enabled).
	PhotoCommentRestore bool `url:"photo_comment_restore,omitempty"`
	// New comment to video notifications ('0' — disabled, '1' — enabled).
	VideoCommentNew bool `url:"video_comment_new,omitempty"`
	// A video comment has been edited ('0' — disabled, '1' — enabled).
	VideoCommentEdit bool `url:"video_comment_edit,omitempty"`
	// A video comment has been deleted ('0' — disabled, '1' — enabled).
	VideoCommentDelete bool `url:"video_comment_delete,omitempty"`
	// A video comment has been restored ('0' — disabled, '1' — enabled).
	VideoCommentRestore bool `url:"video_comment_restore,omitempty"`
	// New comment to market item notifications ('0' — disabled, '1' — enabled).
	MarketCommentNew bool `url:"market_comment_new,omitempty"`
	// A market comment has been edited ('0' — disabled, '1' — enabled).
	MarketCommentEdit bool `url:"market_comment_edit,omitempty"`
	// A market comment has been deleted ('0' — disabled, '1' — enabled).
	MarketCommentDelete bool `url:"market_comment_delete,omitempty"`
	// A market comment has been restored ('0' — disabled, '1' — enabled).
	MarketCommentRestore bool `url:"market_comment_restore,omitempty"`
	// A vote in a public poll has been added ('0' — disabled, '1' — enabled).
	PollVoteNew bool `url:"poll_vote_new,omitempty"`
	// Joined community notifications ('0' — disabled, '1' — enabled).
	GroupJoin bool `url:"group_join,omitempty"`
	// Left community notifications ('0' — disabled, '1' — enabled).
	GroupLeave bool `url:"group_leave,omitempty"`
	// User added to community blacklist
	UserBlock bool `url:"user_block,omitempty"`
	// User removed from community blacklist
	UserUnblock bool `url:"user_unblock,omitempty"`
}

// SetLongPollSettings Sets Long Poll notification settings
func (v APIGroups) SetLongPollSettings(params GroupsSetLongPollSettingsParams) (bool, error) {
	r, err := v.API.Request("groups.setLongPollSettings", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}
