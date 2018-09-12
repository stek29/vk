package vkCallbackApi

import (
	"strconv"

	"github.com/mailru/easyjson"
)

// APINotes implements VK API namespace `notes`
type APINotes struct {
	API *API
}

// NotesGetParams are params for APINotes.Get
type NotesGetParams struct {
	// Note IDs.
	NoteIDs CSVIntSlice `url:"note_ids,omitempty"`
	// Note owner ID.
	UserID int `url:"user_id,omitempty"`
	// Number of notes to return.
	Count int `url:"count,omitempty"`
}

// NotesGetResponse is response for APINotes.Get
//easyjson:json
type NotesGetResponse struct {
	// Total number
	Count int    `json:"count,omitempty"`
	Items []Note `json:"items,omitempty"`
}

// Get Returns a list of notes created by a user.
func (v APINotes) Get(params NotesGetParams) (*NotesGetResponse, error) {
	r, err := v.API.Request("notes.get", params)
	if err != nil {
		return nil, err
	}

	var resp NotesGetResponse
	err = easyjson.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// NotesGetByIDParams are params for APINotes.GetByID
type NotesGetByIDParams struct {
	// Note ID.
	NoteID int `url:"note_id"`
	// Note owner ID.
	OwnerID int `url:"owner_id,omitempty"`
}

// NotesGetByIDResponse is response for APINotes.GetByID
//easyjson:json
type NotesGetByIDResponse Note

// GetByID Returns a note by its ID.
func (v APINotes) GetByID(params NotesGetByIDParams) (*NotesGetByIDResponse, error) {
	r, err := v.API.Request("notes.getById", params)
	if err != nil {
		return nil, err
	}

	var resp NotesGetByIDResponse
	err = easyjson.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// NotesAddParams are params for APINotes.Add
type NotesAddParams struct {
	// Note title.
	Title string `url:"title"`
	// Note text.
	Text           string         `url:"text"`
	PrivacyView    CSVStringSlice `url:"privacy_view,omitempty"`
	PrivacyComment CSVStringSlice `url:"privacy_comment,omitempty"`
}

// NotesAddResponse is response for APINotes.Add
// Note ID
type NotesAddResponse int

// Add Creates a new note for the current user.
func (v APINotes) Add(params NotesAddParams) (NotesAddResponse, error) {
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

// NotesEditParams are params for APINotes.Edit
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
func (v APINotes) Edit(params NotesEditParams) (bool, error) {
	r, err := v.API.Request("notes.edit", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// NotesDeleteParams are params for APINotes.Delete
type NotesDeleteParams struct {
	// Note ID.
	NoteID int `url:"note_id"`
}

// Delete Deletes a note of the current user.
func (v APINotes) Delete(params NotesDeleteParams) (bool, error) {
	r, err := v.API.Request("notes.delete", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// NotesGetCommentsParams are params for APINotes.GetComments
type NotesGetCommentsParams struct {
	// Note ID.
	NoteID int `url:"note_id"`
	// Note owner ID.
	OwnerID int `url:"owner_id,omitempty"`
	// Number of comments to return.
	Count int `url:"count,omitempty"`
}

// NotesGetCommentsResponse is response for APINotes.GetComments
//easyjson:json
type NotesGetCommentsResponse struct {
	// Total number
	Count int `json:"count,omitempty"`
	Items []genTODOType/* objects.json#/definitions/notes_note_comment */ `json:"items,omitempty"`
}

// GetComments Returns a list of comments on a note.
func (v APINotes) GetComments(params NotesGetCommentsParams) (*NotesGetCommentsResponse, error) {
	r, err := v.API.Request("notes.getComments", params)
	if err != nil {
		return nil, err
	}

	var resp NotesGetCommentsResponse
	err = easyjson.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// NotesCreateCommentParams are params for APINotes.CreateComment
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

// NotesCreateCommentResponse is response for APINotes.CreateComment
// Comment ID
type NotesCreateCommentResponse int

// CreateComment Adds a new comment on a note.
func (v APINotes) CreateComment(params NotesCreateCommentParams) (NotesCreateCommentResponse, error) {
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

// NotesEditCommentParams are params for APINotes.EditComment
type NotesEditCommentParams struct {
	// Comment ID.
	CommentID int `url:"comment_id"`
	// Note owner ID.
	OwnerID int `url:"owner_id,omitempty"`
	// New comment text.
	Message string `url:"message,omitempty"`
}

// EditComment Edits a comment on a note.
func (v APINotes) EditComment(params NotesEditCommentParams) (bool, error) {
	r, err := v.API.Request("notes.editComment", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// NotesDeleteCommentParams are params for APINotes.DeleteComment
type NotesDeleteCommentParams struct {
	// Comment ID.
	CommentID int `url:"comment_id"`
	// Note owner ID.
	OwnerID int `url:"owner_id,omitempty"`
}

// DeleteComment Deletes a comment on a note.
func (v APINotes) DeleteComment(params NotesDeleteCommentParams) (bool, error) {
	r, err := v.API.Request("notes.deleteComment", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// NotesRestoreCommentParams are params for APINotes.RestoreComment
type NotesRestoreCommentParams struct {
	// Comment ID.
	CommentID int `url:"comment_id"`
	// Note owner ID.
	OwnerID int `url:"owner_id,omitempty"`
}

// RestoreComment Restores a deleted comment on a note.
func (v APINotes) RestoreComment(params NotesRestoreCommentParams) (bool, error) {
	r, err := v.API.Request("notes.restoreComment", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}
