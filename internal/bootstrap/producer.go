package bootstrap

import (
	"context"
	"time"

	"github.com/vnnyx/real-time-data/internal/config"
	"github.com/vnnyx/real-time-data/internal/gateway/messaging"
	"github.com/vnnyx/real-time-data/internal/repository"
	"github.com/vnnyx/real-time-data/internal/usecase"
)

func StartProducer() {
	logger := config.NewZeroLog()
	cfg, err := config.NewConfig()
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to load config")
	}

	producer, err := config.NewKafkaProducer(cfg, logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to create Kafka producer")
	}
	httpClient := config.NewHttpClient()

	newsProducer := messaging.NewNewsProducer(producer, logger, cfg)
	conn, vectorClient, err := config.NewVectorClientConnection(cfg, logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to create gRPC connection")
	}
	defer conn.Close()
	newsRepo := repository.NewNewsRepository(httpClient, cfg, logger)
	newsUseCase := usecase.NewNewsUseCase(newsRepo, logger, newsProducer, vectorClient)

	// Get news and produce to Kafka every 30 minutes
	signal := make(chan struct{})
	go func() {
		for {
			newsUseCase.GetNews(context.Background())

			// Sleep for 30 minutes before fetching news again
			time.Sleep(30 * time.Minute)
		}
	}()

	<-signal
}
