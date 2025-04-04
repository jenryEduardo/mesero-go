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

func CreatePedido(c *gin.Context){

	var pedido domain.Pedido

	if err:=c.ShouldBindJSON(&pedido);err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"error":"no se encontro nada el archivo json"})
	}

	repo:=infraestructure.NewMySQLRepository()
	useCase:=application.NewSavePedido(repo)

	 id,err:=useCase.Execute(pedido);
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err})
		return
	}
  
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

	c.JSON(http.StatusOK, gin.H{"ok": id, "message": "ID enviado correctamente"})


}