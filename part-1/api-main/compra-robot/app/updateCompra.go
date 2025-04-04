package app

import "api-main/compra-robot/domain"

type UpdateCompra struct {
	repo domain.ICompra
}


func NewUpdateCompra(repo domain.ICompra)*UpdateCompra{
	return &UpdateCompra{repo: repo}
}


func (c *UpdateCompra)Execute(id int,compra *domain.Compra)error{
	return c.repo.Update(id,compra)
}