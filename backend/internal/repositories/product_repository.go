package repositories

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Alym62/crud-korp/internal/models"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetList() ([]models.Product, error) {
	query := "SELECT id, name, description, price, created_at, updated_at, removed FROM product WHERE removed = false"
	rows, err := pr.connection.Query(query)

	if err != nil {
		fmt.Println(err)
		return []models.Product{}, err
	}

	var productList []models.Product
	var product models.Product

	for rows.Next() {
		err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.CreatedAt,
			&product.UpdatedAt,
			&product.Removed,
		)

		if err != nil {
			fmt.Println(err)
			return []models.Product{}, err
		}

		productList = append(productList, product)
	}

	rows.Close()

	return productList, nil
}

func (pr *ProductRepository) Create(product *models.Product) (models.Product, error) {
	var p models.Product

	query, err := pr.connection.Prepare("INSERT INTO product (name, description, price, created_at, updated_at, removed)" +
		"VALUES($1, $2, $3, $4, $5, $6) RETURNING id, name, description, price, created_at, updated_at, removed")

	if err != nil {
		return models.Product{}, err
	}

	err = query.QueryRow(
		product.Name,
		product.Description,
		product.Price,
		product.CreatedAt,
		product.UpdatedAt,
		product.Removed,
	).Scan(
		&p.ID,
		&p.Name,
		&p.Description,
		&p.Price,
		&p.CreatedAt,
		&p.UpdatedAt,
		&p.Removed,
	)
	if err != nil {
		return models.Product{}, err
	}

	query.Close()

	return p, nil
}

func (pr *ProductRepository) GetById(id uint) (*models.Product, error) {
	var p models.Product

	query, err := pr.connection.Prepare(
		"SELECT id, name, description, price, created_at, updated_at, removed FROM product " +
			"WHERE removed = false AND id = $1")

	if err != nil {
		return nil, err
	}

	err = query.QueryRow(id).Scan(
		&p.ID,
		&p.Name,
		&p.Description,
		&p.Price,
		&p.CreatedAt,
		&p.UpdatedAt,
		&p.Removed,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()

	return &p, nil
}

func (pr *ProductRepository) DeleteById(id uint) (*models.Product, error) {
	var p models.Product

	query, err := pr.connection.Prepare(
		"UPDATE product SET removed = true, updated_at = $1 " +
			"WHERE removed = false AND id = $2 RETURNING id, name, description, price, created_at, updated_at, removed")

	if err != nil {
		return nil, err
	}

	err = query.QueryRow(time.Now(), id).Scan(
		&p.ID,
		&p.Name,
		&p.Description,
		&p.Price,
		&p.CreatedAt,
		&p.UpdatedAt,
		&p.Removed,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()

	return &p, nil
}

func (pr *ProductRepository) Update(id uint, product *models.Product) (*models.Product, error) {
	var p models.Product

	query, err := pr.connection.Prepare(
		"UPDATE product SET name = $1, description = $2, price = $3, updated_at = $4 " +
			"WHERE removed = false AND id = $5 RETURNING id, name, description, price, created_at, updated_at, removed")

	if err != nil {
		return nil, err
	}

	err = query.QueryRow(product.Name, product.Description, product.Price, time.Now(), id).Scan(
		&p.ID,
		&p.Name,
		&p.Description,
		&p.Price,
		&p.CreatedAt,
		&p.UpdatedAt,
		&p.Removed,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()

	return &p, nil
}
