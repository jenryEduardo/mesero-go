package controllers

import (
	"fmt"
	"log"
	"net/http"
	"publisher/application"
	"publisher/domain"
	"publisher/infraestructure/adapters"

	"github.com/gin-gonic/gin"
)

func PublisherInRabbit(c *gin.Context) {
	var status domain.EventRabbit

	if err := c.ShouldBindJSON(&status); err != nil {
		log.Println("Error al procesar el JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar el JSON", "details": err.Error()})
		return
	}

	repo, err := adapters.NewRabbitMQ()
	if err!= nil {
		fmt.Println("Ocurrio un percanse al conectarse a RabbitMQ")
	}

	uc := application.NewPublishInRabbit(repo)
	suc, err := uc.Run(&status)

	if err != nil {
		fmt.Println("Error al ejecutar la transacción")
	} else if suc {
		fmt.Println("Transacción exitosa")
	}

	c.JSON(http.StatusOK, gin.H{"message":"Consumo inicado en background"})
}