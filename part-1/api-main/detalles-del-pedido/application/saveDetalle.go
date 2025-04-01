package application

import "api-main/detalles-del-pedido/domain"

type SaveDetalles struct {
	repo domain.IDetallePedido
}

func NewSaveDetalles(repo domain.IDetallePedido) *SaveDetalles {
	return &SaveDetalles{repo:repo}
}

func (c *SaveDetalles) Run(detalles domain.DetallesPedido) error {
	return c.repo.Save(detalles)
}