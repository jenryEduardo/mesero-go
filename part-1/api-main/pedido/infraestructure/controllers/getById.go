package controllers

import (
	"api-main/pedido/application"
	"api-main/pedido/infraestructure"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetByIdPedido(c *gin.Context) {

	idstring:=c.Param("idPedido")

	id,err:=strconv.Atoi(idstring)

	if err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"error":"no se encontro ningun id"})
	}

	repo := infraestructure.NewMySQLRepository()
	useCase := application.NewGetById(repo)

	pedidos, err := useCase.Execute(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no se pudo realizar la solicitud al servidor"})
	}

	c.JSON(http.StatusAccepted, pedidos)

}