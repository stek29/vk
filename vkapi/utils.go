package vkapi

import (
	"encoding/json"
	"strconv"

	"github.com/stek29/vk"
)

// Utils implements VK API namespace `utils`
type Utils struct {
	API vk.API
}

// UtilsCheckLinkParams are params for Utils.CheckLink
type UtilsCheckLinkParams struct {
	// Link to check (e.g., 'http://google.com').
	URL string `url:"url"`
}

// UtilsCheckLinkResponse is response for Utils.CheckLink
//easyjson:json
type UtilsCheckLinkResponse struct {
	// Link status
	Status string `json:"status,omitempty"`
	// Link URL
	Link string `json:"link,omitempty"`
}

// CheckLink Checks whether a link is blocked in VK.
func (v Utils) CheckLink(params UtilsCheckLinkParams) (*UtilsCheckLinkResponse, error) {
	r, err := v.API.Request("utils.checkLink", params)
	if err != nil {
		return nil, err
	}

	var resp UtilsCheckLinkResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UtilsDeleteFromLastShortenedParams are params for Utils.DeleteFromLastShortened
type UtilsDeleteFromLastShortenedParams struct {
	// Link key (characters after vk.cc/).
	Key string `url:"key"`
}

// DeleteFromLastShortened Deletes shortened link from user's list.
func (v Utils) DeleteFromLastShortened(params UtilsDeleteFromLastShortenedParams) (bool, error) {
	r, err := v.API.Request("utils.deleteFromLastShortened", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// UtilsGetLastShortenedLinksParams are params for Utils.GetLastShortenedLinks
type UtilsGetLastShortenedLinksParams struct {
	// Number of links to return.
	Count int `url:"count,omitempty"`
	// Offset needed to return a specific subset of links.
	Offset int `url:"offset,omitempty"`
}

// UtilsGetLastShortenedLinksResponse is response for Utils.GetLastShortenedLinks
//easyjson:json
type UtilsGetLastShortenedLinksResponse struct {
	// Total number of available results
	Count int `json:"count,omitempty"`
	Items []struct {
		// Creation time in Unixtime
		Timestamp int `json:"timestamp,omitempty"`
		// Full URL
		URL string `json:"url,omitempty"`
		// Short link URL
		ShortURL string `json:"short_url,omitempty"`
		// Link key (characters after vk.cc/)
		Key string `json:"key,omitempty"`
		// Total views number
		Views int `json:"views,omitempty"`
		// Access key for private stats
		AccessKey string `json:"access_key,omitempty"`
	} `json:"items,omitempty"`
}

// GetLastShortenedLinks Returns a list of user's shortened links.
func (v Utils) GetLastShortenedLinks(params UtilsGetLastShortenedLinksParams) (*UtilsGetLastShortenedLinksResponse, error) {
	r, err := v.API.Request("utils.getLastShortenedLinks", params)
	if err != nil {
		return nil, err
	}

	var resp UtilsGetLastShortenedLinksResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UtilsGetLinkStatsParams are params for Utils.GetLinkStats
type UtilsGetLinkStatsParams struct {
	// Link key (characters after vk.cc/).
	Key string `url:"key"`
	// Access key for private link stats.
	AccessKey string `url:"access_key,omitempty"`
	// Interval.
	Interval string `url:"interval,omitempty"`
	// Number of intervals to return.
	IntervalsCount int `url:"intervals_count,omitempty"`
	// 1 — to return extended stats data (sex, age, geo). 0 — to return views number only.
	Extended bool `url:"extended,omitempty"`
}

// UtilsGetLinkStatsResponse is response for Utils.GetLinkStats
// Either UtilsGetLinkStatsResponseNormal or UtilsGetLinkStatsResponseExtended, depending on Extended flag
type UtilsGetLinkStatsResponse interface {
	isUtilsGetLinkStats()
}

// UtilsGetLinkStatsResponseNormal is non-extended version of UtilsGetLinkStatsResponse
//easyjson:json
type UtilsGetLinkStatsResponseNormal struct {
	// Link key (characters after vk.cc/)
	Key   string `json:"key,omitempty"`
	Stats []struct {
		// Start time
		Timestamp int `json:"timestamp,omitempty"`
		// Total views number
		Views int `json:"views,omitempty"`
	} `json:"stats,omitempty"`
}

func (UtilsGetLinkStatsResponseNormal) isUtilsGetLinkStats() {}

// UtilsGetLinkStatsResponseExtended is extended version of UtilsGetLinkStatsResponse
//easyjson:json
type UtilsGetLinkStatsResponseExtended struct {
	// Link key (characters after vk.cc/)
	Key   string `json:"key,omitempty"`
	Stats []struct {
		// Start time
		Timestamp int `json:"timestamp,omitempty"`
		// Total views number
		Views  int `json:"views,omitempty"`
		SexAge []struct {
			// Age denotation
			AgeRange string `json:"age_range,omitempty"`
			//  Views by female users
			Female int `json:"female,omitempty"`
			//  Views by male users
			Male int `json:"male,omitempty"`
		} `json:"sex_age,omitempty"`
		Countries []struct {
			// Country ID
			CountryID int `json:"country_id,omitempty"`
			// Views number
			Views int `json:"views,omitempty"`
		} `json:"countries,omitempty"`
		Cities []struct {
			// City ID
			CityID int `json:"city_id,omitempty"`
			// Views number
			Views int `json:"views,omitempty"`
		} `json:"cities,omitempty"`
	} `json:"stats,omitempty"`
}

func (UtilsGetLinkStatsResponseExtended) isUtilsGetLinkStats() {}

// GetLinkStats Returns stats data for shortened link.
func (v Utils) GetLinkStats(params UtilsGetLinkStatsParams) (UtilsGetLinkStatsResponse, error) {
	r, err := v.API.Request("utils.getLinkStats", params)
	if err != nil {
		return nil, err
	}

	var resp UtilsGetLinkStatsResponse
	if params.Extended {
		var tmp UtilsGetLinkStatsResponseExtended
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	} else {
		var tmp UtilsGetLinkStatsResponseNormal
		err = json.Unmarshal(r, &tmp)
		resp = &tmp
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UtilsGetShortLinkParams are params for Utils.GetShortLink
type UtilsGetShortLinkParams struct {
	// URL to be shortened.
	URL string `url:"url"`
	// 1 — private stats, 0 — public stats.
	Private bool `url:"private,omitempty"`
}

// UtilsGetShortLinkResponse is response for Utils.GetShortLink
//easyjson:json
type UtilsGetShortLinkResponse struct {
	// Short link URL
	ShortURL string `json:"short_url,omitempty"`
	// Access key for private stats
	AccessKey string `json:"access_key,omitempty"`
	// Link key (characters after vk.cc/)
	Key string `json:"key,omitempty"`
	// Full URL
	URL string `json:"url,omitempty"`
}

// GetShortLink Allows to receive a link shortened via vk.cc.
func (v Utils) GetShortLink(params UtilsGetShortLinkParams) (*UtilsGetShortLinkResponse, error) {
	r, err := v.API.Request("utils.getShortLink", params)
	if err != nil {
		return nil, err
	}

	var resp UtilsGetShortLinkResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UtilsResolveScreenNameParams are params for Utils.ResolveScreenName
type UtilsResolveScreenNameParams struct {
	// Screen name of the user, community (e.g., 'apiclub,' 'andrew', or 'rules_of_war'), or application.
	ScreenName string `url:"screen_name"`
}

// UtilsResolveScreenNameResponse is response for Utils.ResolveScreenName
//easyjson:json
type UtilsResolveScreenNameResponse struct {
	// Object type
	Type string `json:"type,omitempty"`
	// Object ID
	ObjectID int `json:"object_id,omitempty"`
}

// ResolveScreenName Detects a type of object (e.g., user, community, application) and its ID by screen name.
func (v Utils) ResolveScreenName(params UtilsResolveScreenNameParams) (*UtilsResolveScreenNameResponse, error) {
	r, err := v.API.Request("utils.resolveScreenName", params)
	if err != nil {
		return nil, err
	}

	var resp UtilsResolveScreenNameResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UtilsGetServerTimeResponse is response for Utils.GetServerTime
// Time as Unixtime
type UtilsGetServerTimeResponse int

// GetServerTime Returns the current time of the VK server.
func (v Utils) GetServerTime() (UtilsGetServerTimeResponse, error) {
	r, err := v.API.Request("utils.getServerTime", nil)
	if err != nil {
		return 0, err
	}

	var resp UtilsGetServerTimeResponse

	var cnv int
	cnv, err = strconv.Atoi(string(r))
	resp = UtilsGetServerTimeResponse(cnv)

	if err != nil {
		return 0, err
	}
	return resp, nil
}
