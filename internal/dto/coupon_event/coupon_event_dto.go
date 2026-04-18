package dto

type CouponEvent struct {
	Type          string
	UserID        string
	CouponCode    string
	DiscountType  string
	DiscountValue float64
	MinOrder      float64
	MaxDiscount   float64
	ProductType   string
	ValidFrom     string
	ValidTo       string
	MaxUsage      int
}

