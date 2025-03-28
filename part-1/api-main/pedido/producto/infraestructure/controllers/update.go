package controllers

import (
	"api-main/pedido/producto/application"
	"api-main/pedido/producto/domain"
	"api-main/pedido/producto/infraestructure"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateProduct(c *gin.Context) {

	IDstring:=c.Param("idProducto")

	id,err:=strconv.Atoi(IDstring)

	if err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"error":"no se encontro un id en la solicitud verifique la ruta"})
	}

	var producto domain.Producto

	if err:=c.ShouldBindJSON(&producto);err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"error":"no se pudo encontrar nada el json de la solicitud"})
	}

	repo:=infraestructure.NewMySQLRepository()
	useCase:=application.NewUpdateProduct(repo)

	if err:=useCase.Execute(id,&producto);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"ERROR":"no se pudo realizar la solicitud verifique la sintaxis de mysql"})
	}

	c.JSON(http.StatusCreated,gin.H{"ok":"se creo el producto correctamente"})

}