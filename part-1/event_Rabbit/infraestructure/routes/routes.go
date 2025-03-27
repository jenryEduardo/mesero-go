package routes

import (
	"consumer/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)


func SetUp(router *gin.Engine){

	routes:=router.Group("/consumer")

	routes.POST("/",controllers.SaveInRbbitmq)

}