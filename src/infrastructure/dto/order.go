package dto

import (
	"fmt"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/model"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/model/exception"
	"time"
)

type OrderCreateReq struct {
	CustomerId      int64         `json:"customerId" binding:"required,min=1"`
	ProductId       int64         `json:"productId" binding:"required,min=1"`
	DeliveryDate    time.Time     `json:"deliveryDate" time_format:"2006-01-02T15:04:05.000Z" example:"2090-04-20T15:04:05.000Z" binding:"required"`
	DeliveryAddress model.Address `json:"deliveryAddress" binding:"required"`
}

func (req *OrderCreateReq) Validate() error {
	timeNow := time.Now()
	if req.DeliveryDate.Before(timeNow) {
		return fmt.Errorf("invalid delivery date because is before to %s", timeNow)
	}
	return nil
}

type OrderDTO struct {
	Id              int64          `json:"id" bson:"_id"`
	CustomerId      int64          `json:"customerId" bson:"customerId"`
	CreatedOn       time.Time      `json:"createdOn" bson:"createdOn"`
	UpdatedOn       time.Time      `json:"updatedOn" bson:"updatedOn"`
	DeliveryDate    time.Time      `json:"deliveryDate" bson:"deliveryDate"`
	State           string         `json:"state" bson:"state"`
	Product         *model.Product `json:"product" bson:"product"`
	DeliveryAddress model.Address  `json:"deliveryAddress" bson:"deliveryAddress"`
}

func NewOrderDTOFrom(order model.Order) *OrderDTO {
	return &OrderDTO{
		Id:              order.Id,
		CustomerId:      order.CustomerId,
		CreatedOn:       order.CreatedOn,
		UpdatedOn:       order.UpdatedOn,
		DeliveryDate:    order.DeliveryDate,
		State:           order.StateAsString(),
		Product:         order.Product,
		DeliveryAddress: order.DeliveryAddress,
	}
}

func (dto *OrderDTO) Map() (model.Order, error) {
	orderState, exists := model.GetStateByString(dto.State)
	if !exists {
		return model.Order{}, exception.CannotMapOrderState{State: dto.State}
	}
	return model.Order{
		Id:              dto.Id,
		CustomerId:      dto.CustomerId,
		CreatedOn:       dto.CreatedOn,
		UpdatedOn:       dto.UpdatedOn,
		DeliveryDate:    dto.DeliveryDate,
		State:           orderState,
		Product:         dto.Product,
		DeliveryAddress: dto.DeliveryAddress,
	}, nil
}
