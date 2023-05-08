package query

import (
	"context"
	"fmt"
	"github.com/cassa10/arq2-tp1/src/domain/mock"
	"github.com/cassa10/arq2-tp1/src/domain/model"
	"github.com/cassa10/arq2-tp1/src/domain/model/exception"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GivenFindCustomerByIdQueryAndCustomerThatExistWithSomeId_WhenDoWithSameId_ThenReturnCustomerWithSameIdAndNoError(t *testing.T) {
	findCustomerById, mocks := setUpFindCustomerById(t)
	ctx := context.Background()
	customerIdToFind := int64(4)
	customerToFind := &model.Customer{
		Id: customerIdToFind,
	}
	mocks.CustomerRepo.EXPECT().FindById(ctx, customerIdToFind).Return(customerToFind, nil)

	customer, err := findCustomerById.Do(ctx, customerIdToFind)

	assert.Equal(t, customerToFind, customer)
	assert.NoError(t, err)
}

func Test_GivenFindCustomerByIdQuery_WhenDoWithId_ThenReturnNoCustomerAndAnUnexpectedError(t *testing.T) {
	findCustomerById, mocks := setUpFindCustomerById(t)
	ctx := context.Background()
	customerIdToFind := int64(4)
	errMsg := "unexpected error"
	mocks.CustomerRepo.EXPECT().FindById(ctx, customerIdToFind).Return(nil, fmt.Errorf(errMsg))

	customer, err := findCustomerById.Do(ctx, customerIdToFind)

	assert.Nil(t, customer)
	assert.EqualError(t, err, errMsg)
}

func Test_GivenFindCustomerByIdQuery_WhenDoWithIdThatNotExists_ThenReturnNoCustomerAndCustomerNotFoundErrorWithThatId(t *testing.T) {
	findCustomerById, mocks := setUpFindCustomerById(t)
	ctx := context.Background()
	customerIdToFind := int64(999)
	mocks.CustomerRepo.EXPECT().FindById(ctx, customerIdToFind).Return(nil, exception.CustomerNotFound{Id: customerIdToFind})

	customer, err := findCustomerById.Do(ctx, customerIdToFind)

	assert.Nil(t, customer)
	assert.ErrorIs(t, err, exception.CustomerNotFound{Id: customerIdToFind})
}

func setUpFindCustomerById(t *testing.T) (*FindCustomerById, *mock.InterfaceMocks) {
	mocks := mock.NewInterfaceMocks(t)
	return NewFindCustomerById(mocks.CustomerRepo), mocks
}
