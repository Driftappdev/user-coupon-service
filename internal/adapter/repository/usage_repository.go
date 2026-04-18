package postgres

import (
	"context"
	"database/sql"

	repo "dift_user_insentive/user-coupon-service/internal/interface/repository"
)

type UsageRepository struct {
	db *sql.DB
}

func NewUsageRepository(db *sql.DB) *UsageRepository {
	return &UsageRepository{db: db}
}

// âœ… FIX
var _ repo.UsageRepository = (*UsageRepository)(nil)

func (r *UsageRepository) Create(
	ctx context.Context,
	userID, code, orderID, status string,
) error {

	_, err := r.db.ExecContext(ctx,
		`INSERT INTO coupon_usage (user_id, coupon_code, order_id, status)
		 VALUES ($1,$2,$3,$4)`,
		userID, code, orderID, status,
	)

	return err
}

