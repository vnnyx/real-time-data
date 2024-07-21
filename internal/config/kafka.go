package config

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/rs/zerolog"
)

func NewKafkaConsumer(config *Config, log *zerolog.Logger, groupID string) (*kafka.Consumer, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": config.KafkaBrokerURL,
		"sasl.username":     config.KafkaUsername,
		"sasl.password":     config.KafkaPassword,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
		"sasl.mechanisms":   "SCRAM-SHA-256",
		"security.protocol": "SASL_SSL",
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to create Kafka consumer")
		return nil, err
	}

	return consumer, nil
}

func NewKafkaProducer(config *Config, log *zerolog.Logger) (*kafka.Producer, error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":                     config.KafkaBrokerURL,
		"sasl.username":                         config.KafkaUsername,
		"sasl.password":                         config.KafkaPassword,
		"sasl.mechanisms":                       "SCRAM-SHA-256",
		"security.protocol":                     "SASL_SSL",
		"enable.idempotence":                    true,
		"max.in.flight.requests.per.connection": 1,
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to create Kafka producer")
		return nil, err
	}

	return producer, nil
}
