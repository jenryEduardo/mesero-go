package controllers

import (
	"api-main/mesa/application"
	"api-main/mesa/domain"
	"api-main/mesa/infraestructure"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateMesa(c *gin.Context){

	id_string:=c.Param("idMesa")

	id,err:=strconv.Atoi(id_string)

	if err != nil {
		c.JSON(http.StatusNotFound,gin.H{"error":"no se pudo obtener un id"})
	}

	var mesa domain.Mesa

	if err:=c.ShouldBindJSON(mesa);err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"error":"no se pudo deszerializar el json"})
	}

	repo:=infraestructure.NewMySQLRepository()
	useCase:=application.NewUpdateMesa(repo)

	if err:=useCase.Execute(id,&mesa);err!=nil{
		c.JSON(http.StatusBadGateway,gin.H{"error":err})
	}

	c.JSON(http.StatusAccepted,gin.H{"succesfull":"se actualizo la mesa correctemente"})

}