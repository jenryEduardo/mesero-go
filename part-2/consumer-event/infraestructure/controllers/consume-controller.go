package controllers

import (
	"consumer-event/application"
	"consumer-event/infraestructure/adapters"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func Consumer(c *gin.Context) {


	repo, err := adapters.NewRabbitMQRepository()
if err != nil {
    log.Fatalf("Error inicializando RabbitMQ: %v", err)
}

	useCase := application.NewConsumeRabbit(repo)

	go func() {
		err := useCase.Execute()
		if err != nil {
			fmt.Println(" Error en el caso de uso:", err)
		}
	}()

	c.JSON(200, gin.H{"message": "Consumo iniciado en background"})
}
