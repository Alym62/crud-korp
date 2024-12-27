package usecases

import (
	"github.com/Alym62/crud-korp/internal/models"
	"github.com/Alym62/crud-korp/internal/repository"
)

type ProductUseCase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repository repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		repository: repository,
	}
}

func (pu *ProductUseCase) GetProducts() ([]models.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUseCase) Create(name string, description string, price float64) (models.Product, error) {
	product, err := models.NewProduct(name, description, price)
	if err != nil {
		return models.Product{}, err
	}

	return pu.repository.Create(product)
}