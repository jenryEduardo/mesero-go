package application

import "api-main/pedido/producto/domain"

type GetByIdOProducts struct {
	repo domain.Iproducto
}

func NewGetByIdProducts(repo domain.Iproducto)*GetByIdOProducts{
	return &GetByIdOProducts{repo: repo}
}

func(c *GetByIdOProducts)Execute(id int)([]domain.Producto,error){
	return c.repo.GetById(id)
}