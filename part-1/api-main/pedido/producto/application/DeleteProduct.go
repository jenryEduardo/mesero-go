package application

import "api-main/pedido/producto/domain"

type DeleteProduct struct {
	repo domain.Iproducto
}

func NewDeleteProduct(repo domain.Iproducto)*DeleteProduct{
	return &DeleteProduct{repo: repo}
}

func(c *DeleteProduct)Execute(id int)error{
	return c.repo.Delete(id)
}