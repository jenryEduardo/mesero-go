package routes

import (
<<<<<<< HEAD
	"consumer2/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)


func SetUp(router *gin.Engine){

	routes:=router.Group("/consumer")

	routes.POST("/",controllers.Consumer)

=======
	"consumer-event/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	routes := router.Group("/consume")

	routes.POST("/", controllers.ConsumeRabbitMQ)
>>>>>>> 96f608b264a1a4cf90c5791ecdfd57f89daa857c
}