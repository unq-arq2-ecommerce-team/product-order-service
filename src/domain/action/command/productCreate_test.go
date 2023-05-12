package command

import (
	"context"
	"fmt"
	"github.com/cassa10/arq2-tp1/src/domain/action/query"
	"github.com/cassa10/arq2-tp1/src/domain/mock"
	"github.com/cassa10/arq2-tp1/src/domain/model"
	"github.com/cassa10/arq2-tp1/src/domain/model/exception"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GivenCreateProductCmdAndNewProduct_WhenDo_ThenReturnNoErrorAndANewId(t *testing.T) {
	createProductCmd, mocks := setUpProductCreateCmd(t)
	ctx := context.Background()
	sellerId := int64(123)
	product := model.Product{
		SellerId:    sellerId,
		Name:        "Jabon",
		Description: "Un jabon lindo",
		Price:       0.50,
		Category:    "c1",
		Stock:       60,
	}
	newProductId := int64(666)
	mocks.SellerRepo.EXPECT().FindById(ctx, sellerId).Return(&model.Seller{Id: sellerId}, nil)
	mocks.ProductRepo.EXPECT().FindAllBySellerId(ctx, sellerId).Return([]model.Product{}, nil)
	mocks.ProductRepo.EXPECT().Create(ctx, product).Return(newProductId, nil)

	resProductId, err := createProductCmd.Do(ctx, product)

	assert.Equal(t, newProductId, resProductId)
	assert.NoError(t, err)
}

func Test_GivenCreateProductCmdAndNewProductAndProductRepoCreateError_WhenDo_ThenReturnThatError(t *testing.T) {
	createProductCmd, mocks := setUpProductCreateCmd(t)
	ctx := context.Background()
	sellerId := int64(123)
	product := model.Product{
		SellerId:    sellerId,
		Name:        "Jabon",
		Description: "Un jabon lindo",
		Price:       0.50,
		Category:    "c1",
		Stock:       60,
	}
	msgErr := "unexpected error"
	mocks.SellerRepo.EXPECT().FindById(ctx, sellerId).Return(&model.Seller{Id: sellerId}, nil)
	mocks.ProductRepo.EXPECT().FindAllBySellerId(ctx, sellerId).Return([]model.Product{}, nil)
	mocks.ProductRepo.EXPECT().Create(ctx, product).Return(int64(0), fmt.Errorf(msgErr))

	resProductId, err := createProductCmd.Do(ctx, product)

	assert.Equal(t, int64(0), resProductId)
	assert.EqualError(t, err, msgErr)
}

func Test_GivenCreateProductCmdAndNewProductAndSellerRepoFindByIdError_WhenDo_ThenReturnThatError(t *testing.T) {
	createProductCmd, mocks := setUpProductCreateCmd(t)
	ctx := context.Background()
	sellerId := int64(123)
	product := model.Product{
		SellerId:    sellerId,
		Name:        "Jabon",
		Description: "Un jabon lindo",
		Price:       0.50,
		Category:    "c1",
		Stock:       60,
	}
	mocks.SellerRepo.EXPECT().FindById(ctx, sellerId).Return(nil, exception.SellerNotFound{Id: sellerId})

	resProductId, err := createProductCmd.Do(ctx, product)

	assert.Equal(t, int64(0), resProductId)
	assert.ErrorIs(t, err, exception.SellerNotFound{Id: sellerId})
}

func setUpProductCreateCmd(t *testing.T) (*CreateProduct, *mock.InterfaceMocks) {
	mocks := mock.NewInterfaceMocks(t)
	return NewCreateProduct(mocks.ProductRepo, *query.NewFindSellerById(mocks.SellerRepo)), mocks
}
