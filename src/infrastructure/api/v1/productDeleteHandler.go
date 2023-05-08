package v1

import (
	"github.com/cassa10/arq2-tp1/src/domain/action/command"
	"github.com/cassa10/arq2-tp1/src/domain/model"
	"github.com/cassa10/arq2-tp1/src/domain/model/exception"
	"github.com/cassa10/arq2-tp1/src/infrastructure/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteProductHandler
// @Summary      Endpoint delete product
// @Description  delete product by id
// @Param productId path int true "Product ID" minimum(1)
// @Tags         Product
// @Produce json
// @Success 204
// @Failure 400 {object} dto.ErrorMessage
// @Failure 404 {object} dto.ErrorMessage
// @Failure 406 {object} dto.ErrorMessage
// @Router       /api/v1/seller/product/{productId} [delete]
func DeleteProductHandler(log model.Logger, deleteProductCmd *command.DeleteProduct) gin.HandlerFunc {
	return func(c *gin.Context) {
		productId, err := parsePathParamPositiveIntId(c, "productId")
		if err != nil {
			log.WithFields(logger.Fields{"error": err}).Error("invalid path param")
			writeJsonErrorMessageWithNoDesc(c, http.StatusBadRequest, err)
			return
		}
		err = deleteProductCmd.Do(c.Request.Context(), productId)
		if err != nil {
			switch err.(type) {
			case exception.ProductNotFound:
				writeJsonErrorMessageWithNoDesc(c, http.StatusNotFound, err)
			case exception.ProductCannotDelete:
				writeJsonErrorMessageWithNoDesc(c, http.StatusNotAcceptable, err)
			default:
				defaultInternalServerError(log, c, "uncaught error when delete product", err)
			}
			return
		}
		c.Status(http.StatusNoContent)
	}
}
