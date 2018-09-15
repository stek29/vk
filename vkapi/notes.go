package vkapi

import (
	"encoding/json"
	"strconv"

	"github.com/stek29/vk"
)

// Notes implements VK API namespace `notes`
type Notes struct {
	API vk.API
}

// NotesGetParams are params for Notes.Get
type NotesGetParams struct {
	// Note IDs.
	NoteIDs CSVIntSlice `url:"note_ids,omitempty"`
	// Note owner ID.
	UserID int `url:"user_id,omitempty"`
	// Number of notes to return.
	Count int `url:"count,omitempty"`
}

// NotesGetResponse is response for Notes.Get
//easyjson:json
type NotesGetResponse struct {
	// Total number
	Count int       `json:"count,omitempty"`
	Items []vk.Note `json:"items,omitempty"`
}

// Get Returns a list of notes created by a user.
func (v Notes) Get(params NotesGetParams) (*NotesGetResponse, error) {
	r, err := v.API.Request("notes.get", params)
	if err != nil {
		return nil, err
	}

	var resp NotesGetResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// NotesGetByIDParams are params for Notes.GetByID
type NotesGetByIDParams struct {
	// Note ID.
	NoteID int `url:"note_id"`
	// Note owner ID.
	OwnerID int `url:"owner_id,omitempty"`
}

// NotesGetByIDResponse is response for Notes.GetByID
//easyjson:json
type NotesGetByIDResponse vk.Note

// GetByID Returns a note by its ID.
func (v Notes) GetByID(params NotesGetByIDParams) (*NotesGetByIDResponse, error) {
	r, err := v.API.Request("notes.getById", params)
	if err != nil {
		return nil, err
	}

	var resp NotesGetByIDResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// NotesAddParams are params for Notes.Add
type NotesAddParams struct {
	// Note title.
	Title string `url:"title"`
	// Note text.
	Text           string         `url:"text"`
	PrivacyView    CSVStringSlice `url:"privacy_view,omitempty"`
	PrivacyComment CSVStringSlice `url:"privacy_comment,omitempty"`
}

// NotesAddResponse is response for Notes.Add
// Note ID
type NotesAddResponse int

// Add Creates a new note for the current user.
func (v Notes) Add(params NotesAddParams) (NotesAddResponse, error) {
	r, err := v.API.Request("notes.add", params)
	if err != nil {
		return 0, err
	}

	var resp NotesAddResponse

	var cnv int
	cnv, err = strconv.Atoi(string(r))
	resp = NotesAddResponse(cnv)

	if err != nil {
		return 0, err
	}
	return resp, nil
}

// NotesEditParams are params for Notes.Edit
type NotesEditParams struct {
	// Note ID.
	NoteID int `url:"note_id"`
	// Note title.
	Title string `url:"title"`
	// Note text.
	Text           string         `url:"text"`
	PrivacyView    CSVStringSlice `url:"privacy_view,omitempty"`
	PrivacyComment CSVStringSlice `url:"privacy_comment,omitempty"`
}

// Edit Edits a note of the current user.
func (v Notes) Edit(params NotesEditParams) (bool, error) {
	r, err := v.API.Request("notes.edit", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// NotesDeleteParams are params for Notes.Delete
type NotesDeleteParams struct {
	// Note ID.
	NoteID int `url:"note_id"`
}

// Delete Deletes a note of the current user.
func (v Notes) Delete(params NotesDeleteParams) (bool, error) {
	r, err := v.API.Request("notes.delete", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// NotesGetCommentsParams are params for Notes.GetComments
type NotesGetCommentsParams struct {
	// Note ID.
	NoteID int `url:"note_id"`
	// Note owner ID.
	OwnerID int `url:"owner_id,omitempty"`
	// Number of comments to return.
	Count int `url:"count,omitempty"`
}

// NotesGetCommentsResponse is response for Notes.GetComments
//easyjson:json
type NotesGetCommentsResponse struct {
	// Total number
	Count int `json:"count,omitempty"`
	Items []genTODOType/* objects.json#/definitions/notes_note_comment */ `json:"items,omitempty"`
}

// GetComments Returns a list of comments on a note.
func (v Notes) GetComments(params NotesGetCommentsParams) (*NotesGetCommentsResponse, error) {
	r, err := v.API.Request("notes.getComments", params)
	if err != nil {
		return nil, err
	}

	var resp NotesGetCommentsResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// NotesCreateCommentParams are params for Notes.CreateComment
type NotesCreateCommentParams struct {
	// Note ID.
	NoteID int `url:"note_id"`
	// Note owner ID.
	OwnerID int `url:"owner_id,omitempty"`
	// ID of the user to whom the reply is addressed (if the comment is a reply to another comment).
	ReplyTo int `url:"reply_to,omitempty"`
	// Comment text.
	Message string `url:"message"`
	GUID    string `url:"guid,omitempty"`
}

// NotesCreateCommentResponse is response for Notes.CreateComment
// Comment ID
type NotesCreateCommentResponse int

// CreateComment Adds a new comment on a note.
func (v Notes) CreateComment(params NotesCreateCommentParams) (NotesCreateCommentResponse, error) {
	r, err := v.API.Request("notes.createComment", params)
	if err != nil {
		return 0, err
	}

	var resp NotesCreateCommentResponse

	var cnv int
	cnv, err = strconv.Atoi(string(r))
	resp = NotesCreateCommentResponse(cnv)

	if err != nil {
		return 0, err
	}
	return resp, nil
}

// NotesEditCommentParams are params for Notes.EditComment
type NotesEditCommentParams struct {
	// Comment ID.
	CommentID int `url:"comment_id"`
	// Note owner ID.
	OwnerID int `url:"owner_id,omitempty"`
	// New comment text.
	Message string `url:"message,omitempty"`
}

// EditComment Edits a comment on a note.
func (v Notes) EditComment(params NotesEditCommentParams) (bool, error) {
	r, err := v.API.Request("notes.editComment", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// NotesDeleteCommentParams are params for Notes.DeleteComment
type NotesDeleteCommentParams struct {
	// Comment ID.
	CommentID int `url:"comment_id"`
	// Note owner ID.
	OwnerID int `url:"owner_id,omitempty"`
}

// DeleteComment Deletes a comment on a note.
func (v Notes) DeleteComment(params NotesDeleteCommentParams) (bool, error) {
	r, err := v.API.Request("notes.deleteComment", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// NotesRestoreCommentParams are params for Notes.RestoreComment
type NotesRestoreCommentParams struct {
	// Comment ID.
	CommentID int `url:"comment_id"`
	// Note owner ID.
	OwnerID int `url:"owner_id,omitempty"`
}

// RestoreComment Restores a deleted comment on a note.
func (v Notes) RestoreComment(params NotesRestoreCommentParams) (bool, error) {
	r, err := v.API.Request("notes.restoreComment", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}
