package application

import "api-main/pedido/domain"

type ObtenerTotalPedido struct {
	repo domain.Ipedido
}

func NewObtenerTotal(repo domain.Ipedido)*ObtenerTotalPedido{
	return &ObtenerTotalPedido{repo: repo}
}

func(c *ObtenerTotalPedido)Execute(id int)(float64,error){
	return c.repo.ObtenerTotalPedido(id)
}