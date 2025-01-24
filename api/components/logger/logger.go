package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Init() {
	// Set up pretty logging for development
	if os.Getenv("ENV") != "production" {
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		})
	} else {
		// Production JSON logging
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	}

	// Set global log level
	level := os.Getenv("LOG_LEVEL")
	if level == "" {
		level = "info"
	}
	logLevel, err := zerolog.ParseLevel(level)
	if err != nil {
		log.Warn().Err(err).Str("level", level).Msg("Invalid log level, defaulting to info")
		logLevel = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(logLevel)
}
