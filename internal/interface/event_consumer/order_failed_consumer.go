package event_consumer

import (
	"context"

	dto "dift_user_insentive/user-coupon-service/internal/dto/order_event"
)

type OrderFailedConsumer interface {
	Handle(ctx context.Context, evt *dto.OrderFailedEvent) error
}

