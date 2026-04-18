package servicecore

import (
	middleware "dift_user_insentive/user-coupon-service/internal/servicecore/middleware"
	"github.com/gin-gonic/gin"
)

func UseDefaultMiddlewares(r *gin.Engine) {
	r.Use(
		middleware.Context(),
		middleware.Audit(),
		middleware.Permission(),
		middleware.Idempotency(),
		middleware.Validation(),
	)
}

