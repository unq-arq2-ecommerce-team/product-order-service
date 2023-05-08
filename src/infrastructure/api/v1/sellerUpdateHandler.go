package v1

import (
	"github.com/cassa10/arq2-tp1/src/domain/action/command"
	"github.com/cassa10/arq2-tp1/src/domain/model"
	"github.com/cassa10/arq2-tp1/src/domain/model/exception"
	"github.com/cassa10/arq2-tp1/src/infrastructure/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UpdateSellerHandler
// @Summary      Endpoint update seller
// @Description  update seller
// @Param sellerId path int true "Seller ID" minimum(1)
// @Param Seller body model.UpdateSeller true "It is a seller updatable info."
// @Tags         Seller
// @Produce json
// @Success 204
// @Failure 400 {object} dto.ErrorMessage
// @Failure 404 {object} dto.ErrorMessage
// @Failure 406 {object} dto.ErrorMessage
// @Router       /api/v1/seller/{sellerId} [put]
func UpdateSellerHandler(log model.Logger, updateSellerCmd *command.UpdateSeller) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := parsePathParamPositiveIntId(c, "sellerId")
		if err != nil {
			log.WithFields(logger.Fields{"error": err}).Error("invalid path param")
			writeJsonErrorMessageWithNoDesc(c, http.StatusBadRequest, err)
			return
		}
		var request model.UpdateSeller
		if err := c.BindJSON(&request); err != nil {
			writeJsonErrorMessageInDescAndMessage(c, http.StatusBadRequest, "invalid json body update seller request", err)
			return
		}
		err = updateSellerCmd.Do(c.Request.Context(), id, request)
		if err != nil {
			switch err.(type) {
			case exception.SellerNotFound:
				writeJsonErrorMessageWithNoDesc(c, http.StatusNotFound, err)
			case exception.SellerCannotUpdate, exception.SellerAlreadyExist:
				writeJsonErrorMessageWithNoDesc(c, http.StatusNotAcceptable, err)
			default:
				defaultInternalServerError(log, c, "uncaught error when update seller", err)
			}
			return
		}
		c.Status(http.StatusNoContent)
	}
}
