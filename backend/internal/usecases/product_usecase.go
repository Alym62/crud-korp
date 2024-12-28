package usecases

import (
	"github.com/Alym62/crud-korp/internal/models"
	"github.com/Alym62/crud-korp/internal/repositories"
	"github.com/Alym62/crud-korp/pkg"
)

type ProductUseCase struct {
	repository repositories.ProductRepository
}

func NewProductUseCase(repository repositories.ProductRepository) ProductUseCase {
	return ProductUseCase{
		repository: repository,
	}
}

func (pu *ProductUseCase) GetAllByPage(page int, limit int) (pkg.PageResponse[models.Product], error) {
	return pu.repository.GetAllByPage(page, limit)
}

func (pu *ProductUseCase) GetList() ([]models.Product, error) {
	return pu.repository.GetList()
}

func (pu *ProductUseCase) GetById(id uint) (*models.Product, error) {
	return pu.repository.GetById(id)
}

func (pu *ProductUseCase) Create(name string, description string, price float64) (models.Product, error) {
	product, err := models.NewProduct(name, description, price)
	if err != nil {
		return models.Product{}, err
	}

	return pu.repository.Create(product)
}

func (pu *ProductUseCase) DeleteById(id uint) (*models.Product, error) {
	return pu.repository.DeleteById(id)
}

func (pu *ProductUseCase) Update(id uint, name string, description string, price float64) (*models.Product, error) {
	product, err := models.NewProduct(name, description, price)
	if err != nil {
		return nil, err
	}

	return pu.repository.Update(id, product)
}
