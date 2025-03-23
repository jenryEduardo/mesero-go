package controllers

import (
	"api-main/pedido/application"
	"api-main/pedido/domain"
	"api-main/pedido/infraestructure"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdatePedido(c *gin.Context){

	idString:=c.Param("idPedido")
	id,err:=strconv.Atoi(idString)

	if err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"error":"no se pudo encontrar un id verifique su solicitud"})
	}

	var pedido domain.Pedido

	if err:=c.ShouldBindJSON(pedido);err!=nil{
		c.JSON(http.StatusForbidden,gin.H{"error":"no se pudo decodificar el archivo json"})
	}

	repo:=infraestructure.NewMySQLRepository()
	useCase:=application.NewUpdatePedido(repo)

	if err:=useCase.Execute(id,pedido);err!=nil{
		 c.JSON(http.StatusConflict,gin.H{"error":"no se pudo realizar la solicitud"})
	}


	c.JSON(http.StatusOK,gin.H{"ok":"se actualizo el pedido correctamente"})

}