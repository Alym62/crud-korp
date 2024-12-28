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
		return fmt.Errorf("name must be at least 5 characters, got %d", len(p.Name))
	}

	if len(p.Description) < 10 {
		return fmt.Errorf("description must be at least 10 characters, got %d", len(p.Description))
	}

	if p.Price <= 0 {
		return fmt.Errorf("price must be greater than 0, got %f", p.Price)
	}

	return nil
}
