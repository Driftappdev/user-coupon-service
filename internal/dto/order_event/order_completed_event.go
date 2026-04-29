package dto

type OrderCompletedEvent struct {
	UserID     string
	CouponCode string
	OrderID    string
}

