package application

import "api-main/detalles-del-pedido/domain"

type UpdateDetalle struct {
	repo domain.IDetallePedido
}

func NewUpdateDetalles(repo domain.IDetallePedido) *UpdateDetalle {
	return &UpdateDetalle{repo:repo}
}

func(c *UpdateDetalle) Run(idDetallePedido int, detalles domain.DetallesPedido) error {
	return c.repo.Update(idDetallePedido, detalles)
}