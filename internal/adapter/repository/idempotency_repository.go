package postgres

import (
	"context"
	"database/sql"

	repo "dift_user_insentive/user-coupon-service/internal/interface/repository"
)

type IdempotencyRepository struct {
	db *sql.DB
}

func NewIdempotencyRepository(db *sql.DB) *IdempotencyRepository {
	return &IdempotencyRepository{db: db}
}

// âœ… FIX
var _ repo.IdempotencyRepository = (*IdempotencyRepository)(nil)

func (r *IdempotencyRepository) Exists(ctx context.Context, key string) (bool, error) {
	var exists bool
	err := r.db.QueryRowContext(ctx,
		`SELECT EXISTS (SELECT 1 FROM coupon_idempotency WHERE key=$1)`,
		key,
	).Scan(&exists)
	return exists, err
}

func (r *IdempotencyRepository) Save(ctx context.Context, key string) error {
	_, err := r.db.ExecContext(ctx,
		`INSERT INTO coupon_idempotency (key) VALUES ($1) ON CONFLICT DO NOTHING`,
		key,
	)
	return err
}

