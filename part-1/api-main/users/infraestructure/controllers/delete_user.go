package controllers

import (
	"api-main/users/application"
	"api-main/users/infraestructure"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	// Obtener el id desde los parámetros de la URL
	id_string := c.Param("id")

	// Convertir el id de string a int
	id, err := strconv.Atoi(id_string)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Crear la instancia del repositorio
	repo := infraestructure.NewMySQLRepository()

	// Crear el caso de uso de eliminar usuario
	useCase := application.NewDeleteUser(repo)

	// Ejecutar el caso de uso
	if err := useCase.Execute(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se pudo eliminar el usuario"})
		return
	}

	// Retornar respuesta exitosa
	c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado con éxito"})
}
