package controllers

import (
	"consumer2/application"
	"consumer2/infraestructure/adapters"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Consumer(c *gin.Context) {
	repo, err := adapters.NewRabbitMQRepository()
	if err != nil {
		fmt.Println("❌ Error al conectarse a RabbitMQ:", err)
		c.JSON(500, gin.H{"error": "Error al conectar con RabbitMQ"})
		return
	}

	useCase := application.NewConsume(repo)


	go func() {
		err := useCase.Execute()
		if err != nil {
			fmt.Println(" Error en el caso de uso:", err)
		}
	}()

	c.JSON(200, gin.H{"message": "Consumo iniciado en background"})
}
