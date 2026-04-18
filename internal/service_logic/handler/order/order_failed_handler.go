package handler

import (
	"context"

	dto "dift_user_insentive/user-coupon-service/internal/dto/order_event"
	eventport "dift_user_insentive/user-coupon-service/internal/interface/event_consumer"
	releaseport "dift_user_insentive/user-coupon-service/internal/interface/service/coupon_release"
)

type OrderFailedHandler struct {
	service releaseport.CouponReleaseService
}

// âœ… ensure implement interface
var _ eventport.OrderFailedConsumer = (*OrderFailedHandler)(nil)

func NewOrderFailedHandler(
	svc releaseport.CouponReleaseService,
) *OrderFailedHandler {
	return &OrderFailedHandler{
		service: svc,
	}
}

func (h *OrderFailedHandler) Handle(
	ctx context.Context,
	evt *dto.OrderFailedEvent,
) error {

	if evt == nil {
		return nil
	}

	if evt.UserID == "" || evt.CouponCode == "" {
		return nil
	}

	return h.service.ReleaseCoupon(
		ctx,
		evt.UserID,
		evt.CouponCode,
	)
}

