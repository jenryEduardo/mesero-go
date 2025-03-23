package application

import "api-main/pedido/domain"

type GetByIdPedido struct {
	repo domain.Ipedido
}

func NewGetById(repo domain.Ipedido)*GetByIdPedido{
	return &GetByIdPedido{repo: repo}
}


func (c *GetByIdPedido)Execute(id int)([]domain.Pedido,error){
	return c.repo.GetById(id)
}