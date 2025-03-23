package application

import "api-main/pedido/domain"

type GetAllPedido struct {
	repo domain.Ipedido
}

func NewGetAllPedidos(repo domain.Ipedido)*GetAllPedido{
	return &GetAllPedido{repo: repo}
}

func(c *GetAllPedido)Execute()([]domain.Pedido,error){
	return c.repo.GetAll()
}