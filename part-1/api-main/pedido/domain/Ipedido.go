package domain

type Ipedido interface{
	Save(pedido *Pedido)(int64,error)
	Update(id int,pedido *Pedido)error
	Delete(id int)error
	ObtenerTotalPedido(id int)(float64, error)
	GetAll()([]Pedido,error)
	GetById(id int)([]Pedido,error)
	AgregarNuevoProducto(idPedido int,pedido *DetallesPedido)error
}