package core

type CouponStatus string

const (
	CouponStatusAvailable CouponStatus = "AVAILABLE"
	CouponStatusReserved  CouponStatus = "RESERVED"
	CouponStatusUsed      CouponStatus = "USED"
	CouponStatusExpired   CouponStatus = "EXPIRED"
)

