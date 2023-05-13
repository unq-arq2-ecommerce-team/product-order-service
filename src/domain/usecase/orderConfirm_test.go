package usecase

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/action/command"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/action/query"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/mock"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/model"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/model/exception"
	"testing"
)

func Test_GivenAPendingOrderAndConfirmOrderUseCase_WhenDo_ThenReturnNoErrorAndOrderIsConfirmed(t *testing.T) {
	confirmOrderUseCase, mocks := setUpConfirmOrderUseCase(t)
	ctx := context.Background()
	orderId := int64(9)
	order := &model.Order{
		Id:    orderId,
		State: model.PendingOrderState{},
	}

	orderRepo := *order
	orderRepo.Confirm()
	mocks.OrderRepo.EXPECT().FindById(ctx, orderId).Return(order, nil)
	mocks.OrderRepo.EXPECT().Update(ctx, orderRepo).Return(true, nil)

	err := confirmOrderUseCase.Do(ctx, orderId)

	assert.NoError(t, err)
	assert.Equal(t, model.ConfirmedOrderState{}, order.State)
}

func Test_GivenAConfirmedOrDeliveredOrderAndConfirmOrderUseCase_WhenDo_ThenDoNothingAndReturnNoError(t *testing.T) {
	confirmOrderUseCase, mocks := setUpConfirmOrderUseCase(t)
	ctx := context.Background()
	idConfirmedOrder := int64(4)
	confirmedOrder := &model.Order{
		Id:    idConfirmedOrder,
		State: model.ConfirmedOrderState{},
	}
	idDeliveredOrder := int64(6)
	deliveredOrder := &model.Order{
		Id:    idDeliveredOrder,
		State: model.DeliveredOrderState{},
	}

	copyConfirmedOrder := *confirmedOrder
	copyDeliveredOrder := *deliveredOrder
	mocks.OrderRepo.EXPECT().FindById(ctx, idConfirmedOrder).Return(confirmedOrder, nil)
	mocks.OrderRepo.EXPECT().FindById(ctx, idDeliveredOrder).Return(deliveredOrder, nil)
	mocks.OrderRepo.EXPECT().Update(ctx, gomock.Any()).Times(0)

	err1 := confirmOrderUseCase.Do(ctx, idConfirmedOrder)
	err2 := confirmOrderUseCase.Do(ctx, idDeliveredOrder)

	assert.NoError(t, err1)
	assert.NoError(t, err2)
	assert.Equal(t, copyConfirmedOrder, *confirmedOrder)
	assert.Equal(t, copyDeliveredOrder, *deliveredOrder)
}

func Test_GivenConfirmOrderUseCaseAndAPendingOrderAndOrderRepoFindByIdError_WhenDo_ThenReturnThatError(t *testing.T) {
	confirmOrderUseCase, mocks := setUpConfirmOrderUseCase(t)
	ctx := context.Background()
	orderId := int64(9)

	mocks.OrderRepo.EXPECT().FindById(ctx, orderId).Return(nil, exception.OrderNotFound{Id: orderId})

	err := confirmOrderUseCase.Do(ctx, orderId)

	assert.ErrorIs(t, err, exception.OrderNotFound{Id: orderId})
}

func Test_GivenConfirmOrderUseCaseAndAPendingOrderAndOrderRepoUpdateError_WhenDo_ThenReturnThatError(t *testing.T) {
	confirmOrderUseCase, mocks := setUpConfirmOrderUseCase(t)
	ctx := context.Background()
	orderId := int64(9)
	order := &model.Order{
		Id:    orderId,
		State: model.PendingOrderState{},
	}

	orderRepo := *order
	orderRepo.Confirm()
	mocks.OrderRepo.EXPECT().FindById(ctx, orderId).Return(order, nil)
	mocks.OrderRepo.EXPECT().Update(ctx, orderRepo).Return(false, exception.OrderCannotUpdate{Id: orderId})

	err := confirmOrderUseCase.Do(ctx, orderId)

	assert.ErrorIs(t, err, exception.OrderCannotUpdate{Id: orderId})
}

func setUpConfirmOrderUseCase(t *testing.T) (*ConfirmOrder, *mock.InterfaceMocks) {
	mocks := mock.NewInterfaceMocks(t)
	confirmOrderCmd := *command.NewConfirmOrder(mocks.OrderRepo)
	findOrderByIdQuery := *query.NewFindOrderById(mocks.OrderRepo)
	return NewConfirmOrder(mocks.Logger, confirmOrderCmd, findOrderByIdQuery), mocks
}
