package application

import "api-main/pedido/domain"

type UpdatePedido struct {
	repo domain.Ipedido
}

func NewUpdatePedido(repo domain.Ipedido)*UpdatePedido{
	return &UpdatePedido{repo: repo}
}


func(c *UpdatePedido)Execute(id int,pedido domain.Pedido)error{
	return c.repo.Update(id,&pedido)
}