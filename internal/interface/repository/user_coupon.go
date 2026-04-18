package repository

import (
	"context"

	"dift_user_insentive/user-coupon-service/internal/model"
)

//
// ðŸ”¹ User Coupon Reader
//

type UserCouponReader interface {
	FindByUser(ctx context.Context, userID string) ([]*model.UserCoupon, error)
	FindOne(ctx context.Context, userID, code string) (*model.UserCoupon, error)
}

//
// ðŸ”¹ User Coupon Writer
//

type UserCouponWriter interface {
	Upsert(ctx context.Context, coupon *model.UserCoupon) error
	Delete(ctx context.Context, userID, code string) error
}

//
// ðŸ”¹ Combined Repository
//

type UserCouponRepository interface {
	UserCouponReader
	UserCouponWriter
}

