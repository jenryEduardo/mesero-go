package routes

import (
	"github.com/gin-gonic/gin"
	"api-main/pedido/producto/infraestructure/controllers"
)

func SetUpRoutes(routes *gin.Engine){
	router:=routes.Group("/productos")

	router.POST("/",controllers.CreateProduct)
	router.GET("/",controllers.GetAllProducts)
	router.GET("/:idProducto",controllers.GetByIdOProducts)
	router.DELETE("/:idProducto",controllers.DeleteProduct)
	router.PUT("/:idProducto",controllers.UpdateProduct)
}