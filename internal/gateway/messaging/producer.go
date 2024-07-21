package messaging

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/goccy/go-json"
	"github.com/rs/zerolog"
	"github.com/vnnyx/real-time-data/internal/domain"
)

type Producer[T domain.Event] struct {
	Producer *kafka.Producer
	Topic    string
	Log      *zerolog.Logger
}

func (p *Producer[T]) GetTopic() *string {
	return &p.Topic
}

func (p *Producer[T]) Produce(event T) error {
	value, err := json.Marshal(event)
	if err != nil {
		p.Log.Error().Err(err).Msg("Failed to marshal event")
		return err
	}

	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     p.GetTopic(),
			Partition: kafka.PartitionAny,
		},
		Value: value,
		Key:   []byte(event.GetID()),
	}

	err = p.Producer.Produce(message, nil)
	if err != nil {
		p.Log.Error().Err(err).Msg("Failed to produce message")
		return err
	}

	return nil
}
