package application

import "api-main/users/domain"

type PutUser struct {
	repo domain.Iuser
}

func NewUpdateUser(repo domain.Iuser)*PutUser{
	return &PutUser{repo: repo}
}


func(c *PutUser)Execute(id int,user domain.User)error{
	return c.repo.UpdateUser(id,&user)
}