package main

import (
	mesasRoutes "api-main/mesa/infraestructure/routes"
	userRoutes "api-main/users/infraestructure/routes"
	"log"

	"github.com/gin-gonic/gin"
)


func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("CORS middleware ACTIVADO para:", c.Request.Method, c.Request.URL.Path)

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, Accept, Origin")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			log.Println("Handling OPTIONS request")
			c.AbortWithStatus(204) 
			return
		}

		c.Next()
	}
}


func main() {
	router := gin.New() 
	router.Use(gin.Recovery()) 

	router.Use(CORSMiddleware())

	userRoutes.SetupRoutesCount(router)
	mesasRoutes.SetUpRoutes(router)

	port := ":8080"
	log.Println("Servidor escuchando en el puerto", port)
	log.Fatal(router.Run(port))
}


