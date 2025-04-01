package controllers

import (
	"api-main/detalles-del-pedido/application"
	"api-main/detalles-del-pedido/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SaveDetallesCtrl struct {
	uc *application.SaveDetalles
}

func NewSaveDetallesCtrl(uc *application.SaveDetalles) *SaveDetallesCtrl {
	return &SaveDetallesCtrl{uc: uc}
}

func (ctrl *SaveDetallesCtrl) Run(c *gin.Context) {
	var detalles domain.DetallesPedido

	if err := c.ShouldBindJSON(&detalles); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	err := ctrl.uc.Run(detalles)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"status": true,
			"data": gin.H{
				"type": "Detalles del pedido",
				"ID de los detalles": detalles.IdDetallePedido,
				"attributes": gin.H{
					"ID del pedido": detalles.IdPedido,
					"ID del producto": detalles.IdProducto,
					"Cantidad": detalles.Cantidad,
					"Precio unitario": detalles.PrecioUnitario,
					"Subtotal": detalles.Subtotal,
				},
			},
		})
	}
}