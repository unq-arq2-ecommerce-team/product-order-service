package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/action/command"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/model"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/model/exception"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/infrastructure/dto"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/infrastructure/logger"
	"net/http"
)

// CreateProductHandler
// @Summary      Endpoint create product
// @Description  create product
// @Param sellerId path int true "Seller ID" minimum(1)
// @Param Product body dto.ProductCreateReq true "It is a product creation request."
// @Tags         Product
// @Produce json
// @Success 200 {object} dto.IdResponse
// @Failure 400 {object} dto.ErrorMessage
// @Failure 404 {object} dto.ErrorMessage
// @Router       /api/v1/seller/{sellerId}/product [post]
func CreateProductHandler(log model.Logger, createProductCmd *command.CreateProduct) gin.HandlerFunc {
	return func(c *gin.Context) {
		sellerId, err := parsePathParamPositiveIntId(c, "sellerId")
		if err != nil {
			log.WithFields(logger.Fields{"error": err}).Error("invalid path param")
			writeJsonErrorMessageWithNoDesc(c, http.StatusBadRequest, err)
			return
		}
		var request dto.ProductCreateReq
		if err := c.BindJSON(&request); err != nil {
			writeJsonErrorMessageInDescAndMessage(c, http.StatusBadRequest, "invalid json body product create request", err)
			return
		}
		productId, err := createProductCmd.Do(c.Request.Context(), request.MapToModel(sellerId))
		if err != nil {
			switch err.(type) {
			case exception.SellerNotFound:
				writeJsonErrorMessageWithNoDesc(c, http.StatusNotFound, err)
			default:
				defaultInternalServerError(log, c, "uncaught error when create product", err)
			}
			return
		}
		c.JSON(http.StatusOK, dto.NewIdResponse(productId))
	}
}
