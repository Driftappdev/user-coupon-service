package dto

import pb "dift_user_insentive/user-coupon-service/proto/pb/order_completed"

func FromOrderCompletedProto(evt *pb.OrderCompletedEvent) *OrderCompletedEvent {
	return &OrderCompletedEvent{
		UserID:     evt.UserId,
		CouponCode: evt.CouponCode,
		OrderID:    evt.OrderId,
	}
}

