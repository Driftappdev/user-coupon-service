package postgres

import (
	"context"
	"database/sql"
	"time"

	repo "dift_user_insentive/user-coupon-service/internal/interface/repository"
	"dift_user_insentive/user-coupon-service/internal/model"
)

type UserCouponRepository struct {
	db *sql.DB
}

func NewUserCouponRepository(db *sql.DB) *UserCouponRepository {
	return &UserCouponRepository{db: db}
}

// âœ… FIX: à¹€à¸«à¸¥à¸·à¸­à¹à¸„à¹ˆà¸™à¸µà¹‰à¸žà¸­
var _ repo.UserCouponRepository = (*UserCouponRepository)(nil)

//
// WRITE
//

func (r *UserCouponRepository) Upsert(ctx context.Context, c *model.UserCoupon) error {
	_, err := r.db.ExecContext(ctx, `
	INSERT INTO user_coupons (
		user_id, coupon_code, status,
		discount_type, discount_value,
		min_order, max_discount,
		product_type,
		valid_from, valid_to,
		max_usage, used_count
	)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)
	ON CONFLICT (user_id,coupon_code)
	DO UPDATE SET
		status = EXCLUDED.status,
		discount_type = EXCLUDED.discount_type,
		discount_value = EXCLUDED.discount_value,
		min_order = EXCLUDED.min_order,
		max_discount = EXCLUDED.max_discount,
		product_type = EXCLUDED.product_type,
		valid_from = EXCLUDED.valid_from,
		valid_to = EXCLUDED.valid_to,
		max_usage = EXCLUDED.max_usage,
		used_count = EXCLUDED.used_count
	`,
		c.UserID,
		c.CouponCode,
		c.Status,
		c.DiscountType,
		c.DiscountValue,
		c.MinOrder,
		c.MaxDiscount,
		c.ProductType,
		c.ValidFrom,
		c.ValidTo,
		c.MaxUsage,
		c.UsedCount,
	)
	return err
}

func (r *UserCouponRepository) IncrementUsage(ctx context.Context, userID, code string) error {
	_, err := r.db.ExecContext(ctx,
		`UPDATE user_coupons SET used_count = used_count + 1 WHERE user_id=$1 AND coupon_code=$2`,
		userID, code,
	)
	return err
}

func (r *UserCouponRepository) Delete(ctx context.Context, userID, code string) error {
	_, err := r.db.ExecContext(ctx,
		`DELETE FROM user_coupons WHERE user_id=$1 AND coupon_code=$2`,
		userID, code,
	)
	return err
}

//
// READ
//

func (r *UserCouponRepository) FindByUser(ctx context.Context, userID string) ([]*model.UserCoupon, error) {
	rows, err := r.db.QueryContext(ctx, `
	SELECT user_id, coupon_code, status,
	       discount_type, discount_value,
	       min_order, max_discount,
	       product_type,
	       valid_from, valid_to,
	       max_usage, used_count
	FROM user_coupons
	WHERE user_id=$1
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []*model.UserCoupon

	for rows.Next() {
		c := &model.UserCoupon{}
		if err := rows.Scan(
			&c.UserID,
			&c.CouponCode,
			&c.Status,
			&c.DiscountType,
			&c.DiscountValue,
			&c.MinOrder,
			&c.MaxDiscount,
			&c.ProductType,
			&c.ValidFrom,
			&c.ValidTo,
			&c.MaxUsage,
			&c.UsedCount,
		); err != nil {
			return nil, err
		}
		res = append(res, c)
	}

	return res, rows.Err()
}

func (r *UserCouponRepository) FindOne(ctx context.Context, userID, code string) (*model.UserCoupon, error) {
	row := r.db.QueryRowContext(ctx, `
	SELECT user_id, coupon_code, status,
	       discount_type, discount_value,
	       min_order, max_discount,
	       product_type,
	       valid_from, valid_to,
	       max_usage, used_count
	FROM user_coupons
	WHERE user_id=$1 AND coupon_code=$2
	`, userID, code)

	c := &model.UserCoupon{}
	if err := row.Scan(
		&c.UserID,
		&c.CouponCode,
		&c.Status,
		&c.DiscountType,
		&c.DiscountValue,
		&c.MinOrder,
		&c.MaxDiscount,
		&c.ProductType,
		&c.ValidFrom,
		&c.ValidTo,
		&c.MaxUsage,
		&c.UsedCount,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // âœ… clean à¸à¸§à¹ˆà¸²
		}
		return nil, err
	}

	return c, nil
}

func (r *UserCouponRepository) FindByCode(ctx context.Context, code string) (*model.UserCoupon, error) {
	row := r.db.QueryRowContext(ctx, `
	SELECT user_id, coupon_code, status,
	       discount_type, discount_value,
	       min_order, max_discount,
	       product_type,
	       valid_from, valid_to,
	       max_usage, used_count
	FROM user_coupons
	WHERE coupon_code=$1
	LIMIT 1
	`, code)

	c := &model.UserCoupon{}
	if err := row.Scan(
		&c.UserID,
		&c.CouponCode,
		&c.Status,
		&c.DiscountType,
		&c.DiscountValue,
		&c.MinOrder,
		&c.MaxDiscount,
		&c.ProductType,
		&c.ValidFrom,
		&c.ValidTo,
		&c.MaxUsage,
		&c.UsedCount,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return c, nil
}

//
// EXPIRE (à¸¢à¸±à¸‡à¹ƒà¸Šà¹‰à¹„à¸”à¹‰à¸›à¸à¸•à¸´)
//

func (r *UserCouponRepository) ExpireBefore(ctx context.Context, t time.Time) error {
	_, err := r.db.ExecContext(ctx,
		`UPDATE user_coupons SET status='EXPIRED' WHERE valid_to < $1`,
		t,
	)
	return err
}

