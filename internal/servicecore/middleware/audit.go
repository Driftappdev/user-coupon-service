package middleware

import (
	"time"

	"dift_user_insentive/user-coupon-service/pkg/wrapper"

	"github.com/gin-gonic/gin"
)

func Audit() gin.HandlerFunc {
	logger := wrapper.NewJSONLogger(wrapper.LevelInfo)
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		logger.Info("audit_log",
			wrapper.LogField("request_id", c.GetString("request_id")),
			wrapper.LogField("method", c.Request.Method),
			wrapper.LogField("path", c.FullPath()),
			wrapper.LogField("status", c.Writer.Status()),
			wrapper.LogField("latency_ms", time.Since(start).Milliseconds()),
		)
	}
}

