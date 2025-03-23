package application

import "api-main/mesa/domain"

type DeleteMesa struct {
	repo domain.Imesa
}

func NewDelete(repo domain.Imesa)*DeleteMesa{
	return &DeleteMesa{repo: repo}
}

func (c *DeleteMesa)Execute(id int)error{
	return c.repo.Delete(id)
}