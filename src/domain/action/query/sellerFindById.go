package query

import (
	"context"
	"github.com/cassa10/arq2-tp1/src/domain/model"
)

type FindSellerById struct {
	sellerRepo model.SellerRepository
}

func NewFindSellerById(sellerRepo model.SellerRepository) *FindSellerById {
	return &FindSellerById{
		sellerRepo: sellerRepo,
	}
}

func (q FindSellerById) Do(ctx context.Context, id int64) (*model.Seller, error) {
	return q.sellerRepo.FindById(ctx, id)
}
