package repositories

import (
	"errors"
	"store/models"
)

type CategoryRepository interface {
	GetCategories() []models.Category
	GetCategoryByID(uint) (*models.Category, error)
	InsertCategory(models.Category) models.Category
	UpdateCategory(uint, models.Category) (*models.Category, error)
	DeleteCategory(uint) error
}

type categoryRepository struct {
	categories []models.Category
	lastID     uint
}

func NewCategoryRepository(categories []models.Category) CategoryRepository {
	var lastID uint
	for _, c := range categories {
		if c.ID > lastID {
			lastID = c.ID
		}
	}
	return &categoryRepository{
		categories: categories,
		lastID:     lastID,
	}
}

func (cr *categoryRepository) GetCategories() []models.Category {
	return cr.categories
}

func (cr *categoryRepository) GetCategoryByID(id uint) (*models.Category, error) {
	for i := range cr.categories {
		if cr.categories[i].ID == id {
			return &cr.categories[i], nil
		}
	}

	return nil, errors.New("category not found")
}

func (cr *categoryRepository) InsertCategory(newCategory models.Category) models.Category {
	cr.lastID++
	newCategory.ID = cr.lastID
	cr.categories = append(cr.categories, newCategory)
	return newCategory
}

func (cr *categoryRepository) UpdateCategory(id uint, updateCategory models.Category) (*models.Category, error) {
	for i := range cr.categories {
		if cr.categories[i].ID == id {
			updateCategory.ID = cr.categories[i].ID
			cr.categories[i] = updateCategory
			return &cr.categories[i], nil
		}
	}

	return nil, errors.New("category not found")
}

func (cr *categoryRepository) DeleteCategory(id uint) error {
	for i := range cr.categories {
		if cr.categories[i].ID == id {
			cr.categories = append(
				cr.categories[:i],
				cr.categories[i+1:]...,
			)
			return nil
		}
	}
	return errors.New("category not found")
}
