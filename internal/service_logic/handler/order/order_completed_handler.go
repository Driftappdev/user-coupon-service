package handler

import (
	"context"
	"errors"
	"fmt"

	dto "dift_user_insentive/user-coupon-service/internal/dto/order_event"
	event_consumer "dift_user_insentive/user-coupon-service/internal/interface/event_consumer"
	serviceport "dift_user_insentive/user-coupon-service/internal/interface/service/order_completed"
)

type OrderCompletedHandler struct {
	service serviceport.OrderCompletedService
}

// âœ… compile-time check
var _ event_consumer.OrderCompletedConsumer = (*OrderCompletedHandler)(nil)

func NewOrderCompletedHandler(
	svc serviceport.OrderCompletedService,
) *OrderCompletedHandler {
	return &OrderCompletedHandler{
		service: svc,
	}
}

func (h *OrderCompletedHandler) Handle(
	ctx context.Context,
	evt *dto.OrderCompletedEvent,
) error {

	if evt == nil {
		return errors.New("event is nil")
	}

	if evt.UserID == "" {
		return errors.New("missing user_id")
	}

	if evt.CouponCode == "" {
		return errors.New("missing coupon_code")
	}

	if evt.OrderID == "" {
		return errors.New("missing order_id")
	}

	if err := h.service.UseCoupon(
		ctx,
		evt.UserID,
		evt.CouponCode,
		evt.OrderID,
	); err != nil {
		return fmt.Errorf("use coupon failed: %w", err)
	}

	return nil
}

