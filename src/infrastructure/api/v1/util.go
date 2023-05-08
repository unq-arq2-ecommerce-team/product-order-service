package v1

import (
	"fmt"
	"github.com/cassa10/arq2-tp1/src/domain/model"
	"github.com/cassa10/arq2-tp1/src/infrastructure/dto"
	"github.com/cassa10/arq2-tp1/src/infrastructure/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func parsePathParamPositiveIntId(c *gin.Context, paramKey string) (int64, error) {
	_idParam, _ := c.Params.Get(paramKey)
	id, err := strconv.ParseInt(_idParam, 10, 64)
	if err != nil || id <= 0 {
		return 0, fmt.Errorf("invalid path param %s as positive int64", paramKey)
	}
	return id, err
}

func defaultInternalServerError(log model.Logger, ginContext *gin.Context, additionalLogInfo string, err error) {
	log.WithFields(logger.Fields{"error": err}).Error(additionalLogInfo)
	_writeJsonErrorMessageWithDesc(ginContext, http.StatusInternalServerError, "internal server error", "")
}

func writeJsonErrorMessageWithNoDesc(c *gin.Context, status int, err error) {
	_writeJsonErrorMessageWithDesc(c, status, err.Error(), "")
}

func writeJsonErrorMessageInDescAndMessage(c *gin.Context, status int, msg string, err error) {
	_writeJsonErrorMessageWithDesc(c, status, msg, err.Error())
}

func _writeJsonErrorMessageWithDesc(c *gin.Context, status int, message, desc string) {
	c.JSON(status, dto.NewErrorMessage(message, desc))
}
