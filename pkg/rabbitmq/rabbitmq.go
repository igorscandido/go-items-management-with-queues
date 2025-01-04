package rabbitmq

import (
	"fmt"
	"time"

	"github.com/igorscandido/go-items-management-with-queues/pkg/configs"
	"github.com/streadway/amqp"
)

func NewRabbitMQConnection(configs *configs.Configs) *amqp.Connection {
	url := fmt.Sprintf("amqp://%s:%s@%s:%d%s",
		configs.RabbitMQ.User,
		configs.RabbitMQ.Password,
		configs.RabbitMQ.Address,
		configs.RabbitMQ.Port,
		configs.RabbitMQ.VHost,
	)
	var conn *amqp.Connection
	for i := 0; i < configs.RabbitMQ.RetryPolicy.MaxRetries; i++ {
		var err error
		conn, err = amqp.Dial(url)
		if err != nil {
			fmt.Printf("Failed to connect to RabbitMQ: %v\n", err)
			time.Sleep(time.Duration(configs.RabbitMQ.RetryPolicy.Interval) * time.Second)
			continue
		}
		break
	}

	return conn
}
