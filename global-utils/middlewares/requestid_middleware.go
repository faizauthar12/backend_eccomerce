package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check for incoming header, use it if exists
		requestID := c.Request.Header.Get("X-Request-Id")

		// Create request id with UUID
		if requestID == "" {
			requestID = uuid.New().String()
		}

		// Expose it for use in the application
		c.Set("RequestId", requestID)
		// Set X-Request-Id header
		c.Writer.Header().Set("X-Request-Id", requestID)
		c.Next()
	}
}
