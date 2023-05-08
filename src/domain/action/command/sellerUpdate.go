package command

import (
	"context"
	"github.com/cassa10/arq2-tp1/src/domain/action/query"
	"github.com/cassa10/arq2-tp1/src/domain/model"
)

type UpdateSeller struct {
	sellerRepo          model.SellerRepository
	findSellerByIdQuery query.FindSellerById
}

func NewUpdateSeller(sellerRepo model.SellerRepository, findSeller query.FindSellerById) *UpdateSeller {
	return &UpdateSeller{
		sellerRepo:          sellerRepo,
		findSellerByIdQuery: findSeller,
	}
}

func (c UpdateSeller) Do(ctx context.Context, SellerId int64, updateSeller model.UpdateSeller) error {
	seller, err := c.findSellerByIdQuery.Do(ctx, SellerId)
	if err != nil {
		return err
	}
	seller.Merge(updateSeller)
	if _, err := c.sellerRepo.Update(ctx, *seller); err != nil {
		return err
	}
	return nil
}
