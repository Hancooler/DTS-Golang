package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // blank import for sql driver
)

const (
	Host     = "localhost"
	Port     = 5432
	User     = "postgres"
	Password = "1"
	Dbname   = "postgres"
)

var (
	db  *sql.DB
	err error
)

func sqlInit() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", Host, Port, User, Password, Dbname)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
}
