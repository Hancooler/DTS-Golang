package models

import (
	"gorm.io/gorm"
)

type PeminjamanDetail struct {
	gorm.Model
	BukuID       uint       `json:"buku_id"`
	PeminjamanID uint       `json:"peminjaman_id"`
	Buku         Buku       `json:"buku"`
	Peminjaman   Peminjaman `json:"peminjaman"`
}

type ResponsePeminjamanDetail struct {
	JudulBuku    string `json:"judul_buku"`
	GenreBuku    string `json:"genre_buku"`
	TahunBuku    string `json:"tahun_buku"`
	PenulisBuku  string `json:"penulis_buku"`
	PenerbitBuku string `json:"penerbit_buku"`
}

type RequestPeminjamanDetail struct {
	BukuID       uint `json:"buku_id"`
	PeminjamanID uint `json:"peminjaman_id"`
}

type ResponsePeminjamanDetailByPeminjaman struct {
	Peminjaman ResponsePeminjamanByBuku `json:"peminjaman"`
}
