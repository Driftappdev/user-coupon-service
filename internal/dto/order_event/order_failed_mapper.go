package dto

import pb "dift_user_insentive/user-coupon-service/proto/pb/order_failed"

func FromOrderFailedProto(evt *pb.OrderFailedEvent) *OrderFailedEvent {
	return &OrderFailedEvent{
		UserID:     evt.UserId,
		CouponCode: evt.CouponCode,
		OrderID:    evt.OrderId,
	}
}

