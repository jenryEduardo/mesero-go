package routes

import (
	"api-main/detalles-del-pedido/infraestructure/dependencies"

	"github.com/gin-gonic/gin"
)

func Detallesroutes(router *gin.Engine) {
	routes := router.Group("/detalles")

	save := dependencies.SaveDetalle()
	getAll := dependencies.GetAllDetalles()
	getByID := dependencies.GetByIdDetalles()
	delete := dependencies.DeleteDetalles()
	update := dependencies.UpdateDetalles()

	routes.POST("/", save.Run)
	routes.GET("/", getAll.Run)
	routes.GET("/:idDetallePedido", getByID.Run)
	routes.PUT("/:idDetallePedido", update.Run)
	routes.DELETE("/:idDetallePedido", delete.Run)
}