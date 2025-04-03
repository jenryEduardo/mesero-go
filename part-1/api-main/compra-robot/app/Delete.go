package app

import "api-main/compra-robot/domain"

type Delete struct {
	repo domain.ICompra
}


func NewDelete(repo domain.ICompra)*Delete{
	return &Delete{repo: repo}
}


func (c *Delete)Execute(id int)error{
	return c.repo.Delete(id)
}