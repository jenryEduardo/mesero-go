package application

import "api-main/pedido/producto/domain"

type CreateProduct struct {
	repo domain.Iproducto
}


func NewProducto(repo domain.Iproducto)*CreateProduct{
	return &CreateProduct{repo: repo}
}

func(c  *CreateProduct)Execute(producto *domain.Producto)error{
	return c.repo.Save(producto)
}