package command

import (
	"context"
	"github.com/cassa10/arq2-tp1/src/domain/action/query"
	"github.com/cassa10/arq2-tp1/src/domain/model"
)

type DeleteProduct struct {
	productRepo          model.ProductRepository
	findProductByIdQuery query.FindProductById
}

func NewDeleteProduct(productRepo model.ProductRepository, findProductById query.FindProductById) *DeleteProduct {
	return &DeleteProduct{
		productRepo:          productRepo,
		findProductByIdQuery: findProductById,
	}
}

func (c DeleteProduct) Do(ctx context.Context, id int64) error {
	_, err := c.findProductByIdQuery.Do(ctx, id)
	if err != nil {
		return err
	}
	if _, err := c.productRepo.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
