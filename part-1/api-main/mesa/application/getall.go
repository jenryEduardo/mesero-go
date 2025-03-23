package application

import "api-main/mesa/domain"

type GetAllMesa struct {
	repo domain.Imesa
}

func NewGetAllMesas(repo domain.Imesa)*GetAllMesa{
	return &GetAllMesa{repo: repo}
}

func (c *GetAllMesa)Execute()([]domain.Mesa,error){
	return c.repo.GetAll()
}