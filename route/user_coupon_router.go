package route

import (
	httpadapter "dift_user_insentive/user-coupon-service/internal/adapter/http"
	servicecore "dift_user_insentive/user-coupon-service/internal/servicecore"
	adminshield "github.com/PlatformCore/middleware/adminshield"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RegisterRoutes(
	r *gin.Engine,
	queryHandler *httpadapter.Handler,
) {
	r.GET("/health", func(c *gin.Context) {
		servicecore.HealthController("user-coupon-service")(c)
	})
	r.GET("/metrics/app", gin.WrapH(promhttp.Handler()))

	group := r.Group("/api/v1/user-coupons")

	group.GET("", func(c *gin.Context) {
		queryHandler.GetUserCoupons(c.Writer, c.Request)
	})

	group.GET("/health", func(c *gin.Context) {
		servicecore.HealthController("user-coupon-service")(c)
	})

	admin := r.Group("/api/v1/admin", adminshield.GinRequireRoles("SUPER_ADMIN", "ADMIN", "EDITOR", "VIEWER"))
	admin.GET("/user-coupons", func(c *gin.Context) {
		queryHandler.GetUserCoupons(c.Writer, c.Request)
	})
	admin.GET("/health", func(c *gin.Context) {
		servicecore.HealthController("user-coupon-service")(c)
	})

	r.POST("/internal/admin/control", func(c *gin.Context) {
		secret := strings.TrimSpace(os.Getenv("ADMIN_CONTROL_SHARED_SECRET"))
		if secret != "" && c.GetHeader("X-Admin-Secret") != secret {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		var req map[string]any
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_json"})
			return
		}
		action, _ := req["action"].(string)
		if strings.TrimSpace(action) == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "action_required"})
			return
		}
		c.JSON(http.StatusAccepted, gin.H{"accepted": true, "service": "user-coupon-service", "action": action})
	})
}
