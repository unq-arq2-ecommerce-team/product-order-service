package middleware

import (
	"github.com/cassa10/arq2-tp1/src/infrastructure/logger"
	"github.com/gin-gonic/gin"
)

const headerRequestId = "system-request-id"

func TracingRequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := logger.SetRequestId(c.Request.Context(), c.Request.Header.Get(headerRequestId))
		c.Request = c.Request.WithContext(ctx)
		c.Writer.Header().Set("X-Request-Id", logger.GetRequestId(ctx))
	}
}
