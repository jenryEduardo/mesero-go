package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"second/pedido/domain"
	"github.com/gin-gonic/gin"
)


func GetIdPedido(g *gin.Context) {
	var idPedido domain.Pedido

	if err := g.ShouldBindJSON(&idPedido.IdPedido); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "No se pudo obtener el id"})
		return
	}

	jsonData, err := json.Marshal(idPedido)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error al serializar el JSON"})
		return
	}

	req, err := http.NewRequest("POST", "http://localhost:3010/enviarPedido", bytes.NewBuffer(jsonData))
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la solicitud"})
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error al enviar la solicitud"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		g.JSON(resp.StatusCode, gin.H{"error": "Error en la respuesta del servidor externo"})
		return
	}

	g.JSON(http.StatusOK, gin.H{"message": "Pedido enviado correctamente"})
}
