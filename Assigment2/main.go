package main

import (
	"Assigment2/config"
	crud_order "Assigment2/crud/order"
	crud_product "Assigment2/crud/products"
	"Assigment2/entity"
	"time"
)

func main() {
	config.ConnectDB()

	product := entity.Products{
		KodeProduk: "001",
		NamaProduk: "somay",
		Stok:       "100",
		Harga:      "2000",
	}

	crud_product.InsertProduct(config.DB, &product)

	product := entity.Products{
		KodeProduk: "002",
		NamaProduk: "bakso",
		Stok:       "100",
		Harga:      `2000`,
	}

	crud_product.InsertProduct(config.DB, &product)

	order := entity.Orders{
		Id_order:      0,
		Tanggalorder:  time.Now(),
		Paymentmethod: "",
		Total:         30000,
	}

	crud_order.InsertOrder(config.DB, &order)

	details := []entity.OrderDetails{
		{
			Orders: order,
			Products: product.entity.Products{
				KodeProduct: "001",
			},
			Qty: 10,
		},
		{
			Orders: order,
			Products: product.entity.Products{
				KodeProduct: "002",
			},
			Qty: 10,
		},
	}
}
