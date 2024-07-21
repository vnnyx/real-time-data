package messaging

import (
	"context"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/rs/zerolog"
)

type ConsumerHandler func(message *kafka.Message) error

func ConsumeTopic(ctx context.Context, consumer *kafka.Consumer, topic string, log *zerolog.Logger, handler ConsumerHandler) {
	if err := subscribeToTopic(consumer, topic, log); err != nil {
		return
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Error().Err(err).Msg("Failed to close consumer")
		}
	}()

	run := true

	for run {
		select {
		case <-ctx.Done():
			run = false
		default:
			processMessage(consumer, log, handler)
		}
	}
}

func subscribeToTopic(consumer *kafka.Consumer, topic string, log *zerolog.Logger) error {
	if err := consumer.Subscribe(topic, nil); err != nil {
		log.Error().Err(err).Msg("Failed to subscribe to topic")
		return err
	}
	return nil
}

func processMessage(consumer *kafka.Consumer, log *zerolog.Logger, handler ConsumerHandler) {
	message, err := consumer.ReadMessage(5 * time.Second)
	if err != nil {
		handleReadMessageError(err, log)
		return
	}

	if err := handler(message); err != nil {
		log.Error().Err(err).Msg("Failed to handle message")
		return
	}
}

func handleReadMessageError(err error, log *zerolog.Logger) {
	if kafkaErr, ok := err.(kafka.Error); ok && kafkaErr.Code() == kafka.ErrTimedOut {
		log.Warn().Msg("Timed out waiting for message")
	} else {
		log.Error().Err(err).Msg("Failed to read message")
	}
}
