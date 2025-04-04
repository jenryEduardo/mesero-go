package routes

import (
	"github.com/gin-gonic/gin"
	"api-main/pedido/infraestructure/controllers"
)

func SetUpRoutes(routes *gin.Engine){
	router:=routes.Group("/pedidos")




	router.POST("/",controllers.CreatePedido)
	router.POST("/:idPedido",controllers.AddProductInPedido)
	router.GET("/",controllers.GetAllPedido)
	router.GET("/:idPedido",controllers.GetByIdPedido)
	router.DELETE("/:idPedido",controllers.DeletePedido)
	router.PUT("/:idPedido",controllers.UpdatePedido)
}







