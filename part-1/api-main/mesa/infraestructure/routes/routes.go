package routes

import (
	"github.com/gin-gonic/gin"
	"api-main/mesa/infraestructure/controllers"
)

func SetUpRoutes(routes *gin.Engine){ 


	router:=routes.Group("/mesa")

	router.POST("/",controllers.CreateMesa)
	router.PUT("/:idMesa",controllers.UpdateMesa)
	router.GET("/:idMesa",controllers.GetById)
	router.GET("/",controllers.GetAll)
	router.DELETE("/:idMesa",controllers.DeleteMesa)

}