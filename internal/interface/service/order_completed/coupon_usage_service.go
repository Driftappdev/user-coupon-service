package order_completed

import "context"

type OrderCompletedService interface {
	UseCoupon(
		ctx context.Context,
		userID string,
		code string,
		orderID string,
	) error
}

