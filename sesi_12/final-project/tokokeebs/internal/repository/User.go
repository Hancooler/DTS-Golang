package repository

import (
	"tokokeebs/internal/model"

	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) CreateUser(user *model.User) error {
	return ur.db.Create(user).Error
}

func (ur *UserRepository) GetUserByID(id uint) (*model.User, error) {
	user := &model.User{}
	if err := ur.db.First(user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// Implementasikan metode lainnya sesuai kebutuhan
