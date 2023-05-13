package query

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/mock"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/model"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/model/exception"
	"testing"
)

func Test_GivenFindProductByIdQueryAndProductThatExistWithSomeId_WhenDoWithSameId_ThenReturnProductWithSameIdAndNoError(t *testing.T) {
	findProductById, mocks := setUpFindProductById(t)
	ctx := context.Background()
	productIdToFind := int64(4)
	productToFind := &model.Product{
		Id: productIdToFind,
	}
	mocks.ProductRepo.EXPECT().FindById(ctx, productIdToFind).Return(productToFind, nil)

	product, err := findProductById.Do(ctx, productIdToFind)

	assert.Equal(t, productToFind, product)
	assert.NoError(t, err)
}

func Test_GivenFindProductByIdQuery_WhenDoWithId_ThenReturnNoProductAndAnUnexpectedError(t *testing.T) {
	findProductById, mocks := setUpFindProductById(t)
	ctx := context.Background()
	productIdToFind := int64(4)
	errMsg := "unexpected error"
	mocks.ProductRepo.EXPECT().FindById(ctx, productIdToFind).Return(nil, fmt.Errorf(errMsg))

	product, err := findProductById.Do(ctx, productIdToFind)

	assert.Nil(t, product)
	assert.EqualError(t, err, errMsg)
}

func Test_GivenFindProductByIdQuery_WhenDoWithIdThatNotExists_ThenReturnNoProductAndProductNotFoundErrorWithThatId(t *testing.T) {
	findProductById, mocks := setUpFindProductById(t)
	ctx := context.Background()
	productIdToFind := int64(999)
	mocks.ProductRepo.EXPECT().FindById(ctx, productIdToFind).Return(nil, exception.ProductNotFound{Id: productIdToFind})

	product, err := findProductById.Do(ctx, productIdToFind)

	assert.Nil(t, product)
	assert.ErrorIs(t, err, exception.ProductNotFound{Id: productIdToFind})
}

func setUpFindProductById(t *testing.T) (*FindProductById, *mock.InterfaceMocks) {
	mocks := mock.NewInterfaceMocks(t)
	return NewFindProductById(mocks.ProductRepo), mocks
}
