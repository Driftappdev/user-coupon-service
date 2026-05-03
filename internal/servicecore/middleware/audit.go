package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Audit() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		log.Printf("audit_log request_id=%s method=%s path=%s status=%d latency_ms=%d",
			c.GetString("request_id"),
			c.Request.Method,
			c.FullPath(),
			c.Writer.Status(),
			time.Since(start).Milliseconds(),
		)
	}
}
