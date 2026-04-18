package repository

import "context"

//
// ðŸ”¹ Idempotency Repository
//

type IdempotencyRepository interface {
	Exists(ctx context.Context, key string) (bool, error)
	Save(ctx context.Context, key string) error
}

