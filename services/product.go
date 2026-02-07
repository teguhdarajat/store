package services

import (
	"store/models"
	"store/params"
	"store/repositories"
)

type ProductService interface {
	GetAll() (*params.Products, error)
	GetByID(uint) (*params.Product, error)
	Insert(params.Product) (*params.Product, error)
	Update(params.Product) (*params.Product, error)
	Delete(uint) error
}

type productService struct {
	productRepository repositories.ProductRepository
}

func NewProductService(productRepository repositories.ProductRepository) ProductService {
	return &productService{
		productRepository: productRepository,
	}
}

func toProductParam(product models.Product) params.Product {
	return params.Product{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
		Stock: product.Stock,
	}
}

func toProductParams(products []models.Product) params.Products {
	productsParam := make([]params.Product, len(products))
	for index, product := range products {
		productsParam[index] = toProductParam(product)
	}

	return params.Products{
		Products: productsParam,
	}
}

func toProductModel(product params.Product) models.Product {
	return models.Product{
		Name:  product.Name,
		Price: product.Price,
		Stock: product.Stock,
	}
}

func (ps *productService) GetAll() (*params.Products, error) {
	res, err := ps.productRepository.GetAll()
	if err != nil {
		return nil, err
	}
	products := toProductParams(res)
	return &products, nil
}

func (ps *productService) GetByID(id uint) (*params.Product, error) {
	res, err := ps.productRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	product := toProductParam(*res)
	return &product, err
}

func (ps *productService) Insert(product params.Product) (*params.Product, error) {
	productModel := toProductModel(product)
	res, err := ps.productRepository.Insert(&productModel)
	if err != nil {
		return nil, err
	}

	product = toProductParam(*res)
	return &product, err
}

func (ps *productService) Update(update params.Product) (*params.Product, error) {
	productModel := toProductModel(update)

	err := ps.productRepository.Update(productModel)
	if err != nil {
		return nil, err
	}

	return &update, err
}

func (ps *productService) Delete(id uint) error {
	err := ps.productRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
