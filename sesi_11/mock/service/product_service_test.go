package service

import (
	"mock/entity"
	"mock/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	mock.Mock
}

var productService = ProductService{Repository: productRepository}
var productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}

func TestProductServiceGetOneProductNotFound(t *testing.T) {
	// Prepare the mock behavior
	productRepository.Mock.On("FindById", "1").Return(nil)

	// Call the method under test
	product, err := productService.GetOneProduct("1")

	// Assert the results
	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, "product not found", err.Error(), "error response has to be 'product not found'")
}

func TestProductServiceGetOneProduct(t *testing.T) {
	// Prepare the mock behavior
	product := entity.Product{
		ID:   "2",
		Name: "Sun Glasses",
	}
	productRepository.Mock.On("FindById", "2").Return(&product)

	// Call the method under test
	result, err := productService.GetOneProduct("2")

	// Assert the results
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, product.ID, result.ID, "result has to be '2'")
	assert.Equal(t, product.Name, result.Name, "result has to be 'Sun Glasses'")
	assert.Equal(t, &product, result, "result has to be a product data with ID '2'")
}
