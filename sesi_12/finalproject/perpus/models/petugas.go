package models

import (
	"gorm.io/gorm"
)

type Petugas struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `json:"username"`
	Nip      string `json:"nip"`
	Password string `json:"password"`
	Role     uint   `json:"role"`
}

type ResponsePetugas struct {
	ID       uint
	Name     string `json:"name"`
	Username string `json:"username"`
	Nip      string `json:"nip"`
	Role     uint   `json:"role"`
}

type RequestPetugas struct {
	ID       uint
	Name     string `json:"name"`
	Username string `json:"username"`
	Nip      string `json:"nip"`
	Password string `json:"password"`
	Role     uint   `json:"role"`
}

type TokenRequest struct {
	Nip      string `json:"nip"`
	Password string `json:"password"`
}
