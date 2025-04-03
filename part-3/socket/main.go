package main

import (
	"fmt"
	"log"
	"net/http"
	"socket/infrastructure/routes"
)

func main() {
	// Registrar las rutas
	infrastructure.RegisterRoutes()

	// Iniciar servidor
	fmt.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
