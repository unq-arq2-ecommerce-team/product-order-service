package query

import (
	"context"
	"github.com/cassa10/arq2-tp1/src/domain/model"
)

type FindCustomerById struct {
	customerRepo model.CustomerRepository
}

func NewFindCustomerById(customerRepo model.CustomerRepository) *FindCustomerById {
	return &FindCustomerById{
		customerRepo: customerRepo,
	}
}

func (q FindCustomerById) Do(ctx context.Context, id int64) (*model.Customer, error) {
	return q.customerRepo.FindById(ctx, id)
}
