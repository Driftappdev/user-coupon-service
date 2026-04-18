package repository

import "context"

//
// ðŸ”¹ Usage Repository
//

type UsageRepository interface {
	Create(ctx context.Context, userID, code, orderID, status string) error
}

