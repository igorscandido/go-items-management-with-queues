package usecases

import (
	"context"
	"log"

	"github.com/igorscandido/go-items-management-with-queues/internal/domain"
	"github.com/igorscandido/go-items-management-with-queues/internal/ports"
)

type processInsertedItemUsecase struct {
	rabbitMQConsumer ports.RabbitMQConsumer
	itemsRepository  ports.ItemsRepository
}

func NewProcessInsertedItemUsecase(
	rabbitMQConsumer ports.RabbitMQConsumer,
	itemsRepository ports.ItemsRepository,
) ports.ProcessInsertedItemUsecase {
	return &processInsertedItemUsecase{
		rabbitMQConsumer: rabbitMQConsumer,
		itemsRepository:  itemsRepository,
	}
}

func (u *processInsertedItemUsecase) ProcessItems() error {
	return u.rabbitMQConsumer.ConsumeMessages(u.processInsertedItem)
}

func (u *processInsertedItemUsecase) processInsertedItem(message string) error {
	var item domain.Item
	if err := item.FromJson(message); err != nil {
		log.Printf("error parsing message to item: %v", err)
		return err
	}
	if _, err := u.itemsRepository.InsertItem(context.Background(), &item); err != nil {
		log.Printf("error inserting item: %v", err)
		return err
	}
	return nil
}
