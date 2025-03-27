package application

import "api-main/detalles-del-pedido/domain"

type GetDetalleByID struct {
	repo domain.IDetallePedido
}

func NewGetDetalleByID(repo domain.IDetallePedido) *GetDetalleByID {
	return &GetDetalleByID{repo:repo}
}

func (c *GetDetalleByID) Run(idDetallePedido int) ([]domain.DetallesPedido, error) {
	return c.repo.GetById(idDetallePedido)
}