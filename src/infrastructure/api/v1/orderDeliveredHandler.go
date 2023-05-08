package v1

import (
	"github.com/cassa10/arq2-tp1/src/domain/model"
	"github.com/cassa10/arq2-tp1/src/domain/model/exception"
	"github.com/cassa10/arq2-tp1/src/domain/usecase"
	"github.com/cassa10/arq2-tp1/src/infrastructure/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeliveredOrderHandler
// @Summary      Endpoint delivered order
// @Description  delivered an order
// @Param orderId path int true "Order ID" minimum(1)
// @Tags         Order
// @Produce json
// @Success 204
// @Success 400 {object} dto.ErrorMessage
// @Failure 404 {object} dto.ErrorMessage
// @Failure 406 {object} dto.ErrorMessage
// @Failure 500 {object} dto.ErrorMessage
// @Router       /api/v1/order/{orderId}/delivered [post]
func DeliveredOrderHandler(log model.Logger, deliveredOrderUseCase *usecase.DeliveredOrder) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := parsePathParamPositiveIntId(c, "orderId")
		if err != nil {
			log.WithFields(logger.Fields{"error": err}).Error("invalid path param")
			writeJsonErrorMessageWithNoDesc(c, http.StatusBadRequest, err)
			return
		}
		err = deliveredOrderUseCase.Do(c.Request.Context(), id)
		if err != nil {
			switch err.(type) {
			case exception.OrderNotFound:
				writeJsonErrorMessageWithNoDesc(c, http.StatusNotFound, err)
			case exception.OrderInvalidTransitionState:
				writeJsonErrorMessageWithNoDesc(c, http.StatusNotAcceptable, err)
			case exception.CannotMapOrderState:
				writeJsonErrorMessageWithNoDesc(c, http.StatusInternalServerError, err)
			default:
				defaultInternalServerError(log, c, "uncaught error when delivered order", err)
			}
			return
		}
		c.Status(http.StatusNoContent)
	}
}
