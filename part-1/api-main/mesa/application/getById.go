package application

import "api-main/mesa/domain"

type GetById struct {
	repo domain.Imesa
}

func NewGetById(repo domain.Imesa)*GetById{
	return &GetById{repo: repo}
}


func(c *GetById)Execute(id int)([]domain.Mesa,error){
	return c.repo.GetById(id)
}

