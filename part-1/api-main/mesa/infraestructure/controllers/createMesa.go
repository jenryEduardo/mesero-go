package controllers

import (
	"api-main/mesa/application"
	"api-main/mesa/domain"
	"api-main/mesa/infraestructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateMesa(c *gin.Context) {
	var mesa domain.Mesa

	// Corregido: Usar `&mesa` y devolver 400 en caso de error
	if err := c.ShouldBindJSON(&mesa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se encontraron datos en la solicitud"})
		return

	}

	repo := infraestructure.NewMySQLRepository()
	useCase := application.NewCreateMesa(repo)

	// Corregido: Manejar error correctamente y detener ejecución
	if err := useCase.Execute(&mesa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al ejecutar la solicitud"})
		return
	}

	// Respuesta correcta con 201 Created
	c.JSON(http.StatusCreated, gin.H{
		"message": "Mesa creada exitosamente",
		"data":    mesa,
	})
}
