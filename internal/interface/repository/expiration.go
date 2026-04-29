package repository

import (
	"context"
	"time"
)

//
// ðŸ”¹ Coupon Expiration Repository
//

type CouponExpirationRepository interface {
	ExpireBefore(ctx context.Context, t time.Time) error
}

