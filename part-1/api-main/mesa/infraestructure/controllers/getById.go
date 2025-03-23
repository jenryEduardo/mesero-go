package controllers

import (
	"api-main/mesa/application"
	"api-main/mesa/infraestructure"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetById(c *gin.Context) {

	idMesa_string:=c.Param("idMesa")

	IdMesa,err:=strconv.Atoi(idMesa_string)

	if err!=nil{
		c.JSON(http.StatusForbidden,gin.H{"error":"error no se obtuvo un id"})
	}

	repo := infraestructure.NewMySQLRepository()
	useCase :=application.NewGetById(repo)

	mesas, err := useCase.Execute(IdMesa)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error al ejecutar el metodo"})
	}

	c.JSON(http.StatusOK,mesas)

}