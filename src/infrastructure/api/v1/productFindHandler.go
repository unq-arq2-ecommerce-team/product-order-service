package v1

import (
	"github.com/cassa10/arq2-tp1/src/domain/action/query"
	"github.com/cassa10/arq2-tp1/src/domain/model"
	"github.com/cassa10/arq2-tp1/src/domain/model/exception"
	"github.com/cassa10/arq2-tp1/src/infrastructure/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

// FindProductHandler
// @Summary      Endpoint find product
// @Description  find product
// @Param productId path int true "Product ID" minimum(1)
// @Tags         Product
// @Produce json
// @Success 200 {object} model.Product
// @Success 400 {object} dto.ErrorMessage
// @Failure 404 {object} dto.ErrorMessage
// @Router       /api/v1/seller/product/{productId} [get]
func FindProductHandler(log model.Logger, findProductByIdQuery *query.FindProductById) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := parsePathParamPositiveIntId(c, "productId")
		if err != nil {
			log.WithFields(logger.Fields{"error": err}).Error("invalid path param")
			writeJsonErrorMessageWithNoDesc(c, http.StatusBadRequest, err)
			return
		}
		product, err := findProductByIdQuery.Do(c.Request.Context(), id)
		if err != nil {
			switch err.(type) {
			case exception.ProductNotFound:
				writeJsonErrorMessageWithNoDesc(c, http.StatusNotFound, err)
			default:
				defaultInternalServerError(log, c, "uncaught error when find product", err)
			}
			return
		}
		c.JSON(http.StatusOK, product)
	}
}
