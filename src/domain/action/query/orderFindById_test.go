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

func Test_GivenFindOrderByIdQueryAndOrderThatExistWithSomeId_WhenDoWithSameId_ThenReturnOrderWithSameIdAndNoError(t *testing.T) {
	findOrderById, mocks := setUpFindOrderById(t)
	ctx := context.Background()
	orderIdToFind := int64(4)
	orderToFind := &model.Order{
		Id: orderIdToFind,
	}
	mocks.OrderRepo.EXPECT().FindById(ctx, orderIdToFind).Return(orderToFind, nil)

	order, err := findOrderById.Do(ctx, orderIdToFind)

	assert.Equal(t, orderToFind, order)
	assert.NoError(t, err)
}

func Test_GivenFindOrderByIdQuery_WhenDoWithId_ThenReturnNoOrderAndAnUnexpectedError(t *testing.T) {
	findOrderById, mocks := setUpFindOrderById(t)
	ctx := context.Background()
	orderIdToFind := int64(4)
	errMsg := "unexpected error"
	mocks.OrderRepo.EXPECT().FindById(ctx, orderIdToFind).Return(nil, fmt.Errorf(errMsg))

	order, err := findOrderById.Do(ctx, orderIdToFind)

	assert.Nil(t, order)
	assert.EqualError(t, err, errMsg)
}

func Test_GivenFindOrderByIdQuery_WhenDoWithIdThatNotExists_ThenReturnNoOrderAndOrderNotFoundErrorWithThatId(t *testing.T) {
	findOrderById, mocks := setUpFindOrderById(t)
	ctx := context.Background()
	orderIdToFind := int64(999)
	mocks.OrderRepo.EXPECT().FindById(ctx, orderIdToFind).Return(nil, exception.OrderNotFound{Id: orderIdToFind})

	order, err := findOrderById.Do(ctx, orderIdToFind)

	assert.Nil(t, order)
	assert.ErrorIs(t, err, exception.OrderNotFound{Id: orderIdToFind})
}

func setUpFindOrderById(t *testing.T) (*FindOrderById, *mock.InterfaceMocks) {
	mocks := mock.NewInterfaceMocks(t)
	return NewFindOrderById(mocks.OrderRepo), mocks
}
