package event_consumer

import (
	"context"
	dto "dift_user_insentive/user-coupon-service/internal/dto/coupon_event"
)

type CouponEventConsumer interface {
	Handle(
		ctx context.Context,
		event *dto.CouponEvent,
	) error
}

