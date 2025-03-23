package controllers

import (
	"api-main/mesa/application"
	"api-main/mesa/domain"
	"api-main/mesa/infraestructure"

	// "api-main/mesa/infraestructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateMesa(c *gin.Context) {

	var mesa domain.Mesa

	if err:= c.ShouldBindJSON(mesa);err!=nil{
		c.JSON(http.StatusNoContent,gin.H{"error":"no se encontro datos en la solicitud"})
	}

	repo:=infraestructure.NewMySQLRepository()
	useCase:=application.NewCreateMesa(repo)

	if err:=useCase.Execute(&mesa);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"error al ejecutar la solicitud"})
	}

	c.JSON(http.StatusAccepted,gin.H{"ok":"se creo con exito"}) 

}