package vkCallbackApi

import (
	"encoding/json"
)

// APIPolls implements VK API namespace `polls`
type APIPolls struct {
	API *API
}

// PollsGetByIDParams are params for APIPolls.GetByID
type PollsGetByIDParams struct {
	// ID of the user or community that owns the poll. Use a negative value to designate a community ID.
	OwnerID int `url:"owner_id,omitempty"`
	// '1' – poll is in a board, '0' – poll is on a wall. '0' by default.
	IsBoard bool `url:"is_board,omitempty"`
	// Poll ID.
	PollID int `url:"poll_id"`
}

// PollsGetByIDResponse is response for APIPolls.GetByID
//easyjson:json
type PollsGetByIDResponse Poll

// GetByID Returns detailed information about a poll by its ID.
func (v APIPolls) GetByID(params PollsGetByIDParams) (*PollsGetByIDResponse, error) {
	r, err := v.API.Request("polls.getById", params)
	if err != nil {
		return nil, err
	}

	var resp PollsGetByIDResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PollsAddVoteParams are params for APIPolls.AddVote
type PollsAddVoteParams struct {
	// ID of the user or community that owns the poll. Use a negative value to designate a community ID.
	OwnerID int `url:"owner_id,omitempty"`
	// Poll ID.
	PollID int `url:"poll_id"`
	// Answer ID.
	AnswerID int  `url:"answer_id"`
	IsBoard  bool `url:"is_board,omitempty"`
}

// AddVote Adds the current user's vote to the selected answer in the poll.
func (v APIPolls) AddVote(params PollsAddVoteParams) (bool, error) {
	r, err := v.API.Request("polls.addVote", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// PollsDeleteVoteParams are params for APIPolls.DeleteVote
type PollsDeleteVoteParams struct {
	// ID of the user or community that owns the poll. Use a negative value to designate a community ID.
	OwnerID int `url:"owner_id,omitempty"`
	// Poll ID.
	PollID int `url:"poll_id"`
	// Answer ID.
	AnswerID int  `url:"answer_id"`
	IsBoard  bool `url:"is_board,omitempty"`
}

// DeleteVote Deletes the current user's vote from the selected answer in the poll.
func (v APIPolls) DeleteVote(params PollsDeleteVoteParams) (bool, error) {
	r, err := v.API.Request("polls.deleteVote", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// PollsGetVotersParams are params for APIPolls.GetVoters
type PollsGetVotersParams struct {
	// ID of the user or community that owns the poll. Use a negative value to designate a community ID.
	OwnerID int `url:"owner_id,omitempty"`
	// Poll ID.
	PollID int `url:"poll_id"`
	// Answer IDs.
	AnswerIDs CSVIntSlice `url:"answer_ids"`
	IsBoard   bool        `url:"is_board,omitempty"`
	// '1' — to return only current user's friends, '0' — to return all users (default),
	FriendsOnly bool `url:"friends_only,omitempty"`
	// Offset needed to return a specific subset of voters. '0' — (default)
	Offset int `url:"offset,omitempty"`
	// Number of user IDs to return (if the 'friends_only' parameter is not set, maximum '1000', otherwise '10'). '100' — (default)
	Count int `url:"count,omitempty"`
	// Profile fields to return. Sample values: 'nickname', 'screen_name', 'sex', 'bdate (birthdate)', 'city', 'country', 'timezone', 'photo', 'photo_medium', 'photo_big', 'has_mobile', 'rate', 'contacts', 'education', 'online', 'counters'.
	Fields CSVStringSlice `url:"fields,omitempty"`
	// Case for declension of user name and surname: , 'nom' — nominative (default) , 'gen' — genitive , 'dat' — dative , 'acc' — accusative , 'ins' — instrumental , 'abl' — prepositional
	NameCase string `url:"name_case,omitempty"`
}

// PollsGetVotersResponse is response for APIPolls.GetVoters
//easyjson:json
type PollsGetVotersResponse []struct {
	// Answer ID
	AnswerID int `json:"answer_id,omitempty"`
	Users    struct {
		// Votes number
		Count int `json:"count,omitempty"`
		// User ID
		Items []int `json:"items,omitempty"`
	} `json:"users,omitempty"`
}

// GetVoters Returns a list of IDs of users who selected specific answers in the poll.
func (v APIPolls) GetVoters(params PollsGetVotersParams) (PollsGetVotersResponse, error) {
	r, err := v.API.Request("polls.getVoters", params)
	if err != nil {
		return nil, err
	}

	var resp PollsGetVotersResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// PollsCreateParams are params for APIPolls.Create
type PollsCreateParams struct {
	// question text
	Question string `url:"question,omitempty"`
	// '1' – anonymous poll, participants list is hidden,, '0' – public poll, participants list is available,, Default value is '0'.
	IsAnonymous bool `url:"is_anonymous,omitempty"`
	// If a poll will be added to a communty it is required to send a negative group identifier. Current user by default.
	OwnerID int `url:"owner_id,omitempty"`
	// available answers list, for example: " ["yes","no","maybe"]", There can be from 1 to 10 answers.
	AddAnswers string `url:"add_answers,omitempty"`
}

// PollsCreateResponse is response for APIPolls.Create
//easyjson:json
type PollsCreateResponse Poll

// Create Creates polls that can be attached to the users' or communities' posts.
func (v APIPolls) Create(params PollsCreateParams) (*PollsCreateResponse, error) {
	r, err := v.API.Request("polls.create", params)
	if err != nil {
		return nil, err
	}

	var resp PollsCreateResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PollsEditParams are params for APIPolls.Edit
type PollsEditParams struct {
	// poll owner id
	OwnerID int `url:"owner_id"`
	// edited poll's id
	PollID int `url:"poll_id"`
	// new question text
	Question string `url:"question,omitempty"`
	// answers list, for example: , "["yes","no","maybe"]"
	AddAnswers string `url:"add_answers,omitempty"`
	// object containing answers that need to be edited,, key – answer id, value – new answer text. Example: {"382967099":"option1", "382967103":"option2"}"
	EditAnswers string `url:"edit_answers,omitempty"`
	// list of answer ids to be deleted. For example: "[382967099, 382967103]"
	DeleteAnswers string `url:"delete_answers,omitempty"`
}

// Edit Edits created polls
func (v APIPolls) Edit(params PollsEditParams) (bool, error) {
	r, err := v.API.Request("polls.edit", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}
