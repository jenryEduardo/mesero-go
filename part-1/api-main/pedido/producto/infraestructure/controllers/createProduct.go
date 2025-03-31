package controllers

import (
	"api-main/pedido/producto/application"
	"api-main/pedido/producto/domain"
	"api-main/pedido/producto/infraestructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {

	var producto domain.Producto

	if err:=c.ShouldBindJSON(&producto);err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"error":"no se pudo encontrar nada el json de la solicitud"})
	}

	repo:=infraestructure.NewMySQLRepository()
	useCase:=application.NewProducto(repo)

	if err:=useCase.Execute(&producto);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"ERROR":"no se pudo realizar la solicitud verifique la sintaxis de mysql"})
	}

	c.JSON(http.StatusCreated,gin.H{"ok":"se creo el producto correctamente"})



}