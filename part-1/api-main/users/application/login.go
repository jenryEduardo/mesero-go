package application

import "api-main/users/domain"

type LoginUser struct {
    repo domain.Iuser
}

func NewLoginUser(repo domain.Iuser) *LoginUser {
    return &LoginUser{repo: repo}
}

func (c *LoginUser) Execute(email string, password string) (*domain.User,string, error) {
    return c.repo.Login(email, password)  // Llama a Login desde el repositorio
}
