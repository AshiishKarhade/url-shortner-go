package main

import (
	"database/sql"
	"github.com/AshiishKarhade/url-shortner-go/internal/controller"
	"github.com/AshiishKarhade/url-shortner-go/internal/repository"
	"github.com/AshiishKarhade/url-shortner-go/internal/service"
	"github.com/AshiishKarhade/url-shortner-go/pkg/database"
	"github.com/AshiishKarhade/url-shortner-go/pkg/snowflake"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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

	// Initialize the Snowflake ID generator
	sf, err := snowflake.NewSnowflake(1) // Machine ID 1
	if err != nil {
		log.Fatalf("Failed to initialize snowflake: %v", err)
	}

	urlRepo := repository.NewURLRepository(db)
	urlService := service.NewURLService(urlRepo, sf)
	urlController := controller.NewURLController(urlService)

	router := gin.Default()

	router.POST("/shorten", urlController.ShortenURL)
	router.GET("/:shortURL", urlController.RedirectURL)
	router.PUT("/:shortURL", urlController.UpdateURL)

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
