package ports

type RabbitMQProducer interface {
	PublishMessage(message string) error
	CloseChannel()
}

type RabbitMQConsumer interface {
	ConsumeMessages(processFunc func(string) error) error
	CloseChannel()
}
