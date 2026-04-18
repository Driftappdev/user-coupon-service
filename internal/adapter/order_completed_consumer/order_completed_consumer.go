package adapter

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/proto"

	dto "dift_user_insentive/user-coupon-service/internal/dto/order_event"
	event_consumer "dift_user_insentive/user-coupon-service/internal/interface/event_consumer"
	pb "dift_user_insentive/user-coupon-service/proto/pb/order_completed"
)

type OrderCompletedConsumer struct {
	handler event_consumer.OrderCompletedConsumer
}

func NewOrderCompletedConsumer(
	handler event_consumer.OrderCompletedConsumer,
) *OrderCompletedConsumer {
	return &OrderCompletedConsumer{
		handler: handler,
	}
}

func (c *OrderCompletedConsumer) Handle(
	ctx context.Context,
	subject string,
	payload []byte,
) error {

	pbEvent := &pb.OrderCompletedEvent{}

	if err := proto.Unmarshal(payload, pbEvent); err != nil {
		return fmt.Errorf("failed to unmarshal proto: %w", err)
	}

	dtoEvent := dto.FromOrderCompletedProto(pbEvent)

	return c.handler.Handle(ctx, dtoEvent)
}

