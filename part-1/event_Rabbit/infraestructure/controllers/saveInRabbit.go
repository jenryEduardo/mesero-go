package controllers

import (
	"consumer/application"
	"consumer/infraestructure/adapters"
	"fmt"
	"strconv"
	"github.com/gin-gonic/gin"
)

func SaveInRbbitmq(c *gin.Context) {

	IdString:=c.Param("idPedido")
	id,err:=strconv.Atoi(IdString)


		repo,err:=adapters.NewRabbitMQRepository()
		if err!=nil{
			fmt.Println("ocurrio un error al conectarse a rabbit")
		}

		useCase:=application.NewRabbitSave(repo)
		success,err:=useCase.Execute(id)

		if err!=nil{
			fmt.Println("error al ejecutar la transaccion")
		}else if success{
			fmt.Println("transaccion exitosa",id)
			c.JSON(201, gin.H{"message": "Consumo iniciado en background"})
		}
		
}