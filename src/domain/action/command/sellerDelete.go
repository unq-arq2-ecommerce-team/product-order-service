package command

import (
	"context"
	"github.com/cassa10/arq2-tp1/src/domain/action/query"
	"github.com/cassa10/arq2-tp1/src/domain/model"
)

type DeleteSeller struct {
	sellerRepo          model.SellerRepository
	productRepo         model.ProductRepository
	findSellerByIdQuery query.FindSellerById
}

func NewDeleteSeller(sellerRepo model.SellerRepository, productRepo model.ProductRepository, findSellerById query.FindSellerById) *DeleteSeller {
	return &DeleteSeller{
		sellerRepo:          sellerRepo,
		productRepo:         productRepo,
		findSellerByIdQuery: findSellerById,
	}
}

func (c DeleteSeller) Do(ctx context.Context, id int64) error {
	_, err := c.findSellerByIdQuery.Do(ctx, id)
	if err != nil {
		return err
	}
	if _, err := c.sellerRepo.Delete(ctx, id); err != nil {
		return err
	}
	if _, err := c.productRepo.DeleteAllBySellerId(ctx, id); err != nil {
		return err
	}
	return nil
}
