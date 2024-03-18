package controller

import (
	"assigment3/config"
	"assigment3/model"
	"fmt"
	"math/rand"

	_ "github.com/lib/pq"
)

func UpdateData() {
	config.Connect()

	waterValue := rand.Float64() * 100
	windValue := rand.Float64() * 100

	sensor := model.SensorData{}

	// sqlStatement := `"INSERT INTO sensor_data (water, wind) VALUES ($1, $2)", waterValue, windValue`

	err := config.GetDB().Create(&sensor).Error

	if err != nil {
		panic(err)
	}

	fmt.Printf("Data Updated - Water: %.2f, Wind: %.2f\n", waterValue, windValue)

	waterStatus := "aman"
	if waterValue < 5 {
		waterStatus = "aman"
	} else if waterValue >= 6 && waterValue <= 8 {
		waterStatus = "siaga"
	} else {
		waterStatus = "bahaya"
	}

	windStatus := "aman"
	if windValue < 6 {
		windStatus = "aman"
	} else if windValue >= 7 && windValue <= 15 {
		windStatus = "siaga"
	} else {
		windStatus = "bahaya"
	}

	fmt.Printf("Status Water: %s\n", waterStatus)
	fmt.Printf("Status Wind: %s\n", windStatus)
}
