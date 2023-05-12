package usecase

import (
	"context"
	"fmt"
	"github.com/cassa10/arq2-tp1/src/domain/action/command"
	"github.com/cassa10/arq2-tp1/src/domain/action/query"
	"github.com/cassa10/arq2-tp1/src/domain/mock"
	"github.com/cassa10/arq2-tp1/src/domain/model"
	"github.com/cassa10/arq2-tp1/src/domain/model/exception"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_GivenCreateOrderUseCaseAndProductWithStockAndCustomerIdThatExistAndOrderInfo_WhenDo_ThenReturnNoError(t *testing.T) {
	createOrderUseCase, mocks := setUpCreateOrderUseCase(t)
	ctx := context.Background()
	customerId := int64(12)
	productId := int64(444)
	deliveryDate := time.Now()
	deliveryAddress := model.Address{}
	product := &model.Product{Id: productId, Stock: 10}

	newOrderId := int64(777)
	mocks.ProductRepo.EXPECT().FindById(ctx, productId).Return(product, nil)
	//uses gomock.Any() because model.NewOrder(...) returns time.Now in some fields
	mocks.OrderRepo.EXPECT().Create(ctx, gomock.Any()).Return(newOrderId, nil)

	orderId, err := createOrderUseCase.Do(ctx, customerId, productId, deliveryDate, deliveryAddress)

	assert.NoError(t, err)
	assert.Equal(t, newOrderId, orderId)
}

func Test_GivenCreateOrderUseCaseAndProductWithNoStockAndCustomerIdThatExistAndOrderInfo_WhenDo_ThenReturnProductWithNoStockError(t *testing.T) {
	createOrderUseCase, mocks := setUpCreateOrderUseCase(t)
	ctx := context.Background()
	customerId := int64(12)
	productId := int64(444)
	deliveryDate := time.Now()
	deliveryAddress := model.Address{}
	product := &model.Product{Id: productId, Stock: 0}

	mocks.ProductRepo.EXPECT().FindById(ctx, productId).Return(product, nil)

	orderId, err := createOrderUseCase.Do(ctx, customerId, productId, deliveryDate, deliveryAddress)

	assert.ErrorIs(t, err, exception.ProductWithNoStock{Id: productId})
	assert.Equal(t, int64(0), orderId)
}

func Test_GivenCreateOrderUseCaseAndProductWithStockAndCustomerIdThatExistAndOrderInfoAndOrderRepoCreateError_WhenDo_ThenReturnThatError(t *testing.T) {
	createOrderUseCase, mocks := setUpCreateOrderUseCase(t)
	ctx := context.Background()
	customerId := int64(12)
	productId := int64(444)
	deliveryDate := time.Now()
	deliveryAddress := model.Address{}
	product := &model.Product{Id: productId, Stock: 10}

	msgError := "error in create order"
	mocks.ProductRepo.EXPECT().FindById(ctx, productId).Return(product, nil)
	//uses gomock.Any() because model.NewOrder(...) returns time.Now in some fields
	mocks.OrderRepo.EXPECT().Create(ctx, gomock.Any()).Return(int64(0), fmt.Errorf(msgError))

	orderId, err := createOrderUseCase.Do(ctx, customerId, productId, deliveryDate, deliveryAddress)

	assert.EqualError(t, err, msgError)
	assert.Equal(t, int64(0), orderId)
}

func Test_GivenCreateOrderUseCaseAndProductNoExistAndCustomerIdThatExistAndOrderInfoAndProductRepoFindByIdError_WhenDo_ThenReturnThatError(t *testing.T) {
	createOrderUseCase, mocks := setUpCreateOrderUseCase(t)
	ctx := context.Background()
	customerId := int64(12)
	productId := int64(444)
	deliveryDate := time.Now()
	deliveryAddress := model.Address{}

	expectedErr := exception.ProductNotFound{Id: productId}
	mocks.ProductRepo.EXPECT().FindById(ctx, productId).Return(nil, expectedErr)

	orderId, err := createOrderUseCase.Do(ctx, customerId, productId, deliveryDate, deliveryAddress)

	assert.ErrorIs(t, err, expectedErr)
	assert.Equal(t, int64(0), orderId)
}

func setUpCreateOrderUseCase(t *testing.T) (*CreateOrder, *mock.InterfaceMocks) {
	mocks := mock.NewInterfaceMocks(t)
	createOrderCmd := *command.NewCreateOrder(mocks.OrderRepo)
	findProductByIdQuery := *query.NewFindProductById(mocks.ProductRepo)
	return NewCreateOrder(mocks.Logger, createOrderCmd, findProductByIdQuery), mocks
}
