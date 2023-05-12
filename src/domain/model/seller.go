package model

import (
	"context"
	"github.com/cassa10/arq2-tp1/src/domain/util"
)

type Seller struct {
	Id    int64  `json:"id" binding:"required"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (s *Seller) String() string {
	return util.ParseStruct("Seller", s)
}

//go:generate mockgen -destination=../mock/sellerRepository.go -package=mock -source=seller.go
type SellerRepository interface {
	FindById(ctx context.Context, id int64) (*Seller, error)
}
