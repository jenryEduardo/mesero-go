package main

import (
	mesasRoutes "api-main/mesa/infraestructure/routes"
	userRoutes "api-main/users/infraestructure/routes"
	productosRoutes "api-main/pedido/producto/infraestructure/routes"
	pedidosRoutes "api-main/pedido/infraestructure/routes"
	"github.com/gin-contrib/cors"
	"log"

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
	productosRoutes.SetUpRoutes(router)
	pedidosRoutes.SetUpRoutes(router)
	port := ":8080"
	log.Println("Servidor escuchando en el puerto", port)
	log.Fatal(router.Run(port))
}


