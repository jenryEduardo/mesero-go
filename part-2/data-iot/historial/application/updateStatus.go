package application

import "iot/historial/domain"

type UpdateStatus struct {
	repo domain.IHistorial
}

func NewUpdateStatus(repo domain.IHistorial) *UpdateStatus {
	return &UpdateStatus{repo:repo}
}

func(c *UpdateStatus) Run(idPedido int, nuevoStatus string) error {
	return c.repo.UpdateStatus(idPedido, nuevoStatus)
}