package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Task struct {
	ID          int
	Title       string
	Description string
	Status      string
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Veritabanı bağlantısı başarılı.")
}
