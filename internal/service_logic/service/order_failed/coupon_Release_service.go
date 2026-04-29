package release

import (
	"context"
	"time"

	port "dift_user_insentive/user-coupon-service/internal/interface/repository"
	serviceport "dift_user_insentive/user-coupon-service/internal/interface/service/coupon_release"
	shared "dift_user_insentive/user-coupon-service/internal/service_logic/service/core"
)

type CouponReleaseService struct {
	repo port.UserCouponRepository
}

var _ serviceport.CouponReleaseService = (*CouponReleaseService)(nil)

func NewCouponReleaseService(
	r port.UserCouponRepository,
) *CouponReleaseService {
	return &CouponReleaseService{
		repo: r,
	}
}

func (s *CouponReleaseService) ReleaseCoupon(
	ctx context.Context,
	userID string,
	code string,
) error {

	c, err := s.repo.FindOne(ctx, userID, code)
	if err != nil {
		return err
	}
	if c == nil {
		return shared.ErrCouponNotFound
	}

	if !c.Release(time.Now()) {
		return nil
	}

	return s.repo.Upsert(ctx, c)
}

