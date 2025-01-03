package rabbitmq

import (
	"log"

	"github.com/igorscandido/go-items-management-with-queues/internal/ports"
	"github.com/streadway/amqp"
)

type rabbitMQProducer struct {
	channel   *amqp.Channel
	queueName string
}

func NewRabbitMQProducer(connection *amqp.Connection, queueName string) (ports.RabbitMQProducer, error) {
	ch, err := connection.Channel()
	if err != nil {
		return nil, err
	}

	_, err = ch.QueueDeclare(
		queueName,
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,
	)
	if err != nil {
		ch.Close()
		return nil, err
	}

	return &rabbitMQProducer{
		channel:   ch,
		queueName: queueName,
	}, nil
}

func (p *rabbitMQProducer) PublishMessage(message string) error {
	err := p.channel.Publish(
		"",          // exchange
		p.queueName, // routing key
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		return err
	}

	log.Printf("Message sent to queue %s: %s", p.queueName, message)
	return nil
}

func (p *rabbitMQProducer) CloseChannel() {
	p.channel.Close()
}
