package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // PostgreSQL driver
	"log"
	"time"
)

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func DefaultConfig() *Config {
	return &Config{
		Host:     "localhost",
		Port:     5432,
		Username: "postgres",
		Password: "password",
		DBName:   "url_shortner",
		SSLMode:  "disable",
	}
}

func ConnectPostgres() (*sql.DB, error) {
	config := DefaultConfig()
	return ConnectPostgresWithConfig(config)
}

func ConnectPostgresWithConfig(config *Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.Username, config.Password, config.DBName, config.SSLMode)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	log.Println("Successfully connected to PostgreSQL database")

	// Initialise the database schema if needed
	initQuery := `
		CREATE TABLE IF NOT EXISTS urls (
			id BIGINT PRIMARY KEY,
			long_url TEXT NOT NULL,
			short_url VARCHAR(10) UNIQUE NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);

		CREATE INDEX IF NOT EXISTS idx_short_url ON urls(short_url);
	`

	if _, err := db.Exec(initQuery); err != nil {
		return nil, fmt.Errorf("failed to initialize database schema: %w", err)
	}
	log.Println("Database schema initialized successfully")

	return db, nil
}
