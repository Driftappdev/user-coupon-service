package http

import (
	"context"
	"dift_user_insentive/user-coupon-service/internal/model"
)

type CouponQueryPort interface {
	ListUserCoupons(ctx context.Context, userID string) ([]*model.UserCoupon, error)
}

