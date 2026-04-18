-- Table: user_coupons
CREATE TABLE IF NOT EXISTS user_coupons (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(50) NOT NULL,
    coupon_code VARCHAR(50) NOT NULL,
    description TEXT,
    discount_amount NUMERIC(10,2) DEFAULT 0,
    min_order_amount NUMERIC(10,2) DEFAULT 0,
    product_type VARCHAR(50),
    status VARCHAR(20) DEFAULT 'active', -- active, used, expired
    valid_from TIMESTAMP NOT NULL,
    valid_until TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Index for faster lookups by user
CREATE INDEX IF NOT EXISTS idx_user_coupons_user_id ON user_coupons(user_id);
CREATE INDEX IF NOT EXISTS idx_user_coupons_coupon_code ON user_coupons(coupon_code);
CREATE INDEX IF NOT EXISTS idx_user_coupons_status ON user_coupons(status);
