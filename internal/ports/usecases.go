package ports

import "github.com/igorscandido/go-items-management-with-queues/internal/domain"

type InsertItemUsecase interface {
	PublishItemToQueue(*domain.Item) error
}

type ProcessInsertedItemUsecase interface {
	ProcessItems() error
}
