package command

import (
	"context"
	"github.com/cassa10/arq2-tp1/src/domain/action/query"
	"github.com/cassa10/arq2-tp1/src/domain/mock"
	"github.com/cassa10/arq2-tp1/src/domain/model"
	"github.com/cassa10/arq2-tp1/src/domain/model/exception"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GivenUpdateCustomerCmdAndCustomerIdAndCustomerUpdate_WhenDo_ThenReturnNoError(t *testing.T) {
	updateCustomerCmd, mocks := setUpCustomerUpdateCmd(t)
	ctx := context.Background()
	customerIdToFind := int64(4)
	customerToUpdate := &model.Customer{
		Id:        customerIdToFind,
		Firstname: "pepe",
		Lastname:  "grillo",
		Email:     "pepe_grillo@gmail.com",
	}
	updateCustomerInfo := model.UpdateCustomer{
		Firstname: "jorge",
		Lastname:  "sarasa",
		Email:     "js@mail.com",
	}
	mocks.CustomerRepo.EXPECT().FindById(ctx, customerIdToFind).Return(customerToUpdate, nil)
	customerToUpdate.Merge(updateCustomerInfo)
	mocks.CustomerRepo.EXPECT().Update(ctx, *customerToUpdate).Return(true, nil)

	err := updateCustomerCmd.Do(ctx, customerIdToFind, updateCustomerInfo)
	assert.NoError(t, err)
}

func Test_GivenUpdateCustomerCmdAndCustomerIdAndCustomerUpdateAndAnyErrorInFindCustomer_WhenDo_ThenReturnThatError(t *testing.T) {
	updateCustomerCmd, mocks := setUpCustomerUpdateCmd(t)
	ctx := context.Background()
	customerIdToFind := int64(4)
	updateCustomerInfo := model.UpdateCustomer{
		Firstname: "jorge",
		Lastname:  "sarasa",
		Email:     "js@mail.com",
	}
	mocks.CustomerRepo.EXPECT().FindById(ctx, customerIdToFind).Return(nil, exception.CustomerNotFound{Id: customerIdToFind})
	err := updateCustomerCmd.Do(ctx, customerIdToFind, updateCustomerInfo)
	assert.ErrorIs(t, err, exception.CustomerNotFound{Id: customerIdToFind})
}

func Test_GivenUpdateCustomerCmdAndCustomerIdAndCustomerUpdateAndAnyErrorInUpdateCustomer_WhenDo_ThenReturnThatError(t *testing.T) {
	updateCustomerCmd, mocks := setUpCustomerUpdateCmd(t)
	ctx := context.Background()
	customerIdToFind := int64(4)
	customerToUpdate := &model.Customer{
		Id:        customerIdToFind,
		Firstname: "pepe",
		Lastname:  "grillo",
		Email:     "pepe_grillo@gmail.com",
	}
	updateCustomerInfo := model.UpdateCustomer{
		Firstname: "jorge",
		Lastname:  "sarasa",
		Email:     "js@mail.com",
	}
	mocks.CustomerRepo.EXPECT().FindById(ctx, customerIdToFind).Return(customerToUpdate, nil)
	customerToUpdate.Merge(updateCustomerInfo)
	mocks.CustomerRepo.EXPECT().Update(ctx, *customerToUpdate).Return(false, exception.CustomerAlreadyExist{Email: updateCustomerInfo.Email})

	err := updateCustomerCmd.Do(ctx, customerIdToFind, updateCustomerInfo)
	assert.ErrorIs(t, err, exception.CustomerAlreadyExist{Email: updateCustomerInfo.Email})
}

func setUpCustomerUpdateCmd(t *testing.T) (*UpdateCustomer, *mock.InterfaceMocks) {
	mocks := mock.NewInterfaceMocks(t)
	return NewUpdateCustomer(mocks.CustomerRepo, *query.NewFindCustomerById(mocks.CustomerRepo)), mocks
}
