package item

import "gorm.io/gorm"

type repository interface {
	Get() ([]Item, error)
	Create(item Item) (Item, error)
}

type itemRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *itemRepository {
	return &itemRepository{db}
}

func (r *itemRepository) Get() ([]Item, error) {
	var items []Item
	_, err := r.db.Find(&items).Error

}
