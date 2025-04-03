package application

import "iot/historial/domain"

type FindColor struct {
	repo domain.IHistorial
}

func NewFindColor(repo domain.IHistorial) *FindColor {
	return &FindColor{repo: repo}
}

func (c *FindColor) Run(idPedido int) (int, error) {
	return c.repo.FindColor(idPedido)
}