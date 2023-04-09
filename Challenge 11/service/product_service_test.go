package service

import (
	"Challenge_11/entity"
	"Challenge_11/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = ProductService{Repository: productRepository}

func TestProductServiceGetOneProductNotFound(t *testing.T) {
	productRepository.Mock.On("FindById", "1").Return(nil)

	product, err := productService.GetOneProduct("1")

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, "Product not found", err.Error(), "error response has to be 'Product not found'")
}

func TestProductServiceGetOneProduct(t *testing.T) {
	product := entity.Product{
		Id:   "2",
		Name: "Buku novel",
	}

	productRepository.Mock.On("FindById", "2").Return(product)

	result, err := productService.GetOneProduct("2")

	assert.Nil(t, err)

	assert.NotNil(t, result)

	assert.Equal(t, product.Id, result.Id, "Result has to be '2'")
	assert.Equal(t, product.Name, result.Name, "Result has to be 'Buku novel'")
	assert.Equal(t, &product, result, "Result has to be a product data with Id '2'")
}

func TestProductServiceGetAllProductsNotFound(t *testing.T) {
	productRepository.Mock.On("FindAll").Return(nil)

	product, err := productService.GetAllProducts()

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, "Product not found", err.Error(), "error response has to be 'Product not found'")
}

func TestProductServiceGetAllProducts(t *testing.T) {
	var products = []entity.Product{
		{Id: "3", Name: "Ensiklopedia,"},
		{Id: "4", Name: "Majalah"},
	}

	productRepository.Mock.On("FindAll").Return(products)

	result, err := productService.GetAllProducts()

	assert.Nil(t, err)

	assert.NotNil(t, result)

	assert.Equal(t, products[0].Id, result.Id, "Result has to be '3'")
	assert.Equal(t, products[0].Name, result.Name, "Result has to be 'Ensiklopedia'")

}
