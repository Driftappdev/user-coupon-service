package dto

type OrderFailedEvent struct {
	UserID     string
	CouponCode string
	OrderID    string
}

