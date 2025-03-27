package application

import "api-main/users/domain"

type GetAllUser struct {
	repo domain.Iuser
}

func NewGetAllUser(repo domain.Iuser)*GetAllUser{
	return &GetAllUser{repo: repo}
}


func(c *GetAllUser)Execute()([]domain.User,error){
	return c.repo.GetAllUser()
}