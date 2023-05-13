package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/action/command"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/model"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/infrastructure/logger"
	"net/http"
)

// DeleteAllBySellerHandler
// @Summary      Endpoint delete all products of seller
// @Description  delete all products by seller id
// @Param sellerId path int true "Seller ID" minimum(1)
// @Tags         Product
// @Produce json
// @Success 204
// @Failure 400 {object} dto.ErrorMessage
// @Router       /api/v1/seller/{sellerId}/product/all [delete]
func DeleteAllBySellerHandler(log model.Logger, deleteAllProductsBySellerIdCmd *command.DeleteAllProductsBySellerId) gin.HandlerFunc {
	return func(c *gin.Context) {
		sellerId, err := parsePathParamPositiveIntId(c, "sellerId")
		if err != nil {
			log.WithFields(logger.Fields{"error": err}).Error("invalid path param")
			writeJsonErrorMessageWithNoDesc(c, http.StatusBadRequest, err)
			return
		}
		err = deleteAllProductsBySellerIdCmd.Do(c.Request.Context(), sellerId)
		if err != nil {
			defaultInternalServerError(log, c, "uncaught error when delete product", err)
			return
		}
		c.Status(http.StatusNoContent)
	}
}
