package controllers

import (
	"api-main/mesa/application"
	"api-main/mesa/infraestructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {

	repo:=infraestructure.NewMySQLRepository()
	useCase:=application.NewGetAllMesas(repo)

	mesas,err:=useCase.Execute()

	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"error al ejecutar el metodo"})
	}

	c.JSON(http.StatusOK,mesas)

}