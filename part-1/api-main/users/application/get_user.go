package application

import(
	"api-main/users/domain"
)

type GetUserById struct{
	repo domain.Iuser
}

func NewGetUserById(repo domain.Iuser)*GetUserById{
	return &GetUserById{repo: repo}
}

func (c *GetUserById)Execute(id int)([]domain.User,error){
	return c.repo.GetUserById(id)
}



