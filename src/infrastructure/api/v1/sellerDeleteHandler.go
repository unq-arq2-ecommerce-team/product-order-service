package v1

import (
	"github.com/cassa10/arq2-tp1/src/domain/action/command"
	"github.com/cassa10/arq2-tp1/src/domain/model"
	"github.com/cassa10/arq2-tp1/src/domain/model/exception"
	"github.com/cassa10/arq2-tp1/src/infrastructure/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteSellerHandler
// @Summary      Endpoint delete seller
// @Description  delete seller by id
// @Param sellerId path int true "Seller ID" minimum(1)
// @Tags         Seller
// @Produce json
// @Success 204
// @Failure 400 {object} dto.ErrorMessage
// @Failure 404 {object} dto.ErrorMessage
// @Failure 406 {object} dto.ErrorMessage
// @Router       /api/v1/seller/{sellerId} [delete]
func DeleteSellerHandler(log model.Logger, deleteSellerCmd *command.DeleteSeller) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := parsePathParamPositiveIntId(c, "sellerId")
		if err != nil {
			log.WithFields(logger.Fields{"error": err}).Error("invalid path param")
			writeJsonErrorMessageWithNoDesc(c, http.StatusBadRequest, err)
			return
		}
		err = deleteSellerCmd.Do(c.Request.Context(), id)
		if err != nil {
			switch err.(type) {
			case exception.SellerNotFound:
				writeJsonErrorMessageWithNoDesc(c, http.StatusNotFound, err)
			case exception.SellerCannotDelete:
				writeJsonErrorMessageWithNoDesc(c, http.StatusNotAcceptable, err)
			default:
				defaultInternalServerError(log, c, "uncaught error when delete seller", err)
			}
			return
		}
		c.Status(http.StatusNoContent)
	}
}
