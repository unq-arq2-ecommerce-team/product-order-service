package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/action/command"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/model"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/model/exception"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/infrastructure/logger"
	"net/http"
)

// UpdateProductHandler
// @Summary      Endpoint update product
// @Description  update product
// @Param productId path int true "Product ID" minimum(1)
// @Param Product body model.UpdateProduct true "It is a product updatable info."
// @Tags         Product
// @Produce json
// @Success 204
// @Failure 400 {object} dto.ErrorMessage
// @Failure 404 {object} dto.ErrorMessage
// @Failure 406 {object} dto.ErrorMessage
// @Router       /api/v1/seller/product/{productId} [put]
func UpdateProductHandler(log model.Logger, updateProductCmd *command.UpdateProduct) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := parsePathParamPositiveIntId(c, "productId")
		if err != nil {
			log.WithFields(logger.Fields{"error": err}).Error("invalid path param")
			writeJsonErrorMessageWithNoDesc(c, http.StatusBadRequest, err)
			return
		}
		var request model.UpdateProduct
		if err := c.BindJSON(&request); err != nil {
			writeJsonErrorMessageInDescAndMessage(c, http.StatusBadRequest, "invalid json body update product", err)
			return
		}
		err = updateProductCmd.Do(c.Request.Context(), id, request)
		if err != nil {
			switch err.(type) {
			case exception.ProductNotFound:
				writeJsonErrorMessageWithNoDesc(c, http.StatusNotFound, err)
			case exception.ProductCannotUpdate:
				writeJsonErrorMessageWithNoDesc(c, http.StatusNotAcceptable, err)
			default:
				defaultInternalServerError(log, c, "uncaught error when update product", err)
			}
			return
		}
		c.Status(http.StatusNoContent)
	}
}
