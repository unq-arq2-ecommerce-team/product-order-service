package usecase

import (
	"context"
	"github.com/cassa10/arq2-tp1/src/domain/action/command"
	"github.com/cassa10/arq2-tp1/src/domain/action/query"
	"github.com/cassa10/arq2-tp1/src/domain/model"
	"github.com/cassa10/arq2-tp1/src/infrastructure/logger"
)

type DeliveredOrder struct {
	baseLogger         model.Logger
	findOrderByIdQuery query.FindOrderById
	deliveredOrderCmd  command.DeliveredOrder
}

func NewDeliveredOrder(baseLogger model.Logger, deliveredOrderCmd command.DeliveredOrder, findOrderByIdQuery query.FindOrderById) *DeliveredOrder {
	return &DeliveredOrder{
		baseLogger:         baseLogger.WithFields(logger.Fields{"useCase": "DeliveredOrder"}),
		deliveredOrderCmd:  deliveredOrderCmd,
		findOrderByIdQuery: findOrderByIdQuery,
	}
}

func (u DeliveredOrder) Do(ctx context.Context, orderId int64) error {
	log := u.baseLogger.WithFields(logger.Fields{"orderId": orderId})
	order, err := u.findOrderByIdQuery.Do(ctx, orderId)
	if err != nil {
		log.WithFields(logger.Fields{"error": err}).Error("error when find order")
		return err
	}
	log = log.WithFields(logger.Fields{"orderState": order.State})
	err = u.deliveredOrderCmd.Do(ctx, order)
	if err != nil {
		log.WithFields(logger.Fields{"error": err}).Error("error when delivered order")
		return err
	}
	log.Info("successful order delivered")
	return nil
}
