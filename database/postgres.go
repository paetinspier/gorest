package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

var DB *sql.DB

func StartPostgresDB() error {
	host := os.Getenv("DB_HOST")
	if host == "" {
		return errors.New("missing DB_HOST environment variable")
	}
	port := os.Getenv("DB_PORT")
	if port == "" {
		return errors.New("missing DB_PORT environment variable")
	}
	user := os.Getenv("DB_USERNAME")
	if user == "" {
		return errors.New("missing DB_USERNAME environment variable")
	}
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		return errors.New("missing DB_PASSWORD environment variable")
	}
	dbname := os.Getenv("DB_DATABASE")
	if dbname == "" {
		return errors.New("missing DB_DATABASE environment variable")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		return errors.New("can't verify a connection")
	}

	fmt.Println("Successfully connected!")

	return nil
}

func ClosePostgresDB() {
	err := DB.Close()
	if err != nil {
		panic(err)
	}
}
