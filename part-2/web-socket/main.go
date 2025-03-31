package main

import (
    "log"
    "net/http"
    "websocket/app"
    "websocket/infraestructure/controllers"
    "websocket/infraestructure/routes"
    "websocket/infraestructure/adapters"
    "websocket/domain"
)

func main() {
    // Configurar WebSocket Server
    wsServer := adapters.NewWebSocketServer()
    go wsServer.Run()

    // Crear el controlador WebSocket
    wsController := controllers.NewWebSocketController(wsServer)

    // Inicializar las rutas
    routes.InitializeRoutes(wsController)

    // Simular la recepción de un pedido desde una API
    pedido := domain.Pedido{ID: "12345"}

    // Crear el servicio de pedidos
    pedidoService := app.NewPedidoService(wsController)

    // Enviar el pedido a través de WebSocket
    if err := pedidoService.EnviarPedido(pedido); err != nil {
        log.Fatalf("Error al enviar pedido: %v", err)
    }

    log.Println("Servidor WebSocket iniciado en :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
