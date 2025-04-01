package controllers

import (
	"api-main/detalles-del-pedido/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteDetallesCtrl struct {
	uc *application.DeleteDetalles
}

func NewDeleteRSCtrl(uc *application.DeleteDetalles) *DeleteDetallesCtrl {
	return &DeleteDetallesCtrl{uc: uc}
}

func (ctrl *DeleteDetallesCtrl) Run(c *gin.Context) {
	id := c.Param("idDetallePedido")
	idDetallePedido, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El id debe ser númerico"})
		return
	}

	err = ctrl.uc.Run(idDetallePedido)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": true,
			"data": "Detalles eliminados",
		})
	}
}