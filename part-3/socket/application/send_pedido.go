package application

import (
	"encoding/json"
	"log"
	"socket/domain"
	"socket/infrastructure/adapters"
)

type PedidoService struct{}

func (p PedidoService) SendPedido(pedido domain.Pedido) {
	pedidoJSON, err := json.Marshal(pedido)
	if err != nil {
		log.Println("Error al convertir pedido a JSON:", err)
		return
	}
	adapters.Broadcast <- string(pedidoJSON)
}


// SendPedido se encarga de enviar el pedido a WebSocket
func SendPedido(pedido domain.Pedido) {
	// Convertimos el pedido en JSON
	pedidoJSON, err := json.Marshal(pedido)
	if err != nil {
		log.Println("Error al convertir pedido a JSON:", err)
		return
	}

	// Enviar al canal WebSocket
	adapters.Broadcast <- string(pedidoJSON)
}
