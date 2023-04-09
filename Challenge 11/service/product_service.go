package service

import (
	"Challenge_11/entity"
	"Challenge_11/repository"
	"errors"
)

type ProductService struct {
	Repository repository.ProductRepository
}

func (service ProductService) GetOneProduct(id string) (*entity.Product, error) {
	product := service.Repository.FindById(id)

	if product == nil {
		return nil, errors.New("Product not found")
	}

	return product, nil
}

func (service ProductService) GetAllProducts() (*entity.Product, error) {
	product := service.Repository.FindAll()

	if product == nil {
		return nil, errors.New("Product not found")
	}

	return product, nil
}
