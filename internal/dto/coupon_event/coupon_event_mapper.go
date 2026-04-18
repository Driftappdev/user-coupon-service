package dto

import couponpb "dift_user_insentive/user-coupon-service/proto/pb/coupon"

func FromProto(evt *couponpb.CouponEvent) *CouponEvent {
	return &CouponEvent{
		Type:          evt.Type.String(),
		UserID:        evt.UserId,
		CouponCode:    evt.CouponCode,
		DiscountType:  evt.DiscountType,
		DiscountValue: evt.DiscountValue,
		MinOrder:      evt.MinOrder,
		MaxDiscount:   evt.MaxDiscount,
		ProductType:   evt.ProductType,
		ValidFrom:     evt.ValidFrom,
		ValidTo:       evt.ValidTo,
		MaxUsage:      int(evt.MaxUsage),
	}
}

