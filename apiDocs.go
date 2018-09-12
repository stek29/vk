package vk

import (
	"encoding/json"
)

// APIDocs implements VK API namespace `docs`
type APIDocs struct {
	API *API
}

// DocsGetParams are params for APIDocs.Get
type DocsGetParams struct {
	// Number of documents to return. By default, all documents.
	Count int `url:"count,omitempty"`
	// Offset needed to return a specific subset of documents.
	Offset int `url:"offset,omitempty"`
	// ID of the user or community that owns the documents. Use a negative value to designate a community ID.
	OwnerID int `url:"owner_id,omitempty"`
}

// DocsGetResponse is response for APIDocs.Get
//easyjson:json
type DocsGetResponse struct {
	// Total number
	Count int        `json:"count,omitempty"`
	Items []Document `json:"items,omitempty"`
}

// Get Returns detailed information about user or community documents.
func (v APIDocs) Get(params DocsGetParams) (*DocsGetResponse, error) {
	r, err := v.API.Request("docs.get", params)
	if err != nil {
		return nil, err
	}

	var resp DocsGetResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DocsGetByIDParams are params for APIDocs.GetByID
type DocsGetByIDParams struct {
	// Document IDs. Example: , "66748_91488,66748_91455",
	Docs CSVStringSlice `url:"docs"`
}

// DocsGetByIDResponse is response for APIDocs.GetByID
//easyjson:json
type DocsGetByIDResponse []Document

// GetByID Returns information about documents by their IDs.
func (v APIDocs) GetByID(params DocsGetByIDParams) (DocsGetByIDResponse, error) {
	r, err := v.API.Request("docs.getById", params)
	if err != nil {
		return nil, err
	}

	var resp DocsGetByIDResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DocsGetUploadServerParams are params for APIDocs.GetUploadServer
type DocsGetUploadServerParams struct {
	// Community ID (if the document will be uploaded to the community).
	GroupID int `url:"group_id,omitempty"`
}

// DocsGetUploadServerResponse is response for APIDocs.GetUploadServer
//easyjson:json
type DocsGetUploadServerResponse struct {
	// Upload URL
	UploadURL string `json:"upload_url,omitempty"`
}

// GetUploadServer Returns the server address for document upload.
func (v APIDocs) GetUploadServer(params DocsGetUploadServerParams) (*DocsGetUploadServerResponse, error) {
	r, err := v.API.Request("docs.getUploadServer", params)
	if err != nil {
		return nil, err
	}

	var resp DocsGetUploadServerResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DocsGetWallUploadServerParams are params for APIDocs.GetWallUploadServer
type DocsGetWallUploadServerParams struct {
	// Community ID (if the document will be uploaded to the community).
	GroupID int `url:"group_id,omitempty"`
}

// DocsGetWallUploadServerResponse is response for APIDocs.GetWallUploadServer
//easyjson:json
type DocsGetWallUploadServerResponse struct {
	// Upload URL
	UploadURL string `json:"upload_url,omitempty"`
}

// GetWallUploadServer Returns the server address for document upload onto a user's or community's wall.
func (v APIDocs) GetWallUploadServer(params DocsGetWallUploadServerParams) (*DocsGetWallUploadServerResponse, error) {
	r, err := v.API.Request("docs.getWallUploadServer", params)
	if err != nil {
		return nil, err
	}

	var resp DocsGetWallUploadServerResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DocsGetMessagesUploadServerParams are params for APIDocs.GetMessagesUploadServer
type DocsGetMessagesUploadServerParams struct {
	// Document type.
	Type string `url:"type,omitempty"`
	// Destination ID. "For user: 'User ID', e.g. '12345'. For chat: '2000000000' + 'Chat ID', e.g. '2000000001'. For community: '- Community ID', e.g. '-12345'. "
	PeerID int `url:"peer_id,omitempty"`
}

// DocsGetMessagesUploadServerResponse is response for APIDocs.GetMessagesUploadServer
//easyjson:json
type DocsGetMessagesUploadServerResponse struct {
	// Upload URL
	UploadURL string `json:"upload_url,omitempty"`
}

// GetMessagesUploadServer Returns the server address for document upload.
func (v APIDocs) GetMessagesUploadServer(params DocsGetMessagesUploadServerParams) (*DocsGetMessagesUploadServerResponse, error) {
	r, err := v.API.Request("docs.getMessagesUploadServer", params)
	if err != nil {
		return nil, err
	}

	var resp DocsGetMessagesUploadServerResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DocsSaveParams are params for APIDocs.Save
type DocsSaveParams struct {
	// This parameter is returned when the file is [vk.com/dev/upload_files_2|uploaded to the server].
	File string `url:"file"`
	// Document title.
	Title string `url:"title,omitempty"`
	// Document tags.
	Tags string `url:"tags,omitempty"`
}

// DocsSaveResponse is response for APIDocs.Save
//easyjson:json
type DocsSaveResponse []Document

// Save Saves a document after [vk.com/dev/upload_files_2|uploading it to a server].
func (v APIDocs) Save(params DocsSaveParams) (DocsSaveResponse, error) {
	r, err := v.API.Request("docs.save", params)
	if err != nil {
		return nil, err
	}

	var resp DocsSaveResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DocsDeleteParams are params for APIDocs.Delete
type DocsDeleteParams struct {
	// ID of the user or community that owns the document. Use a negative value to designate a community ID.
	OwnerID int `url:"owner_id"`
	// Document ID.
	DocID int `url:"doc_id"`
}

// Delete Deletes a user or community document.
func (v APIDocs) Delete(params DocsDeleteParams) (bool, error) {
	r, err := v.API.Request("docs.delete", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// DocsAddParams are params for APIDocs.Add
type DocsAddParams struct {
	// ID of the user or community that owns the document. Use a negative value to designate a community ID.
	OwnerID int `url:"owner_id"`
	// Document ID.
	DocID int `url:"doc_id"`
	// Access key. This parameter is required if 'access_key' was returned with the document's data.
	AccessKey string `url:"access_key,omitempty"`
}

// DocsAddResponse is response for APIDocs.Add
//easyjson:json
type DocsAddResponse struct {
	// Doc ID
	ID int `json:"id,omitempty"`
}

// Add Copies a document to a user's or community's document list.
func (v APIDocs) Add(params DocsAddParams) (*DocsAddResponse, error) {
	r, err := v.API.Request("docs.add", params)
	if err != nil {
		return nil, err
	}

	var resp DocsAddResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DocsGetTypesParams are params for APIDocs.GetTypes
type DocsGetTypesParams struct {
	// ID of the user or community that owns the documents. Use a negative value to designate a community ID.
	OwnerID int `url:"owner_id"`
}

// DocsGetTypesResponse is response for APIDocs.GetTypes
//easyjson:json
type DocsGetTypesResponse struct {
	// Total number
	Count int `json:"count,omitempty"`
	Items []struct {
		// Doc type ID
		ID int `json:"id,omitempty"`
		// Doc type title
		Title string `json:"title,omitempty"`
		// Number of docs
		Count int `json:"count,omitempty"`
	} `json:"items,omitempty"`
}

// GetTypes Returns documents types available for current user.
func (v APIDocs) GetTypes(params DocsGetTypesParams) (*DocsGetTypesResponse, error) {
	r, err := v.API.Request("docs.getTypes", params)
	if err != nil {
		return nil, err
	}

	var resp DocsGetTypesResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DocsSearchParams are params for APIDocs.Search
type DocsSearchParams struct {
	// Search query string.
	Q         string `url:"q"`
	SearchOwn bool   `url:"search_own,omitempty"`
	// Number of results to return.
	Count int `url:"count,omitempty"`
	// Offset needed to return a specific subset of results.
	Offset int `url:"offset,omitempty"`
}

// DocsSearchResponse is response for APIDocs.Search
//easyjson:json
type DocsSearchResponse struct {
	// Total number
	Count int        `json:"count,omitempty"`
	Items []Document `json:"items,omitempty"`
}

// Search Returns a list of documents matching the search criteria.
func (v APIDocs) Search(params DocsSearchParams) (*DocsSearchResponse, error) {
	r, err := v.API.Request("docs.search", params)
	if err != nil {
		return nil, err
	}

	var resp DocsSearchResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DocsEditParams are params for APIDocs.Edit
type DocsEditParams struct {
	// User ID or community ID. Use a negative value to designate a community ID.
	OwnerID int `url:"owner_id"`
	// Document ID.
	DocID int `url:"doc_id"`
	// Document title.
	Title string `url:"title,omitempty"`
	// Document tags.
	Tags CSVStringSlice `url:"tags,omitempty"`
}

// Edit Edits a document.
func (v APIDocs) Edit(params DocsEditParams) (bool, error) {
	r, err := v.API.Request("docs.edit", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}
