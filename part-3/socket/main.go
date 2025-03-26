package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

// Mapa para almacenar clientes conectados
var clients = make(map[*websocket.Conn]bool)

// Canal para transmitir mensajes a todos los clientes
var broadcast = make(chan string)

// Upgrader para convertir una solicitud HTTP en WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return r.Host == "localhost:4000" // Permite solo conexiones desde localhost
	},
}

// Maneja nuevas conexiones de clientes
func HandlerConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error al actualizar a WebSocket:", err)
		return
	}
	defer conn.Close()

	// Registrar cliente
	clients[conn] = true
	fmt.Println("Nuevo cliente conectado!")

	// Escuchar mensajes del cliente
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Cliente desconectado:", err)
			delete(clients, conn) // Eliminar cliente desconectado
			break
		}

		// Enviar mensaje recibido al canal broadcast
		broadcast <- string(msg)
	}
}

// Maneja la transmisión de mensajes a todos los clientes
func HandleMessages() {
	for {

		msg := <-broadcast // Espera mensajes en el canal
		// Enviar mensaje a todos los clientes conectados
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				log.Println("Error enviando mensaje:", err)
				client.Close()
				delete(clients, client) // Eliminar cliente desconectado
			}
		}
	}
}

func main() {
	// Ruta WebSocket
	http.HandleFunc("/ws", HandlerConnections)

	// Goroutine para manejar mensajes y enviarlos a todos los clientes
	go HandleMessages()

	fmt.Println("Servidor WebSocket corriendo en ws://localhost:8080/ws")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
