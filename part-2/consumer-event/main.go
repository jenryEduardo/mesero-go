package main

import (
	"log"

	countsRoutes "consumer2/infraestructure/routes"
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

	countsRoutes.SetUp(router)
	

	port := ":3002"
	log.Println("Servidor escuchando en el puerto", port)
	log.Fatal(router.Run(port))

		
}