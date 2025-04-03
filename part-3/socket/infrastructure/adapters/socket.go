package adapters


import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

// Mapa de clientes conectados
var Clients = make(map[*websocket.Conn]bool)

// Canal de mensajes
var Broadcast = make(chan string)

// Configuración del WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Handler para WebSocket
func HandlerConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error al actualizar a WebSocket:", err)
		return
	}
	defer conn.Close()

	// Registrar cliente
	Clients[conn] = true
	fmt.Println("Nuevo cliente conectado!")

	// Escuchar mensajes
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Println("Cliente desconectado:", err)
			delete(Clients, conn)
			break
		}
	}
}

// Función para enviar mensajes a los clientes conectados
func HandleMessages() {
	for {
		msg := <-Broadcast
		for client := range Clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				log.Println("Error enviando mensaje:", err)
				client.Close()
				delete(Clients, client)
			}
		}
	}
}
