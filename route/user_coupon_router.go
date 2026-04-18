package route

import (
	httpadapter "dift_user_insentive/user-coupon-service/internal/adapter/http"
	servicecore "dift_user_insentive/user-coupon-service/internal/servicecore"
	wrapper "dift_user_insentive/user-coupon-service/pkg/wrapper"
	adminmiddleware "dift_user_insentive/user-coupon-service/pkg/wrapper/adminmiddleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	r *gin.Engine,
	queryHandler *httpadapter.Handler,
) {
	r.GET("/health", func(c *gin.Context) {
		servicecore.HealthController("user-coupon-service")(c)
	})
	r.GET("/metrics/app", gin.WrapF(wrapper.MetricsHandler()))

	group := r.Group("/api/v1/user-coupons")

	group.GET("", func(c *gin.Context) {
		queryHandler.GetUserCoupons(c.Writer, c.Request)
	})

	group.GET("/health", func(c *gin.Context) {
		servicecore.HealthController("user-coupon-service")(c)
	})

	admin := r.Group("/api/v1/admin", adminmiddleware.GinRequireRoles("SUPER_ADMIN", "ADMIN", "EDITOR", "VIEWER"))
	admin.GET("/user-coupons", func(c *gin.Context) {
		queryHandler.GetUserCoupons(c.Writer, c.Request)
	})
	admin.GET("/health", func(c *gin.Context) {
		servicecore.HealthController("user-coupon-service")(c)
	})
}
