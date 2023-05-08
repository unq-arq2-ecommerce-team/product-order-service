package usecase

import (
	"context"
	"github.com/cassa10/arq2-tp1/src/domain/action/command"
	"github.com/cassa10/arq2-tp1/src/domain/action/query"
	"github.com/cassa10/arq2-tp1/src/domain/model"
	"github.com/cassa10/arq2-tp1/src/infrastructure/logger"
)

type ConfirmOrder struct {
	baseLogger         model.Logger
	findOrderByIdQuery query.FindOrderById
	confirmOrderCmd    command.ConfirmOrder
}

func NewConfirmOrder(baseLogger model.Logger, confirmOrderCmd command.ConfirmOrder, findOrderByIdQuery query.FindOrderById) *ConfirmOrder {
	return &ConfirmOrder{
		baseLogger:         baseLogger.WithFields(logger.Fields{"useCase": "ConfirmOrder"}),
		confirmOrderCmd:    confirmOrderCmd,
		findOrderByIdQuery: findOrderByIdQuery,
	}
}

func (u ConfirmOrder) Do(ctx context.Context, orderId int64) error {
	log := u.baseLogger.WithFields(logger.Fields{"orderId": orderId})
	order, err := u.findOrderByIdQuery.Do(ctx, orderId)
	if err != nil {
		log.WithFields(logger.Fields{"error": err}).Error("error when find order")
		return err
	}
	log = log.WithFields(logger.Fields{"orderState": order.State})
	err = u.confirmOrderCmd.Do(ctx, order)
	if err != nil {
		log.WithFields(logger.Fields{"error": err}).Error("error when confirm order")
		return err
	}
	log.Info("successful order confirmed")
	return nil
}
