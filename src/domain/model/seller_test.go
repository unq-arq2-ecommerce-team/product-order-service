package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Seller_String(t *testing.T) {
	seller1 := Seller{
		Name:     "pepe",
		Email:    "pepegrillo@mail.com",
		Products: []Product{},
	}

	sellerProducts := []Product{{
		Name:        "Jabon",
		Description: "Un jabon lindo",
		SellerId:    1,
		Price:       0.50,
		Category:    "c1",
		Stock:       6,
	}}

	seller2 := Seller{
		Id:       2,
		Name:     "sarasa",
		Email:    "sarasa@mail.com",
		Products: sellerProducts,
	}
	assert.Equal(t, `[Seller]{"id":0,"name":"pepe","email":"pepegrillo@mail.com","products":[]}`, seller1.String())
	assert.Equal(t, `[Seller]{"id":2,"name":"sarasa","email":"sarasa@mail.com","products":[{"id":0,"sellerId":1,"name":"Jabon","description":"Un jabon lindo","price":0.5,"category":"c1","stock":6}]}`, seller2.String())
}

func Test_Seller_Merge(t *testing.T) {
	seller := Seller{
		Id:    5,
		Name:  "pepe",
		Email: "pepegrillo@mail.com",
	}
	seller.Merge(UpdateSeller{
		Name:  "apple station",
		Email: "apple@asd.com",
	})

	resultSeller := Seller{
		Id:    5,
		Name:  "apple station",
		Email: "apple@asd.com",
	}
	assert.Equal(t, resultSeller, seller)
}
