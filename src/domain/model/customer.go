package model

import (
	"context"
	"github.com/cassa10/arq2-tp1/src/domain/util"
)

type Customer struct {
	Id        int64  `json:"id" bson:"_id"`
	Firstname string `json:"firstname" bson:"firstname"`
	Lastname  string `json:"lastname" bson:"lastname"`
	Email     string `json:"email" bson:"email"`
}

func (c *Customer) Merge(updateCustomer UpdateCustomer) {
	c.Firstname = updateCustomer.Firstname
	c.Lastname = updateCustomer.Lastname
	c.Email = updateCustomer.Email
}

func (c *Customer) String() string {
	return util.ParseStruct("Customer", c)
}

//go:generate mockgen -destination=../mock/customerRepository.go -package=mock -source=customer.go
type CustomerRepository interface {
	FindById(ctx context.Context, id int64) (*Customer, error)
	FindByEmail(ctx context.Context, email string) (*Customer, error)
	Create(ctx context.Context, customer Customer) (int64, error)
	Update(ctx context.Context, customer Customer) (bool, error)
	Delete(ctx context.Context, id int64) (bool, error)
}
