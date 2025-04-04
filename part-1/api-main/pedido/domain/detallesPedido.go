package domain

type DetallesPedido struct {
	IdProducto     int
	NombreProducto string
	Cantidad       int
	PrecioUnitario float64
	Subtotal       float64
}