package reserve

import (
	"context"
	"time"

	port "dift_user_insentive/user-coupon-service/internal/interface/repository"
	serviceport "dift_user_insentive/user-coupon-service/internal/interface/service/coupon_reserve"
	shared "dift_user_insentive/user-coupon-service/internal/service_logic/service/core"
)

type CouponReserveService struct {
	repo port.UserCouponRepository
}

var _ serviceport.CouponReserveService = (*CouponReserveService)(nil)

func NewCouponReserveService(
	r port.UserCouponRepository,
) *CouponReserveService {
	return &CouponReserveService{
		repo: r,
	}
}

func (s *CouponReserveService) ReserveCoupon(
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

	if !c.Reserve(time.Now()) {
		return shared.ErrCouponNotAvailable
	}

	return s.repo.Upsert(ctx, c)
}

