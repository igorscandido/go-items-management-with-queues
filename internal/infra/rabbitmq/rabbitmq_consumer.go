package rabbitmq

import (
	"log"
	"sync"

	"github.com/igorscandido/go-items-management-with-queues/internal/ports"
	"github.com/streadway/amqp"
)

type rabbitMQConsumer struct {
	channel   *amqp.Channel
	queueName string
}

func NewRabbitMQConsumer(connection *amqp.Connection, queueName string) (ports.RabbitMQConsumer, error) {
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

	return &rabbitMQConsumer{
		channel:   ch,
		queueName: queueName,
	}, nil
}

func (c *rabbitMQConsumer) ConsumeMessages(processFunc func(string) error) error {
	msgs, err := c.channel.Consume(
		c.queueName,
		"",    // consumer tag
		false, // auto-ack (set to false for manual ack)
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,
	)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Printf("Reading messages from queue...")
		for msg := range msgs {
			if err := processFunc(string(msg.Body)); err != nil {
				log.Printf("Failed to process message: %v", err)
				msg.Nack(false, true)
			} else {
				msg.Ack(false)
			}
		}
	}()

	wg.Wait()
	return nil
}

func (c *rabbitMQConsumer) CloseChannel() {
	c.channel.Close()
}
