package application

import "api-main/users/domain"

type PostUser struct {
	repo domain.Iuser
}


func NewPostUser(repo domain.Iuser)*PostUser{
	return &PostUser{repo: repo}
}

func(c *PostUser)Execute(user domain.User)error{
	return c.repo.Save(&user)
}