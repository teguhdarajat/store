package services

import (
	"store/models"
	"store/params"
	"store/repositories"
)

type CategoryService interface {
	GetAll() (*params.Categories, error)
	GetByID(uint) (*params.Category, error)
	Insert(params.Category) (*params.Category, error)
	Update(params.Category) (*params.Category, error)
	Delete(uint) error
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

func (cs *categoryService) GetAll() (*params.Categories, error) {
	res, err := cs.categoryRepository.GetAll()
	if err != nil {
		return nil, err
	}
	category := toCategoryParams(res)
	return &category, nil
}

func (cs *categoryService) GetByID(id uint) (*params.Category, error) {
	res, err := cs.categoryRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	category := toCategoryParam(*res)
	return &category, err
}

func (cs *categoryService) Insert(category params.Category) (*params.Category, error) {
	categoryModel := toCategoryModel(category)
	res, err := cs.categoryRepository.Insert(&categoryModel)
	if err != nil {
		return nil, err
	}

	category = toCategoryParam(*res)
	return &category, err
}

func (cs *categoryService) Update(update params.Category) (*params.Category, error) {
	categoryModel := toCategoryModel(update)

	err := cs.categoryRepository.Update(categoryModel)
	if err != nil {
		return nil, err
	}

	return &update, err
}

func (cs *categoryService) Delete(id uint) error {
	err := cs.categoryRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
