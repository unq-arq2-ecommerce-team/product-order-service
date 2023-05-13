package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/action/query"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/model"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/model/exception"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/infrastructure/dto"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/infrastructure/logger"
	"net/http"
)

// FindOrderHandler
// @Summary      Endpoint find order
// @Description  find order
// @Param orderId path int true "Order ID" minimum(1)
// @Tags         Order
// @Produce json
// @Success 200 {object} dto.OrderDTO
// @Success 400 {object} dto.ErrorMessage
// @Failure 404 {object} dto.ErrorMessage
// @Failure 500 {object} dto.ErrorMessage
// @Router       /api/v1/order/{orderId} [get]
func FindOrderHandler(log model.Logger, findOrderByIdQuery *query.FindOrderById) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := parsePathParamPositiveIntId(c, "orderId")
		if err != nil {
			log.WithFields(logger.Fields{"error": err}).Error("invalid path param")
			writeJsonErrorMessageWithNoDesc(c, http.StatusBadRequest, err)
			return
		}
		order, err := findOrderByIdQuery.Do(c.Request.Context(), id)
		if err != nil {
			switch err.(type) {
			case exception.OrderNotFound:
				writeJsonErrorMessageWithNoDesc(c, http.StatusNotFound, err)
			case exception.CannotMapOrderState:
				writeJsonErrorMessageWithNoDesc(c, http.StatusInternalServerError, err)
			default:
				defaultInternalServerError(log, c, "uncaught error when find order", err)
			}
			return
		}
		c.JSON(http.StatusOK, dto.NewOrderDTOFrom(*order))
	}
}
