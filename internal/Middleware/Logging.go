package Middleware

import (
	"github.com/akmyrzza/yandex-url/internal/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func Logging(logger *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		logger.LoggerZap.Info("Request",
			zap.String("path", c.Request.URL.Path),
			zap.String("method", c.Request.Method),
			zap.Duration("duration", time.Since(start)),
		)

		logger.LoggerZap.Info("Responese",
			zap.Int("statusCode", c.Writer.Status()),
			zap.Int64("contentLength", int64(c.Writer.Size())),
		)
	}
}
