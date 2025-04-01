package controllers

import (
	"api-main/detalles-del-pedido/application"
	"api-main/detalles-del-pedido/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateDetallesCtrl struct {
	uc *application.UpdateDetalle
}

func NewUpdateDetallesCtrl (uc *application.UpdateDetalle) *UpdateDetallesCtrl {
	return &UpdateDetallesCtrl{uc:uc}
}

func (ctrl *UpdateDetallesCtrl) Run(c *gin.Context) {
	var detalles domain.DetallesPedido 

	if err := c.ShouldBindJSON(&detalles); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error":"status es requerido"})
		return
	}

	id := c.Param("idEstado")
	idEstado, err := strconv.Atoi(id)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"El id debe ser númerico",
		})
		return
	}

	if err := ctrl.uc.Run(idEstado, detalles); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
	return
}