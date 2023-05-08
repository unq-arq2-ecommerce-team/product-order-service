package query

import (
	"context"
	"github.com/cassa10/arq2-tp1/src/domain/model"
)

type FindOrderById struct {
	orderRepo model.OrderRepository
}

func NewFindOrderById(orderRepo model.OrderRepository) *FindOrderById {
	return &FindOrderById{
		orderRepo: orderRepo,
	}
}

func (q FindOrderById) Do(ctx context.Context, id int64) (*model.Order, error) {
	return q.orderRepo.FindById(ctx, id)
}
