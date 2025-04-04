package application

import "api-main/pedido/domain"

type SavePedido struct {
	repo domain.Ipedido
}

func NewSavePedido(repo domain.Ipedido)*SavePedido{
	return &SavePedido{repo: repo}
}

func (c *SavePedido)Execute(pedido domain.Pedido)(int64,error){
	return c.repo.Save(&pedido)
}