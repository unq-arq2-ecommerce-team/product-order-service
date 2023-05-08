package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ProductSearchFilter_New(t *testing.T) {
	name1, ctg1 := "someName", "someCtg2"
	productSearchFilterFromNew1 := NewProductSearchFilter(name1, ctg1, nil, nil)
	productSearchFilter1 := ProductSearchFilter{
		Name:     name1,
		Category: ctg1,
	}
	name2, ctg2, priceMin, priceMax := "someName", "someCtg2", 0.5, 10.7
	productSearchFilterFromNew2 := NewProductSearchFilter(name2, ctg2, &priceMin, &priceMax)
	productSearchFilter2 := ProductSearchFilter{
		Name:     name2,
		Category: ctg2,
		PriceMin: &priceMin,
		PriceMax: &priceMax,
	}
	assert.Equal(t, productSearchFilter1, productSearchFilterFromNew1)
	assert.Equal(t, productSearchFilter2, productSearchFilterFromNew2)
}

func Test_ProductSearchFilter_String(t *testing.T) {
	productSearchFilter1 := NewProductSearchFilter("someName", "someCtg2", nil, nil)
	priceMin := 0.5
	priceMax := 10.7
	productSearchFilter2 := NewProductSearchFilter("asd", "cat2", &priceMin, &priceMax)
	assert.Equal(t, `[ProductSearchFilter]{"name":"someName","category":"someCtg2","priceMin":null,"priceMax":null}`, productSearchFilter1.String())
	assert.Equal(t, `[ProductSearchFilter]{"name":"asd","category":"cat2","priceMin":0.5,"priceMax":10.7}`, productSearchFilter2.String())
}

func Test_ProductSearchFilter_ContainsAnyPriceFilter(t *testing.T) {
	priceMin := 0.5
	priceMax := 10.7
	productSearchFilter1 := NewProductSearchFilter("", "", nil, nil)
	productSearchFilter2 := NewProductSearchFilter("", "", &priceMin, &priceMax)
	productSearchFilter3 := NewProductSearchFilter("", "", nil, &priceMax)
	productSearchFilter4 := NewProductSearchFilter("", "", &priceMin, nil)
	assert.False(t, productSearchFilter1.ContainsAnyPriceFilter())
	assert.True(t, productSearchFilter2.ContainsAnyPriceFilter())
	assert.True(t, productSearchFilter3.ContainsAnyPriceFilter())
	assert.True(t, productSearchFilter4.ContainsAnyPriceFilter())
}

func Test_ProductSearchFilter_GetPriceMinOrDefault(t *testing.T) {
	priceMin1 := 0.5
	priceMin2 := 100.99
	productSearchFilter1 := NewProductSearchFilter("", "", nil, nil)
	productSearchFilter2 := NewProductSearchFilter("", "", &priceMin1, nil)
	productSearchFilter3 := NewProductSearchFilter("", "", &priceMin2, nil)
	assert.Equal(t, priceMinDefault, productSearchFilter1.GetPriceMinOrDefault())
	assert.Equal(t, priceMin1, productSearchFilter2.GetPriceMinOrDefault())
	assert.Equal(t, priceMin2, productSearchFilter3.GetPriceMinOrDefault())
}

func Test_ProductSearchFilter_GetPriceMaxOrDefault(t *testing.T) {
	priceMax1 := 0.5
	priceMax2 := 100.99
	productSearchFilter1 := NewProductSearchFilter("", "", nil, nil)
	productSearchFilter2 := NewProductSearchFilter("", "", nil, &priceMax1)
	productSearchFilter3 := NewProductSearchFilter("", "", nil, &priceMax2)
	assert.Equal(t, priceMaxDefault, productSearchFilter1.GetPriceMaxOrDefault())
	assert.Equal(t, priceMax1, productSearchFilter2.GetPriceMaxOrDefault())
	assert.Equal(t, priceMax2, productSearchFilter3.GetPriceMaxOrDefault())
}
