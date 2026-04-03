package database

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres driver
)

var DB *sqlx.DB

func Connect() {
	var err error
	dsn := os.Getenv("DB_URL")
	
	DB, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Database unreachable: %v", err)
	}

	log.Println("Database connection established")
}


