package domain

type IPedido interface {
	PublishPedido(pedido Pedido) error
}