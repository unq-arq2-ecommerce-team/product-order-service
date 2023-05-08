package model

import "github.com/cassa10/arq2-tp1/src/domain/util"

const (
	defaultPage  = 0
	minPageLimit = 1
	maxPageLimit = 9999999999999

	defaultPageSize  = 10
	minPageSizeLimit = 1
	maxPageSizeLimit = 500
)

type PagingRequest struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

func NewPagingRequest(page, size int) PagingRequest {
	return PagingRequest{
		Page: getPage(page),
		Size: getSize(size),
	}
}

func getPage(page int) int {
	if page < minPageLimit || page > maxPageLimit {
		return defaultPage
	}
	return page - 1
}

func getSize(size int) int {
	if size < minPageSizeLimit || size > maxPageSizeLimit {
		return defaultPageSize
	}
	return size
}

func (pr PagingRequest) String() string {
	return util.ParseStruct("PagingRequest", pr)
}

type Paging struct {
	Total       int `json:"total"`
	PageSize    int `json:"pageSize"`
	Pages       int `json:"pages"`
	CurrentPage int `json:"currentPage"`
}

func NewPaging(total int, pageSize int, pages int, currentPage int) Paging {
	return Paging{Total: total, PageSize: pageSize, Pages: pages, CurrentPage: currentPage}
}

func NewEmptyPage() Paging {
	return NewPaging(0, 0, 0, 0)
}

func (p Paging) String() string {
	return util.ParseStruct("Paging", p)
}
