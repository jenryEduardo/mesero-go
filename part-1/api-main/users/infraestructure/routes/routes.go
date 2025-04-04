package routes


import (
	"github.com/gin-gonic/gin"
	"api-main/users/infraestructure/controllers"
)


func SetupRoutesCount(router *gin.Engine) {

	routes:=router.Group("/usuarios")

	{
		routes.POST("/", controllers.CreateUser)
		routes.GET("/:id", controllers.GetUsers)
		routes.GET("/", controllers.GetuserAll)
		routes.PUT("/actualizar/:id", controllers.UpdateUser)
		routes.DELETE("/eliminar-usuario/:id",controllers.DeleteUser)
		routes.POST("/login",controllers.Login)	
		}	
}	