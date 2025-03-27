package controllers

import (
	"api-main/detalles-del-pedido/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetDetallesByIDCtrl struct {
	uc *application.GetDetalleByID
}

func NewGetRsIDCtrl(uc *application.GetDetalleByID) *GetDetallesByIDCtrl {
	return &GetDetallesByIDCtrl{uc: uc}
}

func (ctrl *GetDetallesByIDCtrl) Run(c *gin.Context) {
	id := c.Param("idDetallePedido")
	idDetallePedido, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	detalles, err := ctrl.uc.Run(idDetallePedido)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": true,
			"data": detalles,
		})
	}
}