package application

import "api-main/producto/domain"

type GetAllProducts struct {
	repo domain.Iproducto
}

func NewGetALLproducts(repo domain.Iproducto)*GetAllProducts{
	return &GetAllProducts{repo: repo}
}

func (c *GetAllProducts)Execute()([]domain.Producto,error){
	return c.repo.GetAll()
}