package routes

import (
	"consumer2/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)


func SetUp(router *gin.Engine){

	routes:=router.Group("/consumer")

	routes.POST("/",controllers.Consumer)

}