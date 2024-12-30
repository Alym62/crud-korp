package models

import (
	"fmt"
	"time"
)

type Product struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Removed     bool      `json:"removed"`
}

func NewProduct(name string, description string, price float64) (*Product, error) {
	p := &Product{
		Name:        name,
		Description: description,
		Price:       price,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Removed:     false,
	}

	err := p.isValid()
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Product) isValid() error {
	if len(p.Name) < 5 {
		return fmt.Errorf("o nome não pode ter o tamanho menor que 5")
	}

	if len(p.Description) < 10 {
		return fmt.Errorf("a descrição não pode ter um tamanho menor que 10")
	}

	if p.Price <= 0 {
		return fmt.Errorf("o preço não pode ter um valor menor ou igual a zero")
	}

	return nil
}
