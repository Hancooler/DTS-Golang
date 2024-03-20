package repository

import (
	"tokokeebs/internal/model"

	"github.com/jinzhu/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db}
}

func (or *OrderRepository) CreateOrder(order *model.Order) error {
	return or.db.Create(order).Error
}

// Implementasikan metode lainnya sesuai kebutuhan
