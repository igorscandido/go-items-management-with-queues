package ports

type RabbitMQProducer interface {
	PublishMessage(message string) error
	CloseChannel()
}
