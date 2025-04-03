package main

import (
	"log"
	Routes "publisher/infraestructure/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Configurar CORS
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"*"},
		AllowCredentials: true,
	}))

	Routes.Routes(r)

	port := ":3003"
	log.Println("Servidor escuchando en el puerto", port)
	log.Fatal(r.Run(port))
}