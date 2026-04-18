package core

import "errors"

var (
	ErrCouponNotFound     = errors.New("coupon not found")
	ErrCouponNotReserved  = errors.New("coupon not reserved")
	ErrCouponNotAvailable = errors.New("coupon not available")
)

