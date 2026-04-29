package http

import (
	"encoding/json"
	"net/http"

	port "dift_user_insentive/user-coupon-service/internal/interface/http"
)

type Handler struct {
	query port.CouponQueryPort
}

func New(query port.CouponQueryPort) *Handler {
	return &Handler{
		query: query,
	}
}

type GetUserCouponsResponse struct {
	Coupons interface{} `json:"coupons"`
}

func (h *Handler) GetUserCoupons(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "user_id is required", http.StatusBadRequest)
		return
	}

	coupons, err := h.query.ListUserCoupons(ctx, userID)
	if err != nil {
		http.Error(w, "failed to fetch coupons", http.StatusInternalServerError)
		return
	}

	resp := GetUserCouponsResponse{
		Coupons: coupons,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}

