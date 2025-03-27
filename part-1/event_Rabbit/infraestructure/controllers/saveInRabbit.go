package controllers

import (
	"consumer/application"
	"consumer/domain"
	"consumer/infraestructure/adapters"
	"fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

func SaveInRbbitmq(c *gin.Context) {


		var cuenta domain.RabbitMQ
		if err := c.ShouldBindJSON(&cuenta); err != nil {
			log.Println("Error al procesar el JSON:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar el JSON", "details": err.Error()})
			return
		}
	

	

		repo,err:=adapters.NewRabbitMQRepository()
		if err!=nil{
			fmt.Println("ocurrio un error al conectarse a rabbit")
		}

		useCase:=application.NewRabbitSave(repo)
		success,err:=useCase.Execute(&cuenta)

		if err!=nil{
			fmt.Println("error al ejecutar la transaccion")
		}else if success{
			fmt.Println("transaccion exitosa")
		}
		
		c.JSON(200, gin.H{"message": "Consumo iniciado en background"})

}