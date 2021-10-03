package vkapi

import (
	"encoding/json"

	"github.com/stek29/vk"
)

// Database implements VK API namespace `database`
type Database struct {
	API vk.API
}

// DatabaseGetCountriesParams are params for Database.GetCountries
type DatabaseGetCountriesParams struct {
	// '1' — to return a full list of all countries, '0' — to return a list of countries near the current user's country (default).
	NeedAll bool `url:"need_all,omitempty"`
	// Country codes in [vk.com/dev/country_codes|ISO 3166-1 alpha-2] standard.
	Code string `url:"code,omitempty"`
	// Offset needed to return a specific subset of countries.
	Offset int `url:"offset,omitempty"`
	// Number of countries to return.
	Count int `url:"count,omitempty"`
}

// DatabaseGetCountriesResponse is response for Database.GetCountries
//easyjson:json
type DatabaseGetCountriesResponse struct {
	// Total number
	Count int             `json:"count,omitempty"`
	Items []vk.BaseObject `json:"items,omitempty"`
}

// GetCountries Returns a list of countries.
func (v Database) GetCountries(params DatabaseGetCountriesParams) (*DatabaseGetCountriesResponse, error) {
	r, err := v.API.Request("database.getCountries", params)
	if err != nil {
		return nil, err
	}

	var resp DatabaseGetCountriesResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DatabaseGetRegionsParams are params for Database.GetRegions
type DatabaseGetRegionsParams struct {
	// Country ID, received in [vk.com/dev/database.getCountries|database.getCountries] method.
	CountryID int `url:"country_id"`
	// Search query.
	Q string `url:"q,omitempty"`
	// Offset needed to return specific subset of regions.
	Offset int `url:"offset,omitempty"`
	// Number of regions to return.
	Count int `url:"count,omitempty"`
}

// DatabaseGetRegionsResponse is response for Database.GetRegions
//easyjson:json
type DatabaseGetRegionsResponse struct {
	// Total number
	Count int             `json:"count,omitempty"`
	Items []vk.BaseObject `json:"items,omitempty"`
}

// GetRegions Returns a list of regions.
func (v Database) GetRegions(params DatabaseGetRegionsParams) (*DatabaseGetRegionsResponse, error) {
	r, err := v.API.Request("database.getRegions", params)
	if err != nil {
		return nil, err
	}

	var resp DatabaseGetRegionsResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DatabaseGetCountriesByIDParams are params for Database.GetCountriesByID
type DatabaseGetCountriesByIDParams struct {
	// Country IDs.
	CountryIDs CSVIntSlice `url:"country_ids,omitempty"`
}

// DatabaseGetCountriesByIDResponse is response for Database.GetCountriesByID
//easyjson:json
type DatabaseGetCountriesByIDResponse []vk.BaseObject

// GetCountriesByID Returns information about countries by their IDs.
func (v Database) GetCountriesByID(params DatabaseGetCountriesByIDParams) (DatabaseGetCountriesByIDResponse, error) {
	r, err := v.API.Request("database.getCountriesById", params)
	if err != nil {
		return nil, err
	}

	var resp DatabaseGetCountriesByIDResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DatabaseGetCitiesParams are params for Database.GetCities
type DatabaseGetCitiesParams struct {
	// Country ID.
	CountryID int `url:"country_id"`
	// Region ID.
	RegionID int `url:"region_id,omitempty"`
	// Search query.
	Q string `url:"q,omitempty"`
	// '1' — to return all cities in the country, '0' — to return major cities in the country (default),
	NeedAll bool `url:"need_all,omitempty"`
	// Offset needed to return a specific subset of cities.
	Offset int `url:"offset,omitempty"`
	// Number of cities to return.
	Count int `url:"count,omitempty"`
}

// DatabaseGetCitiesResponse is response for Database.GetCities
//easyjson:json
type DatabaseGetCitiesResponse struct {
	// Total number
	Count int               `json:"count,omitempty"`
	Items []vk.DatabaseCity `json:"items,omitempty"`
}

// GetCities Returns a list of cities.
func (v Database) GetCities(params DatabaseGetCitiesParams) (*DatabaseGetCitiesResponse, error) {
	r, err := v.API.Request("database.getCities", params)
	if err != nil {
		return nil, err
	}

	var resp DatabaseGetCitiesResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DatabaseGetCitiesByIDParams are params for Database.GetCitiesByID
type DatabaseGetCitiesByIDParams struct {
	// City IDs.
	CityIDs CSVIntSlice `url:"city_ids,omitempty"`
}

// DatabaseGetCitiesByIDResponse is response for Database.GetCitiesByID
//easyjson:json
type DatabaseGetCitiesByIDResponse []vk.BaseObject

// GetCitiesByID Returns information about cities by their IDs.
func (v Database) GetCitiesByID(params DatabaseGetCitiesByIDParams) (DatabaseGetCitiesByIDResponse, error) {
	r, err := v.API.Request("database.getCitiesById", params)
	if err != nil {
		return nil, err
	}

	var resp DatabaseGetCitiesByIDResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DatabaseGetUniversitiesParams are params for Database.GetUniversities
type DatabaseGetUniversitiesParams struct {
	// Search query.
	Q string `url:"q,omitempty"`
	// Country ID.
	CountryID int `url:"country_id,omitempty"`
	// City ID.
	CityID int `url:"city_id,omitempty"`
	// Offset needed to return a specific subset of universities.
	Offset int `url:"offset,omitempty"`
	// Number of universities to return.
	Count int `url:"count,omitempty"`
}

// DatabaseGetUniversitiesResponse is response for Database.GetUniversities
//easyjson:json
type DatabaseGetUniversitiesResponse struct {
	// Total number
	Count int             `json:"count,omitempty"`
	Items []vk.BaseObject `json:"items,omitempty"`
}

// GetUniversities Returns a list of higher education institutions.
func (v Database) GetUniversities(params DatabaseGetUniversitiesParams) (*DatabaseGetUniversitiesResponse, error) {
	r, err := v.API.Request("database.getUniversities", params)
	if err != nil {
		return nil, err
	}

	var resp DatabaseGetUniversitiesResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DatabaseGetSchoolsParams are params for Database.GetSchools
type DatabaseGetSchoolsParams struct {
	// Search query.
	Q string `url:"q,omitempty"`
	// City ID.
	CityID int `url:"city_id"`
	// Offset needed to return a specific subset of schools.
	Offset int `url:"offset,omitempty"`
	// Number of schools to return.
	Count int `url:"count,omitempty"`
}

// DatabaseGetSchoolsResponse is response for Database.GetSchools
//easyjson:json
type DatabaseGetSchoolsResponse struct {
	// Total number
	Count int             `json:"count,omitempty"`
	Items []vk.BaseObject `json:"items,omitempty"`
}

// GetSchools Returns a list of schools.
func (v Database) GetSchools(params DatabaseGetSchoolsParams) (*DatabaseGetSchoolsResponse, error) {
	r, err := v.API.Request("database.getSchools", params)
	if err != nil {
		return nil, err
	}

	var resp DatabaseGetSchoolsResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DatabaseGetSchoolClassesParams are params for Database.GetSchoolClasses
type DatabaseGetSchoolClassesParams struct {
	// Country ID.
	CountryID int `url:"country_id,omitempty"`
}

// DatabaseGetSchoolClassesResponse is response for Database.GetSchoolClasses
// Class ID or letter
//easyjson:json
type DatabaseGetSchoolClassesResponse [][]genTODOType

// GetSchoolClasses Returns a list of school classes specified for the country.
func (v Database) GetSchoolClasses(params DatabaseGetSchoolClassesParams) (DatabaseGetSchoolClassesResponse, error) {
	r, err := v.API.Request("database.getSchoolClasses", params)
	if err != nil {
		return nil, err
	}

	var resp DatabaseGetSchoolClassesResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DatabaseGetFacultiesParams are params for Database.GetFaculties
type DatabaseGetFacultiesParams struct {
	// University ID.
	UniversityID int `url:"university_id"`
	// Offset needed to return a specific subset of faculties.
	Offset int `url:"offset,omitempty"`
	// Number of faculties to return.
	Count int `url:"count,omitempty"`
}

// DatabaseGetFacultiesResponse is response for Database.GetFaculties
//easyjson:json
type DatabaseGetFacultiesResponse struct {
	// Total number
	Count int             `json:"count,omitempty"`
	Items []vk.BaseObject `json:"items,omitempty"`
}

// GetFaculties Returns a list of faculties (i.e., university departments).
func (v Database) GetFaculties(params DatabaseGetFacultiesParams) (*DatabaseGetFacultiesResponse, error) {
	r, err := v.API.Request("database.getFaculties", params)
	if err != nil {
		return nil, err
	}

	var resp DatabaseGetFacultiesResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DatabaseGetChairsParams are params for Database.GetChairs
type DatabaseGetChairsParams struct {
	// id of the faculty to get chairs from
	FacultyID int `url:"faculty_id"`
	// offset required to get a certain subset of chairs
	Offset int `url:"offset,omitempty"`
	// amount of chairs to get
	Count int `url:"count,omitempty"`
}

// DatabaseGetChairsResponse is response for Database.GetChairs
//easyjson:json
type DatabaseGetChairsResponse struct {
	// Total number
	Count int             `json:"count,omitempty"`
	Items []vk.BaseObject `json:"items,omitempty"`
}

// GetChairs Returns list of chairs on a specified faculty.
func (v Database) GetChairs(params DatabaseGetChairsParams) (*DatabaseGetChairsResponse, error) {
	r, err := v.API.Request("database.getChairs", params)
	if err != nil {
		return nil, err
	}

	var resp DatabaseGetChairsResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DatabaseGetMetroStationsParams are params for Database.GetMetroStations
type DatabaseGetMetroStationsParams struct {
	CityID   int  `url:"city_id"`
	Offset   int  `url:"offset,omitempty"`
	Count    int  `url:"count,omitempty"`
	Extended bool `url:"extended,omitempty"`
}

// DatabaseGetMetroStationsResponse is response for Database.GetMetroStations
//easyjson:json
type DatabaseGetMetroStationsResponse struct {
	// Total number
	Count int `json:"count,omitempty"`
	Items []struct {
		// City ID
		CityID int `json:"city_id,omitempty"`
		// Hex color code without #
		Color string `json:"color,omitempty"`
		// Station ID
		ID int `json:"id,omitempty"`
		// Station name
		Name string `json:"name,omitempty"`
	} `json:"items,omitempty"`
}

// GetMetroStations Get metro stations by city
func (v Database) GetMetroStations(params DatabaseGetMetroStationsParams) (*DatabaseGetMetroStationsResponse, error) {
	r, err := v.API.Request("database.getMetroStations", params)
	if err != nil {
		return nil, err
	}

	var resp DatabaseGetMetroStationsResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DatabaseGetMetroStationsByIDParams are params for Database.GetMetroStationsByID
type DatabaseGetMetroStationsByIDParams struct {
	StationIDs CSVIntSlice `url:"station_ids,omitempty"`
}

// DatabaseGetMetroStationsByIDResponse is response for Database.GetMetroStationsByID
//easyjson:json
type DatabaseGetMetroStationsByIDResponse []struct {
	// City ID
	CityID int `json:"city_id,omitempty"`
	// Hex color code without #
	Color string `json:"color,omitempty"`
	// Station ID
	ID int `json:"id,omitempty"`
	// Station name
	Name string `json:"name,omitempty"`
}

// GetMetroStationsByID Get metro station by his id
func (v Database) GetMetroStationsByID(params DatabaseGetMetroStationsByIDParams) (DatabaseGetMetroStationsByIDResponse, error) {
	r, err := v.API.Request("database.getMetroStationsById", params)
	if err != nil {
		return nil, err
	}

	var resp DatabaseGetMetroStationsByIDResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
