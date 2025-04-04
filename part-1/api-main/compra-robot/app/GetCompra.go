package app

import "api-main/compra-robot/domain"

type GetCompras struct {
	repo domain.ICompra
}


func NewGetCompras(repo domain.ICompra)*GetCompras{
	return &GetCompras{repo: repo}
}


func (c *GetCompras)Execute()([]domain.Compra,error){
	return c.repo.GetAll()
}