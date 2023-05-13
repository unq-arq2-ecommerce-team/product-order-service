package command

import (
	"context"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/action/query"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/model"
)

type CreateProduct struct {
	productRepo         model.ProductRepository
	findSellerByIdQuery query.FindSellerById
}

func NewCreateProduct(productRepo model.ProductRepository, findSellerByIdQuery query.FindSellerById) *CreateProduct {
	return &CreateProduct{
		productRepo:         productRepo,
		findSellerByIdQuery: findSellerByIdQuery,
	}
}

func (c CreateProduct) Do(ctx context.Context, product model.Product) (int64, error) {
	_, err := c.findSellerByIdQuery.Do(ctx, product.SellerId)
	if err != nil {
		return 0, err
	}
	productId, err := c.productRepo.Create(ctx, product)
	if err != nil {
		return 0, err
	}
	return productId, nil
}
