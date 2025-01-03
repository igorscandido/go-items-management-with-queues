package usecases

import (
	"github.com/igorscandido/go-items-management-with-queues/internal/domain"
	"github.com/igorscandido/go-items-management-with-queues/internal/ports"
)

type createItemUsecase struct {
	rabbitmqProducer ports.RabbitMQProducer
}

func NewCreateItemUsecase(rabbitmqProducer ports.RabbitMQProducer) ports.CreateItemUsecase {
	return &createItemUsecase{
		rabbitmqProducer,
	}
}

func (u *createItemUsecase) PublishItemToQueue(item *domain.Item) error {
	return u.rabbitmqProducer.PublishMessage(item.ToJson())
}
