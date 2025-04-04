package controllers

import (
	"api-main/users/infraestructure"
	"net/http"
	"github.com/gin-gonic/gin"
)

// Login controlador para el login del usuario
func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Crear un repositorio de MySQL
	repo := infraestructure.NewMySQLRepository()
	// Crear un caso de uso para login
	user, token, err := repo.Login(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Devolver el token JWT al usuario
	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": token,
	})
}
