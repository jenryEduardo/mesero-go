package domain


type Pedido struct{
	IdPedido int
	IdMesa int
	Nombre_cliente string
	Status string
	Total float64
	Detalles []DetallesPedido
}