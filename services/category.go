package services

import (
	"store/models"
	"store/params"
	"store/repositories"
)

type CategoryService interface {
	GetCategories() params.Categories
	GetCategoryByID(uint) (*params.Category, error)
	InsertCategory(params.Category) params.Category
	UpdateCategory(uint, params.Category) (*params.Category, error)
	DeleteCategory(uint) error
}

type categoryService struct {
	categoryRepository repositories.CategoryRepository
}

func NewCategoryService(categoryRepository repositories.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepository: categoryRepository,
	}
}

func toCategoryParam(category models.Category) params.Category {
	return params.Category{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}
}

func toCategoryParams(categories []models.Category) params.Categories {
	categoriesParam := make([]params.Category, len(categories))
	for index, category := range categories {
		categoriesParam[index] = toCategoryParam(category)
	}

	return params.Categories{
		Categories: categoriesParam,
	}
}

func toCategoryModel(category params.Category) models.Category {
	return models.Category{
		Name:        category.Name,
		Description: category.Description,
	}
}

func (cs *categoryService) GetCategories() params.Categories {
	res := cs.categoryRepository.GetCategories()
	return toCategoryParams(res)
}

func (cs *categoryService) GetCategoryByID(id uint) (*params.Category, error) {
	res, err := cs.categoryRepository.GetCategoryByID(id)
	if err != nil {
		return nil, err
	}

	category := toCategoryParam(*res)
	return &category, err
}

func (cs *categoryService) InsertCategory(category params.Category) params.Category {
	categoryModel := toCategoryModel(category)
	res := cs.categoryRepository.InsertCategory(categoryModel)
	return toCategoryParam(res)
}

func (cs *categoryService) UpdateCategory(id uint, updateCategory params.Category) (*params.Category, error) {
	categoryModel := toCategoryModel(updateCategory)

	res, err := cs.categoryRepository.UpdateCategory(id, categoryModel)
	if err != nil {
		return nil, err
	}

	category := toCategoryParam(*res)
	return &category, err
}

func (cs *categoryService) DeleteCategory(id uint) error {
	err := cs.categoryRepository.DeleteCategory(id)
	if err != nil {
		return err
	}

	return nil
}
