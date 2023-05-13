package command

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/action/query"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/mock"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/model"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/model/exception"
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
