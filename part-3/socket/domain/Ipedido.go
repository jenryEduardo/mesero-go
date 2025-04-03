package domain


type PedidoSender interface {
	SendPedido(pedido Pedido)
}
