package query

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/mock"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/model"
	"testing"
)

func Test_GivenSearchProductsWithTheirFilterAndPagingReqAndProductsWhichMatch_WhenDoWithThatFiltersAndPagingReq_ThenReturnTheProductsAndTheirPagingAndNoError(t *testing.T) {
	searchProducts, mocks := setUpSearchProducts(t)
	ctx := context.Background()
	searchFilters := model.ProductSearchFilter{}
	pagingReq := model.NewPagingRequest(0, 5)
	productsExpected := []model.Product{
		{Id: 1},
		{Id: 2},
		{Id: 6},
		{Id: 25},
	}
	pagingExpected := model.NewPaging(4, 5, 1, 0)
	mocks.ProductRepo.EXPECT().Search(ctx, searchFilters, pagingReq).Return(productsExpected, pagingExpected, nil)

	products, paging, err := searchProducts.Do(ctx, searchFilters, pagingReq)

	assert.Equal(t, productsExpected, products)
	assert.Equal(t, pagingExpected, paging)
	assert.NoError(t, err)
}

func Test_GivenSearchProductsWithTheirFilterAndPagingReqAndNoProductsMatch_WhenDoWithThatFiltersAndPagingReq_ThenReturnNoProductsWithPagingAndNoError(t *testing.T) {
	searchProducts, mocks := setUpSearchProducts(t)
	ctx := context.Background()
	priceMin := float64(99999999999)
	searchFilters := model.ProductSearchFilter{
		PriceMin: &priceMin,
	}
	pagingReq := model.NewPagingRequest(0, 5)
	productsExpected := []model.Product{}
	pagingExpected := model.NewPaging(0, 5, 0, 0)
	mocks.ProductRepo.EXPECT().Search(ctx, searchFilters, pagingReq).Return(productsExpected, pagingExpected, nil)

	products, paging, err := searchProducts.Do(ctx, searchFilters, pagingReq)

	assert.Equal(t, productsExpected, products)
	assert.Equal(t, pagingExpected, paging)
	assert.NoError(t, err)
}

func Test_GivenSearchProductsWithProductRepoWithError_WhenDoWithSomeFiltersAndPagingReq_ThenReturnRepoError(t *testing.T) {
	searchProducts, mocks := setUpSearchProducts(t)
	ctx := context.Background()
	searchFilters := model.ProductSearchFilter{}
	pagingReq := model.NewPagingRequest(0, 5)
	errMsg := "unexpected error"
	pagingExpected := model.NewEmptyPage()
	mocks.ProductRepo.EXPECT().Search(ctx, searchFilters, pagingReq).Return(nil, pagingExpected, fmt.Errorf(errMsg))

	products, paging, err := searchProducts.Do(ctx, searchFilters, pagingReq)

	assert.Nil(t, products)
	assert.Equal(t, pagingExpected, paging)
	assert.EqualError(t, err, errMsg)
}

func setUpSearchProducts(t *testing.T) (*SearchProducts, *mock.InterfaceMocks) {
	mocks := mock.NewInterfaceMocks(t)
	return NewSearchProducts(mocks.ProductRepo), mocks
}
