package application

import "iot/historial/domain"

type GetStatus struct {
	repo domain.IHistorial
}

func NewGetStatus(repo domain.IHistorial) *GetStatus {
	return &GetStatus{repo:repo}
}

func (c *GetStatus) Run() error {
	return c.repo.GetStatus()
}