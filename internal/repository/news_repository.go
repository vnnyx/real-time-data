package repository

import (
	"context"

	"github.com/vnnyx/real-time-data/internal/domain"
)

type NewsRepository interface {
	GetNews(ctx context.Context, query string) (*domain.NewsAPIResponse, error)
}
