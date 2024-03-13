package entity

import "time"

type Products struct {
	Kodeproduct string `gorm:"column:kode_product;primary_key`
	Namaproduct string `gorm:"column:nama_product"`
	Stok        int    `gorm:"column:stok"`
	Harga       int    `gorm:"column:harga"`
}

type Orders struct {
	Id_order      int       `gorm:"column:id_order;primary_key:auto_increment"`
	Tanggalorder  time.Time `gorm:"column:tanggal_order"`
	Paymentmethod string    `gorm:"column:payment_method"`
	Total         int       `gorm:"column:total"`
}

type OrderDetails struct {
	Idorderdetail       int      `gorm:"column:id_order_detail;primary_key:auto_increment"`
	OrdersIdOrder       int      `gorm:"column:id_order"`
	Orders              Orders   `gorm:"foreignKey:OrdersIdOrder"`
	ProductsKodeproduct string   `gorm:"column:kode_product"`
	Product             Products `gorm:"foreignKey:ProductsKodeproduct"`
	Qty                 int      `gorm:"column:qty"`
}
