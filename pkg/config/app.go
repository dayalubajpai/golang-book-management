package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "admin:adminpassword@tcp(127.0.0.1:3306)/bookstore")

	if err != nil {
		log.Fatal("Error while connecting to the database", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error while connecting to the database", err)
	}

	fmt.Println("✅ Successfully connected to MySQL!")
	return db
}

var DB *sql.DB = ConnectDB()

func CloseDB(db *sql.DB) error {
	return db.Close()
}
