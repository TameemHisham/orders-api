package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/TameemHisham/orders-api/application"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Import PostgreSQL driver
)

func main() {
	// Load the .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get environment variables
	PORT := os.Getenv("PORT")
	DB_PORT := os.Getenv("DB_PORT")
	PASSWORD := os.Getenv("PASSWORD")
	DB_USER := os.Getenv("DB_USER")
	DB_NAME := os.Getenv("DB_NAME")

	log.Printf("DB_PORT: %v\nPORT: %v\nPASSWORD: %v\nDB_USER: %v\nDB_NAME: %v\n", DB_PORT, PORT, PASSWORD, DB_USER, DB_NAME)

	// Construct the database connection string
	connStr := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s?sslmode=disable", DB_USER, PASSWORD, DB_PORT, DB_NAME)

	// Open a connection to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}
	defer db.Close() // Defer closing the database connection

	// Ping the database to verify connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	} else {
		fmt.Println("DB connected successfully")
	}
	CreateDatabase(db)
	// Initialize and start the application
	app := application.New()
	err = app.Start(context.TODO(), PORT)
	if err != nil {
		log.Fatalf("Failed to start app: %v", err)
	}
}


func CreateDatabase (db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS product
	(
		ID SERIAL PRIMARY KEY,
		FirstName VARCHAR(100) NOT NULL,
		LastName VARCHAR(100) NOT NULL,
		Price NUMERIC(6,2) NOT NULL,
		Availability BOOLEAN,
		created timestamp DEFAULT NOW()
	)`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Error creation database connection: %v", err)
	}

}