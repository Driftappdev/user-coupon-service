package grpc

import (
	"context"

	"dift_user_insentive/user-coupon-service/internal/model"
)

type CouponQueryPort interface {
	ListUserCoupons(
		ctx context.Context,
		userID string,
	) ([]*model.UserCoupon, error)

	ValidateCoupon(
		ctx context.Context,
		userID string,
		code string,
	) (bool, string)
}

