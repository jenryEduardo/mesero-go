package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"socket/domain"
)

var pedidoService domain.PedidoSender

func SetPedidoService(service domain.PedidoSender) {
	pedidoService = service
}

func PedidoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var pedido domain.Pedido
	err := json.NewDecoder(r.Body).Decode(&pedido)
	if err != nil {
		http.Error(w, "Error al decodificar JSON", http.StatusBadRequest)
		return
	}

	// Usar la interfaz en lugar de una dependencia directa
	pedidoService.SendPedido(pedido)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Pedido enviado correctamente")
}
