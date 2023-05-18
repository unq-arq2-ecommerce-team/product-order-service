package usecase

import (
	"context"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/action/command"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/action/query"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/model"
)

type ConfirmOrder struct {
	baseLogger         model.Logger
	findOrderByIdQuery query.FindOrderById
	confirmOrderCmd    command.ConfirmOrder
}

func NewConfirmOrder(baseLogger model.Logger, confirmOrderCmd command.ConfirmOrder, findOrderByIdQuery query.FindOrderById) *ConfirmOrder {
	return &ConfirmOrder{
		baseLogger:         baseLogger.WithFields(model.LoggerFields{"useCase": "ConfirmOrder"}),
		confirmOrderCmd:    confirmOrderCmd,
		findOrderByIdQuery: findOrderByIdQuery,
	}
}

func (u ConfirmOrder) Do(ctx context.Context, orderId int64) error {
	log := u.baseLogger.WithRequestId(ctx).WithFields(model.LoggerFields{"orderId": orderId})
	order, err := u.findOrderByIdQuery.Do(ctx, orderId)
	if err != nil {
		log.WithFields(model.LoggerFields{"error": err}).Error("error when find order")
		return err
	}
	log = log.WithFields(model.LoggerFields{"orderState": order.State})
	err = u.confirmOrderCmd.Do(ctx, order)
	if err != nil {
		log.WithFields(model.LoggerFields{"error": err}).Error("error when confirm order")
		return err
	}
	log.Info("successful order confirmed")
	return nil
}
