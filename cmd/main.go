package main

import (
	"database/sql"
	"github.com/AshiishKarhade/url-shortner-go/pkg/database"
	"log"
)

func main() {

	db, err := database.ConnectPostgres()
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Failed to close database connection: %v", err)
		}
	}(db)
}
