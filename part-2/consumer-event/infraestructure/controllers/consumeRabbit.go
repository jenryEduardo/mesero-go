package controllers

import (
	"consumer-event/infraestructure/adapters"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Inicia el consumidor de RabbitMQ en segundo plano
func StartRabbitMQConsumer() {
	repo, err := adapters.NewRabbitMQRepository()
	if err != nil {
		log.Fatalf("Error inicializando RabbitMQ: %v", err)
	}
	
	fmt.Println("Iniciando consumidor de RabbitMQ...")
	errors := repo.ConsumeTransaction()
	if errors != nil {
		log.Fatal("Error en el consumidor:", err)
	}
}

// Handler HTTP para iniciar el consumo manualmente
func ConsumeRabbitMQ(c *gin.Context) {
	go StartRabbitMQConsumer() // Ejecutar en background

	c.JSON(http.StatusOK, gin.H{"message": "Consumo de RabbitMQ iniciado en background"})
}
