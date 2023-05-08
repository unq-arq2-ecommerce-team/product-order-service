package model

import (
	"context"
	"github.com/cassa10/arq2-tp1/src/domain/util"
)

type Seller struct {
	Id       int64     `json:"id" bson:"_id" binding:"required"`
	Name     string    `json:"name" bson:"name" binding:"required"`
	Email    string    `json:"email" bson:"email" binding:"required,email"`
	Products []Product `json:"products" bson:"-"`
}

func (s *Seller) Merge(updateSeller UpdateSeller) {
	s.Name = updateSeller.Name
	s.Email = updateSeller.Email
}

func (s *Seller) String() string {
	return util.ParseStruct("Seller", s)
}

//go:generate mockgen -destination=../mock/sellerRepository.go -package=mock -source=seller.go
type SellerRepository interface {
	FindById(ctx context.Context, id int64) (*Seller, error)
	FindByName(ctx context.Context, name string) (*Seller, error)
	Create(ctx context.Context, seller Seller) (int64, error)
	Update(ctx context.Context, seller Seller) (bool, error)
	Delete(ctx context.Context, id int64) (bool, error)
}
