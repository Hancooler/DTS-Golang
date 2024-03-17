package models

import (
	"sesi_10/jwt/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	FullName string    `gorm:"not null" json:"full_name" form:"full_name" validate:"required-Your name is required"`
	Email    string    `gorm:"not null;unique" json:"email" form:"email" validate:"required-Your Email is required"`
	password string    `gorm:"not null" json:"password" form:"password" valid:"required-Password is required,minstingleng(6)password has ti have a minimum length of 6 characters"`
	Products []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; json:"products"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}

	u.password = helpers.HassPass(u.password)
	err = nil
	return
}
