package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Customer_String(t *testing.T) {
	customer1 := Customer{
		Firstname: "pepe",
		Lastname:  "grillo",
		Email:     "pepegrillo@mail.com",
	}
	customer2 := Customer{
		Id:        2,
		Firstname: "sarasa",
		Lastname:  "asaras",
		Email:     "sarasa@mail.com",
	}
	assert.Equal(t, `[Customer]{"id":0,"firstname":"pepe","lastname":"grillo","email":"pepegrillo@mail.com"}`, customer1.String())
	assert.Equal(t, `[Customer]{"id":2,"firstname":"sarasa","lastname":"asaras","email":"sarasa@mail.com"}`, customer2.String())
}

func Test_Customer_Merge(t *testing.T) {
	customer := Customer{
		Id:        5,
		Firstname: "pepe",
		Lastname:  "grillo",
		Email:     "pepegrillo@mail.com",
	}
	customer.Merge(UpdateCustomer{
		Firstname: "jorgito",
		Lastname:  "lalala",
		Email:     "jorgito@asd.com",
	})

	resultCustomer := Customer{
		Id:        5,
		Firstname: "jorgito",
		Lastname:  "lalala",
		Email:     "jorgito@asd.com",
	}
	assert.Equal(t, resultCustomer, customer)
}
