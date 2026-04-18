package dto

import (
	"dift_user_insentive/user-coupon-service/internal/model"
	usercouponpb "dift_user_insentive/user-coupon-service/proto/pb/usercoupon"
	"time"
)

func ToProtoList(list []*model.UserCoupon) []*usercouponpb.UserCoupon {
	var res []*usercouponpb.UserCoupon

	for _, c := range list {
		res = append(res, ToProto(c))
	}

	return res
}

func ToProto(c *model.UserCoupon) *usercouponpb.UserCoupon {
	if c == nil {
		return nil
	}

	return &usercouponpb.UserCoupon{
		UserId:        c.UserID,
		CouponCode:    c.CouponCode,
		Status:        string(c.Status),
		DiscountType:  c.DiscountType,
		DiscountValue: c.DiscountValue,
		MinOrder:      c.MinOrder,
		MaxDiscount:   c.MaxDiscount,
		ProductType:   c.ProductType,
		MaxUsage:      int32(c.MaxUsage),
		UsageCount:    int32(c.UsedCount),
		ValidFrom:     formatTime(c.ValidFrom),
		ValidTo:       formatTime(c.ValidTo),
	}
}

func formatTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format(time.RFC3339)
}

