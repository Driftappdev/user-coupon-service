package adminmiddleware

import (
	adminshield "github.com/driftappdev/libpackage/filemods/middleware/adminshield/admin-middleware"
	"github.com/gin-gonic/gin"
)

func GinRequireRoles(roles ...string) gin.HandlerFunc {
	return adminshield.GinRequireRoles(roles...)
}
