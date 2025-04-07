package controllers

import (
	"api-main/pedido/application"
	"api-main/pedido/domain"
	"api-main/pedido/infraestructure"
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePedido(c *gin.Context) {
	var pedido domain.Pedido

	// Validar JSON de entrada
	if err := c.ShouldBindJSON(&pedido); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se pudo procesar el JSON de entrada"})
		return
	}

	// Crear repositorio y caso de uso
	repo := infraestructure.NewMySQLRepository()
	useCase := application.NewSavePedido(repo)

	// Ejecutar el caso de uso para guardar el pedido
	id, err := useCase.Execute(pedido)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Crear JSON para enviar al otro servidor
	payload, err := json.Marshal(gin.H{"id": id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el JSON de salida"})
		return
	}

	// Enviar el JSON al otro servidor
	resp, err := http.Post("http://localhost:3002/consumer/", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al enviar el ID al otro servidor"})
		return
	}
	defer resp.Body.Close()

	// Verificar si el servidor externo responde correctamente
	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "El servidor externo devolvió un estado inesperado", "status": resp.StatusCode})
		return
	}

	// Respuesta exitosa
	c.JSON(http.StatusOK, gin.H{"ok": id, "message": "ID enviado correctamente"})
}
