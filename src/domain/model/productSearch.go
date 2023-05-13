package model

import "github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/util"

const (
	priceMinDefault = float64(0)
	priceMaxDefault = float64(999999999999999999)
)

type ProductSearchFilter struct {
	Name     string   `json:"name"`
	Category string   `json:"category"`
	SellerId int64    `json:"sellerId"`
	PriceMin *float64 `json:"priceMin"`
	PriceMax *float64 `json:"priceMax"`
}

func NewProductSearchFilter(name, category string, sellerId int64, priceMin, priceMax *float64) ProductSearchFilter {
	return ProductSearchFilter{
		Name:     name,
		Category: category,
		SellerId: sellerId,
		PriceMin: priceMin,
		PriceMax: priceMax,
	}
}

func (f *ProductSearchFilter) ContainsAnyPriceFilter() bool {
	return f.PriceMin != nil || f.PriceMax != nil
}

func (f *ProductSearchFilter) GetPriceMinOrDefault() float64 {
	if f.PriceMin == nil {
		return priceMinDefault
	}
	return *f.PriceMin
}

func (f *ProductSearchFilter) GetPriceMaxOrDefault() float64 {
	if f.PriceMax == nil {
		return priceMaxDefault
	}
	return *f.PriceMax
}

func (f *ProductSearchFilter) String() string {
	return util.ParseStruct("ProductSearchFilter", f)
}
