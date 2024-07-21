package usecase

import (
	"context"
	"sync"

	"github.com/goccy/go-json"
	"github.com/rs/zerolog"
	"github.com/vnnyx/real-time-data/internal/config"
	"github.com/vnnyx/real-time-data/internal/domain"
	"github.com/vnnyx/real-time-data/internal/gateway/messaging"
	"github.com/vnnyx/real-time-data/internal/repository"
	"github.com/vnnyx/real-time-data/pb/vector"
)

type newsUseCase struct {
	newsRepo     repository.NewsRepository
	log          *zerolog.Logger
	newsProducer *messaging.NewsProducer
	vectorClient vector.VectorClient
}

func NewNewsUseCase(newsRepo repository.NewsRepository, log *zerolog.Logger, newsProducer *messaging.NewsProducer, vectorClient vector.VectorClient) NewsUseCase {
	return &newsUseCase{
		newsRepo:     newsRepo,
		log:          log,
		newsProducer: newsProducer,
		vectorClient: vectorClient,
	}
}

func (nu *newsUseCase) GetNews(ctx context.Context) (*domain.NewsAPIResponse, error) {
	newsData, err := nu.newsRepo.GetNews(ctx, "bitcoin OR ethereum OR solana OR doge coin OR shiba inu coin OR meme coin OR cryptocurrency OR digital currency OR blockchain OR decentralized finance OR coin OR token")
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup

	for _, article := range newsData.Articles {
		wg.Add(1)

		go func(article domain.Article) {
			defer wg.Done()

			nu.parseDataAndProduce(article)
		}(article)
	}
	wg.Wait()

	return newsData, nil
}

func (nu *newsUseCase) StoreToVectorDB(ctx context.Context, articles *domain.NewsEvent) error {
	vCtx, cancel := config.NewVectorClientContext()
	defer cancel()

	jsonData, err := json.Marshal(articles)
	if err != nil {
		nu.log.Error().Err(err).Msg("Failed to marshal data")
		return err
	}

	res, err := nu.vectorClient.StoreToVectorDB(vCtx, &vector.DataVectorRequest{
		Data: string(jsonData),
		Type: "News",
	})
	if err != nil {
		nu.log.Error().Err(err).Msg("Failed to store data to vector db")
		return err
	}
	nu.log.Info().Msgf("Stored data to vector db: %s -> %v", articles.Title, res.Value)

	return nil
}

func (nu *newsUseCase) parseDataAndProduce(newsData domain.Article) error {
	event := newsData.ToNewsEvent()
	if err := nu.newsProducer.Produce(event); err != nil {
		nu.log.Error().Err(err).Msg("Failed to produce message")
		return err
	}

	nu.log.Info().Msg("Produced message to Kafka")

	return nil
}
