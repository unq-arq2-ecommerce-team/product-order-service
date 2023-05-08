package command

import (
	"context"
	"github.com/cassa10/arq2-tp1/src/domain/mock"
	"github.com/cassa10/arq2-tp1/src/domain/model"
	"github.com/cassa10/arq2-tp1/src/domain/model/exception"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_GivenCreateOrderCmdAndNewOrder_WhenDo_ThenReturnNoErrorAndANewId(t *testing.T) {
	createOrderCmd, mocks := setUpOrderCreateCmd(t)
	ctx := context.Background()
	order := model.NewOrder(int64(1), &model.Product{Id: 10}, time.Now(), model.Address{})
	newOrderId := int64(871)
	mocks.OrderRepo.EXPECT().Create(ctx, order).Return(newOrderId, nil)

	resOrderId, err := createOrderCmd.Do(ctx, order)

	assert.Equal(t, newOrderId, resOrderId)
	assert.NoError(t, err)
}

func Test_GivenCreateOrderCmdAndNewOrderAndOrderRepoCreateError_WhenDo_ThenReturnThatError(t *testing.T) {
	createOrderCmd, mocks := setUpOrderCreateCmd(t)
	ctx := context.Background()
	order := model.NewOrder(int64(1), &model.Product{Id: 10}, time.Now(), model.Address{})
	mocks.OrderRepo.EXPECT().Create(ctx, order).Return(int64(0), exception.ProductWithNoStock{Id: order.GetProductId()})

	resOrderId, err := createOrderCmd.Do(ctx, order)

	assert.Equal(t, int64(0), resOrderId)
	assert.ErrorIs(t, err, exception.ProductWithNoStock{Id: order.GetProductId()})
}

func setUpOrderCreateCmd(t *testing.T) (*CreateOrder, *mock.InterfaceMocks) {
	mocks := mock.NewInterfaceMocks(t)
	return NewCreateOrder(mocks.OrderRepo), mocks
}
