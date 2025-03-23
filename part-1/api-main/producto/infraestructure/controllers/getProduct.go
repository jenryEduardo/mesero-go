package controllers

import (
	"api-main/producto/application"
	"api-main/producto/infraestructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllProducts(c *gin.Context){

	repo:=infraestructure.NewMySQLRepository()
	useCase:=application.NewGetALLproducts(repo)

	 data,err:=useCase.Execute()

	 if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"verifique la solicitud o la sintaxis de mysql"})
	 }


	c.JSON(http.StatusCreated,data)
}