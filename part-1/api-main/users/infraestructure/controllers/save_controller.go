package controllers

import (
	"api-main/users/application"
	"api-main/users/domain"
	"api-main/users/infraestructure"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context){

	var user domain.User

	// Intenta deserializar el JSON
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println("Error al procesar el JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar el JSON", "details": err.Error()})
		return
	}


	repo:=infraestructure.NewMySQLRepository()
	useCase:=application.NewPostUser(repo)

	if err := useCase.Execute(user);err!=nil{
		fmt.Println("error al guardar los datos",err)
	}
	
	c.JSON(http.StatusOK,gin.H{"ok":"dato guardado exitosamente", "user":user})


}