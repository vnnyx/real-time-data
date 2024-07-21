package messaging

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/goccy/go-json"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/vnnyx/real-time-data/internal/domain"
	"github.com/vnnyx/real-time-data/internal/usecase"
)

type NewsConsumer struct {
	newsUseCase usecase.NewsUseCase
	log         *zerolog.Logger
	consumer    *kafka.Consumer
}

func NewNewsConsumer(newsUseCase usecase.NewsUseCase, log *zerolog.Logger, consumer *kafka.Consumer) *NewsConsumer {
	return &NewsConsumer{
		newsUseCase: newsUseCase,
		log:         log,
		consumer:    consumer,
	}
}

func (nc *NewsConsumer) Consume(message *kafka.Message) error {
	newsEvent := new(domain.NewsEvent)
	if err := json.Unmarshal(message.Value, newsEvent); err != nil {
		nc.log.Error().Err(err).Msg("Failed to unmarshal message")
		return err
	}

	err := nc.newsUseCase.StoreToVectorDB(context.Background(), newsEvent)
	if err != nil {
		return err
	}

	if _, err := nc.consumer.CommitMessage(message); err != nil {
		log.Error().Err(err).Msg("Failed to commit message")
	}

	nc.log.Info().Msg("Message committed")

	return nil
}
