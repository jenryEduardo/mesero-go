package adapters


import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"socket/domain"
	"socket/infrastructure/controllers"
	"github.com/gorilla/websocket"
)

// Mapa para clientes que envían status
var ClientsStatus = make(map[*websocket.Conn]bool)

// Upgrader para WebSocket
var upgraderStatus = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Handler para recibir status de clientes
func HandlerStatusConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgraderStatus.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error al actualizar a WebSocket:", err)
		return
	}
	defer conn.Close()

	ClientsStatus[conn] = true
	fmt.Println("Cliente conectado para enviar status!")

	// Escuchar mensajes de status
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Cliente desconectado:", err)
			delete(ClientsStatus, conn)
			break
		}

		// Convertir mensaje JSON en estructura
		var statusMessage domain.Pedido
		err = json.Unmarshal(msg, &statusMessage)
		if err != nil {
			log.Println("Error al decodificar mensaje de status:", err)
			continue
		}

		// Enviar status al controlador de la API
		controllers.PedidoHandler(w,r)
	}
}
