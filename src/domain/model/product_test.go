package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Product_String(t *testing.T) {
	product1 := Product{
		Name:        "Jabon",
		Description: "Un jabon lindo",
		SellerId:    1,
		Price:       0.50,
		Category:    "c1",
		Stock:       6,
	}
	product2 := Product{
		Id:          5,
		Name:        "Galletitas",
		Description: "Galletitas sonrisa",
		SellerId:    2,
		Price:       100,
		Category:    "c2",
		Stock:       2,
	}
	assert.Equal(t, `[Product]{"id":0,"sellerId":1,"name":"Jabon","description":"Un jabon lindo","price":0.5,"category":"c1","stock":6}`, product1.String())
	assert.Equal(t, `[Product]{"id":5,"sellerId":2,"name":"Galletitas","description":"Galletitas sonrisa","price":100,"category":"c2","stock":2}`, product2.String())
}

func Test_Product_Merge(t *testing.T) {
	product := Product{
		Id:          4,
		Name:        "Jabon",
		Description: "Un jabon lindo",
		SellerId:    1,
		Price:       0.50,
		Category:    "c1",
		Stock:       6,
	}
	product.Merge(UpdateProduct{
		Name:        "Jabon 2lt",
		Description: "Un jabon lindo con 2 lt",
		Price:       3.50,
		Category:    "Limpieza",
	})

	resultProduct := Product{
		Id:          4,
		SellerId:    1,
		Stock:       6,
		Name:        "Jabon 2lt",
		Description: "Un jabon lindo con 2 lt",
		Price:       3.50,
		Category:    "Limpieza",
	}
	assert.Equal(t, resultProduct, product)
}

func Test_Product_ValidStock(t *testing.T) {
	product1 := Product{
		Id:          4,
		Name:        "Jabon",
		Description: "Un jabon lindo",
		SellerId:    1,
		Price:       0.50,
		Category:    "c1",
		Stock:       0,
	}

	product2 := Product{
		Id:          4,
		Name:        "Jabon",
		Description: "Un jabon lindo",
		SellerId:    1,
		Price:       0.50,
		Category:    "c1",
		Stock:       1,
	}

	assert.False(t, product1.ValidStock())
	assert.True(t, product2.ValidStock())
}

func Test_Product_ReduceStock(t *testing.T) {
	product1 := Product{
		Id:          4,
		Name:        "Jabon",
		Description: "Un jabon lindo",
		SellerId:    1,
		Price:       0.50,
		Category:    "c1",
		Stock:       0,
	}

	product2 := Product{
		Id:          4,
		Name:        "Jabon",
		Description: "Un jabon lindo",
		SellerId:    1,
		Price:       0.50,
		Category:    "c1",
		Stock:       1,
	}

	assert.False(t, product1.ReduceStock())
	assert.True(t, product2.ReduceStock())
	assert.Equal(t, 0, product1.Stock)
	assert.Equal(t, 0, product2.Stock)
}
