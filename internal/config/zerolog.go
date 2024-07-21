package config

import (
	"os"

	"github.com/rs/zerolog"
)

func NewZeroLog() *zerolog.Logger {
	log := zerolog.New(os.Stderr).With().Timestamp().Logger().Level(zerolog.TraceLevel)
	return &log
}
