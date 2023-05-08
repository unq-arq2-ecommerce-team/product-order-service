package model

import (
	"context"
	"github.com/cassa10/arq2-tp1/src/domain/util"
)

type Product struct {
	Id          int64   `json:"id" bson:"_id"`
	SellerId    int64   `json:"sellerId" bson:"sellerId"`
	Name        string  `json:"name" bson:"name"`
	Description string  `json:"description" bson:"description"`
	Price       float64 `json:"price" bson:"price"`
	Category    string  `json:"category" bson:"category"`
	Stock       int     `json:"stock" bson:"stock"`
}

func (p *Product) Merge(updateProduct UpdateProduct) {
	p.Name = updateProduct.Name
	p.Description = updateProduct.Description
	p.Price = updateProduct.Price
	p.Category = updateProduct.Category
}

// ValidStock returns true when: stock > 0
func (p *Product) ValidStock() bool {
	return p.Stock > 0
}

// ReduceStock if stock < 0 returns false; otherwise decrease in 1 product stock and returns true
func (p *Product) ReduceStock() bool {
	if !p.ValidStock() {
		return false
	}
	p.Stock--
	return true
}

func (p *Product) String() string {
	return util.ParseStruct("Product", p)
}

//go:generate mockgen -destination=../mock/productRepository.go -package=mock -source=product.go
type ProductRepository interface {
	FindById(ctx context.Context, id int64) (*Product, error)
	Create(ctx context.Context, product Product) (int64, error)
	Update(ctx context.Context, product Product) (bool, error)
	Delete(ctx context.Context, id int64) (bool, error)
	DeleteAllBySellerId(ctx context.Context, sellerId int64) (bool, error)
	FindAllBySellerId(ctx context.Context, sellerId int64) ([]Product, error)
	Search(ctx context.Context, filters ProductSearchFilter, pagingReq PagingRequest) ([]Product, Paging, error)
}
