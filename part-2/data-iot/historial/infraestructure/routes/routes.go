package routes

import (
	"iot/historial/infraestructure/dependencies"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	routes := router.Group("/mqtt")

	postColor := dependencies.FindColor()

	routes.GET("/color", postColor.Run)
	
}