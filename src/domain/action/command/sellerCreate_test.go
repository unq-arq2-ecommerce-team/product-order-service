package command

import (
	"context"
	"github.com/cassa10/arq2-tp1/src/domain/mock"
	"github.com/cassa10/arq2-tp1/src/domain/model"
	"github.com/cassa10/arq2-tp1/src/domain/model/exception"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GivenCreateSellerCmdAndNewSeller_WhenDo_ThenReturnNoErrorAndANewId(t *testing.T) {
	createSellerCmd, mocks := setUpSellerCreateCmd(t)
	ctx := context.Background()
	seller := model.Seller{
		Name:  "pepe",
		Email: "pepegrillo@mail.com",
	}
	newSellerId := int64(15)
	mocks.SellerRepo.EXPECT().Create(ctx, seller).Return(newSellerId, nil)

	resSellerId, err := createSellerCmd.Do(ctx, seller)

	assert.Equal(t, newSellerId, resSellerId)
	assert.NoError(t, err)
}

func Test_GivenCreateSellerCmdAndNewSellerAndSellerRepoCreateError_WhenDo_ThenReturnThatError(t *testing.T) {
	createSellerCmd, mocks := setUpSellerCreateCmd(t)
	ctx := context.Background()
	seller := model.Seller{
		Name:  "pepe",
		Email: "pepegrillo@mail.com",
	}
	mocks.SellerRepo.EXPECT().Create(ctx, seller).Return(int64(0), exception.SellerAlreadyExist{Name: seller.Name})

	resSellerId, err := createSellerCmd.Do(ctx, seller)

	assert.Equal(t, int64(0), resSellerId)
	assert.ErrorIs(t, err, exception.SellerAlreadyExist{Name: seller.Name})
}

func setUpSellerCreateCmd(t *testing.T) (*CreateSeller, *mock.InterfaceMocks) {
	mocks := mock.NewInterfaceMocks(t)
	return NewCreateSeller(mocks.SellerRepo), mocks
}
