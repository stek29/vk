package vkCallbackApi

import (
	"github.com/mailru/easyjson"
)

// APIOrders implements VK API namespace `orders`
type APIOrders struct {
	API *API
}

// OrdersGetParams are params for APIOrders.Get
type OrdersGetParams struct {
	// number of returned orders.
	Count int `url:"count,omitempty"`
	// if this parameter is set to 1, this method returns a list of test mode orders. By default — 0.
	TestMode bool `url:"test_mode,omitempty"`
}

// OrdersGetResponse is response for APIOrders.Get
//easyjson:json
type OrdersGetResponse []struct {
	// Order ID
	ID int `json:"id,omitempty"`
	// App order ID
	AppOrderID int `json:"app_order_id,omitempty"`
	// Order status
	Status string `json:"status,omitempty"`
	// User ID
	UserID int `json:"user_id,omitempty"`
	// Receiver ID
	ReceiverID int `json:"receiver_id,omitempty"`
	// Order item
	Item string `json:"item,omitempty"`
	// Amount
	Amount int `json:"amount,omitempty"`
	// Date of creation in Unixtime
	Date int `json:"date,omitempty"`
	// Transaction ID
	TransactionID int `json:"transaction_id,omitempty"`
	// Cancel transaction ID
	CancelTransactionID int `json:"cancel_transaction_id,omitempty"`
}

// Get Returns a list of orders.
func (v APIOrders) Get(params OrdersGetParams) (OrdersGetResponse, error) {
	r, err := v.API.Request("orders.get", params)
	if err != nil {
		return nil, err
	}

	var resp OrdersGetResponse
	err = easyjson.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// OrdersGetByIDParams are params for APIOrders.GetByID
type OrdersGetByIDParams struct {
	// order ID.
	OrderID int `url:"order_id,omitempty"`
	// order IDs (when information about several orders is requested).
	OrderIDs CSVIntSlice `url:"order_ids,omitempty"`
	// if this parameter is set to 1, this method returns a list of test mode orders. By default — 0.
	TestMode bool `url:"test_mode,omitempty"`
}

// OrdersGetByIDResponse is response for APIOrders.GetByID
//easyjson:json
type OrdersGetByIDResponse []struct {
	// Order ID
	ID int `json:"id,omitempty"`
	// App order ID
	AppOrderID int `json:"app_order_id,omitempty"`
	// Order status
	Status string `json:"status,omitempty"`
	// User ID
	UserID int `json:"user_id,omitempty"`
	// Receiver ID
	ReceiverID int `json:"receiver_id,omitempty"`
	// Order item
	Item string `json:"item,omitempty"`
	// Amount
	Amount int `json:"amount,omitempty"`
	// Date of creation in Unixtime
	Date int `json:"date,omitempty"`
	// Transaction ID
	TransactionID int `json:"transaction_id,omitempty"`
	// Cancel transaction ID
	CancelTransactionID int `json:"cancel_transaction_id,omitempty"`
}

// GetByID Returns information about orders by their IDs.
func (v APIOrders) GetByID(params OrdersGetByIDParams) (OrdersGetByIDResponse, error) {
	r, err := v.API.Request("orders.getById", params)
	if err != nil {
		return nil, err
	}

	var resp OrdersGetByIDResponse
	err = easyjson.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// OrdersChangeStateParams are params for APIOrders.ChangeState
type OrdersChangeStateParams struct {
	// order ID.
	OrderID int `url:"order_id"`
	// action to be done with the order. Available actions: *cancel — to cancel unconfirmed order. *charge — to confirm unconfirmed order. Applies only if processing of [vk.com/dev/payments_status|order_change_state] notification failed. *refund — to cancel confirmed order.
	Action string `url:"action"`
	// internal ID of the order in the application.
	AppOrderID int `url:"app_order_id,omitempty"`
	// if this parameter is set to 1, this method returns a list of test mode orders. By default — 0.
	TestMode bool `url:"test_mode,omitempty"`
}

// OrdersChangeStateResponse is response for APIOrders.ChangeState
// New state
type OrdersChangeStateResponse string

// ChangeState Changes order status.
func (v APIOrders) ChangeState(params OrdersChangeStateParams) (OrdersChangeStateResponse, error) {
	r, err := v.API.Request("orders.changeState", params)
	if err != nil {
		return "", err
	}

	var resp OrdersChangeStateResponse

	resp = OrdersChangeStateResponse(string(r))

	if err != nil {
		return "", err
	}
	return resp, nil
}

// OrdersGetAmountParams are params for APIOrders.GetAmount
type OrdersGetAmountParams struct {
	UserID int            `url:"user_id"`
	Votes  CSVStringSlice `url:"votes"`
}

// OrdersGetAmountResponse is response for APIOrders.GetAmount
//easyjson:json
type OrdersGetAmountResponse struct {
	Amounts []struct {
		// Votes number
		Votes string `json:"votes,omitempty"`
		// Votes amount in user's currency
		Amount int `json:"amount,omitempty"`
		// Amount description
		Description string `json:"description,omitempty"`
	} `json:"amounts,omitempty"`
	// Currency name
	Currency string `json:"currency,omitempty"`
}

// GetAmount does orders.getAmount
func (v APIOrders) GetAmount(params OrdersGetAmountParams) (*OrdersGetAmountResponse, error) {
	r, err := v.API.Request("orders.getAmount", params)
	if err != nil {
		return nil, err
	}

	var resp OrdersGetAmountResponse
	err = easyjson.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
