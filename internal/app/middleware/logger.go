package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"urlshortener/internal/app/config"
)

func Logger(l *zap.Logger, cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		// Сведения о запросах должны содержать URI, метод запроса и время, затраченное на его выполнение.
		latency := time.Since(startTime)
		l.Info("Request",
			zap.String("path", c.Request.URL.Path),
			zap.String("method", c.Request.Method),
			zap.Duration("duration", latency),
		)

		// Сведения об ответах должны содержать код статуса и размер содержимого ответа.
		status := c.Writer.Status()
		size := c.Writer.Size()
		l.Info("Response",
			zap.Int("status_code", status),
			zap.Int("size", size),
		)
	}
}
