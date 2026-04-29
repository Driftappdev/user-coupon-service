package adapter

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/proto"

	dto "dift_user_insentive/user-coupon-service/internal/dto/order_event"
	eventport "dift_user_insentive/user-coupon-service/internal/interface/event_consumer"
	pb "dift_user_insentive/user-coupon-service/proto/pb/order_failed"
)

type OrderFailedConsumer struct {
	handler eventport.OrderFailedConsumer
}

func NewOrderFailedConsumer(
	handler eventport.OrderFailedConsumer,
) *OrderFailedConsumer {
	return &OrderFailedConsumer{
		handler: handler,
	}
}

func (c *OrderFailedConsumer) Handle(
	ctx context.Context,
	subject string,
	payload []byte,
) error {

	pbEvent := &pb.OrderFailedEvent{}

	if err := proto.Unmarshal(payload, pbEvent); err != nil {
		return fmt.Errorf("failed to unmarshal order failed event: %w", err)
	}

	// optional guard
	if pbEvent.UserId == "" || pbEvent.CouponCode == "" {
		return nil
	}

	dtoEvent := dto.FromOrderFailedProto(pbEvent)

	return c.handler.Handle(ctx, dtoEvent)
}

