package v1

import (
	"github.com/cassa10/arq2-tp1/src/domain/model"
	"github.com/cassa10/arq2-tp1/src/domain/model/exception"
	"github.com/cassa10/arq2-tp1/src/domain/usecase"
	"github.com/cassa10/arq2-tp1/src/infrastructure/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateOrderHandler
// @Summary      Endpoint create order
// @Description  create order
// @Param Order body dto.OrderCreateReq true "It is a order creation request."
// @Tags         Order
// @Produce json
// @Success 200 {object} dto.IdResponse
// @Failure 400 {object} dto.ErrorMessage
// @Failure 404 {object} dto.ErrorMessage
// @Failure 406 {object} dto.ErrorMessage
// @Router       /api/v1/order [post]
func CreateOrderHandler(log model.Logger, createOrder *usecase.CreateOrder) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.OrderCreateReq
		if err := c.BindJSON(&req); err != nil {
			writeJsonErrorMessageInDescAndMessage(c, http.StatusBadRequest, "invalid json body order create req", err)
			return
		}
		if err := req.Validate(); err != nil {
			writeJsonErrorMessageWithNoDesc(c, http.StatusBadRequest, err)
			return
		}
		orderId, err := createOrder.Do(c.Request.Context(), req.CustomerId, req.ProductId, req.DeliveryDate, req.DeliveryAddress)
		if err != nil {
			switch err.(type) {
			case exception.ProductNotFound:
				writeJsonErrorMessageWithNoDesc(c, http.StatusNotFound, err)
			case exception.ProductWithNoStock:
				writeJsonErrorMessageWithNoDesc(c, http.StatusNotAcceptable, err)
			default:
				defaultInternalServerError(log, c, "uncaught error when create order", err)
			}
			return
		}
		c.JSON(http.StatusOK, dto.NewIdResponse(orderId))
	}
}
