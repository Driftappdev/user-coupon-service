package model

import "time"

type OrderCouponStatus string

const (
	OrderCouponApplied  OrderCouponStatus = "applied"
	OrderCouponReverted OrderCouponStatus = "reverted"
)

type OrderCoupon struct {
	ID string

	OrderID string
	UserID  string

	UserCouponID string
	CouponCode   string

	DiscountAmount float64

	Status OrderCouponStatus

	RedeemedAt time.Time
	CreatedAt  time.Time
}

