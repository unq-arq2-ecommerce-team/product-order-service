package model

import (
	"context"
	"github.com/cassa10/arq2-tp1/src/domain/util"
	"time"
)

type Order struct {
	Id              int64
	CustomerId      int64
	CreatedOn       time.Time
	UpdatedOn       time.Time
	DeliveryDate    time.Time
	State           OrderState
	Product         *Product
	DeliveryAddress Address
}

func NewOrder(customerId int64, product *Product, deliveryDate time.Time, deliveryAddress Address) Order {
	return Order{
		CustomerId:      customerId,
		CreatedOn:       time.Now(),
		UpdatedOn:       time.Now(),
		DeliveryDate:    deliveryDate,
		State:           PendingOrderState{},
		Product:         product,
		DeliveryAddress: deliveryAddress,
	}
}

func (o *Order) GetProductId() int64 {
	return o.Product.Id
}

// Confirm returns true when order mutates
func (o *Order) Confirm() bool {
	return o.State.Confirm(o)
}

// Delivered returns true when order mutates
func (o *Order) Delivered() bool {
	return o.State.Delivered(o)
}

func (o *Order) IsConfirmed() bool {
	return o.State.IsConfirmed()
}

func (o *Order) IsDelivered() bool {
	return o.State.IsDelivered()
}

func (o *Order) StateAsString() string {
	return o.State.String()
}

func (o *Order) String() string {
	return util.ParseStruct("Order", o)
}

//go:generate mockgen -destination=../mock/orderRepository.go -package=mock -source=order.go
type OrderRepository interface {
	FindById(ctx context.Context, id int64) (*Order, error)
	Create(ctx context.Context, order Order) (int64, error)
	Update(ctx context.Context, order Order) (bool, error)
}
