package application

import (
	"api-main/users/domain"
	"fmt"
)

type GetUserByEmail struct {
	repo domain.Iuser
}

func NewGetUserByEmail(repo domain.Iuser) *GetUserByEmail {
	return &GetUserByEmail{repo: repo}
}

// Execute obtiene un usuario por email
func (g *GetUserByEmail) Execute(email string) (*domain.User, error) {
	// Llamar al repositorio para obtener el usuario por email
	user, err := g.repo.GetUserByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("usuario no encontrado: %w", err)
	}

	return user, nil
}
