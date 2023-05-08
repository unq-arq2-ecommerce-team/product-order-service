package dto

import "github.com/cassa10/arq2-tp1/src/domain/model"

type PagingParamQuery struct {
	Page     int `form:"page"`
	PageSize int `form:"pageSize"`
}

func (pq *PagingParamQuery) MapToPageRequest() model.PagingRequest {
	return model.NewPagingRequest(pq.Page, pq.PageSize)
}
