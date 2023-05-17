package usecase

import (
	"context"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/action/command"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/action/query"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/model"
)

type DeliveredOrder struct {
	baseLogger         model.Logger
	findOrderByIdQuery query.FindOrderById
	deliveredOrderCmd  command.DeliveredOrder
}

func NewDeliveredOrder(baseLogger model.Logger, deliveredOrderCmd command.DeliveredOrder, findOrderByIdQuery query.FindOrderById) *DeliveredOrder {
	return &DeliveredOrder{
		baseLogger:         baseLogger.WithFields(model.LoggerFields{"useCase": "DeliveredOrder"}),
		deliveredOrderCmd:  deliveredOrderCmd,
		findOrderByIdQuery: findOrderByIdQuery,
	}
}

func (u DeliveredOrder) Do(ctx context.Context, orderId int64) error {
	log := u.baseLogger.WithFields(model.LoggerFields{"orderId": orderId})
	order, err := u.findOrderByIdQuery.Do(ctx, orderId)
	if err != nil {
		log.WithFields(model.LoggerFields{"error": err}).Error("error when find order")
		return err
	}
	log = log.WithFields(model.LoggerFields{"orderState": order.State})
	err = u.deliveredOrderCmd.Do(ctx, order)
	if err != nil {
		log.WithFields(model.LoggerFields{"error": err}).Error("error when delivered order")
		return err
	}
	log.Info("successful order delivered")
	return nil
}
