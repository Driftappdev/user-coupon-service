package model

import "time"

type Status string

const (
	StatusAssigned Status = "assigned" // = AVAILABLE
	StatusReserved Status = "reserved" // ðŸ”¥ à¹€à¸žà¸´à¹ˆà¸¡
	StatusRedeemed Status = "redeemed"
	StatusExpired  Status = "expired"
	StatusDisabled Status = "disabled"
)

type UserCoupon struct {
	ID string

	// ownership
	UserID     string
	CouponCode string

	// campaign
	CampaignID string

	// coupon state
	Status Status

	// discount
	DiscountType  string
	DiscountValue float64

	// usage conditions
	MinOrder    float64
	MaxDiscount float64
	ProductType string

	// validity period
	ValidFrom time.Time
	ValidTo   time.Time

	// usage control
	MaxUsage  int
	UsedCount int

	// source
	IssuedBy string

	// usage tracking
	LastUsedAt *time.Time

	// optimistic locking
	Version int

	// audit
	CreatedAt time.Time
	UpdatedAt time.Time
}

// ==============================
// ðŸ” Validation
// ==============================

func (c *UserCoupon) IsExpired(now time.Time) bool {
	if c.ValidTo.IsZero() {
		return false
	}
	return now.After(c.ValidTo)
}

// à¹ƒà¸Šà¹‰à¸ªà¸³à¸«à¸£à¸±à¸š "à¸à¹ˆà¸­à¸™ reserve"
func (c *UserCoupon) CanUse(now time.Time) bool {

	if c.Status != StatusAssigned {
		return false
	}

	if now.Before(c.ValidFrom) {
		return false
	}

	if c.IsExpired(now) {
		return false
	}

	if c.MaxUsage > 0 && c.UsedCount >= c.MaxUsage {
		return false
	}

	return true
}

// ==============================
// ðŸ”’ Reservation Flow
// ==============================

// Reserve = à¸à¸±à¸™ coupon
func (c *UserCoupon) Reserve(now time.Time) bool {

	if !c.CanUse(now) {
		return false
	}

	c.Status = StatusReserved
	c.UpdatedAt = now

	return true
}

// Release = à¸„à¸·à¸™ coupon (order failed)
func (c *UserCoupon) Release(now time.Time) bool {

	if c.Status != StatusReserved {
		return false
	}

	c.Status = StatusAssigned
	c.UpdatedAt = now

	return true
}

// Redeem = à¹ƒà¸Šà¹‰ coupon à¸ˆà¸£à¸´à¸‡ (order success)
func (c *UserCoupon) Redeem(now time.Time) bool {

	// ðŸ”¥ à¸•à¹‰à¸­à¸‡ reserved à¸à¹ˆà¸­à¸™à¹€à¸—à¹ˆà¸²à¸™à¸±à¹‰à¸™
	if c.Status != StatusReserved {
		return false
	}

	// validate à¸‹à¹‰à¸³à¸à¸±à¸™ edge case (optional à¹à¸•à¹ˆà¹à¸™à¸°à¸™à¸³)
	if c.IsExpired(now) {
		return false
	}

	c.UsedCount++
	c.LastUsedAt = &now
	c.UpdatedAt = now

	// à¸–à¹‰à¸²à¹ƒà¸Šà¹‰à¸„à¸£à¸š limit â†’ à¸›à¸´à¸” coupon
	if c.MaxUsage > 0 && c.UsedCount >= c.MaxUsage {
		c.Status = StatusRedeemed
	}

	return true
}

