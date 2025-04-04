package application

import "api-main/pedido/producto/domain"

type UpdateProduct struct {
	repo domain.Iproducto
}


func NewUpdateProduct(repo domain.Iproducto)*UpdateProduct{
	return &UpdateProduct{repo: repo}
}

func (c *UpdateProduct)Execute(id int,product *domain.Producto)error{
	return c.repo.Update(id,product)
}
