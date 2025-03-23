package controllers

import (
	"api-main/pedido/application"
	"api-main/pedido/infraestructure"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeletePedido(c *gin.Context){

	id_string:= c.Param("idPedido")

	id,err:=strconv.Atoi(id_string)

	if err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"error":"no se pudo obtener el id"})
	}

	repo:=infraestructure.NewMySQLRepository()
	useCase:=application.NewDelete(repo)

	if err:=useCase.Execute(id);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"no se pudo realizar la solicitud"})
	}

	c.JSON(http.StatusAccepted,gin.H{"ok":"se elimino correctamente"})

}