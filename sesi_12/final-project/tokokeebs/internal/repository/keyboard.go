package repository

import (
	"tokokeebs/internal/model"

	"github.com/jinzhu/gorm"
)

type KeyboardRepository struct {
	db *gorm.DB
}

func NewKeyboardRepository(db *gorm.DB) *KeyboardRepository {
	return &KeyboardRepository{db}
}

func (kr *KeyboardRepository) CreateKeyboard(keyboard *model.Keyboard) error {
	return kr.db.Create(keyboard).Error
}

// Implementasikan metode lainnya sesuai kebutuhan
