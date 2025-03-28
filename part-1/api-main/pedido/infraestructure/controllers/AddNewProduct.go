package controllers

import (
	"api-main/pedido/application"
	"api-main/pedido/domain"
	"api-main/pedido/infraestructure"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddProductInPedido(c *gin.Context){


	id_string:=c.Param("idPedido")

	id,err:=strconv.Atoi(id_string)

	if err!=nil{
		c.JSON(http.StatusForbidden,gin.H{"ERROR":"no se encontro ningun id"})
	}

	var data domain.DetallesPedido

	if err:=c.ShouldBindJSON(&data);err!=nil{
		c.JSON(http.StatusForbidden,gin.H{"error":"no se pudo deserializar el json"})
	}

	repo:=infraestructure.NewMySQLRepository()
	useCase:=application.NewAddProduct(repo)

	if err:=useCase.Execute(id,&data);err!=nil{
		c.JSON(http.StatusConflict,gin.H{"error":err})
		return 
	}

	c.JSON(http.StatusOK,gin.H{"ok":"producto agregado"})

}