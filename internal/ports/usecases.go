package ports

import "github.com/igorscandido/go-items-management-with-queues/internal/domain"

type CreateItemUsecase interface {
	PublishItemToQueue(*domain.Item) error
}
