package middleware

import "github.com/gin-gonic/gin"

func Idempotency() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

