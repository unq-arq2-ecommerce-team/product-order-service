package command

import (
	"context"
	"github.com/cassa10/arq2-tp1/src/domain/mock"
	"github.com/cassa10/arq2-tp1/src/domain/model"
	"github.com/cassa10/arq2-tp1/src/domain/model/exception"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GivenCreateCustomerCmdAndNewCustomer_WhenDo_ThenReturnNoErrorAndANewId(t *testing.T) {
	createCustomerCmd, mocks := setUpCustomerCreateCmd(t)
	ctx := context.Background()
	customer := model.Customer{
		Firstname: "pepe",
		Lastname:  "grillo",
		Email:     "pepegrillo@mail.com",
	}
	newCustomerId := int64(15)
	mocks.CustomerRepo.EXPECT().Create(ctx, customer).Return(newCustomerId, nil)

	resCustomerId, err := createCustomerCmd.Do(ctx, customer)

	assert.Equal(t, newCustomerId, resCustomerId)
	assert.NoError(t, err)
}

func Test_GivenCreateCustomerCmdAndNewCustomerAndCustomerRepoCreateError_WhenDo_ThenReturnThatError(t *testing.T) {
	createCustomerCmd, mocks := setUpCustomerCreateCmd(t)
	ctx := context.Background()
	customer := model.Customer{
		Firstname: "pepe",
		Lastname:  "grillo",
		Email:     "pepegrillo@mail.com",
	}
	mocks.CustomerRepo.EXPECT().Create(ctx, customer).Return(int64(0), exception.CustomerAlreadyExist{Email: customer.Email})

	resCustomerId, err := createCustomerCmd.Do(ctx, customer)

	assert.Equal(t, int64(0), resCustomerId)
	assert.ErrorIs(t, err, exception.CustomerAlreadyExist{Email: customer.Email})
}

func setUpCustomerCreateCmd(t *testing.T) (*CreateCustomer, *mock.InterfaceMocks) {
	mocks := mock.NewInterfaceMocks(t)
	return NewCreateCustomer(mocks.CustomerRepo), mocks
}
