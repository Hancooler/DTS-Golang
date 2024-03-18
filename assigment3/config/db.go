package config

import (
	"assigment3/model"
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1"
	dbname   = "db-go-sql"
)

var (
	db  *gorm.DB
	err error
)

func Connect() {
	psqlInfo := fmt.Sprintf(`
	host=%s 
	port=%d 
	user=%s 	
	password=%s 
	dbname=%s 
	sslmode=disable`,
		host, port, user, password, dbname)

	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Sukses konek ke DB\n", psqlInfo)

	db.AutoMigrate(&model.SensorData{})
}

func GetDB() *gorm.DB {
	return db
}
