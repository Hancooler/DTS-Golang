package repository

import (
	"tokokeebs/internal/model"

	"github.com/jinzhu/gorm"
)

type OrderItemRepository struct {
	db *gorm.DB
}

func NewOrderItemRepository(db *gorm.DB) *OrderItemRepository {
	return &OrderItemRepository{db}
}

func (oir *OrderItemRepository) CreateOrderItem(orderItem *model.OrderItem) error {
	return oir.db.Create(orderItem).Error
}

// Implementasikan metode lainnya sesuai kebutuhan
