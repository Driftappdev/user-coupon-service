package query

import (
	"context"
	"time"

	repo "dift_user_insentive/user-coupon-service/internal/interface/repository"
	"dift_user_insentive/user-coupon-service/internal/model"
)

type CouponQueryService struct {
	repository repo.UserCouponRepository
}

func NewCouponQueryService(
	r repo.UserCouponRepository,
) *CouponQueryService {
	return &CouponQueryService{
		repository: r,
	}
}

func (s *CouponQueryService) ListUserCoupons(
	ctx context.Context,
	userID string,
) ([]*model.UserCoupon, error) {

	return s.repository.FindByUser(ctx, userID)
}

func (s *CouponQueryService) ValidateCoupon(
	ctx context.Context,
	userID string,
	code string,
) (bool, string) {

	c, err := s.repository.FindOne(ctx, userID, code)
	if err != nil || c == nil {
		return false, "coupon not found"
	}

	if !c.CanUse(time.Now()) {
		return false, "coupon not usable"
	}

	return true, ""
}

