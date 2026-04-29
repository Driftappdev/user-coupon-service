package natsadapter

import (
	"context"

	dto "dift_user_insentive/user-coupon-service/internal/dto/coupon_event"
	port "dift_user_insentive/user-coupon-service/internal/interface/event_consumer"
	couponpb "dift_user_insentive/user-coupon-service/proto/pb/coupon"

	"google.golang.org/protobuf/proto"
)

type Consumer struct {
	handler port.CouponEventConsumer
}

func New(handler port.CouponEventConsumer) *Consumer {
	return &Consumer{handler: handler}
}

func (c *Consumer) Handle(
	ctx context.Context,
	subject string,
	payload []byte,
) error {

	var evt couponpb.CouponEvent

	if err := proto.Unmarshal(payload, &evt); err != nil {
		return err
	}

	// âœ… à¹à¸›à¸¥à¸‡à¸•à¸£à¸‡à¸™à¸µà¹‰
	dtoEvent := dto.FromProto(&evt)

	return c.handler.Handle(ctx, dtoEvent)
}

