package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	NewsAPIKey     string `mapstructure:"NEWS_API_KEY"`
	NewsAPIBaseURL string `mapstructure:"NEWS_API_BASE_URL"`
	KafkaBrokerURL string `mapstructure:"KAFKA_BROKER_URL"`
	KafkaUsername  string `mapstructure:"KAFKA_USERNAME"`
	KafkaPassword  string `mapstructure:"KAFKA_PASSWORD"`
	VectorHost     string `mapstructure:"VECTOR_HOST"`
	NewsTopic      string `mapstructure:"NEWS_TOPIC"`
}

func NewConfig() (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
