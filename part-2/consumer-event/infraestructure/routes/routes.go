package routes

import (
	"consumer-event/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	routes := router.Group("/consume")

	routes.POST("/", controllers.ConsumeRabbitMQ)

}