package main

import (
<<<<<<< HEAD
	"log"

	countsRoutes "consumer2/infraestructure/routes"
=======
	consumerRoutes "consumer-event/infraestructure/routes"
	"log"

>>>>>>> 96f608b264a1a4cf90c5791ecdfd57f89daa857c
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

<<<<<<< HEAD

func main(){
=======
func main() {
>>>>>>> 96f608b264a1a4cf90c5791ecdfd57f89daa857c
	router := gin.Default()

	// Configurar CORS
	router.Use(cors.New(cors.Config{
<<<<<<< HEAD
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"*"},
		AllowCredentials: true,
	}))

	countsRoutes.SetUp(router)
	

	port := ":3002"
	log.Println("Servidor escuchando en el puerto", port)
	log.Fatal(router.Run(port))

		
=======
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	consumerRoutes.Routes(router)

	port := ":3000"
	log.Println("Servidor escuchando en el puerto", port)
	log.Fatal(router.Run(port))
>>>>>>> 96f608b264a1a4cf90c5791ecdfd57f89daa857c
}