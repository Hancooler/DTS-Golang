package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	UserID     uint
	TotalPrice float64
	Status     string
	CreatedAt  time.Time
	User       User `gorm:"foreignkey:UserID"`
}
