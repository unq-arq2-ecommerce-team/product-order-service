package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_PagingRequest_New_WithDefaults(t *testing.T) {
	page1, size1 := 0, 0
	pagingRequestFromNew1 := NewPagingRequest(page1, size1)

	page2, size2 := -1, -1
	pagingRequestFromNew2 := NewPagingRequest(page2, size2)

	page3, size3 := maxPageLimit+1, maxPageSizeLimit+1
	pagingRequestFromNew3 := NewPagingRequest(page3, size3)

	pagingRequestDefault := PagingRequest{
		Page: defaultPage,
		Size: defaultPageSize,
	}
	assert.Equal(t, pagingRequestDefault, pagingRequestFromNew1)
	assert.Equal(t, pagingRequestDefault, pagingRequestFromNew2)
	assert.Equal(t, pagingRequestDefault, pagingRequestFromNew3)
}

func Test_PagingRequest_New(t *testing.T) {
	page1, size1 := maxPageLimit, maxPageSizeLimit
	pagingRequestFromNew1 := NewPagingRequest(page1, size1)
	pagingRequest1 := PagingRequest{
		Page: page1 - 1,
		Size: size1,
	}
	page2, size2 := 5, 250
	pagingRequestFromNew2 := NewPagingRequest(page2, size2)
	pagingRequest2 := PagingRequest{
		Page: page2 - 1,
		Size: size2,
	}

	assert.Equal(t, pagingRequest1, pagingRequestFromNew1)
	assert.Equal(t, pagingRequest2, pagingRequestFromNew2)
}

func Test_PagingRequest_String(t *testing.T) {
	pagingRequest1 := NewPagingRequest(2, 20)
	pagingRequest2 := NewPagingRequest(5, 250)
	assert.Equal(t, `[PagingRequest]{"page":1,"size":20}`, pagingRequest1.String())
	assert.Equal(t, `[PagingRequest]{"page":4,"size":250}`, pagingRequest2.String())
}

func Test_Paging_New(t *testing.T) {
	total1, pageSize1, pages1, currentPage1 := 60, 20, 3, 1
	pagingFromNew1 := NewPaging(total1, pageSize1, pages1, currentPage1)
	paging1 := Paging{
		Total:       total1,
		PageSize:    pageSize1,
		Pages:       pages1,
		CurrentPage: currentPage1,
	}
	total2, pageSize2, pages2, currentPage2 := 20, 5, 4, 0
	pagingFromNew2 := NewPaging(total2, pageSize2, pages2, currentPage2)
	paging2 := Paging{
		Total:       total2,
		PageSize:    pageSize2,
		Pages:       pages2,
		CurrentPage: currentPage2,
	}
	assert.Equal(t, paging1, pagingFromNew1)
	assert.Equal(t, paging2, pagingFromNew2)
}

func Test_Paging_EmptyPage(t *testing.T) {
	emptyPage := NewEmptyPage()
	paging := Paging{
		Total:       0,
		PageSize:    0,
		Pages:       0,
		CurrentPage: 0,
	}
	assert.Equal(t, paging, emptyPage)
}

func Test_Paging_String(t *testing.T) {
	paging1 := NewPaging(60, 20, 3, 1)
	paging2 := NewPaging(20, 5, 4, 0)
	assert.Equal(t, `[Paging]{"total":60,"pageSize":20,"pages":3,"currentPage":1}`, paging1.String())
	assert.Equal(t, `[Paging]{"total":20,"pageSize":5,"pages":4,"currentPage":0}`, paging2.String())
}
