package usecases

import (
	"github.com/Alym62/crud-korp/internal/models"
	"github.com/Alym62/crud-korp/internal/repositories"
)

type AuthUseCase struct {
	repository repositories.AuthRepository
}

func NewAuthUseCase(repository repositories.AuthRepository) AuthUseCase {
	return AuthUseCase{
		repository: repository,
	}
}

func (au *AuthUseCase) Login(username string) (*models.User, error) {
	return au.repository.Login(username)
}
