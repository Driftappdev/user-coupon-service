package middleware

import (
	"net/http"

	"dift_user_insentive/user-coupon-service/pkg/wrapper"

	"github.com/gin-gonic/gin"
)

func Validation() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := wrapper.CleanString(c.GetHeader("X-Request-ID"))
		if requestID != "" {
			if err := wrapper.MaxLen("X-Request-ID", requestID, 128); err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		}
		c.Next()
	}
}

