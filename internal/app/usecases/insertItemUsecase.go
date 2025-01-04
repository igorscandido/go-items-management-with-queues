package usecases

import (
	"github.com/igorscandido/go-items-management-with-queues/internal/domain"
	"github.com/igorscandido/go-items-management-with-queues/internal/ports"
)

type insertItemUsecase struct {
	rabbitmqProducer ports.RabbitMQProducer
}

func NewInsertItemUsecase(rabbitmqProducer ports.RabbitMQProducer) ports.InsertItemUsecase {
	return &insertItemUsecase{
		rabbitmqProducer,
	}
}

func (u *insertItemUsecase) PublishItemToQueue(item *domain.Item) error {
	return u.rabbitmqProducer.PublishMessage(item.ToJson())
}
