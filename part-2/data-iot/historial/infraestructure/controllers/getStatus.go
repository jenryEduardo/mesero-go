package controllers

import (
	"log"
	"net/http"

	"iot/historial/infraestructure"

	"github.com/gin-gonic/gin"
)

type GetStatusCtrl struct {
	repo *infraestructure.MySQLRepository
}

func NewGetStatusCtrl(repo *infraestructure.MySQLRepository) *GetStatusCtrl {
	return &GetStatusCtrl{repo: repo}
}

func (ctrl *GetStatusCtrl) Run(c *gin.Context) {
	go func() {
		err := ctrl.repo.GetStatus()
		if err != nil {
			log.Println("Error al escuchar MQTT:", err)
		}
	}()

	c.JSON(http.StatusOK, gin.H{"mensaje": "Escuchando status desde MQTT"})
}
