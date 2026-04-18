package servicecore

import "github.com/gin-gonic/gin"

func HealthController(serviceName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"service": serviceName, "status": "ok"})
	}
}

