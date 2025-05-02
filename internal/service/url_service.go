package service

import (
	"context"
	"github.com/AshiishKarhade/url-shortner-go/internal/models"
	"github.com/AshiishKarhade/url-shortner-go/internal/repository"
	"github.com/AshiishKarhade/url-shortner-go/pkg/hashing"
	"github.com/AshiishKarhade/url-shortner-go/pkg/snowflake"
)

type URLService struct {
	urlRepo   *repository.URLRepository
	snowflake *snowflake.Snowflake
}

func NewURLService(urlRepo *repository.URLRepository, snowflake *snowflake.Snowflake) *URLService {
	return &URLService{
		urlRepo:   urlRepo,
		snowflake: snowflake,
	}
}

func (s *URLService) ShortenURL(ctx context.Context, longURL string) (*models.URL, error) {
	// create a unique ID using snowflake ID generator algorithm
	id := s.snowflake.NextID()
	// convert to base62
	shortURL := hashing.GenerateShortURL(id)

	url := &models.URL{
		ID:       id,
		LongURL:  longURL,
		ShortURL: shortURL,
	}
	// save to db
	if err := s.urlRepo.CreateURL(ctx, url); err != nil {
		return nil, err
	}
	return url, nil
}

func (s *URLService) GetLongURL(ctx context.Context, shortURL string) (string, error) {
	url, err := s.urlRepo.GetByShortURL(ctx, shortURL)
	if err != nil {
		return "", err
	}
	return url.LongURL, nil
}

func (s *URLService) UpdateURL(ctx context.Context, shortURL string, newLongURL string) error {
	// update the long URL in the database
	_, err := s.GetLongURL(ctx, shortURL)
	if err != nil {
		return err
	}
	return s.urlRepo.Update(ctx, shortURL, newLongURL)
}
