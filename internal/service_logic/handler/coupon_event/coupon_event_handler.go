package handler_logic

import (
	"context"
	"errors"
	"fmt"
	"time"

	dto "dift_user_insentive/user-coupon-service/internal/dto/coupon_event"
	eventport "dift_user_insentive/user-coupon-service/internal/interface/event_consumer"
	serviceport "dift_user_insentive/user-coupon-service/internal/interface/service/coupon_event"
	"dift_user_insentive/user-coupon-service/internal/model"
)

type CouponEventHandler struct {
	command serviceport.CouponCommandService
}

// âœ… compile-time check
var _ eventport.CouponEventConsumer = (*CouponEventHandler)(nil)

func NewCouponEventHandler(
	cmd serviceport.CouponCommandService,
) *CouponEventHandler {
	return &CouponEventHandler{
		command: cmd,
	}
}

func (h *CouponEventHandler) Handle(
	ctx context.Context,
	evt *dto.CouponEvent,
) error {

	if evt == nil {
		return errors.New("event is nil")
	}

	switch evt.Type {

	case "CREATED", "UPDATED":
		return h.handleUpsert(ctx, evt)

	case "REDEEMED":
		// TODO: implement redeem logic
		return nil

	case "DEACTIVATED":
		// TODO: implement deactivate logic
		return nil

	default:
		// ðŸ‘‰ à¸à¸±à¸™ event à¹à¸›à¸¥à¸ à¹†
		return fmt.Errorf("unsupported event type: %s", evt.Type)
	}
}

func (h *CouponEventHandler) handleUpsert(
	ctx context.Context,
	evt *dto.CouponEvent,
) error {

	vf, err := time.Parse(time.RFC3339, evt.ValidFrom)
	if err != nil {
		return fmt.Errorf("invalid valid_from: %w", err)
	}

	vt, err := time.Parse(time.RFC3339, evt.ValidTo)
	if err != nil {
		return fmt.Errorf("invalid valid_to: %w", err)
	}

	uc := &model.UserCoupon{
		UserID:        evt.UserID,
		CouponCode:    evt.CouponCode,
		Status:        model.StatusAssigned,
		DiscountType:  evt.DiscountType,
		DiscountValue: evt.DiscountValue,
		MinOrder:      evt.MinOrder,
		MaxDiscount:   evt.MaxDiscount,
		ProductType:   evt.ProductType,
		ValidFrom:     vf,
		ValidTo:       vt,
		MaxUsage:      evt.MaxUsage,
	}

	return h.command.AssignCoupon(ctx, uc)
}

