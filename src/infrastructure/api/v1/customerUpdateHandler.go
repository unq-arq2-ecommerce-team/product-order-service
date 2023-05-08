package v1

import (
	"github.com/cassa10/arq2-tp1/src/domain/action/command"
	"github.com/cassa10/arq2-tp1/src/domain/model"
	"github.com/cassa10/arq2-tp1/src/domain/model/exception"
	"github.com/cassa10/arq2-tp1/src/infrastructure/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UpdateCustomerHandler
// @Summary      Endpoint update customer
// @Description  update customer
// @Param customerId path int true "Customer ID" minimum(1)
// @Param Customer body model.UpdateCustomer true "It is a customer updatable info."
// @Tags         Customer
// @Produce json
// @Success 204
// @Failure 400 {object} dto.ErrorMessage
// @Failure 404 {object} dto.ErrorMessage
// @Failure 406 {object} dto.ErrorMessage
// @Router       /api/v1/customer/{customerId} [put]
func UpdateCustomerHandler(log model.Logger, updateCustomerCmd *command.UpdateCustomer) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := parsePathParamPositiveIntId(c, "customerId")
		if err != nil {
			log.WithFields(logger.Fields{"error": err}).Error("invalid path param")
			writeJsonErrorMessageWithNoDesc(c, http.StatusBadRequest, err)
			return
		}
		var request model.UpdateCustomer
		if err := c.BindJSON(&request); err != nil {
			writeJsonErrorMessageInDescAndMessage(c, http.StatusBadRequest, "invalid json body update customer request", err)
			return
		}
		err = updateCustomerCmd.Do(c.Request.Context(), id, request)
		if err != nil {
			switch err.(type) {
			case exception.CustomerNotFound:
				writeJsonErrorMessageWithNoDesc(c, http.StatusNotFound, err)
			case exception.CustomerCannotUpdate, exception.CustomerAlreadyExist:
				writeJsonErrorMessageWithNoDesc(c, http.StatusNotAcceptable, err)
			default:
				defaultInternalServerError(log, c, "uncaught error when update customer", err)
			}
			return
		}
		c.Status(http.StatusNoContent)
	}
}
