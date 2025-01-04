package repositories

import (
	"context"

	"github.com/igorscandido/go-items-management-with-queues/internal/domain"
	"github.com/igorscandido/go-items-management-with-queues/internal/infra/repositories/queries"
	"github.com/igorscandido/go-items-management-with-queues/internal/ports"
)

type itemsRepository struct {
	database ports.Database
}

func NewItemsRepository(database ports.Database) ports.ItemsRepository {
	return &itemsRepository{
		database,
	}
}

func (r *itemsRepository) InsertItem(ctx context.Context, item *domain.Item) (*int, error) {
	row := r.database.
		QueryRow(
			ctx,
			queries.InsertItem,
			item.Name,
			item.Description,
			item.Price,
			item.Stock,
			item.Status,
		)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return nil, err
	}
	return &id, nil
}
