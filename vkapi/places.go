package vkapi

import (
	"encoding/json"

	"github.com/stek29/vk"
)

// Places implements VK API namespace `places`
type Places struct {
	API vk.API
}

// PlacesAddParams are params for Places.Add
type PlacesAddParams struct {
	// ID of the location's type (e.g., '1' — Home, '2' — Work). To get location type IDs, use the [vk.com/dev/places.getTypes|places.getTypes] method.
	Type int `url:"type,omitempty"`
	// Title of the location.
	Title string `url:"title"`
	// Geographical latitude, in degrees (from '-90' to '90').
	Latitude float32 `url:"latitude"`
	// Geographical longitude, in degrees (from '-180' to '180').
	Longitude float32 `url:"longitude"`
	// ID of the location's country. To get country IDs, use the [vk.com/dev/database.getCountries|database.getCountries] method.
	Country int `url:"country,omitempty"`
	// ID of the location's city. To get city IDs, use the [vk.com/dev/database.getCities|database.getCities] method.
	City int `url:"city,omitempty"`
	// Street address of the location (e.g., '125 Elm Street').
	Address string `url:"address,omitempty"`
}

// PlacesAddResponse is response for Places.Add
//easyjson:json
type PlacesAddResponse struct {
	// Place ID
	ID int `json:"id,omitempty"`
}

// Add Adds a new location to the location database.
func (v Places) Add(params PlacesAddParams) (*PlacesAddResponse, error) {
	r, err := v.API.Request("places.add", params)
	if err != nil {
		return nil, err
	}

	var resp PlacesAddResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PlacesGetByIDParams are params for Places.GetByID
type PlacesGetByIDParams struct {
	// Location IDs.
	Places CSVIntSlice `url:"places"`
}

// PlacesGetByIDResponse is response for Places.GetByID
//easyjson:json
type PlacesGetByIDResponse []vk.Place

// GetByID Returns information about locations by their IDs.
func (v Places) GetByID(params PlacesGetByIDParams) (PlacesGetByIDResponse, error) {
	r, err := v.API.Request("places.getById", params)
	if err != nil {
		return nil, err
	}

	var resp PlacesGetByIDResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// PlacesSearchParams are params for Places.Search
type PlacesSearchParams struct {
	// Search query string.
	Q string `url:"q,omitempty"`
	// City ID.
	City int `url:"city,omitempty"`
	// Geographical latitude of the initial search point, in degrees (from '-90' to '90').
	Latitude float32 `url:"latitude"`
	// Geographical longitude of the initial search point, in degrees (from '-180' to '180').
	Longitude float32 `url:"longitude"`
	// Radius of the search zone: '1' — 100 m. (default), '2' — 800 m. '3' — 6 km. '4' — 50 km.
	Radius int `url:"radius,omitempty"`
	// Offset needed to return a specific subset of locations.
	Offset int `url:"offset,omitempty"`
	// Number of locations to return.
	Count int `url:"count,omitempty"`
}

// PlacesSearchResponse is response for Places.Search
//easyjson:json
type PlacesSearchResponse struct {
	// Total number
	Count int        `json:"count,omitempty"`
	Items []vk.Place `json:"items,omitempty"`
}

// Search Returns a list of locations that match the search criteria.
func (v Places) Search(params PlacesSearchParams) (*PlacesSearchResponse, error) {
	r, err := v.API.Request("places.search", params)
	if err != nil {
		return nil, err
	}

	var resp PlacesSearchResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PlacesCheckinParams are params for Places.Checkin
type PlacesCheckinParams struct {
	// Location ID.
	PlaceID int `url:"place_id,omitempty"`
	// Text of the comment on the check-in (255 characters maximum, line breaks not supported).
	Text string `url:"text,omitempty"`
	// Geographical latitude of the check-in, in degrees (from '-90' to '90').
	Latitude float32 `url:"latitude,omitempty"`
	// Geographical longitude of the check-in, in degrees (from '-180' to '180').
	Longitude float32 `url:"longitude,omitempty"`
	// '1' — Check-in will be available only for friends. '0' — Check-in will be available for all users (default).
	FriendsOnly bool `url:"friends_only,omitempty"`
	// List of services or websites (e.g., 'twitter', 'facebook') to which the check-in will be exported, if the user has set up the respective option.
	Services CSVStringSlice `url:"services,omitempty"`
}

// PlacesCheckinResponse is response for Places.Checkin
//easyjson:json
type PlacesCheckinResponse struct {
	// Checkin ID
	ID int `json:"id,omitempty"`
}

// Checkin Checks a user in at the specified location.
func (v Places) Checkin(params PlacesCheckinParams) (*PlacesCheckinResponse, error) {
	r, err := v.API.Request("places.checkin", params)
	if err != nil {
		return nil, err
	}

	var resp PlacesCheckinResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PlacesGetCheckinsParams are params for Places.GetCheckins
type PlacesGetCheckinsParams struct {
	// Geographical latitude of the initial search point, in degrees (from '-90' to '90').
	Latitude float32 `url:"latitude,omitempty"`
	// Geographical longitude of the initial search point, in degrees (from '-180' to '180').
	Longitude float32 `url:"longitude,omitempty"`
	// Location ID of check-ins to return. (Ignored if 'latitude' and 'longitude' are specified.)
	Place  int `url:"place,omitempty"`
	UserID int `url:"user_id,omitempty"`
	// Offset needed to return a specific subset of check-ins. (Ignored if 'timestamp' is not null.)
	Offset int `url:"offset,omitempty"`
	// Number of check-ins to return. (Ignored if 'timestamp' is not null.)
	Count int `url:"count,omitempty"`
	// Specifies that only those check-ins created after the specified timestamp will be returned.
	Timestamp int `url:"timestamp,omitempty"`
	// '1' — to return only check-ins with set geographical coordinates. (Ignored if 'latitude' and 'longitude' are not set.)
	FriendsOnly bool `url:"friends_only,omitempty"`
	// '1' — to return location information with the check-ins. (Ignored if 'place' is not set.),
	NeedPlaces bool `url:"need_places,omitempty"`
}

// PlacesGetCheckinsResponse is response for Places.GetCheckins
//easyjson:json
type PlacesGetCheckinsResponse struct {
	// Total number
	Count int `json:"count,omitempty"`
	Items []struct {
		// Checkin ID
		ID int `json:"id,omitempty"`
		// User ID
		UserID int `json:"user_id,omitempty"`
		// Date when the checkin has been added in Unixtime
		Date int `json:"date,omitempty"`
		// Place latitude
		Latitude float32 `json:"latitude,omitempty"`
		// Place longitude
		Longitude float32 `json:"longitude,omitempty"`
		// Place ID
		PlaceID int `json:"place_id,omitempty"`
		// Comment text
		Text string `json:"text,omitempty"`
		// Distance to the place
		Distance int `json:"distance,omitempty"`
		// Place title
		PlaceTitle string `json:"place_title,omitempty"`
		// Country ID
		PlaceCountry int `json:"place_country,omitempty"`
		// City ID
		PlaceCity int `json:"place_city,omitempty"`
		// Place type
		PlaceType string `json:"place_type,omitempty"`
		// URL of the place's icon
		PlaceIcon string `json:"place_icon,omitempty"`
	} `json:"items,omitempty"`
}

// GetCheckins Returns a list of user check-ins at locations according to the set parameters.
func (v Places) GetCheckins(params PlacesGetCheckinsParams) (*PlacesGetCheckinsResponse, error) {
	r, err := v.API.Request("places.getCheckins", params)
	if err != nil {
		return nil, err
	}

	var resp PlacesGetCheckinsResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PlacesGetTypesResponse is response for Places.GetTypes
//easyjson:json
type PlacesGetTypesResponse []struct {
	// Place type ID
	ID int `json:"id,omitempty"`
	// Place type title
	Title string `json:"title,omitempty"`
	// URL of the place's icon
	Icon string `json:"icon,omitempty"`
}

// GetTypes Returns a list of all types of locations.
func (v Places) GetTypes() (PlacesGetTypesResponse, error) {
	r, err := v.API.Request("places.getTypes", nil)
	if err != nil {
		return nil, err
	}

	var resp PlacesGetTypesResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
