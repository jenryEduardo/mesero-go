package domain

type DetallesPedido struct {
	IdProducto     int
	Cantidad       int
	PrecioUnitario float64
	Subtotal       float64
}