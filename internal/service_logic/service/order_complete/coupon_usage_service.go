package usage

import (
	"context"
	"time"

	repo "dift_user_insentive/user-coupon-service/internal/interface/repository"
	serviceport "dift_user_insentive/user-coupon-service/internal/interface/service/order_completed"
	"dift_user_insentive/user-coupon-service/internal/model"
	shared "dift_user_insentive/user-coupon-service/internal/service_logic/service/core"
)

type CouponUsageService struct {
	userCouponRepo repo.UserCouponRepository
	idemRepo       repo.IdempotencyRepository
	usageRepo      repo.UsageRepository
}

// âœ… ensure implement interface
var _ serviceport.OrderCompletedService = (*CouponUsageService)(nil)

func NewCouponUsageService(
	userCouponRepo repo.UserCouponRepository,
	idemRepo repo.IdempotencyRepository,
	usageRepo repo.UsageRepository,
) *CouponUsageService {
	return &CouponUsageService{
		userCouponRepo: userCouponRepo,
		idemRepo:       idemRepo,
		usageRepo:      usageRepo,
	}
}

func (s *CouponUsageService) UseCoupon(
	ctx context.Context,
	userID string,
	code string,
	orderID string,
) error {

	// âœ… idempotency check
	exists, err := s.idemRepo.Exists(ctx, orderID)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}

	// âœ… find coupon
	coupon, err := s.userCouponRepo.FindOne(ctx, userID, code)
	if err != nil {
		return err
	}
	if coupon == nil {
		return shared.ErrCouponNotFound
	}

	// âœ… redeem
	if !coupon.Redeem(time.Now()) {
		return shared.ErrCouponNotReserved
	}

	// âœ… update coupon
	if err := s.userCouponRepo.Upsert(ctx, coupon); err != nil {
		return err
	}

	// âœ… save usage log
	if err := s.usageRepo.Create(
		ctx,
		userID,
		code,
		orderID,
		string(model.StatusRedeemed),
	); err != nil {
		return err
	}

	// âœ… save idempotency key
	return s.idemRepo.Save(ctx, orderID)
}

