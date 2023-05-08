package command

import (
	"context"
	"github.com/cassa10/arq2-tp1/src/domain/mock"
	"github.com/cassa10/arq2-tp1/src/domain/model"
	"github.com/cassa10/arq2-tp1/src/domain/model/exception"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GivenAConfirmedOrderAndDeliveredOrderCmd_WhenDo_ThenReturnNoErrorAndOrderIsDelivered(t *testing.T) {
	deliveredOrderCmd, mocks := setUpDeliveredOrderCmd(t)
	ctx := context.Background()
	order := &model.Order{
		Id:    int64(4),
		State: model.ConfirmedOrderState{},
	}

	orderRepo := *order
	orderRepo.Delivered()
	mocks.OrderRepo.EXPECT().Update(ctx, orderRepo).Return(true, nil)

	err := deliveredOrderCmd.Do(ctx, order)

	assert.NoError(t, err)
	assert.Equal(t, model.DeliveredOrderState{}, order.State)
}

func Test_GivenADeliveredOrderAndDeliveredOrderCmd_WhenDo_ThenDoNothingAndReturnNoErrorForIdempotency(t *testing.T) {
	deliveredOrderCmd, mocks := setUpDeliveredOrderCmd(t)
	ctx := context.Background()
	order := &model.Order{
		Id:    6,
		State: model.DeliveredOrderState{},
	}
	copyOrder := *order
	mocks.OrderRepo.EXPECT().Update(ctx, gomock.Any()).Times(0)
	err := deliveredOrderCmd.Do(ctx, order)

	assert.NoError(t, err)
	assert.Equal(t, copyOrder, *order)
}

func Test_GivenAPendingOrderAndDeliveredOrderCmd_WhenDo_ThenDoNothingAndReturnInvalidTransitionStateError(t *testing.T) {
	deliveredOrderCmd, mocks := setUpDeliveredOrderCmd(t)
	ctx := context.Background()
	idOrder := int64(4)
	order := &model.Order{
		Id:    idOrder,
		State: model.PendingOrderState{},
	}
	copyOrder := *order
	mocks.OrderRepo.EXPECT().Update(ctx, gomock.Any()).Times(0)
	err := deliveredOrderCmd.Do(ctx, order)

	assert.ErrorIs(t, err, exception.OrderInvalidTransitionState{Id: idOrder})
	assert.Equal(t, copyOrder, *order)
}

func Test_GivenAConfirmedOrderAndDeliveredOrderCmdAndOrderRepoUpdateWithError_WhenDo_ThenReturnThatError(t *testing.T) {
	deliveredOrderCmd, mocks := setUpDeliveredOrderCmd(t)
	ctx := context.Background()
	order := &model.Order{
		Id:    int64(4),
		State: model.ConfirmedOrderState{},
	}

	orderRepo := *order
	orderRepo.Delivered()
	mocks.OrderRepo.EXPECT().Update(ctx, orderRepo).Return(false, exception.OrderCannotUpdate{Id: order.Id})

	err := deliveredOrderCmd.Do(ctx, order)

	assert.ErrorIs(t, err, exception.OrderCannotUpdate{Id: order.Id})
}

func setUpDeliveredOrderCmd(t *testing.T) (*DeliveredOrder, *mock.InterfaceMocks) {
	mocks := mock.NewInterfaceMocks(t)
	return NewDeliveredOrder(mocks.OrderRepo), mocks
}
