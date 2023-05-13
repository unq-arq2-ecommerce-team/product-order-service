package command

import (
	"context"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/action/query"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/model"
)

type UpdateProduct struct {
	productRepo          model.ProductRepository
	findProductByIdQuery query.FindProductById
}

func NewUpdateProduct(productRepo model.ProductRepository, findProduct query.FindProductById) *UpdateProduct {
	return &UpdateProduct{
		productRepo:          productRepo,
		findProductByIdQuery: findProduct,
	}
}

func (c UpdateProduct) Do(ctx context.Context, productId int64, updateProduct model.UpdateProduct) error {
	product, err := c.findProductByIdQuery.Do(ctx, productId)
	if err != nil {
		return err
	}
	product.Merge(updateProduct)
	if _, err := c.productRepo.Update(ctx, *product); err != nil {
		return err
	}
	return nil
}
