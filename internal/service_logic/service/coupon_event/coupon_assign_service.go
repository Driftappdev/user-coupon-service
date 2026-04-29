package assign

import (
	"context"
	"errors"

	repo "dift_user_insentive/user-coupon-service/internal/interface/repository"
	serviceport "dift_user_insentive/user-coupon-service/internal/interface/service/coupon_event"
	"dift_user_insentive/user-coupon-service/internal/model"
)

type CouponAssignService struct {
	writer repo.UserCouponWriter
}

var _ serviceport.CouponCommandService = (*CouponAssignService)(nil)

func NewCouponAssignService(
	writer repo.UserCouponWriter,
) *CouponAssignService {
	return &CouponAssignService{
		writer: writer,
	}
}

func (s *CouponAssignService) AssignCoupon(
	ctx context.Context,
	uc *model.UserCoupon,
) error {

	if uc == nil {
		return errors.New("user coupon is nil")
	}

	return s.writer.Upsert(ctx, uc)
}

