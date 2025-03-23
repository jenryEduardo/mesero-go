package application

import "api-main/pedido/domain"

type DeletePedido struct {
	repo domain.Ipedido
}


func NewDelete(repo domain.Ipedido)*DeletePedido{
	return &DeletePedido{repo: repo}
}

func (c *DeletePedido)Execute(id int)error{
	return c.repo.Delete(id)
}