package expire

import (
	"context"
	"time"

	repo "dift_user_insentive/user-coupon-service/internal/interface/repository"
)

type CouponExpireService struct {
	expirer repo.CouponExpirationRepository
}

func NewCouponExpireService(
	expirer repo.CouponExpirationRepository,
) *CouponExpireService {
	return &CouponExpireService{
		expirer: expirer,
	}
}

func (s *CouponExpireService) Expire(ctx context.Context) error {
	return s.expirer.ExpireBefore(ctx, time.Now())
}

