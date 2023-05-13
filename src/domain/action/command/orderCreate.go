package command

import (
	"context"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/model"
)

type CreateOrder struct {
	orderRepo model.OrderRepository
}

func NewCreateOrder(orderRepo model.OrderRepository) *CreateOrder {
	return &CreateOrder{
		orderRepo: orderRepo,
	}
}

func (c CreateOrder) Do(ctx context.Context, order model.Order) (int64, error) {
	orderId, err := c.orderRepo.Create(ctx, order)
	if err != nil {
		return 0, err
	}
	return orderId, nil
}
