package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_Order_String(t *testing.T) {
	order1 := Order{
		CustomerId:      4,
		CreatedOn:       time.Date(2023, 4, 13, 3, 0, 0, 0, time.UTC),
		UpdatedOn:       time.Date(2023, 4, 13, 9, 0, 0, 0, time.UTC),
		DeliveryDate:    time.Date(2023, 4, 25, 16, 0, 0, 0, time.UTC),
		State:           PendingOrderState{},
		Product:         &Product{},
		DeliveryAddress: Address{},
	}

	order2 := Order{
		Id:           5,
		CustomerId:   2,
		CreatedOn:    time.Date(2023, 8, 23, 3, 0, 0, 0, time.UTC),
		UpdatedOn:    time.Date(2024, 8, 23, 9, 0, 0, 0, time.UTC),
		DeliveryDate: time.Date(2023, 10, 25, 15, 0, 0, 0, time.UTC),
		State:        ConfirmedOrderState{},
		Product: &Product{
			Id:          5,
			Name:        "Galletitas",
			Description: "Galletitas sonrisa",
			SellerId:    2,
			Price:       100,
			Category:    "Almacen",
			Stock:       4,
		},
		DeliveryAddress: Address{
			Street:      "Fake street 123",
			City:        "La Plata",
			State:       "Buenos Aires",
			Country:     "Argentina",
			Observation: "asd",
		},
	}
	assert.Equal(t, `[Order]{"Id":0,"CustomerId":4,"CreatedOn":"2023-04-13T03:00:00Z","UpdatedOn":"2023-04-13T09:00:00Z","DeliveryDate":"2023-04-25T16:00:00Z","State":"PENDING","Product":{"id":0,"sellerId":0,"name":"","description":"","price":0,"category":"","stock":0},"DeliveryAddress":{"street":"","city":"","state":"","country":"","observation":""}}`, order1.String())
	assert.Equal(t, `[Order]{"Id":5,"CustomerId":2,"CreatedOn":"2023-08-23T03:00:00Z","UpdatedOn":"2024-08-23T09:00:00Z","DeliveryDate":"2023-10-25T15:00:00Z","State":"CONFIRMED","Product":{"id":5,"sellerId":2,"name":"Galletitas","description":"Galletitas sonrisa","price":100,"category":"Almacen","stock":4},"DeliveryAddress":{"street":"Fake street 123","city":"La Plata","state":"Buenos Aires","country":"Argentina","observation":"asd"}}`, order2.String())
}

func Test_Order_New(t *testing.T) {
	customerId, deliveryDate := int64(4), time.Date(2023, 4, 25, 16, 0, 0, 0, time.UTC)
	product := &Product{
		Id:          5,
		Name:        "Galletitas",
		Description: "Galletitas sonrisa",
		SellerId:    2,
		Price:       100,
		Category:    "Almacen",
		Stock:       4,
	}
	deliveryAddress := Address{
		Street:      "Fake street 123",
		City:        "La Plata",
		State:       "Buenos Aires",
		Country:     "Argentina",
		Observation: "asd",
	}
	orderFromNew := NewOrder(customerId, product, deliveryDate, deliveryAddress)
	orderRes := Order{
		Id:              0,
		State:           PendingOrderState{},
		CreatedOn:       orderFromNew.CreatedOn,
		UpdatedOn:       orderFromNew.UpdatedOn,
		CustomerId:      customerId,
		DeliveryDate:    deliveryDate,
		Product:         product,
		DeliveryAddress: deliveryAddress,
	}

	assert.Equal(t, orderRes, orderFromNew)
}

func Test_Order_GetProductId(t *testing.T) {
	order1 := Order{
		Product: &Product{},
	}
	order2 := Order{
		Product: &Product{Id: 5},
	}
	assert.Equal(t, int64(0), order1.GetProductId())
	assert.Equal(t, int64(5), order2.GetProductId())
}

func Test_Order_StateAsString(t *testing.T) {
	pendingOrder := Order{State: PendingOrderState{}}
	confirmedOrder := Order{State: ConfirmedOrderState{}}
	deliveredOrder := Order{State: DeliveredOrderState{}}
	assert.Equal(t, pendingOrderState, pendingOrder.StateAsString())
	assert.Equal(t, confirmedOrderState, confirmedOrder.StateAsString())
	assert.Equal(t, deliveredOrderState, deliveredOrder.StateAsString())
}

func Test_GivenOrderPending_WhenReceiveConfirm_ThenChangeOrderStateToConfirmedAndReturnTrue(t *testing.T) {
	pendingOrder := Order{State: PendingOrderState{}}
	isConfirmed := pendingOrder.Confirm()
	assert.Equal(t, ConfirmedOrderState{}, pendingOrder.State)
	assert.True(t, isConfirmed)
}

func Test_GivenAnyOrderNoPending_WhenReceiveConfirm_TheDoNothingAndReturnFalse(t *testing.T) {
	confirmedOrder := Order{State: ConfirmedOrderState{}}
	deliveredOrder := Order{State: DeliveredOrderState{}}
	isConfirmed1 := confirmedOrder.Confirm()
	isConfirmed2 := deliveredOrder.Confirm()
	assert.Equal(t, confirmedOrder, confirmedOrder)
	assert.Equal(t, deliveredOrder, deliveredOrder)
	assert.False(t, isConfirmed1)
	assert.False(t, isConfirmed2)
}

func Test_GivenOrderConfirmed_WhenReceiveDelivered_ThenChangeOrderStateToDeliveredAndReturnTrue(t *testing.T) {
	confirmedOrder := Order{State: ConfirmedOrderState{}}
	isDelivered := confirmedOrder.Delivered()
	assert.Equal(t, DeliveredOrderState{}, confirmedOrder.State)
	assert.True(t, isDelivered)
}

func Test_GivenAnyOrderNoConfirmed_WhenReceiveDelivered_TheDoNothingAndReturnFalse(t *testing.T) {
	pendingOrder := Order{State: PendingOrderState{}}
	deliveredOrder := Order{State: DeliveredOrderState{}}
	isDelivered1 := pendingOrder.Delivered()
	isDelivered2 := deliveredOrder.Delivered()
	assert.Equal(t, pendingOrder, pendingOrder)
	assert.Equal(t, deliveredOrder, deliveredOrder)
	assert.False(t, isDelivered1)
	assert.False(t, isDelivered2)
}

func Test_GivenOrderPending_WhenReceiveIsConfirmedOrDelivered_ThenReturnFalse(t *testing.T) {
	pendingOrder := Order{State: PendingOrderState{}}
	assert.False(t, pendingOrder.IsConfirmed())
	assert.False(t, pendingOrder.IsDelivered())
}

func Test_GivenOrderConfirmed_WhenReceiveIsConfirmed_ThenReturnTrue(t *testing.T) {
	confirmedOrder := Order{State: ConfirmedOrderState{}}
	assert.True(t, confirmedOrder.IsConfirmed())
}

func Test_GivenOrderConfirmed_WhenReceiveIsDelivered_ThenReturnFalse(t *testing.T) {
	confirmedOrder := Order{State: ConfirmedOrderState{}}
	assert.False(t, confirmedOrder.IsDelivered())
}

func Test_GivenOrderDelivered_WhenReceiveIsConfirmedOrDelivered_ThenReturnTrue(t *testing.T) {
	deliveredOrder := Order{State: DeliveredOrderState{}}
	assert.True(t, deliveredOrder.IsConfirmed())
	assert.True(t, deliveredOrder.IsDelivered())
}
