package application

import "api-main/detalles-del-pedido/domain"

type DeleteDetalles struct {
	repo domain.IDetallePedido
}

func NewDeleteDetalles(repo domain.IDetallePedido) *DeleteDetalles {
	return &DeleteDetalles{repo:repo}
}

func(c *DeleteDetalles) Run(idDetallePedido int) error {
	return c.repo.Delete(idDetallePedido)
}