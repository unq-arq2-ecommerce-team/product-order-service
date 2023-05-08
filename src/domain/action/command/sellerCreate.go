package command

import (
	"context"
	"github.com/cassa10/arq2-tp1/src/domain/model"
)

type CreateSeller struct {
	sellerRepo model.SellerRepository
}

func NewCreateSeller(sellerRepo model.SellerRepository) *CreateSeller {
	return &CreateSeller{
		sellerRepo: sellerRepo,
	}
}

func (c CreateSeller) Do(ctx context.Context, seller model.Seller) (int64, error) {
	sellerId, err := c.sellerRepo.Create(ctx, seller)
	if err != nil {
		return 0, err
	}
	return sellerId, nil
}
