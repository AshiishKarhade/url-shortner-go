package controller

import (
	"context"
	"database/sql"
	"github.com/AshiishKarhade/url-shortner-go/internal/models"
	"time"
)

type URLRepository struct {
	db *sql.DB
}

func NewURLRepository(db *sql.DB) *URLRepository {
	return &URLRepository{db: db}
}

func (repo *URLRepository) CreateURL(ctx context.Context, url *models.URL) error {
	query := `INSERT INTO urls (id, long_url, short_url, created_at, updated_at)
				VALUES ($1, $2, $3, $4, $5)
			`
	_, err := repo.db.ExecContext(ctx, query, url.ID, url.LongURL, url.ShortURL, time.Now(), time.Now())
	return err
}

func (repo *URLRepository) GetByShortURL(ctx context.Context, shortURL string) (*models.URL, error) {
	query := `SELECT id, long_url, short_url, created_at, updated_at
				FROM urls
				WHERE short_url = $1
			`
	url := &models.URL{}
	row := repo.db.QueryRowContext(ctx, query, shortURL)
	err := row.Scan(&url.ID, &url.LongURL, &url.ShortURL, &url.CreatedAt, &url.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return url, nil
}

func (repo *URLRepository) Delete(ctx context.Context, shortURL string) error {
	query := `
		DELETE FROM urls
		WHERE short_url = $1
		`

	_, err := repo.db.ExecContext(ctx, query, shortURL)
	return err
}

func (repo *URLRepository) Update(ctx context.Context, shortURL string, newLongURL string) error {
	query := `
		UPDATE urls
		SET long_url = $1, updated_at = $2
		WHERE short_url = $3
	`

	_, err := repo.db.ExecContext(ctx, query, newLongURL, time.Now(), shortURL)
	return err
}
