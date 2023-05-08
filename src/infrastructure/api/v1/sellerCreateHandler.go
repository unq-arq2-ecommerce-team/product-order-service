package v1

import (
	"github.com/cassa10/arq2-tp1/src/domain/action/command"
	"github.com/cassa10/arq2-tp1/src/domain/model"
	"github.com/cassa10/arq2-tp1/src/domain/model/exception"
	"github.com/cassa10/arq2-tp1/src/infrastructure/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateSellerHandler
// @Summary      Endpoint create seller
// @Description  create seller
// @Param Seller body dto.SellerCreateReq true "It is a seller creation request."
// @Tags         Seller
// @Produce json
// @Success 200 {object} dto.IdResponse
// @Failure 400 {object} dto.ErrorMessage
// @Failure 406 {object} dto.ErrorMessage
// @Router		/api/v1/seller [post]
func CreateSellerHandler(log model.Logger, createSellerCmd *command.CreateSeller) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request dto.SellerCreateReq
		if err := c.BindJSON(&request); err != nil {
			writeJsonErrorMessageInDescAndMessage(c, http.StatusBadRequest, "invalid json body seller create request", err)
			return
		}
		sellerId, err := createSellerCmd.Do(c.Request.Context(), request.MapToModel())
		if err != nil {
			switch err.(type) {
			case exception.SellerAlreadyExist:
				writeJsonErrorMessageWithNoDesc(c, http.StatusNotAcceptable, err)
			default:
				defaultInternalServerError(log, c, "uncaught error when create seller", err)
			}
			return
		}
		c.JSON(http.StatusOK, dto.NewIdResponse(sellerId))
	}
}
