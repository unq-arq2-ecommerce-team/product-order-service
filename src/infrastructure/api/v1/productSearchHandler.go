package v1

import (
	"github.com/cassa10/arq2-tp1/src/domain/action/query"
	"github.com/cassa10/arq2-tp1/src/domain/model"
	"github.com/cassa10/arq2-tp1/src/infrastructure/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SearchProductsHandler
// @Summary      Endpoint search products
// @Description  update product
// @Param        page    		query	integer	false	"page request"  		minimum(0) maximum(9999999999999)
// @Param        pageSize		query	integer	false	"pageSize request"  	minimum(1) maximum(500)
// @Param        name    		query	string	false	"filter by name"		example("name")
// @Param        category   	query	string	false	"filter by category"	example("category")
// @Param        priceMin    	query	number	false	"filter by min price"	minimum(0) maximum(999999999999999999)
// @Param        priceMax  		query	number	false	"filter by max price"	minimum(0) maximum(999999999999999999)
// @Tags         Product
// @Produce json
// @Success 200 {object} dto.ProductSearchResponse
// @Failure 400 {object} dto.ErrorMessage
// @Router       /api/v1/seller/product/search [get]
func SearchProductsHandler(log model.Logger, searchProductsQuery *query.SearchProducts) gin.HandlerFunc {
	return func(c *gin.Context) {
		var qs dto.ProductSearchQueryReq
		if err := c.ShouldBindQuery(&qs); err != nil {
			writeJsonErrorMessageWithNoDesc(c, http.StatusBadRequest, err)
			return
		}
		if err := qs.ValidateReq(); err != nil {
			writeJsonErrorMessageWithNoDesc(c, http.StatusBadRequest, err)
			return
		}
		products, paging, err := searchProductsQuery.Do(c.Request.Context(), qs.GetProductSearchFilter(), qs.MapToPageRequest())
		if err != nil {
			defaultInternalServerError(log, c, "uncaught error when update product", err)
			return
		}
		c.JSON(http.StatusOK, dto.NewProductSearchResponse(products, paging))
	}
}
