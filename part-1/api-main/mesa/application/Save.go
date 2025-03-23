package application

import "api-main/mesa/domain"

type CreateMesa struct {
	repo domain.Imesa
}

func NewCreateMesa(repo domain.Imesa)*CreateMesa{
	return &CreateMesa{repo: repo}
}


func(c *CreateMesa)Execute(mesa *domain.Mesa)error{
	return c.repo.Save(mesa)
}

