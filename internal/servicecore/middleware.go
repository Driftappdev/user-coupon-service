package servicecore

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HTTPMiddleware defines standard servicecore HTTP middleware signature.
type HTTPMiddleware func(http.Handler) http.Handler

// DefaultHTTPMiddlewares returns enterprise baseline middleware chain placeholders.
func DefaultHTTPMiddlewares() []HTTPMiddleware {
	return []HTTPMiddleware{}
}

// UseDefaultMiddlewares keeps compatibility with services wiring Gin directly.
func UseDefaultMiddlewares(r *gin.Engine) {
	_ = r
}
