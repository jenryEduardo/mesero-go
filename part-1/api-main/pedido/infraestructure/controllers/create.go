package controllers

import (
	"api-main/pedido/application"
	"api-main/pedido/domain"
	"api-main/pedido/infraestructure"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePedido(c *gin.Context) {
	var pedido domain.Pedido

	if err := c.ShouldBindJSON(&pedido); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "no se encontró nada en el archivo JSON"})
		return
	}

	repo := infraestructure.NewMySQLRepository()
	useCase := application.NewSavePedido(repo)

	id, err := useCase.Execute(pedido)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convertir id a string y usarlo en la URL
	url := fmt.Sprintf("http://localhost:3002/consumer/%d", id)

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creando la solicitud HTTP"})
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al enviar la solicitud al otro servidor"})
		return
	}
	defer resp.Body.Close()

	c.JSON(http.StatusOK, gin.H{"ok": id, "message": "ID enviado correctamente como parámetro en la URL"})
}