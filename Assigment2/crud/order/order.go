package crud

import (
	"Assigment2/entity"

	"gorm.io/gorm"
)

func InsertOrder(db *gorm.DB, order *entity.Orders){
	result := db.Create(order)

	if result.Error != nil {
		panic("Failed to insert order")
		
	}
}

func InsertOrderDetails(db *gorm.DB, OrderDetails *[]entity.OrderDetails){
	result := db.Create(OrderDetails){

		if result.Error != nil {
			panic("Failed to insert order details")
			
		}
}
}