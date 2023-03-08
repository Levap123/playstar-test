package logs

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

type Logger struct {
	zerolog.Logger
}

func New() (*Logger, error) {
	f, err := os.OpenFile("logs.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	logger := zerolog.New(io.MultiWriter(os.Stdout, f)).With().
		Timestamp().
		Str("service", "weather").
		Caller().
		Logger()
	return &Logger{
		logger,
	}, nil
}
