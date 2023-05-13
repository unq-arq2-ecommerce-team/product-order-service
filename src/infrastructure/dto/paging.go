package dto

import "github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/model"

type PagingParamQuery struct {
	Page     int `form:"page"`
	PageSize int `form:"pageSize"`
}

func (pq *PagingParamQuery) MapToPageRequest() model.PagingRequest {
	return model.NewPagingRequest(pq.Page, pq.PageSize)
}
