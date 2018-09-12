package vkCallbackApi

import (
	"encoding/json"
)

// APIUsers implements VK API namespace `users`
type APIUsers struct {
	API *API
}

// UsersGetParams are params for APIUsers.Get
type UsersGetParams struct {
	// User IDs or screen names ('screen_name'). By default, current user ID.
	UserIDs CSVStringSlice `url:"user_ids,omitempty"`
	// Profile fields to return. Sample values: 'nickname', 'screen_name', 'sex', 'bdate' (birthdate), 'city', 'country', 'timezone', 'photo', 'photo_medium', 'photo_big', 'has_mobile', 'contacts', 'education', 'online', 'counters', 'relation', 'last_seen', 'activity', 'can_write_private_message', 'can_see_all_posts', 'can_post', 'universities',
	Fields CSVStringSlice `url:"fields,omitempty"`
	// Case for declension of user name and surname: 'nom' — nominative (default), 'gen' — genitive , 'dat' — dative, 'acc' — accusative , 'ins' — instrumental , 'abl' — prepositional
	NameCase string `url:"name_case,omitempty"`
}

// UsersGetResponse is response for APIUsers.Get
//easyjson:json
type UsersGetResponse []User

// Get Returns detailed information on users.
func (v APIUsers) Get(params UsersGetParams) (UsersGetResponse, error) {
	r, err := v.API.Request("users.get", params)
	if err != nil {
		return nil, err
	}

	var resp UsersGetResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UsersSearchParams are params for APIUsers.Search
type UsersSearchParams struct {
	// Search query string (e.g., 'Vasya Babich').
	Q string `url:"q,omitempty"`
	// Sort order: '1' — by date registered, '0' — by rating
	Sort int `url:"sort,omitempty"`
	// Offset needed to return a specific subset of users.
	Offset int `url:"offset,omitempty"`
	// Number of users to return.
	Count int `url:"count,omitempty"`
	// Profile fields to return. Sample values: 'nickname', 'screen_name', 'sex', 'bdate' (birthdate), 'city', 'country', 'timezone', 'photo', 'photo_medium', 'photo_big', 'has_mobile', 'rate', 'contacts', 'education', 'online',
	Fields CSVStringSlice `url:"fields,omitempty"`
	// City ID.
	City int `url:"city,omitempty"`
	// Country ID.
	Country int `url:"country,omitempty"`
	// City name in a string.
	Hometown string `url:"hometown,omitempty"`
	// ID of the country where the user graduated.
	UniversityCountry int `url:"university_country,omitempty"`
	// ID of the institution of higher education.
	University int `url:"university,omitempty"`
	// Year of graduation from an institution of higher education.
	UniversityYear int `url:"university_year,omitempty"`
	// Faculty ID.
	UniversityFaculty int `url:"university_faculty,omitempty"`
	// Chair ID.
	UniversityChair int `url:"university_chair,omitempty"`
	// '1' — female, '2' — male, '0' — any (default)
	Sex int `url:"sex,omitempty"`
	// Relationship status: '1' — Not married, '2' — In a relationship, '3' — Engaged, '4' — Married, '5' — It's complicated, '6' — Actively searching, '7' — In love
	Status int `url:"status,omitempty"`
	// Minimum age.
	AgeFrom int `url:"age_from,omitempty"`
	// Maximum age.
	AgeTo int `url:"age_to,omitempty"`
	// Day of birth.
	BirthDay int `url:"birth_day,omitempty"`
	// Month of birth.
	BirthMonth int `url:"birth_month,omitempty"`
	// Year of birth.
	BirthYear int `url:"birth_year,omitempty"`
	// '1' — online only, '0' — all users
	Online bool `url:"online,omitempty"`
	// '1' — with photo only, '0' — all users
	HasPhoto bool `url:"has_photo,omitempty"`
	// ID of the country where users finished school.
	SchoolCountry int `url:"school_country,omitempty"`
	// ID of the city where users finished school.
	SchoolCity  int `url:"school_city,omitempty"`
	SchoolClass int `url:"school_class,omitempty"`
	// ID of the school.
	School int `url:"school,omitempty"`
	// School graduation year.
	SchoolYear int `url:"school_year,omitempty"`
	// Users' religious affiliation.
	Religion string `url:"religion,omitempty"`
	// Users' interests.
	Interests string `url:"interests,omitempty"`
	// Name of the company where users work.
	Company string `url:"company,omitempty"`
	// Job position.
	Position string `url:"position,omitempty"`
	// ID of a community to search in communities.
	GroupID  int            `url:"group_id,omitempty"`
	FromList CSVStringSlice `url:"from_list,omitempty"`
}

// UsersSearchResponse is response for APIUsers.Search
//easyjson:json
type UsersSearchResponse struct {
	// Total number of available results
	Count int    `json:"count,omitempty"`
	Items []User `json:"items,omitempty"`
}

// Search Returns a list of users matching the search criteria.
func (v APIUsers) Search(params UsersSearchParams) (*UsersSearchResponse, error) {
	r, err := v.API.Request("users.search", params)
	if err != nil {
		return nil, err
	}

	var resp UsersSearchResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UsersIsAppUserParams are params for APIUsers.IsAppUser
type UsersIsAppUserParams struct {
	UserID int `url:"user_id,omitempty"`
}

// IsAppUser Returns information whether a user installed the application.
func (v APIUsers) IsAppUser(params UsersIsAppUserParams) (bool, error) {
	r, err := v.API.Request("users.isAppUser", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// UsersGetSubscriptionsParams are params for APIUsers.GetSubscriptions
type UsersGetSubscriptionsParams struct {
	// User ID.
	UserID int `url:"user_id,omitempty"`
	// '1' — to return a combined list of users and communities, '0' — to return separate lists of users and communities (default)
	Extended bool `url:"extended,omitempty"`
	// Offset needed to return a specific subset of subscriptions.
	Offset int `url:"offset,omitempty"`
	// Number of users and communities to return.
	Count  int            `url:"count,omitempty"`
	Fields CSVStringSlice `url:"fields,omitempty"`
}

// UsersGetSubscriptionsResponse is response for APIUsers.GetSubscriptions
// Either UsersGetSubscriptionsResponseNormal or UsersGetSubscriptionsResponseExtended, depending on Extended flag
type UsersGetSubscriptionsResponse interface {
	isUsersGetSubscriptions()
}

// UsersGetSubscriptionsResponseNormal is non-extended version of UsersGetSubscriptionsResponse
//easyjson:json
type UsersGetSubscriptionsResponseNormal struct {
	Users struct {
		// Users number
		Count int `json:"count,omitempty"`
		// User ID
		Items []int `json:"items,omitempty"`
	} `json:"users,omitempty"`
	Groups struct {
		// Communities number
		Count int `json:"count,omitempty"`
		// Community ID
		Items []int `json:"items,omitempty"`
	} `json:"groups,omitempty"`
}

func (UsersGetSubscriptionsResponseNormal) isUsersGetSubscriptions() {}

// UsersGetSubscriptionsResponseExtended is extended version of UsersGetSubscriptionsResponse
//easyjson:json
type UsersGetSubscriptionsResponseExtended struct {
	// Total number of available results
	Count int         `json:"count,omitempty"`
	Items genTODOType `json:"items,omitempty"`
}

func (UsersGetSubscriptionsResponseExtended) isUsersGetSubscriptions() {}

// GetSubscriptions Returns a list of IDs of users and communities followed by the user.
func (v APIUsers) GetSubscriptions(params UsersGetSubscriptionsParams) (UsersGetSubscriptionsResponse, error) {
	r, err := v.API.Request("users.getSubscriptions", params)
	if err != nil {
		return nil, err
	}

	var resp UsersGetSubscriptionsResponse
	if params.Extended {
		var tmp UsersGetSubscriptionsResponseExtended
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	} else {
		var tmp UsersGetSubscriptionsResponseNormal
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UsersGetFollowersParams are params for APIUsers.GetFollowers
type UsersGetFollowersParams struct {
	// User ID.
	UserID int `url:"user_id,omitempty"`
	// Offset needed to return a specific subset of followers.
	Offset int `url:"offset,omitempty"`
	// Number of followers to return.
	Count int `url:"count,omitempty"`
	// Profile fields to return. Sample values: 'nickname', 'screen_name', 'sex', 'bdate' (birthdate), 'city', 'country', 'timezone', 'photo', 'photo_medium', 'photo_big', 'has_mobile', 'rate', 'contacts', 'education', 'online'.
	Fields CSVStringSlice `url:"fields,omitempty"`
	// Case for declension of user name and surname: 'nom' — nominative (default), 'gen' — genitive , 'dat' — dative, 'acc' — accusative , 'ins' — instrumental , 'abl' — prepositional
	NameCase string `url:"name_case,omitempty"`
}

// UsersGetFollowersResponse is response for APIUsers.GetFollowers
//easyjson:json
type UsersGetFollowersResponse struct {
	// Total friends number
	Count int `json:"count,omitempty"`
	// User ID
	Items []int `json:"items,omitempty"`
}

// GetFollowers Returns a list of IDs of followers of the user in question, sorted by date added, most recent first.
func (v APIUsers) GetFollowers(params UsersGetFollowersParams) (*UsersGetFollowersResponse, error) {
	r, err := v.API.Request("users.getFollowers", params)
	if err != nil {
		return nil, err
	}

	var resp UsersGetFollowersResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UsersReportParams are params for APIUsers.Report
type UsersReportParams struct {
	// ID of the user about whom a complaint is being made.
	UserID int `url:"user_id"`
	// Type of complaint: 'porn' – pornography, 'spam' – spamming, 'insult' – abusive behavior, 'advertisment' – disruptive advertisements
	Type string `url:"type"`
	// Comment describing the complaint.
	Comment string `url:"comment,omitempty"`
}

// Report Reports (submits a complain about) a user.
func (v APIUsers) Report(params UsersReportParams) (bool, error) {
	r, err := v.API.Request("users.report", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// UsersGetNearbyParams are params for APIUsers.GetNearby
type UsersGetNearbyParams struct {
	// geographic latitude of the place a user is located, in degrees (from -90 to 90)
	Latitude float32 `url:"latitude"`
	// geographic longitude of the place a user is located, in degrees (from -180 to 180)
	Longitude float32 `url:"longitude"`
	// current location accuracy in meters
	Accuracy int `url:"accuracy,omitempty"`
	// time when a user disappears from location search results, in seconds
	Timeout int `url:"timeout,omitempty"`
	// search zone radius type (1 to 4), :* 1 – 300 m,, :* 2 – 2400 m,, :* 3 – 18 km,, :* 4 – 150 km.
	Radius int `url:"radius,omitempty"`
	// list of additional fields to return. Available values: sex, bdate, city, country, photo_50, photo_100, photo_200_orig, photo_200, photo_400_orig, photo_max, photo_max_orig, online, online_mobile, domain, has_mobile, contacts, connections, site, education, universities, schools, can_post, can_see_all_posts, can_see_audio, can_write_private_message, status, last_seen, common_count, relation, relatives, counters, screen_name, maiden_name, timezone, occupation
	Fields CSVStringSlice `url:"fields,omitempty"`
	// Case for declension of user name and surname: , nom –nominative (default) , gen – genitive , dat – dative , acc – accusative , ins – instrumental , abl – prepositional
	NameCase string `url:"name_case,omitempty"`
}

// UsersGetNearbyResponse is response for APIUsers.GetNearby
//easyjson:json
type UsersGetNearbyResponse struct {
	// Users number
	Count int    `json:"count,omitempty"`
	Items []User `json:"items,omitempty"`
}

// GetNearby Indexes current user location and returns nearby users.
func (v APIUsers) GetNearby(params UsersGetNearbyParams) (*UsersGetNearbyResponse, error) {
	r, err := v.API.Request("users.getNearby", params)
	if err != nil {
		return nil, err
	}

	var resp UsersGetNearbyResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
