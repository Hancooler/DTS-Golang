package models

import (
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	Name        string      "json:name"
	Address     string      "json:address"
	PhoneNumber string      "json:phone_number"
	Email       string      "json:email"
	Borrowings  []Borrowing "json:borrowings"
}
