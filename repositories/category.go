package repositories

import (
	"database/sql"
	"errors"
	"store/models"
)

type CategoryRepository interface {
	GetAll() ([]models.Category, error)
	GetByID(uint) (*models.Category, error)
	Insert(*models.Category) (*models.Category, error)
	Update(models.Category) error
	Delete(uint) error
}

type categoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (cr *categoryRepository) GetAll() ([]models.Category, error) {
	query := "SELECT id, name, description FROM categories"
	rows, err := cr.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	categories := make([]models.Category, 0)
	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.ID, &category.Name, &category.Description)
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (cr *categoryRepository) GetByID(id uint) (*models.Category, error) {
	query := "SELECT id, name, description FROM categories WHERE id= $1"

	var category models.Category
	err := cr.db.QueryRow(query, id).Scan(&category.ID, &category.Name, &category.Description)
	if err == sql.ErrNoRows {
		return nil, errors.New("category not found")
	}

	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (cr *categoryRepository) Insert(category *models.Category) (*models.Category, error) {
	query := "INSERT INTO categories (name, description) VALUES ($1, $2) returning id"
	err := cr.db.QueryRow(query, category.Name, category.Description).Scan(&category.ID)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (cr *categoryRepository) Update(category models.Category) error {
	query := "UPDATE categories SET name = $1, description = $2, WHERE id = $3"
	result, err := cr.db.Exec(query, category.Name, category.Description, category.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("category not found")
	}

	return nil
}

func (cr *categoryRepository) Delete(id uint) error {
	query := "DELETE FROM categories WHERE id = $1"
	result, err := cr.db.Exec(query, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("category not found")
	}

	return err
}
