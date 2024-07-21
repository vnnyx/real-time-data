package usecase

import (
	"context"

	"github.com/vnnyx/real-time-data/internal/domain"
)

type NewsUseCase interface {
	GetNews(ctx context.Context) (*domain.NewsAPIResponse, error)
	StoreToVectorDB(ctx context.Context, articles *domain.NewsEvent) error
}
