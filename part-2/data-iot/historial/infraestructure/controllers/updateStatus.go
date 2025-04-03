package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"iot/historial/infraestructure"

	"github.com/gin-gonic/gin"
)

type UpdateStatusCtrl struct {
	repo *infraestructure.MySQLRepository
}

func NewUpdateStatusCtrl(repo *infraestructure.MySQLRepository) *UpdateStatusCtrl {
	return &UpdateStatusCtrl{repo: repo}
}

func (ctrl *UpdateStatusCtrl) Run(c *gin.Context) {
	idPedidoStr := c.Param("idPedido")
	idPedido, err := strconv.Atoi(idPedidoStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "idPedido inválido"})
		return
	}

	var request struct {
		NuevoStatus string `json:"nuevo_status"`
	}

	if err := json.NewDecoder(c.Request.Body).Decode(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error en el JSON"})
		return
	}

	err = ctrl.repo.UpdateStatus(idPedido, request.NuevoStatus)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Estado actualizado"})
}
