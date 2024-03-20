package models

import (
	"gorm.io/gorm"
)

type Buku struct {
	gorm.Model
	Judul    string `json:"judul"`
	Genre    string `json:"genre"`
	Tahun    string `json:"tahun"`
	Stok     uint   `json:"stok"`
	Penulis  string `json:"penulis"`
	Penerbit string `json:"penerbit"`
}

type ResponseBuku struct {
	ID       uint
	Judul    string `json:"judul"`
	Genre    string `json:"genre"`
	Tahun    string `json:"tahun"`
	Stok     uint   `json:"stok"`
	Penulis  string `json:"penulis"`
	Penerbit string `json:"penerbit"`
}

type RequestBuku struct {
	Judul    string `json:"judul"`
	Genre    string `json:"genre"`
	Tahun    string `json:"tahun"`
	Stok     uint   `json:"stok"`
	Penulis  string `json:"penulis"`
	Penerbit string `json:"penerbit"`
}
