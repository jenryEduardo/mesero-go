package application

import "api-main/users/domain"

type DeleteUser struct {
	repo domain.Iuser
}

func NewDeleteUser(repo domain.Iuser)*DeleteUser{
	return &DeleteUser{repo: repo}
}

func(c *DeleteUser)Execute(id int)error{
	return c.repo.DeleteUser(id)
}