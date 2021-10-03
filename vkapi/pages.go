package vkapi

import (
	"encoding/json"
	"strconv"

	"github.com/stek29/vk"
)

// Pages implements VK API namespace `pages`
type Pages struct {
	API vk.API
}

// PagesGetParams are params for Pages.Get
type PagesGetParams struct {
	// Page owner ID.
	OwnerID int `url:"owner_id,omitempty"`
	// Wiki page ID.
	PageID int `url:"page_id,omitempty"`
	// '1' — to return information about a global wiki page
	Global bool `url:"global,omitempty"`
	// '1' — resulting wiki page is a preview for the attached link
	SitePreview bool `url:"site_preview,omitempty"`
	// Wiki page title.
	Title      string `url:"title,omitempty"`
	NeedSource bool   `url:"need_source,omitempty"`
	// '1' — to return the page as HTML,
	NeedHTML bool `url:"need_html,omitempty"`
}

// PagesGetResponse is response for Pages.Get
//easyjson:json
type PagesGetResponse vk.Page

// Get Returns information about a wiki page.
func (v Pages) Get(params PagesGetParams) (*PagesGetResponse, error) {
	r, err := v.API.Request("pages.get", params)
	if err != nil {
		return nil, err
	}

	var resp PagesGetResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PagesSaveParams are params for Pages.Save
type PagesSaveParams struct {
	// Text of the wiki page in wiki-format.
	Text string `url:"text,omitempty"`
	// Wiki page ID. The 'title' parameter can be passed instead of 'pid'.
	PageID int `url:"page_id,omitempty"`
	// ID of the community that owns the wiki page.
	GroupID int `url:"group_id,omitempty"`
	// User ID
	UserID int `url:"user_id,omitempty"`
	// Wiki page title.
	Title string `url:"title,omitempty"`
}

// PagesSaveResponse is response for Pages.Save
// Page ID
type PagesSaveResponse int

// Save Saves the text of a wiki page.
func (v Pages) Save(params PagesSaveParams) (PagesSaveResponse, error) {
	r, err := v.API.Request("pages.save", params)
	if err != nil {
		return 0, err
	}

	var resp PagesSaveResponse

	var cnv int
	cnv, err = strconv.Atoi(string(r))
	resp = PagesSaveResponse(cnv)

	if err != nil {
		return 0, err
	}
	return resp, nil
}

// PagesSaveAccessParams are params for Pages.SaveAccess
type PagesSaveAccessParams struct {
	// Wiki page ID.
	PageID int `url:"page_id"`
	// ID of the community that owns the wiki page.
	GroupID int `url:"group_id,omitempty"`
	UserID  int `url:"user_id,omitempty"`
	// Who can view the wiki page: '1' — only community members, '2' — all users can view the page, '0' — only community managers
	View int `url:"view,omitempty"`
	// Who can edit the wiki page: '1' — only community members, '2' — all users can edit the page, '0' — only community managers
	Edit int `url:"edit,omitempty"`
}

// PagesSaveAccessResponse is response for Pages.SaveAccess
// Page ID
type PagesSaveAccessResponse int

// SaveAccess Saves modified read and edit access settings for a wiki page.
func (v Pages) SaveAccess(params PagesSaveAccessParams) (PagesSaveAccessResponse, error) {
	r, err := v.API.Request("pages.saveAccess", params)
	if err != nil {
		return 0, err
	}

	var resp PagesSaveAccessResponse

	var cnv int
	cnv, err = strconv.Atoi(string(r))
	resp = PagesSaveAccessResponse(cnv)

	if err != nil {
		return 0, err
	}
	return resp, nil
}

// PagesGetHistoryParams are params for Pages.GetHistory
type PagesGetHistoryParams struct {
	// Wiki page ID.
	PageID int `url:"page_id"`
	// ID of the community that owns the wiki page.
	GroupID int `url:"group_id,omitempty"`
	UserID  int `url:"user_id,omitempty"`
}

// PagesGetHistoryResponse is response for Pages.GetHistory
//easyjson:json
type PagesGetHistoryResponse []struct {
	// Version ID
	ID int `json:"id,omitempty"`
	// Page size in bytes
	Length int `json:"length,omitempty"`
	// Date when the page has been edited in Unixtime
	Edited int `json:"edited,omitempty"`
	// Last editor ID
	EditorID int `json:"editor_id,omitempty"`
	// Last editor name
	EditorName string `json:"editor_name,omitempty"`
}

// GetHistory Returns a list of all previous versions of a wiki page.
func (v Pages) GetHistory(params PagesGetHistoryParams) (PagesGetHistoryResponse, error) {
	r, err := v.API.Request("pages.getHistory", params)
	if err != nil {
		return nil, err
	}

	var resp PagesGetHistoryResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// PagesGetTitlesParams are params for Pages.GetTitles
type PagesGetTitlesParams struct {
	// ID of the community that owns the wiki page.
	GroupID int `url:"group_id,omitempty"`
}

// PagesGetTitlesResponse is response for Pages.GetTitles
//easyjson:json
type PagesGetTitlesResponse []vk.Page

// GetTitles Returns a list of wiki pages in a group.
func (v Pages) GetTitles(params PagesGetTitlesParams) (PagesGetTitlesResponse, error) {
	r, err := v.API.Request("pages.getTitles", params)
	if err != nil {
		return nil, err
	}

	var resp PagesGetTitlesResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// PagesGetVersionParams are params for Pages.GetVersion
type PagesGetVersionParams struct {
	VersionID int `url:"version_id"`
	// ID of the community that owns the wiki page.
	GroupID int `url:"group_id,omitempty"`
	UserID  int `url:"user_id,omitempty"`
	// '1' — to return the page as HTML
	NeedHTML bool `url:"need_html,omitempty"`
}

// PagesGetVersionResponse is response for Pages.GetVersion
//easyjson:json
type PagesGetVersionResponse vk.Page

// GetVersion Returns the text of one of the previous versions of a wiki page.
func (v Pages) GetVersion(params PagesGetVersionParams) (*PagesGetVersionResponse, error) {
	r, err := v.API.Request("pages.getVersion", params)
	if err != nil {
		return nil, err
	}

	var resp PagesGetVersionResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PagesParseWikiParams are params for Pages.ParseWiki
type PagesParseWikiParams struct {
	// Text of the wiki page.
	Text string `url:"text"`
	// ID of the group in the context of which this markup is interpreted.
	GroupID int `url:"group_id,omitempty"`
}

// PagesParseWikiResponse is response for Pages.ParseWiki
// HTML source
type PagesParseWikiResponse string

// ParseWiki Returns HTML representation of the wiki markup.
func (v Pages) ParseWiki(params PagesParseWikiParams) (PagesParseWikiResponse, error) {
	r, err := v.API.Request("pages.parseWiki", params)
	if err != nil {
		return "", err
	}

	var resp PagesParseWikiResponse

	resp = PagesParseWikiResponse(string(r))

	if err != nil {
		return "", err
	}
	return resp, nil
}

// PagesClearCacheParams are params for Pages.ClearCache
type PagesClearCacheParams struct {
	// Address of the page where you need to refesh the cached version
	URL string `url:"url"`
}

// ClearCache Allows to clear the cache of particular 'external' pages which may be attached to VK posts.
func (v Pages) ClearCache(params PagesClearCacheParams) (bool, error) {
	r, err := v.API.Request("pages.clearCache", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}
