package ports

import (
	"context"

	"github.com/igorscandido/go-items-management-with-queues/internal/domain"
)

type ItemsRepository interface {
	InsertItem(ctx context.Context, item *domain.Item) (*int, error)
}
