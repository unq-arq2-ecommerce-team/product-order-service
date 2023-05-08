package command

import (
	"context"
	"github.com/cassa10/arq2-tp1/src/domain/model"
	"github.com/cassa10/arq2-tp1/src/domain/model/exception"
)

type DeliveredOrder struct {
	orderRepo model.OrderRepository
}

func NewDeliveredOrder(orderRepo model.OrderRepository) *DeliveredOrder {
	return &DeliveredOrder{
		orderRepo: orderRepo,
	}
}

func (c DeliveredOrder) Do(ctx context.Context, order *model.Order) error {
	//idempotency
	if order.IsDelivered() {
		return nil
	}
	ok := order.Delivered()
	if !ok {
		return exception.OrderInvalidTransitionState{Id: order.Id}
	}
	if _, err := c.orderRepo.Update(ctx, *order); err != nil {
		return err
	}
	return nil
}
