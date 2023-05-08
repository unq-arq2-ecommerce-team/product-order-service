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

func Test_GivenDeleteProductCmdAndProductId_WhenDo_ThenReturnNoError(t *testing.T) {
	productDeleteCmd, mocks := setUpProductDeleteCmd(t)
	ctx := context.Background()
	productId := int64(123)
	mocks.ProductRepo.EXPECT().FindById(ctx, productId).Return(&model.Product{Id: productId}, nil)
	mocks.ProductRepo.EXPECT().Delete(ctx, productId).Return(true, nil)

	err := productDeleteCmd.Do(ctx, productId)

	assert.NoError(t, err)
}

func Test_GivenDeleteProductCmdAndProductIdAndProductRepoDeleteError_WhenDo_ThenReturnThatError(t *testing.T) {
	productDeleteCmd, mocks := setUpProductDeleteCmd(t)
	ctx := context.Background()
	productId := int64(123)
	mocks.ProductRepo.EXPECT().FindById(ctx, productId).Return(&model.Product{Id: productId}, nil)
	mocks.ProductRepo.EXPECT().Delete(ctx, productId).Return(false, exception.ProductCannotDelete{Id: productId})

	err := productDeleteCmd.Do(ctx, productId)

	assert.ErrorIs(t, err, exception.ProductCannotDelete{Id: productId})
}

func Test_GivenDeleteProductCmdAndProductIdAndProductRepoFindByIdError_WhenDo_ThenReturnThatError(t *testing.T) {
	productDeleteCmd, mocks := setUpProductDeleteCmd(t)
	ctx := context.Background()
	productId := int64(123)
	mocks.ProductRepo.EXPECT().FindById(ctx, productId).Return(nil, exception.ProductNotFound{Id: productId})

	err := productDeleteCmd.Do(ctx, productId)

	assert.ErrorIs(t, err, exception.ProductNotFound{Id: productId})
}

func setUpProductDeleteCmd(t *testing.T) (*DeleteProduct, *mock.InterfaceMocks) {
	mocks := mock.NewInterfaceMocks(t)
	return NewDeleteProduct(mocks.ProductRepo, *query.NewFindProductById(mocks.ProductRepo)), mocks
}
