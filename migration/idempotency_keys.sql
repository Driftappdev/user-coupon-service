CREATE TABLE idempotency_keys (
    id BIGSERIAL PRIMARY KEY,
    event_key TEXT UNIQUE NOT NULL,
    processed_at TIMESTAMP NOT NULL
);