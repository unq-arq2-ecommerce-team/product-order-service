package usecase

import (
	"context"
	"github.com/cassa10/arq2-tp1/src/domain/action/command"
	"github.com/cassa10/arq2-tp1/src/domain/action/query"
	"github.com/cassa10/arq2-tp1/src/domain/model"
	"github.com/cassa10/arq2-tp1/src/domain/model/exception"
	"github.com/cassa10/arq2-tp1/src/infrastructure/logger"
	"time"
)

type CreateOrder struct {
	baseLogger            model.Logger
	createOrderCmd        command.CreateOrder
	findProductByIdQuery  query.FindProductById
	findCustomerByIdQuery query.FindCustomerById
}

func NewCreateOrder(baseLogger model.Logger, createOrderCmd command.CreateOrder, findProductByIdQuery query.FindProductById, findCustomerByIdQuery query.FindCustomerById) *CreateOrder {
	return &CreateOrder{
		baseLogger:            baseLogger.WithFields(logger.Fields{"useCase": "CreateOrder"}),
		createOrderCmd:        createOrderCmd,
		findProductByIdQuery:  findProductByIdQuery,
		findCustomerByIdQuery: findCustomerByIdQuery,
	}
}

func (u CreateOrder) Do(ctx context.Context, customerId, productId int64, deliveryDate time.Time, deliveryAddress model.Address) (int64, error) {
	log := u.baseLogger.WithFields(logger.Fields{"customerId": customerId, "productId": productId, "deliveryDate": deliveryDate, "deliveryAddress": deliveryAddress})
	_, err := u.findCustomerByIdQuery.Do(ctx, customerId)
	if err != nil {
		log.WithFields(logger.Fields{"error": err}).Errorf("error when find customer")
		return 0, err
	}
	product, err := u.findProductByIdQuery.Do(ctx, productId)
	if err != nil {
		log.WithFields(logger.Fields{"error": err}).Errorf("error when find product")
		return 0, err
	}
	if !product.ReduceStock() {
		log.Infof("product with stock %v is not available", product.Stock)
		return 0, exception.ProductWithNoStock{Id: productId}
	}
	order := model.NewOrder(customerId, product, deliveryDate, deliveryAddress)
	orderId, err := u.createOrderCmd.Do(ctx, order)
	if err != nil {
		log.WithFields(logger.Fields{"error": err, "order": order}).Errorf("error when create order")
		return 0, err
	}
	log.Infof("successful order created with id %v", orderId)
	return orderId, nil
}
