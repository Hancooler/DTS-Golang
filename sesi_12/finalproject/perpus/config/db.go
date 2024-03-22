package config

import (
	"perpus/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	dsn := "host=localhost user=Postgres password=1 dbname=perpus port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
func main() {
	// Establish database connection
	db, err := ConnectDB()
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.Book{}, &models.Borrowing{}, &models.Member{}, &models.Petugas{})

	// Use db to perform operations...
}
