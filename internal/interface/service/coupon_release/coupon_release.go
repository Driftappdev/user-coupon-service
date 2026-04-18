package coupon_release

import "context"

type CouponReleaseService interface {
	ReleaseCoupon(ctx context.Context, userID, code string) error
}

