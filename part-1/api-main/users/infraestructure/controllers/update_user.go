package controllers

import (
	"api-main/users/application"
	"api-main/users/domain"
	"api-main/users/infraestructure"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {


	id_string:=c.Param("id")

	id_int,err:=strconv.Atoi(id_string)

	if err!=nil{
		c.JSON(http.StatusForbidden,gin.H{"error":"no se pudo obtener el id"})
	}
	
	var user domain.User

	if err:=c.ShouldBindJSON(user);err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"error":"no se pudo deserializar el archivo json"})
	}


	repo:=infraestructure.NewMySQLRepository()
	useCase:=application.NewUpdateUser(repo)

	if err:=useCase.Execute(id_int,user);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"failed":"no se pudo realizar la actualizacion"})
	}


	c.JSON(http.StatusOK,gin.H{"ok":"se realizo correctamente la actualizacion"})

}