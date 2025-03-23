package main


import (
	"log"
	userRoutes "api-main/users/infraestructure/routes"
	mesasRoutes "api-main/mesa/infraestructure/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func main(){
	router := gin.Default()

	// Configurar CORS
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"*"},
		AllowCredentials: true,
	}))

	userRoutes.SetupRoutesCount(router)
	mesasRoutes.SetUpRoutes(router)

	port := ":8080"
	log.Println("Servidor escuchando en el puerto", port)
	log.Fatal(router.Run(port))

		
}