package controllers

import (
	"api-main/users/application"
	"api-main/users/infraestructure"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {

	id_user_string:=c.Param("id")

	id,err:=strconv.Atoi(id_user_string)

	if err!=nil{
		fmt.Println("no se pudo obtener el id")
		return
	}

	repo:=infraestructure.NewMySQLRepository()
	useCase:=application.NewGetUserById(repo)

	users,err:=useCase.Execute(id)

	if err!=nil{
		fmt.Println("no se pudo obtener al usuario",err)
	}

	c.JSON(http.StatusOK,users)

}