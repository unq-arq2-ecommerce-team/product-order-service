package command

import (
	"context"
	"github.com/cassa10/arq2-tp1/src/domain/model"
)

type ConfirmOrder struct {
	orderRepo model.OrderRepository
}

func NewConfirmOrder(orderRepo model.OrderRepository) *ConfirmOrder {
	return &ConfirmOrder{
		orderRepo: orderRepo,
	}
}

func (c ConfirmOrder) Do(ctx context.Context, order *model.Order) error {
	//idempotency
	if order.IsConfirmed() {
		return nil
	}
	order.Confirm()
	if _, err := c.orderRepo.Update(ctx, *order); err != nil {
		return err
	}
	return nil
}
