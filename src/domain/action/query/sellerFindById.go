package query

import (
	"context"
	"github.com/cassa10/arq2-tp1/src/domain/model"
)

type FindSellerById struct {
	sellerRepo  model.SellerRepository
	productRepo model.ProductRepository
}

func NewFindSellerById(sellerRepo model.SellerRepository, productRepo model.ProductRepository) *FindSellerById {
	return &FindSellerById{
		sellerRepo:  sellerRepo,
		productRepo: productRepo,
	}
}

func (q FindSellerById) Do(ctx context.Context, id int64) (*model.Seller, error) {
	seller, err := q.sellerRepo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	sellerProducts, err := q.productRepo.FindAllBySellerId(ctx, seller.Id)
	if err != nil {
		return nil, err
	}
	seller.Products = sellerProducts
	return seller, nil
}
