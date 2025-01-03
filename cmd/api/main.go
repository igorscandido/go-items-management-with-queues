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

	rabbitmqConnection, err := pkgRabbitmq.NewRabbitMQConnection(configs)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to RabbitMQ: %v", err))
	}

	rabbitmqProducer, err := rabbitmq.NewRabbitMQProducer(rabbitmqConnection, "items")
	if err != nil {
		panic(fmt.Sprintf("Failed to create RabbitMQ producer: %v", err))
	}

	itemsService := usecases.NewCreateItemUsecase(rabbitmqProducer)
	handler := http.NewItemsHandler(itemsService)

	r.POST("/items", handler.CreateItem)
	r.Run(":8080")
}
