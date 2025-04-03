package app

import "api-main/compra-robot/domain"

type GetByIdCompra struct {
	repo domain.ICompra
}

func NewGetById(repo domain.ICompra)*GetByIdCompra{
	return &GetByIdCompra{repo: repo}
}

func (c *GetByIdCompra)Execute(id int)([]domain.Compra,error){
	return c.repo.GetById(id)
}