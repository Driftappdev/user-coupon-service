package service

import (
	"context"
	"dift_user_insentive/user-coupon-service/internal/model"
)

type CouponCommandService interface {
	AssignCoupon(ctx context.Context, uc *model.UserCoupon) error
}

