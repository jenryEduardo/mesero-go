package domain


// Pedido define la estructura del pedido que se enviará por WebSocket.
type Pedido struct {
    ID string `json:"idPedido"`
}
