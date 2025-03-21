package controllers

import (
	"api-main/users/application"
	"api-main/users/infraestructure"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {

	id_string:= c.Param("id")

	id,err:= strconv.Atoi(id_string)

	if err!=nil{
		c.JSON(http.StatusForbidden,gin.H{"not found":"no se pudo encontrar un id"})
	}

	repo:=infraestructure.NewMySQLRepository()
	useCase:=application.NewDeleteUser(repo)

	if err:=useCase.Execute(id);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"ERROR":"verifique sus datos"})
	}

	c.JSON(http.StatusOK,gin.H{"ok":"se elimino el usuario con exito"})

}