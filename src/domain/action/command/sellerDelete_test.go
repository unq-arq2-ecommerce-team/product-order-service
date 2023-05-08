package command

import (
	"context"
	"fmt"
	"github.com/cassa10/arq2-tp1/src/domain/action/query"
	"github.com/cassa10/arq2-tp1/src/domain/mock"
	"github.com/cassa10/arq2-tp1/src/domain/model"
	"github.com/cassa10/arq2-tp1/src/domain/model/exception"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GivenDeleteSellerCmdAndSellerId_WhenDo_ThenReturnNoError(t *testing.T) {
	sellerDeleteCmd, mocks := setUpSellerDeleteCmd(t)
	ctx := context.Background()
	sellerId := int64(123)
	mocks.SellerRepo.EXPECT().FindById(ctx, sellerId).Return(&model.Seller{Id: sellerId}, nil)
	mocks.ProductRepo.EXPECT().FindAllBySellerId(ctx, sellerId).Return([]model.Product{}, nil)
	mocks.SellerRepo.EXPECT().Delete(ctx, sellerId).Return(true, nil)
	mocks.ProductRepo.EXPECT().DeleteAllBySellerId(ctx, sellerId).Return(true, nil)

	err := sellerDeleteCmd.Do(ctx, sellerId)

	assert.NoError(t, err)
}

func Test_GivenDeleteSellerCmdAndSellerIdAndProductRepoDeleteAllBySellerIdWithError_WhenDo_ThenReturnThatError(t *testing.T) {
	sellerDeleteCmd, mocks := setUpSellerDeleteCmd(t)
	ctx := context.Background()
	sellerId := int64(123)
	mocks.SellerRepo.EXPECT().FindById(ctx, sellerId).Return(&model.Seller{Id: sellerId}, nil)
	mocks.ProductRepo.EXPECT().FindAllBySellerId(ctx, sellerId).Return([]model.Product{}, nil)
	mocks.SellerRepo.EXPECT().Delete(ctx, sellerId).Return(true, nil)
	msgError := "error when delete all by sellerId"
	mocks.ProductRepo.EXPECT().DeleteAllBySellerId(ctx, sellerId).Return(false, fmt.Errorf(msgError))

	err := sellerDeleteCmd.Do(ctx, sellerId)

	assert.EqualError(t, err, msgError)
}

func Test_GivenDeleteSellerCmdAndSellerIdAndSellerRepoDeleteError_WhenDo_ThenReturnThatError(t *testing.T) {
	sellerDeleteCmd, mocks := setUpSellerDeleteCmd(t)
	ctx := context.Background()
	sellerId := int64(123)
	mocks.SellerRepo.EXPECT().FindById(ctx, sellerId).Return(&model.Seller{Id: sellerId}, nil)
	mocks.ProductRepo.EXPECT().FindAllBySellerId(ctx, sellerId).Return([]model.Product{}, nil)
	mocks.SellerRepo.EXPECT().Delete(ctx, sellerId).Return(false, exception.SellerCannotDelete{Id: sellerId})

	err := sellerDeleteCmd.Do(ctx, sellerId)

	assert.ErrorIs(t, err, exception.SellerCannotDelete{Id: sellerId})
}

func Test_GivenDeleteSellerCmdAndSellerIdAndProductRepoFindAllBySellerIdError_WhenDo_ThenReturnThatError(t *testing.T) {
	sellerDeleteCmd, mocks := setUpSellerDeleteCmd(t)
	ctx := context.Background()
	sellerId := int64(123)
	errMsg := "some error"
	mocks.SellerRepo.EXPECT().FindById(ctx, sellerId).Return(&model.Seller{Id: sellerId}, nil)
	mocks.ProductRepo.EXPECT().FindAllBySellerId(ctx, sellerId).Return([]model.Product{}, fmt.Errorf(errMsg))

	err := sellerDeleteCmd.Do(ctx, sellerId)

	assert.EqualError(t, err, errMsg)
}

func Test_GivenDeleteSellerCmdAndSellerIdAndSellerRepoFindByIdError_WhenDo_ThenReturnThatError(t *testing.T) {
	sellerDeleteCmd, mocks := setUpSellerDeleteCmd(t)
	ctx := context.Background()
	sellerId := int64(123)
	mocks.SellerRepo.EXPECT().FindById(ctx, sellerId).Return(nil, exception.SellerNotFound{Id: sellerId})

	err := sellerDeleteCmd.Do(ctx, sellerId)

	assert.ErrorIs(t, err, exception.SellerNotFound{Id: sellerId})
}

func setUpSellerDeleteCmd(t *testing.T) (*DeleteSeller, *mock.InterfaceMocks) {
	mocks := mock.NewInterfaceMocks(t)
	return NewDeleteSeller(mocks.SellerRepo, mocks.ProductRepo, *query.NewFindSellerById(mocks.SellerRepo, mocks.ProductRepo)), mocks
}
