package controllers

import (
	"api-main/users/application"
	"api-main/users/infraestructure"
	"net/http"
	"github.com/gin-gonic/gin"
)


func GetuserAll(c *gin.Context){


	repo:= infraestructure.NewMySQLRepository()
	useCase:=application.NewGetAllUser(repo)

	data,err:=useCase.Execute();
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"verifique el archivo mysql"})
		return
	}

	c.JSON(http.StatusAccepted,data)

}