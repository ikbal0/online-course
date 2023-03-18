package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func GetConnection() *sql.DB {
	envErr := godotenv.Load(".env")

	if envErr != nil {
		fmt.Printf("Could not load .env file")
		os.Exit(1)
	}

	db, err := sql.Open("mysql", os.Getenv("MYSQL_URI"))

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
