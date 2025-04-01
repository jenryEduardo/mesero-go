package application

import "api-main/detalles-del-pedido/domain"

type GetAllDetalles struct {
	repo domain.IDetallePedido
}

func NewGetAllDetalles(repo domain.IDetallePedido) *GetAllDetalles {
	return &GetAllDetalles{repo:repo}
}

func (c *GetAllDetalles) Run() ([]domain.DetallesPedido, error) {
	return c.repo.GetAll()
}