package helper

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// InitializeDB initialize database
var (
	username = "postgres"
	password = "1"
	dbDriver = "postgres"
	dbPort   = "5432"
	dbName   = "db"
	host     = "localhost"
	db       *sql.DB
	err      error
)

func createRequiredTables() {
	itemsTable := `
			CREATE TABLE IF NOT EXISTS items (
	 			id SERIAL PRIMARY KEY,
	 			itemCode varchar(10) NOT NULL,
	 			description TEXT NOT NULL,
	 			quantity int NOT NULL,
	 			orderId int NOT NULL
	 		);
`
	ordersTable := `
		CREATE TABLE IF NOT EXISTS ordersItems (
	  		id SERIAL PRIMARY KEY,
	 		customerName varchar (255) NOT NULL,
	 		orderedAt timestamptz NOT NULL DEFAULT (now())
	 		);
	 `

	createTableQueries := fmt.Sprintf("%s ", itemsTable)
	_, err = db.Exec(createTableQueries)

	if err != nil {
		log.Fatal("error while creating bukan movies table =>", err.Error())
	}
}

func InitializeDB() {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, dbPort, dbName)

	db, err = sql.Open(dbDriver, dsn)

	if err != nil {
		log.Fatal("error connecting to database", err.Error())
	}

	err = db.Ping()

	if err != nil {
		log.Fatal("error while tyring to ping the database connection", err.Error())
	}

	fmt.Println("successfully connected to my database")
	createRequiredTables()
}

func GetDB() *sql.DB {
	return db
}
