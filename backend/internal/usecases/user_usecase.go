package usecases

import (
	"github.com/Alym62/crud-korp/internal/models"
	"github.com/Alym62/crud-korp/internal/repositories"
)

type UserUseCase struct {
	repository repositories.UserRepository
}

func NewUserUseCase(repository repositories.UserRepository) UserUseCase {
	return UserUseCase{
		repository: repository,
	}
}

func (uc *UserUseCase) GetById(id uint) (*models.User, error) {
	return uc.repository.GetById(id)
}

func (uc *UserUseCase) Create(email string, password string, position string, role models.Role) (models.User, error) {
	user, err := models.NewUser(email, password, position, role)
	if err != nil {
		return models.User{}, err
	}

	return uc.repository.Create(user)
}

func (uc *UserUseCase) DeleteById(id uint) (*models.User, error) {
	return uc.repository.DeleteById(id)
}

func (uc *UserUseCase) Update(id uint, email string, password string, position string, role models.Role) (*models.User, error) {
	user, err := models.NewUser(email, password, position, role)
	if err != nil {
		return nil, err
	}

	return uc.repository.Update(id, user)
}
