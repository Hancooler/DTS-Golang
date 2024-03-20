package models

import (
	"gorm.io/gorm"
)

type Petugas struct {
	gorm.Model
	Name       string      `json:"name"`
	Email      string      `json:"email"`
	Borrowings []Borrowing `json:"borrowings"`
}
