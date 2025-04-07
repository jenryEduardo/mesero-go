package main

import (
	"log"
	"consumer-event/application"
	"consumer-event/infraestructure/adapters"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Configurar CORS
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	// 👉 Inicializar el consumidor aquí directamente al iniciar la app
	go func() {
		repo, err := adapters.NewRabbitMQRepository()
		if err != nil {
			log.Fatalf("❌ Error al conectarse a RabbitMQ: %v", err)
		}

		useCase := application.NewConsume(repo)

		err = useCase.Execute()
		if err != nil {
			log.Printf("❌ Error ejecutando el consumidor: %v", err)
		}
	}()


	// Iniciar servidor
	port := ":3003"
	log.Println("Servidor escuchando en el puerto", port)
	log.Fatal(router.Run(port))
}
