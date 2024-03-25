package pg_storage

import (
	"database/sql"
	_ "embed"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var db *sql.DB = nil

//go:embed init.sql
var initSQL string

func Connect() {
	log.Printf("Connecting to the database")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	username, ok1 := os.LookupEnv("DB_USER")
	password, ok2 := os.LookupEnv("DB_PASSWORD")
	databaseHost, ok3 := os.LookupEnv("DB_HOST")
	databasePort, ok4 := os.LookupEnv("DB_PORT")
	databaseName, ok5 := os.LookupEnv("DB_NAME")

	if !ok1 || !ok2 || !ok3 || !ok4 || !ok5 {
		log.Printf("Database environment variables are not set")
		return
	}

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", databaseHost, username, password, databaseName, databasePort)

	db, err = sql.Open("postgres", connStr)
	if err == nil {
		log.Printf("Succesful connecting to the database")
	}
	ErrorHandler(err, "Error while connecting to the database")
	setupDB()
}

func setupDB() {
	if db == nil {
		log.Printf("Database is not connected")
		return
	}

	_, err := db.Exec(initSQL)

	log.Println("[INFO]: Creating the table")

	ErrorHandler(err, "Error while creating the table")
}
