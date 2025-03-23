package controllers

import (
	"api-main/mesa/application"
	"api-main/mesa/infraestructure"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteMesa(c *gin.Context) {

	idMesa_int:=c.Param("idMesa")

	idMesa,err:=strconv.Atoi(idMesa_int)

	if err!=nil{
		c.JSON(http.StatusForbidden,gin.H{"error":"no se pudo obtener el id"})
	}

	repo:=infraestructure.NewMySQLRepository()
	useCase:=application.NewDelete(repo)

	if err:=useCase.Execute(idMesa);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"err":"error al ejecutar el metodo"}) 
	}

	c.JSON(http.StatusAccepted,gin.H{"ok":"se elimino con exito"}) 
}