package main

import (
	"sesi_10/jwt/database"
	"sesi_10/jwt/routers"
)

func main() {
	database.StartDB()
	r := routers.StartApp()
	r.Run(":5000")
}
