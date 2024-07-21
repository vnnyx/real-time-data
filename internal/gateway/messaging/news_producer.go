package messaging

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/rs/zerolog"
	"github.com/vnnyx/real-time-data/internal/config"
	"github.com/vnnyx/real-time-data/internal/domain"
)

type NewsProducer struct {
	Producer[*domain.NewsEvent]
}

func NewNewsProducer(producer *kafka.Producer, log *zerolog.Logger, cfg *config.Config) *NewsProducer {
	return &NewsProducer{
		Producer[*domain.NewsEvent]{
			Producer: producer,
			Topic:    cfg.NewsTopic,
			Log:      log,
		},
	}
}
