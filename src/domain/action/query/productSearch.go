package query

import (
	"context"
	"github.com/cassa10/arq2-tp1/src/domain/model"
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
