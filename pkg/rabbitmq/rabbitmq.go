package rabbitmq

import (
	"fmt"

	"github.com/igorscandido/go-items-management-with-queues/pkg/configs"
	"github.com/streadway/amqp"
)

func NewRabbitMQConnection(configs *configs.Configs) (*amqp.Connection, error) {
	url := fmt.Sprintf("amqp://%s:%s@%s:%d%s",
		configs.RabbitMQ.User,
		configs.RabbitMQ.Password,
		configs.RabbitMQ.Address,
		configs.RabbitMQ.Port,
		configs.RabbitMQ.VHost,
	)
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
