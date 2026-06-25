package sdklog

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
)

var (
	ErrUnsupportedLogMode  = errors.New("unsupported log mode")
	ErrUnsupportedLogLevel = errors.New("unsupported log level")
)

type LogMode string

const (
	ProductionMode  LogMode = "production"
	DevelopmentMode LogMode = "development"
)

func ParseLogMode(value string) (LogMode, error) {
	switch casted := LogMode(value); casted {
	case ProductionMode, DevelopmentMode:
		return casted, nil
	default:
		return "", ErrUnsupportedLogMode
	}
}

func ParseLogLevel(value string) (slog.Level, error) {
	var level slog.Level
	if err := level.UnmarshalText([]byte(value)); err != nil {
		return 0, ErrUnsupportedLogLevel
	}

	return level, nil
}

func NewLog(opts ...LogOption) (*slog.Logger, error) {
	options := defaultLogOptions()
	for _, opt := range opts {
		if err := opt(options); err != nil {
			return nil, fmt.Errorf("cannot initialize log: %w", err)
		}
	}

	var handlerOpts = slog.HandlerOptions{Level: options.Level}

	var handler slog.Handler
	switch options.Mode {
	case ProductionMode:
		handler = slog.NewJSONHandler(os.Stdout, &handlerOpts)
	case DevelopmentMode:
		handler = slog.NewTextHandler(os.Stdout, &handlerOpts)
	}

	return slog.New(handler), nil
}
