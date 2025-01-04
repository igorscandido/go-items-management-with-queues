package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/igorscandido/go-items-management-with-queues/internal/app/usecases"
	"github.com/igorscandido/go-items-management-with-queues/internal/infra/http"
	"github.com/igorscandido/go-items-management-with-queues/internal/infra/rabbitmq"
	"github.com/igorscandido/go-items-management-with-queues/pkg/configs"
	pkgRabbitmq "github.com/igorscandido/go-items-management-with-queues/pkg/rabbitmq"
)

func main() {
	r := gin.Default()
	configs := configs.NewConfigs()

	rabbitmqConnection := pkgRabbitmq.NewRabbitMQConnection(configs)
	if rabbitmqConnection == nil {
		panic("Failed to connect to RabbitMQ")
	}
	defer rabbitmqConnection.Close()

	rabbitmqProducer, err := rabbitmq.NewRabbitMQProducer(rabbitmqConnection, rabbitmq.InsertItemsQueue)
	if err != nil {
		panic(fmt.Sprintf("Failed to create RabbitMQ producer: %v", err))
	}
	defer rabbitmqProducer.CloseChannel()

	itemsService := usecases.NewInsertItemUsecase(rabbitmqProducer)
	handler := http.NewItemsHandler(itemsService)

	r.POST("/items", handler.InsertItem)
	r.Run(":8080")
}
