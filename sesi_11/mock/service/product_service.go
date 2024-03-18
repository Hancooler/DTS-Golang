package service

import (
	"errors"
	"sesi_11/mock/entity"
	"sesi_11/mock/repository"
)

// ProductService handles product-related operations
type ProductService struct {
	Repository repository.ProductRepository
}

// GetOneProduct retrieves a product by its ID
func (service ProductService) GetOneProduct(id string) (*entity.Product, error) {
	// Find the product by ID using the repository
	product := service.Repository.FindById(id)
	if product == nil {
		// Return an error if the product is not found
		return nil, errors.New("product not found")
	}

	// Return the found product and no error
	return product, nil
}
