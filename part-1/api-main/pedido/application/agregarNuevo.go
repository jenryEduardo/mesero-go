package application

import "api-main/pedido/domain"

type AgregarNuevoProducto struct {
	repo domain.Ipedido
}

func NewAddProduct(repo domain.Ipedido)*AgregarNuevoProducto{
	return &AgregarNuevoProducto{repo: repo}
}


func(c *AgregarNuevoProducto)Execute(idPedido int,pedido *domain.DetallesPedido)error{
	return c.repo.AgregarNuevoProducto(idPedido,pedido)
}