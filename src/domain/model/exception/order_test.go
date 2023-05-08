package exception

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CannotMapOrderStateError(t *testing.T) {
	e1 := CannotMapOrderState{
		State: "some_state",
	}
	e2 := CannotMapOrderState{
		State: "some_state2",
	}
	assert.Equal(t, `cannot map order state some_state`, e1.Error())
	assert.Equal(t, `cannot map order state some_state2`, e2.Error())
}

func Test_OrderNotFoundError(t *testing.T) {
	e1 := OrderNotFound{
		Id: 1,
	}
	e2 := OrderNotFound{
		Id: 2,
	}
	assert.Equal(t, `order with id 1 not found`, e1.Error())
	assert.Equal(t, `order with id 2 not found`, e2.Error())
}

func Test_OrderCannotUpdateError(t *testing.T) {
	e1 := OrderCannotUpdate{
		Id: 1,
	}
	e2 := OrderCannotUpdate{
		Id: 2,
	}
	assert.Equal(t, `order with id 1 cannot update`, e1.Error())
	assert.Equal(t, `order with id 2 cannot update`, e2.Error())
}

func Test_OrderInvalidTransitionStateError(t *testing.T) {
	e1 := OrderInvalidTransitionState{
		Id: 1,
	}
	e2 := OrderInvalidTransitionState{
		Id: 2,
	}
	assert.Equal(t, `invalid transition state for order with id 1`, e1.Error())
	assert.Equal(t, `invalid transition state for order with id 2`, e2.Error())
}
