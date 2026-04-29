package http

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"dift_user_insentive/user-coupon-service/pkg/wrapper"

	"github.com/gin-gonic/gin"
)

type MiddlewareSet struct {
	Logger    *wrapper.Logger
	Limiter   wrapper.RateLimiter
	Extractor wrapper.KeyExtractor
	Tracer    *wrapper.Tracer
}

func NewMiddlewareSet(serviceName string) MiddlewareSet {
	limiter, extractor := wrapper.NewIPSlidingWindow(200, time.Minute)
	return MiddlewareSet{
		Logger:    wrapper.NewJSONLogger(wrapper.LevelInfo),
		Limiter:   limiter,
		Extractor: extractor,
		Tracer:    wrapper.GlobalTracer(serviceName),
	}
}

func AccessLog(set MiddlewareSet) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		ctx, span := wrapper.StartTraceSpan(set.Tracer, c.Request.Context(), "http.request")
		c.Request = c.Request.WithContext(ctx)
		c.Next()
		status := c.Writer.Status()
		set.Logger.Info("http_request",
			wrapper.LogField("method", c.Request.Method),
			wrapper.LogField("path", c.FullPath()),
			wrapper.LogField("status", status),
			wrapper.LogField("latency_ms", time.Since(start).Milliseconds()),
		)
		wrapper.TraceSetAttrs(span,
			wrapper.TraceAttr("http.method", c.Request.Method),
			wrapper.TraceAttr("http.path", c.FullPath()),
			wrapper.TraceAttr("http.status_code", status),
		)
		wrapper.TraceEnd(span)
	}
}

func Recover(set MiddlewareSet) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if rec := recover(); rec != nil {
				set.Logger.Error("panic_recovered", wrapper.LogField("panic", rec))
				err := wrapper.NewBaseError("INTERNAL_PANIC", http.StatusInternalServerError, "internal server error")
				c.AbortWithStatusJSON(wrapper.HTTPStatus(err), wrapper.APIResponse(err))
			}
		}()
		c.Next()
	}
}

func RateLimit(set MiddlewareSet) gin.HandlerFunc {
	retryCfg := wrapper.DefaultRetryConfig(2, 20*time.Millisecond, 100*time.Millisecond)
	return func(c *gin.Context) {
		key := wrapper.LimitKey(set.Extractor, c.Request)
		var decision wrapper.RateLimitResult
		err := wrapper.RetryDo(c.Request.Context(), retryCfg, func(ctx context.Context) error {
			return wrapper.DoWithTimeout(ctx, 250*time.Millisecond, func(inner context.Context) error {
				res, allowErr := wrapper.Allow(set.Limiter, inner, key)
				if allowErr == nil {
					decision = res
				}
				return allowErr
			})
		})
		if err != nil {
			set.Logger.Error("rate_limit_error", wrapper.LogField("error", err.Error()))
			c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{"error": "rate limiter unavailable"})
			return
		}
		if !decision.Allowed {
			retryAfter := strconv.Itoa(int(decision.RetryAfter.Seconds()))
			if retryAfter == "0" {
				retryAfter = "1"
			}
			c.Header("Retry-After", retryAfter)
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "rate limit exceeded"})
			return
		}
		c.Next()
	}
}

