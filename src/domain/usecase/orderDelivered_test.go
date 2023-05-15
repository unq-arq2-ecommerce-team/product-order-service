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

func Test_GivenAConfirmedOrderAndDeliveredOrderUseCase_WhenDo_ThenReturnNoErrorAndOrderIsDelivereded(t *testing.T) {
	deliveredOrderUseCase, mocks := setUpDeliveredOrderUseCase(t)
	ctx := context.Background()
	orderId := int64(9)
	order := &model.Order{
		Id:    orderId,
		State: model.ConfirmedOrderState{},
	}

	orderRepo := *order
	orderRepo.Delivered()
	mocks.OrderRepo.EXPECT().FindById(ctx, orderId).Return(order, nil)
	mocks.OrderRepo.EXPECT().Update(ctx, gomock.Any()).Return(true, nil)

	err := deliveredOrderUseCase.Do(ctx, orderId)

	assert.NoError(t, err)
	assert.Equal(t, model.DeliveredOrderState{}, order.State)
}

func Test_GivenADeliveredOrderAndDeliveredOrderUseCase_WhenDo_ThenDoNothingAndReturnNoErrorForIdempotency(t *testing.T) {
	deliveredOrderUseCase, mocks := setUpDeliveredOrderUseCase(t)
	ctx := context.Background()
	idDeliveredOrder := int64(6)
	deliveredOrder := &model.Order{
		Id:    idDeliveredOrder,
		State: model.DeliveredOrderState{},
	}
	copyDeliveredOrder := *deliveredOrder
	mocks.OrderRepo.EXPECT().FindById(ctx, idDeliveredOrder).Return(deliveredOrder, nil)
	mocks.OrderRepo.EXPECT().Update(ctx, gomock.Any()).Times(0)
	err := deliveredOrderUseCase.Do(ctx, idDeliveredOrder)

	assert.NoError(t, err)
	assert.Equal(t, copyDeliveredOrder, *deliveredOrder)
}

func Test_GivenAPendingOrderAndDeliveredOrderUseCase_WhenDo_ThenDoNothingAndReturnInvalidTransitionStateError(t *testing.T) {
	deliveredOrderUseCase, mocks := setUpDeliveredOrderUseCase(t)
	ctx := context.Background()
	idPendingOrder := int64(4)
	pendingOrder := &model.Order{
		Id:    idPendingOrder,
		State: model.PendingOrderState{},
	}

	copyPendingOrder := *pendingOrder
	mocks.OrderRepo.EXPECT().FindById(ctx, idPendingOrder).Return(pendingOrder, nil)
	mocks.OrderRepo.EXPECT().Update(ctx, gomock.Any()).Times(0)

	err := deliveredOrderUseCase.Do(ctx, idPendingOrder)

	assert.ErrorIs(t, err, exception.OrderInvalidTransitionState{Id: idPendingOrder})
	assert.Equal(t, copyPendingOrder, *pendingOrder)
}

func Test_GivenDeliveredOrderUseCaseAndAConfirmedOrderAndOrderRepoFindByIdError_WhenDo_ThenReturnThatError(t *testing.T) {
	deliveredOrderUseCase, mocks := setUpDeliveredOrderUseCase(t)
	ctx := context.Background()
	orderId := int64(9)

	mocks.OrderRepo.EXPECT().FindById(ctx, orderId).Return(nil, exception.OrderNotFound{Id: orderId})

	err := deliveredOrderUseCase.Do(ctx, orderId)

	assert.ErrorIs(t, err, exception.OrderNotFound{Id: orderId})
}

func Test_GivenDeliveredOrderUseCaseAndAConfirmedOrderAndOrderRepoUpdateError_WhenDo_ThenReturnThatError(t *testing.T) {
	deliveredOrderUseCase, mocks := setUpDeliveredOrderUseCase(t)
	ctx := context.Background()
	orderId := int64(9)
	order := &model.Order{
		Id:    orderId,
		State: model.ConfirmedOrderState{},
	}

	orderRepo := *order
	orderRepo.Delivered()
	mocks.OrderRepo.EXPECT().FindById(ctx, orderId).Return(order, nil)
	mocks.OrderRepo.EXPECT().Update(ctx, gomock.Any()).Return(false, exception.OrderCannotUpdate{Id: orderId})

	err := deliveredOrderUseCase.Do(ctx, orderId)

	assert.ErrorIs(t, err, exception.OrderCannotUpdate{Id: orderId})
}

func setUpDeliveredOrderUseCase(t *testing.T) (*DeliveredOrder, *mock.InterfaceMocks) {
	mocks := mock.NewInterfaceMocks(t)
	deliveredOrderCmd := *command.NewDeliveredOrder(mocks.OrderRepo)
	findOrderByIdQuery := *query.NewFindOrderById(mocks.OrderRepo)
	return NewDeliveredOrder(mocks.Logger, deliveredOrderCmd, findOrderByIdQuery), mocks
}
