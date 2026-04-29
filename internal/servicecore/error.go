package servicecore

import "dift_user_insentive/user-coupon-service/pkg/wrapper"

func NewServiceError(code string, status int, message string) *wrapper.BaseError {
	return wrapper.NewBaseError(code, status, message)
}

