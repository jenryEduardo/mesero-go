package controllers

import (
	"api-main/mesa/application"
	"api-main/mesa/domain"
	"api-main/mesa/infraestructure"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateMesa(c *gin.Context) {
	idString := c.Param("idMesa")

	// Convertir ID a entero
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se pudo obtener un ID válido"})
		return
	}

	var mesa domain.Mesa


	if err:=c.ShouldBindJSON(&mesa);err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"error":"no se pudo deszerializar el json"})


	repo := infraestructure.NewMySQLRepository()
	useCase := application.NewUpdateMesa(repo)

	// Corregido: Manejar error correctamente
	if err := useCase.Execute(id, &mesa); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	// Respuesta correcta con 202 Accepted
	c.JSON(http.StatusAccepted, gin.H{"success": "Se actualizó la mesa correctamente"})
}
}