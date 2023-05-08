package model

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_OrderState_String(t *testing.T) {
	stateOrderPending := PendingOrderState{}
	stateOrderConfirmed := ConfirmedOrderState{}
	stateOrderDelivered := DeliveredOrderState{}

	assert.Equal(t, pendingOrderState, stateOrderPending.String())
	assert.Equal(t, confirmedOrderState, stateOrderConfirmed.String())
	assert.Equal(t, deliveredOrderState, stateOrderDelivered.String())
}

func Test_OrderState_GetStateByString_IsNoCaseSensitive_Ok(t *testing.T) {
	state1, ok1 := GetStateByString(pendingOrderState)
	state2, ok2 := GetStateByString(confirmedOrderState)
	state3, ok3 := GetStateByString(deliveredOrderState)
	state4, ok4 := GetStateByString(strings.ToLower(pendingOrderState))
	state5, ok5 := GetStateByString(strings.ToLower(confirmedOrderState))
	state6, ok6 := GetStateByString(strings.ToLower(deliveredOrderState))
	assert.True(t, ok1)
	assert.Equal(t, PendingOrderState{}, state1)
	assert.True(t, ok2)
	assert.Equal(t, ConfirmedOrderState{}, state2)
	assert.True(t, ok3)
	assert.Equal(t, DeliveredOrderState{}, state3)
	assert.True(t, ok4)
	assert.Equal(t, PendingOrderState{}, state4)
	assert.True(t, ok5)
	assert.Equal(t, ConfirmedOrderState{}, state5)
	assert.True(t, ok6)
	assert.Equal(t, DeliveredOrderState{}, state6)
}

func Test_OrderState_GetStateByString_NotOk(t *testing.T) {
	state1, ok1 := GetStateByString("")
	state2, ok2 := GetStateByString("sarasa")
	assert.False(t, ok1)
	assert.Equal(t, nil, state1)
	assert.False(t, ok2)
	assert.Equal(t, nil, state2)
}

func Test_OrderState_MarshalJSON(t *testing.T) {
	stateOrderPending := PendingOrderState{}
	stateOrderConfirmed := ConfirmedOrderState{}
	stateOrderDelivered := DeliveredOrderState{}

	pendingJson, errPending := stateOrderPending.MarshalJSON()
	confirmedJson, errConfirmed := stateOrderConfirmed.MarshalJSON()
	deliveredJson, errDelivered := stateOrderDelivered.MarshalJSON()

	assert.Equal(t, []byte(fmt.Sprintf(`"%s"`, pendingOrderState)), pendingJson)
	assert.NoError(t, errPending)
	assert.Equal(t, []byte(fmt.Sprintf(`"%s"`, confirmedOrderState)), confirmedJson)
	assert.NoError(t, errConfirmed)
	assert.Equal(t, []byte(fmt.Sprintf(`"%s"`, deliveredOrderState)), deliveredJson)
	assert.NoError(t, errDelivered)
}
