package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Logger middleware для логирования запросов
func Logger(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next()

		duration := time.Since(start)
		statusCode := c.Writer.Status()

		logger.WithFields(logrus.Fields{
			"method":      method,
			"path":        path,
			"status":      statusCode,
			"duration_ms": duration.Milliseconds(),
			"client_ip":   c.ClientIP(),
		}).Info("HTTP request")
	}
}

