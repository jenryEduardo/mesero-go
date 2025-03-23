package controllers

import (
	"api-main/pedido/application"
	"api-main/pedido/infraestructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllPedido(c *gin.Context){


	repo:=infraestructure.NewMySQLRepository()
	useCase:=application.NewGetAllPedidos(repo)

	pedidos,err:=useCase.Execute()

	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"no se pudo realizar la solicitud al servidor"})
	}

	c.JSON(http.StatusAccepted,pedidos)

}