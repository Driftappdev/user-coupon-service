package admin

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	adminport "dift_user_insentive/user-coupon-service/internal/interface/service/admin"
	assignport "dift_user_insentive/user-coupon-service/internal/interface/service/coupon_event"
	releaseport "dift_user_insentive/user-coupon-service/internal/interface/service/coupon_release"
	reserveport "dift_user_insentive/user-coupon-service/internal/interface/service/coupon_reserve"
	"dift_user_insentive/user-coupon-service/internal/model"
	expiresvc "dift_user_insentive/user-coupon-service/internal/service_logic/service/worker"
)

var (
	ErrUnknownAction  = errors.New("unknown admin action")
	ErrInvalidPayload = errors.New("invalid admin payload")
)

type CommandService struct {
	assignService  assignport.CouponCommandService
	reserveService reserveport.CouponReserveService
	releaseService releaseport.CouponReleaseService
	expireService  *expiresvc.CouponExpireService
}

func NewCommandService(
	assignService assignport.CouponCommandService,
	reserveService reserveport.CouponReserveService,
	releaseService releaseport.CouponReleaseService,
	expireService *expiresvc.CouponExpireService,
) *CommandService {
	return &CommandService{
		assignService:  assignService,
		reserveService: reserveService,
		releaseService: releaseService,
		expireService:  expireService,
	}
}

var _ adminport.CommandService = (*CommandService)(nil)

func (s *CommandService) Execute(ctx context.Context, cmd adminport.Command) error {
	switch strings.ToLower(strings.TrimSpace(cmd.Action)) {
	case "coupon.assign":
		userID := payloadString(cmd.Payload, "user_id")
		code := payloadString(cmd.Payload, "coupon_code")
		if userID == "" || code == "" {
			return ErrInvalidPayload
		}
		now := time.Now().UTC()
		validFrom := payloadTime(cmd.Payload, "valid_from", now.Add(-time.Minute))
		validTo := payloadTime(cmd.Payload, "valid_to", now.Add(30*24*time.Hour))
		return s.assignService.AssignCoupon(ctx, &model.UserCoupon{
			ID:            payloadString(cmd.Payload, "id"),
			UserID:        userID,
			CouponCode:    code,
			CampaignID:    payloadString(cmd.Payload, "campaign_id"),
			Status:        model.StatusAssigned,
			DiscountType:  payloadString(cmd.Payload, "discount_type"),
			DiscountValue: payloadFloat64(cmd.Payload, "discount_value"),
			MinOrder:      payloadFloat64(cmd.Payload, "min_order"),
			MaxDiscount:   payloadFloat64(cmd.Payload, "max_discount"),
			ProductType:   payloadString(cmd.Payload, "product_type"),
			ValidFrom:     validFrom,
			ValidTo:       validTo,
			MaxUsage:      int(payloadInt64Default(cmd.Payload, "max_usage", 1)),
			IssuedBy:      "admin-command",
			CreatedAt:     now,
			UpdatedAt:     now,
		})
	case "coupon.reserve":
		userID := payloadString(cmd.Payload, "user_id")
		code := payloadString(cmd.Payload, "coupon_code")
		if userID == "" || code == "" {
			return ErrInvalidPayload
		}
		return s.reserveService.ReserveCoupon(ctx, userID, code)
	case "coupon.release":
		userID := payloadString(cmd.Payload, "user_id")
		code := payloadString(cmd.Payload, "coupon_code")
		if userID == "" || code == "" {
			return ErrInvalidPayload
		}
		return s.releaseService.ReleaseCoupon(ctx, userID, code)
	case "coupon.expire.run":
		return s.expireService.Expire(ctx)
	default:
		return fmt.Errorf("%w: %s", ErrUnknownAction, cmd.Action)
	}
}

func payloadString(payload map[string]any, key string) string {
	v, ok := payload[key]
	if !ok || v == nil {
		return ""
	}
	s, ok := v.(string)
	if !ok {
		return ""
	}
	return strings.TrimSpace(s)
}

func payloadFloat64(payload map[string]any, key string) float64 {
	v, ok := payload[key]
	if !ok || v == nil {
		return 0
	}
	switch n := v.(type) {
	case float64:
		return n
	case float32:
		return float64(n)
	case int:
		return float64(n)
	case int32:
		return float64(n)
	case int64:
		return float64(n)
	default:
		return 0
	}
}

func payloadInt64Default(payload map[string]any, key string, fallback int64) int64 {
	v, ok := payload[key]
	if !ok || v == nil {
		return fallback
	}
	switch n := v.(type) {
	case int:
		return int64(n)
	case int32:
		return int64(n)
	case int64:
		return n
	case float32:
		return int64(n)
	case float64:
		return int64(n)
	default:
		return fallback
	}
}

func payloadTime(payload map[string]any, key string, fallback time.Time) time.Time {
	raw := payloadString(payload, key)
	if raw == "" {
		return fallback
	}
	t, err := time.Parse(time.RFC3339, raw)
	if err != nil {
		return fallback
	}
	return t.UTC()
}
