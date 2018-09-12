package vkapi

import (
	"encoding/json"
	"strconv"

	"github.com/stek29/vk"
)

// Friends implements VK API namespace `friends`
type Friends struct {
	API vk.API
}

// FriendsGetParams are params for Friends.Get
type FriendsGetParams struct {
	// User ID. By default, the current user ID.
	UserID int `url:"user_id,omitempty"`
	// Sort order: , 'name' — by name (enabled only if the 'fields' parameter is used), 'hints' — by rating, similar to how friends are sorted in My friends section, , This parameter is available only for [vk.com/dev/standalone|desktop applications].
	Order string `url:"order,omitempty"`
	// ID of the friend list returned by the [vk.com/dev/friends.getLists|friends.getLists] method to be used as the source. This parameter is taken into account only when the uid parameter is set to the current user ID. This parameter is available only for [vk.com/dev/standalone|desktop applications].
	ListID int `url:"list_id,omitempty"`
	// Number of friends to return.
	Count int `url:"count,omitempty"`
	// Offset needed to return a specific subset of friends.
	Offset int `url:"offset,omitempty"`
	// Profile fields to return. Sample values: 'uid', 'first_name', 'last_name', 'nickname', 'sex', 'bdate' (birthdate), 'city', 'country', 'timezone', 'photo', 'photo_medium', 'photo_big', 'domain', 'has_mobile', 'rate', 'contacts', 'education'.
	Fields CSVStringSlice `url:"fields,omitempty"`
	// Case for declension of user name and surname: , 'nom' — nominative (default) , 'gen' — genitive , 'dat' — dative , 'acc' — accusative , 'ins' — instrumental , 'abl' — prepositional
	NameCase string `url:"name_case,omitempty"`
}

// FriendsGetResponse is response for Friends.Get
//easyjson:json
type FriendsGetResponse struct {
	// Total friends number
	Count int `json:"count,omitempty"`
	// User ID
	Items []int `json:"items,omitempty"`
}

// Get Returns a list of user IDs or detailed information about a user's friends.
func (v Friends) Get(params FriendsGetParams) (*FriendsGetResponse, error) {
	r, err := v.API.Request("friends.get", params)
	if err != nil {
		return nil, err
	}

	var resp FriendsGetResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// FriendsGetOnlineParams are params for Friends.GetOnline
type FriendsGetOnlineParams struct {
	// User ID.
	UserID int `url:"user_id,omitempty"`
	// Friend list ID. If this parameter is not set, information about all online friends is returned.
	ListID int `url:"list_id,omitempty"`
	// '1' — to return an additional 'online_mobile' field, '0' — (default),
	OnlineMobile bool `url:"online_mobile,omitempty"`
	// Sort order: 'random' — random order
	Order string `url:"order,omitempty"`
	// Number of friends to return.
	Count int `url:"count,omitempty"`
	// Offset needed to return a specific subset of friends.
	Offset int `url:"offset,omitempty"`
}

// FriendsGetOnlineResponse is response for Friends.GetOnline
// User ID
//easyjson:json
type FriendsGetOnlineResponse []int

// GetOnline Returns a list of user IDs of a user's friends who are online.
func (v Friends) GetOnline(params FriendsGetOnlineParams) (FriendsGetOnlineResponse, error) {
	r, err := v.API.Request("friends.getOnline", params)
	if err != nil {
		return nil, err
	}

	var resp FriendsGetOnlineResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// FriendsGetMutualParams are params for Friends.GetMutual
type FriendsGetMutualParams struct {
	// ID of the user whose friends will be checked against the friends of the user specified in 'target_uid'.
	SourceUID int `url:"source_uid,omitempty"`
	// ID of the user whose friends will be checked against the friends of the user specified in 'source_uid'.
	TargetUID int `url:"target_uid,omitempty"`
	// IDs of the users whose friends will be checked against the friends of the user specified in 'source_uid'.
	TargetUIDs CSVIntSlice `url:"target_uids,omitempty"`
	// Sort order: 'random' — random order
	Order string `url:"order,omitempty"`
	// Number of mutual friends to return.
	Count int `url:"count,omitempty"`
	// Offset needed to return a specific subset of mutual friends.
	Offset int `url:"offset,omitempty"`
}

// FriendsGetMutualResponse is response for Friends.GetMutual
// User ID
//easyjson:json
type FriendsGetMutualResponse []int

// GetMutual Returns a list of user IDs of the mutual friends of two users.
func (v Friends) GetMutual(params FriendsGetMutualParams) (FriendsGetMutualResponse, error) {
	r, err := v.API.Request("friends.getMutual", params)
	if err != nil {
		return nil, err
	}

	var resp FriendsGetMutualResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// FriendsGetRecentParams are params for Friends.GetRecent
type FriendsGetRecentParams struct {
	// Number of recently added friends to return.
	Count int `url:"count,omitempty"`
}

// FriendsGetRecentResponse is response for Friends.GetRecent
// User ID
//easyjson:json
type FriendsGetRecentResponse []int

// GetRecent Returns a list of user IDs of the current user's recently added friends.
func (v Friends) GetRecent(params FriendsGetRecentParams) (FriendsGetRecentResponse, error) {
	r, err := v.API.Request("friends.getRecent", params)
	if err != nil {
		return nil, err
	}

	var resp FriendsGetRecentResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// FriendsGetRequestsParams are params for Friends.GetRequests
type FriendsGetRequestsParams struct {
	// Offset needed to return a specific subset of friend requests.
	Offset int `url:"offset,omitempty"`
	// Number of friend requests to return (default 100, maximum 1000).
	Count int `url:"count,omitempty"`
	// '1' — to return response messages from users who have sent a friend request or, if 'suggested' is set to '1', to return a list of suggested friends
	Extended bool `url:"extended,omitempty"`
	// '1' — to return a list of mutual friends (up to 20), if any
	NeedMutual bool `url:"need_mutual,omitempty"`
	// '1' — to return outgoing requests, '0' — to return incoming requests (default)
	Out bool `url:"out,omitempty"`
	// Sort order: '1' — by number of mutual friends, '0' — by date
	Sort int `url:"sort,omitempty"`
	// '1' — to return a list of suggested friends, '0' — to return friend requests (default)
	Suggested bool `url:"suggested,omitempty"`
}

// FriendsGetRequestsResponse is response for Friends.GetRequests
// Either FriendsGetRequestsResponseNormal or FriendsGetRequestsResponseExtended, depending on Extended flag
type FriendsGetRequestsResponse interface {
	isFriendsGetRequests()
}

// FriendsGetRequestsResponseNormal is non-extended version of FriendsGetRequestsResponse
//easyjson:json
type FriendsGetRequestsResponseNormal struct {
	// Total requests number
	Count int `json:"count,omitempty"`
	// User ID
	Items []int `json:"items,omitempty"`
	// Total unread requests number
	CountUnread int `json:"count_unread,omitempty"`
}

func (FriendsGetRequestsResponseNormal) isFriendsGetRequests() {}

// FriendsGetRequestsResponseExtended is extended version of FriendsGetRequestsResponse
//easyjson:json
type FriendsGetRequestsResponseExtended struct {
	// Total requests number
	Count int `json:"count,omitempty"`
	Items []genTODOType/* objects.json#/definitions/friends_requests_xtr_message */ `json:"items,omitempty"`
}

func (FriendsGetRequestsResponseExtended) isFriendsGetRequests() {}

// GetRequests Returns information about the current user's incoming and outgoing friend requests.
func (v Friends) GetRequests(params FriendsGetRequestsParams) (FriendsGetRequestsResponse, error) {
	r, err := v.API.Request("friends.getRequests", params)
	if err != nil {
		return nil, err
	}

	var resp FriendsGetRequestsResponse
	if params.Extended {
		var tmp FriendsGetRequestsResponseExtended
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	} else {
		var tmp FriendsGetRequestsResponseNormal
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// FriendsAddParams are params for Friends.Add
type FriendsAddParams struct {
	// ID of the user whose friend request will be approved or to whom a friend request will be sent.
	UserID int `url:"user_id"`
	// Text of the message (up to 500 characters) for the friend request, if any.
	Text string `url:"text,omitempty"`
	// '1' to pass an incoming request to followers list.
	Follow bool `url:"follow,omitempty"`
}

// FriendsAddResponse is response for Friends.Add
// Friend request status
type FriendsAddResponse int

// Add Approves or creates a friend request.
func (v Friends) Add(params FriendsAddParams) (FriendsAddResponse, error) {
	r, err := v.API.Request("friends.add", params)
	if err != nil {
		return 0, err
	}

	var resp FriendsAddResponse

	var cnv int
	cnv, err = strconv.Atoi(string(r))
	resp = FriendsAddResponse(cnv)

	if err != nil {
		return 0, err
	}
	return resp, nil
}

// FriendsEditParams are params for Friends.Edit
type FriendsEditParams struct {
	// ID of the user whose friend list is to be edited.
	UserID int `url:"user_id"`
	// IDs of the friend lists to which to add the user.
	ListIDs CSVIntSlice `url:"list_ids,omitempty"`
}

// Edit Edits the friend lists of the selected user.
func (v Friends) Edit(params FriendsEditParams) (bool, error) {
	r, err := v.API.Request("friends.edit", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// FriendsDeleteParams are params for Friends.Delete
type FriendsDeleteParams struct {
	// ID of the user whose friend request is to be declined or who is to be deleted from the current user's friend list.
	UserID int `url:"user_id"`
}

// FriendsDeleteResponse is response for Friends.Delete
//easyjson:json
type FriendsDeleteResponse struct {
	Success vk.BoolInt `json:"success,omitempty"`
	// Returns 1 if friend has been deleted
	FriendDeleted int `json:"friend_deleted,omitempty"`
	// Returns 1 if out request has been canceled
	OutRequestDeleted int `json:"out_request_deleted,omitempty"`
	// Returns 1 if incoming request has been declined
	InRequestDeleted int `json:"in_request_deleted,omitempty"`
	// Returns 1 if suggestion has been declined
	SuggestionDeleted int `json:"suggestion_deleted,omitempty"`
}

// Delete Declines a friend request or deletes a user from the current user's friend list.
func (v Friends) Delete(params FriendsDeleteParams) (*FriendsDeleteResponse, error) {
	r, err := v.API.Request("friends.delete", params)
	if err != nil {
		return nil, err
	}

	var resp FriendsDeleteResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// FriendsGetListsParams are params for Friends.GetLists
type FriendsGetListsParams struct {
	// User ID.
	UserID int `url:"user_id,omitempty"`
	// '1' — to return system friend lists. By default: '0'.
	ReturnSystem bool `url:"return_system,omitempty"`
}

// FriendsGetListsResponse is response for Friends.GetLists
//easyjson:json
type FriendsGetListsResponse struct {
	// Total communities number
	Count int `json:"count,omitempty"`
	Items []struct {
		// List title
		Name string `json:"name,omitempty"`
		// List ID
		ID int `json:"id,omitempty"`
	} `json:"items,omitempty"`
}

// GetLists Returns a list of the user's friend lists.
func (v Friends) GetLists(params FriendsGetListsParams) (*FriendsGetListsResponse, error) {
	r, err := v.API.Request("friends.getLists", params)
	if err != nil {
		return nil, err
	}

	var resp FriendsGetListsResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// FriendsAddListParams are params for Friends.AddList
type FriendsAddListParams struct {
	// Name of the friend list.
	Name string `url:"name"`
	// IDs of users to be added to the friend list.
	UserIDs CSVIntSlice `url:"user_ids,omitempty"`
}

// FriendsAddListResponse is response for Friends.AddList
//easyjson:json
type FriendsAddListResponse struct {
	// List ID
	ListID int `json:"list_id,omitempty"`
}

// AddList Creates a new friend list for the current user.
func (v Friends) AddList(params FriendsAddListParams) (*FriendsAddListResponse, error) {
	r, err := v.API.Request("friends.addList", params)
	if err != nil {
		return nil, err
	}

	var resp FriendsAddListResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// FriendsEditListParams are params for Friends.EditList
type FriendsEditListParams struct {
	// Name of the friend list.
	Name string `url:"name,omitempty"`
	// Friend list ID.
	ListID int `url:"list_id"`
	// IDs of users in the friend list.
	UserIDs CSVIntSlice `url:"user_ids,omitempty"`
	// (Applies if 'user_ids' parameter is not set.), User IDs to add to the friend list.
	AddUserIDs CSVIntSlice `url:"add_user_ids,omitempty"`
	// (Applies if 'user_ids' parameter is not set.), User IDs to delete from the friend list.
	DeleteUserIDs CSVIntSlice `url:"delete_user_ids,omitempty"`
}

// EditList Edits a friend list of the current user.
func (v Friends) EditList(params FriendsEditListParams) (bool, error) {
	r, err := v.API.Request("friends.editList", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// FriendsDeleteListParams are params for Friends.DeleteList
type FriendsDeleteListParams struct {
	// ID of the friend list to delete.
	ListID int `url:"list_id"`
}

// DeleteList Deletes a friend list of the current user.
func (v Friends) DeleteList(params FriendsDeleteListParams) (bool, error) {
	r, err := v.API.Request("friends.deleteList", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// FriendsGetAppUsersResponse is response for Friends.GetAppUsers
// User ID
//easyjson:json
type FriendsGetAppUsersResponse []int

// GetAppUsers Returns a list of IDs of the current user's friends who installed the application.
func (v Friends) GetAppUsers() (FriendsGetAppUsersResponse, error) {
	r, err := v.API.Request("friends.getAppUsers", nil)
	if err != nil {
		return nil, err
	}

	var resp FriendsGetAppUsersResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// FriendsGetByPhonesParams are params for Friends.GetByPhones
type FriendsGetByPhonesParams struct {
	// List of phone numbers in MSISDN format (maximum 1000). Example: "+79219876543,+79111234567"
	Phones CSVStringSlice `url:"phones,omitempty"`
	// Profile fields to return. Sample values: 'nickname', 'screen_name', 'sex', 'bdate' (birthdate), 'city', 'country', 'timezone', 'photo', 'photo_medium', 'photo_big', 'has_mobile', 'rate', 'contacts', 'education', 'online, counters'.
	Fields CSVStringSlice `url:"fields,omitempty"`
}

// FriendsGetByPhonesResponse is response for Friends.GetByPhones
//easyjson:json
type FriendsGetByPhonesResponse []vk.User

// GetByPhones Returns a list of the current user's friends whose phone numbers, validated or specified in a profile, are in a given list.
func (v Friends) GetByPhones(params FriendsGetByPhonesParams) (FriendsGetByPhonesResponse, error) {
	r, err := v.API.Request("friends.getByPhones", params)
	if err != nil {
		return nil, err
	}

	var resp FriendsGetByPhonesResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DeleteAllRequests Marks all incoming friend requests as viewed.
func (v Friends) DeleteAllRequests() (bool, error) {
	r, err := v.API.Request("friends.deleteAllRequests", nil)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// FriendsGetSuggestionsParams are params for Friends.GetSuggestions
type FriendsGetSuggestionsParams struct {
	// Types of potential friends to return: 'mutual' — users with many mutual friends , 'contacts' — users found with the [vk.com/dev/account.importContacts|account.importContacts] method , 'mutual_contacts' — users who imported the same contacts as the current user with the [vk.com/dev/account.importContacts|account.importContacts] method
	Filter CSVStringSlice `url:"filter,omitempty"`
	// Number of suggestions to return.
	Count int `url:"count,omitempty"`
	// Offset needed to return a specific subset of suggestions.
	Offset int `url:"offset,omitempty"`
	// Profile fields to return. Sample values: 'nickname', 'screen_name', 'sex', 'bdate' (birthdate), 'city', 'country', 'timezone', 'photo', 'photo_medium', 'photo_big', 'has_mobile', 'rate', 'contacts', 'education', 'online', 'counters'.
	Fields CSVStringSlice `url:"fields,omitempty"`
	// Case for declension of user name and surname: , 'nom' — nominative (default) , 'gen' — genitive , 'dat' — dative , 'acc' — accusative , 'ins' — instrumental , 'abl' — prepositional
	NameCase string `url:"name_case,omitempty"`
}

// FriendsGetSuggestionsResponse is response for Friends.GetSuggestions
//easyjson:json
type FriendsGetSuggestionsResponse struct {
	// Total results number
	Count int       `json:"count,omitempty"`
	Items []vk.User `json:"items,omitempty"`
}

// GetSuggestions Returns a list of profiles of users whom the current user may know.
func (v Friends) GetSuggestions(params FriendsGetSuggestionsParams) (*FriendsGetSuggestionsResponse, error) {
	r, err := v.API.Request("friends.getSuggestions", params)
	if err != nil {
		return nil, err
	}

	var resp FriendsGetSuggestionsResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// FriendsAreFriendsParams are params for Friends.AreFriends
type FriendsAreFriendsParams struct {
	// IDs of the users whose friendship status to check.
	UserIDs CSVIntSlice `url:"user_ids"`
	// '1' — to return 'sign' field. 'sign' is md5("{id}_{user_id}_{friends_status}_{application_secret}"), where id is current user ID. This field allows to check that data has not been modified by the client. By default: '0'.
	NeedSign bool `url:"need_sign,omitempty"`
}

// FriendsAreFriendsResponse is response for Friends.AreFriends
//easyjson:json
type FriendsAreFriendsResponse []struct {
	// User ID
	UserID int `json:"user_id,omitempty"`
	// Friend status with the user
	FriendStatus int `json:"friend_status,omitempty"`
	// Message sent with request
	RequestMessage string `json:"request_message,omitempty"`
	// Information whether request is unviewed
	ReadState vk.BoolInt `json:"read_state,omitempty"`
	// MD5 hash for the result validation
	Sign string `json:"sign,omitempty"`
}

// AreFriends Checks the current user's friendship status with other specified users.
func (v Friends) AreFriends(params FriendsAreFriendsParams) (FriendsAreFriendsResponse, error) {
	r, err := v.API.Request("friends.areFriends", params)
	if err != nil {
		return nil, err
	}

	var resp FriendsAreFriendsResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// FriendsGetAvailableForCallParams are params for Friends.GetAvailableForCall
type FriendsGetAvailableForCallParams struct {
	// Profile fields to return. Sample values: 'uid', 'first_name', 'last_name', 'nickname', 'sex', 'bdate' (birthdate), 'city', 'country', 'timezone', 'photo', 'photo_medium', 'photo_big', 'domain', 'has_mobile', 'rate', 'contacts', 'education'.
	Fields CSVStringSlice `url:"fields,omitempty"`
	// Case for declension of user name and surname: , 'nom' — nominative (default) , 'gen' — genitive , 'dat' — dative , 'acc' — accusative , 'ins' — instrumental , 'abl' — prepositional
	NameCase string `url:"name_case,omitempty"`
}

// FriendsGetAvailableForCallResponse is response for Friends.GetAvailableForCall
//easyjson:json
type FriendsGetAvailableForCallResponse struct {
	// Total number
	Count int `json:"count,omitempty"`
	// User ID
	Items []int `json:"items,omitempty"`
}

// GetAvailableForCall Returns a list of friends who can be called by the current user.
func (v Friends) GetAvailableForCall(params FriendsGetAvailableForCallParams) (*FriendsGetAvailableForCallResponse, error) {
	r, err := v.API.Request("friends.getAvailableForCall", params)
	if err != nil {
		return nil, err
	}

	var resp FriendsGetAvailableForCallResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// FriendsSearchParams are params for Friends.Search
type FriendsSearchParams struct {
	// User ID.
	UserID int `url:"user_id"`
	// Search query string (e.g., 'Vasya Babich').
	Q string `url:"q,omitempty"`
	// Profile fields to return. Sample values: 'nickname', 'screen_name', 'sex', 'bdate' (birthdate), 'city', 'country', 'timezone', 'photo', 'photo_medium', 'photo_big', 'has_mobile', 'rate', 'contacts', 'education', 'online',
	Fields CSVStringSlice `url:"fields,omitempty"`
	// Case for declension of user name and surname: 'nom' — nominative (default), 'gen' — genitive , 'dat' — dative, 'acc' — accusative , 'ins' — instrumental , 'abl' — prepositional
	NameCase string `url:"name_case,omitempty"`
	// Offset needed to return a specific subset of friends.
	Offset int `url:"offset,omitempty"`
	// Number of friends to return.
	Count int `url:"count,omitempty"`
}

// FriendsSearchResponse is response for Friends.Search
//easyjson:json
type FriendsSearchResponse struct {
	// Total number
	Count int       `json:"count,omitempty"`
	Items []vk.User `json:"items,omitempty"`
}

// Search Returns a list of friends matching the search criteria.
func (v Friends) Search(params FriendsSearchParams) (*FriendsSearchResponse, error) {
	r, err := v.API.Request("friends.search", params)
	if err != nil {
		return nil, err
	}

	var resp FriendsSearchResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
