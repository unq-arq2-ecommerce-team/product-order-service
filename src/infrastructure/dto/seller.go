package dto

import "github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/model"

type SellerCreateReq struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

func (req *SellerCreateReq) MapToModel() model.Seller {
	return model.Seller{
		Name:  req.Name,
		Email: req.Email,
	}
}
