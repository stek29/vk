package vkCallbackApi

import (
	"strconv"

	"github.com/mailru/easyjson"
)

// APISecure implements VK API namespace `secure`
type APISecure struct {
	API *API
}

// SecureGetAppBalanceResponse is response for APISecure.GetAppBalance
// App balance
type SecureGetAppBalanceResponse int

// GetAppBalance Returns payment balance of the application in hundredth of a vote.
func (v APISecure) GetAppBalance() (SecureGetAppBalanceResponse, error) {
	r, err := v.API.Request("secure.getAppBalance", nil)
	if err != nil {
		return 0, err
	}

	var resp SecureGetAppBalanceResponse

	var cnv int
	cnv, err = strconv.Atoi(string(r))
	resp = SecureGetAppBalanceResponse(cnv)

	if err != nil {
		return 0, err
	}
	return resp, nil
}

// SecureGetTransactionsHistoryResponse is response for APISecure.GetTransactionsHistory
//easyjson:json
type SecureGetTransactionsHistoryResponse []struct {
	// Transaction ID
	ID int `json:"id,omitempty"`
	// From ID
	UIDFrom int `json:"uid_from,omitempty"`
	// To ID
	UIDTo int `json:"uid_to,omitempty"`
	// Votes number
	Votes int `json:"votes,omitempty"`
	// Transaction date in Unixtime
	Date int `json:"date,omitempty"`
}

// GetTransactionsHistory Shows history of votes transaction between users and the application.
func (v APISecure) GetTransactionsHistory() (SecureGetTransactionsHistoryResponse, error) {
	r, err := v.API.Request("secure.getTransactionsHistory", nil)
	if err != nil {
		return nil, err
	}

	var resp SecureGetTransactionsHistoryResponse
	err = easyjson.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SecureGetSMSHistoryParams are params for APISecure.GetSMSHistory
type SecureGetSMSHistoryParams struct {
	UserID int `url:"user_id,omitempty"`
	// filter by start date. It is set as UNIX-time.
	DateFrom int `url:"date_from,omitempty"`
	// filter by end date. It is set as UNIX-time.
	DateTo int `url:"date_to,omitempty"`
	// number of returned posts. By default — 1000.
	Limit int `url:"limit,omitempty"`
}

// SecureGetSMSHistoryResponse is response for APISecure.GetSMSHistory
//easyjson:json
type SecureGetSMSHistoryResponse []struct {
	// Notification ID
	ID int `json:"id,omitempty"`
	// Application ID
	AppID int `json:"app_id,omitempty"`
	// User ID
	UserID int `json:"user_id,omitempty"`
	// Date when message has been sent in Unixtime
	Date int `json:"date,omitempty"`
	// Messsage text
	Message string `json:"message,omitempty"`
}

// GetSMSHistory Shows a list of SMS notifications sent by the application using [vk.com/dev/secure.sendSMSNotification|secure.sendSMSNotification] method.
func (v APISecure) GetSMSHistory(params SecureGetSMSHistoryParams) (SecureGetSMSHistoryResponse, error) {
	r, err := v.API.Request("secure.getSMSHistory", params)
	if err != nil {
		return nil, err
	}

	var resp SecureGetSMSHistoryResponse
	err = easyjson.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SecureSendSMSNotificationParams are params for APISecure.SendSMSNotification
type SecureSendSMSNotificationParams struct {
	// ID of the user to whom SMS notification is sent. The user shall allow the application to send him/her notifications (, +1).
	UserID int `url:"user_id"`
	// 'SMS' text to be sent in 'UTF-8' encoding. Only Latin letters and numbers are allowed. Maximum size is '160' characters.
	Message string `url:"message"`
}

// SendSMSNotification Sends 'SMS' notification to a user's mobile device.
func (v APISecure) SendSMSNotification(params SecureSendSMSNotificationParams) (bool, error) {
	r, err := v.API.Request("secure.sendSMSNotification", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// SecureSendNotificationParams are params for APISecure.SendNotification
type SecureSendNotificationParams struct {
	UserIDs CSVIntSlice `url:"user_ids,omitempty"`
	UserID  int         `url:"user_id,omitempty"`
	// notification text which should be sent in 'UTF-8' encoding ('254' characters maximum).
	Message string `url:"message"`
}

// SecureSendNotificationResponse is response for APISecure.SendNotification
// User ID
//easyjson:json
type SecureSendNotificationResponse []int

// SendNotification Sends notification to the user.
func (v APISecure) SendNotification(params SecureSendNotificationParams) (SecureSendNotificationResponse, error) {
	r, err := v.API.Request("secure.sendNotification", params)
	if err != nil {
		return nil, err
	}

	var resp SecureSendNotificationResponse
	err = easyjson.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SecureSetCounterParams are params for APISecure.SetCounter
type SecureSetCounterParams struct {
	Counters CSVStringSlice `url:"counters,omitempty"`
	UserID   int            `url:"user_id,omitempty"`
	// counter value.
	Counter int `url:"counter,omitempty"`
}

// SetCounter Sets a counter which is shown to the user in bold in the left menu.
func (v APISecure) SetCounter(params SecureSetCounterParams) (bool, error) {
	r, err := v.API.Request("secure.setCounter", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// SecureSetUserLevelParams are params for APISecure.SetUserLevel
type SecureSetUserLevelParams struct {
	Levels CSVStringSlice `url:"levels,omitempty"`
	UserID int            `url:"user_id,omitempty"`
	// level value.
	Level int `url:"level,omitempty"`
}

// SetUserLevel Sets user game level in the application which can be seen by his/her friends.
func (v APISecure) SetUserLevel(params SecureSetUserLevelParams) (bool, error) {
	r, err := v.API.Request("secure.setUserLevel", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// SecureGetUserLevelParams are params for APISecure.GetUserLevel
type SecureGetUserLevelParams struct {
	UserIDs CSVIntSlice `url:"user_ids"`
}

// SecureGetUserLevelResponse is response for APISecure.GetUserLevel
//easyjson:json
type SecureGetUserLevelResponse []struct {
	// User ID
	UID int `json:"uid,omitempty"`
	// Level
	Level int `json:"level,omitempty"`
}

// GetUserLevel Returns one of the previously set game levels of one or more users in the application.
func (v APISecure) GetUserLevel(params SecureGetUserLevelParams) (SecureGetUserLevelResponse, error) {
	r, err := v.API.Request("secure.getUserLevel", params)
	if err != nil {
		return nil, err
	}

	var resp SecureGetUserLevelResponse
	err = easyjson.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SecureAddAppEventParams are params for APISecure.AddAppEvent
type SecureAddAppEventParams struct {
	// ID of a user to save the data
	UserID int `url:"user_id"`
	// there are 2 default activities: , * 1 – level. Works similar to ,, * 2 – points, saves points amount, Any other value is for saving completed missions
	ActivityID int `url:"activity_id"`
	// depends on activity_id: * 1 – number, current level number,, * 2 – number, current user's points amount, , Any other value is ignored
	Value int `url:"value,omitempty"`
}

// AddAppEvent Adds user activity information to an application
func (v APISecure) AddAppEvent(params SecureAddAppEventParams) (bool, error) {
	r, err := v.API.Request("secure.addAppEvent", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}

// SecureCheckTokenParams are params for APISecure.CheckToken
type SecureCheckTokenParams struct {
	// client 'access_token'
	Token string `url:"token,omitempty"`
	// user 'ip address'. Note that user may access using the 'ipv6' address, in this case it is required to transmit the 'ipv6' address. If not transmitted, the address will not be checked.
	IP string `url:"ip,omitempty"`
}

// SecureCheckTokenResponse is response for APISecure.CheckToken
//easyjson:json
type SecureCheckTokenResponse struct {
	// Returns if successfully processed
	Success BoolInt `json:"success,omitempty"`
	// User ID
	UserID int `json:"user_id,omitempty"`
	// Date when access_token has been generated in Unixtime
	Date int `json:"date,omitempty"`
	// Date when access_token will expire in Unixtime
	Expire int `json:"expire,omitempty"`
}

// CheckToken Checks the user authentication in 'IFrame' and 'Flash' apps using the 'access_token' parameter.
func (v APISecure) CheckToken(params SecureCheckTokenParams) (*SecureCheckTokenResponse, error) {
	r, err := v.API.Request("secure.checkToken", params)
	if err != nil {
		return nil, err
	}

	var resp SecureCheckTokenResponse
	err = easyjson.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
