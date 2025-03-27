package application

import "last-api/historial-de-entregas/domain"

type GetAllHistorial struct {
	repo domain.IHistorial
}

func NewGetAllHistorial(repo domain.IHistorial) *GetAllHistorial {
	return &GetAllHistorial{repo:repo}
}

func (c *GetAllHistorial) Run() ([]domain.Historial, error) {
	return c.repo.GetAll()
}