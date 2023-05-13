package command

import (
	"context"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/model"
)

type DeleteAllProductsBySellerId struct {
	productRepo model.ProductRepository
}

func NewDeleteAllProductsBySellerId(productRepo model.ProductRepository) *DeleteAllProductsBySellerId {
	return &DeleteAllProductsBySellerId{
		productRepo: productRepo,
	}
}

func (c DeleteAllProductsBySellerId) Do(ctx context.Context, sellerId int64) error {
	if _, err := c.productRepo.DeleteAllBySellerId(ctx, sellerId); err != nil {
		return err
	}
	return nil
}
