package infrastructure

import (
	"net/http"
	"socket/infrastructure/controllers"
	"socket/infrastructure/adapters"
)

func RegisterRoutes() {
	// Ruta API para pedidos
	http.HandleFunc("/api/pedido", controllers.PedidoHandler)

	// Ruta WebSocket
	http.HandleFunc("/ws", adapters.HandlerConnections)

	// Iniciar el manejador de WebSocket en una goroutine
	go adapters.HandleMessages()
}
