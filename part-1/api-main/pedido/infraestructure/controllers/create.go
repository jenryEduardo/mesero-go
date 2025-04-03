package controllers

import (
	"api-main/pedido/application"
	"api-main/pedido/domain"
	"api-main/pedido/infraestructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePedido(c *gin.Context){

	var pedido domain.Pedido

	if err:=c.ShouldBindJSON(&pedido);err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"error":"no se encontro nada el archivo json"})
	}

	repo:=infraestructure.NewMySQLRepository()
	useCase:=application.NewSavePedido(repo)

	 id,err:=useCase.Execute(pedido);
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err})
	}
  
	c.JSON(http.StatusOK,gin.H{"ok":id})


}