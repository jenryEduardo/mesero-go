package application

import "api-main/mesa/domain"

type UpdateMesa struct {
	repo domain.Imesa
}

func NewUpdateMesa(repo domain.Imesa)*UpdateMesa{
	return &UpdateMesa{repo: repo}
}


func(c *UpdateMesa)Execute(id int,mesa *domain.Mesa)error{
	return c.repo.Update(id,*mesa)
}