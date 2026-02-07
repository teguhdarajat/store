package repositories

import (
	"database/sql"
	"errors"
	"store/models"
)

type ProductRepository interface {
	GetAll() ([]models.Product, error)
	GetByID(uint) (*models.Product, error)
	Insert(*models.Product) (*models.Product, error)
	Update(models.Product) error
	Delete(uint) error
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (pr *productRepository) GetAll() ([]models.Product, error) {
	query := "SELECT id, name, description FROM products"
	rows, err := pr.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := make([]models.Product, 0)
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Stock)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (pr *productRepository) GetByID(id uint) (*models.Product, error) {
	query := "SELECT id, name, description FROM products WHERE id= $1"

	var product models.Product
	err := pr.db.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Price, &product.Stock)
	if err == sql.ErrNoRows {
		return nil, errors.New("product not found")
	}

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (pr *productRepository) Insert(product *models.Product) (*models.Product, error) {
	query := "INSERT INTO products (name, price, stock) VALUES ($1, $2, $3) returning id"
	err := pr.db.QueryRow(query, product.Name, product.Stock, product.Price).Scan(&product.ID)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (pr *productRepository) Update(product models.Product) error {
	query := "UPDATE products SET name = $1, description = $2, WHERE id = $3"
	result, err := pr.db.Exec(query, product.Name, product.Price, product.Stock)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("product not found")
	}

	return nil
}

func (pr *productRepository) Delete(id uint) error {
	query := "DELETE FROM products WHERE id = $1"
	result, err := pr.db.Exec(query, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("product not found")
	}

	return err
}
