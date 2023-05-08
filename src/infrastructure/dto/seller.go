package dto

import "github.com/cassa10/arq2-tp1/src/domain/model"

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
