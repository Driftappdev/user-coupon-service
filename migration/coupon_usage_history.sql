CREATE TABLE coupon_usage_history (
    id BIGSERIAL PRIMARY KEY,
    user_id TEXT NOT NULL,
    coupon_code TEXT NOT NULL,
    order_id TEXT NOT NULL,
    status TEXT NOT NULL,
    used_at TIMESTAMP NOT NULL
);