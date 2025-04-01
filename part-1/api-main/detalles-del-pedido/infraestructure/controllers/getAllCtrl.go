package controllers

import (
	"api-main/detalles-del-pedido/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllDetallesCtrl struct {
	uc *application.GetAllDetalles
}

func NewGetAllDetallesCtrl(uc *application.GetAllDetalles) *GetAllDetallesCtrl {
	return &GetAllDetallesCtrl{uc: uc}
}

func (ctrl *GetAllDetallesCtrl) Run(c *gin.Context) {
	rs, err := ctrl.uc.Run()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": true,
			"data": rs,
		})
	}
}