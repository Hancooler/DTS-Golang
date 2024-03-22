package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"perpus/config"
	"perpus/controller"
	"perpus/middleware"
	routes "perpus/routers"
)

func main() {
	// Inisialisasi router Gin
	r := gin.Default()

	// Koneksi ke database
	dsn := "host=localhost user=postgres password=1 dbname=perpus port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	config.MigrateDB(db)

	// Menambahkan middleware JWT
	r.Use(middleware.AuthMiddleware())

	// Inisialisasi controller dengan koneksi database yang sudah disediakan
	bookController := &controller.BookController{DB: db}
	borrowingController := &controller.BorrowingController{DB: db}
	memberController := &controller.MemberController{DB: db}
	petugasController := &controller.PetugasController{DB: db}

	// Mengatur rute untuk model-model
	routes.SetupAPIRoutes(r, bookController, borrowingController, memberController, petugasController)

	// Menjalankan server pada port 5000
	r.Run(":5000")
}
