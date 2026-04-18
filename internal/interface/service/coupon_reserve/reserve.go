package coupon_reserve

import "context"

type CouponReserveService interface {
	ReserveCoupon(ctx context.Context, userID, code string) error
}

