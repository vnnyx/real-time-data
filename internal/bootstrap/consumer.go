package bootstrap

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/vnnyx/real-time-data/internal/config"
	"github.com/vnnyx/real-time-data/internal/delivery/messaging"
	gateway "github.com/vnnyx/real-time-data/internal/gateway/messaging"
	"github.com/vnnyx/real-time-data/internal/repository"
	"github.com/vnnyx/real-time-data/internal/usecase"
)

func StartConsumer() {
	logger := config.NewZeroLog()

	cfg, err := config.NewConfig()
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to load config")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	producer, err := config.NewKafkaProducer(cfg, logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to create Kafka producer")

	}

	httpClient := config.NewHttpClient()

	newsProducer := gateway.NewNewsProducer(producer, logger, cfg)

	conn, vectorClient, err := config.NewVectorClientConnection(cfg, logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to create gRPC connection")
	}
	defer conn.Close()

	newsRepo := repository.NewNewsRepository(httpClient, cfg, logger)

	newsUseCase := usecase.NewNewsUseCase(newsRepo, logger, newsProducer, vectorClient)

	newsConsumer, err := config.NewKafkaConsumer(cfg, logger, "news-consumer")
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to create Kafka consumer")
	}

	newsHandler := messaging.NewNewsConsumer(newsUseCase, logger, newsConsumer)

	terminateSignals := make(chan os.Signal, 1)
	signal.Notify(terminateSignals, syscall.SIGINT, syscall.SIGTERM)

	go messaging.ConsumeTopic(ctx, newsConsumer, cfg.NewsTopic, logger, newsHandler.Consume)

	s := <-terminateSignals
	logger.Info().Msgf("Got one of stop signals, shutting down worker gracefully, SIGNAL NAME : %s", s.String())
	cancel()

	time.Sleep(5 * time.Second)
}
