package main

import (
	"fmt"

	"github.com/igorscandido/go-items-management-with-queues/internal/app/usecases"
	"github.com/igorscandido/go-items-management-with-queues/internal/infra/rabbitmq"
	"github.com/igorscandido/go-items-management-with-queues/internal/infra/repositories"
	"github.com/igorscandido/go-items-management-with-queues/pkg/configs"
	"github.com/igorscandido/go-items-management-with-queues/pkg/database"
	pkgRabbitmq "github.com/igorscandido/go-items-management-with-queues/pkg/rabbitmq"
)

func main() {
	configs := configs.NewConfigs()

	databaseConnection, err := database.NewPostgresAdapter(configs)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to postgres database: %v", err))
	}
	defer databaseConnection.Close()

	rabbitmqConnection := pkgRabbitmq.NewRabbitMQConnection(configs)
	if rabbitmqConnection == nil {
		panic("failed to connect to RabbitMQ")
	}
	defer rabbitmqConnection.Close()

	rabbitMQConsumer, err := rabbitmq.NewRabbitMQConsumer(rabbitmqConnection, rabbitmq.InsertItemsQueue)
	if err != nil {
		panic(fmt.Sprintf("failed to create RabbitMQ producer: %v", err))
	}
	defer rabbitMQConsumer.CloseChannel()

	itemsRepository := repositories.NewItemsRepository(databaseConnection)

	processInsertedItemUsecase := usecases.NewProcessInsertedItemUsecase(rabbitMQConsumer, itemsRepository)
	processInsertedItemUsecase.ProcessItems()
}
