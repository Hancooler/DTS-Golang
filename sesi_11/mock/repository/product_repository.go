package repository

import "sesi_11/mock/entity"

// ProductRepository represents an interface for product data access
type ProductRepository interface {
	// FindById retrieves a product by its ID
	FindById(id string) *entity.Product
}
