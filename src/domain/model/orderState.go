package model

import (
	"encoding/json"
	"strings"
)

const (
	pendingOrderState   = "PENDING"
	confirmedOrderState = "CONFIRMED"
	deliveredOrderState = "DELIVERED"
)

var stateMapper = map[string]OrderState{
	pendingOrderState:   PendingOrderState{},
	confirmedOrderState: ConfirmedOrderState{},
	deliveredOrderState: DeliveredOrderState{},
}

type OrderState interface {
	// Confirm returns true when order mutates
	Confirm(order *Order) bool
	// Delivered returns true when order mutates
	Delivered(order *Order) bool
	IsConfirmed() bool
	IsDelivered() bool
	String() string
	MarshalJSON() ([]byte, error)
}

func marshalJSONOrderState(orderState OrderState) ([]byte, error) {
	return json.Marshal(orderState.String())
}

type PendingOrderState struct{}

func (pS PendingOrderState) Confirm(order *Order) bool {
	order.State = ConfirmedOrderState{}
	return true
}

func (pS PendingOrderState) Delivered(_ *Order) bool {
	return false
}

func (pS PendingOrderState) IsConfirmed() bool {
	return false
}

func (pS PendingOrderState) IsDelivered() bool {
	return false
}

func (pS PendingOrderState) String() string {
	return pendingOrderState
}

func (pS PendingOrderState) MarshalJSON() ([]byte, error) {
	return marshalJSONOrderState(pS)
}

type ConfirmedOrderState struct{}

func (cS ConfirmedOrderState) Confirm(_ *Order) bool {
	return false
}

func (cS ConfirmedOrderState) Delivered(order *Order) bool {
	order.State = DeliveredOrderState{}
	return true
}

func (cS ConfirmedOrderState) IsConfirmed() bool {
	return true
}

func (cS ConfirmedOrderState) IsDelivered() bool {
	return false
}

func (cS ConfirmedOrderState) String() string {
	return confirmedOrderState
}

func (cS ConfirmedOrderState) MarshalJSON() ([]byte, error) {
	return marshalJSONOrderState(cS)
}

type DeliveredOrderState struct{}

func (dS DeliveredOrderState) Confirm(_ *Order) bool {
	return false
}

func (dS DeliveredOrderState) Delivered(_ *Order) bool {
	return false
}

func (dS DeliveredOrderState) IsConfirmed() bool {
	return true
}

func (dS DeliveredOrderState) IsDelivered() bool {
	return true
}

func (dS DeliveredOrderState) String() string {
	return deliveredOrderState
}

func (dS DeliveredOrderState) MarshalJSON() ([]byte, error) {
	return marshalJSONOrderState(dS)
}

func GetStateByString(state string) (OrderState, bool) {
	orderState, ok := stateMapper[strings.ToUpper(state)]
	return orderState, ok
}
