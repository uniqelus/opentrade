package sdklog

import "log/slog"

type logOptions struct {
	Mode  LogMode
	Level slog.Level
}

func defaultLogOptions() *logOptions {
	return &logOptions{
		Mode:  DevelopmentMode,
		Level: slog.LevelDebug,
	}
}

type LogOption func(*logOptions) error

func WithLogMode(value string) LogOption {
	return func(lo *logOptions) error {
		mode, err := ParseLogMode(value)
		if err != nil {
			return err
		}

		lo.Mode = mode
		return nil
	}
}

func WithLogLevel(value string) LogOption {
	return func(lo *logOptions) error {
		level, err := ParseLogLevel(value)
		if err != nil {
			return err
		}

		lo.Level = level
		return nil
	}
}
