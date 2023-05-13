package dto

import (
	"fmt"
	"github.com/cassa10/arq2-tp1/src/domain/model"
)

type ProductCreateReq struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required,min=0"`
	Category    string  `json:"category" binding:"required"`
	Stock       int     `json:"stock" binding:"required,min=1"`
}

func (req *ProductCreateReq) MapToModel(sellerId int64) model.Product {
	return model.Product{
		SellerId:    sellerId,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Category:    req.Category,
		Stock:       req.Stock,
	}
}

type ProductSearchResponse struct {
	Paging   model.Paging    `json:"paging"`
	Products []model.Product `json:"products"`
}

func NewProductSearchResponse(products []model.Product, paging model.Paging) ProductSearchResponse {
	return ProductSearchResponse{
		Paging:   paging,
		Products: products,
	}
}

type ProductSearchQueryReq struct {
	PagingParamQuery
	Name     string   `form:"name"`
	Category string   `form:"category"`
	SellerId int64    `form:"sellerId"`
	PriceMin *float64 `form:"priceMin"`
	PriceMax *float64 `form:"priceMax"`
}

func (qs ProductSearchQueryReq) ValidateReq() error {
	if qs.PriceMin != nil && qs.PriceMax != nil && *qs.PriceMin > *qs.PriceMax {
		return fmt.Errorf("priceMin is greater than priceMax")
	}
	if qs.SellerId < 0 {
		return fmt.Errorf("sellerId must be greater than 0")
	}
	return nil
}

func (qs ProductSearchQueryReq) GetProductSearchFilter() model.ProductSearchFilter {
	return model.NewProductSearchFilter(qs.Name, qs.Category, qs.SellerId, qs.PriceMin, qs.PriceMax)
}
