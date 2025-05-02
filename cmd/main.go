package main

import (
	"database/sql"
	"fmt"
	"github.com/AshiishKarhade/url-shortner-go/pkg/database"
	"github.com/AshiishKarhade/url-shortner-go/pkg/hashing"
	"github.com/AshiishKarhade/url-shortner-go/pkg/snowflake"
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

	sf, err := snowflake.NewSnowflake(1)
	if err != nil {
		log.Fatalf("Failed to initialize snowflake: %v", err)
	}
	for i := 0; i < 10; i++ {
		nextID := sf.NextID()
		hashed := hashing.GenerateShortURL(nextID)
		fmt.Printf("Generated hash ID: %s\n", hashed)
	}
}
