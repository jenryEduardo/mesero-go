package routes

import (
	"publisher/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	routes := r.Group("/publisherStatus")
	routes.POST("/", controllers.PublisherInRabbit)
}