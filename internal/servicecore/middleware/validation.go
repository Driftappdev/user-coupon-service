package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Validation() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := strings.TrimSpace(c.GetHeader("X-Request-ID"))
		if len(requestID) > 128 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "X-Request-ID length must be <= 128"})
			return
		}
		c.Next()
	}
}
