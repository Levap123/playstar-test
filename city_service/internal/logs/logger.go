package logs

import (
	"os"

	"github.com/rs/zerolog"
)

type Logger struct {
	zerolog.Logger
}

func New() *Logger {
	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Str("service", "city").
		Caller().
		Logger()
	return &Logger{
		logger,
	}
}
