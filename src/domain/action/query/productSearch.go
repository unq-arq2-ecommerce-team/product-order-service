package query

import (
	"context"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/model"
)

type SearchProducts struct {
	productRepo model.ProductRepository
}

func NewSearchProducts(productRepo model.ProductRepository) *SearchProducts {
	return &SearchProducts{
		productRepo: productRepo,
	}
}

func (q SearchProducts) Do(ctx context.Context, filters model.ProductSearchFilter, pagingReq model.PagingRequest) ([]model.Product, model.Paging, error) {
	return q.productRepo.Search(ctx, filters, pagingReq)
}
