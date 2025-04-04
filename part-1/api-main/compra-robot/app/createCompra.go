package app

import "api-main/compra-robot/domain"

type CreateCompra struct {
	repo domain.ICompra
}


func NewCompra(repo domain.ICompra)*CreateCompra{
	return &CreateCompra{repo: repo}
}

func(c *CreateCompra)Execute(compra *domain.Compra)error{
	return c.repo.Save(compra)
}