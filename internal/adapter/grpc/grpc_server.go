package grpc

import (
	"context"

	dto "dift_user_insentive/user-coupon-service/internal/dto/order_grpc"
	port "dift_user_insentive/user-coupon-service/internal/interface/grpc"
	reserveport "dift_user_insentive/user-coupon-service/internal/interface/service/coupon_reserve"
	usercouponpb "dift_user_insentive/user-coupon-service/proto/pb/usercoupon"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	usercouponpb.UnimplementedUserCouponServiceServer

	query      port.CouponQueryPort
	reserveSvc reserveport.CouponReserveService
}

func New(
	query port.CouponQueryPort,
	reserveSvc reserveport.CouponReserveService,
) *Handler {
	return &Handler{
		query:      query,
		reserveSvc: reserveSvc,
	}
}

// -------------------------
// GET USER COUPONS
// -------------------------

func (h *Handler) GetUserCoupons(
	ctx context.Context,
	req *usercouponpb.GetUserCouponsRequest,
) (*usercouponpb.GetUserCouponsResponse, error) {

	if req.UserId == "" {
		return nil, status.Error(codes.InvalidArgument, "user_id is required")
	}

	list, err := h.query.ListUserCoupons(ctx, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user coupons: %v", err)
	}

	return &usercouponpb.GetUserCouponsResponse{
		Coupons: dto.ToProtoList(list),
	}, nil
}

// -------------------------
// VALIDATE
// -------------------------

func (h *Handler) ValidateCoupon(
	ctx context.Context,
	req *usercouponpb.ValidateCouponRequest,
) (*usercouponpb.ValidateCouponResponse, error) {

	if req.UserId == "" || req.CouponCode == "" {
		return nil, status.Error(codes.InvalidArgument, "user_id and coupon_code are required")
	}

	ok, msg := h.query.ValidateCoupon(ctx, req.UserId, req.CouponCode)

	var coupon *usercouponpb.UserCoupon

	if ok {
		list, _ := h.query.ListUserCoupons(ctx, req.UserId)
		for _, c := range list {
			if c.CouponCode == req.CouponCode {
				coupon = dto.ToProto(c)
				break
			}
		}
	}

	return &usercouponpb.ValidateCouponResponse{
		Valid:        ok,
		ErrorMessage: msg,
		Coupon:       coupon,
	}, nil
}

// -------------------------
// ðŸ”¥ RESERVE (NEW)
// -------------------------

func (h *Handler) ReserveCoupon(
	ctx context.Context,
	req *usercouponpb.ReserveCouponRequest,
) (*usercouponpb.ReserveCouponResponse, error) {

	if req.UserId == "" || req.CouponCode == "" {
		return nil, status.Error(codes.InvalidArgument, "user_id and coupon_code are required")
	}

	err := h.reserveSvc.ReserveCoupon(
		ctx,
		req.UserId,
		req.CouponCode,
	)

	if err != nil {
		return &usercouponpb.ReserveCouponResponse{
			Success:      false,
			ErrorMessage: err.Error(),
		}, nil
	}

	return &usercouponpb.ReserveCouponResponse{
		Success: true,
	}, nil
}

