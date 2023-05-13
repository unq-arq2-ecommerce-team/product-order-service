package command

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/action/query"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/mock"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/model"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/model/exception"
	"testing"
)

func Test_GivenUpdateProductCmdAndProductIdAndProductUpdate_WhenDo_ThenReturnNoError(t *testing.T) {
	updateProductCmd, mocks := setUpProductUpdateCmd(t)
	ctx := context.Background()
	productIdToFind := int64(4)
	productToUpdate := &model.Product{
		Id:          productIdToFind,
		Name:        "Jabon",
		Description: "Un jabon lindo",
		Category:    "Limpieza",
		Price:       0.50,
		Stock:       6,
	}
	updateProductInfo := model.UpdateProduct{
		Name:        "Updated Jabon",
		Description: "desc custom",
		Category:    "Limpieza x2",
		Price:       199.99,
	}
	mocks.ProductRepo.EXPECT().FindById(ctx, productIdToFind).Return(productToUpdate, nil)
	productToUpdate.Merge(updateProductInfo)
	mocks.ProductRepo.EXPECT().Update(ctx, *productToUpdate).Return(true, nil)

	err := updateProductCmd.Do(ctx, productIdToFind, updateProductInfo)
	assert.NoError(t, err)
}

func Test_GivenUpdateProductCmdAndProductIdAndProductUpdateAndAnyErrorInFindProduct_WhenDo_ThenReturnThatError(t *testing.T) {
	updateProductCmd, mocks := setUpProductUpdateCmd(t)
	ctx := context.Background()
	productIdToFind := int64(4)
	updateProductInfo := model.UpdateProduct{
		Name:        "Updated Jabon",
		Description: "desc custom",
		Category:    "Limpieza x2",
		Price:       199.99,
	}
	mocks.ProductRepo.EXPECT().FindById(ctx, productIdToFind).Return(nil, exception.ProductNotFound{Id: productIdToFind})
	err := updateProductCmd.Do(ctx, productIdToFind, updateProductInfo)
	assert.ErrorIs(t, err, exception.ProductNotFound{Id: productIdToFind})
}

func Test_GivenUpdateProductCmdAndProductIdAndProductUpdateAndAnyErrorInUpdateProduct_WhenDo_ThenReturnThatError(t *testing.T) {
	updateProductCmd, mocks := setUpProductUpdateCmd(t)
	ctx := context.Background()
	productIdToFind := int64(4)
	productToUpdate := &model.Product{
		Id:          productIdToFind,
		Name:        "Jabon",
		Description: "Un jabon lindo",
		Category:    "Limpieza",
		Price:       0.50,
		Stock:       6,
	}
	updateProductInfo := model.UpdateProduct{
		Name:        "Updated Jabon",
		Description: "desc custom",
		Category:    "Limpieza x2",
		Price:       199.99,
	}
	mocks.ProductRepo.EXPECT().FindById(ctx, productIdToFind).Return(productToUpdate, nil)
	productToUpdate.Merge(updateProductInfo)
	mocks.ProductRepo.EXPECT().Update(ctx, *productToUpdate).Return(false, exception.ProductCannotUpdate{Id: productIdToFind})

	err := updateProductCmd.Do(ctx, productIdToFind, updateProductInfo)
	assert.ErrorIs(t, err, exception.ProductCannotUpdate{Id: productIdToFind})
}

func setUpProductUpdateCmd(t *testing.T) (*UpdateProduct, *mock.InterfaceMocks) {
	mocks := mock.NewInterfaceMocks(t)
	return NewUpdateProduct(mocks.ProductRepo, *query.NewFindProductById(mocks.ProductRepo)), mocks
}
