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

func Test_GivenUpdateSellerCmdAndSellerIdAndSellerUpdate_WhenDo_ThenReturnNoError(t *testing.T) {
	updateSellerCmd, mocks := setUpSellerUpdateCmd(t)
	ctx := context.Background()
	sellerIdToFind := int64(4)
	sellerToUpdate := &model.Seller{
		Id:       sellerIdToFind,
		Name:     "pepeGrillo",
		Email:    "pepe_grillo@gmail.com",
		Products: []model.Product{},
	}
	updateSellerInfo := model.UpdateSeller{
		Name:  "fafafa",
		Email: "fafafa@mail.com",
	}
	mocks.SellerRepo.EXPECT().FindById(ctx, sellerIdToFind).Return(sellerToUpdate, nil)
	mocks.ProductRepo.EXPECT().FindAllBySellerId(ctx, sellerIdToFind).Return([]model.Product{}, nil)
	sellerToUpdate.Merge(updateSellerInfo)
	mocks.SellerRepo.EXPECT().Update(ctx, *sellerToUpdate).Return(true, nil)

	err := updateSellerCmd.Do(ctx, sellerIdToFind, updateSellerInfo)
	assert.NoError(t, err)
}

func Test_GivenUpdateSellerCmdAndSellerIdAndSellerUpdateAndAnyErrorInFindSeller_WhenDo_ThenReturnThatError(t *testing.T) {
	updateSellerCmd, mocks := setUpSellerUpdateCmd(t)
	ctx := context.Background()
	sellerIdToFind := int64(4)
	updateSellerInfo := model.UpdateSeller{
		Name:  "fafafa",
		Email: "fafafa@mail.com",
	}
	mocks.SellerRepo.EXPECT().FindById(ctx, sellerIdToFind).Return(nil, exception.SellerNotFound{Id: sellerIdToFind})
	mocks.ProductRepo.EXPECT().FindAllBySellerId(ctx, sellerIdToFind).Return([]model.Product{}, nil)
	err := updateSellerCmd.Do(ctx, sellerIdToFind, updateSellerInfo)
	assert.ErrorIs(t, err, exception.SellerNotFound{Id: sellerIdToFind})
}

func Test_GivenUpdateSellerCmdAndSellerIdAndSellerUpdateAndAnyErrorInFindSellerProducts_WhenDo_ThenReturnThatError(t *testing.T) {
	updateSellerCmd, mocks := setUpSellerUpdateCmd(t)
	ctx := context.Background()
	sellerIdToFind := int64(4)
	sellerToUpdate := &model.Seller{
		Id:    sellerIdToFind,
		Name:  "pepeGrillo",
		Email: "pepe_grillo@gmail.com",
	}
	updateSellerInfo := model.UpdateSeller{
		Name:  "fafafa",
		Email: "fafafa@mail.com",
	}
	msgErr := "unexpected error"

	mocks.SellerRepo.EXPECT().FindById(ctx, sellerIdToFind).Return(sellerToUpdate, nil)
	mocks.ProductRepo.EXPECT().FindAllBySellerId(ctx, sellerIdToFind).Return(nil, fmt.Errorf(msgErr))
	err := updateSellerCmd.Do(ctx, sellerIdToFind, updateSellerInfo)
	assert.EqualError(t, err, msgErr)
}

func Test_GivenUpdateSellerCmdAndSellerIdAndSellerUpdateAndAnyErrorInUpdateSeller_WhenDo_ThenReturnThatError(t *testing.T) {
	updateSellerCmd, mocks := setUpSellerUpdateCmd(t)
	ctx := context.Background()
	sellerIdToFind := int64(4)
	sellerToUpdate := &model.Seller{
		Id:       sellerIdToFind,
		Name:     "pepeGrillo",
		Email:    "pepe_grillo@gmail.com",
		Products: []model.Product{},
	}
	updateSellerInfo := model.UpdateSeller{
		Name:  "fafafa",
		Email: "fafafa@mail.com",
	}
	mocks.SellerRepo.EXPECT().FindById(ctx, sellerIdToFind).Return(sellerToUpdate, nil)
	mocks.ProductRepo.EXPECT().FindAllBySellerId(ctx, sellerIdToFind).Return([]model.Product{}, nil)
	sellerToUpdate.Merge(updateSellerInfo)
	mocks.SellerRepo.EXPECT().Update(ctx, *sellerToUpdate).Return(false, exception.SellerAlreadyExist{Name: updateSellerInfo.Name})

	err := updateSellerCmd.Do(ctx, sellerIdToFind, updateSellerInfo)
	assert.ErrorIs(t, err, exception.SellerAlreadyExist{Name: updateSellerInfo.Name})
}

func setUpSellerUpdateCmd(t *testing.T) (*UpdateSeller, *mock.InterfaceMocks) {
	mocks := mock.NewInterfaceMocks(t)
	return NewUpdateSeller(mocks.SellerRepo, *query.NewFindSellerById(mocks.SellerRepo, mocks.ProductRepo)), mocks
}
