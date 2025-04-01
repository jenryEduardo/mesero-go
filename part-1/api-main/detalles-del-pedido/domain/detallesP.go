package domain

type DetallesPedido struct {
	IdDetallePedido int 	`json:"idDetallePedido"`
	IdPedido int 			`json:"idPedido"`
	IdProducto int			`json:"idProducto"`
	Cantidad int			`json:"cantidad"`
	PrecioUnitario float32 	`json:"precio_unitario"`
	Subtotal float32		`json:"subtotal"`
}