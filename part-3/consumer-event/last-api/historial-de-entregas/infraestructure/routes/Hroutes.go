package routes

import (
	"last/historial-de-entregas/infraestructure/dependencies"

	"github.com/gin-gonic/gin"
)

func HistorialRoutes(router *gin.Engine) {
	routes := router.Group("/historial")

	SaveH := dependencies.SaveHistorial()
	GetAllH := dependencies.GetAllHistorial()
	GetById := dependencies.GetHByID()
	UpdateH := dependencies.UpdateHistorial()
	DeleteH := dependencies.DeleteHistorial()

	routes.POST("/", SaveH.Run)
	routes.GET("/", GetAllH.Run)
	routes.GET("/:id_historial", GetById.Run)
	routes.PUT("/:id_historial", UpdateH.Run)
	routes.DELETE("/id_historial", DeleteH.Run)
}