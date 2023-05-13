package query

import (
	"context"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/model"
)

type FindProductById struct {
	productRepo model.ProductRepository
}

func NewFindProductById(productRepo model.ProductRepository) *FindProductById {
	return &FindProductById{
		productRepo: productRepo,
	}
}

func (q FindProductById) Do(ctx context.Context, id int64) (*model.Product, error) {
	return q.productRepo.FindById(ctx, id)
}
